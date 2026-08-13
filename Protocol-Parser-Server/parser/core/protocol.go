// 协议接口
package core

type Protocol interface {
	Name() string

	Match(data []byte) bool

	Parse(data []byte) (*ParseResult, error)
}
