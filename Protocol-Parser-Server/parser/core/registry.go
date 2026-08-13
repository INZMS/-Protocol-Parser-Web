// 协议注册表
package core

var protocols []Protocol

func Register(p Protocol) {

	protocols =
		append(protocols, p)

}

func Detect(data []byte) (Protocol, error) {

	for _, p := range protocols {

		if p.Match(data) {

			return p, nil

		}

	}

	return nil, ErrUnknownProtocol

}
