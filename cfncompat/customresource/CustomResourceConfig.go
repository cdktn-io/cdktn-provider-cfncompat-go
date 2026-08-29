// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customresource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CustomResourceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// ARN of the Lambda function or SNS topic that implements the custom resource handler (CloudFormation's `ServiceToken`).
	//
	// Must be a Lambda function ARN (`arn:*:lambda:...`, invoked asynchronously) or an SNS topic ARN (`arn:*:sns:...`, published to).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs/resources/custom_resource#service_token CustomResource#service_token}
	ServiceToken *string `field:"required" json:"serviceToken" yaml:"serviceToken"`
	// CloudFormation-style logical resource id reported in the request event's `LogicalResourceId` field.
	//
	// Typically set by a CDK Terrain synthesis backend to the synthesized CloudFormation logical id; defaults to `"CfncompatCustomResource"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs/resources/custom_resource#logical_resource_id CustomResource#logical_resource_id}
	LogicalResourceId *string `field:"optional" json:"logicalResourceId" yaml:"logicalResourceId"`
	// Arbitrary user-defined properties passed to the handler as the request event's `ResourceProperties`.
	//
	// May be an object/map (arbitrarily nested) or omitted entirely. CloudFormation merges `ServiceToken` into this map when delivering it to the handler; this resource replicates that behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs/resources/custom_resource#resource_properties CustomResource#resource_properties}
	ResourceProperties *map[string]interface{} `field:"optional" json:"resourceProperties" yaml:"resourceProperties"`
	// CloudFormation resource type name reported in the request event's `ResourceType` field.
	//
	// Defaults to `"AWS::CloudFormation::CustomResource"`. May also be set to `Custom::<name>` (matching CloudFormation's convention for named custom resource types), where `<name>` matches `^[A-Za-z0-9_@-]{1,52}$` and the full value is at most 60 characters. Changing this forces replacement: CloudFormation forbids changing a custom resource's type in an update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs/resources/custom_resource#resource_type CustomResource#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// S3 bucket used for this resource's response transport (the pre-signed PUT URL the handler writes its response to).
	//
	// Falls back to the provider's `custom_resource_bucket` if unset; it is an error at apply time if neither is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs/resources/custom_resource#response_bucket CustomResource#response_bucket}
	ResponseBucket *string `field:"optional" json:"responseBucket" yaml:"responseBucket"`
	// Optional S3 key prefix for the response object.
	//
	// The full key is `"<response_key_prefix>cfncompat/<RequestId>.json"` -- include a trailing `/` if you want the prefix to behave like a folder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs/resources/custom_resource#response_key_prefix CustomResource#response_key_prefix}
	ResponseKeyPrefix *string `field:"optional" json:"responseKeyPrefix" yaml:"responseKeyPrefix"`
	// Seconds to wait for the handler's response before failing, mirroring CloudFormation's `ServiceTimeout`.
	//
	// Must be between 1 and 3600 (CloudFormation's own range). Defaults to `3600`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs/resources/custom_resource#service_timeout CustomResource#service_timeout}
	ServiceTimeout *float64 `field:"optional" json:"serviceTimeout" yaml:"serviceTimeout"`
	// CloudFormation-style stack identifier reported in the request event's `StackId` field, passed through verbatim.
	//
	// Typically set by a CDK Terrain synthesis backend to a stack identifier; defaults to `"cfncompat/no-stack-id"`.
	//
	// ~> That default is a **shared sentinel**: every `cfncompat_custom_resource` in the workspace that leaves `stack_id` unset sends the same value. Handlers that treat `StackId` as an ownership key then cannot tell one stack's objects from another's -- CDK's S3 notifications handler, for instance, prefixes every notification `Id` with `{StackId}-` and, on delete, removes exactly the notifications carrying that prefix, so two stacks sharing the sentinel would delete each other's notifications. Wire this to `data.cfncompat_pseudo_parameters.<name>.stack_id` with `stack_name` set (that value is deterministic and stable across applies); leaving it unset emits a warning today and is planned to become an error in v1.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs/resources/custom_resource#stack_id CustomResource#stack_id}
	StackId *string `field:"optional" json:"stackId" yaml:"stackId"`
}

