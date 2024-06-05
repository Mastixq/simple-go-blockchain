package main

import (
	"crypto/sha256"
	"encoding/hex"
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
