package main

import "testing"

//how to know what to test?
//what do I really care about?
//easy assertions about the code you're writting

//to test a function always create a test with Test in the begining of the name
//and then put the name of the target function like
func TestNewDeck(t *testing.T) { //the "t" is our test handler/manipulator
	d := newDeck()

	if len(d) != 40 {
		t.Errorf("Expected deck lenght of 20, but got %v", len(d))
	}
}