package std_lib_practice

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJsonData() []Person {
	file, err := os.Open("mock_data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	var persons = make([]Person, 1000)
	err = dec.Decode(&persons)

	fmt.Println(len(persons))
	return persons
}
