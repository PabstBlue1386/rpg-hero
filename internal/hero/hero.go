package hero

import "fmt"

type Hero struct {
	Name      string
	Level     int
	Gold      int
	Inventory map[string]int
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
