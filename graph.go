package main

import (
	// "fmt"
	// "math"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Graph struct {
	ThreePlayers bool
	Rows         int
	Columns      int
	BoardSize    int
	Board        [][]*Node
	Nodes        []*Node
	/*
	   TerrainCh is indexed by terrain type. Zero is illegal. Types 1, 2, and 3
	   refer to flat, hilly, or mountain and correspond to the cost of moving into
	   that square.
	*/
	TerrainCh [4]string
}

func NewGraph(rawboard [][]string, nplayers int) *Graph {
	threePlayers := nplayers == 3
	re := regexp.MustCompile(`(\d?)(\.(\d?)([xd]?))?`)
	nrows := len(rawboard)
	ncols := len(rawboard[0])
	board := make([][]*Node, nrows)
	graph := new(Graph)
	graph.Nodes = make([]*Node, 0)

	ix := 0
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
			node.Ix = ix
			ix++
		}
		board[i] = newrow
	}
	graph.Board = board
	graph.Rows = nrows
	graph.Columns = ncols
	graph.BoardSize = nrows * ncols
	for _, node := range graph.Nodes {
		node.SetNeighbors(board)
	}
	// yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	graph.TerrainCh[0] = "@"
	graph.TerrainCh[1] = green("-  ")
	graph.TerrainCh[2] = green("~~ ")
	graph.TerrainCh[3] = green("^^^")

	return graph
}

func (g Graph) ResetGraph() {
	for _, node := range g.Nodes {
		node.ResetNode()
	}
}

func (g Graph) PrintBoard() {
	red := color.New(color.FgRed).SprintFunc()
	var s string
	for n := 0; n < g.Columns; n++ {
		s += fmt.Sprintf("| %03d ", n)
	}
	fmt.Printf("   %s|\n", s)

	for nrow, row := range g.Board {
		fmt.Println("   " + strings.Repeat("|—————", g.Columns) + "|")
		r1 := ""
		for _, node := range row {
			r1 += g.TerrainCh[node.Terrain] + node.PrDist() + "|"
		}
		fmt.Printf(" %02d|%s\n", nrow, r1)
		r2 := ""
		for _, node := range row {
			goal := " "
			if node.Goal > 0 {
				goal = red(strconv.Itoa(node.Goal))
			}
			r2 += node.PrWells() + node.FromArrow() + goal + "|"
		}
		fmt.Printf("   |%s\n", r2)
	}

	fmt.Println("   " + strings.Repeat("|—————", g.Columns) + "|")
}
