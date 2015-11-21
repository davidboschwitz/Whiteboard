package router

import (
	"github.com/gorilla/securecookie"
	"github.com/hunterpraska/Whiteboard/auth"
	"html/template"
	"log"
	"net/http"
)

// Secure Cookie variables
var hashKey = securecookie.GenerateRandomKey(32)
var blockKey = securecookie.GenerateRandomKey(32)
var s = securecookie.New(hashKey, blockKey)

// Handles home page requests
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/home.html")
	if err != nil {
		log.Println(err)
		return
	}
	t.Execute(w, nil)
}

// Handles user login. If user is logged in, redirects to '/'.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if auth.LoggedIn(w, r, s) {
		http.Redirect(w, r, "/auth-check", 302)
		return
	}

	if r.Method == "GET" {
		t, err := template.ParseFiles("views/login.html")
		if err != nil {
			log.Println(err)
			return
		}
		t.Execute(w, nil)
	} else {
		// Get values from html form
		user := r.FormValue("user")
		pass := r.FormValue("password")

		// Attempt to validate user, if incorrect info, send user back to login page
		if auth.ValidateLogin(user, pass) {
			// Create secure cookie with login info
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

// Test of user authentication. Redirects user to login page if not logged in.
func AuthCheck(w http.ResponseWriter, r *http.Request) {
	if !auth.LoggedIn(w, r, s) {
		http.Redirect(w, r, "/login", 302)
		return
	}

	t, err := template.ParseFiles("views/authCheck.html")
	if err != nil {
		log.Println(err)
		return
	}
	t.Execute(w, nil)
}
