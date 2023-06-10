package main

import ("fmt"
		"net/http"
		"io"
		"os"
)

func main(){
	//The interface is a collection of methods as well as it is a custom type.
	//You can use interface when in methods or functions you want to pass different types of argument in them just like Println () function. 
	//Or you can also use interface when multiple types implement same interface.
	//In the Go language, it is necessary to implement all the methods declared in the interface for implementing an interface.

	resp,err := http.Get("https://youtube.com/")
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//fmt.Println(resp)
	//  bs :=make([]byte, 9999999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	io.Copy(os.Stdout, resp.Body)
}