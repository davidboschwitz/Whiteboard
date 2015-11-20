package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	PORT = ":8080"
)

var hashKey = securecookie.GenerateRandomKey(32)
var blockKey = securecookie.GenerateRandomKey(32)
var s = securecookie.New(hashKey, blockKey)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/login", LoginHandler).Methods("GET", "POST")
	router.HandleFunc("/auth-check", AuthCheck).Methods("GET")

	http.Handle("/", router)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./views/home.html")
	if err != nil {
		log.Println(err)
		return
	}
	t.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if loggedIn(w, r) {
		http.Redirect(w, r, "/auth-check", 302)
		return
	}

	if r.Method == "GET" {
		t, err := template.ParseFiles("./views/login.html")
		if err != nil {
			log.Println(err)
			return
		}
		t.Execute(w, nil)
	} else {
		user := r.FormValue("user")
		pass := r.FormValue("password")
		if validateLogin(user, pass) {
			value := map[string]string{
				"authenticated": "true",
			}
			if encoded, err := s.Encode("whiteboard", value); err == nil {
				cookie := &http.Cookie{
					Name:  "whiteboard",
					Value: encoded,
					Path:  "/",
				}
				http.SetCookie(w, cookie)
			} else {
				log.Println(err)
				http.Redirect(w, r, "/login", 302)
				return
			}
			http.Redirect(w, r, "/auth-check", 302)
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	}
}

func AuthCheck(w http.ResponseWriter, r *http.Request) {
	if !loggedIn(w, r) {
		http.Redirect(w, r, "/login", 302)
		return
	}

	t, err := template.ParseFiles("./views/authCheck.html")
	if err != nil {
		log.Println(err)
		return
	}
	t.Execute(w, nil)
}

func validateLogin(user, password string) bool {
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

func loggedIn(w http.ResponseWriter, r *http.Request) bool {
	if cookie, err := r.Cookie("whiteboard"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("whiteboard", cookie.Value, &value); err == nil {
			return true
		}
		return false
	}
	return false
}
