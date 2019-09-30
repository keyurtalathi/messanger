package servicehandlers

import (
	"Messanger/apigateway/core"
	"Messanger/apigateway/structures"
	"Messanger/apigateway/utils"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ContactHandler struct {
}

func (p ContactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p ContactHandler) Get(r *http.Request) (string, int) {
	apiUrl := "http://127.0.0.1:1111/contactlist?token="

	data := r.URL.Query()["token"][0]

	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiUrl+data, nil) // URL-encoded payload

	req.Header.Add("Content-Type", "application/json")

	resp, _ := client.Do(req)
	log.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		return utils.ResponseView(
			"{}",
			"something went wrong ",
			false), 400
	}
	resData, _ := ioutil.ReadAll(resp.Body)
	res := string(resData)
	log.Println(res)
	return utils.ResponseView(
		res,
		"All OK",
		true), 200
}

func (p ContactHandler) Put(r *http.Request) (string, int) {
	apiUrl := "http://127.0.0.1:1111/contactlist"

	// data := r.Body
	data, err := ioutil.ReadAll(r.Body)

	client := &http.Client{}
	req, _ := http.NewRequest("PUT", apiUrl, bytes.NewReader(data)) // URL-encoded payload

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

	var contact structures.Contact2
	decode := json.NewDecoder(bytes.NewReader(resData))
	err = decode.Decode(&contact)

	if err != nil {
		panic(err)
	}
	var input structures.ContactPUT
	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&input)

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	if input.Type == "addcontact" {
		_ = core.Create_personal_chats(input.ContactPayload.ContactEmail, contact.Email)
		_ = core.Create_personal_chats(contact.Email, input.ContactPayload.ContactEmail)
	}

	log.Println(res)
	return utils.ResponseView(
		res,
		"All OK",
		true), 200
}

func (p ContactHandler) Post(r *http.Request) (string, int) {
	apiUrl := "http://127.0.0.1:1111/contactlist"

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
			false), 400
	}
	resData, _ := ioutil.ReadAll(resp.Body)

	res := string(resData)

	log.Println(res)
	return utils.ResponseView(
		res,
		"All OK",
		true), 200
}
