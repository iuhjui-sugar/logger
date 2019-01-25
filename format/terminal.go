package format;

import (
    "fmt"
    "time"
    "strings"

    "github.com/keystone-coin/logger/define"
)

type TerminalFormat struct {}

func NewTerminalFormat() *TerminalFormat  {
    f := &TerminalFormat{}
    return f;
}

func (t *TerminalFormat) Format(ie define.IEntity) []byte  {
    level := strings.ToUpper(ie.GetLevel().String());
    stime := time.Unix(int64(ie.GetTimestamp()), 0).Format("2006-01-02 15:04:05");
    smsg  := ie.String();
    s := fmt.Sprintf("[%s] [%s] %s \n",stime,level,smsg);
    return []byte(s);
}
