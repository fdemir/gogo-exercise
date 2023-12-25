package lib

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Data struct {
	amount float64
}

type Block struct {
	index     int
	timestamp time.Time
	data      Data
	prevHash  string
	hash      string
}

type BlockChain struct {
	chain []Block
	mutex sync.Mutex
}

func (b *Block) CalculateHash() string {
	var buf []byte

	buf = append(buf, []byte(strconv.Itoa(b.index))...)
	buf = append(buf, []byte(
		fmt.Sprintf("%f", b.data.amount),
	)...)
	buf = append(buf, []byte(b.prevHash)...)
	b.hash = fmt.Sprintf("%x", sha256.Sum256(buf))

	return b.hash
}

func (bc *BlockChain) AddBlock(
	amount float64,
) {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()
	var prevHash string

	if len(bc.chain) > 0 {
		prevHash = bc.chain[len(bc.chain)-1].hash
	}

	block := &Block{
		index:     len(bc.chain),
		timestamp: time.Now(),
		data:      Data{amount: amount},
		prevHash:  prevHash,
	}

	block.CalculateHash()
	bc.chain = append(bc.chain, *block)
}

// it's for testing purpose only.
func (bc *BlockChain) DangerouslySetBlockAmount(
	index int,
	amount float64,
) {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()
	bc.chain[index].data.amount = amount
	bc.chain[index].CalculateHash()
}

func (bc *BlockChain) GetChain() []Block {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()
	return bc.chain
}

func (bc *BlockChain) GetLatestBlock() *Block {
	return &bc.chain[len(bc.chain)-1]
}

func (bc *BlockChain) String() string {
	var sb strings.Builder

	for _, block := range bc.chain {
		sb.WriteString(fmt.Sprintf("Block %d\n Time %s\n Data %v\n Hash %s\n Amount %s\n\n", block.index, block.timestamp.String(), block.data, block.hash,
			fmt.Sprintf("%f", block.data.amount)))
	}

	return sb.String()
}

func (bc *BlockChain) IsValid() bool {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	for i := 1; i < len(bc.chain); i++ {
		currentBlock := bc.chain[i]
		prevBlock := bc.chain[i-1]

		if currentBlock.hash != currentBlock.CalculateHash() {
			return false
		}

		if currentBlock.prevHash != prevBlock.hash {
			return false
		}
	}

	return true
}
