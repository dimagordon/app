package http

type SignInForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpForm struct {
	Username  string `json:"username"`
	Password1 string `json:"password1"`
	Password2 string `json:"password2"`
}
