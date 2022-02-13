package codepractice

import (
	"fmt"
	"runtime"
	"sync"
)

//use wait group  to make the main method wait for subroutines to complete
var wg sync.WaitGroup

func DemoSimpleRoutine() {

	//add two subroutines foo() and bar() to monitor for completion.
	wg.Add(2)
	go foo1()
	go bar1()

	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	//wait until both the subroutines complete
	wg.Wait()
}

func foo1() {
	for i := 1; i < 10; i++ {
		fmt.Println("Foooo", i)
	}
	//notify the main method after the subroutine execution
	wg.Done()
}
func bar1() {
	for i := 1; i < 10; i++ {
		fmt.Println("Barrr", i)
	}
	wg.Done()
}

func ChannelDemo() {
	wg.Add(2)
	//create a channel of type int
	ch := make(chan int)

	go func() {
		fmt.Println("Producer", runtime.NumGoroutine())
		runtime.Gosched()
		fmt.Println("Producer put value in channel")
		//put value into channel ch<-10
		ch <- 10

		wg.Done()
	}()

	go func() {
		fmt.Println("Consumer", runtime.NumGoroutine())
		fmt.Println("Consumer waiting for value")
		//retrieve value from in the channel <-c
		fmt.Println("Consumer consumed value", <-ch)
		wg.Done()
	}()
	wg.Wait()
}

func UnidirectionalChannel() {
	//Note: we can convert a bidirectional into Unidirectional but vice-versa is not possible
	//receiving only channel. In this code we can use this channel for putting the values but can't retrieve it back
	ch1 := make(chan<- int)

	//from the below channel we can only retrieve value but we can not add value.
	ch2 := make(<-chan int)

	ch1 <- 10
	fmt.Println(<-ch2)

}
