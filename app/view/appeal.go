package view

import (
	"appeals/app/model"
	"fmt"
	"strings"
	"time"
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
	DOB := strings.Split(result.Appeal.Appellant.DOB, " ")
	result.Appeal.Appellant.DOB = DOB[0]
	if len(result.Appeal.AppealID) == 0 && len(result.Appeal.EnrollmentYear) == 4 {
		result.Appeal.AppealID = fmt.Sprintf("ACA%v-%v", result.Appeal.EnrollmentYear[2:], result.Appeal.ID)
	}
	return result
}

// FormatTime (Public) -
func FormatTime(date time.Time, flag bool) string {
	if flag {
		return date.Format("01/02/06 03:04PM")
	}
	if date.IsZero() {
		return ""
	}
	return date.Format("01/02/06")
}
