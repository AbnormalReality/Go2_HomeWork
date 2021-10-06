package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//добавил для общего развития
	numcpu := runtime.NumCPU()
	fmt.Println("NumCPU", numcpu)
	runtime.GOMAXPROCS(numcpu)
	//runtime.GOMAXPROCS(1)

	var total = 0
	var workers = make(chan int, 1000)
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		workers <- 1

		go func(m *sync.Mutex) {
			m.Lock()
			total += <-workers
			m.Unlock()
		}(&m)
	}

	time.Sleep(time.Second * 1)
	fmt.Printf("Variable value %v\n", total)
}