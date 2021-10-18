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

func TestCsvWriter(t *testing.T) {

}
