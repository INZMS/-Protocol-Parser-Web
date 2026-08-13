// 校验
package p2929

import "fmt"

func CheckChecksum(data []byte) error {

	if len(data) < 3 {

		return fmt.Errorf("数据长度不足")

	}

	// 最后一位是包尾0D
	// 倒数第二位是校验

	checkIndex := len(data) - 2

	expect := data[checkIndex]

	var checksum byte

	// 从包头29 29开始异或
	for i := 0; i < checkIndex; i++ {

		checksum ^= data[i]

	}

	if checksum != expect {

		return fmt.Errorf(
			"校验失败,计算:%02X 实际:%02X",
			checksum,
			expect,
		)

	}

	return nil

}
