package hero

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "../item"
)

type Equipment struct {
	Weapon   Item `json:"weapon"`
	Armor    Item `json:"armor"`
	Shoulder Item `json:"shoulder"`
	Pants    Item `json:"pants"`
	Shoes    Item `json:"shoes"`
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
	Items      []Item    `json:"items"`
}

func (h *Hero) ReBuild() {
	h.Demage = h.BaseDemage + h.Eqpmt.Weapon.Info.Demage
	h.TotalLife = h.BaseLife + h.Eqpmt.Armor.Info.Life
}

func (h *Hero) EquipItem(id int) {
	i := id
	if i < 0 || i >= len(h.Items) {
		return
	}
	h.Items[i].Info.Status = "equiped"
	c := h.Items[i]
	h.Items = append(h.Items[:i], h.Items[i+1:]...)
	switch c.Info.Type {
	case 0:
		h.Eqpmt.Weapon.Info.Status = "not used"
		h.Items = append(h.Items, h.Eqpmt.Weapon)
		h.Eqpmt.Weapon = c
	case 1:
		h.Eqpmt.Armor.Info.Status = "not used"
		h.Items = append(h.Items, h.Eqpmt.Armor)
		h.Eqpmt.Armor = c
	case 2:
		h.Eqpmt.Shoulder.Info.Status = "not used"
		h.Items = append(h.Items, h.Eqpmt.Shoulder)
		h.Eqpmt.Shoulder = c
	case 3:
		h.Eqpmt.Pants.Info.Status = "not used"
		h.Items = append(h.Items, h.Eqpmt.Pants)
		h.Eqpmt.Pants = c
	case 4:
		h.Eqpmt.Shoes.Info.Status = "not used"
		h.Items = append(h.Items, h.Eqpmt.Shoes)
		h.Eqpmt.Shoes = c
	default:
		fmt.Printf("Equipment type error:%d.\n", c.Info.Type)
	}
	h.ReBuild()

}

func (h *Hero) SellItem(id int) {
	i := id
	if i < 0 || i >= len(h.Items) {
		return
	}
	c := h.Items[i]
	h.Items = append(h.Items[:i], h.Items[i+1:]...)
	h.Money = h.Money + c.Info.Money
	fmt.Printf("Sell %s change for %d $.\n", c.Info.Name, c.Info.Money)
}

func (h *Hero) UseItem(id int) {
	i := id
	if i < 0 || i >= len(h.Items) {
		return
	}
	c := h.Items[i]
	if c.Info.Type != 6 {
		fmt.Printf("%s can not be used.\n", c.Info.Name)
		return
	}
	h.Items = append(h.Items[:i], h.Items[i+1:]...)
	h.CurLife = h.CurLife + c.Info.Life
	if h.CurLife > h.TotalLife {
		h.CurLife = h.TotalLife
	}
	fmt.Printf("Use %s heal %d hp, current life:%d.\n",
		c.Info.Name, c.Info.Life, h.CurLife)
	return
}

func (h *Hero) ShowEquipItems() {
	fmt.Printf("Weapon:%10s\nArmor:%10s\nShoulder:%10s\nPants:%10s\nShoes:%10s\n",
		h.Eqpmt.Weapon.Info.Name, h.Eqpmt.Armor.Info.Name, h.Eqpmt.Shoulder.Info.Name,
		h.Eqpmt.Pants.Info.Name, h.Eqpmt.Shoes.Info.Name)
}

func (h *Hero) ShowAllItems() {
	for i, c := range h.Items {
		fmt.Printf("Id:%d,%s", i, c.ShowInfo())
	}
}

func (h *Hero) ShowStatus() {
	fmt.Printf("Name:\t%s\nTotalLife:\t%d\nCurLife:\t%d\nDemage:\t%d\nMana:\t%d\nLevel:\t%d\nMoney:\t%d\n",
		h.Name, h.TotalLife, h.CurLife, h.Demage, h.Mana, h.Level, h.Money)
}

func (h *Hero) ReapItem(it *Item) {
	h.Items = append(h.Items, *it)
	fmt.Printf("%s reap %s.\n", h.Name, it.Info.Name)
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

func NewHero() *Hero {
	c := new(Hero)
	c.Items = make([]Item, 30)
	c.LoadFromFile("hero.json")
	c.ReBuild()
	return c
}

func main() {
	c := NewHero()
	c.LoadFromFile("tmp.log")
	c.ShowAllItems()
}
