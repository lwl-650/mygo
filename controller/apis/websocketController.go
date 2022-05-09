package apis

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// https://www.cnblogs.com/chnmig/p/14463837.html  websocket
type WebsocketController struct {
}

var (
	// 消息通道
	news = make(map[string]chan interface{})
	// websocket客户端链接池
	client = make(map[string]*websocket.Conn)
	// 互斥锁，防止程序对统一资源同时进行读写
	mux sync.Mutex
)

// api:/getPushNews接口处理函数
func (WebsocketController) GetPushNews(context *gin.Context) {
	id := context.Query("userId")

	log.Println(id + "=>websocket链接")
	// 升级为websocket长链接
	WsHandler(context.Writer, context.Request, id)
}

// api:/deleteClient接口处理函数
func (WebsocketController) DeleteClient(context *gin.Context) {
	id := context.Param("id")
	// 关闭websocket链接
	conn, exist := getClient(id)
	if exist {
		conn.Close()
		deleteClient(id)
	} else {
		context.JSON(http.StatusOK, gin.H{
			"mesg": "未找到该客户端",
		})
	}
	// 关闭其消息通道
	_, exist = getNewsChannel(id)
	if exist {
		deleteNewsChannel(id)
	}
}

// websocket Upgrader
var wsupgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WsHandler 处理ws请求
func WsHandler(w http.ResponseWriter, r *http.Request, id string) {
	var conn *websocket.Conn
	var err error
	var exist bool
	// 创建一个定时器用于服务端心跳
	pingTicker := time.NewTicker(time.Second * 10)
	conn, err = wsupgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}
	// 把与客户端的链接添加到客户端链接池中
	addClient(id, conn)

	// 获取该客户端的消息通道
	m, exist := getNewsChannel(id)
	fmt.Println(m, exist)
	if !exist {
		m = make(chan interface{})
		addNewsChannel(id, m)
	}

	// 设置客户端关闭ws链接回调函数
	conn.SetCloseHandler(func(code int, text string) error {
		deleteClient(id)
		log.Println(code)
		return nil
	})

	for {
		select {
		case content, _ := <-m:
			// 从消息通道接收消息，然后推送给前端
			fmt.Println("content", content)

			err = conn.WriteJSON(content)
			if err != nil {
				log.Println(err)
				conn.Close()
				deleteClient(id)
				return
			}
		case <-pingTicker.C:
			// 服务端心跳:每20秒ping一次客户端，查看其是否在线
			conn.SetWriteDeadline(time.Now().Add(time.Second * 2))
			err = conn.WriteMessage(websocket.PingMessage, []byte{})
			// fmt.Println("心")
			if err != nil {
				log.Println("send ping err:", err)
				conn.Close()
				deleteClient(id)
				return
			}
		}
	}
}

// 将客户端添加到客户端链接池
func addClient(id string, conn *websocket.Conn) {
	mux.Lock()
	client[id] = conn
	mux.Unlock()
}

// 获取指定客户端链接
func getClient(id string) (conn *websocket.Conn, exist bool) {
	mux.Lock()
	conn, exist = client[id]
	mux.Unlock()
	return
}

// 删除客户端链接
func deleteClient(id string) {
	mux.Lock()
	delete(client, id)
	log.Println(id + "websocket退出")
	mux.Unlock()
}

// 添加用户消息通道
func addNewsChannel(id string, m chan interface{}) {
	fmt.Println(id)
	mux.Lock()
	news[id] = m
	mux.Unlock()
}

// 获取指定用户消息通道
func getNewsChannel(id string) (m chan interface{}, exist bool) {
	mux.Lock()
	m, exist = news[id]
	mux.Unlock()
	return
}

// 删除指定消息通道
func deleteNewsChannel(id string) {
	mux.Lock()
	if m, ok := news[id]; ok {
		close(m)
		delete(news, id)
	}
	mux.Unlock()
}
