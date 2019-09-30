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

type GroupHandler struct {
}

func (p GroupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p GroupHandler) Get(r *http.Request) (string, int) {
	return "GROUP GET Called", 200
}

func (p GroupHandler) Put(r *http.Request) (string, int) {
	apiUrl := "http://127.0.0.1:1111/group"

	data, _ := ioutil.ReadAll(r.Body)
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
	log.Println(res)
	return utils.ResponseView(
		res,
		"All OK",
		true), 200
}

func (p GroupHandler) Post(r *http.Request) (string, int) {
	apiUrl := "http://127.0.0.1:1111/group"

	data, _ := ioutil.ReadAll(r.Body)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", apiUrl, bytes.NewReader(data)) // URL-encoded payload)

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

	var group structures.Group1
	decoder := json.NewDecoder(bytes.NewReader(resData))
	err := decoder.Decode(&group)
	if err != nil {
		panic(err)
	}
	for _, email := range group.GroupMembers {
		_ = core.Create_group_chats(email, group.GroupMembers, group.Id)
	}

	res := string(resData)
	log.Println(res)
	return utils.ResponseView(
		res,
		"All OK",
		true), 200
}
