package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func writePkg(conn net.Conn, data []byte) error {

	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	//发送长度
	write, err := conn.Write(buf[:4])
	if write != 4 || err != nil {
		return err
	}

	//发送长度
	write, err = conn.Write(data)
	if write != int(pkgLen) || err != nil {
		return err
	}
	return nil
}

func readPkg(conn net.Conn) (mes message.Massage, err error) {

	buf := make([]byte, 8086)
	fmt.Println("读取客户端发送的消息")
	_, err = conn.Read(buf[:4])
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	//根据pkgLen读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body error")
		return
	}

	//将pkgLen反序列化
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		return message.Massage{}, err
	}

	return
}

// ServerPressLogin 处理登录请求
func ServerPressLogin(conn net.Conn, msg *message.Massage) error {
	//从msg中取出msg.data，并反序列化为Logmsg
	var logMsg message.LoginMes
	err := json.Unmarshal([]byte(msg.Data), &logMsg)
	if err != nil {
		return err
	}

	var resMsg message.Massage
	resMsg.Type = message.LoginResMesType

	var loginResType message.LoginResMes

	if logMsg.UserId == 123 && logMsg.Password == "123" {
		loginResType.Code = 200
	} else {
		loginResType.Code = 500
		fmt.Println("账号或者密码错误")
	}

	//序列化
	data, err := json.Marshal(loginResType)
	if err != nil {
		fmt.Println("序列化失败：", err)
		return err
	}

	resMsg.Data = string(data)

	marshal, err := json.Marshal(resMsg)
	if err != nil {
		return err
	}

	err = writePkg(conn, marshal)
	if err != nil {
		return err
	}

	return nil
}

// ServerPressMsg 判断消息类型
func ServerPressMsg(conn net.Conn, msg *message.Massage) error {
	switch msg.Type {
	case message.LoginMesType:
		err := ServerPressLogin(conn, msg)
		if err != nil {
			return err
		}

	case message.RegisterMesType:
	default:
		fmt.Println("消息类型不存在")
	}
	return nil
}

func process(conn net.Conn) {

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("闭包")
			return
		}
	}(conn)
	for {
		pkg, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端正常关闭，服务端也正常关闭 ")
				return
			} else {
				fmt.Println("read Pkg error :", err)
				return
			}

		}

		//fmt.Println("msg:", pkg)
		err = ServerPressMsg(conn, &pkg)
		if err != nil {
			return
		}
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
