package model

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
)

// Appeal struct that defines what data a Notice should receive
/* This is subject to change once I add a appeal model */
type Appeal struct {
	ID           string
	FirstName    string
	LastName     string
	Address      string
	City         string
	State        string
	Zip          string
	AppealNumber string
	Comment      template.HTML
}

// GetAppealByID takes in an integer and returns a single result from the
// database
func GetAppealByID(id int) (*Appeal, error) {
	result := &Appeal{}
	row := db.QueryRow(`
		SELECT CC_Person_ID, FNAME, LNAME, Address, City, State, Zip, Appeals_ID
		FROM dbo.tblAppealData
		WHERE CC_Person_ID = $1`, id)
	err := row.Scan(&result.ID, &result.FirstName, &result.LastName, &result.Address, &result.City, &result.State, &result.Zip, &result.AppealNumber)
	switch {
	case err == sql.ErrNoRows:
		return result, fmt.Errorf("Not an Appeal")
	case err != nil:
		return result, err
	}
	return result, nil
}

func ShowAllAppeals() []*Appeal {
	result := []*Appeal{}

	rows, err := db.Query(`
		SELECT CC_Person_ID, FNAME, LNAME, Address, City, State, Zip, Appeals_ID
		FROM dbo.tblAppealData
		ORDER BY CC_Person_ID desc`)
	if err != nil {
		log.Println(err)
		return result
	}
	defer rows.Close()
	for rows.Next() {
		a := &Appeal{}
		err := rows.Scan(&a.ID, &a.FirstName, &a.LastName, &a.Address, &a.City, &a.State, &a.Zip, &a.AppealNumber)
		if err != nil {
			log.Println(err)
			return result
		}
		result = append(result, a)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return result
	}
	return result
}
