package main

import (
	"math/rand"
)

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
	tiles *[][]int
}

var TILES = [][]int{
	{},
	{2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4},
	{2, 2, 2, 2, 2, 2, 5, 5, 5, 5, 5, 5},
	{4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6}}

func NewTiles() *Tiles {
	tiles := make([][]int, 4)
	for i := 1; i < 4; i++ {
		tiles[i] = make([]int, 12)
		copy(tiles[i], TILES[i])
		rand.Shuffle(len(tiles[i]), func(ii, jj int) { tiles[i][ii], tiles[i][jj] = tiles[i][jj], tiles[i][ii] })
	}
	sTiles := new(Tiles)
	sTiles.tiles = &tiles
	return sTiles
}
func (tiles *Tiles) PopTile(nWells int) int {
	l := len((*tiles.tiles)[nWells])
	if l < 1 {
		return -1
	}
	ret := (*tiles.tiles)[nWells][l-1]
	(*tiles.tiles)[nWells] = (*tiles.tiles)[nWells][:l-1]
	return ret
}

// --------------------- Red Action Cards -----------------------

type RedCard struct {
	blackLoco, nLicenses, movement, markers, backwards int
}

type RedCards struct {
	cards *[]RedCard
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
	cards := make([]RedCard, len(RED_CARDS))
	copy(cards, RED_CARDS)
	rand.Shuffle(len(cards), func(ii, jj int) { cards[ii], cards[jj] = cards[jj], cards[ii] })
	redCards := new(RedCards)
	redCards.cards = &cards
	return redCards
}

func (redCards *RedCards) PopRedCard() *RedCard {
	l := len(*redCards.cards)
	if l <= 0 {
		return nil
	}
	ret := &(*redCards.cards)[l-1]
	*redCards.cards = (*redCards.cards)[:l-1]
	return ret
}

// --------------------- Beige Action Cards -----------------------

type BeigeCard struct {
	nLicenses, movement, oilPrice int
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
	licenseCards := new(LicenseCards)
	licenseCards.cards = make([]int, 78)
	for i := 0; i < 39; i++ {
		licenseCards.cards[i] = 1
	}
	for i := 39; i < 78; i++ {
		licenseCards.cards[i] = 2
	}
	rand.Shuffle(len(licenseCards.cards), func(ii, jj int) {
		licenseCards.cards[ii], licenseCards.cards[jj] =
			licenseCards.cards[jj], licenseCards.cards[ii]
	})
	return licenseCards
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
