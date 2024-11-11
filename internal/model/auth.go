package model

type RegisterUser struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
	DOB       string
}

type Login struct {
	Email    string
	Password string
}

type AuthResponse struct {
	User        *User  `json:"user"`
	AccessToken string `json:"access_token"`
}
