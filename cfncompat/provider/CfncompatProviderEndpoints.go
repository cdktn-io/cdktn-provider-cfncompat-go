// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type CfncompatProviderEndpoints struct {
	// Override the default EC2 service endpoint URL, used by the cfncompat_availability_zones data source.
	//
	// The override applies to every region, including one named by an explicit `region` argument on cfncompat_availability_zones -- the request is still signed for that region, but always sent to this endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs#ec2 CfncompatProvider#ec2}
	Ec2 *string `field:"optional" json:"ec2" yaml:"ec2"`
	// Override the default Lambda service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs#lambda CfncompatProvider#lambda}
	Lambda *string `field:"optional" json:"lambda" yaml:"lambda"`
	// Override the default S3 service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs#s3 CfncompatProvider#s3}
	S3 *string `field:"optional" json:"s3" yaml:"s3"`
	// Override the default SNS service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs#sns CfncompatProvider#sns}
	Sns *string `field:"optional" json:"sns" yaml:"sns"`
	// Override the default STS service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs#sts CfncompatProvider#sts}
	Sts *string `field:"optional" json:"sts" yaml:"sts"`
}

