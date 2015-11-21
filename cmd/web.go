package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/gorilla/mux"
	"github.com/hunterpraska/Whiteboard/routers"
	"log"
	"net/http"
)

var CmdWeb = cli.Command{
	Name:  "web",
	Usage: "Start Whiteboard server",
	Description: `Whiteboard server is the only thing you need to
	run, and it takes care of everything else for you`,
	Action: runWeb,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "port, p",
			Value: "8080",
			Usage: "Temporary port number to prevent conflict"},
	},
}

const PORT = ":8080"

func runWeb(ctx *cli.Context) {
	// File server for static files: CSS, JS, etc.
	staticFiles := http.FileServer(http.Dir("resources"))

	webRouter := mux.NewRouter()

	// Application routes
	webRouter.HandleFunc("/", routers.HomeHandler).Methods("GET")
	webRouter.HandleFunc("/login", routers.LoginHandler).Methods("GET", "POST")
	webRouter.HandleFunc("/auth-check", routers.AuthCheck).Methods("GET")

	http.Handle("/r/", http.StripPrefix("/r/", staticFiles))
	http.Handle("/", webRouter)

	// Serve pages over HTTP on PORT
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
