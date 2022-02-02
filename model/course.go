package model

// Definite The User
type User struct {
	Name         string `json:"name"`
	Password     string `json:"password"`
	Address      string `json:"address"`
	PhoneAddress string `json:"phone_address"`
	Age          int    `json:"age"`
}

// Course Maybe Related To User
type Course struct {
	Name   string `json:"name"`
	Price  int64  `json:"price"`
	Author User   `json:"author"`
}
