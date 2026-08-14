// 协议注册表
package core

import "strings"

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

func Get(name string) (Protocol, error) {
	for _, p := range protocols {
		if strings.EqualFold(p.Name(), name) {
			return p, nil
		}
	}
	return nil, ErrUnknownProtocol
}
