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
		block.PrintBlock()
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
