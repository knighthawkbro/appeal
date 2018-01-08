package view

import (
	"appeals/app/model"
)

type Search struct {
	Title string
	//Active       string
	Page         int
	ItemsPerPage int
	Warning      string
	Appeal       *model.Appeal
}

func NewSearch() Search {
	result := Search{
		//Active: "search",
		Title: "Appeals - Search",
	}
	return result
}
