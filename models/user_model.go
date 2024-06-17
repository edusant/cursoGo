package models

type User struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// Dummy data for demonstration
func GetUsers() []User {
	users := []User{
		{ID: 1, Nome: "Alice"},
		{ID: 2, Nome: "Bob"},
		{ID: 3, Nome: "Charlie"},
	}
	return users
}
