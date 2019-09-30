package utils

import (
	"Messanger/chat/structures"
)

func Get_emails_list(obj []structures.ChatsDetails) []string {
	var emailList []string

	for _, ele := range obj {
		if ele.GroupId == "" {
			emailList = append(emailList, ele.ContactList[0])
		} else {
			emailList = append(emailList, ele.ContactList...)
		}
	}
	return emailList
}
func UniqueItems(stringSlice []string) []string {
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
