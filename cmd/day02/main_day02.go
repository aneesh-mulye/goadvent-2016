package main

import (
	"fmt"
	"github.com/aneesh-mulye/goadvent-2016/internal/day02"
	"github.com/aneesh-mulye/goadvent-2016/util"
	"os"
)

func main() {
	input, err := util.DayFileLines(2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't read day's input file: %s", err.Error())
		return
	}

	var code []int
	code, err = day02.Day02p1(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error computing answer to 02.p1: %s", err.Error())
		return
	}
	fmt.Println(code)

	var code2 []rune
	code2, err = day02.Day02p2(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error computing answer to 02.p2: %s", err.Error())
		return
	}
	fmt.Printf("%s\n", string(code2))
}
