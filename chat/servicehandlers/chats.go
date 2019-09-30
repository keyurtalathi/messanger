package servicehandlers

import (
	"Messanger/chat/dao"
	"Messanger/chat/structures"
	"Messanger/chat/utils"
	"encoding/json"

	"log"
	"net/http"
)

type ChatHandler struct {
}

func (p ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p ChatHandler) Get(r *http.Request) (string, int) {
	param := r.URL.Query()["email"][0]

	//var chatDtails []ChatsDetails
	session := dao.Connection_to_mongo()
	defer session.Close()

	result := []structures.ChatsDetails{}
	result = dao.Get_all_chats_by_email(param, result, session)
	// currentUser := dao.Get_user_by_email(param, session)
	emailList := utils.Get_emails_list(result)
	emailList = utils.UniqueItems(emailList)
	//fmt.Println(emailList)

	usersList := dao.Get_users_by_email_list(emailList, session)

	var responseObj []structures.ChatsDetailsReformed
	for _, ele := range result {
		groupid, _ := ele.GroupId.MarshalJSON()
		//fmt.Println(string(groupid))
		if string(groupid) == `""` {
			for _, user := range usersList {
				//fmt.Println(ele.ContactList[0])
				//fmt.Println(user.Email)
				if ele.ContactList[0] == user.Email {
					//fmt.Println("****************")
					x := structures.ChatsDetailsReformed{
						Id:    ele.Id,
						Email: ele.Email,
						ContactList: []structures.ContactListDetails{
							{
								ContactEmail: ele.ContactList[0],
								ContactImage: user.ContactImage,
							},
						},
						MessageList:   ele.MessageList,
						RecentMessage: ele.RecentMessage,
						GroupId:       ele.GroupId,
					}
					responseObj = append(responseObj, x)
				}
			}
		} else {
			groupItem := dao.Get_groups_by_group_id(ele.GroupId, session)
			var contact_list []structures.ContactListDetails
			for _, email := range ele.ContactList {
				for _, user := range usersList {
					if email == user.Email {
						contact_list = append(contact_list, structures.ContactListDetails{
							ContactEmail: email,
							ContactImage: user.ContactImage,
						})
					}
				}
			}
			x := structures.ChatsDetailsReformed{
				Id:            ele.Id,
				Email:         ele.Email,
				ContactList:   contact_list,
				MessageList:   ele.MessageList,
				RecentMessage: ele.RecentMessage,
				GroupId:       ele.GroupId,
				GroupImage:    groupItem.GroupImage,
				GroupName:     groupItem.GroupName,
			}
			responseObj = append(responseObj, x)
		}
	}

	res, error := json.Marshal(responseObj)
	if error != nil {
		log.Fatal(error)
	}
	return string(res), 200
}

func (p ChatHandler) Put(r *http.Request) (string, int) {
	return "ChatHandler PUT Called", 200
}

func (p ChatHandler) Post(r *http.Request) (string, int) {
	var chatpayload structures.ChatPayload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&chatpayload)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	session := dao.Connection_to_mongo()
	_, _ = chatpayload.GroupId.MarshalText()
	if chatpayload.GroupId.Hex() != "" {
		insert := structures.ChatInsert{
			Email:         chatpayload.Email,
			ContactList:   chatpayload.ContactList,
			MessageList:   make([]string, 0),
			RecentMessage: "",
			GroupId:       chatpayload.GroupId,
		}
		dao.Create_chat(insert, session)

		res, error := json.Marshal(insert)
		if error != nil {
			log.Fatal(error)
		}
		return string(res), 200
	}
	defer session.Close()
	result := structures.ChatsDetails{}
	result = dao.Get_chat_by_sender_receiver(
		chatpayload.Email,
		chatpayload.ContactList,
		result,
		session,
	)
	if result.Email != "" {
		res, error := json.Marshal(result)
		if error != nil {
			log.Fatal(error)
		}
		return string(res), 200
	}
	insert := structures.ChatInsert{
		Email:         chatpayload.Email,
		ContactList:   chatpayload.ContactList,
		MessageList:   make([]string, 0),
		RecentMessage: "",
		GroupId:       chatpayload.GroupId,
	}
	dao.Create_chat(insert, session)
	result = dao.Get_chat_by_sender_receiver(
		chatpayload.Email,
		chatpayload.ContactList,
		result,
		session,
	)
	log.Println(result)
	res, error := json.Marshal(result)
	if error != nil {
		log.Fatal(error)
	}
	return string(res), 200
}
