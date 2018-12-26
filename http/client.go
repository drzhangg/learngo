package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var writeStr,readStr = make([]byte,1024), make([]byte,1024)

func read(conn net.Conn) {
	for{
		length,err := conn.Read(readStr)
		if err!=nil{
			fmt.Printf("Error when read from server. Error:%s\n",err)
			os.Exit(0)
		}
		fmt.Println(string(readStr[:length]))
	}
}

func main() {
	var(
		host = "localhost"
		port = "8000"
		remote = host + ":" + port
		reader = bufio.NewReader(os.Stdin)
	)

	con,err := net.Dial("tcp",remote)
	if err != nil {
		fmt.Println("Server not found.")
		os.Exit(-1)
	}
	defer con.Close()

	fmt.Println("Connection OK.")
	fmt.Printf("Enter your name:")
	fmt.Scanf("%s",&writeStr)
	in,err := con.Write([]byte(writeStr))
	if err != nil {
		fmt.Printf("Error when send to server:%d\n",in)
		os.Exit(0)
	}
	fmt.Println("Now begin to talk")
	go read(con)

	for{
		writeStr,_,_ = reader.ReadLine()
		if string(writeStr) == "quit" {
			fmt.Println("Communication terminated.")
			os.Exit(1)
		}

		in,err:= con.Write([]byte(writeStr))
		if err != nil {
			fmt.Printf("Error when send to server:%d\n",in)
			os.Exit(0)
		}
	}
}
