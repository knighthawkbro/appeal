package model

import (
	"fmt"
)

// GetAppealByID takes in an integer and returns a single result from the
// database
func GetAppealByID(id int) (Appeal, error) {
	r := Appeal{}
	db.Debug().Preload("Appellant").Preload("Appellant.Address").Find(&r, "ID = ?", id)
	if r.ID == 0 {
		return r, fmt.Errorf("Not an Appeal")
	}
	fmt.Println(r)
	return r, nil
}

// ShowAllAppeals (Public) -
func ShowAllAppeals() []Appeal {
	appeals := []Appeal{}
	db.Order("ID desc").Limit("10").Preload("Appellant").
		Preload("Appellant.Address").Find(&appeals)
	return appeals
}
