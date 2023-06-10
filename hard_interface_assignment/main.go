package main

import ( 
	"io"
	"os"
)

func main () {
	//fmt.Println(os.Open(os.Args[1]))
	file,_ := os.Open(os.Args[1])
	//The identifier represented by the underscore character(_) is known as a blank identifier. It is used as an anonymous placeholder instead of a regular identifier. 
	//Here _ represents anonymous placeholder for err identifier.

	io.Copy(os.Stdout, file)

}

/*OUTPUT
➜  hard_interface_assignment git:(main) ✗ go run main.go myfile.txt                                         <region:us-east-1>
Someone is sitting in the shade today because someone planted a tree a long time ago.%
*/