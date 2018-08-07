package justify

import (
	"strings"
)

// Justify takes a multiline string and justifies the text to the specified length. The last line is not justified.
func Justify(s string, maxLength int) string {
	lines := strings.Split(s, "\n")

	for i, line := range lines {
		// don't justify the last
		if i == len(lines)-1 {
			lines[i] = line
			continue
		}
		lines[i] = justifyLine(line, maxLength)
	}

	return strings.Join(lines, "\n")
}

// justifyLine with determine how many spaces are needed and then applies the added spaces to a given line.
func justifyLine(line string, length int) string {
	spacesNeeded := length - len(line)
	if spacesNeeded <= 0 {
		return line
	}

	if !strings.Contains(line, " ") {
		return line
	}
	indexes := make([]int, 0) // line index to space count
	for i, rn := range line {
		if string(rn) == " " {
			indexes = append(indexes, i)
		}
	}

	spaceCount := getSpaces(len(indexes), spacesNeeded)

	for i := len(indexes) - 1; i >= 0; i-- {
		line = line[:indexes[i]] + strings.Repeat(" ", spaceCount[i]) + line[indexes[i]:]
	}

	return line
}

// getSpaces distributes the number of needed spaces across the known word space count.
// See tests for examples. TL;DR: slotCount=3, spacesNeeded=10 => [3,4,3] which means
// that the first token needs 3 spaces, the second needs 4, and the third needs 3 spaces.
// The idea is to even spread out spaces, giving weight to the center of a line.
func getSpaces(slotCount int, spacesNeeded int) []int {
	// copy over indexes to prevent mutation
	indexes := make([]int, slotCount)

	totalOffset := 0
	middle := len(indexes) / 2
	toggle := 1
	for spacesNeeded > 0 {
		thisOffset := 0
		if toggle > 0 {
			thisOffset = middle + totalOffset
		} else {
			thisOffset = middle - totalOffset
		}

		toggle *= -1
		// only make the offset bigger after we've offset both sides (so every other toggle, increment offset)
		if toggle < 0 {
			totalOffset++
		}

		// reset to middle with initial toggle if we are out of bounds
		if thisOffset < 0 || thisOffset >= len(indexes) {
			totalOffset = 0
			toggle = 1
			continue
		}
		// assign a space
		indexes[thisOffset]++
		spacesNeeded--
	}

	return indexes
}

// MaxLineLength takes a multi-line input and returns the length of the longest line.
func MaxLineLength(s string) int {
	lines := strings.Split(s, "\n")
	max := 0
	for _, line := range lines {
		if len(line) > max {
			max = len(line)
		}
	}
	return max
}
