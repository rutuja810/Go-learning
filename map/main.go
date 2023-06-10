package main

import "fmt"

func main() {
	var slice = make([]string, 1)
	slice[0] = "rutuja"
	fmt.Println(slice)
	// Ways to declare maps
	//1)
	colours := map[int]string{
		1 : "red",
		2 : "orange",
		3 : "yellow",
	}
	fmt.Println(colours)
	delete(colours, 1)
	fmt.Println(colours)
	//2)
	var colour map[int]string 
	fmt.Println(colour)
	//3)
	var rainbow = make(map[int]string)
		rainbow[1] = "red"
		rainbow[2] = "orange"
		rainbow[3] = "yellow"
	
	fmt.Println(rainbow)
	printmap(rainbow)
}

func printmap(rainbow map[int]string) {
	for index,c := range rainbow {            //index = key in map, c=value in map
		fmt.Println(index, ":", c)
	}
}