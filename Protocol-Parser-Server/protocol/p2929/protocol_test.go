package p2929

import (
	"encoding/hex"
	"strings"
	"testing"
)

func packet(cmd byte, body []byte) []byte {
	length := 4 + len(body) + trailerLength
	data := []byte{0x29, 0x29, cmd, byte(length >> 8), byte(length), 0x01, 0x02, 0x03, 0x04}
	data = append(data, body...)
	var checksum byte
	for _, value := range data {
		checksum ^= value
	}
	return append(data, checksum, 0x0D)
}

func TestParseCommonFields(t *testing.T) {
	result, err := New().Parse(packet(0xD8, []byte{0x01, 0x02}))
	if err != nil {
		t.Fatal(err)
	}
	if result.Length != 13 || len(result.Fields) != 6 {
		t.Fatalf("unexpected result: length=%d fields=%d", result.Length, len(result.Fields))
	}
	if result.Fields[0].Name != "包头" || result.Fields[4].Name != "校验码" {
		t.Fatalf("unexpected fields: %#v", result.Fields)
	}
}

func TestParseLocation(t *testing.T) {
	body := make([]byte, locationBaseLength)
	copy(body, []byte{0x26, 0x08, 0x13, 0x10, 0x20, 0x30, 0x22, 0x32, 0x55, 0x60, 0x11, 0x40, 0x52, 0x81, 0x00, 0x60, 0x01, 0x80, 0x80})
	result, err := New().Parse(packet(0x80, body))
	if err != nil {
		t.Fatal(err)
	}
	if result.MessageName != "一般位置数据" || len(result.Fields) != 16 {
		t.Fatalf("unexpected location result: %#v", result)
	}
	property := result.Data.(*ReportProperty)
	location := property.PropertiesMap["location"].(map[string]interface{})
	if location["lat"].(float64) < 22.54 || location["lng"].(float64) < 114.08 {
		t.Fatalf("unexpected coordinates: %#v", location)
	}
	if property.PropertiesMap["locationStatus"].(map[string]interface{})["valid"] != true {
		t.Fatal("expected valid location")
	}
}

func TestParseDocumentSample(t *testing.T) {
	const sample = "292980013B14941494210417185424022409531141642500000000FF000000FFF750FFFFFFFFFFFFFFFF00001200243436303B30303B393336393B343934320005000429680000040008045F000F00A359475F434430315F52322E373B000600A50000001000060089FFFFFFFF002400A901CC00032499134E1C249913121C249911EF1C000000000000000000000000000000007000B90532633A36313A30343A37393A64333A39632C2D36372C63633A30383A66623A39383A62373A39352C2D37322C30383A39623A34623A39643A39623A35312C2D37332C39633A32313A36613A65343A31353A65382C2D37382C39633A61363A31353A39663A64303A62302C2D3833000600C500001010001600FB3839383630343132313031383730383434363635000500AE0300050005F0000100050009F00121041718590700D50D"
	data, err := hex.DecodeString(sample)
	if err != nil {
		t.Fatal(err)
	}
	result, err := New().Parse(data)
	if err != nil {
		t.Fatal(err)
	}
	if result.Length != 320 || len(result.Fields) != 29 {
		t.Fatalf("length=%d fields=%d", result.Length, len(result.Fields))
	}
	property := result.Data.(*ReportProperty)
	if property.DeviceNo != "13520202020" {
		t.Fatalf("unexpected device number: %s", property.DeviceNo)
	}
	if property.PropertiesMap["softwareVersion"] != "YG_CD01_R2.7" {
		t.Fatalf("unexpected version: %#v", property.PropertiesMap["softwareVersion"])
	}
	if property.PropertiesMap["iccid"] != "89860412101870844665" {
		t.Fatalf("unexpected ICCID: %#v", property.PropertiesMap["iccid"])
	}
}

func TestParseRejectsDeclaredLengthMismatch(t *testing.T) {
	data := packet(0x21, nil)
	data[4]++
	var checksum byte
	for _, value := range data[:len(data)-2] {
		checksum ^= value
	}
	data[len(data)-2] = checksum
	_, err := New().Parse(data)
	if err == nil || !strings.Contains(err.Error(), "长度不一致") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestParseLocationRejectsShortBodyWithoutPanic(t *testing.T) {
	_, err := New().Parse(packet(0x80, make([]byte, 6)))
	if err == nil || !strings.Contains(err.Error(), "长度不足") {
		t.Fatalf("unexpected error: %v", err)
	}
}
