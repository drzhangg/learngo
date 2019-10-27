package main

import (
	"fmt"
	"net"
)

func main() {
	listen, _ := net.Listen("tcp",":9090")

	port := listen.Addr().(*net.TCPAddr).Port
	ip := listen.Addr().(*net.TCPAddr).IP
	fmt.Println(ip,port)


	var ips []string
	addrs,_ := net.InterfaceAddrs()
	for _,a := range addrs{
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	fmt.Println(ips)
}
