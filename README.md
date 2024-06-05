# Simple Blockchain in Go

This repository contains a simple implementation of a blockchain written in Go. It includes functionalities for creating a blockchain, adding transactions, mining new blocks using proof-of-work, and verifying the integrity of the blockchain.

## Features

- **Creating a Blockchain**: Initializes a blockchain with a genesis block.
- **Adding Transactions**: Allows adding new transactions to a pending queue.
- **Mining New Blocks**: Mines new blocks using a simplified proof-of-work algorithm.
- **Verifying Blockchain Integrity**: Ensures integrity by validating each block's hash.

## Code Structure

### `block.go`

Defines the `Block` structure and the `CalculateHash` method.

### `transaction.go`
Defines the `Transaction` structure.

### `blockchain.go`
Defines the `Blockchain` structure and methods for adding transactions and blocks.

### `main.go`
Demonstrates how to use the blockchain.