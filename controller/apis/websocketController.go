package apis

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// https://www.cnblogs.com/chnmig/p/14463837.html  websocket

var list []*websocket.Conn

type WebsocketController struct {
}

var upgrader = websocket.Upgrader{
	// 读取存储空间大小
	ReadBufferSize: 1024,
	// 写入存储空间大小
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func (WebsocketController) GetPushNews(c *gin.Context) {
	var (
		ws  *websocket.Conn
		err error
	)
	ws, err = upgrader.Upgrade(c.Writer, c.Request, nil)
	list = append(list, ws)
	if err != nil {

		log.Print("upgrade:", err)
		return
	}

	defer ws.Close()
	for {

		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		getMap := make(map[string]interface{})
		getMap["data"] = string(message)

		// err = ws.WriteMessage(mt, message)

		// err = ws.WriteJSON(getMap)
		for _, value := range list {
			err = value.WriteJSON(getMap)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}

	}
}
