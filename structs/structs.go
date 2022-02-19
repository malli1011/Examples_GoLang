package structs

import "fmt"

type card struct {
	suit  string
	value string
}

//This method with card pointer receiver type  will update the values of the calling object
func (c *card) updateVal(val string) {
	(*c).suit = val
	fmt.Println("Update Value", *c)
}

//This method with card receiver type  will create a new value instead of updating the existing object
//Go is pass by Value language. This is pass by value, we are passing a copy of the calling object into variable c
func (c card) updateVal2(val string) {
	c.suit = val
	fmt.Println("Update Value2", c)
}

func newDeck() card {
	return card{
		"Ace of Spades",
		"Ace of Spades",
	}

}

func demo() {
	c := newDeck()
	c.updateVal("Two of Spades")
	fmt.Println(c)
	c.updateVal2("Three of Spades")
	fmt.Println(c)
}
