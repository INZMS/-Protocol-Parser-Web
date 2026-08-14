package p2929

import "fmt"

type TLV struct {
	Type   uint16
	Length uint16
	Value  []byte
	Offset int
}

func ParseTLV(data []byte) ([]TLV, error) {
	list := make([]TLV, 0)
	for offset := 0; offset < len(data); {
		if len(data)-offset < 4 {
			return nil, fmt.Errorf("扩展数据在偏移%d处不足4字节", offset)
		}
		length := uint16(data[offset])<<8 | uint16(data[offset+1])
		command := uint16(data[offset+2])<<8 | uint16(data[offset+3])
		if length < 2 {
			return nil, fmt.Errorf("扩展指令0x%04X长度%d无效", command, length)
		}
		total := int(length) + 2
		if offset+total > len(data) {
			return nil, fmt.Errorf("扩展指令0x%04X声明%d字节, 剩余仅%d字节", command, total, len(data)-offset)
		}
		list = append(list, TLV{Type: command, Length: length, Value: data[offset+4 : offset+total], Offset: offset})
		offset += total
	}
	return list, nil
}
