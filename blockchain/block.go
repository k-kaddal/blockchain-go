package blockchain

type Chain struct {
	Blocks []*Block
}

type Block struct {
	Hash      []byte
	Data      []byte
	PrevBlock []byte
	Nonce     int
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	// run proof of work
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func (chain *Chain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genisis", []byte{})
}

func InitChain() *Chain {
	return &Chain{[]*Block{Genesis()}}
}
