package main

import "crypto/sha256"

type MerkleTree struct {
	Root *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {

	newNode := MerkleNode{}
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		newNode.Data = hash[:]
	} else {
		prevHash := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHash)
		newNode.Data = hash[:]
	}
	newNode.Left = left
	newNode.Right = right
	return &newNode
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, leaves := range data {
		node := NewMerkleNode(nil, nil, leaves)
		nodes = append(nodes, *node)
	}

	for i := 0; i < len(data)/2; i++ {
		var newLevel []MerkleNode
		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}
		nodes = newLevel
	}

	mktree := MerkleTree{&nodes[0]}
	return &mktree
}
