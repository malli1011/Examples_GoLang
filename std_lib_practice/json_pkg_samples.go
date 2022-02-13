package std_lib_practice

import (
	"encoding/json"
	"fmt"
	"os"
)

//https://pkg.go.dev/encoding/json#example-Unmarshal

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

// Person using json tags to map Json keys to varibales
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

//This is like toString method override in java
//when ever we print the variable of type person this method gets called
func (p Person) String() string {
	return fmt.Sprintf("%s,%s,%s,%s", p.FirstName, p.LastName, p.Email, p.Gender)
}

func JsonMarshal() {

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	fmt.Println()

	person1 := Person{"Mallikarjun", "Bandi", "mallikarjun@gmail.com", "Male"}
	person2 := Person{"Sakthi", "Bandi", "sakthi@gmail.com", "Female"}

	persons := []Person{person1, person2}

	b, err = json.Marshal(persons)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}

func JsonUnMarshal() {
	s := `{"ID":1,"FirstName":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`
	b := []byte(s)
	var colorGroup ColorGroup
	err := json.Unmarshal(b, &colorGroup)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", colorGroup)
	var persons = make([]Person, 2)
	//persons:=[]Person{}
	b = []byte(`[{"first_name":"Mallikarjun","last_name":"Bandi","email":"test@gmail.com","gender":"Male"},{"FirstName":"Sakthi","LastName":"Bandi","email":"test@gmail.com","gender":"Female"}]`)
	json.Unmarshal(b, &persons)
	for _, per := range persons {
		fmt.Println(per)
	}

}

func validateJson() {
	goodJSON := `{"example": 1}`
	badJSON := `{"example":2:]}}`

	fmt.Println(json.Valid([]byte(goodJSON)), json.Valid([]byte(badJSON)))
}
