package apis

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QiniuController struct{}
type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

func (QiniuController) Upload(c *gin.Context) {

	file, err := c.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, "请求失败")
		return
	}
	//获取文件名
	fileName := file.Filename
	fmt.Println(file.Header, file.Size, "文件名：", fileName)

	// accessKey := "DPfKARFB9gnPl1acy8NzvhMSpOQL9kLSCCN0sTqN"
	// secretKey := "K6fkYEpzhOqh5G6FVmrETyu0-3yO5PJha2lDzLvM"
	// mac := qbox.NewMac(accessKey, secretKey)
	// bucket := "vuetitle"
	// putPolicy := storage.PutPolicy{
	// 	Scope:      bucket,
	// 	ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	// }
	// upToken := putPolicy.UploadToken(mac)

	// fmt.Println(mac)
	// util.Success(c, upToken)

}
