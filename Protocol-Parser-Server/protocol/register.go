// 初始化注册入口
package protocol

import (
	"protocol-parser-server/parser/core"
	"protocol-parser-server/protocol/p2929"
)

func RegisterProtocols() {

	core.Register(
		p2929.New(),
	)

}
