package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"testing"
	"time"
)

func TestCreateBlock(t *testing.T) {
	data := "test data"
	prevHash := []byte("prev hash")
	block := CreateBlock(data, prevHash)

	if block.Data == nil || string(block.Data) != data {
		t.Errorf("Expected data to be '%s', got '%s'", data, block.Data)
	}
	if block.PrevHash == nil || !bytes.Equal(block.PrevHash, prevHash) {
		t.Errorf("Expected previous hash to be '%x', got '%x'", prevHash, block.PrevHash)
	}
	if block.Timestamp <= 0 {
		t.Error("Expected timestamp to be greater than 0")
	}
	if block.Hash == nil || len(block.Hash) == 0 {
		t.Error("Expected hash to be non-empty")
	}
}

func TestDeriveHash(t *testing.T) {
	data := "test data"
	prevHash := []byte("prev hash")
	timestamp := time.Now().Unix()
	block := &Block{timestamp, []byte(data), prevHash, []byte{}}
	block.DeriveHash()

	expectedHash := sha256.Sum256(bytes.Join([][]byte{prevHash, []byte(data), []byte(strconv.FormatInt(timestamp, 10))}, []byte{}))
	if !bytes.Equal(block.Hash, expectedHash[:]) {
		t.Errorf("Expected hash to be '%x', got '%x'", expectedHash, block.Hash)
	}
}

func TestGenesis(t *testing.T) {
	genesis := Genesis()

	if genesis == nil {
		t.Error("Genesis block should not be nil")
	}
	if genesis.PrevHash == nil || len(genesis.PrevHash) != 0 {
		t.Error("Genesis block's previous hash should be empty")
	}
	if string(genesis.Data) != "Genesis" {
		t.Errorf("Expected Genesis block's data to be 'Genesis', got '%s'", genesis.Data)
	}
}
