package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

const RAWBOARD = "data/rawboard.csv"
const verbose = 1

func main() {
	// PqMain()
	// os.Exit(0)
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
	g := NewGraph(rawboard, 4)
	if verbose > 1 {
		fmt.Printf("node = (%v, %v) distance %v \n", node12.Row, node12.Col, node12.Distance)
		fmt.Println(node12.SprintNode())
		node12.Distance = 42
		fmt.Println("dist: ", node12.PrDist())
		g.PrintBoard()
	}

	// visited, goals := dijkstra(g, g.Nodes[0], math.MaxInt, 3)
	visited, goals := dijkstra(g, g.Nodes[0], 3, verbose)
	if verbose > 1 {
		fmt.Printf("%v\n", visited)
		fmt.Printf("%v\n", goals)
		g.PrintBoard()
	}

	iterations := 10000
	start := time.Now().UnixMilli()
	for n := 0; n < iterations; n++ {
		g.ResetGraph()
		_, _ = dijkstra(g, g.Nodes[0], math.MaxInt, verbose)
	}
	end := time.Now().UnixMilli()
	elapsed := end - start
	fmt.Printf("elapsed: %v\n", elapsed)
	fmt.Printf("Time per iteration: %v ms", float64(elapsed)/float64(iterations))
}
