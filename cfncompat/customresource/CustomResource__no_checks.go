// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package customresource

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomResource) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateImportFromParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateMoveToIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateCustomResource_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateCustomResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCustomResource_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateCustomResource_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomResource) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomResource) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomResource) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_CustomResource) validateSetLogicalResourceIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomResource) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomResource) validateSetResourcePropertiesParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomResource) validateSetResourceTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomResource) validateSetResponseBucketParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomResource) validateSetResponseKeyPrefixParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomResource) validateSetServiceTimeoutParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_CustomResource) validateSetServiceTokenParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomResource) validateSetStackIdParameters(val *string) error {
	return nil
}

func validateNewCustomResourceParameters(scope constructs.Construct, id *string, config *CustomResourceConfig) error {
	return nil
}

