package main

import "fmt"

func ReturnGeeks(s int) int{
    return s * 5;
}
 
// main function of package
func main() {
    fmt.Println(ReturnGeeks(5))
}