package main

import (
	"fmt"
)

// Blockchain is the bread and butter here, it is the backbone for the
// sending and receiving of transactions, as well as a representation
// of all transactions that have taken place.
type Blockchain struct {
	chain               []Block
	currentTransactions []Transaction
}

// Init initializes the blockchain, seeding it with a "genesis block"
// (a block with no predecessors)
func (bc *Blockchain) Init() {
	bc.chain = make([]Block, 0)
	bc.currentTransactions = make([]Transaction, 0)

	// Create our genesis block.
	bc.newBlock(100, "1")
}

// Hash hashes a block.
func (bc *Blockchain) Hash() {

}

func (bc *Blockchain) newBlock(proof int, previousHash string) Block {
	block := Block{}

	return block
}

func (bc *Blockchain) newTransaction(sender, reciever string, amount int) int {
	bc.currentTransactions = append(bc.currentTransactions, Transaction{
		sender:   sender,
		receiver: reciever,
		amount:   amount,
	})

	return bc.lastBlock().index + 1
}

func (bc *Blockchain) lastBlock() Block {

}

func main() {
	fmt.Println("BLOCKCHAIN IN GO!")
}
