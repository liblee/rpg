//package skill

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type SkillInfo struct {
	Id         int    `json:"id"`
	Type       int    `json:"type"`
	Name       string `json:"name"`
	Cost       int    `json:"cost"`
	Demage     int    `json:"demage"`
	Luck       int    `json:"luck"`
	Last       int    `json:"last"`
	LastDemage int    `json:"last_demage"`
	Status     int    `json:"status"`
	CoolDown   int    `json:"cool_down"`
	Details    string `json:"details"`
}

type Skill struct {
	Skills []SkillInfo `json:"infos"`
}

func (c *Skill) ShowInfo() {
	for _, i := range c.Skills {
		info := fmt.Sprintf("名称:%s|伤害:%d|持续:%d|持续伤害:%d|冷却时间:%d|描述:%s\n",
			i.Name, i.Demage, i.Last, i.LastDemage, i.CoolDown, i.Details)
		fmt.Printf(info)
	}
}

func (c *Skill) GetItem(id int) *SkillInfo {
	for i, info := range c.Skills {
		if id == info.Id {
			return &c.Skills[i]
		}
	}
	return nil
}

func (c *Skill) LoadFromFile(filepath string) {
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

func NewSkill() *Skill {
	c := new(Skill)
	c.LoadFromFile("skill.json")
	return c
}

var SkillMgr *Skill = NewSkill()

func main() {
	it := NewSkill()
	it.ShowInfo()
}
