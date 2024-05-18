package main

import (
	"fmt"
	"math"
	"time"
)

func TestConfig() {

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

	fmt.Println("TRAIN_COSTS", TRAIN_COSTS)
	fmt.Println("INITIAL_CASH", INITIAL_CASH)
	fmt.Println("INITIAL_PRICE", INITIAL_PRICE)
	fmt.Println("INITIAL_OIL_RIGS", INITIAL_OIL_RIGS)
	fmt.Println("INITIAL_OIL_MARKERS", INITIAL_OIL_MARKERS)
	fmt.Println("TRUCK_INIT_ROWS", TRUCK_INIT_ROWS)
	fmt.Println("NCOMPANIES", NCOMPANIES)
	fmt.Println("STORAGE_TANK_LIMIT", STORAGE_TANK_LIMIT)
	fmt.Println("FORCED_SALE_PRICE", FORCED_SALE_PRICE)
	fmt.Println("BUILDING_COST", BUILDING_COST)
	fmt.Println("TRANSPORT_COST", TRANSPORT_COST)
	fmt.Println("END_OF_GAME_RIG_PRICE", END_OF_GAME_RIG_PRICE)
	fmt.Println("END_OF_GAME_MARKER_PRICE", END_OF_GAME_MARKER_PRICE)
	fmt.Println("GOAL_MULTIPLIER", GOAL_MULTIPLIER)
	fmt.Println("TRUCK_COLUMN_MULTIPLIER", TRUCK_COLUMN_MULTIPLIER)
	fmt.Println("PREV_GOAL_MULTIPLER", PREV_GOAL_MULTIPLER)
	fmt.Println("TRAIN_COLUMN_MULTIPLIER", TRAIN_COLUMN_MULTIPLIER)
	fmt.Println("TR_ACTION_CARDS", TR_ACTION_CARDS)
	fmt.Println("TR_COMPUTE_SCORE", TR_COMPUTE_SCORE)
	fmt.Println("TR_FINAL_PATH", TR_FINAL_PATH)
	fmt.Printf("TILES: %v\n", TILES)
	nt := NewTiles()
	fmt.Printf("tiles: %v\n", nt.Tiles)
	n := nt.PopTile(3)
	if n == -1 {
		fmt.Println("oops.")
	}
	fmt.Println("Tile 1: ", n)
	fmt.Printf("tiles: %v\n", nt.Tiles)

	bc := NewBeigeCards()
	l := len(bc.Cards)
	fmt.Printf("(len: %v) Beige cards: %v\n", l, bc.Cards)
	var top *BeigeCard
	top = bc.PopBeigeCard()
	if top != nil {
		fmt.Println("oops. empty")
	}
	l = len(bc.Cards)
	fmt.Printf("top: %v\nlen: %v\nBeige cards: %v\n", *top, l, bc.Cards)
	fmt.Printf("movement: %v\n", top.Movement)
	rc := NewRedCards()
	fmt.Println("Red cards: ", rc.Cards)
	var redtop *RedCard
	redtop = rc.PopRedCard()
	l = len(rc.Cards)
	fmt.Printf("redtop: %v\nlen: %v\nRed cards: %v\n", *redtop, l, rc.Cards)

	licenseCards := NewLicenseCards()
	fmt.Printf("License cards: %v\n", licenseCards.Cards)
	fmt.Println("Pop: ", licenseCards.PopLicenseCard())
	fmt.Println("Pop: ", licenseCards.PopLicenseCard())
	fmt.Println("Pop: ", licenseCards.PopLicenseCard())

	q := make([]int, 10)
	fmt.Printf("q= %v, %p\n", q, &q)

	q = q[:0]
	fmt.Printf("q= %v, %p\n", q, &q)
	q = append(q, 3)
	fmt.Printf("q= %v, %p\n", q, &q)
}

func one_dijkstra(args argstruct) {
	rawBoard := ReadBoard(args.board, false)
	node12 := NewNode(1, 2, false)
	g := NewGraph(rawBoard, 4)
	if verbose > 1 {
		fmt.Printf("node = (%v, %v) distance %v \n", node12.Row, node12.Col, node12.Distance)
		fmt.Println(node12.SprintNode())
		node12.Distance = 42
		fmt.Println("dist: ", node12.PrDist())
		g.PrintBoard("Initial board", verbose)
	}

	// visited, goals := dijkstra(g, g.Nodes[0], math.MaxInt, 3)
	visited, goals := Dijkstra(g, g.Nodes[0], 3, verbose)
	if verbose > 1 {
		fmt.Printf("%v\n", visited)
		fmt.Printf("%v\n", goals)
		g.PrintBoard("Final board", verbose)
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
		_, _ = Dijkstra(g, g.Nodes[0], math.MaxInt, verbose)
	}
	end = time.Now().UnixMilli()
	elapsed = end - start
	fmt.Printf("elapsed: %v ms\n", elapsed)
	fmt.Printf("Time per iteration: %v ms\n", float64(elapsed)/float64(iterations))

}
