package game

import (
	"math"
	"math/rand"
)

// import dij "giganten/dijkstra"
// --------------------- Tiles -----------------------
// TILES: This dict defines the cardboard pieces that cover the squares
//        containing wells.
//        The key is the number of oil wells, the value is a list of the
//        number of oil markers to be allocated when a derrick is built.
//        This list will be copied, shuffled and allocated to the nodes
//        according to the number of wells on that node. When a player builds
//        a derrick, the number of oil markers corresponding to the tile's
//        value are allocated to the node.

type Tiles struct {
	tiles [][]int
}

var TILES = [][]int{
	{},
	{2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4},
	{2, 2, 2, 2, 2, 2, 5, 5, 5, 5, 5, 5},
	{4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6}}

func NewTiles() *Tiles {
	sTiles := new(Tiles)
	sTiles.tiles = make([][]int, 4)
	for i := 1; i < 4; i++ {
		sTiles.tiles[i] = make([]int, 12)
		copy(sTiles.tiles[i], TILES[i])
		rand.Shuffle(len(sTiles.tiles[i]), func(ii, jj int) { sTiles.tiles[i][ii], sTiles.tiles[i][jj] = sTiles.tiles[i][jj], sTiles.tiles[i][ii] })
	}
	return sTiles
}

func (tiles *Tiles) PopTile(nWells int) int {
	l := len(tiles.tiles[nWells])
	if l < 1 {
		return -1
	}
	ret := tiles.tiles[nWells][l-1]
	tiles.tiles[nWells] = tiles.tiles[nWells][:l-1]
	return ret
}

// TRAIN_COSTS:
// How much it costs to move a player's train one column. Appending maxsize to the
// end prevents train movement past the board.

var TRAIN_COSTS = [...]int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, math.MaxInt}

var INITIAL_CASH = 15_000
var INITIAL_PRICE = 5000
var INITIAL_OIL_RIGS = 5
var INITIAL_OIL_MARKERS = 60

// TRUCK_INIT_ROWS the rows in the first column to place the trucks at the start of the game
var TRUCK_INIT_ROWS = [...]int{1, 5, 7, 9}

var NCOMPANIES = 3

// STORAGE_TANK_LIMIT For Action 8, Storage Tank Limitations, any crude markers in storage tanks
// above the limit must be sold to the bank for $1000 each.
var STORAGE_TANK_LIMIT = 2

var FORCED_SALE_PRICE = 1000

// BUILDING_COST Cost to build a derrick. The index is the number of wells on a site.
var BUILDING_COST = [...]int{0, 4000, 6000, 8000}

// TRANSPORT_COST Cost to transport oil on the black train or on an opponent's train
var TRANSPORT_COST = 3000

// GAME_END_RIG_PRICE The price paid for oilrigs at the end of the game
var GAME_END_RIG_PRICE = [...]int{5000, 4000, 3000, 2000}
var GAME_END_MARKER_PRICE = 1000

// Heuristics for choosing a truck destination

var GOAL_MULTIPLIER = 2
var TRUCK_COLUMN_MULTIPLIER = 1
var PREV_GOAL_MULTIPLER = 1
var TRAIN_COLUMN_MULTIPLIER = 1

// Define trace levels

var TR_ACTION_CARDS = 2
var TR_COMPUTE_SCORE = 2
var TR_FINAL_PATH = 2

// --------------------- Red Action Cards -----------------------

type RedCard struct {
	blackLoco, // Advance the black locomotive.
	nLicenses, // Number of green license cards to deal
	movement, // Advance the player's loco or truck.
	markers, // Take one crude oil marker.
	backwards int // Move all opponents locos backwards.
}

type RedCards struct {
	cards []RedCard
}

var RED_CARDS = []RedCard{
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
	rc := new(RedCards)
	rc.cards = make([]RedCard, len(RED_CARDS))
	copy(rc.cards, RED_CARDS)
	rand.Shuffle(len(rc.cards), func(ii, jj int) { rc.cards[ii], rc.cards[jj] = rc.cards[jj], rc.cards[ii] })
	return rc
}

func (redCards *RedCards) PopRedCard() *RedCard {
	l := len(redCards.cards)
	if l <= 0 {
		return nil
	}
	ret := &redCards.cards[l-1]
	redCards.cards = redCards.cards[:l-1]
	return ret
}

// --------------------- Beige Action Cards -----------------------

type BeigeCard struct {
	nLicenses, // Number of green license cards to deal
	movement, // Advance the player's loco or truck.
	oilPrice int // Change any oil price up or down.
}

type BeigeCards struct {
	cards []BeigeCard
}

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
	bc := new(BeigeCards)
	bc.cards = make([]BeigeCard, len(BEIGE_CARDS))
	copy(bc.cards, BEIGE_CARDS)
	rand.Shuffle(len(bc.cards), func(ii, jj int) { bc.cards[ii], bc.cards[jj] = bc.cards[jj], bc.cards[ii] })
	return bc
}

func (beigeCards *BeigeCards) PopBeigeCard() *BeigeCard {
	l := len(beigeCards.cards)
	if l <= 0 {
		return nil
	}
	ret := &(beigeCards.cards)[l-1]
	beigeCards.cards = beigeCards.cards[:l-1]
	return ret
}

// --------------------- License Cards -----------------------

type LicenseCards struct {
	cards []int
}

func NewLicenseCards() *LicenseCards {
	lc := new(LicenseCards)
	lc.cards = make([]int, 78)
	for i := 0; i < 39; i++ {
		lc.cards[i] = 1
	}
	for i := 39; i < 78; i++ {
		lc.cards[i] = 2
	}
	rand.Shuffle(len(lc.cards), func(ii, jj int) { lc.cards[ii], lc.cards[jj] = lc.cards[jj], lc.cards[ii] })
	return lc
}

func (licenseCards *LicenseCards) PopLicenseCard() int {
	l := len(licenseCards.cards)
	if l < 1 {
		return -1 // empty
	}
	ret := (licenseCards.cards)[l-1]
	licenseCards.cards = (licenseCards.cards)[:l-1]
	return ret
}
