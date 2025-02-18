package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	t := &TreeNode{}

	root := build(t, 0, 5)
	fmt.Println(root)
	printTree(root)
	fmt.Println(nodesPerLevel(root))
}

func nodesPerLevel(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var info []levelInfo
	var output []float64

	_nodesPerLevel(root, 0, &info, &output)
	return output
}

type levelInfo struct {
	totalSum      int
	amountOfNodos int
	avg           float64
}

func _nodesPerLevel(n *TreeNode, l int, info *[]levelInfo, output *[]float64) {
	if n == nil {
		return
	}

	if len(*info) > l {
		(*info)[l].totalSum += n.Val
		(*info)[l].amountOfNodos += 1
		(*info)[l].avg = float64((*info)[l].totalSum) / float64((*info)[l].amountOfNodos)
		(*output)[l] = (*info)[l].avg
	} else {
		newInfo := &levelInfo{
			totalSum:      n.Val,
			amountOfNodos: 1,
			avg:           float64(n.Val),
		}
		*info = append(*info, *newInfo)
		*output = append(*output, float64(n.Val))
	}

	_nodesPerLevel(n.Left, l+1, info, output)
	_nodesPerLevel(n.Right, l+1, info, output)
}

// printTree prints the binary tree in a structured format
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("Tree is empty")
		return
	}

	// Get tree levels
	levels := getLevels(root)

	// Print each level
	for _, level := range levels {
		fmt.Println(strings.Join(level, "  "))
	}
}

// getLevels returns a slice of slices, where each inner slice represents a tree level as strings
func getLevels(root *TreeNode) [][]string {
	if root == nil {
		return nil
	}

	var result [][]string
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var level []string

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node != nil {
				level = append(level, fmt.Sprintf("%d", node.Val))
				queue = append(queue, node.Left, node.Right)
			} else {
				level = append(level, " ")
			}
		}

		// If the level contains only empty spaces, stop
		if isEmptyLevel(level) {
			break
		}

		result = append(result, level)
	}

	return result
}

// isEmptyLevel checks if a level contains only empty spaces
func isEmptyLevel(level []string) bool {
	for _, val := range level {
		if val != " " {
			return false
		}
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func build(t *TreeNode, l int, tl int) *TreeNode {
	if l == tl {
		return t
	}

	t.Val = rand.Intn(41)

	t.Left = &TreeNode{}
	t.Right = &TreeNode{}

	t.Left = build(t.Left, l+1, tl)
	t.Right = build(t.Right, l+1, tl)

	return t
}

func averageOfLevels(root *TreeNode) []float64 {
	return []float64{0.0}
}
