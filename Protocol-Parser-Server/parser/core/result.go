// 协议解析结果
package core

type ParseResult struct {
	Protocol string `json:"protocol"`

	MessageID string `json:"messageId"`

	MessageName string `json:"messageName"`

	Length int `json:"length"`

	Raw string `json:"raw"`

	//原始字段解析
	Fields []Field `json:"fields"`

	// 业务解析结果
	Data interface{} `json:"data"`
}
