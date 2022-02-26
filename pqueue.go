package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strconv"
)

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	// mlg: modified to give lowest
	// fmt.Print(i, j, pq)
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].PqIndex = i
	pq[j].PqIndex = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.PqIndex = n
	*pq = append(*pq, item)
}

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
func (pq *PriorityQueue) update(item *Node, value string, Distance int) {
	item.Distance = Distance
	heap.Fix(pq, item.PqIndex)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
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
	pq.update(item, item.Id, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Node)
		fmt.Printf("%.2d:%s ", item.Distance, item.Id)
	}
	fmt.Println()
}

func idist(distance int) (ret string) {
	if distance == math.MaxInt {
		ret = "∞"
	} else {
		ret = strconv.Itoa(distance)
	}
	return
}

func dijkstra(graph *Graph, root *Node, maxcost int, verbose int) (map[*Node]bool, map[*Node]bool) {
	var updated string
	var current *Node
	var nextnDist int

	graph.ResetGraph()
	root.Distance = 0
	root.PqIndex = 0
	if maxcost > graph.BoardSize {
		maxcost = graph.BoardSize
	}
	unvisited_queue := make(PriorityQueue, 0)
	visited := make(map[*Node]bool)
	goals := make(map[*Node]bool)
	heap.Init(&unvisited_queue)
	heap.Push(&unvisited_queue, root)

	for len(unvisited_queue) > 0 {
		current = heap.Pop(&unvisited_queue).(*Node)
		if verbose >= 3 {
			fmt.Printf("\n****current: %v\n", current.Id)
		}
		visited[current] = true
		if current.Goal > 0 {
			goals[current] = true
		}

		if current.Distance >= maxcost {
			continue
		}
		sort.Sort(NodeList(current.Adjacent))
		for _, nextn := range current.Adjacent {
			// fmt.Println("nextn:", nextn.Id)
			if _, ok := visited[nextn]; ok {
				continue
			}
			if nextn.Derrick || (nextn.Truck != nil) {
				continue
			}
			newDist := current.Distance + nextn.Terrain
			nextnDist = nextn.Distance
			// if the next node has wells, the player may not stop here
			if nextn.Wells != 0 && newDist >= maxcost {
				continue
			}
			updated = "not updated"
			if newDist < nextnDist && newDist <= maxcost {
				nextn.Distance = newDist
				nextn.Previous = current
				heap.Push(&unvisited_queue, nextn)
				updated = "updated" // for logging
			}

			if verbose >= 3 {
				fmt.Printf("    %v: current: %v, next: %v, dist: %v -> %v", updated, current.Id, nextn.Id, idist(nextnDist), nextn.Distance)
			}
		}
	}
	return visited, goals
}