package deal

import (
	"E-TexSub-backend/global"
	"context"
	"fmt"
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
