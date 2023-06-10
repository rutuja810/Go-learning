package main

import "fmt"

type deck []string

func (deckcards deck) print() {
	for i,card := range deckcards {
		fmt. Println(i, card)
	}
	
}

type laptopSize float64
 
func (ls laptopSize) getSizeOfLaptop() laptopSize {
    return ls
}
