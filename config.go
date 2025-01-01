package main

import (
	"math"
	"math/rand/v2"
)

// import dij "giganten/dijkstra"
// --------------------- tiles -----------------------
//
// Comments from Python version:
// TILES: This dict defines the cardboard pieces that cover the squares
//        containing wells.
//        The key is the number of oil wells, the value is a list of the
//        number of oil Markers to be allocated when a derrick is built.
//        This list will be copied, shuffled and allocated to the nodes
//        according to the number of wells on that node. When a player builds
//        a derrick, the number of oil Markers corresponding to the tile's
//        value are allocated to the node.

type Tiles [][]int

var TILES = Tiles{
	{},
	{2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4},
	{2, 2, 2, 2, 2, 2, 5, 5, 5, 5, 5, 5},
	{4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6}}

func NewTiles() *Tiles {
	sTiles := make(Tiles, 4)
	for i := 1; i < 4; i++ {
		sTiles[i] = make([]int, len(TILES[i]))
		copy(sTiles[i], TILES[i])
		rand.Shuffle(len(sTiles[i]),
			func(ii, jj int) {
				sTiles[i][ii], sTiles[i][jj] =
					sTiles[i][jj], sTiles[i][ii]
			})
	}
	return &sTiles
}

func (tiles *Tiles) PopTile(nWells int) int {
	//fmt.Println("PopTile:", tiles)
	//fmt.Println("PopTile:", *tiles)
	t := (*tiles)[nWells]
	l := len(t)
	if l < 1 {
		panic("Ran out of tiles.")
	}
	ret := (*tiles)[nWells][l-1]
	(*tiles)[nWells] = (*tiles)[nWells][:l-1]
	return ret
}

// TRAIN_COSTS:
// How much it costs to move a player's train one column. Appending MaxInt to the
// end prevents train movement past the board.

var TRAIN_COSTS = [...]int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, math.MaxInt}

const INITIAL_CASH = 15_000
const INITIAL_PRICE = 5000
const INITIAL_OIL_RIGS = 5
const INITIAL_OIL_MARKERS = 60
const INITIAL_SINGLE_LICENSES = 39
const INITIAL_DOUBLE_LICENSES = 39

// TRUCK_INIT_ROWS the rows in the first column to place the trucks at the start of the game
// indexed by player number 0...3
var TRUCK_INIT_ROWS = [...]int{1, 5, 7, 9}

const NCOMPANIES = 3

// STORAGE_TANK_LIMIT For Action 8, Storage Tank Limitations, any crude Markers in storage tanks
// above the limit must be sold to the bank for $1000 each.
const STORAGE_TANK_LIMIT = 2

const FORCED_SALE_PRICE = 1000

// BUILDING_COST Cost to build a derrick. The index is the number of wells on a site.
var BUILDING_COST = [...]int{0, 4000, 6000, 8000}

// TRANSPORT_COST Cost to transport oil on the black train or on an opponent's train
const TRANSPORT_COST = 3000

// END_OF_GAME_RIG_PRICE The price paid for oilrigs at the end of the game
var END_OF_GAME_RIG_PRICE = [...]int{5000, 4000, 3000, 2000}

// END_OF_GAME_MARKER_PRICE The price paid for each oil marker in a player's rig or storage tank
const END_OF_GAME_MARKER_PRICE = 1000

// TOTAL_LICENSES The total number of licenses contained in single and double license cards
const TOTAL_LICENSES = 39 + 39*2

// Heuristics for choosing a truck destination

const GOAL_MULTIPLIER = 2
const TRUCK_COLUMN_MULTIPLIER = 1
const PREV_GOAL_MULTIPLER = 1
const TRAIN_COLUMN_MULTIPLIER = 1

// Define trace levels

const TRACE_ACTION_CARDS = 2
const TRACE_COMPUTE_SCORE = 2
const TRACE_FINAL_PATH = 2

// --------------------- Combined Action Cards -----------------------

type ActionCard struct {
	BlackLoco, // Advance the black locomotive.
	NLicenseCards, // Number of green license Cards to deal
	Movement, // Advance the player's loco or truck.
	Markers, // Take one crude oil marker.
	Backwards int // Move all opponents locos Backwards.
	OilPrice int // Change any oil price up or down.
}

// --------------------- Red Action Cards -----------------------

type RedCard struct {
	BlackLoco, // Advance the black locomotive.
	NLicenseCards, // Number of green license Cards to deal
	Movement, // Advance the player's loco or truck.
	Markers, // Take one crude oil marker.
	Backwards int // Move all opponents locos Backwards.
}

type RedCards []RedCard

var RED_CARDS = RedCards{
	{3, 3, 4, 1, 0},
	{2, 4, 8, 0, 2},
	{2, 3, 6, 0, 3},
	{2, 2, 6, 1, 0},
	{1, 4, 2, 1, 0},
	{3, 2, 4, 0, 5},
	{2, 2, 6, 0, 4},
	{3, 3, 4, 1, 0},
	{1, 2, 8, 0, 3},
	{2, 2, 6, 1, 0},
	{3, 2, 6, 0, 4},
	{2, 3, 2, 0, 5},
}

func NewRedCards() *RedCards {
	rc := make(RedCards, len(RED_CARDS))
	copy(rc, RED_CARDS)
	//fmt.Println("len(rc)=", len(rc))
	rand.Shuffle(len(rc), func(ii, jj int) { rc[ii], rc[jj] = rc[jj], rc[ii] })
	return &rc
}

func (redCards *RedCards) PopRedCard() *RedCard {
	l := len(*redCards)
	if l <= 0 {
		return nil
	}
	ret := &(*redCards)[l-1]
	*redCards = (*redCards)[:l-1]
	return ret
}

// --------------------- Beige Action Cards -----------------------

type BeigeCard struct {
	NLicenseCards, // Number of green license Cards to deal
	Movement, // Advance the player's loco or truck.
	OilPrice int // Change any oil price up or down.
}

type BeigeCards []BeigeCard

var BEIGE_CARDS = []BeigeCard{
	{2, 10, 2},
	{7, 4, 0},
	{5, 8, 0},
	{5, 5, 3},
	{4, 10, 0},
	{5, 8, 0},
	{4, 5, 4},
	{4, 10, 0},
	{4, 6, 3},
	{5, 8, 0},
	{5, 8, 0},
	{2, 10, 2},
	{3, 8, 3},
	{6, 6, 0},
	{3, 12, 0},
	{5, 8, 0},
	{4, 10, 0},
	{3, 12, 0},
	{6, 6, 0},
	{5, 8, 0},
	{7, 4, 0},
	{4, 5, 4},
	{6, 6, 0},
	{8, 3, 0},
	{4, 5, 4},
	{6, 6, 0},
	{3, 12, 0},
	{4, 10, 0},
	{6, 4, 2},
	{4, 10, 0}}

func NewBeigeCards() *BeigeCards {
	bc := make(BeigeCards, len(BEIGE_CARDS))
	copy(bc, BEIGE_CARDS)
	rand.Shuffle(len(bc), func(ii, jj int) { bc[ii], bc[jj] = bc[jj], bc[ii] })
	return &bc
}

func (beigeCards *BeigeCards) PopBeigeCard() *BeigeCard {
	l := len(*beigeCards)
	if l <= 0 {
		return nil
	}
	ret := &(*beigeCards)[l-1]
	*beigeCards = (*beigeCards)[:l-1]
	return ret
}

// --------------------- License Cards -----------------------

func NewLicenseCards(numSingle, numDouble int) *[]int {
	totcards := numSingle + numDouble
	lc := make([]int, totcards)
	for i := 0; i < numSingle; i++ {
		lc[i] = 1
	}
	for i := numSingle; i < totcards; i++ {
		lc[i] = 2
	}
	rand.Shuffle(len(lc), func(ii, jj int) { lc[ii], lc[jj] = lc[jj], lc[ii] })
	return &lc
}

func PopLicenseCard(licenseCards *[]int) int {
	l := len(*licenseCards)
	if l < 1 {
		return 0 // empty
	}
	ret := (*licenseCards)[l-1]
	*licenseCards = (*licenseCards)[:l-1]
	return ret
}

func DrawLicenseCard(g *Game) int {

	licenses := PopLicenseCard(g.licensecards)
	if licenses > 0 {
		return licenses
	}
	if len(*g.licensediscards) == 0 {
		panic("Exhausted license cards.")
	}
	g.licensecards = g.licensediscards
	*g.licensediscards = make([]int, 0)

	lc := *g.licensecards
	rand.Shuffle(len(lc), func(ii, jj int) { lc[ii], lc[jj] = lc[jj], lc[ii] })
	licenses = PopLicenseCard(g.licensecards)
	if licenses <= 0 {
		panic("Gee, I thought we had licenses....")
	}
	return licenses
}
