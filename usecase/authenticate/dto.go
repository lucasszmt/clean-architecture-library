package authenticate

type AuthCredentialsDto struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
