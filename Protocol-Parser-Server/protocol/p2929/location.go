// 0x80位置解析
package p2929

import (
	"encoding/hex"
	"protocol-parser-server/parser/core"
)

func ParseLocation(
	p *Protocol2929,
	header *Header,
	data []byte,
) (
	*core.ParseResult, error,
) {

	property, err := ParseReportProperty(
		header,
		data,
	)

	if err != nil {

		return nil, err

	}

	result := &core.ParseResult{

		Protocol: p.Name(),

		MessageID: hex.EncodeToString([]byte{header.Cmd}),

		MessageName: "一般位置数据",

		Length: int(header.Length),

		Data: property,

		Raw: hex.EncodeToString(data),
	}

	return result, nil

}
