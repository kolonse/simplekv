package simplekv

import (
	"io/ioutil"
	"os"
)

type SKV struct {
	Dir string // 存储根路径
}

func (p *SKV) init() error {
	if p.Dir[len(p.Dir)-1] == '/' || p.Dir[len(p.Dir)-1] == '\\' {
		p.Dir = p.Dir[0 : len(p.Dir)-1]
	}
	err := os.MkdirAll(p.Dir, 666)
	return err
}

func (p *SKV) Read(key string) Value {
	filename := p.getFilename(key)

	f, err := os.Open(filename)
	if err != nil {
		return Value{}
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return Value{}
	}
	v := NewValue(buf)
	return v
}

func (p *SKV) Write(key string, value []byte) error {
	return ioutil.WriteFile(p.getFilename(key), value, 666)
}

func (p *SKV) Exist(key string) bool {
	_, err := os.Stat(p.getFilename(key))
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}

	return false
}

func (p *SKV) getFilename(key string) string {
	return p.Dir + "/" + key
}

func NewSKV(name string, dir string) (*SKV, error) {
	if len(dir) == 0 {
		dir = "/tmp/simplekv"
	}
	if dir[len(dir)-1] == '/' || dir[len(dir)-1] == '\\' {
		dir = dir[0 : len(dir)-1]
	}
	p := &SKV{
		Dir: dir + "/" + name,
	}
	err := p.init()
	if err != nil {
		return nil, err
	}
	return p, nil
}
