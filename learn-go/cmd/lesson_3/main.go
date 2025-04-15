package main

import (
	"fmt"
	"time"
)

// printArray()
// printSlice()
// printMap()
// loops()
func main() {
	execTimeLoop()
	// fmt.Println()
}

func execTimeLoop() {
	n := 1000000
	testSlice := []int{}
	testSlice2 := make([]int, 0, n)
	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

func loops() {

	// while loop
	var i int = 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// for loop
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func printMap() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	myMap2 := map[string]uint8{"Adam": 23, "Labour": 30}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])
	// default 0
	fmt.Println(myMap2["Json"])

	delete(myMap2, "Labour")

	//reading a maps value returns value and a boolean of if it exists
	var age, ok = myMap2["Adam"]
	if ok {
		fmt.Printf("this age is %v \n", age)
	} else {
		fmt.Printf("Invalid Name")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	intArr := []int32{3, 4, 5, 6}

	// iterating an array using for
	for index, value := range intArr {
		fmt.Printf("\nIndex: %v & Value: %v", index, value)
	}
}

func printSlice() {
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with capacity %v", len(intSlice), cap(intSlice))

	intSlice2 := []int32{2, 3}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("\nThe length is %v with capacity %v", len(intSlice), cap(intSlice))

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("\nThe length is %v with capacity %v", len(intSlice3), cap(intSlice3))

}

func printArray() {
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	intArr2 := [...]int32{1, 2, 3, 4}
	fmt.Println(intArr2)
}
