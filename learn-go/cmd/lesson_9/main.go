package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

type gasEngine struct {
	gallon float32
	mpg    float32
}

type electricEngine struct {
	kwh   float32
	mpkmh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModal string
	engine   T
}

func main() {
	// exapmleGenericFunc()
	// readDataFromFile()
	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModal: "Civic",
		engine: gasEngine{
			gallon: 12.4,
			mpg:    23,
		},
	}

	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModal: "Model 3",
		engine: electricEngine{
			kwh:   14.3,
			mpkmh: 20,
		},
	}

	fmt.Println(gasCar)
	fmt.Println(electricCar)
}

func readDataFromFile() {
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("\n%v", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	fmt.Printf("\n%v", purchases)
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	data, _ := os.ReadFile(filePath)
	var loaded = []T{}
	json.Unmarshal(data, &loaded)
	return loaded
}

// supports only types like int | float32 | float64
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func exapmleGenericFunc() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice(intSlice))

	var floatSlice = []float32{}
	fmt.Println(sumSlice(floatSlice))

	fmt.Println(isEmpty(intSlice))
	fmt.Println(isEmpty(floatSlice))
}
