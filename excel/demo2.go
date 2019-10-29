package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"src/github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

func main() {
	var fileDatas []File
	var fileData File
	fff := excelize.NewFile()
	files, _ := ioutil.ReadDir("./excel/file/")
	for _, f := range files {
		bytes, err := ioutil.ReadFile("./excel/file/" + f.Name())
		if err != nil {
			fmt.Println("open file err,", err)
		}
		err = json.Unmarshal(bytes, &fileData)
		if err != nil {
			fmt.Println("unmarshal failed,", err)
		}
		fileDatas = append(fileDatas, fileData)
		//fmt.Println(reflect.TypeOf(k1).String())
		fmt.Println(len(fileDatas))
		//fmt.Println(k1)

		for k, v := range fileData.Item {
			//fmt.Println(len(fileData.Item))
			fff.SetCellValue("Sheet"+strconv.Itoa(len(fileDatas)), "A1", "接口名称")
			fff.SetCellValue("Sheet"+strconv.Itoa(len(fileDatas)), "B1", "接口url")
			fff.SetCellValue("Sheet"+strconv.Itoa(len(fileDatas)), "C1", "接口请求方式")
			fff.SetCellValue("Sheet"+strconv.Itoa(len(fileDatas)), "D1", "接口调用参数")
			fff.SetCellValue("Sheet"+strconv.Itoa(len(fileDatas)), "A"+strconv.Itoa(k+2), v.Name)
			fff.SetCellValue("Sheet"+strconv.Itoa(len(fileDatas)), "B"+strconv.Itoa(k+2), v.Request.Url.Raw)
			fff.SetCellValue("Sheet"+strconv.Itoa(len(fileDatas)), "C"+strconv.Itoa(k+2), v.Request.Method)
			fff.SetCellValue("Sheet"+strconv.Itoa(len(fileDatas)), "D"+strconv.Itoa(k+2), v.Request.Body.Raw)
			fff.SetCellValue("Sheet"+strconv.Itoa(len(fileDatas)), "D"+strconv.Itoa(k+2), v.Request.Body.Raw)
			//excelName := `./book1.xlsx`
			excelName := fmt.Sprintf("./%s.xlsx",f.Name())
			err = fff.SaveAs(excelName)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func removeDup(a []int64) []int64 {
	i := 0
	for j := 1; j < len(a); j++ {
		if a[i] != a[j] {
			i++
			a[i] = a[j]
		}
	}
	return a[:i+1]
}

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
