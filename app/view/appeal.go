package view

import (
	"appeals/app/model"
	"fmt"
)

type Appeal struct {
	Title string
	//Active  string
	Appeal *model.Appeal
	URL    string
}

func NewAppeal() Appeal {
	result := Appeal{}
	return result
}

func appealtoVM(c model.Appeal) Appeal {
	return Appeal{
		URL:   fmt.Sprintf("/shop/%v", c.ID),
		Title: fmt.Sprintf("Appeals - %v, %v", c.LastName, c.FirstName),
	}
}
