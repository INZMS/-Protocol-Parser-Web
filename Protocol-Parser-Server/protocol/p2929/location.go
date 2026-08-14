package p2929

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"protocol-parser-server/parser/core"
)

var beijingLocation = time.FixedZone("Asia/Shanghai", 8*60*60)

func ParseLocation(p *Protocol2929, header *Header, data []byte) (*core.ParseResult, error) {
	property, tlvs, err := ParseReportProperty(header, data)
	if err != nil {
		return nil, err
	}
	body := data[headerLength : len(data)-trailerLength]
	location := property.PropertiesMap["location"].(map[string]interface{})
	status := property.PropertiesMap["locationStatus"].(map[string]interface{})
	vehicle := property.PropertiesMap["vehicleStatus"].(map[string]interface{})
	fields := headerFields(header, data)
	fields = append(fields,
		newField(5, "定位时间", 9, body[0:6], timeText(property.Timestamp), "北京时间（UTC+8），YYMMDDHHmmss压缩BCD"),
		newField(6, "纬度", 15, body[6:10], fmt.Sprint(location["lat"]), "DDMM.mmm，最高位为南纬符号"),
		newField(7, "经度", 19, body[10:14], fmt.Sprint(location["lng"]), "DDDMM.mmm，最高位为西经符号"),
		newField(8, "速度", 23, body[14:16], fmt.Sprintf("%v km/h", property.PropertiesMap["speed"]), "压缩BCD"),
		newField(9, "方向", 25, body[16:18], fmt.Sprintf("%v°", property.PropertiesMap["direction"]), "正北0度，顺时针"),
		newField(10, "定位/天线/电源状态", 27, body[18:19], fmt.Sprintf("0x%02X", body[18]), fmt.Sprintf("定位=%v，天线=%v，电源=%v", status["valid"], status["gpsAntenna"], status["power"])),
		newField(11, "未使用LLL", 28, body[19:22], hex.EncodeToString(body[19:22]), "协议保留"),
		newField(12, "车辆状态ABCD", 31, body[22:26], fmt.Sprint(vehicle["raw"]), fmt.Sprintf("%v，信号=%v，需应答=%v", vehicle["transport"], vehicle["signalStrength"], vehicle["needAck"])),
		newField(13, "未使用WWERTYU", 35, body[26:33], hex.EncodeToString(body[26:33]), "协议保留"),
		newField(14, "中心命令", 42, body[33:34], fmt.Sprintf("0x%02X", body[33]), "中心下发的主命令"),
	)
	for _, item := range tlvs {
		start := headerLength + locationBaseLength + item.Offset
		raw := data[start : start+int(item.Length)+2]
		fields = append(fields, newField(len(fields)+1, extensionName(item.Type), start, raw, extensionValue(property, item), fmt.Sprintf("扩展指令0x%04X，长度字段包含指令字", item.Type)))
	}
	fields = append(fields, trailerFields(len(fields)+1, data)...)
	return &core.ParseResult{Protocol: p.Name(), MessageID: hex.EncodeToString([]byte{header.Cmd}), MessageName: MessageName(header.Cmd), Length: len(data), Data: property, Raw: hex.EncodeToString(data), Fields: fields}, nil
}

func extensionName(command uint16) string {
	names := map[uint16]string{0x0024: "单基站", 0x0004: "AD电压", 0x0008: "电池电量", 0x00A3: "软件版本", 0x00A5: "终端型号", 0x0089: "扩展报警状态", 0x00A9: "多基站", 0x00B9: "WiFi热点", 0x00C5: "扩展定位状态", 0x00FB: "SIM ICCID", 0x00AE: "当前工作模式", 0xF000: "预设上报模式", 0xF001: "下次上报时间", 0x0030: "通讯信号强度", 0x0031: "GNSS卫星数"}
	if name, ok := names[command]; ok {
		return name
	}
	return fmt.Sprintf("未知扩展0x%04X", command)
}
func extensionValue(property *ReportProperty, item TLV) string {
	keys := map[uint16]string{0x0024: "singleBaseStation", 0x0004: "voltage", 0x0008: "battery", 0x00A3: "softwareVersion", 0x00A5: "terminalModel", 0x0089: "alarmStatus", 0x00A9: "baseStations", 0x00B9: "wifi", 0x00C5: "extendedLocationStatus", 0x00FB: "iccid", 0x00AE: "workMode", 0xF000: "reportMode", 0xF001: "nextReport", 0x0030: "signalStrength", 0x0031: "satellites"}
	if key, ok := keys[item.Type]; ok {
		value := property.PropertiesMap[key]
		switch item.Type {
		case 0x0008:
			if battery, ok := value.(map[string]interface{}); ok {
				return fmt.Sprintf("%.2f%%（%v/%v）", battery["percent"], battery["count"], battery["total"])
			}
		case 0x0089:
			if alarm, ok := value.(map[string]interface{}); ok {
				return fmt.Sprintf("运动=%v，SOS=%v，拆除=%v", alarm["moving"], alarm["sos"], alarm["removed"])
			}
		case 0x00A9:
			if stations, ok := value.(map[string]interface{}); ok {
				return baseStationsText(stations)
			}
		case 0x00B9:
			if wifi, ok := value.(map[string]interface{}); ok {
				return wifiText(wifi)
			}
		case 0x00C5:
			if status, ok := value.(map[string]interface{}); ok {
				labels := map[interface{}]string{"gps": "GPS定位", "wifi": "WiFi定位", "none": "未定位"}
				if label, exists := labels[status["mode"]]; exists {
					return label
				}
			}
		case 0x00AE:
			if mode, ok := value.(map[string]interface{}); ok {
				labels := map[interface{}]string{1: "闹钟模式", 3: "追踪模式", 4: "星期模式", 6: "月模式"}
				if label, exists := labels[mode["mode"]]; exists {
					return label
				}
			}
		case 0xF000:
			if mode, ok := value.(map[string]interface{}); ok {
				labels := map[interface{}]string{0: "闹钟模式", 1: "定时回传", 2: "星期模式"}
				if label, exists := labels[mode["mode"]]; exists {
					return label
				}
			}
		case 0xF001:
			if next, ok := value.(map[string]interface{}); ok {
				if timestamp, ok := next["timestamp"].(int64); ok {
					return fmt.Sprintf("%s（%v）", timeText(timestamp), next["environment"])
				}
			}
		}
		return fmt.Sprint(value)
	}
	return hex.EncodeToString(item.Value)
}

func baseStationsText(value map[string]interface{}) string {
	parts := []string{fmt.Sprintf("%v个基站（MCC=%v，MNC=%v）", value["count"], value["country"], value["operator"])}
	if stations, ok := value["stations"].([]map[string]int); ok {
		for index, station := range stations {
			parts = append(parts, fmt.Sprintf("#%d LAC=%d，CellID=%d，信号=%d", index+1, station["area"], station["tower"], station["signal"]))
		}
	}
	return strings.Join(parts, "；")
}

func wifiText(value map[string]interface{}) string {
	parts := []string{fmt.Sprintf("%v个WiFi热点", value["count"])}
	if hotspots, ok := value["hotspots"].([]map[string]interface{}); ok {
		for index, hotspot := range hotspots {
			parts = append(parts, fmt.Sprintf("#%d %v（%v dBm）", index+1, hotspot["mac"], hotspot["signal"]))
		}
	}
	return strings.Join(parts, "；")
}
func timeText(timestamp int64) string {
	return time.UnixMilli(timestamp).In(beijingLocation).Format("2006-01-02 15:04:05") + "（北京时间）"
}
