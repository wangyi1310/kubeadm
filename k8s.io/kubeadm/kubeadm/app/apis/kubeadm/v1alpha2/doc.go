/*
Copyright 2018 The Kubernetes Authors.

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

// Package v1alpha2 is the package that contains the libraries that drive the kubeadm binary.
// +k8s:defaulter-gen=TypeMeta
// +groupName=kubeadm.k8s.io
// +k8s:deepcopy-gen=package
// +k8s:conversion-gen=k8s.io/kubeadm/cmd/kubeadm/app/apis/kubeadm
package v1alpha2 // import "k8s.io/kubeadm/cmd/kubeadm/app/apis/kubeadm/v1alpha2"
