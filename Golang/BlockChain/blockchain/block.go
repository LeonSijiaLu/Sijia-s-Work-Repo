package blockchain

import (
	"fmt"
)

type Block struct{
	Hash []byte
	Data []byte
	Pre_Hash []byte // last block hash, allows us to link blocks like a linked list
	Nonce int
}

type BlockChain struct{
	Blocks []*Block
}

func CreateBlock(data string, pre_hash []byte) *Block{
	block := &Block{
		Hash: []byte{},
		Data: []byte(data),
		Pre_Hash: pre_hash,
	}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func (chain *BlockChain) AddBlock(data string){
	chain.Blocks = append(chain.Blocks, CreateBlock(data, chain.Blocks[len(chain.Blocks) - 1].Hash))
}

func Genesis() *Block{
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain{
	return &BlockChain{
		Blocks: []*Block{Genesis()},
	}
}

func (chain *BlockChain) TraverseChain(){
	for _, block := range chain.Blocks{
		fmt.Printf("Data: %s \n", block.Data)
		fmt.Printf("Hash: %x \n", block.Hash)
	}
}