package validation

import (
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/kubeedge/edgemesh/pkg/apis/config/v1alpha1"
	"github.com/kubeedge/kubeedge/pkg/apis/componentconfig/cloudcore/v1alpha1/validation"
)

func ValidateEdgeMeshAgentConfiguration(c *v1alpha1.EdgeMeshAgentConfig) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, ValidateKubeAPIConfig(c.KubeAPIConfig)...)
	allErrs = append(allErrs, ValidateModuleEdgeTunnel(c.Modules.EdgeTunnelConfig)...)
	return allErrs
}

func ValidateEdgeMeshGatewayConfiguration(c *v1alpha1.EdgeMeshGatewayConfig) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, ValidateKubeAPIConfig(c.KubeAPIConfig)...)
	allErrs = append(allErrs, ValidateModuleEdgeTunnel(c.Modules.EdgeTunnelConfig)...)
	return allErrs
}

func ValidateKubeAPIConfig(c *v1alpha1.KubeAPIConfig) field.ErrorList {
	allErrs := field.ErrorList{}
	validation.ValidateKubeAPIConfig(c.KubeAPIConfig)
	// TODO validate metaServerAddress
	return allErrs
}

func ValidateModuleEdgeTunnel(c *v1alpha1.EdgeTunnelConfig) field.ErrorList {
	if !c.Enable {
		return field.ErrorList{}
	}

	allErrs := field.ErrorList{}
	validTransport := IsValidTransport(c.Transport)

	if len(validTransport) > 0 {
		for _, m := range validTransport {
			allErrs = append(allErrs, field.Invalid(field.NewPath("Transport"), c.Transport, m))
		}
	}

	return allErrs
}

func IsValidTransport(transport string) []string {
	var supportedTransports = []string{"tcp", "ws", "quic"}
	for _, tr := range supportedTransports {
		if transport == tr {
			return nil
		}
	}
	return supportedTransports
}
