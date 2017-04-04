package main

import (
	"./hero"
	. "./item"
	. "./level"
	"./monster"
	. "fmt"
	"math/rand"
	"strconv"
	"time"
)

func randInt(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}

func heroAttack(h *hero.Hero, m *monster.Monster) {
	ret := randInt(100)
	td := h.Demage
	if ret <= h.Luck {
		td = h.Demage + h.ExDemage
	}
	m.CurLife = m.CurLife - td
	Printf("%s attack %s deal with %d demage\n", h.Name, m.Name, td)
}

func monsterAttack(m *monster.Monster, h *hero.Hero) {
	ret := randInt(100)
	td := m.Demage
	if ret <= m.Luck {
		td = h.Demage + m.ExDemage
	}
	h.CurLife = h.CurLife - td
	Printf("%s attack %s deal with %d demage\n", m.Name, h.Name, td)
}

func heroReapTerasures(h *hero.Hero, m *monster.Monster) {
	h.Money = h.Money + m.Money
	Printf("%s find %d gold from %s\n", h.Name, m.Money, m.Name)
	total := len(m.Items)
	Printf("%s has %d items\n", m.Name, total)
	r := randInt(100)
	for _, it := range m.Items {
		if r <= ItemMgr.GetItem(it).Luck {
			h.ReapItem(it)
			return
		}
	}
	Printf("%s find nothing else from %s\n", h.Name, m.Name)
}

func battle(h *hero.Hero, m *monster.Monster) {
	var a, b int
	for {
		if a*1000 >= h.Intval {
			heroAttack(h, m)
			a = 0
		}
		if m.CurLife <= 0 {
			Printf("%s is dead\n", m.Name)
			heroReapTerasures(h, m)
			h.SaveToFile("hero.json")
			break
		}
		if b*1000 >= m.Intval {
			monsterAttack(m, h)
			b = 0
		}
		if h.CurLife <= 0 {
			Printf("%s is dead\n", h.Name)
			break
		}
		time.Sleep(1000 * time.Millisecond)
		a = a + 1
		b = b + 1
	}
}

func gamelogic(h *hero.Hero, l *Level) {
	for {
		Printf("input your command:\n")
		var cmd string
		Scanln(&cmd)
		switch cmd {
		case "f":
			m := l.GetOneMonster()
			battle(h, &m)
		case "show":
			h.ShowStatus()
			h.ShowEquipItems()
			h.ShowAllItems()
		case "equip":
			Printf("Input Item id\n")
			Scanln(&cmd)
			id, _ := strconv.Atoi(cmd)
			h.EquipItem(id)
			h.SaveToFile("hero.json")
		case "sell":
			Printf("Input Item id\n")
			Scanln(&cmd)
			id, _ := strconv.Atoi(cmd)
			h.SellItem(id)
			h.SaveToFile("hero.json")
		case "use":
			Printf("Input Item id\n")
			Scanln(&cmd)
			id, _ := strconv.Atoi(cmd)
			h.UseItem(id)
			h.SaveToFile("hero.json")
		case "help":
			fallthrough
		default:
			Println("support cmds:[fight,show,equip,sell,use,help]")
		}
	}
}

func main() {
	h := hero.NewHero()
	h.ShowAllItems()
	l := NewLevel()
	l.LoadLevel("level.json")
	gamelogic(h, l)
	h.ShowAllItems()
}
