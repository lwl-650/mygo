package apis

import (
	"fmt"
	"log"
	"mygo/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// https://www.cnblogs.com/chnmig/p/14463837.html  websocket

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

var clientMap map[string]*websocket.Conn

var pointMap map[string]*websocket.Conn

func init() {
	clientMap = make(map[string]*websocket.Conn)
	pointMap = make(map[string]*websocket.Conn)
}

func (WebsocketController) GetPushNews(c *gin.Context) {

	var (
		ws  *websocket.Conn
		err error
	)
	ws, err = upgrader.Upgrade(c.Writer, c.Request, nil)

	for id, _ := range clientMap {
		if id == c.Query("id") {
			err = ws.WriteJSON("id已存在")
			if err != nil {
				log.Println("write:", err)
				break
			}
			ws.Close()
			return
		}
	}

	clientMap[c.Query("id")] = ws
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer ws.Close()
	fmt.Println(clientMap)
	fmt.Println(len(clientMap))
	for {

		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			delete(clientMap, c.Query("id"))
			fmt.Println(clientMap)
			break
		}
		log.Printf("recv: %s", message)
		getMap := make(map[string]interface{})
		getMap["data"] = string(message)

		// err = ws.WriteMessage(mt, message)

		// err = ws.WriteJSON(getMap)
		fmt.Println(clientMap)

		for _, value := range clientMap {
			err = value.WriteJSON(getMap)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}

	}
}

func (WebsocketController) GetWebsocket(c *gin.Context) {
	id := c.Query("id")
	uid := c.Query("uid")
	for index, value := range clientMap {
		if index == uid || index == id {
			pointMap[index] = value
		}
	}
	fmt.Println(pointMap)
	util.Success(c, pointMap)

}
