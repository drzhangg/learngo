package main

import "fmt"

//电子支付方式
type Alipay struct {
}

//为Alipay添加CanUseFaceID()方法，表示电子支付方式支持刷脸
func (a *Alipay) CanUserFaceID() {
	
}

//现金支付方法
type Cash struct {
}

//为Cash添加Stolen()方法，表示现金支付方式会出现偷窃情况
func (a *Cash) Stolen() {
}

//具备刷脸特性的接口
type ContainCanUseFaceID interface {
	CanUseFaceID()
}

//具备被偷特性的接口
type ContainStolen interface {
	Stolen()
}

//打印支付方式具备的特点
func print(payMethod interface{}) {
	switch payMethod.(type) {
	case ContainCanUseFaceID:
		fmt.Printf("%T can use faceid\n",payMethod)
	case ContainStolen:
		fmt.Printf("%T may be stolen\n",payMethod)
	}
}

func main() {

	print(new(Alipay))

	print(new(Cash))
}

