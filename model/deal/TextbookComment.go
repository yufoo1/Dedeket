package deal

import (
	"Dedeket/global"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type TextbookComment struct {
	TextbookId string `bson:"textbookId"`
	Sender     string `bson:"sender"`
	Comment    string `bson:"comment"`
	CreatedAt  string `bson:"createdAt"`
}

func (TextbookComment) CollectionName() string {
	return "user_textbook_comment"
}

func InsertOneTextbookComment(textbookComment *TextbookComment) {
	_, err := global.MongoDb.Collection(TextbookComment{}.CollectionName()).InsertOne(context.Background(), textbookComment)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteTextbookAllComment(textbookId string) {
	_, err := global.MongoDb.Collection(TextbookComment{}.CollectionName()).DeleteMany(context.Background(), bson.M{"textbookId": textbookId})
	if err != nil {
		fmt.Println(err)
	}
}
