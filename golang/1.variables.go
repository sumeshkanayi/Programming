package main

import "fmt"

func main() {
	//Integer
	var intVar int
	intVar = 100
	fmt.Println(intVar)
	//String
	var StringVar string
	StringVar = "This is a string"
	fmt.Println(StringVar)
	//Another way of declaring variables .Note using var and type
	StringVar2 := " I am a string"
	fmt.Println(StringVar2)
	//Arrays
	stringarray := [3]string{"a", "b", "c"}
	//boolean
	present := false

	if len(stringarray) == 3 {

		present = true
		fmt.Println(present)
	}

}
