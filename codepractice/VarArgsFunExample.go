package codepractice

import "fmt"

func UnfurlingSlice() {
	xi := []int{1, 3, 5, 7, 9, 11}
	//xi... will extract the values from xi
	fmt.Println(sum(xi...))
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sumWithMessage("Hello from var args function!"))
	fmt.Println(sumWithMessage("Hello from var args function!", xi...))

}

/*func sum(vargs ...int) int {
	var sum int
	fmt.Printf("Type of vargs is %[1]T, value of vargs is %[1]v \n", vargs)
	for _, v := range vargs {
		sum += v
	}
	return sum
}*/

func sum(vargs ...int) (sum int) {
	fmt.Printf("Type of vargs is %[1]T, value of vargs is %[1]v \n", vargs)
	for _, v := range vargs {
		sum += v
	}
	return
}

//Return two integers, sum and multiplier
func sumWithMessage(msg string, vargs ...int) (sum, mul int) {
	mul = 1
	fmt.Println("message :", msg)
	for _, v := range vargs {
		sum += v
		mul *= v
	}

	return
}
