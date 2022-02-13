package codepractice

import (
	"fmt"
	"io"
	"os"
)

var x int
var y string
var z bool

type hotdog int

var mytype hotdog

func example1() {
	fmt.Println(x, y, z)
	x = 42
	y = "James Bond"
	z = true
	s := fmt.Sprintf("%v\t%v\t%v\t\n", x, y, z)
	io.WriteString(os.Stdout, s)

	mytype = hotdog(x)
	fmt.Printf("%T\n", mytype)
	fmt.Printf("%T\n", x)
	fmt.Println(mytype)

	z := int(mytype)

	fmt.Printf("Type of z is %T and Value is %v", z, z)

}
