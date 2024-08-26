package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DirectResponseRoute contains configuration for defining direct response routes.
//
// +kubebuilder:object:root=true
// +kubebuilder:metadata:labels={app=gloo-gateway,app.kubernetes.io/name=gloo-gateway}
// +kubebuilder:resource:categories=gloo-gateway,shortName=drr
// +kubebuilder:subresource:status
type DirectResponseRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DirectResponseRouteSpec   `json:"spec,omitempty"`
	Status DirectResponseRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type DirectResponseRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DirectResponseRoute `json:"items"`
}

// DirectResponseRouteSpec describes the desired state of a DirectResponseRoute.
//
// +kubebuilder:validation:XValidation:message="The 'body' field is required when 'code' is a 2xx status code",rule="self.code < 200 || self.code >= 300 || has(self.body)"
type DirectResponseRouteSpec struct {
	// Code defines the HTTP status code to return for this route.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=100
	// +kubebuilder:validation:Maximum=599
	Code *int32 `json:"code"`
	// Body defines the content to be returned in the HTTP response body.
	//
	// +kubebuilder:validation:Optional
	//
	// TODO(tim): Make required? Add validation on length of body?
	Body *string `json:"body,omitempty"`
}

// DirectResponseRouteStatus defines the observed state of a DirectResponseRoute.
type DirectResponseRouteStatus struct {
	// Define observed state fields here. For example, you might include a LastUpdated timestamp.
	// +kubebuilder:validation:Optional
	//
	// TODO(tim): Do we need to add any status fields? If so, investigate how
	// other APIs define status fields with the same pattern.
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// GetCode returns the HTTP status code to return for this route.
func (in *DirectResponseRoute) GetCode() *int32 {
	if in == nil {
		return nil
	}
	return in.Spec.Code
}

// GetBody returns the content to be returned in the HTTP response body.
func (in *DirectResponseRoute) GetBody() *string {
	if in == nil {
		return nil
	}
	return in.Spec.Body
}

// TODO(tim): This is normally scaffolded with kubebuilder, but I didn't see it present
// in the GatewayParameters Go types. Investigate if it's necessary.
func init() {
	SchemeBuilder.Register(&DirectResponseRoute{}, &DirectResponseRouteList{})
}
