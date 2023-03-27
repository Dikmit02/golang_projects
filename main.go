package main

import "fmt"



func main(){
	
	// The append function will not modify the existing card variable
	//instaed it will return the new card variable
	cards := []string{"Ace of Diamond", newCard()}
	cards = append(cards, "Six of Diamond")

	// iterate over a slice
	// here in every loop we are throwing away the previously declared i and card variables
	// thats why we are using :=
	for i, card := range cards{
		fmt.Println(i, card)
	}
	fmt.Println(cards)

}

func newCard() string{
	return "Five of Diamond"
}