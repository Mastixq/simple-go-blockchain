package blockchain

import (
	"bytes"
	"testing"
)

func TestInitBlockChain(t *testing.T) {
	chain := InitBlockChain()

	if chain == nil {
		t.Error("Blockchain should not be nil")
	}
	if len(chain.Blocks) != 1 {
		t.Error("Blockchain should have one block (Genesis)")
	}
	if string(chain.Blocks[0].Data) != "Genesis" {
		t.Errorf("Expected Genesis block's data to be 'Genesis', got '%s'", chain.Blocks[0].Data)
	}
}

func TestAddBlock(t *testing.T) {
	chain := InitBlockChain()
	chain.AddBlock("Test Block 1")

	if len(chain.Blocks) != 2 {
		t.Errorf("Expected blockchain length to be 2, got %d", len(chain.Blocks))
	}
	if string(chain.Blocks[1].Data) != "Test Block 1" {
		t.Errorf("Expected block data to be 'Test Block 1', got '%s'", chain.Blocks[1].Data)
	}
	if !bytes.Equal(chain.Blocks[1].PrevHash, chain.Blocks[0].Hash) {
		t.Error("Previous hash of new block should match hash of the last block")
	}
}
