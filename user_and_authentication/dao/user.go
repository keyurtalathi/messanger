package dao

import (
	"Messanger/user_and_authentication/structures"
	"log"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

func Get_user_by_id(userId bson.ObjectId, userdetails structures.UserDetails, session *mgo.Session) structures.UserDetails {
	collection := session.DB("messanger").C("users")
	collection.Find(bson.M{"_id": userId}).One(&userdetails)
	return userdetails
}
func Get_user_by_email(email string, u structures.UserDetails, session *mgo.Session) structures.UserDetails {
	c := session.DB("messanger").C("users")
	c.Find(bson.M{"email": email}).One(&u)
	return u
}
func Create_user(signUserDetails structures.SignUpUserDetails, session *mgo.Session) {
	c := session.DB("messanger").C("users")
	err := c.Insert(&signUserDetails)
	if err != nil {
		log.Fatal(err)
	}

}
