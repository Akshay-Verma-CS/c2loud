# c2loud - Cloud Agnostic Library

## Why c2loud?

In the rapidly evolving cloud environment, businesses and developers often find themselves locked into a specific cloud provider's ecosystem. `c2loud` aims to alleviate this issue by offering a unified, cloud-agnostic interface that abstracts away the complexities of individual cloud providers. This approach enables developers to integrate multiple cloud services into their applications seamlessly, fostering an environment of flexibility and efficiency.

## Problem It Solves

`c2loud` addresses several challenges:

1. **Cloud Lock-in**: Reduces dependency on a single cloud provider, making it easier to switch or integrate multiple services.
2. **Complexity**: Simplifies the interaction with various cloud APIs by providing a unified interface.
3. **Adaptability**: Allows easy adaptation to new or changing cloud services without significant refactoring.

## Flow and Structure

The `c2loud` library is structured around a central `CloudProvider` interface, which defines common methods applicable across various cloud services. Specific providers like AWS and GCP implement this interface, ensuring that the underlying implementation details are abstracted away.

### Structure Overview

- `cloud/`: The main directory containing all cloud-related code.
  - `aws/`: AWS-specific implementations.
  - `gcp/`: GCP-specific implementations.
  - `interface.go`: Defines the `CloudProvider` interface and related types.

### Flow

1. **Initialization**: Users initialize a cloud provider instance with necessary configurations.
2. **Usage**: Users interact with cloud services using the methods defined in `CloudProvider`.
3. **Extension**: New services or providers can be added by implementing additional methods or interfaces.

## Usage Example

### AWS

```go
// Initialize AWS Provider
provider, err := aws.GetInstance("awsConfigKey", awsConfig)
if err != nil {
    log.Fatal(err)
}

// Use the S3 service
buckets, err := provider.S3Service.ListBuckets()
// ... handle buckets ...
```

### GCP

```go
// Initialize GCP Provider
provider, err := gcp.GetInstance(ctx, "gcpConfigKey", gcpConfig)
if err != nil {
    log.Fatal(err)
}

// Use the Compute Engine service
instances, err := provider.GetComputeService().ListInstances(projectID)
// ... handle instances ...
```

## Creators

### Akshay Verma

![Akshay Verma]

Akshay is a backend enthusiast passionate about building scalable and cloud-agnostic solutions. With `c2loud`, he aims to simplify cloud interactions and promote a more efficient and flexible development ecosystem.

- [GitHub](https://github.com/Akshay-Verma-CS)
- [LinkedIn](https://www.linkedin.com/in/akshay-verma-44a299198/)

## Acknowledgments

Thank the contributors, mentors, or anyone else who helped.

---

Thank you for exploring `c2loud`. Your contributions and feedback are welcome!

```
