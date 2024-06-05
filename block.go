package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	Proof        int
	PreviousHash string
	Hash         string
}

func (b *Block) CalculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + strconv.Itoa(b.Proof) + b.PreviousHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (b *Block) PrintBlock() {
	fmt.Printf("Index: %d\n", b.Index)
	fmt.Printf("Timestamp: %s\n", b.Timestamp)
	fmt.Printf("Transactions: %v\n", b.Transactions)
	fmt.Printf("Proof: %d\n", b.Proof)
	fmt.Printf("PreviousHash: %s\n", b.PreviousHash)
	fmt.Printf("Hash: %s\n\n", b.Hash)
}
