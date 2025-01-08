package main

import "fmt"

const PI = 3.14159

func main(){

	r := 1.0

	perimeter := 2*PI*r
	area := PI*r*r

	fmt.Printf("Perimeter of Circle %f\n", perimeter)
	fmt.Printf("Area of Circle %f\n", area)
}

// to multiply float and int (or vice versa) we need to typecast one of them 