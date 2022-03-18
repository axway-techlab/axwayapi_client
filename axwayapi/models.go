package axwayapi

//---
type Frontend struct {
	Id                     string                     `json:"id,omitempty"`
	OrganizationId         string                     `json:"organizationId"`
	ApiId                  string                     `json:"apiId"`
	Name                   string                     `json:"name"`
	Version                string                     `json:"version"`
	ApiRoutingKey          string                     `json:"apiRoutingKey"`
	Vhost                  string                     `json:"vhost"`
	Path                   string                     `json:"path,omitempty"`
	DescriptionType        string                     `json:"descriptionType,omitempty"`
	DescriptionManual      string                     `json:"descriptionManual,omitempty"`
	DescriptionMarkdown    string                     `json:"descriptionMarkdown,omitempty"`
	DescriptionUrl         string                     `json:"descriptionUrl,omitempty"`
	Summary                string                     `json:"summary,omitempty"`
	Retired                bool                       `json:"retired,omitempty"`
	Expired                bool                       `json:"expired,omitempty"`
	Image                  string                     `json:"image,omitempty"`
	RetirementDate         int                        `json:"retirementDate,omitempty"`
	Deprecated             bool                       `json:"deprecated,omitempty"`
	State                  string                     `json:"state,omitempty"`
	CorsProfiles           []CorsProfile              `json:"corsProfiles,omitempty"`
	SecurityProfiles       []SecurityProfile          `json:"securityProfiles,omitempty"`
	AuthenticationProfiles []AuthenticationProfile    `json:"authenticationProfiles,omitempty"`
	InboundProfiles        map[string]InboundProfile  `json:"inboundProfiles,omitempty"`
	OutboundProfiles       map[string]OutboundProfile `json:"outboundProfiles,omitempty"`
	ServiceProfiles        map[string]ServiceProfile  `json:"serviceProfiles,omitempty"`
	CACerts                []CACert                   `json:"caCerts,omitempty"`
	Tags                   map[string][]string        `json:"tags,omitempty"`
	CustomProperties       map[string]interface{}     `json:"customProperties,omitempty"`
	CreatedBy              string                     `json:"createdBy,omitempty"`
	CreatedOn              int                        `json:"createdOn,omitempty"`
}

type CorsProfile struct {
	//ValueObject
	Name               string   `json:"name"`
	IsDefault          bool     `json:"isDefault"`
	Origins            []string `json:"origins"`
	AllowedHeaders     []string `json:"allowedHeaders"`
	ExposedHeaders     []string `json:"exposedHeaders"`
	SupportCredentials bool     `json:"supportCredentials"`
	MaxAgeSeconds      int      `json:"maxAgeSeconds,omitempty"`
}
type SecurityProfile struct {
	//ValueObject
	Name      string   `json:"name"`
	IsDefault bool     `json:"isDefault"`
	Devices   []Device `json:"devices"`
}
type Device struct {
	//ValueObject
	Name       string                 `json:"name,omitempty"`
	Type       string                 `json:"type,omitempty"`
	Order      int                    `json:"order,omitempty"`
	Properties map[string]interface{} `json:"properties"`
}
type AuthenticationProfile struct {
	//ValueObject
	Name       string                 `json:"name,omitempty"`
	Type       string                 `json:"type,omitempty"`
	IsDefault  bool                   `json:"isDefault,omitempty"`
	Parameters map[string]interface{} `json:"parameters"`
}
type InboundProfile struct {
	//ValueObject
	SecurityProfile string `json:"securityProfile,omitempty"`
	CorsProfile     string `json:"corsProfile,omitempty"`
	MonitorAPI      bool   `json:"monitorAPI,omitempty"`
	MonitorSubject  string `json:"monitorSubject,omitempty"`
}
type OutboundProfile struct {
	//ValueObject
	AuthenticationProfile string       `json:"authenticationProfile,omitempty"`
	RouteType             string       `json:"routeType,omitempty"`
	RequestPolicy         string       `json:"requestPolicy,omitempty"`
	ResponsePolicy        string       `json:"responsePolicy,omitempty"`
	RoutePolicy           string       `json:"routePolicy,omitempty"`
	FaultHandlerPolicy    string       `json:"faultHandlerPolicy,omitempty"`
	ApiId                 string       `json:"apiId,omitempty"`
	ApiMethodId           string       `json:"apiMethodId,omitempty"`
	Parameters            []ParamValue `json:"parameters"`
}
type ParamValue struct {
	//ValueObject
	Name       string `json:"name,omitempty"`
	ParamType  string `json:"paramType,omitempty"`
	Type       string `json:"type,omitempty"`
	Format     string `json:"format,omitempty"`
	Value      string `json:"value,omitempty"`
	Required   bool   `json:"required,omitempty"`
	Exclude    bool   `json:"exclude,omitempty"`
	Additional bool   `json:"additional,omitempty"`
}
type ServiceProfile struct {
	//ValueObject
	ApiId    string `json:"apiId"`
	BasePath string `json:"basePath"`
}
type CACert struct {
	//ValueObject
	CertBlob           string `json:"certBlob"`
	Name               string `json:"name"`
	Alias              string `json:"alias"`
	Subject            string `json:"subject"`
	Issuer             string `json:"issuer"`
	Version            int    `json:"version"`
	NotValidBefore     int    `json:"notValidBefore"`
	NotValidAfter      int    `json:"notValidAfter"`
	SignatureAlgorithm string `json:"signatureAlgorithm"`
	Sha1Fingerprint    string `json:"sha1Fingerprint"`
	Md5Fingerprint     string `json:"md5Fingerprint"`
	Expired            bool   `json:"expired"`
	NotYetValid        bool   `json:"notYetValid"`
	Inbound            bool   `json:"inbound"`
	Outbound           bool   `json:"outbound"`
}

//---
type Backend struct {
	Id                    string                 `json:"id,omitempty"`
	Name                  string                 `json:"name,omitempty"`
	Summary               string                 `json:"summary,omitempty"`
	Description           string                 `json:"description,omitempty"`
	Version               string                 `json:"version,omitempty"`
	BasePath              string                 `json:"basePath,omitempty"`
	ResourcePath          string                 `json:"resourcePath,omitempty"`
	Consumes              []string               `json:"consumes,omitempty"`
	Produces              []string               `json:"produces,omitempty"`
	Integral              bool                   `json:"integral,omitempty"`
	CreatedOn             int                    `json:"createdOn,omitempty"`
	CreatedBy             string                 `json:"createdBy,omitempty"`
	OrganizationId        string                 `json:"organizationId,omitempty"`
	ServiceType           string                 `json:"serviceType,omitempty"`
	HasOriginalDefinition bool                   `json:"hasOriginalDefinition,omitempty"`
	ImportUrl             string                 `json:"importUrl,omitempty"`
	Properties            map[string]interface{} `json:"properties,omitempty"`
	Models                map[string]interface{} `json:"models,omitempty"`
}

// Config -
type Config struct {
	RegistrationEnabled               bool            `json:"registrationEnabled,omitempty"`
	RegTokenEmailEnabled              bool            `json:"regTokenEmailEnabled,omitempty"`
	ApiImportTimeout                  int             `json:"apiImportTimeout,omitempty"`
	IsTrial                           bool            `json:"isTrial,omitempty"`
	PromoteApiViaPolicy               bool            `json:"promoteApiViaPolicy,omitempty"`
	SystemOAuthScopesEnabled          bool            `json:"systemOAuthScopesEnabled,omitempty"`
	OadminSelfServiceEnabled          bool            `json:"oadminSelfServiceEnabled,omitempty"`
	ProductVersion                    string          `json:"productVersion,omitempty"`
	PortalName                        string          `json:"portalName,omitempty"`
	GlobalResponsePolicy              string          `json:"globalResponsePolicy,omitempty"`
	AutoApproveApplications           bool            `json:"autoApproveApplications,omitempty"`
	GlobalRequestPolicy               string          `json:"globalRequestPolicy,omitempty"`
	AutoApproveUserRegistration       bool            `json:"autoApproveUserRegistration,omitempty"`
	DelegateApplicationAdministration bool            `json:"delegateApplicationAdministration,omitempty"`
	ApiDefaultVirtualHost             string          `json:"apiDefaultVirtualHost,omitempty"`
	ApiRoutingKeyLocation             string          `json:"apiRoutingKeyLocation,omitempty"`
	ApplicationScopeRestrictions      bool            `json:"applicationScopeRestrictions,omitempty"`
	BaseOAuth                         bool            `json:"baseOAuth,omitempty"`
	EmailBounceAddress                string          `json:"emailBounceAddress,omitempty"`
	AdvisoryBannerEnabled             bool            `json:"advisoryBannerEnabled,omitempty"`
	UserNameRegex                     string          `json:"userNameRegex,omitempty"`
	ApiImportMimeValidation           bool            `json:"apiImportMimeValidation,omitempty"`
	SessionIdleTimeout                int             `json:"sessionIdleTimeout,omitempty"`
	IsApiPortalConfigured             bool            `json:"isApiPortalConfigured,omitempty"`
	ChangePasswordOnFirstLogin        bool            `json:"changePasswordOnFirstLogin,omitempty"`
	SessionTimeout                    int             `json:"sessionTimeout,omitempty"`
	EmailFrom                         string          `json:"emailFrom,omitempty"`
	ApiRoutingKeyEnabled              bool            `json:"apiRoutingKeyEnabled,omitempty"`
	LoginResponseTime                 int             `json:"loginResponseTime,omitempty"`
	ServerCertificateVerification     bool            `json:"serverCertificateVerification,omitempty"`
	ResetPasswordEnabled              bool            `json:"resetPasswordEnabled,omitempty"`
	AdvisoryBannerText                string          `json:"advisoryBannerText,omitempty"`
	ApiImportEditable                 bool            `json:"apiImportEditable,omitempty"`
	ApiPortalHostname                 string          `json:"apiPortalHostname,omitempty"`
	ApiPortalName                     string          `json:"apiPortalName,omitempty"`
	FaultHandlersEnabled              bool            `json:"faultHandlersEnabled,omitempty"`
	LockUserAccount                   LockUserAccount `json:"lockUserAccount"`
	Architecture                      string          `json:"architecture,omitempty"`
	StrictCertificateChecking         bool            `json:"strictCertificateChecking,omitempty"`
	GlobalPoliciesEnabled             bool            `json:"globalPoliciesEnabled,omitempty"`
	MinimumPasswordLength             int             `json:"minimumPasswordLength,omitempty"`
	PasswordExpiryEnabled             bool            `json:"passwordExpiryEnabled,omitempty"`
	Os                                string          `json:"os,omitempty"`
	LoginNameRegex                    string          `json:"loginNameRegex,omitempty"`
	DefaultTrialDuration              int             `json:"defaultTrialDuration,omitempty"`
	GlobalFaultHandlerPolicy          string          `json:"globalFaultHandlerPolicy,omitempty"`
	PasswordLifetimeDays              int             `json:"passwordLifetimeDays,omitempty"`
	DelegateUserAdministration        bool            `json:"delegateUserAdministration,omitempty"`
	PortalHostname                    string          `json:"portalHostname,omitempty"`
}

type LockUserAccount struct {
	Enabled            bool   `json:"enabled,omitempty"`
	Attempts           int    `json:"attempts,omitempty"`
	TimePeriod         int    `json:"timePeriod,omitempty"`
	TimePeriodUnit     string `json:"timePeriodUnit,omitempty"`
	LockTimePeriod     int    `json:"lockTimePeriod,omitempty"`
	LockTimePeriodUnit string `json:"lockTimePeriodUnit,omitempty"`
}

type Org struct {
	Id             string `json:"id,omitempty"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Email          string `json:"email"`
	Image          string `json:"image"`
	Restricted     bool   `json:"restricted,omitempty"`
	VirtualHost    string `json:"virtualHost"`
	Phone          string `json:"phone"`
	Enabled        bool   `json:"enabled"`
	Development    bool   `json:"development"`
	Dn             string `json:"dn,omitempty"`
	CreatedOn      int    `json:"createdOn,omitempty"`
	StartTrialDate int    `json:"startTrialDate,omitempty"`
	EndTrialDate   int    `json:"endTrialDate,omitempty"`
	TrialDuration  int    `json:"trialDuration,omitempty"`
	IsTrial        bool   `json:"isTrial,omitempty"`
}
type User struct {
	Id             string                 `json:"id,omitempty"`
	OrganizationId string                 `json:"organizationId"`
	Name           string                 `json:"name"`
	Description    string                 `json:"description,omitempty"`
	LoginName      string                 `json:"loginName"`
	Email          string                 `json:"email,omitempty"`
	Phone          string                 `json:"phone,omitempty"`
	Mobile         string                 `json:"mobile,omitempty"`
	Role           string                 `json:"role"`
	Image          string                 `json:"image,omitempty"`
	Enabled        bool                   `json:"enabled"`
	CreatedOn      int                    `json:"createdOn,omitempty"`
	State          string                 `json:"state,omitempty"`
	Type           string                 `json:"type,omitempty"`
	Orgs2Role      map[string]string      `json:"orgs2Role,omitempty"`
	AuthAttrs      map[string]interface{} `json:"authAttrs,omitempty"`
	Dn             string                 `json:"dn,omitempty"`
}
