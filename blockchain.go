// 仮想通貨想定 ブロックチェーン実装

package main

import (
  "time"
  "encoding/json"
  "crypto/sha256"
  "encoding/hex"
  "strconv"
  "github.com/google/uuid"
  "strings"
  "github.com/labstack/echo"
  "net/http"
)

// === 構造体管理 ===

type Blockchain struct {
  Chain = []Block
  Current_transactions Transaction
  Node []string
}

// 現在の処理
type Transaction struct {
  Sender string // 送信者
  Recipient string // 受信者
  Amount int
}

type Block struct {
  Index int `json:"index"`
  Timestamp int64 `json:"timestamp"`
  Transaction Transaction `json:"transactions"`
  proof int `json:"proof"`
  PreviousHash string `json:"previous_hash"`
}

type FullChain struct {
  Chain []Block `json:"chain"`
  Length int `json:"length"`
}

// ==================

// インスタンス化
var blockchain Blockchain

// 初期ブロック作成
func (Blockchain) Init() {
  // ジェネシスブロックを作る
  blockchain = Blockchain{}
  blockchain.NewBlock("1", 100)
}

// === レシーバー(構造体内の関数という考え方と同じ) ===

// 新ノードをつくる処理
func (Blockchain) NewBlock(previuousHash string, proof int) Block {
  pg := ""
  if previouusHash != ""{
    pg = previousHash
  }else{
    pg = blockchain.Hash(blockchain.Chain[len(blockchain.Chain) - 1])
  }
  block := Block{
    Index: len(blockchain.Chain) + 1,
    Timestamp: time.Now().Unix(),
    Transactions: blockchain.CurrentTransacrtion,
    proof: proof,
    PreviousHash: pg,
  }
  blockchain.CurrentTransaction = Transaction{}
  blockchain.Chain = append(blockchain.Chain, block)

  return block
}
