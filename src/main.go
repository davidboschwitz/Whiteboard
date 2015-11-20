package main

import (
	"./auth"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"html/template"
	"log"
	"net/http"
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

////////////////////////////////////////////////////////////////////////////////
// Routes
////////////////////////////////////////////////////////////////////////////////
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./views/home.html")
	if err != nil {
		log.Println(err)
		return
	}
	t.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if auth.LoggedIn(w, r, s) {
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
		if auth.ValidateLogin(user, pass) {
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
	if !auth.LoggedIn(w, r, s) {
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
