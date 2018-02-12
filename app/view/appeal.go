package view

import (
	"appeals/app/model"
	"fmt"
)

type Appeal struct {
	Title  string
	Appeal *model.Appeal
	URL    string
}

func NewAppeal(a *model.Appeal) Appeal {
	result := Appeal{
		URL:    fmt.Sprintf("/appeal/%v", a.ID),
		Title:  fmt.Sprintf("Appeals - %v", a.AppealNumber),
		Appeal: a,
	}
	return result
}
