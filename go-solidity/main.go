package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"

	"github.com/pkg/errors"
	"github.com/t-tojima/sandbox/go-solidity/token"
	"log"
)

//go:generate abigen --sol sol/MyToken.sol --pkg token --out token/token.go

func main() {
	// アカウント作成
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	// Ethereumシミュレーターを起動
	conn := backends.NewSimulatedBackend(core.GenesisAlloc{
		auth.From: {
			PrivateKey: key.D.Bytes(),
			Balance:    big.NewInt(10000000000),
		},
	})

	// トークンのデプロイ
	_, _, token, err := token.DeployMyToken(auth, conn)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get balance"))
	}

	// ブロックに取り込ませる
	conn.Commit()

	// MyTokenのメソッドをCallする
	balance, err := token.BalanceOf(nil, auth.From)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get balance"))
	}
	log.Printf("Balance = %d\n", balance.Int64())
}
