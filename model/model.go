package model

type Man struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Height    int    `json:"height"`
	Weight    int    `json:"weight"`
}
