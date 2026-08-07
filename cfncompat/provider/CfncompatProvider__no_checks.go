// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfncompatProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfncompatProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfncompatProvider) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateCfncompatProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateCfncompatProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCfncompatProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateCfncompatProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CfncompatProvider) validateSetAssumeRoleParameters(val *CfncompatProviderAssumeRole) error {
	return nil
}

func (j *jsiiProxy_CfncompatProvider) validateSetAssumeRoleWithWebIdentityParameters(val *CfncompatProviderAssumeRoleWithWebIdentity) error {
	return nil
}

func (j *jsiiProxy_CfncompatProvider) validateSetEndpointsParameters(val *CfncompatProviderEndpoints) error {
	return nil
}

func (j *jsiiProxy_CfncompatProvider) validateSetInsecureParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CfncompatProvider) validateSetSkipMetadataApiCheckParameters(val interface{}) error {
	return nil
}

func validateNewCfncompatProviderParameters(scope constructs.Construct, id *string, config *CfncompatProviderConfig) error {
	return nil
}

