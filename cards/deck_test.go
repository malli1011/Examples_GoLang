package cards

import (
	"os"
	"reflect"
	"testing"
)

func Test_newDeck1(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Test Failed: expected len %d, actual len %d", 16, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected 'Ace of Spades' as first element but go '%v'", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected 'Four of Clubs' as last element but go '%v'", d[len(d)-1])
	}

}

func Test_newDeck(t *testing.T) {
	tests := []struct {
		name string
		want deck
	}{
		{"first test", deck{"Ace of Spades", "Two of Spades", "Three of Spades", "Four of Spades", "Ace of Diamonds", "Two of Diamonds",
			"Three of Diamonds", "Four of Diamonds", "Ace of Hearts", "Two of Hearts", "Three of Hearts", "Four of Hearts", "Ace of Clubs", "Two of Clubs", "Three of Clubs", "Four of Clubs"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDeck(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDeck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deal(t *testing.T) {
	d1, d2 := deal(newDeck(), 4)

	d1.display()
	d2.display()
}

func Test_saveToFile(t *testing.T) {
	newDeck().saveToFile("My_deck.txt")
}

func Test_newDeckFromFile(t *testing.T) {
	d := newDeckFromFile("My_deck.txt")
	d.shuffle()
	d.display()
}

func Test_SaveToFileAndNewDeckFromFile(t *testing.T) {

	fileName := "_decktesting"
	os.Remove(fileName)

	deck := newDeck()

	deck.saveToFile(fileName)
	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 card in deck, but got %v", len(loadedDeck))
	}
	os.Remove(fileName)
}
