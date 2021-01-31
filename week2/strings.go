package main

import (
	"strings"
	"fmt"
)

func main() {
	/* 	//This is an example of a string
	   	fmt.Println("Hello World")

	   	var ourString string
	   	ourString = "This is a Test"

	   	fmt.Println(ourString)

	   	//This is an example of a Rune
	   	var ourRune rune
	   	ourRune ='g'
	   	fmt.Printf("Type: %T, Value: %v\n", ourRune, ourRune)

	   	//Example of needing special characters
	   	fmt.Println("Don't Touch That, that's my food")
	   	fmt.Println("She said, \"Thats my food too\"") */

	//Example of a string literal
	fmt.Println(`Hello my name is Ali
	and I am the coolest person`)

	//Example of string literal
	howLong := len("bre ak")

	fmt.Println(howLong)

	//String Formatting
	//name := "Ali"
	name := "Ali"
	age := 26
	fmt.Printf("Hello my name is %s and I am %d years old\n", name, age)

	//String Method 
	//Uppercase all
	fmt.Println(strings.ToUpper("Gopher dbhdbhdbhwhbdwhj"))

	//Example Split 
	s := "Why Are We Here"
	sSplit := strings.Split(s, " ")
	fmt.Println(sSplit)

	//String Concatenation 
	fmt.Println("Hello" + "World")

}
