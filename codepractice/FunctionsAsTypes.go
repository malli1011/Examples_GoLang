package codepractice

import "fmt"

func demo() {
	//anonymous function

	func(i int) {
		fmt.Println("Value of i is ", i)
	}(10)

	double := func(i int) int {
		fmt.Println("Value of i is ", i)
		return i + i
	}

	fmt.Printf("type of double %T", double)
	fmt.Println(double(10))

	val := incrementor(0)
	fmt.Printf("%T", val)
	fmt.Println(val())
	fmt.Println(incrementor(0)())
}

//Function as return type

func incrementor(by int) func() int {
	fmt.Println("Executing incrementor")
	var x int
	return func() int {
		fmt.Println("Executing inner function")
		x += by
		return x
	}
}

//callback functions

func CalcSum(i ...int) (total int) {
	for _, val := range i {
		total += val
	}
	return
}

func EvenSum(f func(i ...int) int, vi ...int) int {
	var temp []int

	for _, val := range vi {
		if val%2 == 0 {
			temp = append(temp, val)
		}
	}
	return f(temp...)
}

var foo = func() {
	fmt.Println("Hello world from foo method")
}

func bar(f func()) {
	fmt.Println("Hello world from bar method")
	fmt.Println("Executing foo")
	f()
}

// ClosureDemo closure demo
func ClosureDemo() {
	a := incrementor(2)
	b := incrementor(5)

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}
