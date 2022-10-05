package main

import "fmt"

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
	fmt.Println("GAME_END_RIG_PRICE", GAME_END_RIG_PRICE)
	fmt.Println("GAME_END_MARKER_PRICE", GAME_END_MARKER_PRICE)
	fmt.Println("GOAL_MULTIPLIER", GOAL_MULTIPLIER)
	fmt.Println("TRUCK_COLUMN_MULTIPLIER", TRUCK_COLUMN_MULTIPLIER)
	fmt.Println("PREV_GOAL_MULTIPLER", PREV_GOAL_MULTIPLER)
	fmt.Println("TRAIN_COLUMN_MULTIPLIER", TRAIN_COLUMN_MULTIPLIER)
	fmt.Println("TR_ACTION_CARDS", TR_ACTION_CARDS)
	fmt.Println("TR_COMPUTE_SCORE", TR_COMPUTE_SCORE)
	fmt.Println("TR_FINAL_PATH", TR_FINAL_PATH)
	fmt.Printf("TILES: %v\n", TILES)
	nt := NewTiles()
	fmt.Printf("tiles: %v\n", nt.tiles)
	n := nt.PopTile(3)
	if n == -1 {
		fmt.Println("oops.")
	}
	fmt.Println("Tile 1: ", n)
	fmt.Printf("tiles: %v\n", nt.tiles)

	bc := NewBeigeCards()
	l := len(bc.cards)
	fmt.Printf("(len: %v) Beige cards: %v\n", l, bc.cards)
	var top *BeigeCard
	top = bc.PopBeigeCard()
	if top != nil {
		fmt.Println("oops. empty")
	}
	l = len(bc.cards)
	fmt.Printf("top: %v\nlen: %v\nBeige cards: %v\n", *top, l, bc.cards)
	fmt.Printf("movement: %v\n", top.movement)
	rc := NewRedCards()
	fmt.Println("Red cards: ", rc.cards)
	var redtop *RedCard
	redtop = rc.PopRedCard()
	l = len(rc.cards)
	fmt.Printf("redtop: %v\nlen: %v\nRed cards: %v\n", *redtop, l, rc.cards)

	licenseCards := NewLicenseCards()
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
