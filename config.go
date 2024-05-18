package game

import (
	"math"
	"math/rand"
)

// import dij "giganten/dijkstra"
// --------------------- TileType -----------------------
// TILES: This dict defines the cardboard pieces that cover the squares
//        containing wells.
//        The key is the number of oil wells, the value is a list of the
//        number of oil Markers to be allocated when a derrick is built.
//        This list will be copied, shuffled and allocated to the nodes
//        according to the number of wells on that node. When a player builds
//        a derrick, the number of oil Markers corresponding to the tile's
//        value are allocated to the node.

type TileType struct {
	Tiles [][]int
}

var TILES = [][]int{
	{},
	{2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4},
	{2, 2, 2, 2, 2, 2, 5, 5, 5, 5, 5, 5},
	{4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6}}

func NewTiles() *TileType {
	sTiles := new(TileType)
	sTiles.Tiles = make([][]int, 4)
	for i := 1; i < 4; i++ {
		sTiles.Tiles[i] = make([]int, 12)
		copy(sTiles.Tiles[i], TILES[i])
		rand.Shuffle(len(sTiles.Tiles[i]),
			func(ii, jj int) { sTiles.Tiles[i][ii], sTiles.Tiles[i][jj] = sTiles.Tiles[i][jj], sTiles.Tiles[i][ii] })
	}
	return sTiles
}

func (tiles *TileType) PopTile(nWells int) int {
	l := len(tiles.Tiles[nWells])
	if l < 1 {
		return -1
	}
	ret := tiles.Tiles[nWells][l-1]
	tiles.Tiles[nWells] = tiles.Tiles[nWells][:l-1]
	return ret
}

// TRAIN_COSTS:
// How much it costs to move a player's train one column. Appending maxsize to the
// end prevents train Movement past the board.

var TRAIN_COSTS = [...]int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, math.MaxInt}

var INITIAL_CASH = 15_000
var INITIAL_PRICE = 5000
var INITIAL_OIL_RIGS = 5
var INITIAL_OIL_MARKERS = 60

// TRUCK_INIT_ROWS the rows in the first column to place the trucks at the start of the game
var TRUCK_INIT_ROWS = [...]int{1, 5, 7, 9}

var NCOMPANIES = 3

// STORAGE_TANK_LIMIT For Action 8, Storage Tank Limitations, any crude Markers in storage tanks
// above the limit must be sold to the bank for $1000 each.
var STORAGE_TANK_LIMIT = 2

var FORCED_SALE_PRICE = 1000

// BUILDING_COST Cost to build a derrick. The index is the number of wells on a site.
var BUILDING_COST = [...]int{0, 4000, 6000, 8000}

// TRANSPORT_COST Cost to transport oil on the black train or on an opponent's train
var TRANSPORT_COST = 3000

// END_OF_GAME_RIG_PRICE The price paid for oilrigs at the end of the game
var END_OF_GAME_RIG_PRICE = [...]int{5000, 4000, 3000, 2000}
var END_OF_GAME_MARKER_PRICE = 1000

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
	BlackLoco, // Advance the black locomotive.
	NLicenses, // Number of green license Cards to deal
	Movement, // Advance the player's loco or truck.
	Markers, // Take one crude oil marker.
	Backwards int // Move all opponents locos Backwards.
}

type RedCards struct {
	Cards []RedCard
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
	rc.Cards = make([]RedCard, len(RED_CARDS))
	copy(rc.Cards, RED_CARDS)
	rand.Shuffle(len(rc.Cards), func(ii, jj int) { rc.Cards[ii], rc.Cards[jj] = rc.Cards[jj], rc.Cards[ii] })
	return rc
}

func (redCards *RedCards) PopRedCard() *RedCard {
	l := len(redCards.Cards)
	if l <= 0 {
		return nil
	}
	ret := &redCards.Cards[l-1]
	redCards.Cards = redCards.Cards[:l-1]
	return ret
}

// --------------------- Beige Action Cards -----------------------

type BeigeCard struct {
	NLicenses, // Number of green license Cards to deal
	Movement, // Advance the player's loco or truck.
	OilPrice int // Change any oil price up or down.
}

type BeigeCards struct {
	Cards []BeigeCard
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
	bc.Cards = make([]BeigeCard, len(BEIGE_CARDS))
	copy(bc.Cards, BEIGE_CARDS)
	rand.Shuffle(len(bc.Cards), func(ii, jj int) { bc.Cards[ii], bc.Cards[jj] = bc.Cards[jj], bc.Cards[ii] })
	return bc
}

func (beigeCards *BeigeCards) PopBeigeCard() *BeigeCard {
	l := len(beigeCards.Cards)
	if l <= 0 {
		return nil
	}
	ret := &(beigeCards.Cards)[l-1]
	beigeCards.Cards = beigeCards.Cards[:l-1]
	return ret
}

// --------------------- License Cards -----------------------

type LicenseCards struct {
	Cards []int
}

func NewLicenseCards() *LicenseCards {
	lc := new(LicenseCards)
	lc.Cards = make([]int, 78)
	for i := 0; i < 39; i++ {
		lc.Cards[i] = 1
	}
	for i := 39; i < 78; i++ {
		lc.Cards[i] = 2
	}
	rand.Shuffle(len(lc.Cards), func(ii, jj int) { lc.Cards[ii], lc.Cards[jj] = lc.Cards[jj], lc.Cards[ii] })
	return lc
}

func (licenseCards *LicenseCards) PopLicenseCard() int {
	l := len(licenseCards.Cards)
	if l < 1 {
		return -1 // empty
	}
	ret := (licenseCards.Cards)[l-1]
	licenseCards.Cards = (licenseCards.Cards)[:l-1]
	return ret
}
