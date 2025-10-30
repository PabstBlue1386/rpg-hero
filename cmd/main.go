package main

import (
	"fmt"
	"rpg-hero/internal/hero"
)

func main() {
	arantir := hero.CreateHero("Arantir", 12, 400)
	fmt.Printf("Hero created: %s level %v gold %v\n", arantir.Name, arantir.Level, arantir.Gold)
	arantir.LevelUp()
	arantir.AddGold(500)
	fmt.Printf("Hero created: %s level %v gold %v\n", arantir.Name, arantir.Level, arantir.Gold)

	err := arantir.Save("hero.json")
	if err != nil {
		return
	}

}
