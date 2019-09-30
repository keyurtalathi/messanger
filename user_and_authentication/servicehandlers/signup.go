package servicehandlers

import (
	"Messanger/user_and_authentication/dao"
	"Messanger/user_and_authentication/structures"
	"encoding/json"
	"log"
	"net/http"
)

type SignupHandler struct {
}

func (p SignupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p SignupHandler) Get(r *http.Request) (string, int) {
	return "SignUp GET Called", 200
}

func (p SignupHandler) Put(r *http.Request) (string, int) {
	return "SignUp PUT Called", 200
}

func (p SignupHandler) Post(r *http.Request) (string, int) {
	var u structures.UserDetails
	var signUserDetails structures.SignUpUserDetails
	session := dao.Connection_to_mongo()
	defer session.Close()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&signUserDetails)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	u = dao.Get_user_by_email(signUserDetails.Email, u, session)
	log.Println(u)
	if u.Email != "" {
		return `{"msg":"User Already Present"}`, 400
	}
	dao.Create_user(signUserDetails, session)
	res, e := json.Marshal(signUserDetails)
	log.Print(string(res))

	if e != nil {
		log.Fatal(e)
	}
	return string(res), 200
}
