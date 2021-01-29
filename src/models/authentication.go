package models

// Authentication contains token and id of authenticated user.
type Authentication struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
