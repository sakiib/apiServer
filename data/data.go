package data

import "github.com/sakiib/apiServer/model"

var Users = []model.User{
	model.User{
		ID: "1",
		FirstName: "sakib",
		LastName: "alamin",
		FavouriteMovie: &model.Movie{
			ID: "1",
			Title: "Movie-1",
			Genre: "Comedy",
			Rating: 5,
		},
	},
	model.User{
		ID: "2",
		FirstName: "prangon",
		LastName: "majumder",
		FavouriteMovie: &model.Movie{
			ID: "2",
			Title: "Movie-2",
			Genre: "Sci-Fi",
			Rating: 5,
		},
	},
	model.User{
		ID: "3",
		FirstName: "mehedi",
		LastName: "hasan",
		FavouriteMovie: &model.Movie{
			ID: "3",
			Title: "Movie-3",
			Genre: "Horror",
			Rating: 3,
		},
	},
	model.User{
		ID: "4",
		FirstName: "sahadat",
		LastName: "hossain",
		FavouriteMovie: &model.Movie{
			ID: "4",
			Title: "Movie-4",
			Genre: "Action",
			Rating: 4,
		},
	},
}
