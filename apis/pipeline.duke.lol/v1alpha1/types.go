package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Pipeline is a duked resource defining a CI lifecycle
type Pipeline struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the ddesired behaviour of the pod terminator.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec PipelineSpec `json:"spec"`
}

// PipelineSpec is the spec for a Duked resource
type PipelineSpec struct {
	// Selector is how the target will be selected
	selector map[string]string `json:"selector,omitempty"`
	// sourceRepository is the location of a Deployments code
	sourceRepository string `json:"sourceRepository,omitempty"`
	// dockerfile specifies which dockerfile to use from sourceRepository
	dockerfile string `json:"dockerfile,omitempty"`
	ssh        SSHSettings
	steps      Steps
	notify     Notify
	when       RunWhen
	worker     Worker
}

// SSHSettings defines operations for using SSH keys with duke
type SSHSettings struct {
	// hostPath specifies the location of the .ssh/ directory
	hostPath string `json:"hostPath,omitempty"`
}

// Steps defines pipeline steps (pipeline execution)
type Steps struct {
	// name is the label for a single pipeline step
	name string `json:"name,omitempty"`
	// image is the docker builder image to run steps in
	image string `jsonn:"image,omitempty"`
	// commands is an array of strings defining a custom command to execute
	commands map[string]string `json:"commands,omitempty"`
	// env defines environment variables ingested by the build executor
	env     map[string]string `json:"env,omitempty"`
	secrets Secrets           `json:"secrets,omitempty"`
}

// Secrets defines the location, kind, and type of a kubernetes secret object
type Secrets struct {
	// name is the label for a single secret
	name string `json:"name,omitempty"`
	// TODO
}

// Notify
type Notify struct {
	when  NotifyWhen  `json:"when,omitempty"`
	where NotifyWhere `json:"where,omitempty"`
}

// NotifyWhen defines conditionals for notifications
type NotifyWhen struct {
	// event holds all possible event-based conditionals i.e. "on_success"
	event map[string]string `json:"event,omitempty"`
}

// NotifyWhere defines the external services to POST to for notifications
// New providers can be added via satisfying notify.Interface
type NotifyWhere struct {
	// slack integration
	slack NotifyWhereSlack `json:"slack,omitempty"`
	// email integration
	email NotifyWhereEmail `json:"email,omitempty"`
}

// NotifyWhereSlack defines notifications via slack
type NotifyWhereSlack struct {
	// channel is a list of slack channels
	channel map[string]string `json:"channel,omitempty"`
	// TODO allow fromSecret
	// token is the workspace-specific slack token for authentication
	token string `json:"token,omitempty"`
}

// NotifyWhereEmail defines notifications via email
type NotifyWhereEmail struct {
	// from_address is the outgoing email address
	from_address string `json:"from_address,omitempty"`
	// to_address is a list of addresses to notify
	to_address map[string]string `json:"to_address,omitempty"`
}

// RunWhen defines what events trigger the operator to execute the pipeline
type RunWhen struct {
	// event holds all possible event-based conditionals i.e. "on_commit"
	event map[string]string `json:"event,omitempty"`
	// branch defines a git branch conditional
	branch string `json:"branch,omitempty"`
}

// Worker defines the configuration for a pipeline executor
type Worker struct {
	// replicas controllers how many executors the operator will run at one given time
	replicas int32 `json:"replicas,omitempty"`
	// resources define k8 style cpu and memory limits & requests for executors
	resources Resources `json:"resources,omitempty"`
}

// Resources sets the limits and requests for a container
type Resources struct {
	Requests CPUAndMem `json:"requests,omitempty"`
	Limits   CPUAndMem `json:"limits,omitempty"`
}

// CPUAndMem defines how many cpu and ram the container will request/limit
type CPUAndMem struct {
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

// Status has the status of the cluster
type PipelineStatus struct {
	Phase      Phase       `json:"phase"`
	Conditions []Condition `json:"conditions"`
}

// Phase of the pipeline status
type Phase string

// Condition saves the state information of the pipeline
type Condition struct {
	Type           ConditionType `json:"type"`
	Reason         string        `json:"reason"`
	TransitionTime string        `json:"transitionTime"`
}

// ConditionType defines the condition that the pipelinne can have
type ConditionType string

// Pipeline
type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Pipeline `json:"items"`
}
