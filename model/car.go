package model

type Car struct {
	ID      int    `json:"id"`
	RegNum  string `json:"regNum"`
	Make    string `json:"make"`
	Model   string `json:"model"`
	Year    int    `json:"year"`
	OwnerID int    `json:"ownerID"` // Change type to int
}
