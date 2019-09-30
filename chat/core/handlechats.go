package core

import (
	"Messanger/chat/dao"
	"Messanger/chat/structures"
	
	"strings"

	"github.com/Shopify/sarama"
)

func Handle_messages(msg *sarama.ConsumerMessage) {
	message := string(msg.Value)
	//fmt.Println("***************" + message + "***************")
	msg_list := strings.Split(message, ":")
	msg_personal_group := msg_list[0]
	msg_sender := msg_list[1]
	msg_receiver := msg_list[2]
	msg_type := msg_list[3]
	msg_body := msg_list[4]

	var msg_obj structures.Message
	msg_obj.MessageBody = msg_body
	msg_obj.MessageType = msg_type
	msg_obj.Reciever = []string{msg_receiver}
	msg_obj.Sender = msg_sender
	msg_obj.Status = ""
	session := dao.Connection_to_mongo()
	defer session.Close()

	if msg_personal_group == "personal" {

		Post_msg_personal(session, msg_sender, msg_receiver, msg_obj)

	} else if msg_personal_group == "group" {

		Post_msg_group(session, msg_sender, msg_receiver, msg_obj)
	}
}
