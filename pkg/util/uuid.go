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

package util

import (
	"time"

	"github.com/pborman/uuid"
	"k8s.io/kubernetes/pkg/types"
)

/**
 * The UUID package is naive and can generate identical UUIDs if the time interval is quick enough.
 * Block subsequent UUIDs for 200 Nanoseconds, the UUID uses 100 ns increments, we block for 200 to be safe
 * Blocks in a go routine, so that the caller doesn't have to wait.
 */

// Burst buffer size 100
var uuidChan = make(chan types.UID, 100)

func init() {
	go Forever(func() {
		result := uuid.NewUUID()
		uuidChan <- types.UID(result.String())
	}, 200*time.Nanosecond)
}

func NewUUID() types.UID {
	return <-uuidChan
}
