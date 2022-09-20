package main

import (
	"fmt"
	"sync"
)

func main() {
	lett, num := make(chan bool), make(chan bool)

	wg := sync.WaitGroup{}
	go func() {
		for ch := 'A'; ch < 'Z'; ch += 1 {
			lett <- true
			fmt.Print(string(ch))
			<-num
		}
		close(lett)
	}()

	wg.Add(1)
	go func() {
		start := 1
		for {
			num <- true
			fmt.Print(start)
			start += 1
			_, ok := <-lett
			if !ok {
				break
			}
		}
		wg.Done()
	}()

	<-num

	wg.Wait()
	fmt.Print("\n")
}
