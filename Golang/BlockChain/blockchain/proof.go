package blockchain

import (
	"bytes"
	"math"
	"math/big"
	"crypto/sha256"
	"encoding/binary"
)
// consensus algo
// PROOF OF WORK ALGO


// Steps:
// 1. Grab data from the block     			data := GetBlockData()
// 2. Create a counter starts from 0    	counter := 0
// 3. Create a hash of the data + counter	hash := data + counter
// 4. Check hash 							if hash.TestRequirements() == true


// Requirements:
// 1. The first few bytes must contain 0s

const Difficulty = 12  // the number of consecutive 0s in the front

type ProofOfWork struct{
	Block *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork{
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty)) // left shift
	return &ProofOfWork{
		Block: b,
		Target: target,
	}
}

func (p *ProofOfWork) InitData(nonce int) []byte{
	data := bytes.Join(
		[][]byte{
			p.Block.Pre_Hash,
			p.Block.Data,
			ToHex(int64(nonce)), 
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

func ToHex(num int64) []byte{
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num) // convert num to binary and write to buffer
	if err != nil{
		panic(err)
	}
	return buff.Bytes()
}

func (p *ProofOfWork) Run() (int, []byte){
	var intHash big.Int
	var hash [32]byte
	nonce := 0
	for nonce < math.MaxInt64{
		data := p.InitData(nonce)
		hash := sha256.Sum256(data)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(p.Target) == -1{
			break
		}else{
			nonce ++
		}
	}
	return nonce, hash[:]
}
