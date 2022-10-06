package main

import (
	"flag"
	"fmt"
	"giganten/board"
	"giganten/dijkstra"
	"math"
	"os"
	// "runtime/pprof"
	"strings"
	"time"
)

const RAWBOARD = "data/rawboard.csv"

var verbose int

type argstruct struct {
	board    string
	column   int
	config   bool
	dijkstra bool
	games    int
	maxcost  int
	nplayers int
	pqMain   bool
	print    bool
	row      int
	test     string
	timeit   int
	turns    int
	verbose  int
}

func getArgs() argstruct {
	args := argstruct{}
	flag.StringVar(&args.board, "board", RAWBOARD, `The file containing the board description.`)
	flag.IntVar(&args.column, "column", 0, `Start column. For testing.`)
	//flag.BoolVar(&args.dijkstra, "dijkstra", false, "Test the dijkstra function.")
	//flag.BoolVar(&args.pqMain, "pqMain", false, "Test the pqMain function.")
	flag.IntVar(&args.games, "games", 1, `Number of games to play. Defaults is 1.`)
	flag.IntVar(&args.maxcost, "maxcost", math.MaxInt, `Maximum distance of interest. For testing.`)
	flag.IntVar(&args.nplayers, "nplayers", 4, `The number of players.`)
	flag.BoolVar(&args.print, "print", false, "Print the finished board with distances.")
	flag.IntVar(&args.row, "row", 0, `Start row. For testing.`)
	flag.StringVar(&args.test, "test", "", `Test option:
c - config
d - dijkstra
p - pqMain`)
	flag.IntVar(&args.timeit, "timeit", 0, `Time the dijkstra function with this number of iterations.`)
	flag.IntVar(&args.turns, "turns", math.MaxInt, `Stop the game after this number of turns.`)
	flag.IntVar(&args.verbose, "verbose", 1, `Modify verbosity`)
	flag.Parse()
	if args.test != "" {
		u := strings.ToLower(args.test)
		switch {
		case u == "c":
			args.config = true
		case u == "d":
			args.dijkstra = true
		case u == "p":
			args.pqMain = true
		default:
			panic("Invalid test.")
		}
	}
	return args
}

func main() {

	args := getArgs()
	if args.config {
		TestConfig()
		os.Exit(0)
	}
	if args.dijkstra {
		//rawBoard := ReadBoard(args.board, false)
		// g := NewGraph(rawBoard, args.nplayers)
		//one_dijkstra()
		os.Exit(0)
	}
	if args.pqMain {
		dijkstra.PqMain()
		os.Exit(0)
	}
	var rawBoard [][]string
	verbose = args.verbose
	// fmt.Println("timeit", args.timeit)
	rawBoard = board.ReadBoard(args.board, false)
	if verbose > 1 {
		for _, row := range rawBoard {
			fmt.Printf("%v\n", row)
		}
	}

	node12 := board.NewNode(1, 2, false)
	g := board.NewGraph(rawBoard, 4)
	if verbose > 1 {
		fmt.Printf("node = (%v, %v) distance %v \n", node12.Row, node12.Col, node12.Distance)
		fmt.Println(node12.SprintNode())
		node12.Distance = 42
		fmt.Println("dist: ", node12.PrDist())
		g.PrintBoard()
	}

	// visited, goals := dijkstra(g, g.Nodes[0], math.MaxInt, 3)
	visited, goals := dijkstra.Dijkstra(g, g.Nodes[0], 3, verbose)
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
		_, _ = dijkstra.Dijkstra(g, g.Nodes[0], math.MaxInt, verbose)
	}
	end = time.Now().UnixMilli()
	elapsed = end - start
	fmt.Printf("elapsed: %v\n", elapsed)
	fmt.Printf("Time per iteration: %v ms\n", float64(elapsed)/float64(iterations))
}
