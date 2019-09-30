package servicehandlers

import (
	"Messanger/user_and_authentication/structures"
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ContactHandler struct {
}

type ContactItem struct {
	ContactEmail string `bson:"contact_email"`
	ContactName  string `bson:"contact_name"`
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
type Contact2 struct { //get,put call
	Id           bson.ObjectId `bson:"_id"`
	Email        string        `bson:"email"`
	ContactList  []ContactItem `bson:"contact_list"`
	BlockList    []ContactItem `bson:"block_list"`
	ContactImage string        `bson:"image"`
}

func (p ContactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}
func uniqueItems(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func (p ContactHandler) Get(r *http.Request) (string, int) {
	log.Println("in GET")
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

	res, e := json.Marshal(contactObject)
	if e != nil {
		log.Fatal(e)
	}
	return string(res), 200
}

func (p ContactHandler) Put(r *http.Request) (string, int) {
	log.Println("in PUT")
	var input structures.ContactPUT
	var sessionObject structures.Session
	var contactObject structures.Contact2
	var userObject1 structures.UserDetails
	var userObject2 structures.UserDetails
	session, er := mgo.Dial("127.0.0.1")
	if er != nil {
		panic(er)
	}
	defer session.Close()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	sessionConn := session.DB("messanger").C("sessions")
	sessionConn.Find(bson.M{"token": input.Token}).One(&sessionObject)

	userConn := session.DB("messanger").C("users")
	userConn.Find(bson.M{"_id": sessionObject.UserId}).One(&userObject1)

	contactsConn := session.DB("messanger").C("contacts")
	contactsConn.Find(bson.M{"email": userObject1.Email}).One(&contactObject)

	if input.Type == "addcontact" &&
		input.ContactPayload.ContactEmail != contactObject.Email {
		contactList := contactObject.ContactList

		userConn.Find(bson.M{"email": input.ContactPayload.ContactEmail}).One(&userObject2)
		duplicateFlag := 0
		if userObject2.Email == input.ContactPayload.ContactEmail {
			for i, p := range contactList {
				if p.ContactEmail == input.ContactPayload.ContactEmail {
					duplicateFlag = 1
					log.Println(i)
				}
			}
			if duplicateFlag == 0 {
				var contactItem structures.ContactItem = structures.ContactItem{
					ContactEmail: input.ContactPayload.ContactEmail,
					ContactName:  input.ContactPayload.ContactName}
				contactList = append(contactList, contactItem)
				err = contactsConn.Update(bson.M{"_id": contactObject.Id},
					bson.M{"$set": bson.M{"contact_list": contactList}})

				if err != nil {
					panic(err)
				}
			}
		}

	} else if input.Type == "blockcontact" {
		blockList := contactObject.BlockList
		var contactItem structures.ContactItem = structures.ContactItem{
			ContactEmail: input.ContactPayload.ContactEmail,
			ContactName:  input.ContactPayload.ContactName}
		blockList = append(blockList, contactItem)
		err = contactsConn.Update(bson.M{"_id": contactObject.Id},
			bson.M{"$set": bson.M{"block_list": blockList}})

		if err != nil {
			panic(err)
		}
	} else if input.Type == "changeimage" {
		err = contactsConn.Update(bson.M{"_id": contactObject.Id},
			bson.M{"$set": bson.M{"image": input.ContactImage}})

		if err != nil {
			panic(err)
		}
	} else if input.Type == "unblock" {
		blockList := contactObject.BlockList
		for i, p := range blockList {
			if p.ContactEmail == input.ContactPayload.ContactEmail {
				blockList = append(blockList[:i], blockList[i+1:]...)
				break
			}
		}
		log.Println(blockList)
		err = contactsConn.Update(bson.M{"_id": contactObject.Id},
			bson.M{"$set": bson.M{"block_list": blockList}})

		if err != nil {
			panic(err)
		}
	} else if input.Type == "removecontact" {
		contactList := contactObject.ContactList
		blockList := contactObject.BlockList
		for i, p := range contactList {
			if p.ContactEmail == input.ContactPayload.ContactEmail {
				contactList = append(contactList[:i], contactList[i+1:]...)
				break
			}
		}
		for i, p := range blockList {
			if p.ContactEmail == input.ContactPayload.ContactEmail {
				blockList = append(contactList[:i], contactList[i+1:]...)
				break
			}
		}
		err = contactsConn.Update(bson.M{"_id": contactObject.Id},
			bson.M{"$set": bson.M{"contact_list": contactList,
				"block_list": blockList}})

		if err != nil {
			panic(err)
		}
	}

	contactsConn.Find(bson.M{"email": userObject1.Email}).One(&contactObject)

	res, e := json.Marshal(contactObject)
	log.Print(string(res))
	if e != nil {
		log.Fatal(e)
	}
	return string(res), 200
}

func (p ContactHandler) Post(r *http.Request) (string, int) {
	log.Println("in POST")
	var input ContactPOST
	var contactObject Contact1
	session, er := mgo.Dial("127.0.0.1")
	if er != nil {
		panic(er)
	}
	defer session.Close()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	contactsConn := session.DB("messanger").C("contacts")

	err = contactsConn.Insert(&Contact1{Email: input.Email,
		BlockList:    input.BlockList,
		ContactImage: input.ContactImage})

	log.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	dbSize, err := contactsConn.Count()
	if err != nil {
		log.Fatal(err)
	}

	err = contactsConn.Find(nil).Skip(dbSize - 1).One(&contactObject)
	if err != nil {
		log.Fatal(err)
	}
	res, e := json.Marshal(contactObject)
	if e != nil {
		log.Fatal(e)
	}
	return string(res), 200
}
