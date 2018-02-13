package view

import (
	"appeals/app/model"
	"fmt"
)

// Appeal (Public) -
type Appeal struct {
	Title  string
	Appeal model.Appeal
	URL    string
}

// NewAppeal (Public) -
func NewAppeal(a model.Appeal) Appeal {
	result := Appeal{
		URL:    fmt.Sprintf("/appeal/%v", a.ID),
		Title:  fmt.Sprintf("Appeals - %v", a.AppealID),
		Appeal: a,
	}
	return result
}
