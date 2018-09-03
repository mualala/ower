package spider

import (
	"fmt"

	"github.com/henrylee2cn/pholcus/common/pinyin"
)

// 蜘蛛种类列表
type SpiderSpecies struct {
	spiders []*Spider
	hash    map[string]*Spider
	sorted  bool
}

// 全局蜘蛛种类实例
var Species = &SpiderSpecies{
	spiders: []*Spider{},
	hash:    map[string]*Spider{},
}

// 向蜘蛛种类清单添加新种类
func (ss *SpiderSpecies) Add(sp *Spider) *Spider {
	name := sp.Name
	for i := 2; true; i++ {
		if _, ok := ss.hash[name]; !ok {
			sp.Name = name
			ss.hash[sp.Name] = sp
			break
		}
		name = fmt.Sprintf("%s(%d)", sp.Name, i)
	}
	sp.Name = name
	ss.spiders = append(ss.spiders, sp)
	return sp
}

// 获取全部蜘蛛种类
func (ss *SpiderSpecies) Get() []*Spider {
	if !ss.sorted {
		l := len(ss.spiders)
		initials := make([]string, l)
		newlist := map[string]*Spider{}
		for i := 0; i < l; i++ {
			initials[i] = ss.spiders[i].Name
			newlist[initials[i]] = ss.spiders[i]
		}
		pinyin.SortInitials(initials)
		for i := 0; i < l; i++ {
			ss.spiders[i] = newlist[initials[i]]
		}
		ss.sorted = true
	}
	return ss.spiders
}

func (ss *SpiderSpecies) GetByName(name string) *Spider {
	return ss.hash[name]
}
