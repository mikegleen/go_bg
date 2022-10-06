package main

import (
	"fmt"
	"giganten/game"
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

	fmt.Println("TRAIN_COSTS", game.TRAIN_COSTS)
	fmt.Println("INITIAL_CASH", game.INITIAL_CASH)
	fmt.Println("INITIAL_PRICE", game.INITIAL_PRICE)
	fmt.Println("INITIAL_OIL_RIGS", game.INITIAL_OIL_RIGS)
	fmt.Println("INITIAL_OIL_MARKERS", game.INITIAL_OIL_MARKERS)
	fmt.Println("TRUCK_INIT_ROWS", game.TRUCK_INIT_ROWS)
	fmt.Println("NCOMPANIES", game.NCOMPANIES)
	fmt.Println("STORAGE_TANK_LIMIT", game.STORAGE_TANK_LIMIT)
	fmt.Println("FORCED_SALE_PRICE", game.FORCED_SALE_PRICE)
	fmt.Println("BUILDING_COST", game.BUILDING_COST)
	fmt.Println("TRANSPORT_COST", game.TRANSPORT_COST)
	fmt.Println("GAME_END_RIG_PRICE", game.GAME_END_RIG_PRICE)
	fmt.Println("GAME_END_MARKER_PRICE", game.GAME_END_MARKER_PRICE)
	fmt.Println("GOAL_MULTIPLIER", game.GOAL_MULTIPLIER)
	fmt.Println("TRUCK_COLUMN_MULTIPLIER", game.TRUCK_COLUMN_MULTIPLIER)
	fmt.Println("PREV_GOAL_MULTIPLER", game.PREV_GOAL_MULTIPLER)
	fmt.Println("TRAIN_COLUMN_MULTIPLIER", game.TRAIN_COLUMN_MULTIPLIER)
	fmt.Println("TR_ACTION_CARDS", game.TR_ACTION_CARDS)
	fmt.Println("TR_COMPUTE_SCORE", game.TR_COMPUTE_SCORE)
	fmt.Println("TR_FINAL_PATH", game.TR_FINAL_PATH)
	fmt.Printf("TILES: %v\n", game.TILES)
	nt := game.NewTiles()
	fmt.Printf("tiles: %v\n", nt.tiles)
	n := nt.PopTile(3)
	if n == -1 {
		fmt.Println("oops.")
	}
	fmt.Println("Tile 1: ", n)
	fmt.Printf("tiles: %v\n", nt.tiles)

	bc := game.NewBeigeCards()
	l := len(bc.cards)
	fmt.Printf("(len: %v) Beige cards: %v\n", l, bc.cards)
	var top *game.BeigeCard
	top = bc.PopBeigeCard()
	if top != nil {
		fmt.Println("oops. empty")
	}
	l = len(bc.cards)
	fmt.Printf("top: %v\nlen: %v\nBeige cards: %v\n", *top, l, bc.cards)
	fmt.Printf("movement: %v\n", top.movement)
	rc := game.NewRedCards()
	fmt.Println("Red cards: ", rc.cards)
	var redtop *game.RedCard
	redtop = rc.PopRedCard()
	l = len(rc.cards)
	fmt.Printf("redtop: %v\nlen: %v\nRed cards: %v\n", *redtop, l, rc.cards)

	licenseCards := game.NewLicenseCards()
	fmt.Printf("License cards: %v\n", licenseCards.cards)
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
