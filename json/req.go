package json

import (
	"fmt"
)

type Req struct {
	Config   map[string]interface{} `json:"config"`
	Header   Header                 `json:"header"`
	Elements []*Element             `json:"elements"`
}
type Header struct {
	Template string  `json:"template"`
	Title    Content `json:"title"`
}
type Element struct {
	Tag  string  `json:"tag"`
	Text Content `json:"text"`
}
type Content struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

func Build() *Req {
	req := &Req{}
	req.Config = make(map[string]interface{}, 0)
	req.Config["wide_screen_mode"] = true
	req.Header.Template = "red"

	titleContent := Content{}
	titleContent.Tag = "plain_text"
	titleContent.Content = "高热音频片段状态异常"

	bodyContent := Content{}
	bodyContent.Tag = "lark_md"
	bodyContent.Content = "测试"

	operatorContent := Content{}
	operatorContent.Content = "关注人"
	operatorContent.Tag = "lark_md"

	element1 := Element{Text: bodyContent, Tag: "div"}
	element2 := Element{Text: operatorContent, Tag: "div"}

	req.Header.Title = titleContent
	req.Elements = make([]*Element, 0)
	req.Elements = append(req.Elements, &element1)
	req.Elements = append(req.Elements, &element2)
	str1 := jsonx.ToString(req)
	fmt.Printf(str1)
	return req
}
