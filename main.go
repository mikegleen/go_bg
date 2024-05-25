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

type RawBoardType [][]string

type GameType struct {
	nplayers        int
	black_train_col int
	selling_price   []int
	players         []*PlayerType
	graph           *Graph
	oilmarker_stock int
	beigecards      *BeigeCards
	redcards        *RedCards
}

func newGame(rawboard RawBoardType, args *argstruct) *GameType {
	game := new(GameType)
	game.graph = NewGraph(rawboard, args.nplayers)
	game.nplayers = args.nplayers
	game.black_train_col = 0
	game.selling_price = make([]int, NCOMPANIES)
	for n := 0; n < NCOMPANIES; n++ {
		game.selling_price[n] = INITIAL_PRICE
	}
	game.players = make([]*PlayerType, game.nplayers)
	for p := 0; p < args.nplayers; p++ {
		player := new(PlayerType)
		game.players[p] = player
		player.Id = p
		trucknode := game.graph.Board[p][0]
		game.players[p].TruckNode = trucknode
	}
	game.beigecards = NewBeigeCards()
	game.redcards = make([]*RedCard, len(RED_CARDS))
	return game
}

func printGame(g *GameType) {
	fmt.Printf("Players: %v, Black Train Col: %v\n", g.nplayers, g.black_train_col)
	fmt.Println("Selling Price:", g.selling_price)
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
		switch u {
		case "c":
			args.config = true
		case "d":
			args.dijkstra = true
		case "p":
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
		//rawBoard := ReadRawBoard(args.board, false)
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
	rawBoard = ReadRawBoard(args.board, false)
	if verbose > 1 {
		fmt.Println("Printing rawboard.")
		for _, row := range rawBoard {
			fmt.Printf("%v\n", row)
		}
	}
	for gamen := 0; gamen < args.games; gamen++ {
		game := newGame(rawBoard, &args)
		printGame(game)
	}

	fmt.Println(green("End Giganten."))
}
