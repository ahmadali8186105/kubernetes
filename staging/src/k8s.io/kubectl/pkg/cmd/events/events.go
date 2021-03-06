/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package events

import (
	"context"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/duration"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/client-go/kubernetes"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

const (
	eventsUsageStr = "events"
)

var (
	eventsLong = templates.LongDesc(i18n.T(`
		Experimental: Display events

		Prints a table of the most important information about events.
		You can request events for a namespace, for all namespace, or
		filtered to only those pertaining to a specified resource.`))

	eventsExample = templates.Examples(i18n.T(`
	# List recent events in the default namespace.
	kubectl alpha events

	# List recent events in all namespaces.
	kubectl alpha events --all-namespaces

	# List recent events, then wait for more events and list them as they arrive.
	kubectl alpha events --watch`))
)

type EventsOptions struct {
	AllNamespaces bool
	Namespace     string
	Watch         bool
	ForObject     string

	forResource string // decoded version of ForObject
	forName     string

	ctx    context.Context
	client *kubernetes.Clientset

	genericclioptions.IOStreams
}

func NewEventsOptions(streams genericclioptions.IOStreams, watch bool) *EventsOptions {
	return &EventsOptions{
		IOStreams: streams,
		Watch:     watch,
	}
}

// NewCmdEvents creates a new pod logs command
func NewCmdEvents(f cmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	o := NewEventsOptions(streams, false)

	cmd := &cobra.Command{
		Use:                   eventsUsageStr,
		DisableFlagsInUseLine: true,
		Short:                 i18n.T("Experimental: List events"),
		Long:                  eventsLong,
		Example:               eventsExample,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(o.Complete(f, cmd, args))
			cmdutil.CheckErr(o.Validate())
			cmdutil.CheckErr(o.Run())
		},
	}
	o.AddFlags(cmd)
	return cmd
}

func (o *EventsOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&o.Watch, "watch", "w", o.Watch, "After listing the requested events, watch for more events.")
	cmd.Flags().BoolVarP(&o.AllNamespaces, "all-namespaces", "A", o.AllNamespaces, "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.")
	cmd.Flags().StringVar(&o.ForObject, "for", o.ForObject, "Filter events to only those pertaining to the specified resource.")
}

func (o *EventsOptions) Complete(f cmdutil.Factory, cmd *cobra.Command, args []string) error {
	var err error
	o.Namespace, _, err = f.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return err
	}

	if o.ForObject != "" {
		tuple, found, err := splitResourceTypeName(o.ForObject)
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("--watch-for must be in resource/name form")
		}
		o.forName, o.forResource = tuple.Name, tuple.Resource
	}

	o.ctx = cmd.Context()
	o.client, err = f.KubernetesClientSet()
	if err != nil {
		return err
	}

	return nil
}

func (o EventsOptions) Validate() error {
	// Nothing to do here.
	return nil
}

// Run retrieves events
func (o EventsOptions) Run() error {
	namespace := o.Namespace
	if o.AllNamespaces {
		namespace = ""
	}
	listOptions := metav1.ListOptions{}
	if o.forName != "" {
		listOptions.FieldSelector = fields.AndSelectors(
			fields.OneTermEqualSelector("involvedObject.kind", o.forResource),
			fields.OneTermEqualSelector("involvedObject.name", o.forName)).String()
	}
	el, err := o.client.CoreV1().Events(namespace).List(o.ctx, listOptions)

	if err != nil {
		return err
	}

	w := printers.GetNewTabWriter(o.Out)

	sort.Sort(SortableEvents(el.Items))

	printHeadings(w, o.AllNamespaces)
	for _, e := range el.Items {
		printOneEvent(w, e, o.AllNamespaces)
	}
	w.Flush()
	if o.Watch {
		if len(el.Items) > 0 {
			// start watch from after the most recent item in the list
			listOptions.ResourceVersion = el.Items[len(el.Items)-1].ResourceVersion
		}
		eventWatch, err := o.client.CoreV1().Events(namespace).Watch(o.ctx, listOptions)
		if err != nil {
			return err
		}

		for e := range eventWatch.ResultChan() {
			if e.Type == watch.Deleted { // events are deleted after 1 hour; don't print that
				continue
			}
			event := e.Object.(*corev1.Event)
			printOneEvent(w, *event, o.AllNamespaces)
			w.Flush()
		}
	}

	return nil
}

func printHeadings(w io.Writer, allNamespaces bool) {
	if allNamespaces {
		fmt.Fprintf(w, "NAMESPACE\t")
	}
	fmt.Fprintf(w, "LAST SEEN\tTYPE\tREASON\tOBJECT\tMESSAGE\n")
}

func printOneEvent(w io.Writer, e corev1.Event, allNamespaces bool) {
	var interval string
	if e.Count > 1 {
		interval = fmt.Sprintf("%s (x%d over %s)", translateTimestampSince(e.LastTimestamp.Time), e.Count, translateTimestampSince(e.FirstTimestamp.Time))
	} else {
		interval = translateTimestampSince(eventTime(e))
	}
	if allNamespaces {
		fmt.Fprintf(w, "%v\t", e.Namespace)
	}
	fmt.Fprintf(w, "%s\t%s\t%s\t%s/%s\t%v\n",
		interval,
		e.Type,
		e.Reason,
		e.InvolvedObject.Kind, e.InvolvedObject.Name,
		strings.TrimSpace(e.Message),
	)
}

// SortableEvents implements sort.Interface for []api.Event by time
type SortableEvents []corev1.Event

func (list SortableEvents) Len() int {
	return len(list)
}

func (list SortableEvents) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list SortableEvents) Less(i, j int) bool {
	return eventTime(list[i]).Before(eventTime(list[j]))
}

// Some events have just an EventTime; if LastTimestamp is present we prefer that.
func eventTime(event corev1.Event) time.Time {
	if !event.LastTimestamp.Time.IsZero() {
		return event.LastTimestamp.Time
	}
	return event.EventTime.Time
}

// translateTimestampSince returns the elapsed time since timestamp in
// human-readable approximation.
func translateTimestampSince(timestamp time.Time) string {
	if timestamp.IsZero() {
		return "<unknown>"
	}

	return duration.HumanDuration(time.Since(timestamp))
}
