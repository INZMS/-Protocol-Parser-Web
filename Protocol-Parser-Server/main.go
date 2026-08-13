// main.go
package main

import (
	"protocol-parser-server/protocol"
	"protocol-parser-server/router"
)

func main() {
	// 注册所有协议
	protocol.RegisterProtocols()

	// 初始化HTTP路由
	r := router.InitRouter()

	// 启动服务
	r.Run(":8080")

}
