package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"src/github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

type File struct {
	Item []Date `json:"item"`
}

type Date struct {
	Name    string  `json:"name"`
	Request Request `json:"request"`
}
type Request struct {
	Method string `json:"method"`
	//Header      []Header `json:"header"`
	Body        Body   `json:"body"`
	Url         Url    `json:"url"`
	Description string `json:"description"`
}

type Header struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Body struct {
	//Mode string `json:"mode"`
	Raw string `json:"raw"`
}

type Url struct {
	Raw string `json:"raw"`
}

func main() {

	var fileData File
	fff := excelize.NewFile()
	files, _ := ioutil.ReadDir("./excel/file/")
	for _, f := range files {
		//fmt.Println(f.Name())
		bytes, err := ioutil.ReadFile("./excel/file/" + f.Name())
		if err != nil {
			fmt.Println("open file err,", err)
		}

		//fmt.Println(string(bytes))

		err = json.Unmarshal(bytes, &fileData)
		if err != nil {
			fmt.Println("unmarshal failed,", err)
		}

	}

	//fmt.Println("waimian----",fileData.Item)
	//单个接口

	for k, v := range fileData.Item {
		fmt.Println(k, v)
		var datas = []Date{}
		datas = append(datas, v)

		for k1, v1 := range datas {
			//统一获取异常类型
			//106.52.166.84/api/r/businesdetailsvr/GetAbnTypeData
			//POST
			//v.Request.Body.Raw
			fff.SetCellValue("Sheet1", "A1", "接口名称")
			fff.SetCellValue("Sheet1", "B1", "接口url")
			fff.SetCellValue("Sheet1", "C1", "接口请求方式")
			fff.SetCellValue("Sheet1", "D1", "接口调用参数")
			fff.SetCellValue("Sheet1", "A"+strconv.Itoa(k1+2), v1.Name)
			fff.SetCellValue("Sheet1", "B"+strconv.Itoa(k1+2), v1.Request.Url.Raw)
			fff.SetCellValue("Sheet1", "C"+strconv.Itoa(k1+2), v1.Request.Method)
			fff.SetCellValue("Sheet1", "D"+strconv.Itoa(k1+2), v1.Request.Body.Raw)
		}

	}

	err := fff.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
