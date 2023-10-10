package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("闭包")
			return
		}
	}(conn)
	//调用哟个总控
	procrsser := &Processor{Conn: conn}
	err := procrsser.SendMessage()
	if err != nil {
		fmt.Println("客户端和服务器端通信错误：", err)
		return
	}
}

func main() {
	//提示信息
	//监听端口
	//如果监听成功，等待客户端来连接服务器，并启动一个协程来和客户端保持通信

	fmt.Println("服务端在监听端口8889")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen error", err)
		return
	}

	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			return
		}
	}(listen)

	for {
		fmt.Println("等待客户端连接服务器")
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept error", err)
		}

		go process(accept)
	}
}
