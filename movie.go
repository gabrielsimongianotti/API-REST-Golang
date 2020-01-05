package main

type Movie struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"Director"`
}

type Movies []Movie
