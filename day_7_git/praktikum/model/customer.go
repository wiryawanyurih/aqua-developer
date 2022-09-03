package model

type customer struct {
	ID    int    `json:"id"`
	name  string `json:"name"`
	Email string `json:"email"`
}
