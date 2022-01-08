package main

import (
	"fmt"
	"os"
)

const RAWBOARD = "data/rawboard.csv"

func main() {
	var rawboard [][]string
	if len(os.Args) < 2 {
		fmt.Printf("No parameters, using \"%v\"\n", RAWBOARD)
	}
	// rawboard := ReadBoard(RAWBOARD, true)
	// for _, row := range rawboard {
	// 	fmt.Printf("%v\n", row)
	// }
	rawboard = ReadBoard(RAWBOARD, false)
	for _, row := range rawboard {
		fmt.Printf("%v\n", row)
	}

	node12 := NewNode(1, 2, false)
	fmt.Printf("node = (%v, %v) distance %v \n", node12.Row, node12.Col, node12.Distance)
	fmt.Println(SprintNode(node12))
	node12.Distance = 42
	fmt.Println("dist: ", node12.PrDist())
}
