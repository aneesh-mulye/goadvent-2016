package day02

import "fmt"

func updatePosP1(pos complex128, dir rune) (complex128, error) {
	addends := map[rune]complex128{
		'L': -1 + 0i,
		'R': 1 + 0i,
		'U': 0 + 1i,
		'D': 0 - 1i,
	}
	switch dir {
	case 'L':
		fallthrough
	case 'R':
		if real(pos) == real(addends[dir]) {
			return pos, nil
		} else {
			return pos + addends[dir], nil
		}
	case 'U':
		fallthrough
	case 'D':
		if imag(pos) == imag(addends[dir]) {
			return pos, nil
		} else {
			return pos + addends[dir], nil
		}
	default:
		return 0 + 0i, fmt.Errorf("%c is not a valid direction", dir)
	}
}

func Day02p1(input []string) ([]int, error) {
	posToCode := map[complex128]int{
		-1 + 1i: 1,
		0 + 1i:  2,
		1 + 1i:  3,
		-1 + 0i: 4,
		0 + 0i:  5,
		1 + 0i:  6,
		-1 - 1i: 7,
		0 - 1i:  8,
		1 - 1i:  9,
	}

	var code []int
	pos := 0 + 0i
	for i, line := range input {
		for _, r := range line {
			var err error
			pos, err = updatePosP1(pos, r)
			if err != nil {
				return nil, fmt.Errorf("Error on line %d: %w", i+1, err)
			}
		}
		code = append(code, posToCode[pos])
	}
	return code, nil
}

var posToCodeP2 map[complex128]rune = map[complex128]rune{
	0 + 0i:  '7',
	-1 + 0i: '6',
	-2 + 0i: '5',
	1 + 0i:  '8',
	2 + 0i:  '9',
	0 + 1i:  '3',
	0 + 2i:  '1',
	0 - 1i:  'B',
	0 - 2i:  'D',
	-1 + 1i: '2',
	1 + 1i:  '4',
	-1 - 1i: 'A',
	1 - 1i:  'C',
}

func updatePosP2(pos complex128, dir rune) (complex128, error) {
	addends := map[rune]complex128{
		'L': -1 + 0i,
		'R': 1 + 0i,
		'U': 0 + 1i,
		'D': 0 - 1i,
	}

	if addends[dir] == 0+0i {
		return 0 + 0i, fmt.Errorf("%c is not a valid direction", dir)
	}
	if posToCodeP2[pos+addends[dir]] == 0 {
		return pos, nil
	} else {
		return pos + addends[dir], nil
	}
}

func Day02p2(input []string) ([]rune, error) {

	var code []rune
	pos := -2 + 0i
	for i, line := range input {
		for _, r := range line {
			var err error
			pos, err = updatePosP2(pos, r)
			if err != nil {
				return nil, fmt.Errorf("Error on line %d: %w", i+1, err)
			}
		}
		code = append(code, posToCodeP2[pos])
	}
	return code, nil
}
