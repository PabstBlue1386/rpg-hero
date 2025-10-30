package hero

import (
	"encoding/json"
	"fmt"
	"os"
)

type Hero struct {
	Name      string         `json:"name"`
	Level     int            `json:"level"`
	Gold      int            `json:"gold"`
	Inventory map[string]int `json:"inventory"`
}

func CreateHero(name string, level, gold int) *Hero {

	return &Hero{
		Name:      name,
		Level:     level,
		Gold:      gold,
		Inventory: make(map[string]int),
	}
}

func (h *Hero) LevelUp() {
	h.Level++
	fmt.Println("Level up!")
}

func (h *Hero) AddGold(amount int) {
	if amount <= 0 {
		fmt.Println("Incorrect value amount!")
		return
	}

	h.Gold += amount
	fmt.Println("That's better!")
}

func (h *Hero) Save(path string) error {

	if h.Inventory == nil {
		h.Inventory = make(map[string]int)
	}

	data, err := json.MarshalIndent(h, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return err
}

func (h *Hero) Load(path string) (*Hero, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &h)
	if h.Inventory == nil {
		h.Inventory = make(map[string]int)
	}

	return h, err
}
