package message

const (
	LoginMesType    = "LoginMes"
	LoginMesResType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"` //消息的类型
	Data string `json:"data"` //数据
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code int    `json:"code"`
	Err  string `json:"err"`
}
