package core

import (
	"bytes"

	"log"
	"strings"

	"Messanger/apigateway/kafka"

	"github.com/gorilla/websocket"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Client struct {
	Clients []string
	Block   []string
	Groups  []string
	Conn    *websocket.Conn
	Receive chan []byte
	Email   string
	Hub     *Hub
}
type Groups struct {
	GroupId    bson.ObjectId `bson:"_id"`
	Emails     []string      `bson:"group_members"`
	Name       string        `bson:"group_name"`
	unregister string        `bson:"group_image"`
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)
var EmailClientMap = make(map[string]*Client)
var GroupClientsMap = make(map[string][]string)

func NewClient() *Client {
	return &Client{
		Clients: make([]string, 0),
		Block:   make([]string, 0),
		Groups:  make([]string, 0),
		Receive: make(chan []byte),
	}
}

func MakeGroupClientMap() map[string][]string {
	log.Println("in MakeGroupClientMap")
	var group []Groups
	session, er := mgo.Dial("127.0.0.1")
	if er != nil {
		//fmt.Println("mongo")
		panic(er)
	}
	defer session.Close()
	groupConn := session.DB("messanger").C("groups")
	_ = groupConn.Find(nil).All(&group)
	for _, v := range group {
		GroupClientsMap[v.GroupId.Hex()] = v.Emails
	}
	return GroupClientsMap
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.Hub.Send <- message
	}
	//fmt.Println("in read pump")
}

func (c *Client) WritePump() {
	for {
		select {
		case message, ok := <-c.Receive:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			n := len(c.Receive)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Receive)
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

type Hub struct {
	Register   chan *Client
	Unregister chan *Client
	Send       chan []byte
}

func NewHub() *Hub {
	return &Hub{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Send:       make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			//fmt.Println("**********in register")
			EmailClientMap[client.Email] = client

		case message := <-h.Send:
			result := strings.Split(string(message), ":")
			typeOfMessage := result[0]
			fromWhom := result[1]
			toWhom := result[2]
			_ = result[3]
			_ = result[4]
			_ = strings.Join(result[5:], ":")
			if typeOfMessage == "group" {
				listOfClients := GroupClientsMap[toWhom]
				for _, email := range listOfClients {
					reciever := EmailClientMap[email]
					// //fmt.Print(fromWhom + "**********")
					// //fmt.Println(email)
					if reciever != nil && fromWhom != email {
						reciever.Receive <- []byte(message)
					}
				}
				kafka.Push_messages(string(message))
			} else {
				reciever := EmailClientMap[toWhom]

				if reciever != nil {
					reciever.Receive <- []byte(message)
				}
				kafka.Push_messages(string(message))

			}
		}
	}
}
