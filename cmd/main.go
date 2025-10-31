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

	var arantir2 hero.Hero

	if err := arantir2.Load("hero.json"); err != nil {
		fmt.Println("Error read file:", err)
	}
	fmt.Printf("Hero loaded: %s level %v gold %v\n", arantir2.Name, arantir2.Level, arantir2.Gold)

}
