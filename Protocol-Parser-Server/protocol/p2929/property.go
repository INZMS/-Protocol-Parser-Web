package p2929

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type ReportProperty struct {
	MessageType string `json:"messageType"`

	MessageID string `json:"messageId"`

	DeviceNo string `json:"deviceNo"`

	ProductNo string `json:"productNo"`

	Timestamp int64 `json:"timestamp"`

	ReceiptTimestamp int64 `json:"receiptTimestamp"`

	PropertiesMap map[string]interface{} `json:"propertiesMap"`

	AsyncSave bool `json:"asyncSave"`

	FlowType string `json:"flowType"`
}

// 0x80 一般位置数据解析
func ParseReportProperty(
	header *Header,
	data []byte,
) (
	*ReportProperty,
	error,
) {

	result := &ReportProperty{

		MessageType: "reportProperty",

		PropertiesMap: make(
			map[string]interface{},
		),

		AsyncSave: true,

		FlowType: "upstream",
	}

	/*

		2929:

		29 29
		CMD
		LEN LEN
		IP IP IP IP

		body开始

	*/

	if len(data) < 15 {

		return result, nil

	}

	body := data[9:]

	//解析34字节位置

	location, _ :=
		ParseLocationBase(body)

	result.Timestamp =
		location.DeviceTime

	result.PropertiesMap["location"] =
		map[string]interface{}{

			"lat": location.Lat,

			"lng": location.Lng,
		}

	result.PropertiesMap["speed"] =
		location.Speed

	result.PropertiesMap["direction"] =
		location.Direction

	result.PropertiesMap["hasValidLocation"] =
		location.HasValidLocation

	// 跳过34字节

	if len(body) > 34 {

		body =
			body[34:]

	}

	// ==========================
	// 时间
	// ==========================

	timestamp, err := parse2929Time(
		body[:6],
	)

	if err == nil {

		result.Timestamp =
			timestamp

		result.PropertiesMap["deviceTime"] = timestamp

	}

	// 原始数据保存

	// result.PropertiesMap["rawBody"] = hex.EncodeToString(body)
	result.PropertiesMap["rawExtension"] = hex.EncodeToString(body)

	// ==========================
	// TLV解析
	// ==========================

	if len(body) <= 6 {

		return result, nil

	}

	tlvBody := body[6:]

	tlvs := ParseTLV(
		tlvBody,
	)

	for _, tlv := range tlvs {

		switch tlv.Type {

		/*

			基站信息

			0018 0024
			460;00;26430;215717960

		*/

		case 0x0018:

			result.PropertiesMap["singleBaseStation"] = string(tlv.Value)

		/*

			WIFI

			0006 xxxx

		*/

		case 0x0006:

			result.PropertiesMap["wifi"] = parseWifi(
				tlv.Value,
			)

		/*

			产品信息

			T96_2929_V5.0

		*/

		case 0x0005:

			parseProduct(
				result,
				tlv.Value,
			)

		/*

			设备参数

			电池
			工作模式
			定位状态

		*/

		case 0x0004:

			parseDeviceParam(
				result,
				tlv.Value,
			)

		default:

			result.PropertiesMap[fmt.Sprintf(
				"unknown_%04x",
				tlv.Type,
			)] = hex.EncodeToString(
				tlv.Value,
			)

		}

	}

	return result, nil

}

// ==============================
// WIFI解析
// ==============================

func parseWifi(
	data []byte,
) map[string]interface{} {

	text := string(data)

	return map[string]interface{}{

		"count": 1,

		"hotspots": []string{

			text,
		},
	}

}

// ==============================
// 产品型号解析
// ==============================

func parseProduct(
	result *ReportProperty,
	data []byte,
) {

	text := string(data)

	parts := strings.Split(
		text,
		"_",
	)

	if len(parts) > 0 {

		result.ProductNo =
			parts[0]

	}

}

// ==============================
// 设备参数解析
// 后续继续扩展
// ==============================

func parseDeviceParam(
	result *ReportProperty,
	data []byte,
) {

	if len(data) < 2 {

		return

	}

	value :=
		int(data[0])<<8 |
			int(data[1])

	result.PropertiesMap["deviceParamRaw"] = hex.EncodeToString(data)

	// 暂时按照电压处理

	result.PropertiesMap["deviceBatteryVoltage"] = float64(value) / 1000

}

// ==============================
// 2929 BCD时间
//
// YY MM DD HH MM SS
// ==============================

func parse2929Time(
	data []byte,
) (
	int64, error,
) {

	if len(data) < 6 {

		return 0, nil

	}

	year := 2000 +
		int(bcd(data[0]))

	month := time.Month(
		bcd(data[1]),
	)

	day := bcd(data[2])

	hour := bcd(data[3])

	minute := bcd(data[4])

	second := bcd(data[5])

	t := time.Date(

		year,

		month,

		day,

		hour,

		minute,

		second,

		0,

		time.UTC,
	)

	return t.UnixMilli(), nil

}

// ==============================
// BCD转换
// ==============================

func bcd(
	v byte,
) int {

	return int(v>>4)*10 +
		int(v&0x0f)

}
