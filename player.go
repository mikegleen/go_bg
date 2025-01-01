package main

import (
	"fmt"
	"strconv"
)

type ActionsT struct {
	nlicenses, movement, markers, backwards, oilprice int
}

type Player struct {
	Id             int
	game           *Game
	TruckNode      *Node
	TruckHist      []string
	TrainCol       int // train starts at column zero
	FreeOilRigs    int
	RigsInUse      []*Node
	Cash           int
	StorageTanks   []int // one tank at each company
	Actions        ActionsT
	SingleLicenses int // each SingleLicense is worth one License
	DoubleLicenses int // each DoubleLicense is worth two Licenses
}

func NewPlayer(trucknode *Node, id int) *Player {
	player := new(Player)
	player.Id = id
	player.Cash = INITIAL_CASH
	player.FreeOilRigs = INITIAL_OIL_RIGS
	player.RigsInUse = make([]*Node, 0, 5)
	player.TruckNode = trucknode
	player.TruckHist = make([]string, 0, 20)
	player.StorageTanks = make([]int, NCOMPANIES)
	return player
}

func (p *Player) SetActions(nlicenses, movement, markers, backwards, oilprice int) {
	p.Actions = ActionsT{nlicenses: nlicenses, movement: movement,
		markers: markers, backwards: backwards, oilprice: oilprice}
}

func (p *Player) AdvanceTrain(verbos int) {
	movement := p.Actions.movement
	oldMovement := movement
	oldTrainCol := p.TrainCol
	movement -= p.TruckNode.Distance
	//  cost to move to next column increases as we advance
	for needed := TRAIN_COSTS[p.TrainCol+1]; needed <= movement; {
		movement -= needed
		p.TrainCol += 1
	}
	if verbos >= 2 {
		fmt.Printf("AdvanceTrain: player: %v, Movement: %v -> %v, train col: %v -> %v, truck distance: %v\n",
			p.Id, oldMovement, movement, oldTrainCol, p.TrainCol, p.TruckNode.Distance)
	}
}

func SprintPlayer(p *Player) string {
	if p == nil {
		return "nil"
	}
	return strconv.Itoa(p.Id)
}
