// 包头解析
package p2929

import (
	"fmt"
)

// 2929固定头

type Header struct {

	//主信令

	Cmd byte

	//数据长度

	Length uint16

	//伪IP

	IP []byte
}

//解析2929包头

func ParseHeader(data []byte) (*Header, error) {

	//最小长度

	//
	//29 29 CMD LEN LEN IP IP IP IP
	//

	if len(data) < 9 {

		return nil,
			fmt.Errorf("2929数据长度不足")

	}

	//判断包头

	if data[0] != 0x29 ||
		data[1] != 0x29 {

		return nil,
			fmt.Errorf("不是2929协议")

	}

	header := &Header{

		//CMD
		Cmd: data[2],
		Length: uint16(data[3])<<8 |
			uint16(data[4]),

		//伪IP
		IP: data[5:9],
	}

	// header.Cmd = data[2]

	// //长度

	// header.Length =
	// 	uint16(data[3])<<8 |
	// 		uint16(data[4])

	// //伪IP

	// header.IP =
	// 	data[5:9]

	return header, nil

}
