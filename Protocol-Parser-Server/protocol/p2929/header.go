package p2929

import (
	"encoding/hex"
	"fmt"

	"protocol-parser-server/parser/core"
)

const (
	headerLength  = 9
	trailerLength = 2
)

type Header struct {
	Cmd    byte
	Length uint16
	IP     []byte
}

func ParseHeader(data []byte) (*Header, error) {
	if len(data) < headerLength+trailerLength {
		return nil, fmt.Errorf("2929数据长度不足: 至少需要%d字节", headerLength+trailerLength)
	}
	if data[0] != 0x29 || data[1] != 0x29 {
		return nil, fmt.Errorf("不是2929协议")
	}
	if data[len(data)-1] != 0x0D {
		return nil, fmt.Errorf("2929报文缺少结束符0D")
	}
	header := &Header{Cmd: data[2], Length: uint16(data[3])<<8 | uint16(data[4]), IP: data[5:9]}
	// LEN 表示从包头到结束符的完整报文长度。
	if int(header.Length) != len(data) {
		return nil, fmt.Errorf("2929报文长度不一致: 声明%d字节, 实际%d字节", header.Length, len(data))
	}
	return header, nil
}

func headerFields(header *Header, data []byte) []core.Field {
	return []core.Field{
		newField(1, "包头", 0, data[0:2], "2929", "协议标识"),
		newField(2, "命令字", 2, data[2:3], fmt.Sprintf("0x%02X", header.Cmd), MessageName(header.Cmd)),
		newField(3, "报文长度", 3, data[3:5], fmt.Sprintf("%d", header.Length), "完整报文长度（字节）"),
		newField(4, "伪IP", 5, data[5:9], hex.EncodeToString(header.IP), "终端标识"),
	}
}

func newField(index int, name string, offset int, raw []byte, value, description string) core.Field {
	return core.Field{Index: index, Name: name, Offset: offset, Length: len(raw), Raw: hex.EncodeToString(raw), Value: value, Description: description}
}

func trailerFields(startIndex int, data []byte) []core.Field {
	return []core.Field{
		newField(startIndex, "校验码", len(data)-2, data[len(data)-2:len(data)-1], fmt.Sprintf("0x%02X", data[len(data)-2]), "包头至校验码前所有字节异或"),
		newField(startIndex+1, "结束符", len(data)-1, data[len(data)-1:], "0x0D", "报文结束标识"),
	}
}
