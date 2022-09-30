package main

import (
	"container/heap"
	"fmt"
	"math"

	// "sort"
	"strconv"
)

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Node

// There are mixed receiver types because heap inherits sort.

//goland:noinspection GoMixedReceiverTypes
func (pq PriorityQueue) Len() int { return len(pq) }

//goland:noinspection GoMixedReceiverTypes
func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	// mlg: modified to give lowest
	// fmt.Print(i, j, pq)
	return pq[i].Distance < pq[j].Distance
}

//goland:noinspection GoMixedReceiverTypes
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].PqIndex = i
	pq[j].PqIndex = j
}

//goland:noinspection GoMixedReceiverTypes
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.PqIndex = n
	*pq = append(*pq, item)
}

//goland:noinspection GoMixedReceiverTypes
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil    // avoid memory leak
	item.PqIndex = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
//
//goland:noinspection GoMixedReceiverTypes
func (pq *PriorityQueue) update(item *Node, Distance int) {
	item.Distance = Distance
	heap.Fix(pq, item.PqIndex)
}

// PqMain This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func PqMain() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Node{
			Distance: priority,
			Id:       value,
			PqIndex:  i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Node{
		Id:       "orange",
		Distance: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Node)
		fmt.Printf("%.2d:%s ", item.Distance, item.Id)
	}
	fmt.Println()
}

func iDist(distance int) (ret string) {
	if distance == math.MaxInt {
		ret = "∞"
	} else {
		ret = strconv.Itoa(distance)
	}
	return
}

/*
func selectionSort(items []*Node) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1].Distance > items[j].Distance {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
*/

func insertionSort(A []*Node) {
	for i := 1; i < len(A); i++ {
		key := A[i]
		j := i - 1
		for j > -1 && A[j].Distance > key.Distance {
			A[j+1] = A[j]
			j -= 1
		}
		A[j+1] = key
	}
}

// func dijkstra(graph *Graph, root *Node, maxCost int, verbose int) (map[*Node]bool, map[*Node]bool) {
func dijkstra(graph *Graph, root *Node, maxCost int, verbose int) ([]bool, map[*Node]bool) {
	var updated string
	var current *Node
	var nextnDist int

	graph.ResetGraph()
	root.Distance = 0
	root.PqIndex = 0
	if maxCost > graph.BoardSize {
		maxCost = graph.BoardSize
	}
	unvisitedQueue := make(PriorityQueue, 0)
	// visited := make(map[*Node]bool)
	goals := make(map[*Node]bool)
	visited := make([]bool, graph.BoardSize)
	heap.Init(&unvisitedQueue)
	heap.Push(&unvisitedQueue, root)

	for len(unvisitedQueue) > 0 {
		current = heap.Pop(&unvisitedQueue).(*Node)
		if verbose >= 3 {
			fmt.Printf("\n****current: %v\n", current.Id)
		}
		// visited[current] = true
		visited[current.Ix] = true
		if current.Goal > 0 {
			goals[current] = true
		}

		if current.Distance >= maxCost {
			continue
		}
		insertionSort(current.Adjacent)
		for _, nextN := range current.Adjacent {
			// fmt.Println("nextN:", nextN.Id)
			// if _, ok := visited[nextN]; ok {
			if visited[nextN.Ix] {
				continue
			}
			// You may not enter or cross a node with a derrick or truck.
			if nextN.Derrick || (nextN.Truck != nil) {
				continue
			}
			newDist := current.Distance + nextN.Terrain
			nextnDist = nextN.Distance
			// if the next node has wells, the player may not stop here
			if nextN.Wells != 0 {
				if newDist >= maxCost {
					continue
				}
				// There is still some movement credit left so it might be possible to move to an adjacent node.
				canMove := false
				credit := maxCost - newDist
				for _, nextA := range nextN.Adjacent {
					if visited[nextA.Ix] {
						continue
					}
					if nextA.Terrain <= credit {
						canMove = true
						break
					}
				}
				if !canMove {
					continue
				}
			}
			updated = "not updated"
			if newDist < nextnDist && newDist <= maxCost {
				nextN.Distance = newDist
				nextN.Previous = current
				heap.Push(&unvisitedQueue, nextN)
				updated = "updated" // for logging
			}

			if verbose >= 3 {
				fmt.Printf("    %v: current: %v, next: %v, dist: %v -> %v", updated, current.Id, nextN.Id, iDist(nextnDist), nextN.Distance)
			}
		}
	}
	return visited, goals
}
