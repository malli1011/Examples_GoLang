package codepractice

import "fmt"

type person2 struct {
	name string
	age  int
}

func (person person2) changeMe(p *person2) {

	//We can access the variable from a pointer in one of the below 2 ways
	p.name = "Changed Malli"
	(*p).age = 31
}

/*func (person *person2) changeMePointer(p person2) {

	//We can access the variable from a pointer in one of the below 2 ways
	p.name = "Changed Malli"
	(p).age = 31
}*/

func Demo() {
	p1 := person2{"Malli", 30}
	fmt.Println("Person before Marriage:", p1.name, p1.age)
	p1.changeMe(&p1)
	fmt.Println("Person after Marriage:", p1.name, p1.age)
}
