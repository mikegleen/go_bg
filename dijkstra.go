package main

import (
	"container/heap"
	"fmt"
)

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

func Dijkstra(graph *Graph, root *Node, maxCost int, verbose int) ([]bool, map[*Node]bool) {
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
