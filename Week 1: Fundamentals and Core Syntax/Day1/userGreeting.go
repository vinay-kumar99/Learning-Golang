package main


import "fmt"

// n Go, a directory is a package, and a package can only have one function with a given name
func main(){

	fmt.Print("Enter your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Println("Hello ", name)
}