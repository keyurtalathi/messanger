package main

import (
	"Messanger/apigateway/core"
	"Messanger/apigateway/servicehandlers"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var addr = flag.String("addr", ":1234", "http service address")

func main() {
	flag.Parse()
	p := servicehandlers.PingHandler{}
	a := servicehandlers.AuthenticateHandler{}
	ch := servicehandlers.ChatHandler{}
	c := servicehandlers.ContactHandler{}
	s := servicehandlers.SignupHandler{}
	g := servicehandlers.GroupHandler{}
	u := servicehandlers.UploadHandler{}
	hub := core.NewHub()
	go hub.Run()
	go hub.Run()
	go hub.Run()

	// mux := http.NewServeMux()

	router := mux.NewRouter()
	router.Handle("/api/ping", p)
	router.Handle("/api/upload", u)
	router.Handle("/api/authenticate", a)
	router.Handle("/api/contactlist", c)
	router.Handle("/api/signup", s)
	router.Handle("/api/chat", ch)
	router.Handle("/api/group", g)
	router.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		servicehandlers.Echo(hub, w, r)
	})
	// handlr := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(*addr,
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(router)))
	// x := http.ListenAndServe(*addr, handlr)
	// log.Fatal(x)
}
