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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindKubevirtInfraCsiDriver = "KubevirtInfraCsiDriver"
	ResourceKubevirtInfraCsiDriver     = "kubevirtinfracsidriver"
	ResourceKubevirtInfraCsiDrivers    = "kubevirtinfracsidrivers"
)

// KubevirtInfraCsiDriver defines the schama for KubevirtInfraCsiDriver Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=clusterauths,singular=clusterauth,categories={kubeops,appscode}
type KubevirtInfraCsiDriver struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubevirtInfraCsiDriverSpec `json:"spec,omitempty"`
}

// KubevirtInfraCsiDriverSpec is the schema for KubevirtInfraCsiDriver Operator values file
type KubevirtInfraCsiDriverSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string    `json:"fullnameOverride"`
	ReplicaCount     int32     `json:"replicaCount"`
	Driver           Container `json:"driver"`
	Provisioner      Container `json:"provisioner"`
	Attacher         Container `json:"attacher"`
	Livenessprobe    Container `json:"livenessprobe"`
	Snapshotter      Container `json:"snapshotter"`
	Resizer          Container `json:"resizer"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	PodLabels map[string]string `json:"podLabels"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector" protobuf:"bytes,12,rep,name=nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations" protobuf:"bytes,13,rep,name=tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity" protobuf:"bytes,14,opt,name=affinity"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount"`
	PriorityClassName  string                   `json:"priorityClassName"`
	Tenant             InfraTenant              `json:"tenant"`
}

type InfraTenant struct {
	Namespace               string               `json:"namespace"`
	Labels                  string               `json:"labels"`
	Kubeconfig              string               `json:"kubeconfig"`
	StorageClassEnforcement apiextensionsv1.JSON `json:"storageClassEnforcement"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubevirtInfraCsiDriverList is a list of KubevirtInfraCsiDrivers
type KubevirtInfraCsiDriverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubevirtInfraCsiDriver CRD objects
	Items []KubevirtInfraCsiDriver `json:"items,omitempty"`
}
