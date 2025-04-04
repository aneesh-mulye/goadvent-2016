package main

import (
	"fmt"
	"github.com/aneesh-mulye/goadvent-2016/internal/day03"
	"github.com/aneesh-mulye/goadvent-2016/util"
	"os"
)

func main() {
	input, err := util.DayFileLines(3)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't read day's input file: %s", err.Error())
		return
	}

	var possibleTriangles int
	possibleTriangles, err = day03.Day03p1(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}
	fmt.Println(possibleTriangles)
}
