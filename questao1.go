package main

import (
	"fmt"
	"sync"
)

func addNumbers(i int, valores chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 1 + i*10; j <= 10+i*10; j++ {
		valores <- j
	}
}

func main() {

	valores := make(chan int, 20)
	var wg sync.WaitGroup

	wg.Add(2)
	go addNumbers(1, valores, &wg)
	go addNumbers(10, valores, &wg)

	go func() {
		wg.Wait()
		close(valores)
	}()

	for r := range valores {
		fmt.Println("resultado: ", r)
	}

}
