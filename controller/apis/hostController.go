package apis

import (
	"fmt"
	"mygo/util"
	"net"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type HostController struct {
}

func (HostController) GetCpu(c *gin.Context) {
	rate := make(map[string]interface{})
	// getP := make([]interface{}, 1)
	getP := []interface{}{}
	var total, usedT, freeT uint64

	// 本机信息
	localMachine, _ := host.Info()
	// fmt.Println(localMachine)

	// cpu
	getCpu, _ := cpu.Info()
	// fmt.Println(getCpu)

	// 内存
	memory, _ := mem.VirtualMemory()
	// fmt.Println(memory)
	// 交换分区
	// swapPartition, _ := mem.SwapMemory()
	// fmt.Println(swapPartition)

	// 获取磁盘信息
	info, _ := disk.Partitions(true) //所有分区
	// fmt.Println(info)

	for one := range info {
		// fmt.Println(info[one].Device)
		info2, _ := disk.Usage(info[one].Device)
		total += info2.Total
		usedT += info2.Used
		freeT += info2.Free
		getP = append(getP, info2)
	}

	// cpu使用率
	rate["GetCpuPercent"] = GetCpuPercent()
	// 内存使用率
	rate["GetMemPercent"] = GetMemPercent()
	//磁盘使用率
	rate["GetDiskPercent"] = GetDiskPercent()

	// info2, _ := disk.Usage("D:") //指定某路径的硬盘使用情况
	// fmt.Println(info2)
	// info3, _ := disk.IOCounters() //所有硬盘的io信息
	// fmt.Println(info3)

	rate["localMachine"] = localMachine
	rate["cpu"] = getCpu
	rate["memory"] = memory

	rate["diskUsed"] = usedT
	rate["disk"] = getP
	rate["diskTotal"] = total
	rate["diskHard"] = freeT

	gos := make(map[string]interface{})
	gos["demoPath"] = demoPath()
	gos["goPath"] = goPath()
	gos["version"] = runtime.Version()

	rate["go"] = gos

	util.Success(c, rate)
	go getIp()

}

func getIp() {
	interface_list, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}
	var byName *net.Interface
	var addrList []net.Addr
	var oneAddrs []string
	for _, i := range interface_list {
		byName, err = net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
			return
		}
		addrList, err = byName.Addrs()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, oneAddr := range addrList {
			oneAddrs = strings.SplitN(oneAddr.String(), "/", 2)
			fmt.Println(i.Name, "-", oneAddrs[0])
		}
	}
}

func demoPath() string {
	_, filename, _, ok := runtime.Caller(1)
	var cwdPath string
	if ok {
		cwdPath = path.Join(path.Dir(filename), "../../") // the the main function file directory
	} else {
		cwdPath = "./"
	}
	fmt.Println("demopath:", cwdPath)
	return cwdPath
}

func goPath() string {
	_, filename, _, ok := runtime.Caller(2)
	var cwdPath string
	if ok {
		cwdPath = path.Join(path.Dir(filename), "../../../../../") // the the main function file directory
	} else {
		cwdPath = "./"
	}
	fmt.Println("gopath:", cwdPath)
	return cwdPath
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
