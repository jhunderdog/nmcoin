package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"

	"github.com/jhunderdog/nmcoin/db"
	"github.com/jhunderdog/nmcoin/utils"
)

type blockchain struct {
	NewestHash string `json: "newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.HandleErr(gob.NewDecoder(bytes.NewReader(data)).Decode(b))
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			fmt.Printf("NewestHash: %s\nHeight:%d\n", b.NewestHash, b.Height)
			// searech for checkpoint on the db
			// restore b from bytes
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis")
			} else {
				fmt.Println("Restoring...")
				b.restore(checkpoint)
			}

		})
	}
	fmt.Printf("NewestHash: %s\nHeight:%d\n", b.NewestHash, b.Height)
	return b
}
