package p2929

type LocationData struct {
	DeviceTime int64 `json:"deviceTime"`

	Lat float64 `json:"lat"`

	Lng float64 `json:"lng"`

	Speed string `json:"speed"`

	Direction string `json:"direction"`

	HasValidLocation bool `json:"hasValidLocation"`

	Status byte `json:"status"`
}

func ParseLocationBase(
	body []byte,
) (*LocationData, error) {

	result := &LocationData{}

	if len(body) < 34 {

		return result, nil

	}

	// 时间 YYMMDDHHMMSS

	t, _ := parse2929Time(
		body[0:6],
	)

	result.DeviceTime = t

	// 纬度

	result.Lat =
		parseCoordinate(
			body[6:10],
			false,
		)

	// 经度

	result.Lng =
		parseCoordinate(
			body[10:14],
			true,
		)

	//速度

	result.Speed =
		bcdString(
			body[14:16],
		)

	//方向

	result.Direction =
		bcdString(
			body[16:18],
		)

	//状态

	result.Status =
		body[18]

	// D7定位标志

	if result.Status&0x80 != 0 {

		result.HasValidLocation = true

	}

	return result, nil

}

// 坐标解析
//
// 纬度:
// DDMMmmm
//
// 经度:
// DDDMMmmm
//

func parseCoordinate(
	data []byte,
	isLng bool,
) float64 {

	if len(data) != 4 {

		return 0

	}

	v := bcdString(data)

	if isLng {

		// 经度

		// 11405281

		if len(v) < 8 {

			return 0

		}

		deg :=

			toInt(v[0:3])

		min :=

			toInt(v[3:8])

		return float64(deg) +
			float64(min)/60000

	} else {

		// 纬度

		// 02232556

		if len(v) < 8 {

			return 0

		}

		deg :=

			toInt(v[0:2])

		min :=

			toInt(v[2:7])

		return float64(deg) +
			float64(min)/60000

	}

}

func toInt(
	s string,
) int {

	r := 0

	for _, c := range s {

		r = r*10 +
			int(c-'0')

	}

	return r

}

func bcdString(
	data []byte,
) string {

	buf := make(
		[]byte,
		0,
		len(data)*2,
	)

	for _, v := range data {

		buf = append(
			buf,
			'0'+(v>>4),
			'0'+(v&0x0f),
		)

	}

	return string(buf)

}
