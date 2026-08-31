package findminimummaximumbetweencriticalpoints

import "math"

type ListNode struct {
	Val int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
    if head == nil || head.Next == nil || head.Next.Next == nil {
		return []int{-1, -1}
	}

	prev := head
	curr := head.Next

	index := 2
	
	inf := math.MaxInt
	firstIndex := -1
	prevIndex := -1
	minDist := inf

	for curr != nil && curr.Next != nil {
		isMax := curr.Val > prev.Val && curr.Val > curr.Next.Val
		isMin := curr.Val < prev.Val && curr.Val < curr.Next.Val

		if isMax || isMin {
			if firstIndex == -1 {
				firstIndex = index
			} else {
				dist := index - prevIndex
				if dist < minDist {
					minDist = dist
				}
			}
			prevIndex = index
		}
		
		prev = curr
		curr = curr.Next
		index++
	}

	if minDist == inf {
		return []int{-1, -1}
	}

	maxDist := prevIndex - firstIndex
	return []int{minDist, maxDist}
}