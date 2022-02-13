package codepractice

import "fmt"

type myerror struct {
	msg string
}

func (e myerror) Error() string {
	return "Error:" + e.msg
}

func calcSquare(i int) (int, error) {
	if i < 0 {
		return 0, myerror{fmt.Sprintf("Entered negative number, %d. Expecting only postive numbers", i)}
	} else {
		return i * i, nil
	}
}

func CustomeErrorDemo(i int) {
	val, err := calcSquare(i)
	if err != nil {
		fmt.Printf("%T,%v", err, err.Error())
	}
	return
	fmt.Printf("Square of %d is %d\n", i, val)
}
