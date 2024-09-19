package vayanaTypes

type AuthRequest struct {
	Handle              string `json:"handle"`
	HandleType          string `json:"handleType"`
	Password            string `json:"password"`
	Otp                 string `json:"otp,omitempty"`
	TokenDurationInMins int    `json:"tokenDurationInMins"` //Max 360
}

type AuthResponse struct {
	Data  AuthResponseData `json:"data"`
	Error *Error           `json:"error"`
}

type AuthResponseData struct {
	Token                string               `json:"token"`
	UserId               string               `json:"userId"`
	Expiry               int64                `json:"expiry"`
	Email                string               `json:"email"`
	Mobile               string               `json:"mobile"`
	GivenName            string               `json:"givenName"`
	LastName             string               `json:"lastName"`
	State                string               `json:"state"`
	VerificationStatus   string               `json:"verificationStatus"`
	PasswordLastModified int64                `json:"passwordLastModified"`
	AssociatedOrgs       []OrganisationAccess `json:"associatedOrgs"`
	Timezone             string               `json:"timezone"`
	Locale               string               `json:"locale"`
	CreatedOn            int64                `json:"createdOn"`
	LastUpdated          int64                `json:"lastUpdated"`
}

type OrganisationAccess struct {
	UserAccessInfo UserAccess   `json:"userAccessInfo"`
	Organisation   Organization `json:"organisation"`
}

type UserAccess struct {
	Primary bool `json:"primary"`
	Admin   bool `json:"admin"`
}

type Organization struct {
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	TaxIdentifier     string    `json:"taxIdentifier"`
	TaxIdentifierType string    `json:"taxIdentifierType"`
	Country           string    `json:"country"`
	Services          []Service `json:"services"`
}

type Service struct {
	ServiceCode string `json:"serviceCode"`
	ServiceName string `json:"serviceName"`
}
