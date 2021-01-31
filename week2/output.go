package main

import "fmt"

func main() {
	//here is an example of a variable
	name := "Maria"
	fmt.Println(name)

	name = "Aaron"
	fmt.Println(name)
	//fmt.Print("Hello world!\n")

	//Data Types 
	//Example of a string 
	var ourString string 
	ourString = "hfkjlkdjdjlfj"

	fmt.Printf("Type: %T\n", ourString)

	//Example of int 
	var ourInt int 
	ourInt = 55
	fmt.Printf("Type: %T\n", ourInt)

	//Example of a boolena 
	var ourBool bool
	ourBool = true
	fmt.Printf("Type: %T\n", ourBool)

	//Example of a Float 
	var ourFloat float32
	ourFloat = 5.5436
	fmt.Printf("Type: %T\n", ourFloat)

	//Collection 
	//Example of a slice 
	var ourSlice []string 
	ourSlice = []string{"Hello", "goodbye", "Aaron"}
	fmt.Println(ourSlice)

	var ourIntSlice [5]int 
	ourIntSlice = [5]int{5, 4, 2, 875, 424}
	fmt.Println(ourIntSlice)

	//rune byte 
	var ourRune byte
	ourRune = 'h'
	fmt.Printf("Type: %T\n", ourRune) 
	fmt.Printf("Type: %v\n", ourRune) 


}
