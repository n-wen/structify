package main

import "strings"

func BigCamelMarshal(name string) string {
	if name == "" {
		return ""
	}

	temp := strings.Split(name, "_")
	var s string
	for _, v := range temp {
		vv := []rune(v)
		if len(vv) > 0 {
			if bool(vv[0] >= 'a' && vv[0] <= 'z') { //首字母大写
				vv[0] -= 32
			}
			s += string(vv)
		}
	}

	// s = uncommonInitialismsReplacer.Replace(s)
	//smap.Set(name, s)
	return s
}


type PrintAtom struct {
	lines []string
}

var _interval = "\t"

func (p *PrintAtom) Add(str ...string) {
	var tmp string
	for _, v := range str {
		tmp += v + _interval
	}
	p.lines = append(p.lines, tmp)
}

// Generates Get the generated list.获取生成列表
func (p *PrintAtom) Generates() []string {
	return p.lines
}
