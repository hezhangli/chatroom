package main

import (
	"chatroom/common/message"
	process2 "chatroom/server/process"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// ServerPressMsg 判断消息类型
func (this *Processor) ServerPressMsg(msg *message.Massage) error {

	switch msg.Type {
	case message.LoginMesType:
		up := &process2.UserProcess{Conn: this.Conn}
		err := up.ServerPressLogin(msg)
		if err != nil {
			return err
		}

	case message.RegisterMesType:
	default:
		fmt.Println("消息类型不存在")
	}
	return nil
}

func (this *Processor) SendMessage() error {

	tf := &utils.Transfers{Conn: this.Conn}
	pkg, err := tf.ReadPkg()
	if err != nil {
		if err == io.EOF {
			fmt.Println("客户端正常关闭，服务端也正常关闭 ")
			return err
		} else {
			fmt.Println("read Pkg error :", err)
			return err
		}
	}

	//fmt.Println("msg:", pkg)
	err = this.ServerPressMsg(&pkg)
	if err != nil {
		return err
	}
	return nil
}
