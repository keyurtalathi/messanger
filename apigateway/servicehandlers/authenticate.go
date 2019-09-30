package servicehandlers

import (
	"Messanger/apigateway/utils"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type AuthenticateHandler struct {
}

func (p AuthenticateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p AuthenticateHandler) Get(r *http.Request) (string, int) {
	apiUrl := "http://127.0.0.1:1111/authenticate?token="

	data := r.URL.Query()["token"][0]

	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiUrl+data, nil) // URL-encoded payload

	req.Header.Add("Content-Type", "application/json")

	resp, _ := client.Do(req)
	log.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		return utils.ResponseView(
			"{}",
			"something went wrong in authentication",
			false), 401
	}
	resData, _ := ioutil.ReadAll(resp.Body)
	res := string(resData)
	log.Println(res)
	return utils.ResponseView(
		res,
		"All OK",
		true), 200
}

func (p AuthenticateHandler) Put(r *http.Request) (string, int) {
	return "AUTHENTICATE PUT Called", 200
}

func (p AuthenticateHandler) Post(r *http.Request) (string, int) {
	apiUrl := "http://127.0.0.1:1111/authenticate"

	data, _ := ioutil.ReadAll(r.Body)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", apiUrl, bytes.NewReader(data)) // URL-encoded payload

	req.Header.Add("Content-Type", "application/json")

	resp, _ := client.Do(req)
	log.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		return utils.ResponseView(
			"{}",
			"something went wrong",
			false), 401
	}
	resData, _ := ioutil.ReadAll(resp.Body)
	res := string(resData)
	log.Println(res)
	return utils.ResponseView(
		res,
		"All OK",
		true), 200
}
