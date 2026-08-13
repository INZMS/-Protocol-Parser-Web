// 协议字段
package core

type Field struct {
	Index int `json:"index"`

	Name string `json:"name"`

	Offset int `json:"offset"`

	Length int `json:"length"`

	Raw string `json:"raw"`

	Value string `json:"value"`

	Description string `json:"description"`
}
