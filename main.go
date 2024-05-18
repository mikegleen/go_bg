package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"math"
	"os"
	// "runtime/pprof"
	"strings"
)

const RAWBOARD = "data/rawboard.csv"

var verbose int
var rawBoard [][]string

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
	flag.IntVar(&args.games, "games", 1, `Number of games to play.`)
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
	green := color.New(color.FgGreen).SprintFunc()
	args := getArgs()
	if args.config {
		TestConfig()
		os.Exit(0)
	}
	if args.dijkstra {
		//rawBoard := ReadBoard(args.board, false)
		// g := NewGraph(rawBoard, args.nplayers)
		one_dijkstra(args)
		os.Exit(0)
	}
	if args.pqMain {
		PqMain()
		os.Exit(0)
	}

	verbose = args.verbose
	// fmt.Println("timeit", args.timeit)
	rawBoard = ReadBoard(args.board, false)
	if verbose > 1 {
		fmt.Println("Printing rawboard.")
		for _, row := range rawBoard {
			fmt.Printf("%v\n", row)
		}
	}

	fmt.Println(green("End Giganten."))
}
