package main

import (
	"bytes"
	"fmt"
)

func printTypeValue(slist ...interface{}) string {

	//字节缓冲作为快速字符串连接
	var b bytes.Buffer

	//遍历参数
	for _,s := range slist{


		//将interface{}类型格式化为字符串
		str := fmt.Sprintf("%v",s)

		//类型的字符串描述
		var typeString string

		//对s进行类型断言
		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		}

		b.WriteString("value:")
		b.WriteString(str)
		b.WriteString(" type:")
		b.WriteString(typeString)
		b.WriteString("\n")
	}
	return b.String()
}

func main() {
	fmt.Println(printTypeValue(100, "str", true))
}
