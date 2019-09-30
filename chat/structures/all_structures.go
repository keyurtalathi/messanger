package structures

import "gopkg.in/mgo.v2/bson"

type Message struct {
	Sender      string
	Reciever    []string
	MessageBody string
	MessageType string
	Status      string
}

type ChatsDetails struct {
	Id            bson.ObjectId `json:"id" bson:"_id"`
	Email         string        `bson:"email"`
	ContactList   []string      `bson:"contactList"`
	MessageList   []Message     `bson:"messageList"`
	RecentMessage Message       `bson:"recentMessage"`
	GroupId       bson.ObjectId `bson:"groupId"`
}
type ChatInsert struct {
	Id            bson.ObjectId `bson:"_id,omitempty"`
	Email         string        `bson:"email"`
	ContactList   []string      `bson:"contactList"`
	MessageList   []string      `bson:"messageList"`
	RecentMessage string        `bson:"recentMessage"`
	GroupId       bson.ObjectId `bson:"groupId,omitempty"`
}

type ChatPayload struct {
	Type        string        `json:"type"`
	Email       string        `json:"email"`
	ContactList []string      `json:"contactList"`
	GroupId     bson.ObjectId `json:"groupId,omitempty"`
}
type ContactItem struct {
	ContactEmail string `bson:"contact_email"`
	ContactName  string `bson:"contact_name"`
}
type Contact struct { //get,put call
	Id           bson.ObjectId `bson:"_id"`
	Email        string        `bson:"email"`
	ContactList  []ContactItem `bson:"contact_list"`
	BlockList    []ContactItem `bson:"block_list"`
	ContactImage string        `bson:"image"`
}
type Group struct { //put,get call
	Id           bson.ObjectId `bson:"_id"`
	AdminMembers []string      `bson:"admin_members"`
	GroupMembers []string      `bson:"group_members"`
	GroupName    string        `bson:"name"`
	GroupImage   string        `bson:"image"`
}
type ContactListDetails struct {
	ContactEmail string `bson:"contact_email"`
	ContactImage string `bson:"contact_name"`
}
type ChatsDetailsReformed struct {
	Id            bson.ObjectId        `json:"id" bson:"_id"`
	Email         string               `bson:"email"`
	ContactList   []ContactListDetails `bson:"contactList"`
	MessageList   []Message            `bson:"messageList"`
	RecentMessage Message              `bson:"recentMessage"`
	GroupId       bson.ObjectId        `bson:"groupId"`
	GroupImage    string
	GroupName     string
}
