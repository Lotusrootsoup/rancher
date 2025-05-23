package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Plugin Name",type=string,JSONPath=`.spec.plugin.name`
// +kubebuilder:printcolumn:name="Version",type=string,JSONPath=`.spec.plugin.version`
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.cacheState`
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UIPlugin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Spec is the desired state of the ui plugin.
	Spec UIPluginSpec `json:"spec"`
	// Status is the observed state of the ui plugin.
	// +kubebuilder:validation:Optional
	Status UIPluginStatus `json:"status"`
}

type UIPluginSpec struct {
	Plugin UIPluginEntry `json:"plugin,omitempty"`
}

// UIPluginEntry represents an ui plugin.
type UIPluginEntry struct {
	// Name of the plugin.
	Name string `json:"name,omitempty"`
	// Version of the plugin.
	Version string `json:"version,omitempty"`
	// Endpoint from where to fetch the contents of the plugin.
	Endpoint string `json:"endpoint,omitempty"`
	// CompressedEndpoint link to a targz file that contains the content of the plugin.
	CompressedEndpoint string `json:"compressedEndpoint,omitempty"`
	// NoCache a flag that tells if the plugin should be cached or not.
	// Defaults to false.
	// +kubebuilder:default:=false
	NoCache bool `json:"noCache,omitempty"`
	// NoAuth a flag that tells if the plugin should be accessible without authentication.
	// Defaults to false.
	// +kubebuilder:default:=false
	NoAuth bool `json:"noAuth,omitempty"`
	// Metadata of the plugin.
	Metadata map[string]string `json:"metadata,omitempty"`
}

type UIPluginStatus struct {
	// ObservedGeneration is used by Rancher controller to track the latest generation of the resource that it triggered on.
	ObservedGeneration int64 `json:"observedGeneration"`
	// CacheState is the cache status of the plugin.
	// +nullable
	CacheState string `json:"cacheState,omitempty"`
	// Error is the error message if any.
	Error string `json:"error,omitempty"`
	// Ready is the readiness of the plugin.
	// Defaults to false.
	// +kubebuilder:default:=false
	Ready bool `json:"ready,omitempty"`
	// RetryNumber is the number of times the plugin has been retried.
	RetryNumber int `json:"retryNumber,omitempty"`
	// RetryAt is the time at which the plugin should be retried.
	RetryAt metav1.Time `json:"retryAt,omitempty"`
}
