package main

// Transaction is a representation model of the sending/receiving of resources
// via the Blockchain
type Transaction struct {
	sender   string
	receiver string
	amount   int
}
