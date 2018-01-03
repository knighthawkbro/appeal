package view

import (
	"appeals/app/model"
	"fmt"
)

type Home struct {
	Title   string
	Active  string
	Appeals []*model.Appeal
}

type Appeal struct {
	URL   string
	Title string
}

func NewHome() Home {
	result := Home{
		Active:  "home",
		Title:   "Appeals",
		Appeals: model.ShowAllAppeals(),
	}
	return result
}

func appealtoVM(c model.Appeal) Appeal {
	return Appeal{
		URL:   fmt.Sprintf("/shop/%v", c.ID),
		Title: fmt.Sprintf("Appeals - %v %v", c.FirstName, c.LastName),
	}
}
