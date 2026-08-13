package p2929

type TLV struct {
	Type uint16

	Length uint16

	Value []byte
}

// 解析TLV
func ParseTLV(data []byte) []TLV {

	var list []TLV

	offset := 0

	for offset+4 <= len(data) {

		t := uint16(data[offset])<<8 |
			uint16(data[offset+1])

		l := uint16(data[offset+2])<<8 |
			uint16(data[offset+3])

		offset += 4

		if offset+int(l) > len(data) {

			break

		}

		item := TLV{

			Type: t,

			Length: l,

			Value: data[offset : offset+int(l)],
		}

		list = append(
			list,
			item,
		)

		offset += int(l)

	}

	return list

}
