package splitz_test

import (
	"testing"

	"github.com/gweithio/splitz"
)

func TestSplitzArrays(t *testing.T) {

	erroGot := splitz.Splitz("")

	if len(erroGot) > 0 {
		t.Errorf("got %q, wanted: %q", erroGot[0], "")
	}

	got := splitz.Splitz("TestWord")

	if got[0] != "Test" {
		t.Errorf("got %q, wanted: %q", got[0], "Test")
	}

	if got[1] != "Word" {
		t.Errorf("got %q, wanted: %q", got[1], "World")
	}

	got2 := splitz.Splitz("TestWordAgain")

	if got2[0] != "Test" {
		t.Errorf("got %q, wanted: %q", got2[0], "Test")
	}

	if got2[1] != "Word" {
		t.Errorf("got %q, wanted: %q", got2[1], "World")
	}

	if got2[2] != "Again" {
		t.Errorf("got %q, wanted: %q", got2[1], "Again")
	}

}
