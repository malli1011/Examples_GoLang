package codepractice

import "fmt"

func SelectChannelDemo() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(even, odd, quit)

	receive(even, odd, quit)
	fmt.Println("Exit from main method")
}

func send(e, o, q chan<- int) {
	for i := 1; i < 25; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel:", v)
		case v := <-o:
			fmt.Println("from the odd channel:", v)
		case v := <-q:
			fmt.Println("from the quit channel:", v)
			return
		}

	}

}
