package auth

import (
	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func ValidateLogin(user, password string) bool {
	store, err := ioutil.ReadFile("user.pass")
	if err != nil {
		log.Println(err)
		return false
	}

	userData := strings.Split(string(store), ":")

	if user != userData[0] {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData[1]), []byte(password))
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
