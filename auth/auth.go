package auth

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func ValidateLogin(user, password string, db *bolt.DB) bool {
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("users"))
		if bucket == nil {
			return fmt.Errorf("Bucket pastes not found!")
		}

		if passwordCrypt = bucket.Get([]byte(user)); err == nil {
			return fmt.Errorf("User not found!")
		}

		err = bcrypt.CompareHashAndPassword(passwordCrypt, []byte(password))
		return err
	})
	if err != nil {
		return false
	}
	return true
}

func LoggedIn(w http.ResponseWriter, r *http.Request, s *securecookie.SecureCookie) bool {
	if cookie, err := r.Cookie("whiteboard"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("whiteboard", cookie.Value, &value); err == nil {
			return true
		}
		return false
	}
	return false
}
