package ldbclone

import "bytes"

type SkipList struct {
	head SkipNode
}

type SkipNode struct {
	maxLanes int
	maxIdx   int
	Key      []byte
	Val      []byte
	Paths    []SkipNode
}

func NewSkipNode(maxLanes int) SkipNode {
	return SkipNode{
		maxLanes: maxLanes,
		Paths:    make([]SkipNode, maxLanes),
	}
}

func (l *SkipList) Search(searchVal []byte) []byte {
	currentNode := l.head
	for {
		if res := bytes.Compare(searchVal, currentNode.Key); res == 0 {
			return currentNode.Val
		} else if res == -1 {
			currentNode = *currentNode.FindNode(searchVal)
		} else {
			return nil
		}
	}

}

func (s *SkipNode) FindNode(searchVal []byte) *SkipNode {
	for i := s.maxIdx; i >= 0; i-- {
		if cmp := bytes.Compare(searchVal, s.Paths[i].Key); cmp >= 0 {
			return &s.Paths[i]
		}
	}
	return &s.Paths[0]
}

//  7
//  |--------------|
//  5 -> 6 -> 7 -> 8
