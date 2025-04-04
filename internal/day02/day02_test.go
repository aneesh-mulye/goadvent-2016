package day02_test

import (
	"fmt"
	"github.com/aneesh-mulye/goadvent-2016/internal/day02"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func equalCodeP[T comparable](a, b []T) bool {
	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range len(a) {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestDay02p1(t *testing.T) {

	r := require.New(t)
	a := assert.New(t)

	exInput := []string{"ULL", "RRDDD", "LURDL", "UUUUD"}
	exOutput := []int{1, 9, 8, 5}

	got, err := day02.Day02p1(exInput)
	r.NoError(err)
	a.Equal(true, equalCodeP(got, exOutput),
		fmt.Sprintf("ex %v got %v", exOutput, got))
}

func TestDay02p2(t *testing.T) {

	r := require.New(t)
	a := assert.New(t)

	exInput := []string{"ULL", "RRDDD", "LURDL", "UUUUD"}
	exOutput := []rune{'5', 'D', 'B', '3'}

	got, err := day02.Day02p2(exInput)
	r.NoError(err)
	a.Equal(true, equalCodeP(got, exOutput),
		fmt.Sprintf("ex %v got %v", exOutput, got))
}
