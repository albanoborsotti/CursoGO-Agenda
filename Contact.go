package main

type Contact struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Adress   string `json:"adress"`
	Email    string `json:"email"`
}

func getEmail(contac *Contact) string {
	return contac.Email
}
