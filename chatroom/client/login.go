package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"learngo/chatroom/common/message"
	"net"
)

func login(userId int, userPwd string) (err error) {
	//	fmt.Printf("userId = %d ,userPwd =  %s\n", userId, userPwd)

	//连接服务端
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net dial failed,err = ", err)
		return err
	}
	defer conn.Close()

	var msg message.Message
	var logMsg message.LoginMes

	msg.Type = message.LoginMesType

	logMsg.UserId = userId
	logMsg.UserPwd = userPwd

	//序列化logMsg数据，存到msg.Data里
	data, err := json.Marshal(logMsg)
	if err != nil {
		fmt.Println("logMsg marshal failed,err = ", err)
		return err
	}

	msg.Data = string(data)
	//序列化msg数据
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("msg marshal failed, err = ", err)
		return err
	}

	//这时候data就是要发送的信息
	//先把data的长度发送给服务器
	//先获取到 data的长度转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(bytes[:4]) //获取切片长度，从开始到4前面结束
	if n != 4 || err != nil {
		fmt.Println("conn write failed, err = ", err)
		return err
	}
	return
}
