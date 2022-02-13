package codepractice

import (
	"fmt"
	"os"
)

// OpenFile example of opening a file in the read-only mode
func OpenFile(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	//closing the file
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error occurred while closing the file: ", file.Name())
		}
	}(file)

	/*// this deletes "new_file.txt" from our current working directory
	err := os.Remove("new_file.txt")
	// this deletes "test_directory" and all of its contents
	err := os.RemoveAll("test_directory")*/

}

// OpenFileInWriteAppend is an example for opening the file in write append mode
func OpenFileInWriteAppend(fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error occurred while closing the file: ", file.Name())
		}
	}(file)
}

func CreateFile(fileName string) {
	// this creates a new file, or truncates the already existing file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// this creates a new file using the os.O_CREATE flag and the 0666 permission
	//Using os.OpenFile() with the os.O_CREATE flag will not truncate the file if it already exists, as opposed to os.Create()
	file2, err2 := os.OpenFile("new_file.txt", os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err2)
	}
	defer file2.Close()
}
