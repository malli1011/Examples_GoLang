package codepractice

import "fmt"

// any variable of type programmer can invoke this method.
//programmer type is the receiver of this method.
func (pro2 programmer) display() {
	pro2.person.display()
	fmt.Println("Language :", pro2.language)
}

func (pro2 person) display() {
	fmt.Printf("FirstName: %s, Age: %d, Sex:%c", pro2.name, pro2.age, pro2.sex)
	fmt.Print("\nFavourite Flavours: ")
	for _, v := range pro2.favFlavors {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

func ReceiverExample() {
	pro := programmer{
		person: person{
			"Malli",
			30,
			'M',
			[]string{"vanilla", "pista", "coffee"},
		},
		language: "Java",
	}
	pro.display()
	InterfaceDemo(pro.person)
	InterfaceDemo(pro)
}
