// Package cloud defines interfaces and structures for interacting with various cloud providers.
package cloud

// CloudProvider is an interface that all cloud provider implementations must satisfy.
// It defines common methods that are agnostic of the specific provider.
type CloudProvider interface {
    // CreateResource creates a new resource in the cloud.
    CreateResource(name string, options ResourceOptions) (Resource, error)

    // DeleteResource deletes a resource from the cloud.
    DeleteResource(id string) error

    // ListResources lists all resources managed by the cloud provider.
    ListResources() ([]Resource, error)

    // GetResource fetches details of a specific resource.
    GetResource(id string) (Resource, error)
}

// Resource represents a general cloud resource.
type Resource struct {
    ID      string
    Name    string
    Type    string
    Status  string
    Details map[string]interface{} // Additional details specific to the resource type
}

// ResourceOptions represents options for creating a cloud resource.
type ResourceOptions struct {
    Type    string
    Config  map[string]interface{} // Configuration specific to the resource type
}

// NewProviderFunc is a function type that creates a new instance of a CloudProvider.
// Each cloud provider package should implement this function.
type NewProviderFunc func(config map[string]interface{}) (CloudProvider, error)
