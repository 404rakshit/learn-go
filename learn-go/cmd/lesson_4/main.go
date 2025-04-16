package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := []rune("résumé")
	indexed := myString[1]
	fmt.Printf("%v %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("The lenght of 'myString' is %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)
	concatStrWithBuilder()
}

func concatStrWithBuilder() {
	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var strBuilder strings.Builder

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}

// it creates new string everyIteration, cuz golang has immutable string use string builder instead
func concateStr() {

	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var catStr = ""

	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)
}
