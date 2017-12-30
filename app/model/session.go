package model

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    string
	CreatedAt int64
}

func (user *User) CreateSession() (Session, error) {
	result := Session{
		UUID:      createUUID(),
		Email:     user.Email,
		UserID:    user.Username,
		CreatedAt: time.Now().Unix(),
	}
	_, err := db.Exec(`
		INSERT INTO dbo.tblSessions ( uuid, email, user_id, created_at )
		VALUES ( $1, $2, $3, $4 )`,
		result.UUID, result.Email, result.UserID, result.CreatedAt)
	if err != nil {
		return Session{}, err
	}

	return result, nil
}

func DestroySession(uuid string) error {
	_, err := db.Exec(`
		DELETE FROM dbo.tblSessions
		WHERE uuid = $1`, uuid)
	if err != nil {
		return err
	}
	return nil
}

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

// Session.Check is a per request check on the database to see if a session is still active,
// It kills expired sessions and updates
func (s *Session) Check() (bool, error) {
	row := db.QueryRow(`
		SELECT id, uuid, email, user_id, created_at 
		FROM dbo.tblSessions 
		WHERE uuid = $1`, s.UUID)
	err := row.Scan(&s.ID, &s.UUID, &s.Email, &s.UserID, &s.CreatedAt)
	switch {
	case err == sql.ErrNoRows:
		return false, fmt.Errorf("Session doesn't exist")
	case err != nil:
		return false, err
	}
	if s.CreatedAt > (time.Now().Unix() + 3600) {
		_, err := db.Exec(`
			UPDATE dbo.tblSessions
			SET created_at = $1
			WHERE uuid = $2`, time.Now().Unix(), s.UUID)
		if err != nil {
			return true, err
		}
	}
	DestroyExpired()
	return true, nil
}

func DestroyExpired() {

	_, err := db.Exec(`
		DELETE FROM dbo.tblSessions
		WHERE created_at < $1`, time.Now().Unix()-3600)
	if err != nil {
		log.Println(err)
	}
}
