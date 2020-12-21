package model

type User struct {
	ID             string `json:"id"`
	FirstName      string `json:"firstname"`
	LastName       string `json:"lastname"`
	FavouriteMovie *Movie `json:"favouriteMovie"`
}
