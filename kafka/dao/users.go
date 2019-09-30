package dao

import (
	"Messanger/kafka/structures"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

func Get_user_by_email(email string, session *mgo.Session) structures.Contact {
	var objects_list structures.Contact
	c := session.DB("messanger").C("contacts")
	c.Find(bson.M{"email": email}).One(&objects_list)
	return objects_list
}

func Get_users_by_email_list(email_list []string, session *mgo.Session) []structures.Contact {
	var objects_list []structures.Contact
	c := session.DB("messanger").C("contacts")
	c.Find(bson.M{"email": bson.M{"$in": email_list}}).All(&objects_list)
	return objects_list
}
func Get_groups_by_group_id(groupId bson.ObjectId, session *mgo.Session) structures.Group {
	var objects_list structures.Group
	c := session.DB("messanger").C("groups")
	c.Find(bson.M{"_id": groupId}).One(&objects_list)
	return objects_list
}
