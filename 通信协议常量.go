package convention

import "time"

const (
	ReadTimeout = time.Minute * 5 //当服务器1分钟内没有收到任何数据，断开客户端连接
)

/*
//数据包结构
const (
	TypeLen       = 2                                   // 消息类型字节数组长度
	LenLen        = 2                                   // 消息长度字节数组长度
	SeqLen        = 4                                   // 消息seq字节数组长度
	UnixLen       = 8                                   //时间戳长度
	ContentMaxLen = 1024 * 16                           // 消息体最大长度
	HeadLen       = TypeLen + LenLen + SeqLen + UnixLen // 消息头部字节数组长度（消息类型字节数组长度+消息长度字节数组长度）
	BufLen        = ContentMaxLen + HeadLen             // 缓冲buffer字节数组长度
)
*/

const (
	CodeLen                  = 2 //uint16 Code长度
	CodeLenLen               = 0 + CodeLen
	ContentLen               = 2 //uint16   内容长度
	CodeContentLenLen        = CodeLenLen + ContentLen
	SeqLen                   = 4 //uint32  Seq长度
	CodeContentSeqLenLen     = CodeContentLenLen + SeqLen
	UnixLen                  = 8 //uint64 unix长度
	CodeContentSeqLenUnixLen = CodeContentSeqLenLen + UnixLen

	ContentMaxLen = 1024 * 16                // 消息体最大长度
	HeadLen       = CodeContentSeqLenUnixLen // 消息头部字节数组长度（消息类型字节数组长度+消息长度字节数组长度）
	BufLen        = ContentMaxLen + HeadLen  // 缓冲buffer字节数组长度
)
