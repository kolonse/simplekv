package simplekv

import (
	"testing"
)

func TestKV(t *testing.T) {
	kv, _ := NewSKV("test", "/test2")
	kv.Write("test1", []byte("test1"))
	str := kv.Read("test1").ToString()
	if str != "test1" {
		t.Error("读取值与写入值不一致 read:", str)
	}

	kv.Write("test2", []byte("{\"key\":1}"))
	obj := kv.Read("test2").ToJsonObject()

	if obj["key"] == nil || obj["key"].(float64) != 1 {
		t.Error("读取值与写入值不一致 read:", obj["key"].(float64))
	}

	obj2 := kv.Read("test2").ToJson().(map[string]interface{})
	if obj2["key"] == nil || obj2["key"].(float64) != 1 {
		t.Error("读取值与写入值不一致 read:", obj2["key"].(float64))
	}

}

func TestKVExist(t *testing.T) {
	kv, err := NewSKV("test", "/test2")
	if err != nil {
		t.Error(err.Error())
	}
	if !kv.Exist("test2") {
		t.Error("文件存在却判断不存在")
	}
}
