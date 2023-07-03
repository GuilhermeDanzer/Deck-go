package main

func main() {
	cards := newDeckFromFile("my_cards")
	newDeck()
	cards.shuffle()
	cards.print()

}
