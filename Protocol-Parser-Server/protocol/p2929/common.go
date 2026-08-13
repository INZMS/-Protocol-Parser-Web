package p2929

import (
	"encoding/hex"

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

		Length: int(header.Length),

		Raw: hex.EncodeToString(data),

		Fields: append(headerFields(header, data), trailerFields(5, data)...),

		Data: map[string]interface{}{
			"body": hex.EncodeToString(data[headerLength : len(data)-trailerLength]),
		},
	}

	return result, nil

}
