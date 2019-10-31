package main

import (
	"fmt"
	"sync"
)

const (
	CONSTMULTI = 2
)

// Nodes[0] is winner
type LoserTree struct {
	Nodes []*Node
}

type Node struct {
	Value int64
	From  *Entrance
}

type Entrance struct {
	From []int64
	To   int
	Now  int
}

func buildLoserTree(channels [][]int64) *LoserTree {
	length := len(channels)
	beforeNodesNum := 1
	for beforeNodesNum < length {
		beforeNodesNum *= 2
	}

	loserTree := &LoserTree{
		Nodes: make([]*Node, beforeNodesNum+length),
	}

	entrances := make([]*Entrance, length)
	for index, _ := range entrances {
		entrances[index] = &Entrance{
			From: channels[index],
			To:   0,
			Now:  0,
		}
	}

	for i := beforeNodesNum; i < beforeNodesNum+length; i++ {
		loserTree.Nodes[i] = &Node{
			Value: entrances[i-beforeNodesNum].From[0],
			From:  entrances[i-beforeNodesNum],
		}
		entrances[i-beforeNodesNum].To = i
	}

	for i := beforeNodesNum; i < beforeNodesNum+length; i++ {
		index := i
		winner := loserTree.Nodes[index]
		for index >= 1 {
			if loserTree.Nodes[index/2] == nil {
				loserTree.Nodes[index/2] = &Node{
					Value: winner.Value,
					From:  winner.From,
				}
			} else {
				if winner.Value < loserTree.Nodes[index/2].Value {
					// winner still win
				} else {
					// winner lose, loser win
					winner, loserTree.Nodes[index/2] = loserTree.Nodes[index/2], winner
				}
			}
			index /= 2
		}
		loserTree.Nodes[0] = winner
	}
	return loserTree
}

func (loserTree *LoserTree) GetWinner() int64 {
	winner := loserTree.Nodes[0]
	if loserTree.Nodes[1] == nil {
		loserTree.Nodes[0] = nil
		return winner.Value
	}
	winnerFrom := winner.From
	winnerFrom.Now += 1
	nowIndex := winnerFrom.To
	if winnerFrom.Now >= len(winnerFrom.From) {
		loserTree.Nodes[nowIndex] = nil
		nowIndex /= 2
	} else {
		loserTree.Nodes[nowIndex] = &Node{
			Value: winnerFrom.From[winnerFrom.Now],
			From:  winnerFrom,
		}
	}

	// be challenger
	challenger := loserTree.Nodes[nowIndex]
	for challenger == nil {
		challenger = loserTree.Nodes[nowIndex/2]
		nowIndex /= 2
	}
	loserTree.Nodes[nowIndex] = nil
	loserTree.Nodes[0].Value = challenger.Value + 1
	// start challenge
	for nowIndex >= 1 {
		if challenger.Value < loserTree.Nodes[nowIndex/2].Value {
			// win
		} else {
			// lose
			challenger, loserTree.Nodes[nowIndex/2] = loserTree.Nodes[nowIndex/2], challenger
		}
		nowIndex /= 2
	}
	loserTree.Nodes[0] = challenger

	return winner.Value
}

func (loserTree *LoserTree) Print() {
	for _, node := range loserTree.Nodes {
		if node == nil {
			fmt.Printf("* ")
		} else {
			fmt.Printf("%d ", node.Value)
		}
	}
	fmt.Println()
}

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	// fmt.Println("Sort: ", src)
	MULTI := CONSTMULTI
	length := len(src)
	if length <= 1 {
		return
	}
	if length == 2 {
		if src[0] > src[1] {
			src[0], src[1] = src[1], src[0]
		}
		return
	}
	if length < MULTI {
		MULTI = length
	}
	results := make([][]int64, MULTI)
	start := 0
	for i := 0; i < MULTI; i++ {
		end := start + length/MULTI
		if i == MULTI-1 && end != length {
			end = length
		}
		results[i] = make([]int64, end-start)
		results[i] = src[start:end]
		start = end
	}
	wg := new(sync.WaitGroup)
	for index, _ := range results {
		wg.Add(1)
		go MergeSort(results[index])
		wg.Done()
	}
	wg.Wait()

	fmt.Println("Sort:", results)
	loserTree := buildLoserTree(results)

	newSrc := make([]int64, len(src))
	for i := 0; i < length; i++ {
		loserTree.Print()
		newSrc[i] = loserTree.GetWinner()
		fmt.Println("winner:", newSrc[i])
	}
	copy(src, newSrc)
	return
}
