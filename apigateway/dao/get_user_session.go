package dao

import (
	"Messanger/apigateway/structures"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

func Get_user_session(token string, result structures.Session, session *mgo.Session) structures.Session {
	collection := session.DB("messanger").C("sessions")
	collection.Find(bson.M{"token": token}).One(&result)
	return result
}
