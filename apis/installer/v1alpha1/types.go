/*
Copyright AppsCode Inc. and Contributors

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

package v1alpha1

import (
	core "k8s.io/api/core/v1"
)

type ImageRef struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

type Container struct {
	ImageRef   `json:",inline"`
	PullPolicy string `json:"pullPolicy"`
	// Compute Resources required by the sidecar container.
	// +optional
	Resources core.ResourceRequirements `json:"resources"`
	// Security options the pod should run with.
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
}

type ServiceAccountSpec struct {
	Create    bool `json:"create"`
	Automount bool `json:"automount"`
	//+optional
	Name *string `json:"name"`
	//+optional
	Annotations map[string]string `json:"annotations"`
}
