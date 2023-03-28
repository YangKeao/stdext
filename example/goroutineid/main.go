package main

import (
	"fmt"
	"sync"

	"github.com/YangKeao/stdext/goroutine"
)

func main() {
	id := goroutine.Id()
	fmt.Println("goroutine id in main:", id)

	fmt.Println("spawn 100 goroutine and see their id")
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		i := i
		wg.Add(1)
		go func() {
			id := goroutine.Id()
			fmt.Printf("goroutine id in goroutine %d is %d\n", i, id)
			wg.Done()
		}()
	}
	wg.Wait()
}
