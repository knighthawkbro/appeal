package model

import (
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	ldap "gopkg.in/ldap.v2"
)

type User struct {
	ID        int
	Email     string
	username  string
	FirstName string
	LastName  string
	LastLogin *time.Time
}

// Need a read-only user to get attributes back from ldap search
const bindusername = "CCA-ReadOnly@msd.govt.state.ma.us"
const bindpassword = "Password1"

const servername = "MSD-INF-SDC-001.msd.govt.state.ma.us"
const port = 389

const baseDN = "ou=Users,ou=CCA,dc=msd,dc=govt,dc=state,dc=ma,dc=us"
const adminsDN = "CN=Administrators (CCA),OU=Groups,OU=CCA,DC=msd,DC=govt,DC=state,DC=ma,DC=us"
const appealsDN = "CN=Appeals (CCA),OU=Groups,OU=CCA,DC=msd,DC=govt,DC=state,DC=ma,DC=us"

func Login(username, password string) (*User, error) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", servername, port))
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to LDAP Server")
	}
	defer l.Close()
	l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err = l.Bind(bindusername, bindpassword); err != nil {
		return nil, fmt.Errorf("Failed to bind Read-Only User")
	}
	Groups := fmt.Sprintf("(|(memberof=%v)(memberof=%v))",
		ldap.EscapeFilter(adminsDN),
		ldap.EscapeFilter(appealsDN))
	usedUsername := username
	if username[0:3] == "msd" {
		username = username[4:]
	}
	if strings.Contains(username, "@state.ma.us") {
		x := strings.Split(username, "@")
		username = x[0] + "@massmail.state.ma.us"
	}
	username = ldap.EscapeFilter(username)
	searchRequest := ldap.NewSearchRequest(
		baseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=user)(|(sAMAccountName=%s)(mail=%s))%s)", username, username, Groups),
		[]string{"dn", "mail", "givenName", "sn"}, nil,
	)
	sr, err := l.Search(searchRequest)
	if err != nil {
		return nil, err
	}
	if len(sr.Entries) != 1 {
		return nil, fmt.Errorf("User not found, Contact your Admin if this is incorrect")
	}
	err = l.Bind(sr.Entries[0].DN, password)
	if err != nil {
		return nil, fmt.Errorf("Password incorrect")
	}
	var result = &User{
		Email:     sr.Entries[0].GetAttributeValue("mail"),
		FirstName: sr.Entries[0].GetAttributeValue("givenName"),
		LastName:  sr.Entries[0].GetAttributeValue("sn"),
	}
	t := time.Now()
	_, err = db.Exec(`
		INSERT INTO dbo.tblUserLog ( email, username, firstname, lastname, lastlogin )
		VALUES ( $1, $2, $3, $4, $5 )`,
		result.Email, usedUsername, result.FirstName, result.LastName, t)

	if err != nil {
		return nil, fmt.Errorf("Could not insert new record into database, Error: %v", err)
	}

	return result, nil
}
