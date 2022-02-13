package codepractice

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func FanOUtDemo() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)
	go fanout(c1, c2)

	for val := range c2 {
		fmt.Println(val)
	}

	fmt.Println("Completed the main method")
}

func populate(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanout(c1, c2 chan int) {

	var wg sync.WaitGroup
	for i := range c1 {
		wg.Add(1)
		go func(v1 int) {
			c2 <- timeTakingTask(v1)
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(c2)
}

func timeTakingTask(v1 int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	return v1 + rand.Intn(500)
}
