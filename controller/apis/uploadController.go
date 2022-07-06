package apis

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

type UploadController struct{}

func (UploadController) Xlsx(c *gin.Context) {
	file, err := c.FormFile("file")
	fileName := file.Filename
	if err == nil {
		fmt.Println(file, file.Size, "文件名：", fileName)

		// AnalyzeExcel(file)
	}
}

//解析一个学号+姓名的表，并返回一个n×2的学号姓名字符串数组
func AnalyzeExcel(path string) [1000][2]string {
	xlFile, err := xlsx.OpenFile(path) //打开文件
	if err != nil {
		log.Println(err)
	}
	var result [1000][2]string
	for _, sheet := range xlFile.Sheets { //遍历sheet层
		for rowIndex, row := range sheet.Rows { //遍历row层
			if rowIndex > 0 {
				if len(row.Cells) < 2 {
					break
				}
				for cellIndex, cell := range row.Cells { //遍历cell层
					text := cell.String() //把单元格的内容转成string
					if len(text) == 0 {
						break
					}
					result[rowIndex-1][cellIndex] = text //为数组赋值
				}
			}
		}
		break //这里直接break是因为我这个测试用的文件只有1个sheet
	}
	for i := 0; i < 3; i++ {
		fmt.Println(result[i][0] + result[i][1])
	}
	return result
}
