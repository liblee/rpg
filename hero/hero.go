package hero

//package main

import (
	. "../item"
	//	. "../skill"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Equipment struct {
	Weapon   int `json:"weapon"`
	Armor    int `json:"armor"`
	Shoulder int `json:"shoulder"`
	Pants    int `json:"pants"`
	Shoes    int `json:"shoes"`
}

type Hero struct {
	Name       string    `json:"name"`
	BaseLife   int       `json:"baselife"`
	TotalLife  int       `json:"totallife"`
	CurLife    int       `json:"curlife"`
	BaseDemage int       `json:"basedemage"`
	Demage     int       `json:"demage"`
	ExDemage   int       `json:"exdemage"`
	Luck       int       `json:"luck"`
	Intval     int       `json:"intavl"`
	Mana       int       `json:"mana"`
	Level      int       `json:"level"`
	NextExp    int       `json:"nextexp"`
	CurExp     int       `json:"curexp"`
	Money      int       `json:"money"`
	Eqpmt      Equipment `json:"equipment"`
	Items      []int     `json:"items"`
	Skills     []int     `json:"skills"`
}

func (h *Hero) ReBuild() {
	h.Demage = h.BaseDemage + ItemMgr.GetItem(h.Eqpmt.Weapon).Demage
	h.TotalLife = h.BaseLife + ItemMgr.GetItem(h.Eqpmt.Armor).Life
}

const (
	WEAPON = iota
	ARMOR
	SHOULDER
	PANTS
	SHOES
)

func (h *Hero) EquipItem(id int) {
	i := id
	if i < 0 || i >= len(h.Items) {
		return
	}
	ItemMgr.GetItem(h.Items[i]).Status = "equiped"
	c := h.Items[i]
	h.Items = append(h.Items[:i], h.Items[i+1:]...)
	switch ItemMgr.GetItem(c).Type {
	case WEAPON:
		ItemMgr.GetItem(h.Eqpmt.Weapon).Status = "not used"
		h.Items = append(h.Items, h.Eqpmt.Weapon)
		h.Eqpmt.Weapon = c
	case ARMOR:
		ItemMgr.GetItem(h.Eqpmt.Armor).Status = "not used"
		h.Items = append(h.Items, h.Eqpmt.Armor)
		h.Eqpmt.Armor = c
	case SHOULDER:
		ItemMgr.GetItem(h.Eqpmt.Shoulder).Status = "not used"
		h.Items = append(h.Items, h.Eqpmt.Shoulder)
		h.Eqpmt.Shoulder = c
	case PANTS:
		ItemMgr.GetItem(h.Eqpmt.Pants).Status = "not used"
		h.Items = append(h.Items, h.Eqpmt.Pants)
		h.Eqpmt.Pants = c
	case SHOES:
		ItemMgr.GetItem(h.Eqpmt.Shoes).Status = "not used"
		h.Items = append(h.Items, h.Eqpmt.Shoes)
		h.Eqpmt.Shoes = c
	default:
		fmt.Printf("Equipment type error:%d.\n", ItemMgr.GetItem(c).Type)
	}
	h.ReBuild()
}

func (h *Hero) SellItem(id int) {
	i := id
	if i < 0 || i >= len(h.Items) {
		return
	}
	c := ItemMgr.GetItem(h.Items[i])
	h.Items = append(h.Items[:i], h.Items[i+1:]...)
	h.Money = h.Money + c.Money
	fmt.Printf("Sell %s change for %d $.\n", c.Name, c.Money)
}

func (h *Hero) UseItem(id int) {
	i := id
	if i < 0 || i >= len(h.Items) {
		return
	}
	c := ItemMgr.GetItem(h.Items[i])
	if c.Type != 6 {
		fmt.Printf("%s can not be used.\n", c.Name)
		return
	}
	h.Items = append(h.Items[:i], h.Items[i+1:]...)
	h.CurLife = h.CurLife + c.Life
	if h.CurLife > h.TotalLife {
		h.CurLife = h.TotalLife
	}
	fmt.Printf("Use %s heal %d hp, current life:%d.\n",
		c.Name, c.Life, h.CurLife)
	return
}

func (h *Hero) ShowEquipItems() {
	fmt.Printf("Weapon:%10s\nArmor:%10s\nShoulder:%10s\nPants:%10s\nShoes:%10s\n",
		ItemMgr.GetItem(h.Eqpmt.Weapon).Name, ItemMgr.GetItem(h.Eqpmt.Armor).Name, ItemMgr.GetItem(h.Eqpmt.Shoulder).Name,
		ItemMgr.GetItem(h.Eqpmt.Pants).Name, ItemMgr.GetItem(h.Eqpmt.Shoes).Name)

}

func (h *Hero) ShowAllItems() {
	for i, c := range h.Items {
		fmt.Printf("Id:%d,%+v\n", i, ItemMgr.GetItem(c))
	}
}

func (h *Hero) ShowStatus() {
	fmt.Printf("Name:\t%s\nTotalLife:\t%d\nCurLife:\t%d\nDemage:\t%d\nMana:\t%d\nLevel:\t%d\nMoney:\t%d\n",
		h.Name, h.TotalLife, h.CurLife, h.Demage, h.Mana, h.Level, h.Money)
}

func (h *Hero) ReapItem(it int) {
	h.Items = append(h.Items, it)
	fmt.Printf("%s reap %s.\n", h.Name, ItemMgr.GetItem(it).Name)
}

func (cl *Hero) LoadFromFile(filepath string) {
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

func (h *Hero) SaveToFile(filepath string) {
	w, err := json.Marshal(h)
	if err != nil {
		fmt.Println("marshal json failed.", err)
		return
	}
	var d1 = []byte(w)
	err = ioutil.WriteFile(filepath, d1, 0666)
	if err != nil {
		fmt.Println("write to file failed.")
		return
	}
	fmt.Println("write to file ok.")

}

func (h *Hero) run() {
	for {
		select {
		case <-time.After(time.Second * 5):
			h.CurLife = h.CurLife + 10
			if h.CurLife > h.TotalLife {
				h.CurLife = h.TotalLife
			}
			//println("5s timer")

		case <-time.After(time.Second * 10):
			//println("10s timer")
		}
	}
}

func NewHero() *Hero {
	c := new(Hero)
	c.Items = make([]int, 30)
	c.LoadFromFile("hero.json")
	c.ReBuild()
	go c.run()
	return c
}

func main() {
	c := NewHero()
	c.LoadFromFile("hero.json")
	c.ShowAllItems()
}
