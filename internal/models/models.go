package models


// Config structure
type Config struct{
	BorderColor 	string `json:"border_color"`
	BackgroundColor string `json:"background_color"`
	
}


// Xkcd Json Structure
type XkcdJsonStruct struct {
	Year  string `json:"year"`
	Month string `json:"month"`
	Day   string `json:"day"`
	Title string `json:"title"`
	Alt   string `json:"alt"`
	Img   string `json:"img"`
	Num   int    `json:"num"`
}