// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ClusterValidationID cluster validation id
//
// swagger:model cluster-validation-id
type ClusterValidationID string

func NewClusterValidationID(value ClusterValidationID) *ClusterValidationID {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ClusterValidationID.
func (m ClusterValidationID) Pointer() *ClusterValidationID {
	return &m
}

const (

	// ClusterValidationIDMachineCidrDefined captures enum value "machine-cidr-defined"
	ClusterValidationIDMachineCidrDefined ClusterValidationID = "machine-cidr-defined"

	// ClusterValidationIDClusterCidrDefined captures enum value "cluster-cidr-defined"
	ClusterValidationIDClusterCidrDefined ClusterValidationID = "cluster-cidr-defined"

	// ClusterValidationIDServiceCidrDefined captures enum value "service-cidr-defined"
	ClusterValidationIDServiceCidrDefined ClusterValidationID = "service-cidr-defined"

	// ClusterValidationIDNoCidrsOverlapping captures enum value "no-cidrs-overlapping"
	ClusterValidationIDNoCidrsOverlapping ClusterValidationID = "no-cidrs-overlapping"

	// ClusterValidationIDNetworksSameAddressFamilies captures enum value "networks-same-address-families"
	ClusterValidationIDNetworksSameAddressFamilies ClusterValidationID = "networks-same-address-families"

	// ClusterValidationIDNetworkPrefixValid captures enum value "network-prefix-valid"
	ClusterValidationIDNetworkPrefixValid ClusterValidationID = "network-prefix-valid"

	// ClusterValidationIDMachineCidrEqualsToCalculatedCidr captures enum value "machine-cidr-equals-to-calculated-cidr"
	ClusterValidationIDMachineCidrEqualsToCalculatedCidr ClusterValidationID = "machine-cidr-equals-to-calculated-cidr"

	// ClusterValidationIDAPIVipsDefined captures enum value "api-vips-defined"
	ClusterValidationIDAPIVipsDefined ClusterValidationID = "api-vips-defined"

	// ClusterValidationIDAPIVipsValid captures enum value "api-vips-valid"
	ClusterValidationIDAPIVipsValid ClusterValidationID = "api-vips-valid"

	// ClusterValidationIDIngressVipsDefined captures enum value "ingress-vips-defined"
	ClusterValidationIDIngressVipsDefined ClusterValidationID = "ingress-vips-defined"

	// ClusterValidationIDIngressVipsValid captures enum value "ingress-vips-valid"
	ClusterValidationIDIngressVipsValid ClusterValidationID = "ingress-vips-valid"

	// ClusterValidationIDAllHostsAreReadyToInstall captures enum value "all-hosts-are-ready-to-install"
	ClusterValidationIDAllHostsAreReadyToInstall ClusterValidationID = "all-hosts-are-ready-to-install"

	// ClusterValidationIDSufficientMastersCount captures enum value "sufficient-masters-count"
	ClusterValidationIDSufficientMastersCount ClusterValidationID = "sufficient-masters-count"

	// ClusterValidationIDDNSDomainDefined captures enum value "dns-domain-defined"
	ClusterValidationIDDNSDomainDefined ClusterValidationID = "dns-domain-defined"

	// ClusterValidationIDPullSecretSet captures enum value "pull-secret-set"
	ClusterValidationIDPullSecretSet ClusterValidationID = "pull-secret-set"

	// ClusterValidationIDNtpServerConfigured captures enum value "ntp-server-configured"
	ClusterValidationIDNtpServerConfigured ClusterValidationID = "ntp-server-configured"

	// ClusterValidationIDLsoRequirementsSatisfied captures enum value "lso-requirements-satisfied"
	ClusterValidationIDLsoRequirementsSatisfied ClusterValidationID = "lso-requirements-satisfied"

	// ClusterValidationIDOcsRequirementsSatisfied captures enum value "ocs-requirements-satisfied"
	ClusterValidationIDOcsRequirementsSatisfied ClusterValidationID = "ocs-requirements-satisfied"

	// ClusterValidationIDOdfRequirementsSatisfied captures enum value "odf-requirements-satisfied"
	ClusterValidationIDOdfRequirementsSatisfied ClusterValidationID = "odf-requirements-satisfied"

	// ClusterValidationIDCnvRequirementsSatisfied captures enum value "cnv-requirements-satisfied"
	ClusterValidationIDCnvRequirementsSatisfied ClusterValidationID = "cnv-requirements-satisfied"

	// ClusterValidationIDLvmRequirementsSatisfied captures enum value "lvm-requirements-satisfied"
	ClusterValidationIDLvmRequirementsSatisfied ClusterValidationID = "lvm-requirements-satisfied"

	// ClusterValidationIDMceRequirementsSatisfied captures enum value "mce-requirements-satisfied"
	ClusterValidationIDMceRequirementsSatisfied ClusterValidationID = "mce-requirements-satisfied"

	// ClusterValidationIDMtvRequirementsSatisfied captures enum value "mtv-requirements-satisfied"
	ClusterValidationIDMtvRequirementsSatisfied ClusterValidationID = "mtv-requirements-satisfied"

	// ClusterValidationIDOscRequirementsSatisfied captures enum value "osc-requirements-satisfied"
	ClusterValidationIDOscRequirementsSatisfied ClusterValidationID = "osc-requirements-satisfied"

	// ClusterValidationIDNetworkTypeValid captures enum value "network-type-valid"
	ClusterValidationIDNetworkTypeValid ClusterValidationID = "network-type-valid"

	// ClusterValidationIDPlatformRequirementsSatisfied captures enum value "platform-requirements-satisfied"
	ClusterValidationIDPlatformRequirementsSatisfied ClusterValidationID = "platform-requirements-satisfied"

	// ClusterValidationIDNodeFeatureDiscoveryRequirementsSatisfied captures enum value "node-feature-discovery-requirements-satisfied"
	ClusterValidationIDNodeFeatureDiscoveryRequirementsSatisfied ClusterValidationID = "node-feature-discovery-requirements-satisfied"

	// ClusterValidationIDNvidiaGpuRequirementsSatisfied captures enum value "nvidia-gpu-requirements-satisfied"
	ClusterValidationIDNvidiaGpuRequirementsSatisfied ClusterValidationID = "nvidia-gpu-requirements-satisfied"

	// ClusterValidationIDPipelinesRequirementsSatisfied captures enum value "pipelines-requirements-satisfied"
	ClusterValidationIDPipelinesRequirementsSatisfied ClusterValidationID = "pipelines-requirements-satisfied"

	// ClusterValidationIDServicemeshRequirementsSatisfied captures enum value "servicemesh-requirements-satisfied"
	ClusterValidationIDServicemeshRequirementsSatisfied ClusterValidationID = "servicemesh-requirements-satisfied"

	// ClusterValidationIDServerlessRequirementsSatisfied captures enum value "serverless-requirements-satisfied"
	ClusterValidationIDServerlessRequirementsSatisfied ClusterValidationID = "serverless-requirements-satisfied"

	// ClusterValidationIDOpenshiftAiRequirementsSatisfied captures enum value "openshift-ai-requirements-satisfied"
	ClusterValidationIDOpenshiftAiRequirementsSatisfied ClusterValidationID = "openshift-ai-requirements-satisfied"

	// ClusterValidationIDAuthorinoRequirementsSatisfied captures enum value "authorino-requirements-satisfied"
	ClusterValidationIDAuthorinoRequirementsSatisfied ClusterValidationID = "authorino-requirements-satisfied"

	// ClusterValidationIDNmstateRequirementsSatisfied captures enum value "nmstate-requirements-satisfied"
	ClusterValidationIDNmstateRequirementsSatisfied ClusterValidationID = "nmstate-requirements-satisfied"

	// ClusterValidationIDAmdGpuRequirementsSatisfied captures enum value "amd-gpu-requirements-satisfied"
	ClusterValidationIDAmdGpuRequirementsSatisfied ClusterValidationID = "amd-gpu-requirements-satisfied"

	// ClusterValidationIDKmmRequirementsSatisfied captures enum value "kmm-requirements-satisfied"
	ClusterValidationIDKmmRequirementsSatisfied ClusterValidationID = "kmm-requirements-satisfied"

	// ClusterValidationIDNodeHealthcheckRequirementsSatisfied captures enum value "node-healthcheck-requirements-satisfied"
	ClusterValidationIDNodeHealthcheckRequirementsSatisfied ClusterValidationID = "node-healthcheck-requirements-satisfied"

	// ClusterValidationIDSelfNodeRemediationRequirementsSatisfied captures enum value "self-node-remediation-requirements-satisfied"
	ClusterValidationIDSelfNodeRemediationRequirementsSatisfied ClusterValidationID = "self-node-remediation-requirements-satisfied"

	// ClusterValidationIDFenceAgentsRemediationRequirementsSatisfied captures enum value "fence-agents-remediation-requirements-satisfied"
	ClusterValidationIDFenceAgentsRemediationRequirementsSatisfied ClusterValidationID = "fence-agents-remediation-requirements-satisfied"
)

// for schema
var clusterValidationIdEnum []interface{}

func init() {
	var res []ClusterValidationID
	if err := json.Unmarshal([]byte(`["machine-cidr-defined","cluster-cidr-defined","service-cidr-defined","no-cidrs-overlapping","networks-same-address-families","network-prefix-valid","machine-cidr-equals-to-calculated-cidr","api-vips-defined","api-vips-valid","ingress-vips-defined","ingress-vips-valid","all-hosts-are-ready-to-install","sufficient-masters-count","dns-domain-defined","pull-secret-set","ntp-server-configured","lso-requirements-satisfied","ocs-requirements-satisfied","odf-requirements-satisfied","cnv-requirements-satisfied","lvm-requirements-satisfied","mce-requirements-satisfied","mtv-requirements-satisfied","osc-requirements-satisfied","network-type-valid","platform-requirements-satisfied","node-feature-discovery-requirements-satisfied","nvidia-gpu-requirements-satisfied","pipelines-requirements-satisfied","servicemesh-requirements-satisfied","serverless-requirements-satisfied","openshift-ai-requirements-satisfied","authorino-requirements-satisfied","nmstate-requirements-satisfied","amd-gpu-requirements-satisfied","kmm-requirements-satisfied","node-healthcheck-requirements-satisfied","self-node-remediation-requirements-satisfied","fence-agents-remediation-requirements-satisfied"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterValidationIdEnum = append(clusterValidationIdEnum, v)
	}
}

func (m ClusterValidationID) validateClusterValidationIDEnum(path, location string, value ClusterValidationID) error {
	if err := validate.EnumCase(path, location, value, clusterValidationIdEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cluster validation id
func (m ClusterValidationID) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClusterValidationIDEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cluster validation id based on context it is used
func (m ClusterValidationID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
