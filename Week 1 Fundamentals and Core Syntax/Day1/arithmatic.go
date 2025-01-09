package main

import "fmt"


func main(){

	var a int
	var b int


	fmt.Print("Enter two number separated by space\n")
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println("Sum of two number is ", a+b)
	fmt.Println("Difference of two number is ", a-b)
	fmt.Println("Product of two numer is ", a*b)
	fmt.Println("quotient of two number is ", a/b)

}