package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Massage struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的类型
}

// LoginMes 定义两个类型
type LoginMes struct {
	UserId   int    `json:"userId"`
	Password string `json:"password"`
}

// LoginResMes 返回响应及错误信息
type LoginResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
