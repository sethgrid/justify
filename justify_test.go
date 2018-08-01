package justify

import (
	"fmt"
	"strings"
	"testing"
)

func TestJustify(t *testing.T) {
	in := strings.TrimSpace(`
I can't help but think that this is
something that will prove to be harder than
originally imagined. We will see how
this goes.
`)
	expected := strings.TrimSpace(`
I  can't  help  but   think  that  this  is
something that will prove to be harder than
originally  imagined.   We   will  see  how
this goes.
`)

	actual := Justify(in, MaxLineLength(in))
	if actual != expected {
		t.Errorf("Did not match.\nGot\n%s\n\nWant\n%s\n", actual, expected)
	}
}

func TestMaxLineLength(t *testing.T) {
	in := strings.TrimSpace(`
This line is 26 characters
But this line is 30 characters
The final result
should be 30.
	`)

	if got, want := MaxLineLength(in), 30; got != want {
		t.Errorf("got %d, want %d for max line length", got, want)
	}
}

func TestJustifyLine(t *testing.T) {
	tests := []struct {
		line        string
		spacesToAdd int
		output      string
	}{
		{
			line:        "this is something",
			spacesToAdd: 5,
			output:      "this   is    something",
		},
	}
	for _, test := range tests {
		if got, want := justifyLine(test.line, len(test.line)+test.spacesToAdd), test.output; got != want {
			t.Errorf("got\n%s\nwant\n%s", got, want)
		}
	}
}

func TestGetSpaces(t *testing.T) {
	tests := []struct {
		spaceSlots   int
		spacesNeeded int
		expected     []int
	}{
		{
			spaceSlots:   3,
			spacesNeeded: 10,
			expected:     []int{3, 4, 3},
		},
		{
			spaceSlots:   5,
			spacesNeeded: 14,
			expected:     []int{3, 3, 3, 3, 2},
		},
	}

	for _, test := range tests {
		actual := getSpaces(test.spaceSlots, test.spacesNeeded)
		if !sliceEqual(actual, test.expected) {
			t.Errorf("got %#v, want %#v", actual, test.expected)
		}
	}
}

func sliceEqual(m1 []int, m2 []int) bool {
	if fmt.Sprintf("%#v", m1) == fmt.Sprintf("%#v", m2) {
		return true
	}
	return false
}
