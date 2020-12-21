package data

import "github.com/sakiib/apiServer/model"

var Users = []model.User{
	model.User{
		ID:        "1",
		FirstName: "sakib",
		LastName:  "alamin",
		FavouriteMovies: []model.Movie{
			model.Movie{ID: "1", Title: "movie-1", Genre: "rom-com", Rating: 5},
			model.Movie{ID: "2", Title: "movie-2", Genre: "action", Rating: 2},
		},
	},
	model.User{
		ID:        "2",
		FirstName: "prangon",
		LastName:  "majumder",
		FavouriteMovies: []model.Movie{
			model.Movie{ID: "3", Title: "movie-3", Genre: "horror", Rating: 3},
			model.Movie{ID: "4", Title: "movie-4", Genre: "action", Rating: 4},
		},
	},
	model.User{
		ID:        "3",
		FirstName: "mehedi",
		LastName:  "hasan",
		FavouriteMovies: []model.Movie{
			model.Movie{ID: "5", Title: "movie-5", Genre: "comedy", Rating: 5},
			model.Movie{ID: "6", Title: "movie-6", Genre: "sci-fi", Rating: 3},
		},
	},
	model.User{
		ID:        "4",
		FirstName: "sahadat",
		LastName:  "hossain",
		FavouriteMovies: []model.Movie{
			model.Movie{ID: "7", Title: "movie-7", Genre: "romantic", Rating: 3},
			model.Movie{ID: "8", Title: "movie-8", Genre: "action", Rating: 2},
		},
	},
}
