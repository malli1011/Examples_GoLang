package codepractice

import (
	"fmt"
	"sync"
)

func FanInDemo() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	//send
	go send2(even, odd)

	go receive2(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("Exit from main method")
}

func send2(e, o chan<- int) {
	for i := 1; i < 25; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(o)
	close(e)
}

func receive2(e, o <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range e {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)
}
