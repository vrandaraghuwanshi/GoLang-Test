package main

func main() {
	//slice of bytes example
	// greeting := "Hi there"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// cards.saveToFile("my_card")
	// cards.print()

	//error handling
	// cards := newDeckFromFile("my_card")
	// cards.print()

	//shuffle cards function
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// fmt.Println(cards)
	// fmt.Println(cards.toString())
	// variable declaration

	// var card string = "Ace of Spades"
	// card := "Ace of Spades" //new variable decalration

	// card := newCard() //assign function to variable

	// card = "Five of Diamonds"

	//slice

	// cards := []string{"Ace of Spades", newCard()}

	// cards := deck{"Ace of Spades", newCard()}

	// cards = append(cards, "Six of Spades")

	// cards := newDeck()

	// // deal(cards, 5)

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	//for loop with slice

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// fmt.Println(cards)
	// cards.print()
}

//new function

// func newCard() string {
// 	return "Five Of Diamonds"
// }

// func newCardInt() int {
// 	return 12
// }
