package core

import (
	"bytes"
	
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func Create_personal_chats(sender string, receiver string) string {
	apiUrl := "http://127.0.0.1:2428/chat"
	x := `{"email":"` + sender + `","type":"personal","contactList":["` + receiver + `"]}`
	//fmt.Println(x)
	data := []byte(x)
	body := bytes.NewReader(data)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", apiUrl, body) // URL-encoded payload

	req.Header.Add("Content-Type", "application/json")

	resp, _ := client.Do(req)
	log.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		return "{}"
	}
	resData, _ := ioutil.ReadAll(resp.Body)
	res := string(resData)
	return res
}

func Create_group_chats(sender string, receiver []string, groupId bson.ObjectId) string {
	apiUrl := "http://127.0.0.1:2428/chat"

	strgroupid, _ := groupId.MarshalJSON()
	x := `{"email":"` + sender + `","type":"group","contactList":["` + strings.Join(receiver, `","`) + `"],"groupId":` + string(strgroupid) + `}`
	//fmt.Println(x)
	data := []byte(x)
	body := bytes.NewReader(data)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", apiUrl, body) // URL-encoded payload

	req.Header.Add("Content-Type", "application/json")

	resp, _ := client.Do(req)
	log.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		return "{}"
	}
	resData, _ := ioutil.ReadAll(resp.Body)
	res := string(resData)
	return res
}
