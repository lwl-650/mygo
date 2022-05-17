package apis

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebController struct {
	ws       *websocket.Conn
	inchan   chan []byte
	outchan  chan []byte
	exitchan chan bool
	isClose  bool
	err      error
}

// h := make(chan int)  声明通道

// https://blog.csdn.net/qq_43923045/article/details/122563592?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522165278010416781435454368%2522%252C%2522scm%2522%253A%252220140713.130102334..%2522%257D&request_id=165278010416781435454368&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~baidu_landing_v2~default-2-122563592-null-null.142^v10^pc_search_result_control_group,157^v4^new_style&utm_term=golang+websocket%E5%B9%BF%E6%92%AD&spm=1018.2226.3001.4187
func (con WebController) GetWeb(c *gin.Context) {
	var upgrader = websocket.Upgrader{
		// 读取存储空间大小
		ReadBufferSize: 1024,
		// 写入存储空间大小
		WriteBufferSize: 1024,
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	con.ws, con.err = upgrader.Upgrade(c.Writer, c.Request, nil)
	if con.err != nil {
		log.Print("upgrade:", con.err)
		return
	}
	defer con.ws.Close()
	for {

		_, message, err := con.ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		// getMap := make(map[string]interface{})
		// getMap["data"] = string(message)

		// err = ws.WriteMessage(mt, message)

		err = con.ws.WriteMessage(1, []byte(message))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func (con *WebController) Read() ([]byte, error) {
	var data []byte
	if con.isClose {
		return nil, errors.New("connection is closed")
	}
	data = <-con.inchan
	return data, nil
}

func (con *WebController) WriteMessage(data []byte) error {
	if con.isClose {
		return errors.New("connection is closed")
	}
	con.outchan <- data
	return nil
}
func (con *WebController) writeLoop() {
	for {
		// 每个case语句里必须是一个IO操作
		select {
		case data := <-con.outchan:
			if err := con.ws.WriteMessage(websocket.TextMessage, data); err != nil {
				con.exitchan <- true
				return
			}
		case <-con.exitchan:
			return
		}
	}
}
