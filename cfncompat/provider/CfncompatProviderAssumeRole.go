// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type CfncompatProviderAssumeRole struct {
	// ARN of the IAM role to assume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.2.0/docs#role_arn CfncompatProvider#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The duration of the role session, e.g. "1h". Parsed with Go's time.ParseDuration (valid units: ns, us, ms, s, m, h).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.2.0/docs#duration CfncompatProvider#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// External identifier to use when assuming the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.2.0/docs#external_id CfncompatProvider#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// IAM policy in JSON format used as a session policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.2.0/docs#policy CfncompatProvider#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// ARNs of IAM policies used as managed session policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.2.0/docs#policy_arns CfncompatProvider#policy_arns}
	PolicyArns *[]*string `field:"optional" json:"policyArns" yaml:"policyArns"`
	// Session name to use when assuming the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.2.0/docs#session_name CfncompatProvider#session_name}
	SessionName *string `field:"optional" json:"sessionName" yaml:"sessionName"`
	// Source identity to use when assuming the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.2.0/docs#source_identity CfncompatProvider#source_identity}
	SourceIdentity *string `field:"optional" json:"sourceIdentity" yaml:"sourceIdentity"`
	// Map of assume role session tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.2.0/docs#tags CfncompatProvider#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Set of assume role session tag keys to pass to any subsequent sessions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.2.0/docs#transitive_tag_keys CfncompatProvider#transitive_tag_keys}
	TransitiveTagKeys *[]*string `field:"optional" json:"transitiveTagKeys" yaml:"transitiveTagKeys"`
}

