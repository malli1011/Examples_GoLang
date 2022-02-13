package codepractice

import (
	"fmt"
	"time"
)

func UnidirChannelDemo() {
	//create a bidirectional channel
	ch := make(chan int)
	//assign a bidirectional to unidirectional
	go foo3(ch)
	go bar3(ch)

	//main method will wait until the subroutines completes as we are accessing a value from the channel
	//we don't need to use sync.WaitGroup here.
	fmt.Println("Channel value in Main", <-ch)
}

//accepts only receiving channel
func foo3(c chan<- int) {
	c <- 42
	c <- 33
}

//accepts only sending channel
func bar3(c <-chan int) {
	fmt.Println("Channel value in Bar", <-c)
}

func DemoChannelRange() {
	ch := make(chan int)

	go func() {
		for i := 1; i < 10; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
		//close channel
		close(ch)
	}()

	// the range will pull all the values from channel until the channel is closed.
	for val := range ch {
		fmt.Println(time.Now().Format(time.Layout), val)
	}

}
