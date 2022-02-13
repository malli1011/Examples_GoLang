package codepractice

import (
	"fmt"
)

type person struct {
	name       string
	age        int
	sex        rune
	favFlavors []string
}

type programmer struct {
	person
	language string
}

func DisplayStructVal() {
	p := person{
		name:       "Malli",
		age:        33,
		sex:        'M',
		favFlavors: []string{"vanilla", "pista", "coffee"},
	}
	fmt.Println(p)
	fmt.Println(p.name, p.age, p.sex)
	for _, v := range p.favFlavors {
		fmt.Println(v)
	}

	p.name = "Arjun"
	fmt.Println(p.name, p.age, p.sex)

	pro1 := programmer{
		person:   p,
		language: "GoLang",
	}
	fmt.Println(pro1)

	defer fmt.Println("Executing defer statement1: ", pro1.name, pro1.age, pro1.language)

	pro2 := programmer{
		person: person{
			name: "Sakthi",
			age:  30,
		},
		language: "Kotlin",
	}
	fmt.Println(pro2)

	defer fmt.Println("Executing defer statement2: ", pro2.name, pro2.age, pro2.language)
	pro2.display()

}

func MapOfPersonsDemo() {
	var persons = make(map[string]person, 2)
	p1 := person{
		name: "Malli",
		age:  20,
	}

	p2 := person{
		name: "Sakthi",
		age:  20,
	}

	persons[p1.name] = p1
	persons[p2.name] = p2

	for key, val := range persons {
		fmt.Println(key, val)
	}
	AnonymousStruck()
}

func AnonymousStruck() {
	s := struct {
		person
		name string
		age  int
		sex  rune
	}{
		name: "malli",
		age:  33,
		sex:  'M',
		person: person{
			name:       "Arjun",
			age:        30,
			sex:        'M',
			favFlavors: []string{"hot", "spicy"},
		},
	}

	fmt.Println("Anonymous Structs: ", s.name, s.age, s.sex)
	fmt.Println("Anonymous Structs inherited data: ", s.person.name, s.person.age, string(s.person.sex), s.favFlavors[0], "and", s.favFlavors[1])
}
