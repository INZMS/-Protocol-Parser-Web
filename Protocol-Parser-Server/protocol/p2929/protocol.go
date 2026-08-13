// 2929协议解析器（协议入口/CMD分发）
package p2929

import (
	"protocol-parser-server/parser/core"
)

type Protocol2929 struct{}

// 创建2929协议实例

func New() core.Protocol {

	return &Protocol2929{}

}

// 协议名称

func (p *Protocol2929) Name() string {

	return "2929"

}

// 协议识别

func (p *Protocol2929) Match(data []byte) bool {

	if len(data) < 2 {

		return false

	}

	return data[0] == 0x29 &&
		data[1] == 0x29

}

// 协议解析入口
//
// 所有2929数据包都会进入这里
//
// 根据CMD分发到具体解析器

func (p *Protocol2929) Parse(data []byte) (*core.ParseResult, error) {

	// 解析固定包头
	err := CheckChecksum(data)

	if err != nil {

		return nil, err

	}

	header, err := ParseHeader(data)

	if err != nil {

		return nil, err

	}

	// 根据主信令分发

	switch header.Cmd {

	// 中心确认
	// case 0x21:

	// 	return ParseAck(
	// 		p,
	// 		header,
	// 		data,
	// 	)

	// 调度文本信息
	// case 0x3A:

	// 	return ParseDispatch(
	// 		p,
	// 		header,
	// 		data,
	// 	)

	// 校时应答
	// case 0xD7:

	// 	return ParseTimeSync(
	// 		p,
	// 		header,
	// 		data,
	// 	)

	// 一般位置数据
	case 0x80:

		return ParseLocation(
			p,
			header,
			data,
		)

	// 终端确认
	// case 0x85:

	// 	return ParseTerminalAck(
	// 		p,
	// 		header,
	// 		data,
	// 	)

	// // 申请设置参数
	// case 0xD8:

	// 	return ParseConfigRequest(
	// 		p,
	// 		header,
	// 		data,
	// 	)

	// 未实现指令
	default:

		return ParseCommon(
			p,
			header,
			data,
		)

	}

}
