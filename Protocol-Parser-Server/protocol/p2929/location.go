package p2929

import (
	"encoding/hex"
	"fmt"

	"protocol-parser-server/parser/core"
)

func ParseLocation(p *Protocol2929, header *Header, data []byte) (*core.ParseResult, error) {
	property, err := ParseReportProperty(header, data)
	if err != nil {
		return nil, err
	}
	body := data[headerLength : len(data)-trailerLength]
	properties := property.PropertiesMap
	coordinates := properties["location"].(map[string]interface{})
	fields := headerFields(header, data)
	fields = append(fields,
		newField(5, "定位时间", 9, body[0:6], fmt.Sprint(property.Timestamp), "设备上报时间（Unix毫秒）"),
		newField(6, "纬度", 15, body[6:10], fmt.Sprint(coordinates["lat"]), "BCD转十进制度"),
		newField(7, "经度", 19, body[10:14], fmt.Sprint(coordinates["lng"]), "BCD转十进制度"),
		newField(8, "速度", 23, body[14:16], fmt.Sprint(properties["speed"]), "BCD速度值"),
		newField(9, "方向", 25, body[16:18], fmt.Sprint(properties["direction"]), "BCD方向值"),
		newField(10, "状态", 27, body[18:19], fmt.Sprintf("0x%02X", body[18]), fmt.Sprintf("定位有效=%t", properties["hasValidLocation"])),
		newField(11, "位置保留域", 28, body[19:locationBaseLength], hex.EncodeToString(body[19:locationBaseLength]), "位置数据固定区保留字段"),
	)
	if len(body) > locationBaseLength {
		fields = append(fields, newField(12, "扩展数据", headerLength+locationBaseLength, body[locationBaseLength:], hex.EncodeToString(body[locationBaseLength:]), "TLV扩展域"))
	}
	fields = append(fields, trailerFields(len(fields)+1, data)...)
	return &core.ParseResult{Protocol: p.Name(), MessageID: hex.EncodeToString([]byte{header.Cmd}), MessageName: MessageName(header.Cmd), Length: int(header.Length), Data: property, Raw: hex.EncodeToString(data), Fields: fields}, nil
}
