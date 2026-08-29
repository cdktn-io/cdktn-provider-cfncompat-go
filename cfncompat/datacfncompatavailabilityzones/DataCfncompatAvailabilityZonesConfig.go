// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datacfncompatavailabilityzones

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataCfncompatAvailabilityZonesConfig struct {
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
	// The region to list Availability Zones for, i.e. the `Fn::GetAZs` argument. An empty string or an unset value means the provider's own region -- exactly as [the `Fn::GetAZs` reference](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-getavailabilityzones.html) specifies that "specifying an empty string is equivalent to specifying `AWS::Region`". Echoed back as the region that was actually queried.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs/data-sources/availability_zones#region DataCfncompatAvailabilityZones#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

