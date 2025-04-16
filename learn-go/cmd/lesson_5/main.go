package main

import "fmt"

type gasEngine struct {
	mpg    uint8
	gallon uint8
	owner
	int
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type engine interface {
	milesLeft() uint8
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallon
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	myGasEngine := gasEngine{3, 4, owner{"Hello"}, 10}
	myEleEngine := electricEngine{3, 4}
	canMakeIt(myGasEngine, 10)
	canMakeIt(myEleEngine, 10)
	fmt.Println(myGasEngine.mpg, myGasEngine.gallon, myGasEngine.name, myGasEngine.int)
}
