package axwayapi

type Config struct {
	RegistrationEnabled               bool            `json:"registrationEnabled"`
	RegTokenEmailEnabled              bool            `json:"regTokenEmailEnabled"`
	ApiImportTimeout                  int             `json:"apiImportTimeout,omitempty"`
	IsTrial                           bool            `json:"isTrial"`
	PromoteApiViaPolicy               bool            `json:"promoteApiViaPolicy"`
	SystemOAuthScopesEnabled          bool            `json:"systemOAuthScopesEnabled"`
	OadminSelfServiceEnabled          bool            `json:"oadminSelfServiceEnabled"`
	ProductVersion                    string          `json:"productVersion,omitempty"`
	PortalName                        string          `json:"portalName,omitempty"`
	GlobalResponsePolicy              string          `json:"globalResponsePolicy,omitempty"`
	AutoApproveApplications           bool            `json:"autoApproveApplications"`
	GlobalRequestPolicy               string          `json:"globalRequestPolicy,omitempty"`
	AutoApproveUserRegistration       bool            `json:"autoApproveUserRegistration"`
	DelegateApplicationAdministration bool            `json:"delegateApplicationAdministration"`
	ApiDefaultVirtualHost             string          `json:"apiDefaultVirtualHost,omitempty"`
	ApiRoutingKeyLocation             string          `json:"apiRoutingKeyLocation,omitempty"`
	ApplicationScopeRestrictions      bool            `json:"applicationScopeRestrictions"`
	BaseOAuth                         bool            `json:"baseOAuth"`
	EmailBounceAddress                string          `json:"emailBounceAddress,omitempty"`
	AdvisoryBannerEnabled             bool            `json:"advisoryBannerEnabled"`
	UserNameRegex                     string          `json:"userNameRegex,omitempty"`
	ApiImportMimeValidation           bool            `json:"apiImportMimeValidation"`
	SessionIdleTimeout                int             `json:"sessionIdleTimeout,omitempty"`
	IsApiPortalConfigured             bool            `json:"isApiPortalConfigured"`
	ChangePasswordOnFirstLogin        bool            `json:"changePasswordOnFirstLogin"`
	SessionTimeout                    int             `json:"sessionTimeout,omitempty"`
	EmailFrom                         string          `json:"emailFrom,omitempty"`
	ApiRoutingKeyEnabled              bool            `json:"apiRoutingKeyEnabled"`
	LoginResponseTime                 int             `json:"loginResponseTime,omitempty"`
	ServerCertificateVerification     bool            `json:"serverCertificateVerification"`
	ResetPasswordEnabled              bool            `json:"resetPasswordEnabled"`
	AdvisoryBannerText                string          `json:"advisoryBannerText,omitempty"`
	ApiImportEditable                 bool            `json:"apiImportEditable"`
	ApiPortalHostname                 string          `json:"apiPortalHostname,omitempty"`
	ApiPortalName                     string          `json:"apiPortalName,omitempty"`
	FaultHandlersEnabled              bool            `json:"faultHandlersEnabled"`
	LockUserAccount                   LockUserAccount `json:"lockUserAccount"`
	Architecture                      string          `json:"architecture,omitempty"`
	StrictCertificateChecking         bool            `json:"strictCertificateChecking"`
	GlobalPoliciesEnabled             bool            `json:"globalPoliciesEnabled"`
	MinimumPasswordLength             int             `json:"minimumPasswordLength,omitempty"`
	PasswordExpiryEnabled             bool            `json:"passwordExpiryEnabled"`
	Os                                string          `json:"os,omitempty"`
	LoginNameRegex                    string          `json:"loginNameRegex,omitempty"`
	DefaultTrialDuration              int             `json:"defaultTrialDuration,omitempty"`
	GlobalFaultHandlerPolicy          string          `json:"globalFaultHandlerPolicy,omitempty"`
	PasswordLifetimeDays              int             `json:"passwordLifetimeDays,omitempty"`
	DelegateUserAdministration        bool            `json:"delegateUserAdministration"`
	PortalHostname                    string          `json:"portalHostname,omitempty"`
}

type LockUserAccount struct {
	Enabled            bool   `json:"enabled"`
	Attempts           int    `json:"attempts,omitempty"`
	TimePeriod         int    `json:"timePeriod,omitempty"`
	TimePeriodUnit     string `json:"timePeriodUnit,omitempty"`
	LockTimePeriod     int    `json:"lockTimePeriod,omitempty"`
	LockTimePeriodUnit string `json:"lockTimePeriodUnit,omitempty"`
}

// GetConfig - Returns a specifc Config
func (c *Client) GetConfig() (ret *Config, err error) {
	ret = &Config{}
	err = c.get(ret, "config", 200)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// UpdateConfig - Updates an Config
func (c *Client) UpdateConfig(config *Config) (error) {
	return c.put(config, "config")
}
