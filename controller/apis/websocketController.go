package apis

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// https://www.cnblogs.com/chnmig/p/14463837.html   websocket
type WebsocketController struct {
}

var upgrader = websocket.Upgrader{
	// 这个是校验请求来源
	// 在这里我们不做校验，直接return true
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// api:/getPushNews接口处理函数
func (WebsocketController) GetPushNews(context *gin.Context) {
	// id := context.Query("userId")

	client, _ := upgrader.Upgrade(context.Writer, context.Request, nil)
	for {
		// 每隔两秒给前端推送一句消息“hello, WebSocket”
		err := client.WriteMessage(websocket.TextMessage, []byte("hello, WebSocket"))
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second * 2)
	}

}
