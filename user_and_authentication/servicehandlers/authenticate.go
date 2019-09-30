package servicehandlers

import (
	"Messanger/user_and_authentication/structures"
	"encoding/json"

	"log"
	"net/http"

	"github.com/google/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AuthenticateHandler struct {
}

func (p AuthenticateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p AuthenticateHandler) Get(r *http.Request) (string, int) {
	log.Println("in authenticate get")
	var contactObject structures.Contact2
	var sessionObject structures.Session
	var userObject structures.UserDetails
	session, er := mgo.Dial("127.0.0.1")
	if er != nil {
		panic(er)
	}
	defer session.Close()
	keys, ok := r.URL.Query()["token"]
	log.Println(keys[0])
	log.Println(ok)
	sessionConn := session.DB("messanger").C("sessions")
	sessionConn.Find(bson.M{"token": keys[0]}).One(&sessionObject)

	userConn := session.DB("messanger").C("users")
	userConn.Find(bson.M{"_id": sessionObject.UserId}).One(&userObject)

	contactsConn := session.DB("messanger").C("contacts")
	contactsConn.Find(bson.M{"email": userObject.Email}).One(&contactObject)

	var usercontact structures.UserContact = structures.UserContact{UserObject: userObject, ContactObject: contactObject}
	res, e := json.Marshal(usercontact)
	if e != nil {
		log.Fatal(e)
	}
	return string(res), 200
}

func (p AuthenticateHandler) Put(r *http.Request) (string, int) {
	log.Println("in authenticate put")
	return "AUTHENTICATE PUT Called", 200
}

func (p AuthenticateHandler) Post(r *http.Request) (string, int) {
	log.Println("in authenticate post")
	var u structures.User
	var result structures.UserDetails
	session, er := mgo.Dial("127.0.0.1")
	if er != nil {
		panic(er)
	}
	defer session.Close()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	log.Printf("%s \n", u.Email)
	log.Printf("%s \n", u.Password)
	c := session.DB("messanger").C("users")
	c.Find(bson.M{"email": u.Email}).One(&result)

	if result.Password != u.Password {
		return "unauthorised", 401
	}
	u1 := uuid.Must(uuid.NewRandom())
	c = session.DB("messanger").C("sessions")
	token := u1.String()
	//fmt.Printf(token)
	var ses structures.Session = structures.Session{ID: token, Token: token, UserId: result.Id}
	err = c.Insert(&ses)
	if err != nil {
		log.Fatal(err)
	}

	res, e := json.Marshal(ses)
	if e != nil {
		log.Fatal(e)
	}
	return string(res), 200
}
