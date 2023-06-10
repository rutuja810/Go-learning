package main

import "fmt"

func main()  {
	// card := "Hey"
	// fmt.Println(card)
	// fmt.Println("==================")
	// var deckcard string = "Hey"
	// fmt.Println(deckcard)
	// fmt.Println("==================")
	// cards := []string {"Rutuja", "Rutuja", "Rutuja"}
	// cards = append(cards, "Vikrant")
	// for i, card := range cards{
	// 	fmt.Println(i, card)
	// }
	// fmt.Println("==================")
	// deckcards := deck{"Rutuja", "Rutuja", "Rutuja"} //<- run both the go files main.go and deck.go "go run main.go deck.go"

	// for i, card := range deckcards{
	// 	fmt.Println(i, card)
	// }
	// fmt.Println("==================")
	// deckcards.print()
	// fmt.Println("==================")

	// ls := laptopSize(3.14)
	// fmt.Println(ls.getSizeOfLaptop())
	// fmt.Println("==================")

	numbers := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range numbers {
		if(number % 2 == 0) {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}

/*
➜  cards git:(main) ✗ go run main.go deck.go                                                               
Hey
==================
Hey
==================
0 Rutuja
1 Rutuja
2 Rutuja
3 Vikrant
==================
0 Rutuja
1 Rutuja
2 Rutuja
==================
0 Rutuja
1 Rutuja
2 Rutuja
==================
3.14
==================
*/