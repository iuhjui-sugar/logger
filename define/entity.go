package define;

import (
    "time"

    "github.com/keystone-coin/logger/dot"
)


type IEntity interface{
    GetLevel()Level
    SetLevel(Level)
    GetTimestamp()uint64
    SetTimestamp(uint64)
    SetTemplate(string)
    GetTemplate()string
    Derive(map[string]string) IEntity
    String()string
}

type Entity struct {
    level     Level
    timestamp uint64
    template  string
    data      map[string]string
}

func NewEntity(level Level) *Entity {
    entity := &Entity{};
    entity.level = level;
    entity.timestamp = uint64(time.Now().Unix());
    return entity;
}

func (e *Entity) GetLevel() Level{
    return e.level;
}

func (e *Entity) SetLevel(l Level){
    e.level = l;
    return;
}

func (e *Entity) GetTimestamp() uint64 {
    return e.timestamp;
}

func (e *Entity) SetTimestamp(timestamp uint64) {
    e.timestamp = timestamp;
    return;
}

func (e *Entity) GetTemplate() string{
    return e.template;
}

func (e *Entity) SetTemplate(template string) {
    e.template = template;
    return;
}

func (e *Entity) Derive(data map[string]string) IEntity {
    derive := NewEntity(e.level);
    derive.SetTemplate(e.template);
    derive.data = data;
    return derive;
}

func (e *Entity) String() string  {
    str,_:= dot.New(e.template).Complie(e.data);
    return str;
}
