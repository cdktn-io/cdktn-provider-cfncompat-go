// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package providerfunctions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (c *jsiiProxy_CfncompatProviderFunctions) validateBase64Parameters(content *string) error {
	if content == nil {
		return fmt.Errorf("parameter content is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateCidrParameters(ipBlock *string, count *float64, cidrBits *float64) error {
	if ipBlock == nil {
		return fmt.Errorf("parameter ipBlock is required, but nil was provided")
	}

	if count == nil {
		return fmt.Errorf("parameter count is required, but nil was provided")
	}

	if cidrBits == nil {
		return fmt.Errorf("parameter cidrBits is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateConditionAndParameters(conditions *[]interface{}) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}
	for idx_562a77, v := range *conditions {
		switch v.(type) {
		case *bool:
			// ok
		case bool:
			// ok
		case cdktn.IResolvable:
			// ok
		default:
			if !_jsii_.IsAnonymousProxy(v) {
				return fmt.Errorf("parameter conditions[%#v] must be one of the allowed types: *bool, cdktn.IResolvable; received %#v (a %T)", idx_562a77, v, v)
			}
		}
	}

	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateConditionContainsParameters(listOfStrings *[]*string, value *string) error {
	if listOfStrings == nil {
		return fmt.Errorf("parameter listOfStrings is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateConditionEachMemberEqualsParameters(listOfStrings *[]*string, value *string) error {
	if listOfStrings == nil {
		return fmt.Errorf("parameter listOfStrings is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateConditionEachMemberInParameters(stringsToCheck *[]*string, stringsToMatch *[]*string) error {
	if stringsToCheck == nil {
		return fmt.Errorf("parameter stringsToCheck is required, but nil was provided")
	}

	if stringsToMatch == nil {
		return fmt.Errorf("parameter stringsToMatch is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateConditionEqualsParameters(value1 interface{}, value2 interface{}) error {
	if value1 == nil {
		return fmt.Errorf("parameter value1 is required, but nil was provided")
	}

	if value2 == nil {
		return fmt.Errorf("parameter value2 is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateConditionIfParameters(condition interface{}) error {
	if condition == nil {
		return fmt.Errorf("parameter condition is required, but nil was provided")
	}
	switch condition.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case cdktn.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(condition) {
			return fmt.Errorf("parameter condition must be one of the allowed types: *bool, cdktn.IResolvable; received %#v (a %T)", condition, condition)
		}
	}

	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateConditionNotParameters(condition interface{}) error {
	if condition == nil {
		return fmt.Errorf("parameter condition is required, but nil was provided")
	}
	switch condition.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case cdktn.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(condition) {
			return fmt.Errorf("parameter condition must be one of the allowed types: *bool, cdktn.IResolvable; received %#v (a %T)", condition, condition)
		}
	}

	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateConditionOrParameters(conditions *[]interface{}) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}
	for idx_562a77, v := range *conditions {
		switch v.(type) {
		case *bool:
			// ok
		case bool:
			// ok
		case cdktn.IResolvable:
			// ok
		default:
			if !_jsii_.IsAnonymousProxy(v) {
				return fmt.Errorf("parameter conditions[%#v] must be one of the allowed types: *bool, cdktn.IResolvable; received %#v (a %T)", idx_562a77, v, v)
			}
		}
	}

	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateFindInMapParameters(mapping interface{}, topLevelKey *string, secondLevelKey *string, defaultValue *[]interface{}) error {
	if mapping == nil {
		return fmt.Errorf("parameter mapping is required, but nil was provided")
	}

	if topLevelKey == nil {
		return fmt.Errorf("parameter topLevelKey is required, but nil was provided")
	}

	if secondLevelKey == nil {
		return fmt.Errorf("parameter secondLevelKey is required, but nil was provided")
	}

	if defaultValue == nil {
		return fmt.Errorf("parameter defaultValue is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateJoinParameters(delimiter *string, values *[]*string) error {
	if delimiter == nil {
		return fmt.Errorf("parameter delimiter is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateLengthOfParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateSelectParameters(index *float64, objects interface{}) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	if objects == nil {
		return fmt.Errorf("parameter objects is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateSplitParameters(delimiter *string, source *string) error {
	if delimiter == nil {
		return fmt.Errorf("parameter delimiter is required, but nil was provided")
	}

	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateSubParameters(template *string, variables *map[string]*string) error {
	if template == nil {
		return fmt.Errorf("parameter template is required, but nil was provided")
	}

	if variables == nil {
		return fmt.Errorf("parameter variables is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfncompatProviderFunctions) validateToJsonStringParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateNewCfncompatProviderFunctionsParameters(providerLocalName *string) error {
	if providerLocalName == nil {
		return fmt.Errorf("parameter providerLocalName is required, but nil was provided")
	}

	return nil
}

