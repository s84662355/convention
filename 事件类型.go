package convention

//事件类型 从1递增
const (
	EventConnectSucceed = iota + 1 //客户端 连接成功事件
	EventTest
	EventLogin            ////服务器端 登陆事件
	EventLoginSucceed     ////客户端 登陆成功事件
	EventHeartBeat        ////服务器 心跳事件
	EventConnectClose     ///客户端 关闭连接事件
	EventLoginFail        ///客户端 登陆失败事件
	EventLoginRepeat      ///客户端 重复登陆
	EventSendChatMessage  ///服务器端 发送聊天消息事件
	EventReachChatMessage ///客户端 获取聊天消息事件

)
