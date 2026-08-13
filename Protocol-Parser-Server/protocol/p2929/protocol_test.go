package p2929

import (
	"strings"
	"testing"
)

func packet(cmd byte, body []byte) []byte {
	length := headerLength + len(body) + trailerLength
	data := []byte{0x29, 0x29, cmd, byte(length >> 8), byte(length), 0x01, 0x02, 0x03, 0x04}
	data = append(data, body...)
	var checksum byte
	for _, value := range data {
		checksum ^= value
	}
	return append(data, checksum, 0x0D)
}

func TestParseCommonFields(t *testing.T) {
	result, err := New().Parse(packet(0x21, []byte{0x01, 0x02}))
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
	if result.MessageName != "一般位置数据" || len(result.Fields) != 13 {
		t.Fatalf("unexpected location result: %#v", result)
	}
	property := result.Data.(*ReportProperty)
	location := property.PropertiesMap["location"].(map[string]interface{})
	if location["lat"].(float64) < 22.54 || location["lng"].(float64) < 114.08 {
		t.Fatalf("unexpected coordinates: %#v", location)
	}
	if property.PropertiesMap["hasValidLocation"] != true {
		t.Fatal("expected valid location")
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
