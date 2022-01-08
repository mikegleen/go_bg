package main

import (
	"fmt"
	"math"
	"strings"
)

// yellow := color.New(color.FgYellow).SprintFunc()
// red := color.New(color.FgRed).SprintFunc()
// green := color.New(color.FgGreen).SprintFunc()
const leftwardsArrow string = "\u2190"
const upwardsArrow string = "\u2191"
const rightwardsArrow string = "\u2192"
const downwardsArrow string = "\u2193"

// terrainCh := []string{"@", green("-  "), green("~~ "), green("^^^")}

type Node struct {
	Row     int
	Col     int
	Id      string
	Terrain int
	// wells: int in 0..3: the number of wells on the square. If non-zero
	// the square is covered with a tile at the start of the game. Wells
	// are assigned when the Graph is instantiated.
	Wells int
	// oil_reserve: the number on the bottom side of the tile covering a
	// square with well(s) or the number of plastic markers under a derrick.
	// If wells == 2 then we cannot peek at the oil reserve in deciding
	// whether to drill.
	OilReserve int
	Exhausted  bool
	Goal       int
	Derrick    bool
	Truck      *Player
	Adjacent   []*Node
	Cell       string

	// Fields set by dijkstra
	Distance int
	Previous *Node
}

func NewNode(row, col int, derrick bool) *Node {
	node := new(Node)
	node.Row = row
	node.Col = col
	node.Id = fmt.Sprintf("%d,%d", row, col)
	node.Derrick = derrick
	node.Distance = math.MaxInt
	return node
}

func SprintPreviousNode(n *Node) string {
	ret := ""
	if n.Previous != nil {
		ret = n.Previous.Id
	}
	return ret
}

func SprintNode(n *Node) string {

	s := fmt.Sprintf("%s t: %d, w: %d ", n.Id, n.Terrain, n.Wells)
	s += fmt.Sprintf("ex=%t, goal=%d, derrick=%t, truck=%s, ", n.Exhausted, n.Goal, n.Derrick, SprintPlayer(n.Truck))
	s += fmt.Sprintf("previous={%s}, ", SprintPreviousNode(n))
	d := fmt.Sprintf("%d", n.Distance)
	if n.Distance == math.MaxInt {
		d = "∞"
	}
	s += fmt.Sprintf("dist: %s, ", d)
	adjacents := "{"
	for _, adj := range n.Adjacent {
		adjacents += adj.Id + ","
	}
	adjacents = strings.TrimSuffix(adjacents, ",")
	adjacents += "}"
	s += fmt.Sprintf("adjacent: %s", adjacents)
	return s
}

func (node *Node) PrDist() string {
	dist := node.Distance
	var ret string
	if dist != math.MaxInt {
		ret = fmt.Sprintf("%2d", dist)
	} else {
		ret = "  "
	}
	return ret
}

func (node *Node) PrWells() string {
	if node.Exhausted {
		return "X  "
	}
	well := "w"
	if node.Derrick {
		well = "D"
	}
	return strings.Repeat(well, node.Wells) + strings.Repeat(" ", 3-node.Wells)
}

func (node *Node) FromArrow() string {
	previous := node.Previous
	if previous == nil {
		return " "
	}
	if node.Row == previous.Row {
		if node.Row > previous.Row {
			return leftwardsArrow
		}
		return rightwardsArrow
	}
	if node.Row > previous.Row {
		return upwardsArrow
	}
	return downwardsArrow
}

func (n *Node) set1Neighbor(board [][]*Node, nrow, ncol int) {
	neighbor := board[nrow][ncol]
	n.Adjacent = append(n.Adjacent, neighbor)
	if n.Wells > 0 && !n.Derrick {
		if neighbor.Wells == 0 {
			neighbor.Goal += 1
		}
	}
}
func (n *Node) SetNeighbors(board [][]*Node) {

	lastrow := len(board) - 1
	lastcol := len(board[0]) - 1
	if n.Row > 0 {
		n.set1Neighbor(board, n.Row-1, n.Col)
	}
	if n.Col > 0 {
		n.set1Neighbor(board, n.Row, n.Col-1)
	}
	if n.Row < lastrow {
		n.set1Neighbor(board, n.Row+1, n.Col)
	}
	if n.Col < lastcol {
		n.set1Neighbor(board, n.Row, n.Col+1)
	}
}
