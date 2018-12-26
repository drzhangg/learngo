package main

import "fmt"

//声明一个解析错误
type ParseError struct {
	FileName string
	Line int
}

//实现error接口，返回错误描述
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d",e.FileName,e.Line)
}

//创建一些解析错误
func newParseError(filename string, line int) error {
	return &ParseError{filename,line}
}

func main() {
	var e error
	e = newParseError("main.go",1)

	fmt.Println(e.Error())

	switch detail := e.(type) {
	case *ParseError:
		fmt.Printf("Filename:%s Line:%d\n",detail.FileName,detail.Line)
	default:
		fmt.Println("other error")
	}
}
