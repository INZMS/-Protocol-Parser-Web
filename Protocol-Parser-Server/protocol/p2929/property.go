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
	ProductNo        string                 `json:"productNo,omitempty"`
	Timestamp        int64                  `json:"timestamp"`
	ReceiptTimestamp int64                  `json:"receiptTimestamp,omitempty"`
	PropertiesMap    map[string]interface{} `json:"propertiesMap"`
	AsyncSave        bool                   `json:"asyncSave"`
	FlowType         string                 `json:"flowType"`
}

func ParseReportProperty(header *Header, data []byte) (*ReportProperty, []TLV, error) {
	if len(data) < headerLength+locationBaseLength+trailerLength {
		return nil, nil, fmt.Errorf("0x80报文长度不足: 至少需要%d字节, 实际%d字节", headerLength+locationBaseLength+trailerLength, len(data))
	}
	body := data[headerLength : len(data)-trailerLength]
	location, err := ParseLocationBase(body)
	if err != nil {
		return nil, nil, err
	}
	result := &ReportProperty{MessageType: "reportProperty", MessageID: fmt.Sprintf("%02x", header.Cmd), DeviceNo: decodePseudoIP(header.IP), Timestamp: location.DeviceTime, PropertiesMap: map[string]interface{}{}, AsyncSave: true, FlowType: "upstream"}
	result.PropertiesMap["location"] = map[string]interface{}{"lat": location.Lat, "lng": location.Lng}
	result.PropertiesMap["speed"] = location.Speed
	result.PropertiesMap["direction"] = location.Direction
	result.PropertiesMap["locationStatus"] = map[string]interface{}{"valid": location.HasValidLocation, "gpsAntenna": location.GPSAntenna, "power": location.PowerStatus, "raw": fmt.Sprintf("%02X", location.Status)}
	result.PropertiesMap["vehicleStatus"] = map[string]interface{}{"raw": fmt.Sprintf("%08X", location.VehicleStatus), "needAck": location.NeedAck, "transport": location.Transport, "signalStrength": location.SignalStrength}
	result.PropertiesMap["centerCommand"] = fmt.Sprintf("0x%02X", location.CenterCommand)
	result.PropertiesMap["deviceTime"] = location.DeviceTime
	result.PropertiesMap["deviceTimeText"] = timeText(location.DeviceTime)
	result.PropertiesMap["timeZone"] = "Asia/Shanghai"
	extension := body[locationBaseLength:]
	result.PropertiesMap["rawExtension"] = hex.EncodeToString(extension)
	tlvs, err := ParseTLV(extension)
	if err != nil {
		return nil, nil, err
	}
	for _, item := range tlvs {
		parseExtension(result, item)
	}
	return result, tlvs, nil
}

func parseExtension(result *ReportProperty, item TLV) {
	raw := hex.EncodeToString(item.Value)
	switch item.Type {
	case 0x0024:
		result.PropertiesMap["singleBaseStation"] = string(item.Value)
	case 0x0004:
		if len(item.Value) == 3 {
			result.PropertiesMap["voltage"] = float64(bcdInt(item.Value)) / 100000
		} else {
			result.PropertiesMap["voltageRaw"] = raw
		}
	case 0x0008:
		if len(item.Value) >= 2 {
			count := int(item.Value[0])<<8 | int(item.Value[1])
			result.PropertiesMap["battery"] = map[string]interface{}{"count": count, "total": 1500, "percent": float64(count) / 15}
		}
	case 0x00A3:
		result.PropertiesMap["softwareVersion"] = strings.TrimRight(string(item.Value), "\x00;")
	case 0x00A5:
		result.ProductNo = raw
		result.PropertiesMap["terminalModel"] = raw
	case 0x0089:
		result.PropertiesMap["alarmStatus"] = map[string]interface{}{"raw": raw, "moving": bit(item.Value, 9) == 0, "sos": bit(item.Value, 10) == 0, "removed": bit(item.Value, 12) == 0}
	case 0x00A9:
		result.PropertiesMap["baseStations"] = parseBaseStations(item.Value)
	case 0x00B9:
		result.PropertiesMap["wifi"] = parseWifi(item.Value)
	case 0x00C5:
		result.PropertiesMap["extendedLocationStatus"] = parseExtendedLocation(item.Value)
	case 0x00FB:
		result.PropertiesMap["iccid"] = strings.TrimRight(string(item.Value), "\x00")
	case 0x00AE:
		result.PropertiesMap["workMode"] = map[string]interface{}{"mode": first(item.Value), "raw": raw}
	case 0xF000:
		result.PropertiesMap["reportMode"] = map[string]interface{}{"mode": first(item.Value), "raw": raw}
	case 0xF001:
		result.PropertiesMap["nextReport"] = parseNextReport(item.Value)
	case 0x0030:
		result.PropertiesMap["signalStrength"] = first(item.Value)
	case 0x0031:
		result.PropertiesMap["satellites"] = first(item.Value)
	default:
		result.PropertiesMap[fmt.Sprintf("unknown_%04x", item.Type)] = raw
	}
}

func parseWifi(data []byte) map[string]interface{} {
	if len(data) == 0 {
		return map[string]interface{}{"count": 0, "hotspots": []string{}}
	}
	parts := strings.Split(string(data[1:]), ",")
	return map[string]interface{}{"count": int(data[0]), "hotspots": parts}
}
func parseBaseStations(data []byte) map[string]interface{} {
	result := map[string]interface{}{"raw": hex.EncodeToString(data)}
	if len(data) < 4 {
		return result
	}
	result["country"] = int(data[0])<<8 | int(data[1])
	result["operator"] = int(data[2])
	result["count"] = int(data[3])
	stations := []map[string]int{}
	for i := 4; i+5 <= len(data); i += 5 {
		stations = append(stations, map[string]int{"area": int(data[i])<<8 | int(data[i+1]), "tower": int(data[i+2])<<8 | int(data[i+3]), "signal": int(data[i+4])})
	}
	result["stations"] = stations
	return result
}
func parseExtendedLocation(data []byte) map[string]interface{} {
	value := uint32(0)
	for _, v := range data {
		value = value<<8 | uint32(v)
	}
	mode := "unknown"
	if value&0x08 != 0 {
		mode = "gps"
	} else if value&0x10 != 0 {
		mode = "wifi"
	} else {
		mode = "none"
	}
	return map[string]interface{}{"raw": fmt.Sprintf("%08X", value), "mode": mode}
}
func parseNextReport(data []byte) interface{} {
	if len(data) < 7 {
		return hex.EncodeToString(data)
	}
	timestamp, err := parse2929Time(data[:6])
	if err != nil {
		return hex.EncodeToString(data)
	}
	environment := map[byte]string{0: "全天空", 1: "半天空", 2: "地下室"}[data[6]]
	return map[string]interface{}{"timestamp": timestamp, "environment": environment}
}
func bit(data []byte, index uint) byte {
	value := uint64(0)
	for _, v := range data {
		value = value<<8 | uint64(v)
	}
	return byte(value >> index & 1)
}
func first(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	return int(data[0])
}

func decodePseudoIP(data []byte) string {
	if len(data) != 4 {
		return ""
	}
	marker := 0
	groups := make([]string, 4)
	for i, v := range data {
		if v&0x80 != 0 {
			marker |= 1 << (3 - i)
		}
		groups[i] = fmt.Sprintf("%02d", v&0x7F)
	}
	return fmt.Sprintf("1%02d%s", marker+30, strings.Join(groups, ""))
}
func parse2929Time(data []byte) (int64, error) {
	if len(data) != 6 {
		return 0, fmt.Errorf("BCD时间需要6字节, 实际%d字节", len(data))
	}
	for _, v := range data {
		if v>>4 > 9 || v&15 > 9 {
			return 0, fmt.Errorf("无效BCD值%02X", v)
		}
	}
	year := 2000 + bcd(data[0])
	month := time.Month(bcd(data[1]))
	day := bcd(data[2])
	hour := bcd(data[3])
	minute := bcd(data[4])
	second := bcd(data[5])
	value := time.Date(year, month, day, hour, minute, second, 0, beijingLocation)
	if value.Year() != year || value.Month() != month || value.Day() != day || value.Hour() != hour || value.Minute() != minute || value.Second() != second {
		return 0, fmt.Errorf("无效日期时间")
	}
	return value.UnixMilli(), nil
}
func bcd(v byte) int { return int(v>>4)*10 + int(v&15) }
