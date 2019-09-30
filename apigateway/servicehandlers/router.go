package servicehandlers

import (
	"fmt"
	"net/http"
)

type HttpServiceHandler interface {
	Get(*http.Request) (string, int)
	Put(*http.Request) (string, int)
	Post(*http.Request) (string, int)
}

func methodRouter(p HttpServiceHandler, w http.ResponseWriter, r *http.Request) {

	var response string
	var code int
	if r.Method == "GET" {
		response, code = p.Get(r)
	} else if r.Method == "PUT" {
		response, code = p.Put(r)
	} else if r.Method == "POST" {
		response, code = p.Post(r)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	fmt.Fprintf(w, response)
}
