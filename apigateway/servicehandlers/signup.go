package servicehandlers

import (
	"Messanger/apigateway/utils"
	"bytes"

	"io/ioutil"
	"log"
	"net/http"
)

type SignupHandler struct {
}

func (p SignupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p SignupHandler) Get(r *http.Request) (string, int) {
	return "signup GET Called", 200
}

func (p SignupHandler) Put(r *http.Request) (string, int) {
	return "signup PUT Called", 200
}

func (p SignupHandler) Post(r *http.Request) (string, int) {
	apiUrl := "http://127.0.0.1:1111/signup"

	data, _ := ioutil.ReadAll(r.Body)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", apiUrl, bytes.NewReader(data)) // URL-encoded payload

	req.Header.Add("Content-Type", "application/json")

	resp, _ := client.Do(req)
	log.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		return utils.ResponseView(
			"{}",
			"something went wrong in authentication",
			false), 400
	}
	resData, _ := ioutil.ReadAll(resp.Body)
	res := string(resData)
	//fmt.Println(res)
	if res == `{"msg":"User Already Present"}` {
		return utils.ResponseView(
			"{}",
			"already exists",
			false), 400
	}
	log.Println(res)

	apiUrlContact := "http://127.0.0.1:1111/contactlist"

	// dataContact := r.Body
	dataContact := data
	log.Println("************" + string(dataContact))
	clientContact := &http.Client{}
	reqConatact, _ := http.NewRequest("POST", apiUrlContact, bytes.NewReader(dataContact)) // URL-encoded payload

	req.Header.Add("Content-Type", "application/json")

	respContact, _ := clientContact.Do(reqConatact)
	log.Println(respContact.StatusCode)
	if respContact.StatusCode != 200 {
		return utils.ResponseView(
			"{}",
			"something went wrong in authentication",
			false), 400
	}
	resDataContact, _ := ioutil.ReadAll(respContact.Body)
	resContact := string(resDataContact)

	log.Println(resContact)
	return utils.ResponseView(
		`{"user":`+res+`,
		"conatact":`+resContact+`}`,
		"All OK",
		true), 200
}
