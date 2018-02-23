package main

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/t-tojima/sandbox/go-solidity/token"
	"log"
	"math/big"
)

func generateAccount() (*ecdsa.PrivateKey, *bind.TransactOpts, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil, err
	}

	auth := bind.NewKeyedTransactor(key)
	return key, auth, nil
}

func printTokens() {
}

func main() {
	key1, auth1, _ := generateAccount()
	key2, auth2, _ := generateAccount()

	conn := backends.NewSimulatedBackend(core.GenesisAlloc{
		auth1.From: {
			PrivateKey: key1.D.Bytes(),
			Balance:    big.NewInt(10000000000000),
		},
		auth2.From: {
			PrivateKey: key2.D.Bytes(),
			Balance:    big.NewInt(10000000000000),
		},
	})

	radexAddr, _, radex, err := token.DeployRadex(auth1, conn)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to deploy radex contract"))
	}
	log.Printf("Radex address = %s\n", radexAddr.Hex())
	tkn1Addr, _, tkn1, err := token.DeployMyToken(auth1, conn)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to deploy token contract"))
	}
	log.Printf("Tkn1 address = %s\n", tkn1Addr.Hex())
	tkn2Addr, _, tkn2, err := token.DeployMyToken(auth2, conn)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to deploy token contract"))
	}
	log.Printf("Tkn2 address = %s\n", tkn2Addr.Hex())

	conn.Commit()

	{
		balance, _ := tkn1.BalanceOf(nil, auth1.From)
		log.Printf("[Auth1] Balance of Tkn1 = %d\n", balance.Int64())
	}
	{
		balance, _ := tkn1.BalanceOf(nil, auth2.From)
		log.Printf("[Auth2] Balance of Tkn1 = %d\n", balance.Int64())
	}
	{
		balance, _ := tkn2.BalanceOf(nil, auth1.From)
		log.Printf("[Auth1] Balance of Tkn2 = %d\n", balance.Int64())
	}
	{
		balance, _ := tkn2.BalanceOf(nil, auth2.From)
		log.Printf("[Auth2] Balance of Tkn2 = %d\n", balance.Int64())
	}

	_, err = tkn1.Transfer(auth1, radexAddr, big.NewInt(50000), nil)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to transfer token to radex"))
	}
	_, err = tkn2.Transfer(auth2, radexAddr, big.NewInt(50000), nil)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to transfer token to radex"))
	}

	conn.Commit()
	log.Println("Transfer executed")

	{
		balance, _ := radex.BalanceOf(nil, tkn1Addr, auth1.From)
		log.Printf("[Radex / Auth1] Balance of Tkn1 = %d\n", balance.Int64())
	}
	{
		balance, _ := radex.BalanceOf(nil, tkn1Addr, auth2.From)
		log.Printf("[Radex / Auth2] Balance of Tkn1 = %d\n", balance.Int64())
	}
	{
		balance, _ := radex.BalanceOf(nil, tkn2Addr, auth1.From)
		log.Printf("[Radex / Auth1] Balance of Tkn2 = %d\n", balance.Int64())
	}
	{
		balance, _ := radex.BalanceOf(nil, tkn2Addr, auth2.From)
		log.Printf("[Radex / Auth2] Balance of Tkn2 = %d\n", balance.Int64())
	}

	_, err = radex.CreateOrder(auth1, tkn1Addr, tkn2Addr, big.NewInt(50000), big.NewInt(1), big.NewInt(1))
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to create order to radex"))
	}

	conn.Commit()
	log.Println("Create order executed")

	{
		balance, _ := radex.BalanceOf(nil, tkn1Addr, auth1.From)
		log.Printf("[Radex / Auth1] Balance of Tkn1 = %d\n", balance.Int64())
	}
	{
		balance, _ := radex.BalanceOf(nil, tkn1Addr, auth2.From)
		log.Printf("[Radex / Auth2] Balance of Tkn1 = %d\n", balance.Int64())
	}
	{
		balance, _ := radex.BalanceOf(nil, tkn2Addr, auth1.From)
		log.Printf("[Radex / Auth1] Balance of Tkn2 = %d\n", balance.Int64())
	}
	{
		balance, _ := radex.BalanceOf(nil, tkn2Addr, auth2.From)
		log.Printf("[Radex / Auth2] Balance of Tkn2 = %d\n", balance.Int64())
	}

	iter, err := radex.FilterNewOrder(nil, []common.Address{auth1.From},
		[]common.Address{tkn1Addr}, []common.Address{tkn2Addr})
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get filter NewOrder"))
	}

	if !iter.Next() {
		log.Fatal("failed get log")
	}
	orderId := iter.Event.Id

	_, err = radex.ExecuteOrder(auth2, orderId, big.NewInt(50000))
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to execute order"))
	}

	conn.Commit()
	log.Println("Execute Order executed")

	{
		balance, _ := radex.BalanceOf(nil, tkn1Addr, auth1.From)
		log.Printf("[Radex / Auth1] Balance of Tkn1 = %d\n", balance.Int64())
	}
	{
		balance, _ := radex.BalanceOf(nil, tkn1Addr, auth2.From)
		log.Printf("[Radex / Auth2] Balance of Tkn1 = %d\n", balance.Int64())
	}
	{
		balance, _ := radex.BalanceOf(nil, tkn2Addr, auth1.From)
		log.Printf("[Radex / Auth1] Balance of Tkn2 = %d\n", balance.Int64())
	}
	{
		balance, _ := radex.BalanceOf(nil, tkn2Addr, auth2.From)
		log.Printf("[Radex / Auth2] Balance of Tkn2 = %d\n", balance.Int64())
	}

	{
		balance, _ := radex.BalanceOf(nil, tkn1Addr, auth1.From)
		radex.Redeem(auth1, tkn1Addr, balance)
	}
	{
		balance, _ := radex.BalanceOf(nil, tkn2Addr, auth1.From)
		radex.Redeem(auth1, tkn2Addr, balance)
	}
	log.Println("Redeem auth1 tokens")

	{
		balance, _ := radex.BalanceOf(nil, tkn1Addr, auth2.From)
		radex.Redeem(auth2, tkn1Addr, balance)
	}
	{
		balance, _ := radex.BalanceOf(nil, tkn2Addr, auth2.From)
		radex.Redeem(auth2, tkn2Addr, balance)
	}
	log.Println("Redeem auth2 tokens")

	conn.Commit()

	{
		balance, _ := tkn1.BalanceOf(nil, auth1.From)
		log.Printf("[Auth1] Balance of Tkn1 = %d\n", balance.Int64())
	}
	{
		balance, _ := tkn1.BalanceOf(nil, auth2.From)
		log.Printf("[Auth2] Balance of Tkn1 = %d\n", balance.Int64())
	}
	{
		balance, _ := tkn2.BalanceOf(nil, auth1.From)
		log.Printf("[Auth1] Balance of Tkn2 = %d\n", balance.Int64())
	}
	{
		balance, _ := tkn2.BalanceOf(nil, auth2.From)
		log.Printf("[Auth2] Balance of Tkn2 = %d\n", balance.Int64())
	}
	{
		balance, _ := tkn1.BalanceOf(nil, radexAddr)
		log.Printf("[Radex] Balance of Tkn1 = %d\n", balance.Int64())
	}
	{
		balance, _ := tkn2.BalanceOf(nil, radexAddr)
		log.Printf("[Radex] Balance of Tkn2 = %d\n", balance.Int64())
	}

}
