package codepractice

import "fmt"

func PointerExample() {
	var a *int
	a = new(int)

	*a = 10

	b := new(string)

	*b = `Malli`
	c := a

	fmt.Println(*a)
	fmt.Println(&a)
	fmt.Println(a)

	fmt.Printf("Type and Value of 'a' before modification %[1]T, %[1]d", a)

	fmt.Println(b)
	fmt.Println(*b)

	*c = 500

	fmt.Println("Value of 'a' after modification : ", *a)

}
