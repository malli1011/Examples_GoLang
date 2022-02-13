package codepractice

import (
	"errors"
	"fmt"
	"io/ioutil"
)

// Implementation with errors.New
// Output: cannot divide by zero
func divide1(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return num1 / num2, nil
}

// Implementation with fmt.Errrof
// Output: cannot divide 10.00 by 0.00
func divide2(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("cannot divide %.2f by %.2f", num1, num2)
	}
	return num1 / num2, nil
}

// The function openFile returns a custom error message if opening the file fails
func openFile(filename string) error {
	if _, err := ioutil.ReadFile(filename); err != nil {
		return fmt.Errorf("error opening %s: %w", filename, err)
	}

	return nil
}

func WrapError() {
	err := openFile("new_file.txt")

	if err != nil {
		fmt.Printf("error running main.go: %s \n", err.Error()) // Print the wrapped error message

		unwrappedErr := errors.Unwrap(err)                 // This line unwraps the error and gives original error
		fmt.Printf("unwrapped error: %v \n", unwrappedErr) // Print the original error message
	}
}

func brokenMessage() error {
	// create the main error here - you can use errors.New or fmt.Errorf
	err := errors.New("error:ice cream machine is broken")

	// create the wrapped error here using fmt.Errorf and the %w verb
	wrappedErr := fmt.Errorf("error:can't serve you an ice cream main cause of %w", err)

	return wrappedErr
}
