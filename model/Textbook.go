package model

import (
	"E-TexSub-backend/global"
	"fmt"
)

type Textbook struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Writer      string `json:"writer" db:"writer"`
	Class       string `json:"class" db:"class"`
	Description string `json:"description" db:"description"`
	CreatedAt   string `json:"createdAt" db:"createdAt"`
	CreatedBy   string `json:"createdBy" db:"createdBy"`
}

func InsertTextbook(textbook *Textbook) {
	_, err := global.MysqlDb.Exec("insert into textbook(name, writer, class, description, createdAt, createBy)values(?, ?, ?, ?, ?, ?)",
		textbook.Name, textbook.Writer, textbook.Class, textbook.Description, textbook.CreatedAt, textbook.CreatedBy)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
}
