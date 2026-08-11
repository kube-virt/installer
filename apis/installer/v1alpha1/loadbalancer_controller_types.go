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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindLoadbalancerController = "LoadbalancerController"
	ResourceLoadbalancerController     = "loadbalancercontroller"
	ResourceLoadbalancerControllers    = "loadbalancercontrollers"
)

// LoadbalancerController defines the schama for LoadbalancerController Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=clusterauths,singular=clusterauth,categories={kubeops,appscode}
type LoadbalancerController struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadbalancerControllerSpec `json:"spec,omitempty"`
}

// LoadbalancerControllerSpec is the schema for LoadbalancerController Operator values file
type LoadbalancerControllerSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string `json:"fullnameOverride"`

	// Mode is either "openstack" or "harvester". In harvester mode, a kubeconfig
	// secret is mounted; openstack-creds are NOT used.
	Mode string `json:"mode"`

	LoadbalancerController LoadbalancerControllerDeploymentSpec `json:"loadbalancerController"`
	OpenstackCreds         OpenstackCredsSpec                   `json:"openstackCreds"`
	DBCredentials          DBCredentialsSpec                    `json:"dbCredentials"`
	Harvester              HarvesterSpec                        `json:"harvester"`
	KeepalivedGrpc         KeepalivedGrpcSpec                   `json:"keepalivedGrpc"`
}

// ImagePullSpec is a container image reference along with its pull policy.
type ImagePullSpec struct {
	ImageRef   `json:",inline"`
	PullPolicy string `json:"pullPolicy"`
}

// LoadbalancerControllerDeploymentSpec is the schema for the loadbalancer-controller Deployment.
type LoadbalancerControllerDeploymentSpec struct {
	ReplicaCount       int32                     `json:"replicaCount"`
	Image              ImagePullSpec             `json:"image"`
	ServiceAccountName string                    `json:"serviceAccountName"`
	Env                LoadbalancerControllerEnv `json:"env"`
	//+optional
	Resources core.ResourceRequirements `json:"resources"`
}

// LoadbalancerControllerEnv holds environment configuration for the loadbalancer-controller container.
type LoadbalancerControllerEnv struct {
	//+optional
	OsTenantId string `json:"osTenantId"`
	//+optional
	OsNetworkId string `json:"osNetworkId"`
	//+optional
	OsSubnetId       string `json:"osSubnetId"`
	ClusterNamespace string `json:"clusterNamespace"`
}

// OpenstackCredsSpec is the schema for the openstack-creds Secret (used only when mode=openstack).
type OpenstackCredsSpec struct {
	//+optional
	OsAuthUrl  string `json:"osAuthUrl"`
	OsUsername string `json:"osUsername"`
	//+optional
	OsPassword   string `json:"osPassword"`
	OsDomainName string `json:"osDomainName"`
	OsRegionName string `json:"osRegionName"`
	//+optional
	OsProjectId string `json:"osProjectId"`
}

// DBCredentialsSpec is the schema for the db-credentials Secret.
type DBCredentialsSpec struct {
	//+optional
	DBHost     string `json:"dbHost"`
	DBPort     string `json:"dbPort"`
	DBUser     string `json:"dbUser"`
	DBPassword string `json:"dbPassword"`
	DBName     string `json:"dbName"`
	DBSslmode  string `json:"dbSslmode"`
}

// HarvesterSpec is the schema for the harvester kubeconfig (used only when mode=harvester).
// The kubeconfig content is stored in a Secret and mounted into the controller pod
// at /etc/harvester/kubeconfig.
type HarvesterSpec struct {
	KubeconfigSecretName string `json:"kubeconfigSecretName"`
	// Paste the full kubeconfig as a multiline string value, or leave empty and
	// create the secret manually.
	//+optional
	Kubeconfig string `json:"kubeconfig"`
}

// KeepalivedGrpcSpec is the schema for the keepalived-grpc DaemonSet.
type KeepalivedGrpcSpec struct {
	Image       ImagePullSpec `json:"image"`
	HostNetwork bool          `json:"hostNetwork"`
	HostPID     bool          `json:"hostPID"`
	//+optional
	Resources core.ResourceRequirements `json:"resources"`
	// KeepalivedConfigPath is the hostPath for the keepalived config directory.
	KeepalivedConfigPath string `json:"keepalivedConfigPath"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LoadbalancerControllerList is a list of LoadbalancerControllers
type LoadbalancerControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LoadbalancerController CRD objects
	Items []LoadbalancerController `json:"items,omitempty"`
}
