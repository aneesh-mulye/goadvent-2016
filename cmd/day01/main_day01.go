package main

import (
	"fmt"
	"github.com/aneesh-mulye/goadvent-2016/internal/day01"
	"github.com/aneesh-mulye/goadvent-2016/util"
	"os"
)

func main() {
	input, err := util.DayFileLines(1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't read day's input file: %s", err.Error())
		return
	}

	var answer uint
	answer, err = day01.Day01p1(input[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error computing answer to 01.p1: %s", err.Error())
		return
	}
	fmt.Println(answer)

	answer, err = day01.Day01p2(input[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error computing answer to 01.p1: %s", err.Error())
		return
	}
	fmt.Println(answer)

}
