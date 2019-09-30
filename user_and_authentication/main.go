package main

import (
	"Messanger/user_and_authentication/servicehandlers"
	"flag"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	flag.Parse()
	p := servicehandlers.PingHandler{}
	a := servicehandlers.AuthenticateHandler{}
	c := servicehandlers.ContactHandler{}
	s := servicehandlers.SignupHandler{}
	g := servicehandlers.GroupHandler{}

	mux := http.NewServeMux()
	mux.Handle("/ping", p)
	mux.Handle("/authenticate", a)
	mux.Handle("/contactlist", c)
	mux.Handle("/signup", s)
	mux.Handle("/group", g)
	handlr := cors.Default().Handler(mux)
	x := http.ListenAndServe(":1111", handlr)
	log.Fatal(x)
}
