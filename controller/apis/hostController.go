package apis

import (
	"fmt"
	"mygo/util"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/net"
)

type HostController struct {
}

func (HostController) GetCpu(c *gin.Context) {

	info, _ := net.Connections("all") //可填入tcp、udp、tcp4、udp4等等
	fmt.Println(info)
	util.Success(c, info)
}
