// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package providerfunctions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-cfncompat.providerFunctions.CfncompatProviderFunctions",
		reflect.TypeOf((*CfncompatProviderFunctions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "base64", GoMethod: "Base64"},
			_jsii_.MemberMethod{JsiiMethod: "cidr", GoMethod: "Cidr"},
			_jsii_.MemberMethod{JsiiMethod: "conditionAnd", GoMethod: "ConditionAnd"},
			_jsii_.MemberMethod{JsiiMethod: "conditionContains", GoMethod: "ConditionContains"},
			_jsii_.MemberMethod{JsiiMethod: "conditionEachMemberEquals", GoMethod: "ConditionEachMemberEquals"},
			_jsii_.MemberMethod{JsiiMethod: "conditionEachMemberIn", GoMethod: "ConditionEachMemberIn"},
			_jsii_.MemberMethod{JsiiMethod: "conditionEquals", GoMethod: "ConditionEquals"},
			_jsii_.MemberMethod{JsiiMethod: "conditionIf", GoMethod: "ConditionIf"},
			_jsii_.MemberMethod{JsiiMethod: "conditionNot", GoMethod: "ConditionNot"},
			_jsii_.MemberMethod{JsiiMethod: "conditionOr", GoMethod: "ConditionOr"},
			_jsii_.MemberMethod{JsiiMethod: "findInMap", GoMethod: "FindInMap"},
			_jsii_.MemberMethod{JsiiMethod: "join", GoMethod: "Join"},
			_jsii_.MemberMethod{JsiiMethod: "lengthOf", GoMethod: "LengthOf"},
			_jsii_.MemberMethod{JsiiMethod: "select", GoMethod: "Select"},
			_jsii_.MemberMethod{JsiiMethod: "split", GoMethod: "Split"},
			_jsii_.MemberMethod{JsiiMethod: "sub", GoMethod: "Sub"},
			_jsii_.MemberMethod{JsiiMethod: "toJsonString", GoMethod: "ToJsonString"},
		},
		func() interface{} {
			return &jsiiProxy_CfncompatProviderFunctions{}
		},
	)
}
