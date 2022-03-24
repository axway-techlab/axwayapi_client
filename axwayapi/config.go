package axwayapi

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

// GetConfig - Returns a specifc Config
func (c *Client) GetConfig() (ret *Config, err error) {
	ret = &Config{}
	err = c.get(ret, "config")
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UpdateConfig - Updates an Config
func (c *Client) UpdateConfig(config *Config) (error) {
	return c.put(config, "config")
}
