package view

import (
	"appeals/app/model"
)

type Home struct {
	Title string
	//Active  string
	Appeals []model.Appeal
}

func NewHome() Home {
	result := Home{
		//Active: "home",
		Title:   "Appeals",
		Appeals: model.ShowAllAppeals(),
	}
	return result
}
