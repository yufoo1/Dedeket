package api

import (
	"E-TexSub-backend/model/chat"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	upgrade = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

var connMap = make(map[string]*chat.Connection) // 用户通过token连接websocket(使用token可以避免用户伪造他人)，后端通过token获取用户的username，username作为coonMap的key
// 单独发送消息时，用户需要同时传入接收方的username

func WsHandler(c *gin.Context) {
	fmt.Println("connecting...")
	WsServer(c.Writer, c.Request)
}

func WsServer(w http.ResponseWriter, r *http.Request) {
	var (
		//websocket 长连接
		wsConn *websocket.Conn
		err    error
		conn   *chat.Connection
		data   []byte
	)

	token := r.URL.Query().Get("token")
	//var usernameArr []string
	//err = global.MysqlDb.Select(&usernameArr, "select username from user_token where token=?", token)
	//sourceUsername := usernameArr[0]
	//targetUsername := r.URL.Query().Get("targetUsername")

	//header中添加Upgrade:websocket
	if wsConn, err = upgrade.Upgrade(w, r, nil); err != nil {
		return
	}

	if conn, err = chat.InitConnection(wsConn); err != nil {
		goto ERR
	}

	fmt.Println("connect successfully!")
	//fmt.Println(sourceUsername)
	connMap[token] = conn

	for {
		//if data, err = connMap[sourceUsername].ReadMessage(); err != nil {
		//	goto ERR
		//}
		//if err = connMap[targetUsername].WriteMessage(data); err != nil {
		//	goto ERR
		//}
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		} else {
			fmt.Println("sending message...")
			var message = new(chat.Message)
			message.Username = "yufoo1"
			message.Data = string(data)
			message.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
			fmt.Println(message)
			chat.InsertOneMessage(message)
		}
		//if err = conn.WriteMessage(data); err != nil {
		//	goto ERR
		//}
		for key := range connMap {
			if err = connMap[key].WriteMessage(data); err != nil {
				goto ERR
			}
		}
	}

ERR:
	conn.Close()
}
