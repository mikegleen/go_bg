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
	fmt.Printf("TILES: %v\n", TILES)
	nt := NewTiles()
	fmt.Printf("tiles: %v\n", *nt.tiles)
	n := nt.PopTile(3)
	if n == -1 {
		fmt.Println("oops.")
	}
	fmt.Println("Tile 1: ", n)
	fmt.Printf("tiles: %v\n", *nt.tiles)

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
	fmt.Println("Red cards: ", *rc.cards)
	var redtop *RedCard
	redtop = rc.PopRedCard()
	l = len(*rc.cards)
	fmt.Printf("redtop: %v\nlen: %v\nRed cards: %v\n", *redtop, l, *rc.cards)

	licenseCards := NewLicenseCards()
	fmt.Printf("License cards: %v\n", licenseCards.cards)
	fmt.Println("Pop: ", licenseCards.PopLicenseCard())
	fmt.Println("Pop: ", licenseCards.PopLicenseCard())
	fmt.Println("Pop: ", licenseCards.PopLicenseCard())
}
