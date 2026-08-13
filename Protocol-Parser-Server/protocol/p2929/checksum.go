package p2929

import "fmt"

func CheckChecksum(data []byte) error {
	if len(data) < headerLength+trailerLength {
		return fmt.Errorf("数据长度不足")
	}
	if data[len(data)-1] != 0x0D {
		return fmt.Errorf("2929报文缺少结束符0D")
	}
	checkIndex := len(data) - 2
	expect := data[checkIndex]
	var checksum byte
	for i := 0; i < checkIndex; i++ {
		checksum ^= data[i]
	}
	if checksum != expect {
		return fmt.Errorf("校验失败,计算:%02X 实际:%02X", checksum, expect)
	}
	return nil
}
