package structures

import (
	"gopkg.in/mgo.v2/bson"
)

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
	Email    string
	Password string
}

type ContactPUT struct {
	Token          string
	Type           string
	ContactPayload ContactItem
	ContactImage   string
}

type ContactPOST struct {
	Email        string
	ContactImage string
	ContactList  []ContactItem
	BlockList    []ContactItem
}

type Contact1 struct { //post call
	Id           bson.ObjectId `bson:"_id,omitempty"`
	Email        string        `bson:"email"`
	ContactList  []ContactItem `bson:"contact_list"`
	BlockList    []ContactItem `bson:"block_list"`
	ContactImage string        `bson:"image"`
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

type GroupPUT struct {
	Token             string
	Type              string
	Group             Group2
	RemoveGroupMember string
	AddGroupMembers   []string
	GroupName         string
	GroupImage        string
	AddAdmin          string
}
type Group1 struct { //post call
	Id           bson.ObjectId `bson:"_id,omitempty"`
	AdminMembers []string      `bson:"admin_members"`
	GroupMembers []string      `bson:"group_members"`
	GroupName    string        `bson:"name"`
	GroupImage   string        `bson:"image"`
}
type Group2 struct { //put,get call
	Id           bson.ObjectId `bson:"_id"`
	AdminMembers []string      `bson:"admin_members"`
	GroupMembers []string      `bson:"group_members"`
	GroupName    string        `bson:"name"`
	GroupImage   string        `bson:"image"`
}
type SignUpUserDetails struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
	Alias    string `bson:"alias"`
	Phone    int    `bson:"phone"`
	Email    string `bson:"email"`
}

type UserContact struct {
	UserObject    UserDetails
	ContactObject Contact2
}
