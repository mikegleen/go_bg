package main

import (
	"fmt"
	"math"
	"strings"
)

const leftwardsArrow string = "\u2190"
const upwardsArrow string = "\u2191"
const rightwardsArrow string = "\u2192"
const downwardsArrow string = "\u2193"

// terrainCh := []string{"@", green("-  "), green("~~ "), green("^^^")}

type Node struct {
	Row     int
	Col     int
	Ix      int // set when the Graph is created
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
	Goal       int // count of adjacent nodes with unbuilt wells
	Derrick    bool
	Truck      *Player // set when a truck moves here
	Adjacent   []*Node // will be populated by SetNeighbors
	Cell       string  // this node's string from rawboard, for debugging

	// Fields set by dijkstra
	Distance int
	PqIndex  int
	Previous *Node
}

type NodeList []*Node // implements Sort.Interface

func (nl NodeList) Len() int           { return len(nl) }
func (nl NodeList) Less(i, j int) bool { return nl[i].Distance > nl[j].Distance }
func (nl NodeList) Swap(i, j int)      { nl[i], nl[j] = nl[j], nl[i] }

func NewNode(row, col int, derrick bool) *Node {
	node := new(Node)
	node.Row = row
	node.Col = col
	node.Id = fmt.Sprintf("<%d,%d>", row, col)
	node.Derrick = derrick
	node.Distance = math.MaxInt
	return node
}

func (node *Node) ResetNode() {
	node.Distance = math.MaxInt
	node.Previous = nil
}

func SprintPreviousNode(n *Node) string {
	ret := ""
	if n.Previous != nil {
		ret = n.Previous.Id
	}
	return ret
}

func (n *Node) SprintNode() string {
	/*
		Implement the function of the Python __str__ method
	*/

	tf := func(b bool) string {
		t := "F"
		if b {
			t = "T"
		}
		return t
	}

	s := fmt.Sprintf("%s t: %d, w: %d ", n.Id, n.Terrain, n.Wells)
	s += fmt.Sprintf("ex=%s, goal=%d, derrick=%s, truck=%s, ", tf(n.Exhausted), n.Goal, tf(n.Derrick), SprintPlayer(n.Truck))
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
		if node.Col > previous.Col {
			return leftwardsArrow
		}
		return rightwardsArrow
	}
	if node.Row > previous.Row {
		return upwardsArrow
	}
	return downwardsArrow
}

func (n *Node) SetNeighbors(board [][]*Node) {
	/*
		This is called by NewGraph.

		:param board: the board from a Graph instance

		adjacent contains a list of nodes next to this node.
		A node can have up to 4 adjacent, reduced if it is on an edge.

		:return: None. The adjacent list for this node is set.
	*/
	set1Neighbor := func(nrow, ncol int) {
		neighbor := board[nrow][ncol]
		n.Adjacent = append(n.Adjacent, neighbor)
		if n.Wells > 0 && !n.Derrick {
			if neighbor.Wells == 0 {
				neighbor.Goal += 1
			}
		}
	}

	lastrow := len(board) - 1
	lastcol := len(board[0]) - 1
	if n.Row > 0 {
		set1Neighbor(n.Row-1, n.Col)
	}
	if n.Col > 0 {
		set1Neighbor(n.Row, n.Col-1)
	}
	if n.Row < lastrow {
		set1Neighbor(n.Row+1, n.Col)
	}
	if n.Col < lastcol {
		set1Neighbor(n.Row, n.Col+1)
	}
}

func (n *Node) AddDerrick() {
	if n.Derrick {
		panic("node " + n.Id + " already has derrick.")
	}
	n.Derrick = true
	for _, node := range n.Adjacent {
		node.Goal--
		if node.Goal <= 0 {
			panic("node " + node.Id + " Goal <= 0, adding derrick to node " + n.Id)
		}
	}
}

func (n *Node) RemoveDerrick() {
	if !n.Derrick {
		panic("node " + n.Id + " attempting to remove nonexisting derrick.")
	}
	n.Derrick = false
	n.Exhausted = true
}
