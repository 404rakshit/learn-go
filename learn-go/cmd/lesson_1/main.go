package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	printBoolean()
}

func printInt() {
	var intNum int16 = 32767
	intNum = intNum + 1
	fmt.Println(intNum)
}

func printFloat() {
	var floatNum float64 = 1234567.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.3
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	intNum1 := 3
	intNum2 := 2
	fmt.Println(intNum1 / intNum2)
}

func printString() {
	var myString string = "a"
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString(myString))

	var myRune rune = 'a'
	// prints ascii code
	fmt.Println(myRune)
}

func printBoolean() {
	var boolean bool = true
	fmt.Println(boolean)

	myVar := "text"
	fmt.Println(myVar)

	const myConst string = "text"
	fmt.Println(myConst)
}
