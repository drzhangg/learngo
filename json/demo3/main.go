package main

import (
	"encoding/json"
	"fmt"
)

type Snapshot struct {
	ID       int     `json:"ID"`
	TagName  string  `json:"TagName"`
	Time     string  `json:"Time"`
	Value    float64 `json:"Value"`
	Quality  int     `json:"Quality"`
	Error    int     `json:"Error"`
	ErrorMsg string  `json:"ErrorMsg"`
}
type mydata struct {
	md []Snapshot
}

func main() {

	str := `[
  {
    "ID": 1878,
    "TagName": "SF2WT.x1_zjs_sfc_ps20kt_4-5_100-1239_pv:1",
    "Time": "2019-11-27 22:28:12.000",
    "Value": "1009.375",
    "Quality": 0,
    "Error": 0,
    "ErrorMsg": "错误码为：0;错误信息为：操作成功完成"
  },
  {
    "ID": 1879,
    "TagName": "SF2WT.x1_zjs_sfc_ps20kt_4-5_100-1239_sum:1",
    "Time": "2019-11-27 22:28:10.000",
    "Value": "842147",
    "Quality": 0,
    "Error": 0,
    "ErrorMsg": "错误码为：0;错误信息为：操作成功完成"
  }]
`
	s := []Snapshot{}
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

}
