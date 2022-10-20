/*
Copyright 2022 The KubeBlocks Authors

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	DoNotTerminate TerminationPolicyType = "DoNotTerminate"
	Halt           TerminationPolicyType = "Halt"
	Delete         TerminationPolicyType = "Delete"
	WipeOut        TerminationPolicyType = "WipeOut"
)

type TerminationPolicyType string

type PodAntiAffinity string

const (
	Preferred PodAntiAffinity = "Preferred"
	Required  PodAntiAffinity = "Required"
)

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	// ref ClusterDefinition, immutable
	// +kubebuilder:validation:Required
	ClusterDefRef string `json:"clusterDefinitionRef"`

	// ref AppVersion
	// +kubebuilder:validation:Required
	AppVersionRef string `json:"appVersionRef"`

	// +optional
	Components []ClusterComponent `json:"components,omitempty"`

	// Affinity describes affinities which specific by users
	// +optional
	Affinity *Affinity `json:"affinity,omitempty"`

	// One of DoNotTerminate, Halt, Delete, WipeOut.
	// Defaults to Halt.
	// DoNotTerminate means block delete operation.
	// Halt means delete resources such as sts,deploy,svc,pdb, but keep pvcs.
	// Delete is based on Halt and delete pvcs.
	// WipeOut is based on Delete and wipe out all snapshots and snapshot data from bucket.
	// +kubebuilder:default=Halt
	// +kubebuilder:validation:Enum={DoNotTerminate,Halt,Delete,WipeOut}
	TerminationPolicy TerminationPolicyType `json:"terminationPolicy,omitempty"`
}

// ClusterStatus defines the observed state of Cluster
type ClusterStatus struct {
	// observedGeneration is the most recent generation observed for this
	// Cluster. It corresponds to the Cluster's generation, which is
	// updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// phase - in list of [Running, Failed, Creating, Updating, Deleting, Deleted]
	// +kubebuilder:validation:Enum={Running,Failed,Creating,Updating,Deleting,Deleted}
	Phase Phase `json:"phase,omitempty"`

	// Message cluster details message in current phase
	// +optional
	Message string `json:"message,omitempty"`

	// Components record the current status information of all components of the cluster
	// +optional
	Components map[string]*ClusterStatusComponent `json:"components,omitempty"`

	// Operations declares which operations the cluster supports
	// +optional
	Operations Operations `json:"operations,omitempty"`

	ClusterDefinitionStatusGeneration `json:",inline"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:categories={dbaas,all}
//+kubebuilder:printcolumn:name="APP-VERSION",type="string",JSONPath=".spec.appVersionRef",description="Cluster Application Version."
//+kubebuilder:printcolumn:name="PHASE",type="string",JSONPath=".status.phase",description="Cluster Status."
//+kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"

// Cluster is the Schema for the clusters API
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterList contains a list of Cluster
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

type ClusterComponent struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=12
	Name string `json:"name"`

	// component name in ClusterDefinition
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=12
	Type string `json:"type"`

	// default value in ClusterDefinition
	Replicas int `json:"replicas,omitempty"`

	// Affinity describes affinities which specific by users
	// +optional
	Affinity *Affinity `json:"affinity,omitempty"`

	// Monitor which is a switch to enable monitoring, default is false
	// DBaas provides an extension mechanism to support component level monitoring,
	// which will scrape metrics auto or manually from servers in component and export
	// metrics to Time Series Database.
	// +kubebuilder:validation:Required
	// +kubebuilder:default=false
	Monitor bool `json:"monitor"`

	// +optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// +optional
	VolumeClaimTemplates []ClusterComponentVolumeClaimTemplate `json:"volumeClaimTemplates,omitempty"`

	// serviceType determines how the Service is exposed. Valid
	// options are ClusterIP, NodePort, and LoadBalancer.
	// "ClusterIP" allocates a cluster-internal IP address for load-balancing
	// to endpoints. Endpoints are determined by the selector or if that is not
	// specified, by manual construction of an Endpoints object or
	// EndpointSlice objects. If clusterIP is "None", no virtual IP is
	// allocated and the endpoints are published as a set of endpoints rather
	// than a virtual IP.
	// "NodePort" builds on ClusterIP and allocates a port on every node which
	// routes to the same endpoints as the clusterIP.
	// "LoadBalancer" builds on NodePort and creates an external load-balancer
	// (if supported in the current cloud) which routes to the same endpoints
	// as the clusterIP.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
	// +kubebuilder:default=ClusterIP
	// +kubebuilder:validation:Enum={ClusterIP,NodePort,LoadBalancer}
	// +optional
	ServiceType corev1.ServiceType `json:"serviceType,omitempty"`
}

// ClusterStatusComponent record components status information
type ClusterStatusComponent struct {
	// Type component type
	// +optional
	Type string `json:"type,omitempty"`

	// Phase - in list of [Running, Failed, Creating, Updating, Deleting, Deleted]
	// +kubebuilder:validation:Enum={Running,Failed,Creating,Updating,Deleting,Deleted}
	Phase Phase `json:"phase,omitempty"`

	// Message record the component details message in current phase
	// +optional
	Message string `json:"message,omitempty"`
}

type ClusterComponentVolumeClaimTemplate struct {
	Name string                           `json:"name"`
	Spec corev1.PersistentVolumeClaimSpec `json:"spec,omitempty"`
}

type Affinity struct {
	// PodAntiAffinity defines pods of component anti-affnity
	// Defaults to Preferred
	// Preferred means try spread pods by topologyKey
	// Required means must spread pods by topologyKey
	// +kubebuilder:validation:Enum={Preferred,Required}
	// +optional
	PodAntiAffinity PodAntiAffinity `json:"podAntiAffinity,omitempty"`
	// TopologyKeys describe topologyKeys for `topologySpreadConstraint` and `podAntiAffinity` in ClusterDefinition API
	// +optional
	TopologyKeys []string `json:"topologyKeys"`
	// NodeLabels describe constrain which nodes pod can be scheduled on based on node labels
	// +optional
	NodeLabels map[string]string `json:"nodeLabels,omitempty"`
}

type Operations struct {
	// Upgradable whether the cluster supports upgrade. if multiple appVersions existed, it is true
	// +optional
	Upgradable bool `json:"upgradable,omitempty"`

	// VerticalScalable which components of the cluster support verticalScaling.
	// +optional
	VerticalScalable []string `json:"verticalScalable,omitempty"`

	// Restartable which components of the cluster support restart.
	// +optional
	Restartable []string `json:"restartable,omitempty"`

	// VolumeExpandable which components of the cluster and its volumeClaimTemplates support volumeExpansion.
	// +optional
	VolumeExpandable []*OperationComponent `json:"volumeExpandable,omitempty"`

	// HorizontalScalable which components of the cluster support horizontalScaling, and the replicas range limit.
	// +optional
	HorizontalScalable []*OperationComponent `json:"horizontalScalable,omitempty"`
}

type OperationComponent struct {
	// Name component name
	// +kubebuilder:validation:Required
	Name string `json:"name,omitempty"`

	// Min minimum of replicas when operation is horizontalScaling,
	// +optional
	Min int `json:"min,omitempty"`

	// Max maximum of replicas when operation is horizontalScaling
	// +optional
	Max int `json:"max,omitempty"`

	// VolumeClaimTemplateNames which VolumeClaimTemplate of the component support volumeExpansion
	// +optional
	VolumeClaimTemplateNames []string `json:"volumeClaimTemplateNames,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
