package main

import (
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
	BoardSize    int       // rows * columns
	Board        [][]*Node // rows, columns
	Nodes        []*Node   // 1d array easier to iterate over all nodes
	/*
	   TerrainCh is indexed by terrain type. Zero is illegal. Types 1, 2, and 3
	   refer to flat, hilly, or mountain and correspond to the cost of moving into
	   that square. The values will be ASCII graphics to display on the map.
	*/
	TerrainCh [4]string
}

func NewGraph(rawBoard [][]string, nPlayers int) *Graph {
	threePlayers := nPlayers == 3
	re := regexp.MustCompile(`(\d?)(\.(\d?)([xd]?))?`)
	nRows := len(rawBoard)
	nCols := len(rawBoard[0])
	board := make([][]*Node, nRows)
	graph := new(Graph)
	graph.Nodes = make([]*Node, 0)

	ix := 0
	for i := 0; i < nRows; i++ {
		newRow := make([]*Node, nCols)
		for j := 0; j < nCols; j++ {
			node := NewNode(i, j, false)
			m := re.FindStringSubmatch(rawBoard[i][j])
			if m == nil {
				panic(fmt.Sprintf("Bad format rawBoard[%d][%d] = '%s'", i, j, rawBoard[i][j]))
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
			newRow[j] = node
			graph.Nodes = append(graph.Nodes, node)
			node.Ix = ix
			ix++
		}
		board[i] = newRow
	}
	graph.Board = board
	graph.Rows = nRows
	graph.Columns = nCols
	graph.BoardSize = nRows * nCols
	for _, node := range graph.Nodes {
		node.SetNeighbors(board)
	}
	// yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	graph.TerrainCh[0] = "@"
	//graph.TerrainCh[1] = green("   ")
	//graph.TerrainCh[2] = green("~~ ")
	//graph.TerrainCh[3] = green("^^^")
	graph.TerrainCh[1] = green(" ")
	graph.TerrainCh[2] = green("~")
	graph.TerrainCh[3] = green("^")

	return graph
}

func (g Graph) ResetGraph() {
	for _, node := range g.Nodes {
		node.ResetNode()
	}
}

func (g Graph) PrintBoard(message string, verbos int) {
	if verbos > 1 {
		fmt.Println(message)
	}
	red := color.New(color.FgHiRed).SprintFunc()
	yellow := color.New(color.BgHiYellow).Add(color.FgBlack).SprintFunc()
	var s string
	for n := 0; n < g.Columns; n++ {
		s += fmt.Sprintf("| %03d ", n)
	}
	fmt.Printf("   %s|\n", s)

	for nRow, row := range g.Board {
		fmt.Println("   " + strings.Repeat("|—————", g.Columns) + "|")
		r1 := ""
		r2 := ""
		for _, node := range row {
			terrain := g.TerrainCh[node.Terrain]
			truck := " "
			if node.Truck != nil {
				truck = yellow(strconv.Itoa(node.Truck.Id))
			}
			r1 += terrain + truck + " " + node.PrDist() + "|"
			goal := " "
			if node.Goal > 0 {
				goal = red(strconv.Itoa(node.Goal))
			}
			r2 += node.PrWells() + node.FromArrow() + goal + "|"
		}
		//for _, node := range row {
		//}
		fmt.Printf(" %02d|%s\n", nRow, r1)
		fmt.Printf("   |%s\n", r2)
	}

	fmt.Println("   " + strings.Repeat("|—————", g.Columns) + "|")
}
