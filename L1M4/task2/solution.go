package main

import (
	"fmt"
	"container/list"
)

// State represents a single configuration of missionaries, cannibals, and boat position.
type State struct {
	ML, CL, MR, CR int    // Missionaries & Cannibals on left and right
	Boat           string // "L" for left, "R" for right
	Parent         *State // To track the path
}

// isValid checks if the state is valid (no missionaries eaten).
func (s State) isValid() bool {
	if (s.ML < s.CL && s.ML > 0) || (s.MR < s.CR && s.MR > 0) {
		return false
	}
	return s.ML >= 0 && s.CL >= 0 && s.MR >= 0 && s.CR >= 0
}

// isGoal checks if we reached the goal state.
func (s State) isGoal() bool {
	return s.ML == 0 && s.CL == 0 && s.MR == 3 && s.CR == 3
}

// possibleMoves represents allowed boat movements.
var possibleMoves = [][]int{
	{1, 0}, {2, 0}, {0, 1}, {0, 2}, {1, 1}, // 1M, 2M, 1C, 2C, 1M1C
}

// bfs solves the problem using BFS.
func bfs() {
	start := State{3, 3, 0, 0, "L", nil}
	visited := make(map[State]bool)
	queue := list.New()

	queue.PushBack(start)

	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(State)

		if curr.isGoal() {
			printSolution(&curr)
			return
		}

		for _, move := range possibleMoves {
			newState := getNextState(curr, move)

			if newState.isValid() && !visited[newState] {
				newState.Parent = &curr
				visited[newState] = true
				queue.PushBack(newState)
			}
		}
	}
	fmt.Println("No solution found.")
}

// getNextState generates the next possible state.
func getNextState(s State, move []int) State {
	m, c := move[0], move[1]
	if s.Boat == "L" {
		return State{s.ML - m, s.CL - c, s.MR + m, s.CR + c, "R", nil}
	}
	return State{s.ML + m, s.CL + c, s.MR - m, s.CR - c, "L", nil}
}

// printSolution prints the path in the requested format.
func printSolution(s *State) {
	var path []*State
	for s != nil {
		path = append([]*State{s}, path...)
		s = s.Parent
	}

	fmt.Println("Solution Path:")
	for _, state := range path {
		fmt.Printf("Left: (%dM, %dC) | Right: (%dM, %dC) | Boat: %s\n",
			state.ML, state.CL, state.MR, state.CR, state.Boat)
	}
}

func main() {
	bfs()
}
