package model

type User struct {
	ID             string  `json:"id"`
	FirstName      string  `json:"firstname"`
	LastName       string  `json:"lastname"`
	FavouriteMovies []Movie `json:"favouriteMovies"`
}
