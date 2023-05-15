package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first crd of ace of spade but got %v", d[0])
	}
}

func SaveDeckAndGetDeckFromFile(t *testing.T) {
	os.Remove("_docktesting")
	d := newDeck()
	d.saveToFile("_docktesting")
	loadedDeck := newDeckFromFile("_docktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(loadedDeck))
	}

}
