package servicehandlers

import (
	"Messanger/apigateway/utils"

	"io/ioutil"
	"log"
	"net/http"
)

type ChatHandler struct {
}

func (p ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p ChatHandler) Get(r *http.Request) (string, int) {
	//fmt.Println("In chat apigateway get")
	apiUrl := "http://127.0.0.1:2428/chat?email="

	data := r.URL.Query()["email"][0]

	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiUrl+data, nil) // URL-encoded payload

	req.Header.Add("Content-Type", "application/json")

	resp, _ := client.Do(req)
	log.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		return utils.ResponseView(
			"{}",
			"something went wrong ",
			false), 402
	}
	resData, _ := ioutil.ReadAll(resp.Body)
	res := string(resData)
	log.Println(res)
	return utils.ResponseView(
		res,
		"All OK",
		true), 200
}

func (p ChatHandler) Put(r *http.Request) (string, int) {
	return "ChatHandler PUT Called", 200
}

func (p ChatHandler) Post(r *http.Request) (string, int) {
	return "ChatHandler POST Called", 200
}
