package deal

import (
	"Dedeket/global"
	"fmt"
)

type Textbook struct {
	Id          int    `json:"id" db:"id"`
	BookName    string `json:"bookName" db:"bookName"`
	Writer      string `json:"writer" db:"writer"`
	Class       string `json:"class" db:"class"`
	Description string `json:"description" db:"description"`
	CreatedAt   string `json:"createdAt" db:"createdAt"`
	Seller      string `json:"seller" db:"seller"`
	College     string `json:"college" db:"college"`
	Remain      int64  `json:"remain" db:"remain"`
	Total       int64  `json:"total" db:"total"`
}

func InsertTextbook(textbook *Textbook) {
	_, err := global.MysqlDb.Exec("insert into textbook(bookName, writer, class, description, createdAt, seller, total, remain, college)values(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		textbook.BookName, textbook.Writer, textbook.Class, textbook.Description, textbook.CreatedAt, textbook.Seller, textbook.Total, textbook.Total, textbook.College)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
}
