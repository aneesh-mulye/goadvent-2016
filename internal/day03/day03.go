package day03

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func parseInputLine(s string) ([]int, error) {
	sf := strings.Fields(s)
	ret := make([]int, 0)
	for _, field := range sf {
		i, err := strconv.Atoi(field)
		if err != nil {
			return nil, fmt.Errorf("Couldn't parse '%s' into an integer: %w",
				field, err)
		}
		ret = append(ret, i)
	}
	return ret, nil
}

// WARNING!!!
// MUTATES THE INPUT SLICE!
// WARNING!!!
func possibleTriangle(sides []int) (bool, error) {
	if len(sides) != 3 {
		return false, fmt.Errorf("Triangle has 3 sides; %d given", len(sides))
	}
	slices.Sort(sides)
	if sides[0]+sides[1] > sides[2] {
		return true, nil
	}
	return false, nil
}

func Day03p1(input []string) (int, error) {
	var total int
	var err error
	for i, line := range input {
		var sides []int
		sides, err = parseInputLine(line)
		if err != nil {
			return 0, fmt.Errorf("Error parsing line %d: %w", i+1, err)
		}
		var possible bool
		possible, err := possibleTriangle(sides)
		if err != nil {
			return 0, fmt.Errorf("Error on line %d: %w", i+1, err)
		}
		if possible {
			total++
		}
	}
	return total, nil
}

func Day03p2(input []string) (int, error) {
	var total int
	var err error
	var triples [][]int
	for i, line := range input {
		var sides []int
		sides, err = parseInputLine(line)
		if err != nil {
			return 0, fmt.Errorf("Error parsing line %d: %w", i+1, err)
		}
		triples = append(triples, sides)
	}
	parsed := make([]int, 3*len(triples))
	ind := 0
	for i := range 3 {
		for _, triple := range triples {
			parsed[ind] = triple[i]
			ind++
		}
	}
	ind = 0
	for _, triple := range triples {
		for i := range 3 {
			triple[i] = parsed[ind]
			ind++
		}
	}
	for _, triple := range triples {
		var possible bool
		possible, err = possibleTriangle(triple)
		if err != nil {
			return 0, err
		}
		if possible {
			total++
		}
	}
	return total, nil
}
