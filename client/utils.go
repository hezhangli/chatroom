package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
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
