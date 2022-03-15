package axwayapi

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

// Order -
//type Config struct {
//	registrationEnabled    bool         `json:"registrationEnabled,omitempty"`
//	Items []ConfigItem `json:"items,omitempty"`
//}

// OrderItem -
//type ConfigItem struct {
//	Coffee   Coffee `json:"coffee"`
//	Quantity int    `json:"quantity"`
//}

// Coffee -
type Coffee struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Teaser      string             `json:"teaser"`
	Collection  string             `json:"collection"`
	Origin      string             `json:"origin"`
	Color       string             `json:"color"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Image       string             `json:"image"`
	Ingredient  []CoffeeIngredient `json:"ingredients"`
}

// Ingredient -
type CoffeeIngredient struct {
	ID       int    `json:"ingredient_id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

// Ingredient -
type Ingredient struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}
