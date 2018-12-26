package handleSession

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleSession(conn net.Conn, exitChan chan int) {

	fmt.Println("Session started:")

	//创建一个网络连接数据的读取器
	reader := bufio.NewReader(conn)

	for{
		str,err := reader.ReadString('\n')

		if err == nil {
			//去掉字符串尾部的回车
			str = strings.TrimSpace(str)

			//处理Telnet指令
			if !processTelnetCommand(str,exitChan){
				conn.Close()
				break
			}

			conn.Write([]byte(str + "\r\n"))
		}else {
			fmt.Println("Session closed")
			conn.Close()
			break
		}
	}
}
