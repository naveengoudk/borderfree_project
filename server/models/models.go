package models

type Product struct {
	User        string `json:"user"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
}
type Users struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
