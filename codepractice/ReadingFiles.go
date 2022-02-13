package codepractice

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))
}

func ReadFileLineByLine(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// split each scanned line into words
	// read the file word by word
	// scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// ReadFileInChunks
//When working with big text files, a more efficient approach is to read the file in chunks;
func ReadFileInChunks(fileName string, chunkSize int) {
	file, err := os.Open("test_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create a slice of bytes buffer with the previously defined chunk size
	buf := make([]byte, chunkSize)

	for {
		readTotal, err := file.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break //after reading the last chunk break the loop
			}
			log.Fatal(err)
		}
		fmt.Println(string(buf[:readTotal]))
	}
}
