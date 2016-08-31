/*
Copyright 2016 The Kubernetes Authors.

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
package kubemaster

import (
	"bytes"
	"fmt"
	"os"
	"path"

	kubeadmapi "k8s.io/kubernetes/pkg/kubeadm/api"
	kubeadmutil "k8s.io/kubernetes/pkg/kubeadm/util"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/util/uuid"
)

func generateTokenIfNeeded(params *kubeadmapi.BootstrapParams) error {
	ok, err := kubeadmutil.UseGivenTokenIfValid(params)
	if !ok {
		if err != nil {
			return err
		}
		return kubeadmutil.GenerateToken(params)
	}

	return nil
}

func CreateTokenAuthFile(params *kubeadmapi.BootstrapParams) error {
	if err := generateTokenIfNeeded(params); err != nil {
		return err
	}
	if err := os.MkdirAll(path.Join(params.EnvParams["host_pki_path"]), 0700); err != nil {
		return err
	}
	// <random-token>,<username>,<uid>,system:kubelet-bootstrap
	serialized := fmt.Sprintf("%s,kubeadm-node-csr,%s,system:kubelet-bootstrap\n", params.Discovery.BearerToken, uuid.NewUUID())
	if err := cmdutil.DumpReaderToFile(bytes.NewReader([]byte(serialized)), path.Join(params.EnvParams["host_pki_path"], "tokens.csv")); err != nil {
		return err
	}
	return nil
}
