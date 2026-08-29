// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type CfncompatProviderAssumeRoleWithWebIdentity struct {
	// ARN of the IAM role to assume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs#role_arn CfncompatProvider#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The duration of the role session, e.g. "1h". Parsed with Go's time.ParseDuration (valid units: ns, us, ms, s, m, h).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs#duration CfncompatProvider#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// IAM policy in JSON format used as a session policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs#policy CfncompatProvider#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// ARNs of IAM policies used as managed session policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs#policy_arns CfncompatProvider#policy_arns}
	PolicyArns *[]*string `field:"optional" json:"policyArns" yaml:"policyArns"`
	// Session name to use when assuming the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs#session_name CfncompatProvider#session_name}
	SessionName *string `field:"optional" json:"sessionName" yaml:"sessionName"`
	// Value of a web identity token from an OIDC/OAuth provider. One of `web_identity_token` or `web_identity_token_file` is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs#web_identity_token CfncompatProvider#web_identity_token}
	WebIdentityToken *string `field:"optional" json:"webIdentityToken" yaml:"webIdentityToken"`
	// File containing a web identity token from an OIDC/OAuth provider. One of `web_identity_token_file` or `web_identity_token` is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs#web_identity_token_file CfncompatProvider#web_identity_token_file}
	WebIdentityTokenFile *string `field:"optional" json:"webIdentityTokenFile" yaml:"webIdentityTokenFile"`
}

