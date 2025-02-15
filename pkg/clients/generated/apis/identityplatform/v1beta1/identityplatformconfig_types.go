// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConfigAnonymous struct {
	/* Whether anonymous user auth is enabled for the project or not. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`
}

type ConfigBlockingFunctions struct {
	/* Map of Trigger to event type. Key should be one of the supported event types: "beforeCreate", "beforeSignIn" */
	// +optional
	Triggers map[string]ConfigTriggers `json:"triggers,omitempty"`
}

type ConfigChangeEmailTemplate struct {
	/* Immutable. Email body */
	// +optional
	Body *string `json:"body,omitempty"`

	/* Email body format Possible values: BODY_FORMAT_UNSPECIFIED, PLAIN_TEXT, HTML */
	// +optional
	BodyFormat *string `json:"bodyFormat,omitempty"`

	/* Reply-to address */
	// +optional
	ReplyTo *string `json:"replyTo,omitempty"`

	/* Sender display name */
	// +optional
	SenderDisplayName *string `json:"senderDisplayName,omitempty"`

	/* Local part of From address */
	// +optional
	SenderLocalPart *string `json:"senderLocalPart,omitempty"`

	/* Subject of the email */
	// +optional
	Subject *string `json:"subject,omitempty"`
}

type ConfigClient struct {
	/* Configuration related to restricting a user's ability to affect their account. */
	// +optional
	Permissions *ConfigPermissions `json:"permissions,omitempty"`
}

type ConfigDnsInfo struct {
	/* Whether to use custom domain. */
	// +optional
	UseCustomDomain *bool `json:"useCustomDomain,omitempty"`
}

type ConfigEmail struct {
	/* Whether email auth is enabled for the project or not. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`

	/* Whether a password is required for email auth or not. If true, both an email and password must be provided to sign in. If false, a user may sign in via either email/password or email link. */
	// +optional
	PasswordRequired *bool `json:"passwordRequired,omitempty"`
}

type ConfigMfa struct {
	/* Whether MultiFactor Authentication has been enabled for this project. Possible values: STATE_UNSPECIFIED, DISABLED, ENABLED, MANDATORY */
	// +optional
	State *string `json:"state,omitempty"`
}

type ConfigMonitoring struct {
	/* Configuration for logging requests made to this project to Stackdriver Logging */
	// +optional
	RequestLogging *ConfigRequestLogging `json:"requestLogging,omitempty"`
}

type ConfigMultiTenant struct {
	/* Whether this project can have tenants or not. */
	// +optional
	AllowTenants *bool `json:"allowTenants,omitempty"`

	/*  */
	// +optional
	DefaultTenantLocationRef *v1alpha1.ResourceRef `json:"defaultTenantLocationRef,omitempty"`
}

type ConfigNotification struct {
	/* Default locale used for email and SMS in IETF BCP 47 format. */
	// +optional
	DefaultLocale *string `json:"defaultLocale,omitempty"`

	/* Options for email sending. */
	// +optional
	SendEmail *ConfigSendEmail `json:"sendEmail,omitempty"`

	/* Options for SMS sending. */
	// +optional
	SendSms *ConfigSendSms `json:"sendSms,omitempty"`
}

type ConfigPassword struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *ConfigValueFrom `json:"valueFrom,omitempty"`
}

type ConfigPermissions struct {
	/* When true, end users cannot delete their account on the associated project through any of our API methods */
	// +optional
	DisabledUserDeletion *bool `json:"disabledUserDeletion,omitempty"`

	/* When true, end users cannot sign up for a new account on the associated project through any of our API methods */
	// +optional
	DisabledUserSignup *bool `json:"disabledUserSignup,omitempty"`
}

type ConfigPhoneNumber struct {
	/* Whether phone number auth is enabled for the project or not. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`

	/* A map of that can be used for phone auth testing. */
	// +optional
	TestPhoneNumbers map[string]string `json:"testPhoneNumbers,omitempty"`
}

type ConfigQuota struct {
	/* Quota for the Signup endpoint, if overwritten. Signup quota is measured in sign ups per project per hour per IP. */
	// +optional
	SignUpQuotaConfig *ConfigSignUpQuotaConfig `json:"signUpQuotaConfig,omitempty"`
}

type ConfigRequestLogging struct {
	/* Whether logging is enabled for this project or not. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`
}

type ConfigResetPasswordTemplate struct {
	/* Email body */
	// +optional
	Body *string `json:"body,omitempty"`

	/* Email body format Possible values: BODY_FORMAT_UNSPECIFIED, PLAIN_TEXT, HTML */
	// +optional
	BodyFormat *string `json:"bodyFormat,omitempty"`

	/* Reply-to address */
	// +optional
	ReplyTo *string `json:"replyTo,omitempty"`

	/* Sender display name */
	// +optional
	SenderDisplayName *string `json:"senderDisplayName,omitempty"`

	/* Local part of From address */
	// +optional
	SenderLocalPart *string `json:"senderLocalPart,omitempty"`

	/* Subject of the email */
	// +optional
	Subject *string `json:"subject,omitempty"`
}

type ConfigRevertSecondFactorAdditionTemplate struct {
	/* Immutable. Email body */
	// +optional
	Body *string `json:"body,omitempty"`

	/* Email body format Possible values: BODY_FORMAT_UNSPECIFIED, PLAIN_TEXT, HTML */
	// +optional
	BodyFormat *string `json:"bodyFormat,omitempty"`

	/* Reply-to address */
	// +optional
	ReplyTo *string `json:"replyTo,omitempty"`

	/* Sender display name */
	// +optional
	SenderDisplayName *string `json:"senderDisplayName,omitempty"`

	/* Local part of From address */
	// +optional
	SenderLocalPart *string `json:"senderLocalPart,omitempty"`

	/* Subject of the email */
	// +optional
	Subject *string `json:"subject,omitempty"`
}

type ConfigSendEmail struct {
	/* action url in email template. */
	// +optional
	CallbackUri *string `json:"callbackUri,omitempty"`

	/* Email template for change email */
	// +optional
	ChangeEmailTemplate *ConfigChangeEmailTemplate `json:"changeEmailTemplate,omitempty"`

	/* Information of custom domain DNS verification. */
	// +optional
	DnsInfo *ConfigDnsInfo `json:"dnsInfo,omitempty"`

	/* The method used for sending an email. Possible values: METHOD_UNSPECIFIED, DEFAULT, CUSTOM_SMTP */
	// +optional
	Method *string `json:"method,omitempty"`

	/* Email template for reset password */
	// +optional
	ResetPasswordTemplate *ConfigResetPasswordTemplate `json:"resetPasswordTemplate,omitempty"`

	/* Email template for reverting second factor addition emails */
	// +optional
	RevertSecondFactorAdditionTemplate *ConfigRevertSecondFactorAdditionTemplate `json:"revertSecondFactorAdditionTemplate,omitempty"`

	/* Use a custom SMTP relay */
	// +optional
	Smtp *ConfigSmtp `json:"smtp,omitempty"`

	/* Email template for verify email */
	// +optional
	VerifyEmailTemplate *ConfigVerifyEmailTemplate `json:"verifyEmailTemplate,omitempty"`
}

type ConfigSendSms struct {
	/* Whether to use the accept_language header for SMS. */
	// +optional
	UseDeviceLocale *bool `json:"useDeviceLocale,omitempty"`
}

type ConfigSignIn struct {
	/* Whether to allow more than one account to have the same email. */
	// +optional
	AllowDuplicateEmails *bool `json:"allowDuplicateEmails,omitempty"`

	/* Configuration options related to authenticating an anonymous user. */
	// +optional
	Anonymous *ConfigAnonymous `json:"anonymous,omitempty"`

	/* Configuration options related to authenticating a user by their email address. */
	// +optional
	Email *ConfigEmail `json:"email,omitempty"`

	/* Configuration options related to authenticated a user by their phone number. */
	// +optional
	PhoneNumber *ConfigPhoneNumber `json:"phoneNumber,omitempty"`
}

type ConfigSignUpQuotaConfig struct {
	/* Corresponds to the 'refill_token_count' field in QuotaServer config */
	// +optional
	Quota *int `json:"quota,omitempty"`

	/* How long this quota will be active for */
	// +optional
	QuotaDuration *string `json:"quotaDuration,omitempty"`

	/* When this quota will take affect */
	// +optional
	StartTime *string `json:"startTime,omitempty"`
}

type ConfigSmtp struct {
	/* SMTP relay host */
	// +optional
	Host *string `json:"host,omitempty"`

	/* SMTP relay password */
	// +optional
	Password *ConfigPassword `json:"password,omitempty"`

	/* SMTP relay port */
	// +optional
	Port *int `json:"port,omitempty"`

	/* SMTP security mode. Possible values: SECURITY_MODE_UNSPECIFIED, SSL, START_TLS */
	// +optional
	SecurityMode *string `json:"securityMode,omitempty"`

	/* Sender email for the SMTP relay */
	// +optional
	SenderEmail *string `json:"senderEmail,omitempty"`

	/* SMTP relay username */
	// +optional
	Username *string `json:"username,omitempty"`
}

type ConfigTriggers struct {
	/*  */
	// +optional
	FunctionUriRef *v1alpha1.ResourceRef `json:"functionUriRef,omitempty"`

	/* When the trigger was changed. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

type ConfigValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *v1alpha1.ResourceRef `json:"secretKeyRef,omitempty"`
}

type ConfigVerifyEmailTemplate struct {
	/* Immutable. Email body */
	// +optional
	Body *string `json:"body,omitempty"`

	/* Email body format Possible values: BODY_FORMAT_UNSPECIFIED, PLAIN_TEXT, HTML */
	// +optional
	BodyFormat *string `json:"bodyFormat,omitempty"`

	/* Reply-to address */
	// +optional
	ReplyTo *string `json:"replyTo,omitempty"`

	/* Sender display name */
	// +optional
	SenderDisplayName *string `json:"senderDisplayName,omitempty"`

	/* Local part of From address */
	// +optional
	SenderLocalPart *string `json:"senderLocalPart,omitempty"`

	/* Subject of the email */
	// +optional
	Subject *string `json:"subject,omitempty"`
}

type IdentityPlatformConfigSpec struct {
	/* List of domains authorized for OAuth redirects */
	// +optional
	AuthorizedDomains []string `json:"authorizedDomains,omitempty"`

	/* Configuration related to blocking functions. */
	// +optional
	BlockingFunctions *ConfigBlockingFunctions `json:"blockingFunctions,omitempty"`

	/* Options related to how clients making requests on behalf of a project should be configured. */
	// +optional
	Client *ConfigClient `json:"client,omitempty"`

	/* Configuration for this project's multi-factor authentication, including whether it is active and what factors can be used for the second factor */
	// +optional
	Mfa *ConfigMfa `json:"mfa,omitempty"`

	/* Configuration related to monitoring project activity. */
	// +optional
	Monitoring *ConfigMonitoring `json:"monitoring,omitempty"`

	/* Configuration related to multi-tenant functionality. */
	// +optional
	MultiTenant *ConfigMultiTenant `json:"multiTenant,omitempty"`

	/* Configuration related to sending notifications to users. */
	// +optional
	Notification *ConfigNotification `json:"notification,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Configuration related to quotas. */
	// +optional
	Quota *ConfigQuota `json:"quota,omitempty"`

	/* Configuration related to local sign in methods. */
	// +optional
	SignIn *ConfigSignIn `json:"signIn,omitempty"`
}

type ConfigChangeEmailTemplateStatus struct {
	/* Output only. Whether the body or subject of the email is customized. */
	Customized bool `json:"customized,omitempty"`
}

type ConfigClientStatus struct {
	/* Output only. API key that can be used when making requests for this project. */
	ApiKey string `json:"apiKey,omitempty"`

	/* Output only. Firebase subdomain. */
	FirebaseSubdomain string `json:"firebaseSubdomain,omitempty"`
}

type ConfigDnsInfoStatus struct {
	/* Output only. The applied verified custom domain. */
	CustomDomain string `json:"customDomain,omitempty"`

	/* Output only. The current verification state of the custom domain. The custom domain will only be used once the domain verification is successful. Possible values: VERIFICATION_STATE_UNSPECIFIED, NOT_STARTED, IN_PROGRESS, FAILED, SUCCEEDED */
	CustomDomainState string `json:"customDomainState,omitempty"`

	/* Output only. The timestamp of initial request for the current domain verification. */
	DomainVerificationRequestTime string `json:"domainVerificationRequestTime,omitempty"`

	/* Output only. The custom domain that's to be verified. */
	PendingCustomDomain string `json:"pendingCustomDomain,omitempty"`
}

type ConfigEmailStatus struct {
	/* Output only. Hash config information. */
	HashConfig ConfigHashConfigStatus `json:"hashConfig,omitempty"`
}

type ConfigHashConfigStatus struct {
	/* Output only. Different password hash algorithms used in Identity Toolkit. Possible values: HASH_ALGORITHM_UNSPECIFIED, HMAC_SHA256, HMAC_SHA1, HMAC_MD5, SCRYPT, PBKDF_SHA1, MD5, HMAC_SHA512, SHA1, BCRYPT, PBKDF2_SHA256, SHA256, SHA512, STANDARD_SCRYPT */
	Algorithm string `json:"algorithm,omitempty"`

	/* Output only. Memory cost for hash calculation. Used by scrypt and other similar password derivation algorithms. See https://tools.ietf.org/html/rfc7914 for explanation of field. */
	MemoryCost int `json:"memoryCost,omitempty"`

	/* Output only. How many rounds for hash calculation. Used by scrypt and other similar password derivation algorithms. */
	Rounds int `json:"rounds,omitempty"`

	/* Output only. Non-printable character to be inserted between the salt and plain text password in base64. */
	SaltSeparator string `json:"saltSeparator,omitempty"`

	/* Output only. Signer key in base64. */
	SignerKey string `json:"signerKey,omitempty"`
}

type ConfigNotificationStatus struct {
	/*  */
	SendEmail ConfigSendEmailStatus `json:"sendEmail,omitempty"`

	/*  */
	SendSms ConfigSendSmsStatus `json:"sendSms,omitempty"`
}

type ConfigResetPasswordTemplateStatus struct {
	/* Output only. Whether the body or subject of the email is customized. */
	Customized bool `json:"customized,omitempty"`
}

type ConfigRevertSecondFactorAdditionTemplateStatus struct {
	/* Output only. Whether the body or subject of the email is customized. */
	Customized bool `json:"customized,omitempty"`
}

type ConfigSendEmailStatus struct {
	/*  */
	ChangeEmailTemplate ConfigChangeEmailTemplateStatus `json:"changeEmailTemplate,omitempty"`

	/*  */
	DnsInfo ConfigDnsInfoStatus `json:"dnsInfo,omitempty"`

	/*  */
	ResetPasswordTemplate ConfigResetPasswordTemplateStatus `json:"resetPasswordTemplate,omitempty"`

	/*  */
	RevertSecondFactorAdditionTemplate ConfigRevertSecondFactorAdditionTemplateStatus `json:"revertSecondFactorAdditionTemplate,omitempty"`

	/*  */
	VerifyEmailTemplate ConfigVerifyEmailTemplateStatus `json:"verifyEmailTemplate,omitempty"`
}

type ConfigSendSmsStatus struct {
	/* Output only. The template to use when sending an SMS. */
	SmsTemplate ConfigSmsTemplateStatus `json:"smsTemplate,omitempty"`
}

type ConfigSignInStatus struct {
	/*  */
	Email ConfigEmailStatus `json:"email,omitempty"`

	/* Output only. Hash config information. */
	HashConfig ConfigHashConfigStatus `json:"hashConfig,omitempty"`
}

type ConfigSmsTemplateStatus struct {
	/* Output only. The SMS's content. Can contain the following placeholders which will be replaced with the appropriate values: %APP_NAME% - For Android or iOS apps, the app's display name. For web apps, the domain hosting the application. %LOGIN_CODE% - The OOB code being sent in the SMS. */
	Content string `json:"content,omitempty"`
}

type ConfigVerifyEmailTemplateStatus struct {
	/* Output only. Whether the body or subject of the email is customized. */
	Customized bool `json:"customized,omitempty"`
}

type IdentityPlatformConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   IdentityPlatformConfig's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/*  */
	Client ConfigClientStatus `json:"client,omitempty"`
	/*  */
	Notification ConfigNotificationStatus `json:"notification,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/*  */
	SignIn ConfigSignInStatus `json:"signIn,omitempty"`
	/* Output only. The subtype of this config. Possible values: SUBTYPE_UNSPECIFIED, IDENTITY_PLATFORM, FIREBASE_AUTH */
	Subtype string `json:"subtype,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IdentityPlatformConfig is the Schema for the identityplatform API
// +k8s:openapi-gen=true
type IdentityPlatformConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdentityPlatformConfigSpec   `json:"spec,omitempty"`
	Status IdentityPlatformConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IdentityPlatformConfigList contains a list of IdentityPlatformConfig
type IdentityPlatformConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityPlatformConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IdentityPlatformConfig{}, &IdentityPlatformConfigList{})
}
