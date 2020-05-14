package main

import (
	"fmt"
	"BlockChain/blockchain"
)

func main(){
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First")
	chain.AddBlock("Second")
	chain.AddBlock("Third")
	fmt.Println(blockchain.NewProof(chain.Blocks[0]))
}