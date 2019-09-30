package servicehandlers

import (
	"log"
	"net/http"

	"Messanger/apigateway/core"
	"Messanger/apigateway/dao"
	"Messanger/apigateway/structures"
	"Messanger/apigateway/utils"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}}

// use default options

func Echo(hub *core.Hub, w http.ResponseWriter, r *http.Request) string {
	var result structures.Session
	var userdetails structures.UserDetails
	session := dao.Connection_to_mongo()
	defer session.Close()
	_ = core.MakeGroupClientMap()
	x_auth_token := r.URL.Query().Get("token")
	//fmt.Println(groups)
	//fmt.Println("x-auth-tokennnn =  " + x_auth_token)

	result = dao.Get_user_session(x_auth_token, result, session)
	if result.Token == "" {
		return utils.ResponseView(
			"{}",
			"something went wrong in authentication",
			false)
	}
	userdetails = dao.Get_user_by_id(result.UserId, userdetails, session)

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return utils.ResponseView(
			"{}",
			"error updating socket",
			false)
	}
	//	defer c.Close()
	client := &core.Client{Hub: hub, Email: userdetails.Email, Conn: c, Receive: make(chan []byte, 1000)}
	client = core.CreateContactBlockGroupLists(client)
	hub.Register <- client
	//fmt.Println(client.Email)
	go client.WritePump()
	go client.ReadPump()
	return utils.ResponseView(
		"{}",
		"Successful",
		true)
}
