// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	// fmt.Println("Hello")
	countToHundred()
}

func log(n int) {
	fmt.Println(n)
}

func countToHundred() {
	var n int = 0
	for i := 0; i < 100000000; i++ {
		n += 1
	}
	log(n)
}
