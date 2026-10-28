// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-cfncompat-go/cfncompat/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-cfncompat-go/cfncompat/v2/provider/internal"
	"github.com/cdktn-io/cdktn-provider-cfncompat-go/cfncompat/v2/providerfunctions"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs cfncompat}.
type CfncompatProvider interface {
	cdktn.TerraformProvider
	AccessKey() *string
	SetAccessKey(val *string)
	AccessKeyInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AssumeRole() *CfncompatProviderAssumeRole
	SetAssumeRole(val *CfncompatProviderAssumeRole)
	AssumeRoleInput() *CfncompatProviderAssumeRole
	AssumeRoleWithWebIdentity() *CfncompatProviderAssumeRoleWithWebIdentity
	SetAssumeRoleWithWebIdentity(val *CfncompatProviderAssumeRoleWithWebIdentity)
	AssumeRoleWithWebIdentityInput() *CfncompatProviderAssumeRoleWithWebIdentity
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CustomResourceBucket() *string
	SetCustomResourceBucket(val *string)
	CustomResourceBucketInput() *string
	Endpoints() *CfncompatProviderEndpoints
	SetEndpoints(val *CfncompatProviderEndpoints)
	EndpointsInput() *CfncompatProviderEndpoints
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Provider-defined functions of the cfncompat provider.
	Functions() providerfunctions.CfncompatProviderFunctions
	HttpProxy() *string
	SetHttpProxy(val *string)
	HttpProxyInput() *string
	HttpsProxy() *string
	SetHttpsProxy(val *string)
	HttpsProxyInput() *string
	Insecure() interface{}
	SetInsecure(val interface{})
	InsecureInput() interface{}
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	NoProxy() *string
	SetNoProxy(val *string)
	NoProxyInput() *string
	Profile() *string
	SetProfile(val *string)
	ProfileInput() *string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SecretKey() *string
	SetSecretKey(val *string)
	SecretKeyInput() *string
	SharedConfigFiles() *[]*string
	SetSharedConfigFiles(val *[]*string)
	SharedConfigFilesInput() *[]*string
	SharedCredentialsFiles() *[]*string
	SetSharedCredentialsFiles(val *[]*string)
	SharedCredentialsFilesInput() *[]*string
	SkipMetadataApiCheck() interface{}
	SetSkipMetadataApiCheck(val interface{})
	SkipMetadataApiCheckInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetAccessKey()
	ResetAlias()
	ResetAssumeRole()
	ResetAssumeRoleWithWebIdentity()
	ResetCustomResourceBucket()
	ResetEndpoints()
	ResetHttpProxy()
	ResetHttpsProxy()
	ResetInsecure()
	ResetMaxRetries()
	ResetNoProxy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProfile()
	ResetRegion()
	ResetSecretKey()
	ResetSharedConfigFiles()
	ResetSharedCredentialsFiles()
	ResetSkipMetadataApiCheck()
	ResetToken()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfncompatProvider
type jsiiProxy_CfncompatProvider struct {
	internal.Type__cdktnTerraformProvider
}

func (j *jsiiProxy_CfncompatProvider) AccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) AccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) AssumeRole() *CfncompatProviderAssumeRole {
	var returns *CfncompatProviderAssumeRole
	_jsii_.Get(
		j,
		"assumeRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) AssumeRoleInput() *CfncompatProviderAssumeRole {
	var returns *CfncompatProviderAssumeRole
	_jsii_.Get(
		j,
		"assumeRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) AssumeRoleWithWebIdentity() *CfncompatProviderAssumeRoleWithWebIdentity {
	var returns *CfncompatProviderAssumeRoleWithWebIdentity
	_jsii_.Get(
		j,
		"assumeRoleWithWebIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) AssumeRoleWithWebIdentityInput() *CfncompatProviderAssumeRoleWithWebIdentity {
	var returns *CfncompatProviderAssumeRoleWithWebIdentity
	_jsii_.Get(
		j,
		"assumeRoleWithWebIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) CustomResourceBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customResourceBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) CustomResourceBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customResourceBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) Endpoints() *CfncompatProviderEndpoints {
	var returns *CfncompatProviderEndpoints
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) EndpointsInput() *CfncompatProviderEndpoints {
	var returns *CfncompatProviderEndpoints
	_jsii_.Get(
		j,
		"endpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) Functions() providerfunctions.CfncompatProviderFunctions {
	var returns providerfunctions.CfncompatProviderFunctions
	_jsii_.Get(
		j,
		"functions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) HttpProxy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) HttpProxyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) HttpsProxy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) HttpsProxyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) Insecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) InsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) NoProxy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) NoProxyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) SecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) SecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) SharedConfigFiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedConfigFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) SharedConfigFilesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedConfigFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) SharedCredentialsFiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedCredentialsFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) SharedCredentialsFilesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedCredentialsFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) SkipMetadataApiCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipMetadataApiCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) SkipMetadataApiCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipMetadataApiCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfncompatProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs cfncompat} Resource.
func NewCfncompatProvider(scope constructs.Construct, id *string, config *CfncompatProviderConfig) CfncompatProvider {
	_init_.Initialize()

	if err := validateNewCfncompatProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfncompatProvider{}

	_jsii_.Create(
		"@cdktn/provider-cfncompat.provider.CfncompatProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cdktn-io/cfncompat/0.3.0/docs cfncompat} Resource.
func NewCfncompatProvider_Override(c CfncompatProvider, scope constructs.Construct, id *string, config *CfncompatProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cfncompat.provider.CfncompatProvider",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetAccessKey(val *string) {
	_jsii_.Set(
		j,
		"accessKey",
		val,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetAssumeRole(val *CfncompatProviderAssumeRole) {
	if err := j.validateSetAssumeRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assumeRole",
		val,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetAssumeRoleWithWebIdentity(val *CfncompatProviderAssumeRoleWithWebIdentity) {
	if err := j.validateSetAssumeRoleWithWebIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assumeRoleWithWebIdentity",
		val,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetCustomResourceBucket(val *string) {
	_jsii_.Set(
		j,
		"customResourceBucket",
		val,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetEndpoints(val *CfncompatProviderEndpoints) {
	if err := j.validateSetEndpointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoints",
		val,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetHttpProxy(val *string) {
	_jsii_.Set(
		j,
		"httpProxy",
		val,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetHttpsProxy(val *string) {
	_jsii_.Set(
		j,
		"httpsProxy",
		val,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetInsecure(val interface{}) {
	if err := j.validateSetInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecure",
		val,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetNoProxy(val *string) {
	_jsii_.Set(
		j,
		"noProxy",
		val,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetProfile(val *string) {
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetSecretKey(val *string) {
	_jsii_.Set(
		j,
		"secretKey",
		val,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetSharedConfigFiles(val *[]*string) {
	_jsii_.Set(
		j,
		"sharedConfigFiles",
		val,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetSharedCredentialsFiles(val *[]*string) {
	_jsii_.Set(
		j,
		"sharedCredentialsFiles",
		val,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetSkipMetadataApiCheck(val interface{}) {
	if err := j.validateSetSkipMetadataApiCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipMetadataApiCheck",
		val,
	)
}

func (j *jsiiProxy_CfncompatProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

// Generates CDKTN code for importing a CfncompatProvider resource upon running "cdktn plan <stack-name>".
func CfncompatProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateCfncompatProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-cfncompat.provider.CfncompatProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfncompatProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfncompatProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cfncompat.provider.CfncompatProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CfncompatProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfncompatProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cfncompat.provider.CfncompatProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CfncompatProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfncompatProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cfncompat.provider.CfncompatProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfncompatProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-cfncompat.provider.CfncompatProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfncompatProvider) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfncompatProvider) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfncompatProvider) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := c.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetAccessKey() {
	_jsii_.InvokeVoid(
		c,
		"resetAccessKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		c,
		"resetAlias",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetAssumeRole() {
	_jsii_.InvokeVoid(
		c,
		"resetAssumeRole",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetAssumeRoleWithWebIdentity() {
	_jsii_.InvokeVoid(
		c,
		"resetAssumeRoleWithWebIdentity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetCustomResourceBucket() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomResourceBucket",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetEndpoints() {
	_jsii_.InvokeVoid(
		c,
		"resetEndpoints",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetHttpProxy() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpProxy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetHttpsProxy() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpsProxy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetInsecure() {
	_jsii_.InvokeVoid(
		c,
		"resetInsecure",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetNoProxy() {
	_jsii_.InvokeVoid(
		c,
		"resetNoProxy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetProfile() {
	_jsii_.InvokeVoid(
		c,
		"resetProfile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetSecretKey() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetSharedConfigFiles() {
	_jsii_.InvokeVoid(
		c,
		"resetSharedConfigFiles",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetSharedCredentialsFiles() {
	_jsii_.InvokeVoid(
		c,
		"resetSharedCredentialsFiles",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetSkipMetadataApiCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetSkipMetadataApiCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) ResetToken() {
	_jsii_.InvokeVoid(
		c,
		"resetToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfncompatProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfncompatProvider) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

