/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

// A set of common functions needed by cmd/kubectl and pkg/kubectl packages.
package kubectl

import (
	"fmt"
	"strings"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/meta"
)

const kubectlAnnotationPrefix = "kubectl.kubernetes.io/"

type NamespaceInfo struct {
	Namespace string
}

func listOfImages(spec *api.PodSpec) []string {
	var images []string
	for _, container := range spec.Containers {
		images = append(images, container.Image)
	}
	return images
}

func makeImageList(spec *api.PodSpec) string {
	return strings.Join(listOfImages(spec), ",")
}

// OutputVersionMapper is a RESTMapper that will prefer mappings that
// correspond to a preferred output version (if feasible)
type OutputVersionMapper struct {
	meta.RESTMapper
	OutputVersion string
}

// RESTMapping implements meta.RESTMapper by prepending the output version to the preferred version list.
func (m OutputVersionMapper) RESTMapping(kind string, versions ...string) (*meta.RESTMapping, error) {
	preferred := []string{m.OutputVersion}
	for _, version := range versions {
		if len(version) > 0 {
			preferred = append(preferred, version)
		}
	}
	// if the caller wants to use the default version list, try with the preferred version, and on
	// error, use the default behavior.
	if len(preferred) == 1 {
		if m, err := m.RESTMapper.RESTMapping(kind, preferred...); err == nil {
			return m, nil
		}
		preferred = nil
	}
	return m.RESTMapper.RESTMapping(kind, preferred...)
}

// GroupExpander handles resource values that include a qualified group (jobs.extensions) and
// verifies the specified group is empty or matches the resource group.
// TODO: allow resolution of resources that exist in multiple groups by knowing about all
// possible groups and checking them in order
type GroupExpander struct {
	meta.RESTMapper
}

func (e GroupExpander) VersionAndKindForResource(resource string) (defaultVersion, kind string, err error) {
	if !strings.Contains(resource, ".") {
		return e.RESTMapper.VersionAndKindForResource(resource)
	}
	parts := strings.SplitN(resource, ".", 2)
	resource, expectGroup := parts[0], parts[1]
	group, err := e.RESTMapper.GroupForResource(resource)
	if err != nil {
		return "", "", err
	}
	if group != expectGroup {
		return "", "", fmt.Errorf("resource %q is part of group %q, not %q", resource, group, expectGroup)
	}
	return e.RESTMapper.VersionAndKindForResource(resource)
}

// ShortcutExpander is a RESTMapper that can be used for Kubernetes
// resources that expands shortened resources. It cannot handle fully
// qualified resource values.
type ShortcutExpander struct {
	meta.RESTMapper
}

// VersionAndKindForResource implements meta.RESTMapper. It expands the resource first, then invokes the wrapped
// mapper.
func (e ShortcutExpander) VersionAndKindForResource(resource string) (defaultVersion, kind string, err error) {
	resource = expandResourceShortcut(resource)
	defaultVersion, kind, err = e.RESTMapper.VersionAndKindForResource(resource)
	return defaultVersion, kind, err
}

// ResourceIsValid takes a string and checks if it's a valid resource.
// It expands the resource first, then invokes the wrapped mapper.
func (e ShortcutExpander) ResourceIsValid(resource string) bool {
	return e.RESTMapper.ResourceIsValid(expandResourceShortcut(resource))
}

// GroupForResource expands the resource and then calls the wrapped mapper.
func (e ShortcutExpander) GroupForResource(resource string) (string, error) {
	return e.RESTMapper.GroupForResource(expandResourceShortcut(resource))
}

// AliasesForResource expands the resource and then calls the wrapped mapper.
func (e ShortcutExpander) AliasesForResource(resource string) ([]string, bool) {
	return e.RESTMapper.AliasesForResource(expandResourceShortcut(resource))
}

// ResourceSingularizer expands the resource and then calls the wrapped mapper.
func (e ShortcutExpander) ResourceSingularizer(resource string) (singular string, err error) {
	return e.RESTMapper.ResourceSingularizer(expandResourceShortcut(resource))
}

// expandResourceShortcut will return the expanded version of resource
// (something that a pkg/api/meta.RESTMapper can understand), if it is
// indeed a shortcut. Otherwise, will return resource unmodified.
func expandResourceShortcut(resource string) string {
	shortForms := map[string]string{
		// Please keep this alphabetized
		"cs":     "componentstatuses",
		"ev":     "events",
		"ep":     "endpoints",
		"hpa":    "horizontalpodautoscalers",
		"limits": "limitranges",
		"no":     "nodes",
		"ns":     "namespaces",
		"po":     "pods",
		"pv":     "persistentvolumes",
		"pvc":    "persistentvolumeclaims",
		"quota":  "resourcequotas",
		"rc":     "replicationcontrollers",
		"ds":     "daemonsets",
		"svc":    "services",
		"ing":    "ingresses",
	}
	if expanded, ok := shortForms[resource]; ok {
		return expanded
	}
	return resource
}
