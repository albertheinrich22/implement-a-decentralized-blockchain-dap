package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type BlockchainParser struct {
	privateKey *ecdsa.PrivateKey
}

func NewBlockchainParser(privKey string) (*BlockchainParser, error) {
	privateKey, err := hex.DecodeString(privKey)
	if err != nil {
		return nil, err
	}
	privKeyECDSA, err := btcec.ParsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return &BlockchainParser{privateKey: privKeyECDSA}, nil
}

func (p *BlockchainParser) ParseTransaction(txHex string) (*types.Transaction, error) {
	tx, err := hex.DecodeString(txHex)
	if err != nil {
		return nil, err
	}
	return types TRANSACTION(tx)
}

func (p *BlockchainParser) ParseBlock(blockHex string) (*types.Block, error) {
	block, err := hex.DecodeString(blockHex)
	if err != nil {
		return nil, err
	}
	return types.NewBlockFromByte(block)
}

func main() {
	privKey := "0x12cd3456789012345678901234567890123456789012345678901234567890123"
	parser, err := NewBlockchainParser(privKey)
	if err != nil {
		log.Fatal(err)
	}

	txHex := "0xf861808085048080608553594509101909101909101909101909101909101909101909101"
	tx, err := parser.ParseTransaction(txHex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction parsed: %+v\n", tx)

	blockHex := "0x123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345"
	block, err := parser.ParseBlock(blockHex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Block parsed: %+v\n", block)
}