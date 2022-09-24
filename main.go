package main

import (
	"flag"
	"fmt"
	"math"

	// "runtime/pprof"
	"time"
)

const RAWBOARD = "data/rawboard.csv"

var verbose int

type argstruct struct {
	board    string
	column   int
	dijkstra bool
	games    int
	maxcost  int
	nplayers int
	print    bool
	row      int
	timeit   int
	turns    int
	verbose  int
}

func getargs() argstruct {
	args := argstruct{}
	flag.StringVar(&args.board, "board", RAWBOARD, `The file containing the board description.`)
	flag.IntVar(&args.column, "column", 0, `Start column. For testing.`)
	flag.BoolVar(&args.dijkstra, "dijkstra", false, "Test the dijkstra function.")
	flag.IntVar(&args.games, "games", 1, `Number of games to play. Defaults is 1.`)
	flag.IntVar(&args.maxcost, "maxcost", math.MaxInt, `Maximum distance of interest. For testing.`)
	flag.IntVar(&args.nplayers, "nplayers", 4, `The number of players; the default is 4.`)
	flag.BoolVar(&args.print, "print", false, "Print the finished board with distances.")
	flag.IntVar(&args.row, "row", 0, `Start row. For testing.`)
	flag.IntVar(&args.timeit, "timeit", 0, `Time the dijkstra function with this number of iterations.`)
	flag.IntVar(&args.turns, "turns", math.MaxInt, `Stop the game after this number of turns.`)
	flag.IntVar(&args.verbose, "verbose", 1, `Modify verbosity`)
	flag.Parse()
	return args
}

func main() {

	// cpufile, err := os.Create("cpu.pprof")
	// if err != nil {
	// 	panic(err)
	// }
	// err = pprof.StartCPUProfile(cpufile)
	// if err != nil {
	// 	panic(err)
	// }
	// defer cpufile.Close()
	// defer pprof.StopCPUProfile()

	var rawboard [][]string
	args := getargs()
	verbose = args.verbose
	// fmt.Println("timeit", args.timeit)
	rawboard = ReadBoard(args.board, false)
	if verbose > 1 {
		for _, row := range rawboard {
			fmt.Printf("%v\n", row)
		}
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

	iterations := args.timeit
	start := time.Now().UnixMilli()
	for n := 0; n < iterations; n++ {
		g.ResetGraph()
	}
	end := time.Now().UnixMilli()
	elapsed := end - start
	fmt.Printf("reset time: %v\n", elapsed)
	start = time.Now().UnixMilli()
	for n := 0; n < iterations; n++ {
		g.ResetGraph()
		_, _ = dijkstra(g, g.Nodes[0], math.MaxInt, verbose)
	}
	end = time.Now().UnixMilli()
	elapsed = end - start
	fmt.Printf("elapsed: %v\n", elapsed)
	fmt.Printf("Time per iteration: %v ms\n", float64(elapsed)/float64(iterations))
}
