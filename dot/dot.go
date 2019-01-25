package dot;

import (
	"bytes"
	"errors"
	"strings"
)

type DoT struct {
	template string
	start    string
	end      string
}

func New(template string) *DoT {
    dot := new(DoT);
    dot.template = template;
    dot.start    = "{{"
    dot.end      = "}}"
	return dot;
}

func (dot *DoT) Complie(params map[string]string) (string,error) {
	buffer   := bytes.NewBuffer(nil);
	template := []byte(dot.template);
	start    := []byte(dot.start);
	end      := []byte(dot.end);
	for {
		// 处理开头标签
		n := bytes.Index(template,start)
		if n < 0 {
			buffer.Write(template);
			break;
		}
		_, err := buffer.Write(template[:n])
		if err != nil {
			return "",err;
		}
		// 处理结尾标签
		template = template[n+len(start):]
		n = bytes.Index(template,end)
		// 没有找到结尾标签，将开头标签写入并退出
		if n < 0 {
			return "",errors.New("doT end tag no exists");
		}
		// 替换对应参数
		name := strings.TrimSpace(string(template[:n]))
		value,exists := params[name];
		if exists == true {
			buffer.Write([]byte(value));
		}
		// 切除结尾标签
		template = template[n+len(end):];
	}
	return string(buffer.Bytes()),nil;
}
