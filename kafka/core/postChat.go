package core

import (
	"Messanger/kafka/dao"
	"Messanger/kafka/structures"
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Post_msg_personal(session *mgo.Session, msgSender string, msgReceiver string, msgBody structures.Message) {

	var chatD structures.ChatsDetails
	msgReceiverL := []string{msgReceiver}
	chatD = dao.Get_chat_by_sender_receiver(msgSender, msgReceiverL, chatD, session)
	mList := chatD.MessageList
	mList = append(mList, msgBody)
	dao.Update_chat(session, msgSender, msgReceiverL, mList, msgBody)
	listSender := []string{msgSender}
	dao.Update_chat(session, msgReceiverL[0], listSender, mList, msgBody)

}

func Post_msg_group(session *mgo.Session, msgSender string, groupId string, msgBody structures.Message) {
	var chatD []structures.ChatsDetails
	//msgReceiverL := []string{msgReceiver}
	grpId := bson.ObjectIdHex(groupId)
	chatD = dao.Get_all_chat_by_group(session, grpId, chatD)
	fmt.Println(chatD)

	for _, elem := range chatD {
		mList := elem.MessageList
		mList = append(mList, msgBody)
		dao.Update_chat_group(session, msgSender, mList, grpId, msgBody)

	}

	// mList := chatD.MessageList
	// mList = append(mList, msg)
	// dao.Update_chat(session, msgSender, msgReceiver, msg)
	// listSender := []string{msgSender}
	// dao.Update_chat(session, msgReceiver[0], listSender, msg)

}
