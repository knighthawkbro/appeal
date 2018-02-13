package model

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

// Session (Public) -
type Session struct {
	gorm.Model

	UUID   string
	Email  string
	UserID string
}

// CreateSession (Public) -
func (user *User) CreateSession() (Session, error) {
	result := Session{
		UUID:   createUUID(),
		Email:  user.Email,
		UserID: user.Username,
	}
	db.Create(&result)
	return result, nil
}

// DestroySession (Public) -
func DestroySession(uuid string) error {
	session := Session{}
	db.Where("uuid = ?", uuid).First(&session)
	if session.ID == 0 {
		return fmt.Errorf("Could not the record")
	}
	db.Delete(&session)
	return nil
}

// createUUID (Private) -
func createUUID() string {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

// Sessions (Public) -
func Sessions(w http.ResponseWriter, r *http.Request) (Session, error) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		return Session{}, err
	}
	result := Session{UUID: cookie.Value}
	if ok, _ := result.Check(); !ok {
		return result, fmt.Errorf("Invalid session")
	}
	return result, nil
}

// Check (Public) - is a per request check on the database to see if a session is still active,
// It kills expired sessions and updates
func (s *Session) Check() (bool, error) {
	row := Session{}
	db.Find(&row, "uuid = ?", s.UUID)
	if row.UUID == "" {
		return false, fmt.Errorf("Session doesn't exist")
	}
	if row.CreatedAt.Unix() > (time.Now().Unix() + 3600) {
		db.Model(&row).Update("created_at", time.Now())
	}
	DestroyExpired()
	return true, nil
}

// DestroyExpired (Public) -  Need to fix this function
func DestroyExpired() {
	sessions := []Session{}
	db.Find(&sessions)
	for _, sess := range sessions {
		if sess.CreatedAt.Unix() < time.Now().Unix()-3600 {
			db.Delete(&sess)
		}
	}
}
