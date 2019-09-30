package main

import (
	"Messanger/chat/servicehandlers"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	flag.Parse()
	p := servicehandlers.PingHandler{}
	c := servicehandlers.ChatHandler{}
	router := mux.NewRouter()

	router.HandleFunc("/upload", servicehandlers.UploadFile)
	router.Handle("/ping", p)
	router.Handle("/chat", c)
	log.Fatal(http.ListenAndServe(":2428",
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(router)))

}
