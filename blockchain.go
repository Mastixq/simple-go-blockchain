package main

import (
	"time"
)

type Blockchain struct {
	Blocks              []Block
	PendingTransactions []Transaction
	Difficulty          int
}

func NewBlockchain() *Blockchain {
	b := &Blockchain{Difficulty: 31}
	genesisBlock := Block{Index: 0, Timestamp: time.Now().String(), Proof: 1, PreviousHash: "0"}
	genesisBlock.Hash = genesisBlock.CalculateHash()
	b.Blocks = append(b.Blocks, genesisBlock)
	return b
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, amount float64) {
	transaction := Transaction{Sender: sender, Recipient: recipient, Amount: amount}
	bc.PendingTransactions = append(bc.PendingTransactions, transaction)
}

func (bc *Blockchain) ProofOfWork(previousProof int) int {
	proof := 0
	for !isValidProof(previousProof, proof, bc.Difficulty) {
		proof++
	}
	return proof
}

func isValidProof(previousProof, proof, difficulty int) bool {
	var sum = previousProof + proof
	return sum%difficulty == 0
}

func (bc *Blockchain) AddBlock(proof int) {
	previousBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := Block{
		Index:        len(bc.Blocks),
		Timestamp:    time.Now().String(),
		Transactions: bc.PendingTransactions,
		Proof:        proof,
		PreviousHash: previousBlock.Hash,
	}
	newBlock.Hash = newBlock.CalculateHash()

	if isValidChain(bc.Blocks) {
		bc.Blocks = append(bc.Blocks, newBlock)
		bc.PendingTransactions = []Transaction{}
	}
}

func isValidChain(blocks []Block) bool {
	for i := 1; i < len(blocks); i++ {
		if blocks[i].PreviousHash != blocks[i-1].Hash {
			return false
		}
	}
	return true
}
