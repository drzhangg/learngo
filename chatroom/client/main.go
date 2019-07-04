package main

import "fmt"

//定义两个变量，一个表示用户id，一个表示用户密码
var userId int
var userPwd string

func main() {

	//接收用户的选择
	var key int
	//判断是否还继续显示菜单
	var loop = true

	fmt.Println("------欢迎登录多人聊天系统------")
	fmt.Println("\t\t\t 1 登录聊天室")
	fmt.Println("\t\t\t 2 注册用户")
	fmt.Println("\t\t\t 3 退出聊天室")
	fmt.Println("\t\t\t 请选择(1-3)")

	fmt.Scanf("%d\n", &key)

	for loop {
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出聊天室")
			loop = false
		default:
			break
		}
	}

	//根据用户的输入，显示新的提示信息
	if key == 1 {
		//用户登录
		fmt.Println("请输入用户的id：")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户的密码：")
		fmt.Scanf("%s\n", &userPwd)
		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("密码错误")
		}else {
			fmt.Println("登录成功")
		}
		//fmt.Println(loop)

	}
}
