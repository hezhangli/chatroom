package process

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) ServerPressLogin(msg *message.Massage) error {
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

	tf := &utils.Transfers{Conn: this.Conn}

	err = tf.WritePkg(marshal)
	if err != nil {
		return err
	}

	return nil
}
