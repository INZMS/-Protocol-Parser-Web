package p2929

import "fmt"

const locationBaseLength = 34

type LocationData struct {
	DeviceTime       int64   `json:"deviceTime"`
	Lat              float64 `json:"lat"`
	Lng              float64 `json:"lng"`
	Speed            string  `json:"speed"`
	Direction        string  `json:"direction"`
	HasValidLocation bool    `json:"hasValidLocation"`
	Status           byte    `json:"status"`
}

func ParseLocationBase(body []byte) (*LocationData, error) {
	if len(body) < locationBaseLength {
		return nil, fmt.Errorf("0x80位置数据长度不足: 至少需要%d字节, 实际%d字节", locationBaseLength, len(body))
	}
	t, err := parse2929Time(body[0:6])
	if err != nil {
		return nil, fmt.Errorf("位置时间无效: %w", err)
	}
	result := &LocationData{
		DeviceTime: t,
		Lat:        parseCoordinate(body[6:10], false),
		Lng:        parseCoordinate(body[10:14], true),
		Speed:      bcdString(body[14:16]),
		Direction:  bcdString(body[16:18]),
		Status:     body[18],
	}
	result.HasValidLocation = result.Status&0x80 != 0
	return result, nil
}

func parseCoordinate(data []byte, isLng bool) float64 {
	if len(data) != 4 {
		return 0
	}
	v := bcdString(data)
	if isLng {
		return float64(toInt(v[0:3])) + float64(toInt(v[3:8]))/60000
	}
	return float64(toInt(v[0:2])) + float64(toInt(v[2:7]))/60000
}

func toInt(s string) int {
	r := 0
	for _, c := range s {
		r = r*10 + int(c-'0')
	}
	return r
}

func bcdString(data []byte) string {
	buf := make([]byte, 0, len(data)*2)
	for _, v := range data {
		buf = append(buf, '0'+(v>>4), '0'+(v&0x0f))
	}
	return string(buf)
}
