package apis

import (
	"mygo/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type HostController struct {
}

func (HostController) GetCpu(c *gin.Context) {
	rate := make(map[string]interface{})
	// rate["cpu使用率"] = GetCpuPercent()
	// rate["内存使用率"] = GetMemPercent()
	// rate["磁盘使用率"] = GetDiskPercent()

	// 本机信息
	// localMachine, _ := host.Info()
	// fmt.Println(localMachine)

	// cpu
	// getCpu, _ := cpu.Info()
	// fmt.Println(getCpu)

	// 内存
	// memory, _ := mem.VirtualMemory()
	// fmt.Println(memory)
	// 交换分区
	// swapPartition, _ := mem.SwapMemory()
	// fmt.Println(swapPartition)

	// 获取磁盘信息
	info, _ := disk.Partitions(true) //所有分区
	// fmt.Println(info)
	info2, _ := disk.Usage("D:") //指定某路径的硬盘使用情况
	// fmt.Println(info2)
	info3, _ := disk.IOCounters() //所有硬盘的io信息
	// fmt.Println(info3)

	rate["所有分区"] = info
	rate["指定某路径的硬盘使用情况"] = info2
	rate["所有硬盘的io信息"] = info3
	util.Success(c, rate)
}

// cpu使用率
func GetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

// 内存使用率
func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

// 磁盘使用率
func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}
