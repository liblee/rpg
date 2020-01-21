//package buff

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type BuffInfo struct {
	Id         int    `json:"id"`
	Type       int    `json:"type"`
	Name       string `json:"name"`
	Count      int    `json:"count"`
	Demage     int    `json:"demage"`
	Luck       int    `json:"luck"`
	Last       int    `json:"last"`
	LastDemage int    `json:"last_demage"`
	Status     int    `json:"status"`
	CoolDown   int    `json:"cool_down"`
	Details    string `json:"details"`
}

type Buff struct {
	Buffs []BuffInfo `json:"infos"`
}

func (c *Buff) ShowInfo() {
	for _, i := range c.Buffs {
		info := fmt.Sprintf("名称:%s|伤害:%d|持续:%d|持续伤害:%d|冷却时间:%d|描述:%s\n",
			i.Name, i.Demage, i.Last, i.LastDemage, i.CoolDown, i.Details)
		fmt.Printf(info)
	}
}

func (c *Buff) GetItem(id int) *BuffInfo {
	for i, info := range c.Buffs {
		if id == info.Id {
			return &c.Buffs[i]
		}
	}
	return nil
}

func (c *Buff) LoadFromFile(filepath string) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return
	}
	if err := json.Unmarshal(bytes, c); err != nil {
		fmt.Println("Unmarshal json failed.", err)
		return
	}
}

func NewBuff() *Buff {
	c := new(Buff)
	c.LoadFromFile("skill.json")
	return c
}

var BuffMgr *Buff = NewBuff()

func main() {
	it := NewBuff()
	it.ShowInfo()
}
