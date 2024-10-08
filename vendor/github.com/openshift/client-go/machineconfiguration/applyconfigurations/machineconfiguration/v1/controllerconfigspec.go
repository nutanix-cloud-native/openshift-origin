// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	configv1 "github.com/openshift/api/config/v1"
	machineconfigurationv1 "github.com/openshift/api/machineconfiguration/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// ControllerConfigSpecApplyConfiguration represents a declarative configuration of the ControllerConfigSpec type for use
// with apply.
type ControllerConfigSpecApplyConfiguration struct {
	ClusterDNSIP                   *string                                   `json:"clusterDNSIP,omitempty"`
	CloudProviderConfig            *string                                   `json:"cloudProviderConfig,omitempty"`
	Platform                       *string                                   `json:"platform,omitempty"`
	EtcdDiscoveryDomain            *string                                   `json:"etcdDiscoveryDomain,omitempty"`
	KubeAPIServerServingCAData     []byte                                    `json:"kubeAPIServerServingCAData,omitempty"`
	RootCAData                     []byte                                    `json:"rootCAData,omitempty"`
	CloudProviderCAData            []byte                                    `json:"cloudProviderCAData,omitempty"`
	AdditionalTrustBundle          []byte                                    `json:"additionalTrustBundle,omitempty"`
	ImageRegistryBundleUserData    []ImageRegistryBundleApplyConfiguration   `json:"imageRegistryBundleUserData,omitempty"`
	ImageRegistryBundleData        []ImageRegistryBundleApplyConfiguration   `json:"imageRegistryBundleData,omitempty"`
	PullSecret                     *corev1.ObjectReferenceApplyConfiguration `json:"pullSecret,omitempty"`
	InternalRegistryPullSecret     []byte                                    `json:"internalRegistryPullSecret,omitempty"`
	Images                         map[string]string                         `json:"images,omitempty"`
	BaseOSContainerImage           *string                                   `json:"baseOSContainerImage,omitempty"`
	BaseOSExtensionsContainerImage *string                                   `json:"baseOSExtensionsContainerImage,omitempty"`
	OSImageURL                     *string                                   `json:"osImageURL,omitempty"`
	ReleaseImage                   *string                                   `json:"releaseImage,omitempty"`
	Proxy                          *configv1.ProxyStatus                     `json:"proxy,omitempty"`
	Infra                          *configv1.Infrastructure                  `json:"infra,omitempty"`
	DNS                            *configv1.DNS                             `json:"dns,omitempty"`
	IPFamilies                     *machineconfigurationv1.IPFamiliesType    `json:"ipFamilies,omitempty"`
	NetworkType                    *string                                   `json:"networkType,omitempty"`
	Network                        *NetworkInfoApplyConfiguration            `json:"network,omitempty"`
}

// ControllerConfigSpecApplyConfiguration constructs a declarative configuration of the ControllerConfigSpec type for use with
// apply.
func ControllerConfigSpec() *ControllerConfigSpecApplyConfiguration {
	return &ControllerConfigSpecApplyConfiguration{}
}

// WithClusterDNSIP sets the ClusterDNSIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterDNSIP field is set to the value of the last call.
func (b *ControllerConfigSpecApplyConfiguration) WithClusterDNSIP(value string) *ControllerConfigSpecApplyConfiguration {
	b.ClusterDNSIP = &value
	return b
}

// WithCloudProviderConfig sets the CloudProviderConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CloudProviderConfig field is set to the value of the last call.
func (b *ControllerConfigSpecApplyConfiguration) WithCloudProviderConfig(value string) *ControllerConfigSpecApplyConfiguration {
	b.CloudProviderConfig = &value
	return b
}

// WithPlatform sets the Platform field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Platform field is set to the value of the last call.
func (b *ControllerConfigSpecApplyConfiguration) WithPlatform(value string) *ControllerConfigSpecApplyConfiguration {
	b.Platform = &value
	return b
}

// WithEtcdDiscoveryDomain sets the EtcdDiscoveryDomain field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EtcdDiscoveryDomain field is set to the value of the last call.
func (b *ControllerConfigSpecApplyConfiguration) WithEtcdDiscoveryDomain(value string) *ControllerConfigSpecApplyConfiguration {
	b.EtcdDiscoveryDomain = &value
	return b
}

// WithKubeAPIServerServingCAData adds the given value to the KubeAPIServerServingCAData field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the KubeAPIServerServingCAData field.
func (b *ControllerConfigSpecApplyConfiguration) WithKubeAPIServerServingCAData(values ...byte) *ControllerConfigSpecApplyConfiguration {
	for i := range values {
		b.KubeAPIServerServingCAData = append(b.KubeAPIServerServingCAData, values[i])
	}
	return b
}

// WithRootCAData adds the given value to the RootCAData field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RootCAData field.
func (b *ControllerConfigSpecApplyConfiguration) WithRootCAData(values ...byte) *ControllerConfigSpecApplyConfiguration {
	for i := range values {
		b.RootCAData = append(b.RootCAData, values[i])
	}
	return b
}

// WithCloudProviderCAData adds the given value to the CloudProviderCAData field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the CloudProviderCAData field.
func (b *ControllerConfigSpecApplyConfiguration) WithCloudProviderCAData(values ...byte) *ControllerConfigSpecApplyConfiguration {
	for i := range values {
		b.CloudProviderCAData = append(b.CloudProviderCAData, values[i])
	}
	return b
}

// WithAdditionalTrustBundle adds the given value to the AdditionalTrustBundle field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AdditionalTrustBundle field.
func (b *ControllerConfigSpecApplyConfiguration) WithAdditionalTrustBundle(values ...byte) *ControllerConfigSpecApplyConfiguration {
	for i := range values {
		b.AdditionalTrustBundle = append(b.AdditionalTrustBundle, values[i])
	}
	return b
}

// WithImageRegistryBundleUserData adds the given value to the ImageRegistryBundleUserData field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ImageRegistryBundleUserData field.
func (b *ControllerConfigSpecApplyConfiguration) WithImageRegistryBundleUserData(values ...*ImageRegistryBundleApplyConfiguration) *ControllerConfigSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithImageRegistryBundleUserData")
		}
		b.ImageRegistryBundleUserData = append(b.ImageRegistryBundleUserData, *values[i])
	}
	return b
}

// WithImageRegistryBundleData adds the given value to the ImageRegistryBundleData field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ImageRegistryBundleData field.
func (b *ControllerConfigSpecApplyConfiguration) WithImageRegistryBundleData(values ...*ImageRegistryBundleApplyConfiguration) *ControllerConfigSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithImageRegistryBundleData")
		}
		b.ImageRegistryBundleData = append(b.ImageRegistryBundleData, *values[i])
	}
	return b
}

// WithPullSecret sets the PullSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PullSecret field is set to the value of the last call.
func (b *ControllerConfigSpecApplyConfiguration) WithPullSecret(value *corev1.ObjectReferenceApplyConfiguration) *ControllerConfigSpecApplyConfiguration {
	b.PullSecret = value
	return b
}

// WithInternalRegistryPullSecret adds the given value to the InternalRegistryPullSecret field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the InternalRegistryPullSecret field.
func (b *ControllerConfigSpecApplyConfiguration) WithInternalRegistryPullSecret(values ...byte) *ControllerConfigSpecApplyConfiguration {
	for i := range values {
		b.InternalRegistryPullSecret = append(b.InternalRegistryPullSecret, values[i])
	}
	return b
}

// WithImages puts the entries into the Images field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Images field,
// overwriting an existing map entries in Images field with the same key.
func (b *ControllerConfigSpecApplyConfiguration) WithImages(entries map[string]string) *ControllerConfigSpecApplyConfiguration {
	if b.Images == nil && len(entries) > 0 {
		b.Images = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Images[k] = v
	}
	return b
}

// WithBaseOSContainerImage sets the BaseOSContainerImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BaseOSContainerImage field is set to the value of the last call.
func (b *ControllerConfigSpecApplyConfiguration) WithBaseOSContainerImage(value string) *ControllerConfigSpecApplyConfiguration {
	b.BaseOSContainerImage = &value
	return b
}

// WithBaseOSExtensionsContainerImage sets the BaseOSExtensionsContainerImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BaseOSExtensionsContainerImage field is set to the value of the last call.
func (b *ControllerConfigSpecApplyConfiguration) WithBaseOSExtensionsContainerImage(value string) *ControllerConfigSpecApplyConfiguration {
	b.BaseOSExtensionsContainerImage = &value
	return b
}

// WithOSImageURL sets the OSImageURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OSImageURL field is set to the value of the last call.
func (b *ControllerConfigSpecApplyConfiguration) WithOSImageURL(value string) *ControllerConfigSpecApplyConfiguration {
	b.OSImageURL = &value
	return b
}

// WithReleaseImage sets the ReleaseImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReleaseImage field is set to the value of the last call.
func (b *ControllerConfigSpecApplyConfiguration) WithReleaseImage(value string) *ControllerConfigSpecApplyConfiguration {
	b.ReleaseImage = &value
	return b
}

// WithProxy sets the Proxy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Proxy field is set to the value of the last call.
func (b *ControllerConfigSpecApplyConfiguration) WithProxy(value configv1.ProxyStatus) *ControllerConfigSpecApplyConfiguration {
	b.Proxy = &value
	return b
}

// WithInfra sets the Infra field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Infra field is set to the value of the last call.
func (b *ControllerConfigSpecApplyConfiguration) WithInfra(value configv1.Infrastructure) *ControllerConfigSpecApplyConfiguration {
	b.Infra = &value
	return b
}

// WithDNS sets the DNS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DNS field is set to the value of the last call.
func (b *ControllerConfigSpecApplyConfiguration) WithDNS(value configv1.DNS) *ControllerConfigSpecApplyConfiguration {
	b.DNS = &value
	return b
}

// WithIPFamilies sets the IPFamilies field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IPFamilies field is set to the value of the last call.
func (b *ControllerConfigSpecApplyConfiguration) WithIPFamilies(value machineconfigurationv1.IPFamiliesType) *ControllerConfigSpecApplyConfiguration {
	b.IPFamilies = &value
	return b
}

// WithNetworkType sets the NetworkType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkType field is set to the value of the last call.
func (b *ControllerConfigSpecApplyConfiguration) WithNetworkType(value string) *ControllerConfigSpecApplyConfiguration {
	b.NetworkType = &value
	return b
}

// WithNetwork sets the Network field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Network field is set to the value of the last call.
func (b *ControllerConfigSpecApplyConfiguration) WithNetwork(value *NetworkInfoApplyConfiguration) *ControllerConfigSpecApplyConfiguration {
	b.Network = value
	return b
}
