package zrouter

import (
	"github.com/p0nley/zinx/ziface"
	"github.com/p0nley/zinx/zlog"
	"github.com/p0nley/zinx/znet"
)

type HelloZinxRouter struct {
	znet.BaseRouter
}

//HelloZinxRouter Handle
func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	zlog.Debug("Call HelloZinxRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	zlog.Debug("recv from client : msgId=0", ", data=", string(request.GetData()))

	err := request.GetConnection().SendBuffMsg([]byte("Hello Zinx Router V0.10"))
	if err != nil {
		zlog.Error(err)
	}
}
