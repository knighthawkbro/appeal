package model

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
)

// Appeal struct that defines what data a Notice should receive
type Appeal struct {
	ID                    string
	Person                person
	AppealNumber          string
	IsOnline              bool
	NoticeDate            string
	NoticeNumber          string
	DateEntered           string
	DateReceived          string
	Reason                reason
	BusinessEvent         string
	IDR                   idr
	PublicMEC             bool
	publicMecText         string
	TaxFilingIssue        bool
	Expedite              bool
	AidPendingApplied     bool
	AidPendingAppliedDate string
	AidPendingRemoved     bool
	AidPendingRemovedDate string
	ReviewOutreach        string
	Hearing               bool
	HearingDate           string
	OutreachNotes         string
	DismissedInvalid      string
	Dismissed             string
	PendingInfo           bool
	WDDate                string
	DateFinalLetterSent   string
	LetterSentBy          string
	Comment               template.HTML
}

type hearing struct {
}

type address struct {
	Address string
	City    string
	State   string
	Zip     string
}

type person struct {
	FirstName string
	LastName  string
	MemberID  string
	Address   address
	Phone     string
	Email     string
	Rep       rep
}

type rep struct {
	FirstName string
	LastName  string
	Address   address
	phone     string
}

type reason struct {
	Income          bool
	Residency       bool
	OtherInsurance  bool
	LawfulPresence  bool
	Incarceration   bool
	PremiumWavier   bool
	OtherReason     bool
	OtherReasonText string
}

type idr struct {
	IDRIncome          bool
	IDRResidency       bool
	IDRSep             bool
	IDRESI             bool
	IDRLawfulPresence  bool
	IDRIncarceration   bool
	IDRPWDenial        bool
	IDROtherReason     bool
	IDROtherReasonText string
}

// GetAppealByID takes in an integer and returns a single result from the
// database
func GetAppealByID(id int) (*Appeal, error) {
	r := &Appeal{}
	row := db.QueryRow(`
		SELECT CC_Person_ID, FNAME, LNAME, Address, City, State, Zip, Appeals_ID
		FROM dbo.tblAppealData
		WHERE CC_Person_ID = $1`, id)
	err := row.Scan(&r.ID, &r.Person.FirstName, &r.Person.LastName, &r.Person.Address.Address,
		&r.Person.Address.City, &r.Person.Address.State, &r.Person.Address.Zip, &r.AppealNumber)
	switch {
	case err == sql.ErrNoRows:
		return r, fmt.Errorf("Not an Appeal")
	case err != nil:
		return r, err
	}
	return r, nil
}

func count() (count int) {
	row := db.QueryRow(`
		SELECT count(*)
		FROM dbo.tblAppealData`)
	row.Scan(&count)
	log.Println(count)
	return count
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
		err := rows.Scan(&a.ID, &a.Person.FirstName, &a.Person.LastName, &a.Person.Address.Address,
			&a.Person.Address.City, &a.Person.Address.State, &a.Person.Address.Zip, &a.AppealNumber)
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
