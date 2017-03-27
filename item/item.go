package item
//package main

import (
	"fmt"
	"io/ioutil"
"encoding/json"
)

type ItemInfo struct {
	Id      int    `json:"id"`
	Type    int    `json:"type"`
	Name    string `json:"name"`
	Cost    int    `json:"cost"`
	Demage  int    `json:"demage"`
	Luck    int    `json:"luck"`
	Life    int    `json:"life"`
	Skillid int    `json:"skillid"`
	Money   int    `json:"money"`
	Status  string `json:"status"`
	Details string `json:"details"`
}

type Item struct {
	Infos []ItemInfo `json:"infos"`
}

func (c *Item) ShowInfo() {
	for _, i := range c.Infos{
		info := fmt.Sprintf("Name:%s|Demage:%d|Value:%d|State:%s\n",
		i.Name, i.Demage, i.Money, i.Status)
		fmt.Printf(info)
	}
}

func (c *Item) GetItem(id int)ItemInfo{
	for _, info :=range c.Infos{
		if id == info.Id {
			return info
		}
	}
}

func (c *Item) LoadFromFile(filepath string) {
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

func NewItem() *Item {
	c := new(Item)
	c.LoadFromFile("item.json")
	return c
}

func main(){
	it := NewItem()
	it.ShowInfo()
}

