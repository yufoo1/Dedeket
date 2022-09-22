package model

import (
	"E-TexSub-backend/global"
	"fmt"
)

type Textbook struct {
	Name        string `json:"name"`
	Writer      string `json:"writer"`
	Class       string `json:"class"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	CreatedBy   string `json:"createdBy"`
}

func InsertTextbook(textbook *Textbook) {
	_, err := global.MysqlDb.Exec("insert into textbook(name, writer, class, description, createdAt, createBy)values(?, ?, ?, ?, ?, ?)",
		textbook.Name, textbook.Writer, textbook.Class, textbook.Description, textbook.CreatedAt, textbook.CreatedBy)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
}
