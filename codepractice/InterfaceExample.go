package codepractice

import "fmt"

//https://www.ardanlabs.com/blog/2015/09/composition-with-go.html
//polymorphism
//any type (struct) which has method display is also of human type
type human interface {
	display()
}

func InterfaceDemo(h human) {
	val := fmt.Sprintf("I accept any value of type human %[1]T , %[1]v", h)
	fmt.Println(val)
}

func display(h human) {
	switch h.(type) {
	case person:
		fmt.Println("Human is of type Person")
		h.display()
	case programmer:
		fmt.Println("Human is of type Programmer")
		h.display()
	default:
		fmt.Println("Human is of type something else")

	}
}

func TestInterface() {
	per := person{
		"Sakthi",
		30,
		'F',
		[]string{"vanilla", "pista", "coffee"},
	}

	pro := programmer{
		person: person{
			"Malli",
			30,
			'M',
			[]string{"vanilla", "pista", "coffee"},
		},
		language: "Java",
	}

	display(per)
	display(pro)
}
