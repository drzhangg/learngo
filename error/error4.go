package main

import (
	"fmt"
	"math"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0,fmt.Errorf("Area calculation failed,radius %0.2f is less than zero",radius)
			//errors.New("Area calculation failed,radius is less than zero")
	}
	return math.Pi * radius * radius,nil
}

func main() {
	radius := -20.0
	area,err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Area of circle %0.2f",area)
}
