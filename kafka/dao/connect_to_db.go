package dao

import "gopkg.in/mgo.v2"

func Connection_to_mongo() *mgo.Session {
	session, er := mgo.Dial("127.0.0.1")
	if er != nil {
		panic(er)
	}
	return session
}
