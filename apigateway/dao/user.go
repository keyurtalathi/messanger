package dao

import (
	"Messanger/apigateway/structures"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

func Get_user_by_id(userId bson.ObjectId, userdetails structures.UserDetails, session *mgo.Session) structures.UserDetails {
	collection := session.DB("messanger").C("users")
	collection.Find(bson.M{"_id": userId}).One(&userdetails)
	return userdetails
}
