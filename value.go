package simplekv

import (
	"encoding/json"
)

type Value struct {
	V []byte
}

func (v Value) ToString() string {
	return string(v.V)
}

func (v Value) ToJsonArray() []interface{} {
	r := make([]interface{}, 1)
	err := json.Unmarshal(v.V, &r)
	if err != nil {
		return nil
	}
	return r
}

func (v Value) ToJsonObject() map[string]interface{} {
	r := make(map[string]interface{})
	err := json.Unmarshal(v.V, &r)
	if err != nil {
		return nil
	}
	return r
}

func (v Value) ToJson() interface{} {
	var r interface{}
	err := json.Unmarshal(v.V, &r)
	if err != nil {
		return nil
	}
	return r
}

func NewValue(buf []byte) Value {
	return Value{
		V: buf,
	}
}
