// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package providerfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfncompatProviderFunctions) validateBase64Parameters(content *string) error {
	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateCidrParameters(ipBlock *string, count *float64, cidrBits *float64) error {
	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateConditionAndParameters(conditions *[]interface{}) error {
	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateConditionContainsParameters(listOfStrings *[]*string, value *string) error {
	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateConditionEachMemberEqualsParameters(listOfStrings *[]*string, value *string) error {
	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateConditionEachMemberInParameters(stringsToCheck *[]*string, stringsToMatch *[]*string) error {
	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateConditionEqualsParameters(value1 interface{}, value2 interface{}) error {
	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateConditionIfParameters(condition interface{}) error {
	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateConditionNotParameters(condition interface{}) error {
	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateConditionOrParameters(conditions *[]interface{}) error {
	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateFindInMapParameters(mapping interface{}, topLevelKey *string, secondLevelKey *string, defaultValue *[]interface{}) error {
	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateJoinParameters(delimiter *string, values *[]*string) error {
	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateLengthOfParameters(value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateSelectParameters(index *float64, objects interface{}) error {
	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateSplitParameters(delimiter *string, source *string) error {
	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateSubParameters(template *string, variables *map[string]*string) error {
	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateToJsonStringParameters(value interface{}) error {
	return nil
}

func validateNewCfncompatProviderFunctionsParameters(providerLocalName *string) error {
	return nil
}

