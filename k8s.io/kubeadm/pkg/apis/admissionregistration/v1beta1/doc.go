/*
Copyright 2017 The Kubernetes Authors.

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

// +k8s:conversion-gen=k8s.io/kubeadm/pkg/apis/admissionregistration
// +k8s:conversion-gen-external-types=k8s.io/api/admissionregistration/v1beta1
// +k8s:defaulter-gen=TypeMeta
// +k8s:defaulter-gen-input=../../../../vendor/k8s.io/api/admissionregistration/v1beta1

// Package v1beta1 is the v1beta1 version of the API.
// AdmissionConfiguration and AdmissionPluginConfiguration are legacy static admission plugin configuration
// InitializerConfiguration, ValidatingWebhookConfiguration, and MutatingWebhookConfiguration are for the
// new dynamic admission controller configuration.
// +groupName=admissionregistration.k8s.io
package v1beta1 // import "k8s.io/kubeadm/pkg/apis/admissionregistration/v1beta1"
