package main

import "fmt"

func main() {
	// slicePointerRef()
	pointerFuncDemo()
}
func pointerFuncDemo() {
	var thing [5]float64 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\n Memory Location for thing2: %p", &thing)
	var result = square(&thing)
	fmt.Printf("\n The value for result: %v", result)
	fmt.Printf("\n The value for thing: %v", thing)
}

// pointers and functions goes best, as passing any arg to a func creates a new Variable unless you pass their pointers instead
func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\n Memory Location for thing2: %p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}

func slicePointerRef() {
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
}

func pointerReference() {
	// if not init, it shows up as nil during pointer dereferencing below
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)
	p = &i
	// changing the value of i usiong pointer referncing
	*p = 1
	fmt.Printf("\n The value for i is %v", i)
	var k int32 = 5
	i = k
	fmt.Printf("\n The value for i is %v", i)
}
