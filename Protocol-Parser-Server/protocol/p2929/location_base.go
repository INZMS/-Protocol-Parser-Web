package p2929

import "fmt"

const locationBaseLength = 34

type LocationData struct {
	DeviceTime       int64   `json:"deviceTime"`
	Lat              float64 `json:"lat"`
	Lng              float64 `json:"lng"`
	Speed            int     `json:"speed"`
	Direction        int     `json:"direction"`
	HasValidLocation bool    `json:"hasValidLocation"`
	GPSAntenna       string  `json:"gpsAntenna"`
	PowerStatus      string  `json:"powerStatus"`
	Status           byte    `json:"status"`
	VehicleStatus    uint32  `json:"vehicleStatus"`
	NeedAck          bool    `json:"needAck"`
	Transport        string  `json:"transport"`
	SignalStrength   int     `json:"signalStrength"`
	CenterCommand    byte    `json:"centerCommand"`
}

func ParseLocationBase(body []byte) (*LocationData, error) {
	if len(body) < locationBaseLength {
		return nil, fmt.Errorf("0x80位置数据长度不足: 至少需要%d字节, 实际%d字节", locationBaseLength, len(body))
	}
	t, err := parse2929Time(body[0:6])
	if err != nil {
		return nil, fmt.Errorf("位置时间无效: %w", err)
	}
	status := body[18]
	vehicle := uint32(body[22])<<24 | uint32(body[23])<<16 | uint32(body[24])<<8 | uint32(body[25])
	result := &LocationData{DeviceTime: t, Lat: parseCoordinate(body[6:10], false), Lng: parseCoordinate(body[10:14], true), Speed: bcdInt(body[14:16]), Direction: bcdInt(body[16:18]), Status: status, VehicleStatus: vehicle, CenterCommand: body[33]}
	result.HasValidLocation = status&0x80 != 0
	switch (status >> 5) & 3 {
	case 3:
		result.GPSAntenna = "正常"
	case 1:
		result.GPSAntenna = "开路"
	default:
		result.GPSAntenna = "未知"
	}
	switch (status >> 3) & 3 {
	case 3:
		result.PowerStatus = "正常"
	case 2:
		result.PowerStatus = "主电源掉电"
	case 1:
		result.PowerStatus = "主电源电压过低"
	default:
		result.PowerStatus = "未知"
	}
	communication := body[24]
	result.NeedAck = communication&0x40 != 0
	if communication&0x20 != 0 {
		result.Transport = "TCP"
	} else {
		result.Transport = "UDP"
	}
	result.SignalStrength = int(communication & 0x1F)
	return result, nil
}

func parseCoordinate(data []byte, longitude bool) float64 {
	if len(data) != 4 {
		return 0
	}
	value := append([]byte(nil), data...)
	negative := value[0]&0x80 != 0
	value[0] &= 0x7F
	digits := bcdString(value)
	degreeStart, minuteStart := 1, 3
	if longitude {
		degreeStart, minuteStart = 0, 3
	}
	coordinate := float64(toInt(digits[degreeStart:minuteStart])) + float64(toInt(digits[minuteStart:]))/60000
	if negative {
		return -coordinate
	}
	return coordinate
}
func toInt(s string) int {
	value := 0
	for _, c := range s {
		value = value*10 + int(c-'0')
	}
	return value
}
func bcdString(data []byte) string {
	result := make([]byte, 0, len(data)*2)
	for _, v := range data {
		result = append(result, '0'+(v>>4), '0'+(v&15))
	}
	return string(result)
}
func bcdInt(data []byte) int { return toInt(bcdString(data)) }
