package main

import "testing"

func TestRate(t *testing.T) {

	got := rate(0)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

}

func TestWeekday(t *testing.T) {

	got := weekday(0)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestCost(t *testing.T) {
	cost, readingSkipped := cost(101, 0, 0, 1)
	want := float64(0)

	if cost != want {
		t.Errorf("got %f, wanted %f", cost, want)
	}
	if !readingSkipped {
		t.Errorf("readingSkipped is %t, expected true", readingSkipped)
	}
}
