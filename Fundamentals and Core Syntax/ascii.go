package main


import "fmt"
import "strconv"


func main(){

	var a int

	fmt.Scan(&a)
	
	s := strconv.Itoa(a);
	ascii := []rune(s)[0]
	fmt.Println(ascii);
}