// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"

	applyconfigurationapiv1alpha1 "github.com/kgateway-dev/kgateway/v2/api/applyconfiguration/api/v1alpha1"
	apiv1alpha1 "github.com/kgateway-dev/kgateway/v2/api/v1alpha1"
	scheme "github.com/kgateway-dev/kgateway/v2/pkg/client/clientset/versioned/scheme"
)

// DirectResponsesGetter has a method to return a DirectResponseInterface.
// A group's client should implement this interface.
type DirectResponsesGetter interface {
	DirectResponses(namespace string) DirectResponseInterface
}

// DirectResponseInterface has methods to work with DirectResponse resources.
type DirectResponseInterface interface {
	Create(ctx context.Context, directResponse *apiv1alpha1.DirectResponse, opts v1.CreateOptions) (*apiv1alpha1.DirectResponse, error)
	Update(ctx context.Context, directResponse *apiv1alpha1.DirectResponse, opts v1.UpdateOptions) (*apiv1alpha1.DirectResponse, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, directResponse *apiv1alpha1.DirectResponse, opts v1.UpdateOptions) (*apiv1alpha1.DirectResponse, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*apiv1alpha1.DirectResponse, error)
	List(ctx context.Context, opts v1.ListOptions) (*apiv1alpha1.DirectResponseList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apiv1alpha1.DirectResponse, err error)
	Apply(ctx context.Context, directResponse *applyconfigurationapiv1alpha1.DirectResponseApplyConfiguration, opts v1.ApplyOptions) (result *apiv1alpha1.DirectResponse, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, directResponse *applyconfigurationapiv1alpha1.DirectResponseApplyConfiguration, opts v1.ApplyOptions) (result *apiv1alpha1.DirectResponse, err error)
	DirectResponseExpansion
}

// directResponses implements DirectResponseInterface
type directResponses struct {
	*gentype.ClientWithListAndApply[*apiv1alpha1.DirectResponse, *apiv1alpha1.DirectResponseList, *applyconfigurationapiv1alpha1.DirectResponseApplyConfiguration]
}

// newDirectResponses returns a DirectResponses
func newDirectResponses(c *GatewayV1alpha1Client, namespace string) *directResponses {
	return &directResponses{
		gentype.NewClientWithListAndApply[*apiv1alpha1.DirectResponse, *apiv1alpha1.DirectResponseList, *applyconfigurationapiv1alpha1.DirectResponseApplyConfiguration](
			"directresponses",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *apiv1alpha1.DirectResponse { return &apiv1alpha1.DirectResponse{} },
			func() *apiv1alpha1.DirectResponseList { return &apiv1alpha1.DirectResponseList{} },
		),
	}
}
