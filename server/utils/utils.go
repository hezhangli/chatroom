package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfers struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfers) WritePkg(data []byte) (err error) {

	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	//发送长度
	write, err := this.Conn.Write(this.Buf[:4])
	if write != 4 || err != nil {
		fmt.Println("conn write err:", err)
		return
	}

	//发送长度
	write, err = this.Conn.Write(data)
	if write != int(pkgLen) || err != nil {
		fmt.Println("conn write err:", err)
		return
	}
	return
}

func (this *Transfers) ReadPkg() (mes message.Massage, err error) {

	//buf := make([]byte, 8086)
	fmt.Println("读取客户端发送的消息")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body error")
		return
	}

	//将pkgLen反序列化
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json unmarshal err:", err)
		return
	}

	return
}
