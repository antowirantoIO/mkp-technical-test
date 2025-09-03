package model

type Auth struct {
	// Login user id
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
