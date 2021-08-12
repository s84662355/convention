package convention

//事件类型 从1递增
const (
	EventConnectSucceed = iota + 1 //客户端 连接成功事件
	EventLogin                     ////服务器端 登陆事件
	EventHeartBeat                 ////服务器 心跳事件
)
