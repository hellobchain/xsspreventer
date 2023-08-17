package xsspreventer

import "html"

var _ XSSPreventer = (*PreventXSSStruct)(nil)

// XSS防止接口
type XSSPreventer interface {
	PreventXSS(input string) string
}

// 实现XSS防止接口的结构体
type PreventXSSStruct struct{}

func NewPreventXSSStruct() XSSPreventer {
	return &PreventXSSStruct{}
}

// 实现PreventXSS方法
func (p *PreventXSSStruct) PreventXSS(input string) string {
	return html.EscapeString(input)
}
