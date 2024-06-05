package main

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()

	bc.AddTransaction("Alice", "Bob", 10.0)
	bc.AddTransaction("Bob", "Charlie", 5.0)

	lastProof := bc.Blocks[len(bc.Blocks)-1].Proof
	proof := bc.ProofOfWork(lastProof)
	bc.AddBlock(proof)

	fmt.Println("Blockchain:")
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Transactions: %v\n", block.Transactions)
		fmt.Printf("Proof: %d\n", block.Proof)
		fmt.Printf("PreviousHash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}

	bc.AddTransaction("Charlie", "Alice", 2.5)
	lastProof = bc.Blocks[len(bc.Blocks)-1].Proof
	proof = bc.ProofOfWork(lastProof)
	bc.AddBlock(proof)

	fmt.Println("Updated Blockchain:")
	for _, block := range bc.Blocks {
		block.PrintBlock()
	}
}
