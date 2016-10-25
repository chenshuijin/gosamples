package main

import (
	"fmt"
	"gopkg.in/pg.v4"
)

var (
	predefinemessage_select = `select key,language,picture_height,picture_width,text,picture_identifier from predefined_message_translations`
)

type PredefineMessage struct {
	Key, Language, Text, Picture_height, Picture_width, Picture_identifier string
}

func GetPredefineMessages() []PredefineMessage {
	db := pg.Connect(pgDbConf)
	defer db.Close()
	var messages []PredefineMessage
	_, err := db.Query(&messages, predefinemessage_select)
	if err != nil {
		fmt.Println("db query err:", err)
	}
	//	fmt.Println("messages:", messages)
	return messages
}
