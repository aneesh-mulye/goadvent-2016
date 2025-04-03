package day01

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day01p1(input string) (uint, error) {
	instructions := strings.Split(input, ", ")

	dir := 0 + 1i
	location := 0 + 0i
	for _, inst := range instructions {
		// First, update the direction
		switch inst[0] {
		case 'R':
			dir *= 0 - 1i
		case 'L':
			dir *= 0 + 1i
		}
		// Then, parse the number remaining and add that to the location
		n, err := strconv.Atoi(inst[1:])
		if err != nil {
			return 0, fmt.Errorf("Couldn't convert instruction to integer: %w", err)
		}
		location += dir * complex(float64(n), 0)
	}

	return uint(math.Abs(real(location)) + math.Abs(imag(location))), nil
}
func Day01p2(input string) (uint, error) {
	instructions := strings.Split(input, ", ")

	visited := map[complex128]bool{
		// Not sure if this is neeeded? not uncommenting unless problem.
		//complex(0, 0): true
	}

	dir := 0 + 1i
	location := 0 + 0i
out:
	for _, inst := range instructions {
		// First, update the direction
		switch inst[0] {
		case 'R':
			dir *= 0 - 1i
		case 'L':
			dir *= 0 + 1i
		}
		// Then, parse the number remaining and add that to the location
		n, err := strconv.Atoi(inst[1:])
		if err != nil {
			return 0, fmt.Errorf("Couldn't convert instruction to integer: %w", err)
		}
		for range n {
			location += dir
			if visited[location] {
				break out
			}
			visited[location] = true
		}
	}

	return uint(math.Abs(real(location)) + math.Abs(imag(location))), nil
}
