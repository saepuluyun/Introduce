package main

func main() {

	// Variabel Declaration
	// card := newCard()

	// fmt.Println(card)

	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	cards := newDeck()
	// Looping
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// Call display
	// cards.print()

	// Pick from array
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	
	// greeting := "saepul uyun"
	// fmt.Println([]byte(greeting))

	// fmt.Println(cards.toString())

	// cards.saveToFile("my_cards")
	// cards := newDeckFormFile("my_cards")
	// cards.print()

	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
