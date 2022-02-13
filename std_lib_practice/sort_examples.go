package std_lib_practice

import (
	"fmt"
	"sort"
)

// ByFirstName
//Sort list (slice) of persons by name
//create a custom type, with underlying type as []Person
// ByFirstName implements sort.Interface for []Person based on
// the firstname field.
type ByFirstName []Person

func (a ByFirstName) Len() int           { return len(a) }
func (a ByFirstName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFirstName) Less(i, j int) bool { return a[i].FirstName < a[j].FirstName }

type ByLastName []Person

func (a ByLastName) Len() int           { return len(a) }
func (a ByLastName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLastName) Less(i, j int) bool { return a[i].LastName < a[j].LastName }

func ExampleSortCustomType() {
	people := ReadJsonData()
	fmt.Println(len(people))
	people = people[0:10]
	fmt.Println(len(people))
	fmt.Println(people)
	sort.Sort(ByFirstName(people))
	fmt.Println(people)
	sort.Sort(ByLastName(people))
	fmt.Println(people)

}
