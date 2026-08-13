// CMD定义

package p2929

func MessageName(cmd byte) string {

	switch cmd {

	case 0x21:

		return "中心确认"

	case 0x3A:

		return "调度文本"

	case 0x80:

		return "一般位置数据"

	case 0x85:

		return "终端确认"

	case 0xD7:

		return "校时"

	case 0xD8:

		return "申请设置参数"

	default:

		return "未知指令"

	}

}
