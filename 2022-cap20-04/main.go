package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador int

const QuantGoRoutines = 50000

var mu sync.Mutex

func main() {
	GoRoutines(QuantGoRoutines)
	wg.Wait()
	fmt.Println("Total Go Routines:", QuantGoRoutines, "\n Contador:", contador)

}

func GoRoutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
			mu.Unlock()
		}()
	}
}
