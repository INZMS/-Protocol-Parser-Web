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
	if len(data) < 7 {
		return nil, fmt.Errorf("2929数据长度不足: 至少需要7字节")
	}
	if data[0] != 0x29 || data[1] != 0x29 {
		return nil, fmt.Errorf("不是2929协议")
	}
	if data[len(data)-1] != 0x0D {
		return nil, fmt.Errorf("2929报文缺少结束符0D")
	}
	header := &Header{Cmd: data[2], Length: uint16(data[3])<<8 | uint16(data[4])}
	if header.Cmd != 0x21 {
		if len(data) < headerLength+trailerLength {
			return nil, fmt.Errorf("2929命令0x%02X缺少4字节伪IP", header.Cmd)
		}
		header.IP = append([]byte(nil), data[5:9]...)
	}
	// LEN 表示包长字段之后第一字节至包尾的长度，整包长度为 LEN+5。
	if int(header.Length)+5 != len(data) {
		return nil, fmt.Errorf("2929报文长度不一致: 声明包长%d字节, 预期整包%d字节, 实际%d字节", header.Length, int(header.Length)+5, len(data))
	}
	return header, nil
}

func headerFields(header *Header, data []byte) []core.Field {
	fields := []core.Field{
		newField(1, "包头", 0, data[0:2], "2929", "协议标识"),
		newField(2, "命令字", 2, data[2:3], fmt.Sprintf("0x%02X", header.Cmd), MessageName(header.Cmd)),
		newField(3, "包长", 3, data[3:5], fmt.Sprintf("%d", header.Length), "从伪IP首字节至包尾的长度"),
	}
	if len(header.IP) == 4 {
		fields = append(fields, newField(4, "伪IP", 5, data[5:9], hex.EncodeToString(header.IP), fmt.Sprintf("终端标识，设备号%s", decodePseudoIP(header.IP))))
	}
	return fields
}

func bodyOffset(header *Header) int {
	if header.Cmd == 0x21 {
		return 5
	}
	return headerLength
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
