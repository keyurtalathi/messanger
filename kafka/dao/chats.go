package dao

import (
	"Messanger/kafka/structures"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Get_all_chats_by_email(
	email string,
	result []structures.ChatsDetails,
	session *mgo.Session) []structures.ChatsDetails {
	c := session.DB("messanger").C("chats")
	c.Find(bson.M{"email": email}).All(&result)
	return result
}

func Get_chat_by_sender_receiver(sender_email string,
	receiver_email []string,
	result structures.ChatsDetails,
	session *mgo.Session) structures.ChatsDetails {
	c := session.DB("messanger").C("chats")
	c.Find(bson.M{"email": sender_email, "contactList": receiver_email}).One(&result)
	return result
}

func Create_chat(input structures.ChatInsert, session *mgo.Session) {
	c := session.DB("messanger").C("chats")
	err := c.Insert(&input)
	if err != nil {
		log.Fatal(err)
	}
}

func Update_chat(session *mgo.Session, msgSender string, msgReceiver []string, msg []structures.Message, msgBody structures.Message) {
	c := session.DB("messanger").C("chats")
	find := bson.M{"email": msgSender, "contactList": msgReceiver}
	change := bson.M{"$set": bson.M{"messageList": msg, "recentMessage": msgBody}}
	c.Update(find, change)

}

func Get_all_chat_by_group(session *mgo.Session, groupId bson.ObjectId, result []structures.ChatsDetails) []structures.ChatsDetails {

	c := session.DB("messanger").C("chats")
	c.Find(bson.M{"groupId": groupId}).All(&result)
	return result
}

func Update_chat_group(session *mgo.Session, msgSender string, msg []structures.Message, groupId bson.ObjectId, msgBody structures.Message) {
	c := session.DB("messanger").C("chats")
	find := bson.M{"groupId": groupId}
	change := bson.M{"$set": bson.M{"messageList": msg, "recentMessage": msgBody}}
	c.UpdateAll(find, change)

}
