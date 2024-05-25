package main

import (
	"fmt"
	"strconv"
)

type ActionsT struct {
	nlicenses, movement, markers, backwards, oilprice int
}

type PlayerType struct {
	Id             int
	TruckNode      *Node
	TruckHist      []string
	TrainCol       int
	FreeOilRigs    int
	RigsInUse      []*Node
	Cash           int
	StorageTanks   []int // one tank at each company
	Actions        ActionsT
	SingleLicenses LicenseCards
	DoubleLicenses LicenseCards
	NLicenses      int
}

func (p *PlayerType) SetActions(nlicenses, movement, markers, backwards, oilprice int) {
	p.Actions = ActionsT{nlicenses: nlicenses, movement: movement,
		markers: markers, backwards: backwards, oilprice: oilprice}
}

func (p *PlayerType) AdvanceTrain(verbos int) {
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

func SprintPlayer(p *PlayerType) string {
	if p == nil {
		return "nil"
	}
	return strconv.Itoa(p.Id)
}
