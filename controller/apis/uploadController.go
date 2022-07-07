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
		fmt.Println(file.Size, "文件名：", fileName)
		c.SaveUploadedFile(file, fileName)
		ExcelParse(fileName)
	}
}

func ExcelParse(fileName string) []string {
	// filePath := "../../" + fileName
	xlFile, err := xlsx.OpenFile("D:/work/新建文件夹 (2)/mygo/z.xlsx")
	fmt.Println(err, "123123123")
	//获取行数
	length := len(xlFile.Sheets[0].Rows)
	//开辟除表头外的行数的数组内存
	resourceArr := make([]string, length-1)
	//遍历sheet
	for _, sheet := range xlFile.Sheets {
		//遍历每一行
		for rowIndex, row := range sheet.Rows {
			//跳过第一行表头信息
			if rowIndex == 0 {
				// for _, cell := range row.Cells {
				//  text := cell.String()
				//  fmt.Printf("%s\n", text)
				// }
				continue
			}
			//遍历每一个单元
			for cellIndex, cell := range row.Cells {
				text := cell.String()
				if text != "" {
					//如果是每一行的第一个单元格
					if cellIndex == 0 {
						resourceArr[rowIndex-1] = text
					}
				}
			}
		}
	}
	fmt.Println(resourceArr)
	return resourceArr

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
	fmt.Println(result, "resultresultresultresult")
	return result
}
