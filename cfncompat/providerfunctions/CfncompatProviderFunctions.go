// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package providerfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cfncompat-go/cfncompat/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Provider-defined functions of the cfncompat provider.
type CfncompatProviderFunctions interface {
	// Returns the Base64 representation of the input string, matching the semantics of CloudFormation's [`Fn::Base64`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-base64.html) intrinsic function. This is typically used to pass encoded data (e.g. EC2 `UserData`) through Terraform configuration that mirrors a CloudFormation template.
	Base64(content *string) *string
	// Returns an array of `count` CIDR address blocks, split from `ip_block`.
	//
	// Matches the semantics of CloudFormation's [`Fn::Cidr`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-cidr.html) intrinsic function.
	//
	// `cidr_bits` is the number of subnet bits for the generated CIDRs: the resulting prefix length is `32 - cidr_bits` for an IPv4 `ip_block`, or `128 - cidr_bits` for an IPv6 `ip_block`. For example, `cidr_bits = 8` against an IPv4 block produces `/24` subnets. `count` must be between 1 and 256. Subnets are allocated consecutively starting from the beginning of `ip_block`, and `ip_block` must be large enough to contain `count` subnets of the requested size. Both IPv4 and IPv6 blocks are supported.
	Cidr(ipBlock *string, count *float64, cidrBits *float64) *[]*string
	// Returns `true` if all the specified `conditions` evaluate to `true`, or returns `false` if any one of the conditions evaluates to `false`.
	//
	// Acts as an AND operator.
	//
	// Matches the semantics of CloudFormation's [`Fn::And`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-conditions.html#intrinsic-function-reference-conditions-and) intrinsic function (CDK binding: `Fn.conditionAnd`).
	//
	// CloudFormation's `Fn::And` takes a list of condition expressions (references to named `Conditions`, or nested condition functions) that CloudFormation evaluates lazily. This provider function instead takes already-evaluated boolean values as its `conditions` — the synthesis backend is responsible for evaluating any nested condition expressions to booleans before calling this function.
	//
	// The minimum number of `conditions` you can include is 2, and the maximum is 10.
	ConditionAnd(conditions *[]interface{}) cdktn.IResolvable
	// Returns `true` if `value` matches at least one member of `list_of_strings`.
	//
	// Matches the semantics of CloudFormation's [`Fn::Contains`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-rules.html#fn-contains) function. CloudFormation classifies `Fn::Contains` as a [rule-specific intrinsic function](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/rules-section-structure.html#template-constraint-rules-syntax): in a template it is only valid within the `RuleCondition` or `Assert` fields of the template's `Rules` section, not as a general-purpose intrinsic used in `Resources` or `Outputs`. As a `provider::cfncompat::*` function it has no such placement restriction and can be called anywhere a boolean value is needed.
	//
	// `list_of_strings` must not be null and must not contain null entries.
	ConditionContains(listOfStrings *[]*string, value *string) cdktn.IResolvable
	// Returns `true` if every member of `list_of_strings` equals `value`, mirroring CloudFormation's rule-specific [`Fn::EachMemberEquals`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-rules.html#fn-eachmemberequals) intrinsic function.
	//
	// AWS's documentation does not specify the behavior when `list_of_strings` is empty. This provider chooses vacuous truth: an empty list returns `true`, since "every member equals value" is trivially satisfied when there are no members.
	//
	// In CloudFormation, `Fn::EachMemberEquals` is a Rules-section-only function typically used with `Fn::ValueOfAll` to validate that a tag or attribute is consistent across every AWS-specific parameter of a given type; this provider-function form operates directly on an already-resolved list of strings rather than on template Rules-section constructs like `Fn::ValueOfAll`.
	ConditionEachMemberEquals(listOfStrings *[]*string, value *string) cdktn.IResolvable
	// Returns `true` if each member in `strings_to_check` is present in `strings_to_match`, and `false` otherwise.
	//
	// Matches the semantics of CloudFormation's [`Fn::EachMemberIn`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-rules.html#fn-eachmemberin) rule-specific intrinsic function.
	//
	// **Empty `strings_to_check`:** the AWS documentation does not specify a result for an empty `strings_to_check` list. This implementation returns `true` (vacuous truth: there are no members of `strings_to_check` that fail to appear in `strings_to_match`), matching how universally-quantified checks over an empty set are conventionally defined.
	//
	// **Deviation from CloudFormation:** in a CloudFormation template, `Fn::EachMemberIn` is normally used together with `Fn::ValueOfAll` and `Fn::RefAll` to compare attribute values across every parameter of a given AWS-specific parameter type. Those two functions depend on template-wide `Parameters` declarations and stack-account/Region context that this provider-defined function form has no access to; callers must resolve `strings_to_check` and `strings_to_match` to concrete lists of strings themselves before calling this function.
	ConditionEachMemberIn(stringsToCheck *[]*string, stringsToMatch *[]*string) cdktn.IResolvable
	// Compares if two values are equal. Returns `true` if the two values are equal, or `false` if they aren't.
	//
	// Matches the semantics of CloudFormation's [`Fn::Equals`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-conditions.html#intrinsic-function-reference-conditions-equals) intrinsic function.
	//
	// **Coercion rule:** CloudFormation template values are fundamentally strings, so `value_1` and `value_2` are compared by their canonical string representation rather than by native type: a string compares as-is, a number compares by its canonical decimal representation (so `1`, `1.0`, and `"1"` are all equal, and `-0` is equal to `0`), and a boolean compares as the literal `true` or `false`. `value_1` and `value_2` must each be a string, number, or boolean, and must not be null; non-primitive arguments (list, set, map/object, or tuple) and null arguments produce an error.
	ConditionEquals(value1 interface{}, value2 interface{}) cdktn.IResolvable
	// Returns `value_if_true` if `condition` evaluates to `true`, or `value_if_false` if `condition` evaluates to `false`.
	//
	// Matches the semantics of CloudFormation's [`Fn::If`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-conditions.html#intrinsic-function-reference-conditions-if) intrinsic function.
	//
	// **Deviation from `Fn::If`:** in a CloudFormation template, `Fn::If`'s first argument is the *name* of a condition declared in the template's `Conditions` section (for example `!If [CreateNewSecurityGroup, ..., ...]`), and CloudFormation resolves that name to a boolean by evaluating the referenced condition at stack create/update time. This provider function has no template `Conditions` section to resolve names against, so `condition` here is the already-evaluated **boolean value** of that condition — the CDK Terrain synthesis backend is responsible for evaluating the named condition (typically itself built from `provider::cfncompat::condition_*` function calls) and passing the resulting boolean through.
	//
	// `value_if_true` and `value_if_false` accept any type and are not required to share a type with each other, since only one of them is ever returned; the return type is dynamic to accommodate this. Only the selected branch is validated (e.g. for unknown-ness) — the branch that is not selected is ignored entirely, matching CloudFormation's lazy evaluation of `Fn::If` branches.
	ConditionIf(condition interface{}, valueIfTrue interface{}, valueIfFalse interface{}) cdktn.IResolvable
	// Returns `true` for a `condition` that evaluates to `false`, or `false` for a `condition` that evaluates to `true`.
	//
	// `Fn::Not` acts as a NOT operator.
	//
	// Matches the semantics of CloudFormation's [`Fn::Not`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-conditions.html#intrinsic-function-reference-conditions-not) intrinsic function.
	//
	// **Deviation from CloudFormation:** `Fn::Not` takes a single-element list containing the condition (`"Fn::Not": [{{condition}}]`). As a provider-defined function, which does not support single-element list arguments distinct from scalars, `condition_not` instead takes the condition directly as a single boolean argument.
	ConditionNot(condition interface{}) cdktn.IResolvable
	// Returns `true` if any one of the specified `conditions` evaluates to `true`, or returns `false` if all the conditions evaluate to `false`.
	//
	// Acts as an OR operator.
	//
	// Matches the semantics of CloudFormation's [`Fn::Or`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-conditions.html#intrinsic-function-reference-conditions-or) intrinsic function (CDK binding: `Fn.conditionOr`).
	//
	// CloudFormation's `Fn::Or` takes a list of condition expressions (references to named `Conditions`, or nested condition functions) that CloudFormation evaluates lazily. This provider function instead takes already-evaluated boolean values as its `conditions` - the synthesis backend is responsible for evaluating any nested condition expressions to booleans before calling this function.
	//
	// The minimum number of `conditions` you can include is 2, and the maximum is 10.
	ConditionOr(conditions *[]interface{}) cdktn.IResolvable
	// Returns the value corresponding to `top_level_key` and `second_level_key` in the two-level `mapping`.
	//
	// Matches the semantics of CloudFormation's [`Fn::FindInMap`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-findinmap.html) intrinsic function.
	//
	// **Deviation from template-form CloudFormation:** in a CloudFormation template, `Fn::FindInMap`'s first argument is the *logical name* of a mapping declared in the template's `Mappings` section. This provider function has no `Mappings` section to resolve a name against, so it takes the mapping *itself* as the first argument: `mapping` is a two-level object/map (`top_level_key` -> `second_level_key` -> leaf value), where each leaf value is either a string or a list of strings, matching the value shapes CloudFormation allows in a `Mappings` entry.
	//
	// If `top_level_key` or `second_level_key` is not present in `mapping`, this function raises an error naming the missing key and the level (top-level or second-level) it was expected at - unless a `default_value` argument is supplied, in which case that default is returned instead. This mirrors the optional `DefaultValue` parameter of CloudFormation's `Fn::FindInMap` (a standard parameter that no longer requires the `AWS::LanguageExtensions` transform). `default_value` is variadic and accepts zero or one arguments; supplying more than one is an error.
	FindInMap(mapping interface{}, topLevelKey *string, secondLevelKey *string, defaultValue *[]interface{}) cdktn.IResolvable
	// Appends a set of values into a single value, separated by the specified `delimiter`.
	//
	// If `delimiter` is the empty string, the values are concatenated with no delimiter.
	//
	// Matches the semantics of CloudFormation's [`Fn::Join`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-join.html) intrinsic function.
	//
	// The delimiter occurs between fragments only; it never terminates the final value. An empty `values` list returns the empty string. Every element of `values` must be a known, non-null string (CloudFormation's `Fn::Join` list of values must not contain null entries).
	Join(delimiter *string, values *[]*string) *string
	// Returns the number of elements within an array.
	//
	// Matches the semantics of CloudFormation's [`Fn::Length`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-length.html) intrinsic function, which **requires the `AWS::LanguageExtensions` transform** to be enabled on the CloudFormation template in order to be used. AWS CDK exposes this intrinsic as `Fn.len`; the generated CDKTN TypeScript binding for this provider function is `lengthOf` (`length` collides with the built-in array/string property name in TypeScript/JavaScript).
	//
	// `value` accepts a list, tuple, or set (CloudFormation's `Fn::Length` operates on the array produced by a `Ref`, a nested intrinsic function such as `Fn::Split`, or a literal array). Any other underlying type — a string, number, boolean, or map/object — is an error; the AWS documentation only documents array inputs for this function, so this provider does not attempt to guess a length for non-array values (e.g. treating a string's character count as its "length").
	LengthOf(value interface{}) *float64
	// Returns a single object from a list of objects by `index`.
	//
	// Matches the semantics of CloudFormation's [`Fn::Select`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-select.html) intrinsic function.
	//
	// `index` must be a value from zero to N-1, where N represents the number of elements in `objects`. `objects` must not be null, nor can it have null entries; CloudFormation documents both conditions as resulting in a stack error, so this function returns an error rather than a partial/best-effort result.
	//
	// `objects` accepts either a list or a tuple, so it can represent both CloudFormation's homogeneous lists (e.g. a resolved `CommaDelimitedList` parameter) and heterogeneous tuples produced by other `provider::cfncompat::*` functions. The return type is dynamic because tuple elements may be of any type.
	Select(index *float64, objects interface{}) cdktn.IResolvable
	// Splits `source` into a list of strings at each occurrence of `delimiter`, mirroring CloudFormation's [`Fn::Split`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-split.html) intrinsic function.
	//
	// Consecutive occurrences of the delimiter (or a delimiter at the start/end of `source`) produce empty string entries in the result, e.g. `split("|", "a||c|")` returns `["a", "", "c", ""]`.
	//
	// `delimiter` must be a non-empty string; CloudFormation does not document behavior for an empty delimiter, so this provider treats it as an error rather than guessing at undocumented Go-specific splitting behavior.
	Split(delimiter *string, source *string) *[]*string
	// Implements the [`Fn::Sub`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-sub.html) intrinsic function.
	//
	// Every `${VarName}` occurrence in `template` is replaced with `variables["VarName"]`. Write `${!VarName}` to produce the literal text `${VarName}` without performing a lookup (the `!` is stripped and the rest of the substitution is left untouched). Substituted values are **not** re-scanned for further `${...}` substitution (no recursion).
	//
	// **Deviation from CloudFormation**: in a CloudFormation template, `${Name}` may also refer to a template parameter name, a resource logical ID, or a resource attribute (`${MyInstance.PublicIp}`), which CloudFormation resolves the same way as `Ref`/`Fn::GetAtt`. This provider function has no access to a template's parameters/resources, so it only supports the explicit key-value `variables` map form of `Fn::Sub`. Any such name must be resolved by the calling synthesis backend (e.g. CDK Terrain) into an explicit entry of `variables` before calling this function — `${Name}` with no matching entry in `variables` is an error.
	Sub(template *string, variables *map[string]*string) *string
	// Implements CloudFormation's [`Fn::ToJsonString`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-ToJsonString.html) intrinsic function (`AWS::LanguageExtensions`) / AWS CDK's `Fn.toJsonString()`.
	//
	// **Deviation from CFN template syntax:** in a CloudFormation template, `Fn::ToJsonString` requires the `AWS::LanguageExtensions` transform to be declared, and its documented syntax only shows an `Object` or `Array` as the top-level argument. As a provider-defined function this is a pure serializer with no transform to declare, so `value` may be any dynamic value — object/map, list/tuple, or primitive (string/number/bool), including `null`, arbitrarily nested — matching AWS CDK's `Fn.toJsonString(value: any)`, which likewise accepts any value. When CDK synthesizes a CloudFormation template that calls `Fn.toJsonString()`, CDK automatically adds the `AWS::LanguageExtensions` transform to the template on your behalf; this function is the synth-time equivalent used before that template exists.
	//
	// `value` is serialized to compact JSON (no extraneous whitespace). Object/map keys are sorted lexicographically byte-wise so output is deterministic regardless of input key order (CloudFormation does not document a specific key order for `Fn::ToJsonString`, and Go maps are unordered, so this function imposes and documents sorted-key output). Numbers are rendered as canonical decimal literals (shortest round-tripping representation, no trailing zeros, no scientific notation) for typical values; extremely large or small magnitudes may still require many digits since JSON has no exponent-free size limit guarantee.
	//
	// Matches the official example: `{"key1": "value1", "key2": <resolved-value>}` serializes to `"{\"key1\":\"value1\",\"key2\":\"resolved-value\"}"`.
	ToJsonString(value interface{}) *string
}

// The jsii proxy struct for CfncompatProviderFunctions
type jsiiProxy_CfncompatProviderFunctions struct {
	_ byte // padding
}

func NewCfncompatProviderFunctions(providerLocalName *string) CfncompatProviderFunctions {
	_init_.Initialize()

	if err := validateNewCfncompatProviderFunctionsParameters(providerLocalName); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfncompatProviderFunctions{}

	_jsii_.Create(
		"@cdktn/provider-cfncompat.providerFunctions.CfncompatProviderFunctions",
		[]interface{}{providerLocalName},
		&j,
	)

	return &j
}

func NewCfncompatProviderFunctions_Override(c CfncompatProviderFunctions, providerLocalName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cfncompat.providerFunctions.CfncompatProviderFunctions",
		[]interface{}{providerLocalName},
		c,
	)
}

func (c *jsiiProxy_CfncompatProviderFunctions) Base64(content *string) *string {
	if err := c.validateBase64Parameters(content); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"base64",
		[]interface{}{content},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProviderFunctions) Cidr(ipBlock *string, count *float64, cidrBits *float64) *[]*string {
	if err := c.validateCidrParameters(ipBlock, count, cidrBits); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"cidr",
		[]interface{}{ipBlock, count, cidrBits},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProviderFunctions) ConditionAnd(conditions *[]interface{}) cdktn.IResolvable {
	if err := c.validateConditionAndParameters(conditions); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"conditionAnd",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProviderFunctions) ConditionContains(listOfStrings *[]*string, value *string) cdktn.IResolvable {
	if err := c.validateConditionContainsParameters(listOfStrings, value); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"conditionContains",
		[]interface{}{listOfStrings, value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProviderFunctions) ConditionEachMemberEquals(listOfStrings *[]*string, value *string) cdktn.IResolvable {
	if err := c.validateConditionEachMemberEqualsParameters(listOfStrings, value); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"conditionEachMemberEquals",
		[]interface{}{listOfStrings, value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProviderFunctions) ConditionEachMemberIn(stringsToCheck *[]*string, stringsToMatch *[]*string) cdktn.IResolvable {
	if err := c.validateConditionEachMemberInParameters(stringsToCheck, stringsToMatch); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"conditionEachMemberIn",
		[]interface{}{stringsToCheck, stringsToMatch},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProviderFunctions) ConditionEquals(value1 interface{}, value2 interface{}) cdktn.IResolvable {
	if err := c.validateConditionEqualsParameters(value1, value2); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"conditionEquals",
		[]interface{}{value1, value2},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProviderFunctions) ConditionIf(condition interface{}, valueIfTrue interface{}, valueIfFalse interface{}) cdktn.IResolvable {
	if err := c.validateConditionIfParameters(condition); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"conditionIf",
		[]interface{}{condition, valueIfTrue, valueIfFalse},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProviderFunctions) ConditionNot(condition interface{}) cdktn.IResolvable {
	if err := c.validateConditionNotParameters(condition); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"conditionNot",
		[]interface{}{condition},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProviderFunctions) ConditionOr(conditions *[]interface{}) cdktn.IResolvable {
	if err := c.validateConditionOrParameters(conditions); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"conditionOr",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProviderFunctions) FindInMap(mapping interface{}, topLevelKey *string, secondLevelKey *string, defaultValue *[]interface{}) cdktn.IResolvable {
	if err := c.validateFindInMapParameters(mapping, topLevelKey, secondLevelKey, defaultValue); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"findInMap",
		[]interface{}{mapping, topLevelKey, secondLevelKey, defaultValue},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProviderFunctions) Join(delimiter *string, values *[]*string) *string {
	if err := c.validateJoinParameters(delimiter, values); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"join",
		[]interface{}{delimiter, values},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProviderFunctions) LengthOf(value interface{}) *float64 {
	if err := c.validateLengthOfParameters(value); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"lengthOf",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProviderFunctions) Select(index *float64, objects interface{}) cdktn.IResolvable {
	if err := c.validateSelectParameters(index, objects); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"select",
		[]interface{}{index, objects},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProviderFunctions) Split(delimiter *string, source *string) *[]*string {
	if err := c.validateSplitParameters(delimiter, source); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"split",
		[]interface{}{delimiter, source},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProviderFunctions) Sub(template *string, variables *map[string]*string) *string {
	if err := c.validateSubParameters(template, variables); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"sub",
		[]interface{}{template, variables},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProviderFunctions) ToJsonString(value interface{}) *string {
	if err := c.validateToJsonStringParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"toJsonString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

