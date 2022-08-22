package main


import "fmt"

func main(){
	cards := []string{"ace", "of", "spades", newCard()}
	cards = append(cards, "joker")
	fmt.Println(cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string{
	return "five of diamonds"
}