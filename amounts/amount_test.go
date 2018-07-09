package amounts

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a1 := Parse("$10.30")
	a2 := Parse("$9.70")
	a3 := a1.Add(a2)

	if a3.String() != "$20" {
		t.Errorf("Sum is incorrect, got: %s, want: %s.", a3.String(), "$20")
	}
}

func TestAddZero(t *testing.T) {
	a1 := Parse("$0.01")
	a2 := Parse("$0")
	a3 := a1.Add(a2)

	if a3.String() != "$0.01" {
		t.Errorf("Sum is incorrect, got: %s, want: %s.", a3.String(), "$0.01")
	}
}
