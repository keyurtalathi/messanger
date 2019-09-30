package core

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Contact struct {
	Id          bson.ObjectId       `json:"id" bson:"_id"`
	Email       string              `bson:"email"`
	ContactList []map[string]string `bson:"contact_list"`
	BlockList   []string            `bson:"block_list"`
	UserId      bson.ObjectId       `bson:"user_id"`
	Name        string              `bson:"name"`
}

func CreateContactBlockGroupLists(c *Client) *Client {
	var contactObject Contact
	var groupresult []struct {
		GroupIds bson.ObjectId `bson:"group_id"`
	}

	session, er := mgo.Dial("127.0.0.1")
	var contactList []string
	var blockList []string
	if er != nil {
		panic(er)
	}
	defer session.Close()
	contactsConn := session.DB("messanger").C("contacts")
	contactsConn.Find(bson.M{"email": c.Email}).One(&contactObject)
	for i := range contactObject.ContactList {
		for k := range contactObject.ContactList[i] {
			contactList = append(contactList, k)
		}
	}
	chatConn := session.DB("messanger").C("chat")
	_ = chatConn.Find(bson.M{"email": c.Email}).Select(bson.M{"group_id": 1}).All(&groupresult)
	blockList = contactObject.BlockList
	c.Clients = contactList
	c.Block = blockList
	for _, v := range groupresult {
		c.Groups = append(c.Groups, v.GroupIds.Hex())
	}
	return c
}
