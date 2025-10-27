package main

import (
	"fmt"
	"rpg-hero/internal/hero"
)

func main() {
	arantir := hero.CreateHero("Arantir", 12, 400)
	fmt.Printf("Hero created: %s level %v gold %v", arantir.Name, arantir.Level, arantir.Gold)
}
