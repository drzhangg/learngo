package main

import "fmt"

type areaError1 struct {
	err string
	length float64
	width float64
}

func (e *areaError1) Error() string {
	return e.err
}

func (e *areaError1) lengthNegative() bool {
	return e.length <0
}

func (e *areaError1) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width <0 {
		if err == "" {
			err = "width is less than zero"
		}else {
			err += ",width is less than zero"
		}
	}
	if err != "" {
		return 0,&areaError1{err,length,width}
	}
	return length * width,nil
}

func main() {
	length ,width := -5.0,-9.0
	area,err := rectArea(length,width)
	if err != nil {
		if err, ok := err.(*areaError1); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)
			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)
			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect",area)
}
