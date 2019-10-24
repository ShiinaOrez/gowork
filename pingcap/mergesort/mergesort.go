package main

import (
	"sync"
    _ "fmt"
)

const (
	CONSTMULTI = 4
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
    		To: 0,
    		Now: 0,
    	}
    }

    for i:=beforeNodesNum; i<beforeNodesNum+length; i++ {
    	loserTree.Nodes[i] = &Node{
    		Value: entrances[i-beforeNodesNum].From[0],
    		From: entrances[i-beforeNodesNum],
    	}
    	entrances[i-beforeNodesNum].To = i
    }

    for i:=beforeNodesNum; i<beforeNodesNum+length; i++ {
    	index := i
    	winner := loserTree.Nodes[index]
    	for index >= 1 {
    		if loserTree.Nodes[index/2] == nil {
    			loserTree.Nodes[index/2] = &Node{
    				Value: winner.Value,
    				From: winner.From,
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
    }
    return loserTree
}

func (loserTree *LoserTree) GetWinner() int64 {
	winner := loserTree.Nodes[0].Value
	entrance := loserTree.Nodes[0].From
    challengerIndex := entrance.To

    entrance.Now += 1
    if entrance.Now >= len(entrance.From) {
    	// entrance is empty
    	loserTree.Nodes[entrance.To] = nil
    	challengerIndex = challengerIndex / 2
    }

    challenger := loserTree.Nodes[challengerIndex]

    loserTree.Nodes[0] = nil
    for challengerIndex >= 1 {
    	if loserTree.Nodes[challengerIndex/2] == nil {
    		loserTree.Nodes[challengerIndex/2] = challenger
    	} else {
	    	if challenger.Value < loserTree.Nodes[challengerIndex/2].Value {
    			// challenger win
    		} else {
    			// challenger lose
    			challenger, loserTree.Nodes[challengerIndex/2] = loserTree.Nodes[challengerIndex/2], challenger
    		}
    	}
    	challengerIndex /= 2
    }

    return winner
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
	for i:=0; i<MULTI; i++ {
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

	loserTree := buildLoserTree(results)

    newSrc := make([]int64, len(src))
    for i:=0; i<length; i++ {
    	newSrc[i] = loserTree.GetWinner()
    }
    copy(src, newSrc)
    return
}
