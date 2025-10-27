package hero

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
