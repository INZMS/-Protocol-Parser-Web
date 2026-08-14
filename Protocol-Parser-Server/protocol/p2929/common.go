package p2929

import (
	"encoding/hex"
	"fmt"

	"protocol-parser-server/parser/core"
)

// 未知CMD通用解析

func ParseCommon(
	p *Protocol2929,
	header *Header,
	data []byte,
) (
	*core.ParseResult,
	error,
) {

	result := &core.ParseResult{

		Protocol: p.Name(),

		MessageID: hex.EncodeToString(
			[]byte{
				header.Cmd,
			},
		),

		MessageName: MessageName(header.Cmd),

		Length: len(data),

		Raw: hex.EncodeToString(data),
	}
	fields := headerFields(header, data)
	body := data[bodyOffset(header) : len(data)-trailerLength]
	if header.Cmd == 0x21 && len(body) >= 3 {
		fields = append(fields,
			newField(len(fields)+1, "原包校验码", 5, body[0:1], fmt.Sprintf("0x%02X", body[0]), "收到的数据包校验值"),
			newField(len(fields)+1, "原包主信令", 6, body[1:2], fmt.Sprintf("0x%02X", body[1]), "收到的数据包主信令"),
			newField(len(fields)+1, "原包子信令", 7, body[2:3], fmt.Sprintf("0x%02X", body[2]), "无子信令时保留"),
		)
	}
	result.Fields = append(fields, trailerFields(len(fields)+1, data)...)
	result.Data = map[string]interface{}{"body": hex.EncodeToString(body)}

	return result, nil

}
