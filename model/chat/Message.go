package chat

import (
	"E-TexSub-backend/global"
	"context"
	"fmt"
)

type Message struct {
	Username  string `bson:"username"`
	Data      string `bson:"data"`
	CreatedAt string `bson:"createdAt"`
}

func (Message) CollectionName() string {
	return "message_user"
}

func InsertOneMessage(message *Message) {
	_, err := global.MongoDb.Collection(Message{}.CollectionName()).InsertOne(context.Background(), message)
	if err != nil {
		fmt.Println(err)
	}
}
