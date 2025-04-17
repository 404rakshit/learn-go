package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// goroutines

var m = sync.RWMutex{} // or sync.Mutex{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

var wg = sync.WaitGroup{}

func main() {
	countWithRoutines()
}

// waitGroups
func countWithRoutines() {
	count := func() {
		var res int
		for i := 0; i < 1000000; i++ {
			res += i
		}
		wg.Done()
	}

	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go count()
	}
	wg.Wait()
	fmt.Printf("\n Total execution time: %v", time.Since(t0))
}

func usualDbCall() {
	dbCall := func(i int) {
		var delay float32 = rand.Float32() * 2000
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Println("The result from the data base is", dbData[i])
	}

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		dbCall(i)
	}
	fmt.Printf("\n Total execution time: %v", time.Since(t0))
}

func dbCallwithWaitGroups() {
	dbCall := func(i int) {
		var delay float32 = rand.Float32() * 2000
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Println("The result from the data base is", dbData[i])
		wg.Done()
	}

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\n Total execution time: %v", time.Since(t0))
}

func dbCallwithGoRoutines() {
	dbCall := func(i int) {
		var delay float32 = rand.Float32() * 2000
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Println("The result from the data base is", dbData[i])
	}

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		go dbCall(i)
	}
	fmt.Printf("\n Total execution time: %v", time.Since(t0))
}

// mutex performs lock & unlock on the current memory
func dbCallwithWaitGroupsMutex() {
	dbCall := func(i int) {
		var delay float32 = 2000
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Println("The result from the data base is", dbData[i])
		// m.Lock()
		// results = append(results, dbData[i])
		// m.Unlock()
		save(dbData[i])
		log()
		wg.Done()
	}

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\n Total execution time: %v", time.Since(t0))
	fmt.Printf("\n The final results are: %v", results)
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\n The current result are: %v", results)
	m.RUnlock()
}
