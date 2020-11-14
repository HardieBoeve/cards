package main

import (
	"fmt"
	"github.com/HardieBoeve/cards/internal/game"
)

func main() {
	cards := game.NewDeck()
	hand, remainingCards := game.Deal(cards, 5)
	hand.Print()
	fmt.Println("")
	remainingCards.Print()
}