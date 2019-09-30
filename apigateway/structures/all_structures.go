package structures

import "gopkg.in/mgo.v2/bson"

type UserDetails struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
	Alias    string        `bson:"alias"`
	Phone    int           `bson:"phone"`
	Email    string        `bson:"email"`
}

type Session struct {
	ID     string        `bson:"_id"`
	UserId bson.ObjectId `bson:"user"`
	Token  string        `bson:"token"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type ContactPUT struct {
	Token          string
	Type           string
	ContactPayload ContactItem
	ContactImage   string
}
type ContactItem struct {
	ContactEmail string `bson:"contact_email"`
	ContactName  string `bson:"contact_name"`
}
type Contact2 struct { //get,put call
	Id           bson.ObjectId `bson:"_id"`
	Email        string        `bson:"email"`
	ContactList  []ContactItem `bson:"contact_list"`
	BlockList    []ContactItem `bson:"block_list"`
	ContactImage string        `bson:"image"`
}

type GroupPOST struct {
	Token        string
	GroupMembers []string
	GroupName    string
	GroupImage   string
}
type Group1 struct { //post call
	Id           bson.ObjectId `bson:"_id,omitempty"`
	AdminMembers []string      `bson:"admin_members"`
	GroupMembers []string      `bson:"group_members"`
	GroupName    string        `bson:"name"`
	GroupImage   string        `bson:"image"`
}
