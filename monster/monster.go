package monster

//package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//	. "../item"
)

type Monster struct {
	Name      string `json:"name"`
	TotalLife int    `json:"totallife"`
	CurLife   int    `json:"curlife"`
	Demage    int    `json:"demage"`
	ExDemage  int    `json:"exdemage"`
	Luck      int    `json:"luck"`
	Intval    int    `json:"Intval"`
	Mana      int    `json:"mana"`
	Money     int    `json:"money"`
	Items     []int  `json:"items"`
}

func (h *Monster) ShowAllItems() {
	for _, c := range h.Items {
		fmt.Printf("%v\n", c)
	}
}

func (cl *Monster) LoadFromFile(filepath string) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return
	}
	if err := json.Unmarshal(bytes, cl); err != nil {
		fmt.Println("Unmarshal json failed.", err)
		return
	}
}

func NewMonster() *Monster {
	c := new(Monster)
	c.Items = make([]int, 30)
	return c
}

func main() {
	c := NewMonster()
	c.LoadFromFile("../data/monster.json")
	c.ShowAllItems()
}
