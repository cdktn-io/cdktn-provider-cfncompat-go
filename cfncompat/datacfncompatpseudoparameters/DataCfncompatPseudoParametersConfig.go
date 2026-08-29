// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacfncompatpseudoparameters

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCfncompatPseudoParametersConfig struct {
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
	// CloudFormation's [`AWS::NotificationARNs`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/pseudo-parameter-reference.html), echoed back unchanged; defaults to `[]` when not set. There is no CloudFormation stack to send notifications for, so the synthesis backend (e.g. CDK Terrain, passing `StackProps.notificationArns`) supplies the list it would have passed to `CreateStack`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs/data-sources/pseudo_parameters#notification_arns DataCfncompatPseudoParameters#notification_arns}
	NotificationArns *[]*string `field:"optional" json:"notificationArns" yaml:"notificationArns"`
	// CloudFormation's [`AWS::StackName`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/pseudo-parameter-reference.html), echoed back unchanged. There is no real CloudFormation stack behind a Terraform apply, so the synthesis backend (e.g. CDK Terrain, passing `Stack.stackName`) supplies the name it uses for the stack. Must be a non-empty string when set (an empty name would derive a `stack_id` for a stack that cannot exist).
	//
	// It is also the only input `stack_id` is derived from: leaving it null makes `stack_id` null, and a `cfncompat_custom_resource.stack_id` fed from a null value falls back to that resource's shared `"cfncompat/no-stack-id"` default -- which every stack in the workspace shares. Set `stack_name` whenever a custom-resource handler keys on `StackId`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs/data-sources/pseudo_parameters#stack_name DataCfncompatPseudoParameters#stack_name}
	StackName *string `field:"optional" json:"stackName" yaml:"stackName"`
}

