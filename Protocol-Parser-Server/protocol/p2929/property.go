package p2929

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type ReportProperty struct {
	MessageType      string                 `json:"messageType"`
	MessageID        string                 `json:"messageId"`
	DeviceNo         string                 `json:"deviceNo"`
	ProductNo        string                 `json:"productNo"`
	Timestamp        int64                  `json:"timestamp"`
	ReceiptTimestamp int64                  `json:"receiptTimestamp"`
	PropertiesMap    map[string]interface{} `json:"propertiesMap"`
	AsyncSave        bool                   `json:"asyncSave"`
	FlowType         string                 `json:"flowType"`
}

func ParseReportProperty(header *Header, data []byte) (*ReportProperty, error) {
	if len(data) < headerLength+locationBaseLength+trailerLength {
		return nil, fmt.Errorf("0x80报文长度不足: 至少需要%d字节, 实际%d字节", headerLength+locationBaseLength+trailerLength, len(data))
	}
	body := data[headerLength : len(data)-trailerLength]
	location, err := ParseLocationBase(body)
	if err != nil {
		return nil, err
	}
	result := &ReportProperty{MessageType: "reportProperty", MessageID: fmt.Sprintf("%02x", header.Cmd), DeviceNo: hex.EncodeToString(header.IP), Timestamp: location.DeviceTime, PropertiesMap: make(map[string]interface{}), AsyncSave: true, FlowType: "upstream"}
	result.PropertiesMap["location"] = map[string]interface{}{"lat": location.Lat, "lng": location.Lng}
	result.PropertiesMap["speed"] = location.Speed
	result.PropertiesMap["direction"] = location.Direction
	result.PropertiesMap["status"] = location.Status
	result.PropertiesMap["hasValidLocation"] = location.HasValidLocation
	result.PropertiesMap["deviceTime"] = location.DeviceTime
	extension := body[locationBaseLength:]
	result.PropertiesMap["rawExtension"] = hex.EncodeToString(extension)
	for _, tlv := range ParseTLV(extension) {
		switch tlv.Type {
		case 0x0018:
			result.PropertiesMap["singleBaseStation"] = string(tlv.Value)
		case 0x0006:
			result.PropertiesMap["wifi"] = parseWifi(tlv.Value)
		case 0x0005:
			parseProduct(result, tlv.Value)
		case 0x0004:
			parseDeviceParam(result, tlv.Value)
		default:
			result.PropertiesMap[fmt.Sprintf("unknown_%04x", tlv.Type)] = hex.EncodeToString(tlv.Value)
		}
	}
	return result, nil
}

func parseWifi(data []byte) map[string]interface{} {
	return map[string]interface{}{"count": 1, "hotspots": []string{string(data)}}
}
func parseProduct(result *ReportProperty, data []byte) {
	parts := strings.Split(string(data), "_")
	if len(parts) > 0 {
		result.ProductNo = parts[0]
	}
}
func parseDeviceParam(result *ReportProperty, data []byte) {
	if len(data) < 2 {
		return
	}
	value := int(data[0])<<8 | int(data[1])
	result.PropertiesMap["deviceParamRaw"] = hex.EncodeToString(data)
	result.PropertiesMap["deviceBatteryVoltage"] = float64(value) / 1000
}

func parse2929Time(data []byte) (int64, error) {
	if len(data) != 6 {
		return 0, fmt.Errorf("BCD时间需要6字节, 实际%d字节", len(data))
	}
	for _, value := range data {
		if value>>4 > 9 || value&0x0f > 9 {
			return 0, fmt.Errorf("无效BCD值%02X", value)
		}
	}
	year := 2000 + bcd(data[0])
	month := time.Month(bcd(data[1]))
	day := bcd(data[2])
	hour := bcd(data[3])
	minute := bcd(data[4])
	second := bcd(data[5])
	t := time.Date(year, month, day, hour, minute, second, 0, time.UTC)
	if t.Year() != year || t.Month() != month || t.Day() != day || t.Hour() != hour || t.Minute() != minute || t.Second() != second {
		return 0, fmt.Errorf("无效日期时间")
	}
	return t.UnixMilli(), nil
}
func bcd(v byte) int { return int(v>>4)*10 + int(v&0x0f) }
