package servicehandlers

import (
	"Messanger/apigateway/utils"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type UploadHandler struct {
}

func (p UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p UploadHandler) Get(r *http.Request) (string, int) {
	return "Upload GET Called", 200
}
func (p UploadHandler) Put(r *http.Request) (string, int) {
	return "Upload PutT Called", 200
}

func (p UploadHandler) Post(r *http.Request) (string, int) {
	apiUrl := "http://127.0.0.1:2428/upload"

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
