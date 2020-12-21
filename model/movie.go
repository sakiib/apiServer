package model

type Movie struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
	Rating int `json:"rating"`
}