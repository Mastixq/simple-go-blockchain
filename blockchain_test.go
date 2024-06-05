package main

import (
	"testing"
)

func TestBlockchain(t *testing.T) {
	bc := NewBlockchain()
	if len(bc.Blocks) != 1 {
		t.Fatalf("Expected genesis block, but got %d blocks", len(bc.Blocks))
	}

	bc.AddTransaction("Alice", "Bob", 10)
	if len(bc.PendingTransactions) != 1 {
		t.Fatalf("Expected 1 pending transaction, but got %d", len(bc.PendingTransactions))
	}

	proof := bc.ProofOfWork(bc.Blocks[len(bc.Blocks)-1].Proof)
	bc.AddBlock(proof)

	if len(bc.Blocks) != 2 {
		t.Fatalf("Expected 2 blocks, but got %d blocks", len(bc.Blocks))
	}

	if len(bc.PendingTransactions) != 0 {
		t.Fatalf("Expected 0 pending transactions, but got %d", len(bc.PendingTransactions))
	}

	if !isValidChain(bc.Blocks) {
		t.Fatalf("Blockchain is not valid")
	}
}
