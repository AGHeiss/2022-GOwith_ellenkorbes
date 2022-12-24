package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var contador int32

const QuantGoRoutines int = 50000

func CriarGoRoutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			// func AddInt32(addr *int32, delta int32) (new int32)
			// func LoadInt32(addr *int32) (val int32)
			atomic.AddInt32(&contador, 1)
			v := atomic.LoadInt32(&contador)
			runtime.Gosched()
			fmt.Println(v)
			wg.Done()

		}()
	}
}

func main() {
	CriarGoRoutines(QuantGoRoutines)
	wg.Wait()
	fmt.Println("Total Go Routines:", QuantGoRoutines, "Total contador:", contador)
}
