package main

// Block is a single block within the blockchain.
type Block struct {
	index        int
	timestamp    float64
	transactions []Transaction
	proof        int
	previousHash string
}
