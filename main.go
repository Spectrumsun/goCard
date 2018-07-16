package main


func main() {
	// cards := newDeck()
	// hand, reaminingCards := deal(cards, 5)
	// hand.print()
	// reaminingCards.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	// cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.print()



}

