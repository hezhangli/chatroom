package main

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func login(userId int, password string) error {
	//fmt.Printf("userId %d, password %s\n", userId, password)
	//return nil

	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		return err
	}
	//延时关闭
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	//2.准备通过conn发送信息给服务器
	var msg message.Massage
	msg.Type = message.LoginMesType
	//3.创建一个loginMas结构体

	var loginMsg message.LoginMes
	loginMsg.UserId = userId
	loginMsg.Password = password
	//4.将loginMse序列化
	data, err := json.Marshal(loginMsg)
	if err != nil {
		return err
	}
	//5.将data赋值给msg.Data字段
	msg.Data = string(data)
	//6.将msg进行序列化
	data, err = json.Marshal(msg)
	if err != nil {
		return err
	}
	//7.data为要发送的数据
	//7.1先把data的长度发送给服务器
	//将获取到的data长度转换为表示长度的byte切片

	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	//发送长度
	write, err := conn.Write(buf[:4])
	if write != 4 || err != nil {
		return err
	}

	//fmt.Println("客户端发送消息的长度ok")
	//发送内容本身
	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	//休眠20s
	//time.Sleep(20 * time.Second)
	//fmt.Println("休眠了20s ")

	pkg, err := utils.readPkg(conn)
	if err != nil {
		return err
	}
	var loginResMsg message.LoginResMes
	err = json.Unmarshal([]byte(pkg.Data), &loginResMsg)
	if err != nil {
		return err
	}
	if loginResMsg.Code == 200 {
		fmt.Println("登录成功")
	}
	if loginResMsg.Code == 500 {
		return err
	}
	return nil
}
