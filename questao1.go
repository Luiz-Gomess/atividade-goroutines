package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func main() {

	valores := make(chan int, 20)
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			for j := 1 + i*10; j <= 10 + i*10; j ++ {
		 		valores <- j
		 }
		defer wg.Done()
		}()
	} 

	wg.Wait()	
	close(valores)
	
	for r := range valores{
		fmt.Println("resultado: ", r)
	}

}