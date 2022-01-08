package main

import (
	// "fmt"
	// "math"
	"fmt"
	"regexp"
	"strconv"
	// "github.com/fatih/color"
)

type Graph struct {
	ThreePlayers bool
	Rows         int
	Columns      int
	Board        [][]*Node
	Nodes        []*Node
}

func NewGraph(rawboard [][]string, nplayers int) *Graph {
	threePlayers := nplayers == 3
	re := regexp.MustCompile(`(\d?)(\.(\d?)([xd]?))?`)
	nrows := len(rawboard)
	ncols := len(rawboard[0])
	board := make([][]*Node, nrows)
	graph := new(Graph)
	graph.Nodes = make([]*Node, 0)

	for i := 0; i < nrows; i++ {
		newrow := make([]*Node, ncols)
		for j := 0; j < ncols; j++ {
			node := NewNode(i, j, false)
			m := re.FindStringSubmatch(rawboard[i][j])
			if m == nil {
				panic(fmt.Sprintf("Bad format rawboard[%d][%d] = '%s'", i, j, rawboard[i][j]))
			}
			node.Terrain = 1
			if m[1] != "" {
				node.Terrain, _ = strconv.Atoi(m[1])
			}
			node.Wells, _ = strconv.Atoi(m[3])
			if m[4] == "x" && threePlayers {
				node.Wells = 0
			}
			// For testing, force a derrick
			if m[4] == "d" {
				node.Derrick = true
			}
			newrow[j] = node
			graph.Nodes = append(graph.Nodes, node)
		}
		board[i] = newrow
	}
	graph.Board = board
	graph.Rows = nrows
	graph.Columns = ncols
	for _, node := range graph.Nodes {
		node.SetNeighbors(board)
	}
	return graph
}

func (g Graph) PrintBoard() {

}
