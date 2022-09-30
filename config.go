package main

import "math/rand"

var TILES = [][]int{
	{},
	{2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4},
	{2, 2, 2, 2, 2, 2, 5, 5, 5, 5, 5, 5},
	{4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6}}

func NewTiles() [][]int {
	tiles := make([][]int, 4)
	for i := 1; i < 4; i++ {
		tiles[i] = make([]int, 12)
		copy(tiles[i], TILES[i])
		rand.Shuffle(len(tiles[i]), func(ii, jj int) { tiles[i][ii], tiles[i][jj] = tiles[i][jj], tiles[i][ii] })
	}
	return tiles
}

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

func NewRedCards() []RedCard {
	cards := make([]RedCard, len(RED_CARDS))
	copy(cards, RED_CARDS)
	rand.Shuffle(len(cards), func(ii, jj int) { cards[ii], cards[jj] = cards[jj], cards[ii] })
	return cards
}

func (redCards *RedCards) PopRedCard() RedCard {
	l := len(*redCards.cards)
	ret := (*redCards.cards)[l-1]
	*redCards.cards = (*redCards.cards)[:l-1]
	return ret
}

type BeigeCard struct {
	nLicenses, movement, oilPrice int
}

type BeigeCards struct {
	cards *[]BeigeCard
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
	cards := make([]BeigeCard, len(BEIGE_CARDS))
	copy(cards, BEIGE_CARDS)
	rand.Shuffle(len(cards), func(ii, jj int) { cards[ii], cards[jj] = cards[jj], cards[ii] })
	beigeCards := new(BeigeCards)
	beigeCards.cards = &cards
	return beigeCards
}

func (beigeCards *BeigeCards) PopBeigeCard() BeigeCard {
	l := len(*beigeCards.cards)
	ret := (*beigeCards.cards)[l-1]
	*beigeCards.cards = (*beigeCards.cards)[:l-1]
	return ret
}

type LicenseCards struct {
	cards *[]int
}

func NewLicenseCards() *LicenseCards {
	cards := make([]int, 78)
	for i := 0; i < 39; i++ {
		cards[i] = 1
	}
	for i := 39; i < 78; i++ {
		cards[i] = 2
	}
	rand.Shuffle(len(cards), func(ii, jj int) { cards[ii], cards[jj] = cards[jj], cards[ii] })
	licenseCards := new(LicenseCards)
	licenseCards.cards = &cards
	return licenseCards
}

func (licenseCards *LicenseCards) PopLicenseCard() int {
	l := len(*licenseCards.cards)
	ret := (*licenseCards.cards)[l-1]
	*licenseCards.cards = (*licenseCards.cards)[:l-1]
	return ret
}

//func (cards *[]) Pop() int {
//	res := (*s)[len(*s)-1]
//	*s = (*s)[:len(*s)-1]
//	return res
//}
