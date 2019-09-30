package servicehandlers

import (
	"Messanger/user_and_authentication/structures"
	"encoding/json"

	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type GroupHandler struct {
}

func (p GroupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p GroupHandler) Get(r *http.Request) (string, int) {
	log.Println("in GET")
	var groupObject structures.Group2
	session, er := mgo.Dial("127.0.0.1")
	if er != nil {
		panic(er)
	}
	defer session.Close()
	keys, ok := r.URL.Query()["groupId"]
	var groupId = bson.ObjectIdHex(keys[0])
	log.Println(keys[0])
	log.Println(ok)
	groupConn := session.DB("messanger").C("groups")
	groupConn.Find(bson.M{"_id": groupId}).One(&groupObject)
	res, e := json.Marshal(groupObject)
	if e != nil {
		log.Fatal(e)
	}
	return string(res), 200
}

func (p GroupHandler) Put(r *http.Request) (string, int) {
	log.Println("in PUT")
	var input structures.GroupPUT
	var sessionObject structures.Session
	var userObject1, userObject2 structures.UserDetails
	var groupObject structures.Group2

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
	//fmt.Println("-----" + sessionObject.Token)
	userConn := session.DB("messanger").C("users")
	userConn.Find(bson.M{"_id": sessionObject.UserId}).One(&userObject1)
	//fmt.Println("---usr--" + userObject1.Username)
	groupConn := session.DB("messanger").C("groups")
	groupConn.Find(bson.M{"_id": input.Group.Id}).One(&groupObject)
	//fmt.Println("-----" + groupObject.GroupName)
	isAdmin := 0
	isMember := 0

	if input.Type == "addadmin" {
		adminList := groupObject.AdminMembers
		//fmt.Println("-----" + adminList[0])
		for i := 0; i < len(adminList); i++ {
			if adminList[i] == userObject1.Email {
				isAdmin = 1
				break
			}
		}
		log.Println(adminList)
		membersList := groupObject.GroupMembers
		for i := 0; i < len(membersList); i++ {
			if membersList[i] == input.AddAdmin {
				isMember = 1
				break
			}
		}
		if isAdmin == 1 && isMember == 1 {
			adminList = append(adminList, input.AddAdmin)
			err = groupConn.Update(bson.M{"_id": groupObject.Id},
				bson.M{"$set": bson.M{"admin_members": adminList}})

			if err != nil {
				panic(err)
			}
		}
	} else if input.Type == "removemember" {
		adminList := groupObject.AdminMembers
		membersList := groupObject.GroupMembers
		for i := 0; i < len(adminList); i++ {
			if adminList[i] == userObject1.Email {
				isAdmin = 1
				break
			}
		}
		for i := 0; i < len(membersList); i++ {
			if membersList[i] == input.AddAdmin {
				isMember = 1
				break
			}
		}
		if isAdmin == 1 && isMember == 1 {
			for i := range membersList {
				if input.RemoveGroupMember == membersList[i] {
					membersList = append(membersList[:i], membersList[i+1:]...)
					break
				}
			}
			err = groupConn.Update(bson.M{"_id": groupObject.Id},
				bson.M{"$set": bson.M{"group_members": membersList}})

			if err != nil {
				panic(err)
			}
		}
	} else if input.Type == "addmembers" {
		adminList := groupObject.AdminMembers
		membersList := groupObject.GroupMembers
		inputMembersList := input.AddGroupMembers
		for i := 0; i < len(adminList); i++ {
			if adminList[i] == userObject1.Email {
				isAdmin = 1
				break
			}
		}
		for i := 0; i < len(inputMembersList); i++ {
			userConn.Find(bson.M{"email": inputMembersList[i]}).One(&userObject2)
			if userObject2.Email != inputMembersList[i] {
				inputMembersList = append(inputMembersList[:i], inputMembersList[i+1:]...)
			}
		}
		if isAdmin == 1 {
			membersList = append(membersList, inputMembersList...)
			err = groupConn.Update(bson.M{"_id": groupObject.Id},
				bson.M{"$set": bson.M{"group_members": membersList}})

			if err != nil {
				panic(err)
			}
		}
	} else if input.Type == "changeimage" {
		membersList := groupObject.GroupMembers
		for i := 0; i < len(membersList); i++ {
			if membersList[i] == userObject1.Email {
				isMember = 1
				break
			}
		}
		if isMember == 1 {
			err = groupConn.Update(bson.M{"_id": groupObject.Id},
				bson.M{"$set": bson.M{"image": input.GroupImage}})

			if err != nil {
				panic(err)
			}
		}

	} else if input.Type == "changename" {
		membersList := groupObject.GroupMembers
		for i := 0; i < len(membersList); i++ {
			if membersList[i] == userObject1.Email {
				isMember = 1
				break
			}
		}
		if isMember == 1 {
			err = groupConn.Update(bson.M{"_id": groupObject.Id},
				bson.M{"$set": bson.M{"name": input.GroupName}})

			if err != nil {
				panic(err)
			}
		}

	}
	groupConn.Find(bson.M{"_id": input.Group.Id}).One(&groupObject)
	res, e := json.Marshal(groupObject)
	if e != nil {
		log.Fatal(e)
	}
	return string(res), 200
}

func (p GroupHandler) Post(r *http.Request) (string, int) {
	log.Println("in POST")
	var input structures.GroupPOST
	var sessionObject structures.Session
	var userObject structures.UserDetails
	var groupObject structures.Group1
	var admin_members []string

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
	userConn.Find(bson.M{"_id": sessionObject.UserId}).One(&userObject)

	input.GroupMembers = append(input.GroupMembers, userObject.Email)
	groupConn := session.DB("messanger").C("groups")

	admin_members = append(admin_members, userObject.Email)
	err = groupConn.Insert(&structures.Group1{AdminMembers: admin_members,
		GroupName:    input.GroupName,
		GroupImage:   input.GroupImage,
		GroupMembers: input.GroupMembers})

	log.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	dbSize, err := groupConn.Count()
	if err != nil {
		log.Fatal(err)
	}

	err = groupConn.Find(nil).Skip(dbSize - 1).One(&groupObject)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(groupObject)
	res, e := json.Marshal(groupObject)
	if e != nil {
		log.Fatal(e)
	}
	return string(res), 200
}
