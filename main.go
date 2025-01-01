package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"math"
	"math/rand/v2"
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

type RawBoard [][]string

type Game struct {
	nplayers        int
	black_train_col int
	selling_price   []int
	players         []*Player
	graph           *Graph
	oilmarker_stock int
	beigecards      *BeigeCards
	beigediscards   *BeigeCards
	redcards        *RedCards
	reddiscards     *RedCards
	licensecards    *[]int
	licensediscards *[]int
	tiles           *Tiles
}

func newGame(rawboard RawBoard, args *argstruct) *Game {
	fmt.Println("begin newgame")
	game := new(Game)
	game.graph = NewGraph(rawboard, args.nplayers)
	game.nplayers = args.nplayers
	game.oilmarker_stock = INITIAL_OIL_MARKERS
	game.black_train_col = 0

	// Set the initial selling price for each of the companies
	game.selling_price = make([]int, NCOMPANIES)
	for n := 0; n < NCOMPANIES; n++ {
		game.selling_price[n] = INITIAL_PRICE
	}
	game.players = make([]*Player, game.nplayers)
	for p := 0; p < args.nplayers; p++ {
		trucknode := game.graph.Board[TRUCK_INIT_ROWS[p]][0]
		player := NewPlayer(trucknode, p)
		player.game = game
		game.players[p] = player
		trucknode.Truck = player
	}
	game.beigecards = NewBeigeCards()
	game.redcards = NewRedCards()
	game.licensecards = NewLicenseCards(INITIAL_SINGLE_LICENSES, INITIAL_DOUBLE_LICENSES)
	game.licensediscards = new([]int)
	game.tiles = NewTiles()

	// Place tiles on the board
	for _, node := range game.graph.Nodes {
		if node.Wells > 0 {
			node.OilReserve = game.tiles.PopTile(node.Wells)
			if args.verbose >= 3 {
				fmt.Printf("Node col, row: %d, %d oil reserve %d\n", node.Col, node.Row,
					node.OilReserve)
			}
		}
	}
	return game
}

func (g *Game) move_black_train(spaces_to_move int) bool {
	g.black_train_col += spaces_to_move
	game_ended := g.black_train_col >= g.graph.Columns
	return game_ended
}

func (g *Game) audit_licenses() {
	licenses := 0
	for _, player := range g.players {
		licenses += player.SingleLicenses
		licenses += player.DoubleLicenses
	}
	for _, license_value := range *g.licensecards {
		licenses += license_value
	}
	if licenses != TOTAL_LICENSES {
		panic(fmt.Sprintf("total licenses %d, should be %d", licenses, TOTAL_LICENSES))
	}
}

func (g *Game) draw_red_action_card() *RedCard {
	card := g.redcards.PopRedCard()
	if card != nil {
		return card
	}
	// Move the cards from the discard pile to the active pile
	for len(*g.reddiscards) > 0 {
		l := len(*g.reddiscards)
		r := (*g.reddiscards)[l-1]
		*g.redcards = append(*g.redcards, r)
		*g.reddiscards = (*g.reddiscards)[:l-1]
	}
	if len(*g.redcards) == 0 {
		panic("out of red action cards")
	}
	rc := *g.redcards
	rand.Shuffle(len(rc), func(ii, jj int) { rc[ii], rc[jj] = rc[jj], rc[ii] })
	card = g.redcards.PopRedCard()
	return card
}

func (g *Game) deal_licenses(player Player) {

}

func printGame(g *Game) {
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

func oneTurn(turn int, playerlist []*Player, game *Game) bool {

	return true
}

func PlayGame(game *Game, args *argstruct) {
	gameEnded := false
	turn := 0
	playerList := make([]*Player, 0)
	for !gameEnded {
		turn++
		for startPlayerNum, _ := range game.players {
			playerNum := startPlayerNum
			for n := 0; n < game.nplayers; n++ {
				playerList = append(playerList, game.players[playerNum])
				playerNum++
				if playerNum >= game.nplayers {
					playerNum = 0
				}
			}
			gameEnded = oneTurn(turn, playerList, game)
			if args.verbose >= 2 {
				game.audit_licenses()
			}
		}
	}
}

func main() {
	green := color.New(color.FgGreen).SprintFunc()
	args := getArgs()
	if args.config {
		rawBoard := ReadRawBoard(args.board, false)
		TestConfig(rawBoard, &args)
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
		fmt.Println(game.tiles)
		// printGame(game)
		if args.print {
			printGame(game)
			game.graph.PrintBoard("board:", 1)
		}
	}

	fmt.Println(green("End Giganten."))
}
