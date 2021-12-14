package main

import (
	"fmt"
	"net"
)

func TcpStart(port int) {
	fmt.Println("Reverse TCP Server Start... ")
	listen,err := net.Listen("tcp",fmt.Sprintf("0.0.0.0:%d",port)) // 改为从配置文件中进行读取
	if err != nil {
		fmt.Sprintf("[!] Listen Fail err: %s\n", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Sprintf("[!] Listen Accept Fail err: %s\n", err)
			continue
		}
		go AcceptProcess(conn)
	}
}

func AcceptProcess(conn net.Conn){
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {}
	}(conn)

	for  {
		buf := make([]byte, 512)
		num, err := conn.Read(buf)
		if err != nil {
			fmt.Sprintf("[!] Accept Data Reading err: %s\n",err)
			break
		}
		hexStr := fmt.Sprintf("%x", buf[:num])
		switch hexStr {
			case "300c020101600702010304008000":
				res := &result{
					host:   conn.RemoteAddr().String(),
					name:   "Log4j2Vuln",
					finger: hexStr,
				}
				conn.Close()
				socketChan <- true
				resultChan <- res
			default:
				conn.Close()
		}
	}
}

