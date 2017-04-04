package level

import (
	. "../monster"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type Level struct {
	Name     string   `json:"name"`
	Paths    []string `json:"path"`
	Monsters []Monster
}

func (l *Level) LoadLevel(path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return
	}
	if err := json.Unmarshal(bytes, l); err != nil {
		fmt.Println("Unmarshal json failed.", err)
		return
	}

	for _, s := range l.Paths {
		fmt.Printf("list:%s\n", s)
		m := NewMonster()
		m.LoadFromFile(s)
		l.Monsters = append(l.Monsters, *m)
	}
}

func (l *Level) ShowMonsters() {
	for _, m := range l.Monsters {
		fmt.Printf("level:%s|monster:%s.\n", l.Name, m.Name)
	}
}

func randInt(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}

func (l *Level) GetOneMonster() Monster {
	r := randInt(len(l.Monsters))
	m := l.Monsters[r]
	fmt.Printf("Monster:%s\n", m.Name)
	return m
}

func NewLevel() *Level {
	l := new(Level)
	l.Monsters = make([]Monster, 0)
	return l
}

func main() {
	l := NewLevel()
	l.LoadLevel("../level.json")
	l.ShowMonsters()
	l.GetOneMonster()
}
