package day01_test

import (
	"fmt"
	"github.com/aneesh-mulye/goadvent-2016/internal/day01"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDay01p1(t *testing.T) {
	r := require.New(t)
	a := assert.New(t)

	cases := map[string]uint{
		"L0":     0,
		"R0":     0,
		"L5":     5,
		"R5":     5,
		"L1, R0": 1,
		"R1, L0": 1,
		"L3, R3": 6,
		"R3, L3": 6,
	}

	for i, o := range cases {
		got, err := day01.Day01p1(i)
		r.NoError(err)
		a.Equal(got, o, fmt.Sprintf("P1: Instructions '%s' should result in %d", i, o))
	}
}

func TestDay01p2(t *testing.T) {
	r := require.New(t)
	a := assert.New(t)

	cases := map[string]uint{
		"R8, R4, R4, R8": 4,
	}

	for i, o := range cases {
		got, err := day01.Day01p2(i)
		r.NoError(err)
		a.Equal(got, o, fmt.Sprintf("P2: Instructions '%s' should result in %d", i, o))
	}
}
