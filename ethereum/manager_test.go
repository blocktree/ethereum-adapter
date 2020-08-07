/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package ethereum

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/v2/common"
	"github.com/blocktree/openwallet/v2/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"path/filepath"
	"strings"
	"testing"
)

var (
	tw *WalletManager
)

func init() {

	tw = testNewWalletManager()
}

func testNewWalletManager() *WalletManager {
	wm := NewWalletManager()

	//读取配置
	absFile := filepath.Join("conf", "ETH.ini")
	//log.Debug("absFile:", absFile)
	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		panic(err)
	}
	wm.LoadAssetsConfig(c)
	wm.WalletClient.Debug = true
	return wm
}

func TestFixGasLimit(t *testing.T) {
	fixGasLimitStr := "sfsd"
	fixGasLimit := new(big.Int)
	fixGasLimit.SetString(fixGasLimitStr, 10)
	fmt.Printf("fixGasLimit: %d\n", fixGasLimit.Int64())
}

func TestWalletManager_GetAddrBalance(t *testing.T) {
	wm := testNewWalletManager()
	balance, err := wm.GetAddrBalance("", "pending")
	if err != nil {
		t.Errorf("GetAddrBalance2 error: %v", err)
		return
	}
	ethB := common.BigIntToDecimals(balance, wm.Decimal())
	log.Infof("ethB: %v", ethB)
}

func TestWalletManager_SetNetworkChainID(t *testing.T) {
	wm := testNewWalletManager()
	id, err := wm.SetNetworkChainID()
	if err != nil {
		t.Errorf("SetNetworkChainID error: %v", err)
		return
	}
	log.Infof("chainID: %d", id)
}

func TestWalletManager_GetTransactionFeeEstimated(t *testing.T) {
	wm := testNewWalletManager()
	txFee, err := wm.GetTransactionFeeEstimated(
		"0xc883aaf61a15da53cb3071e52b1760eb1355ba78",
		"0xf33c594038f41d5fa2e8b7d8b491aba2aca650b1",
		big.NewInt(1000000),
		nil)
	if err != nil {
		t.Errorf("GetTransactionFeeEstimated error: %v", err)
		return
	}
	log.Infof("txfee: %v", txFee)
}

func TestWalletManager_GetGasPrice(t *testing.T) {
	wm := testNewWalletManager()
	price, err := wm.GetGasPrice()
	if err != nil {
		t.Errorf("GetGasPrice error: %v", err)
		return
	}
	log.Infof("price: %v", price.String())
}

func TestWalletManager_GetTransactionCount(t *testing.T) {
	wm := testNewWalletManager()
	count, err := wm.GetTransactionCount("")
	if err != nil {
		t.Errorf("GetTransactionCount error: %v", err)
		return
	}
	log.Infof("count: %v", count)
}

func TestWalletManager_IsContract(t *testing.T) {
	wm := testNewWalletManager()
	a, err := wm.IsContract("0x3440f720862aa7dfd4f86ecc78542b3ded900c02")
	log.Infof("IsContract: %v", a)
	if err != nil {
		t.Errorf("IsContract error: %v", err)
		return
	}

	c, _ := wm.IsContract("0x627b11ead4eb39ebe61a70ab3d6fe145e5d06ab6")
	log.Infof("IsContract: %v", c)

}

func TestWalletManager_DecodeReceiptLogResult(t *testing.T) {
	wm := testNewWalletManager()
	abiJSON := `
[{"inputs":[{"internalType":"contract KeyValueStorage","name":"storage_","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"implementation","type":"address"}],"name":"Upgraded","type":"event"},{"payable":true,"stateMutability":"payable","type":"fallback"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getOwner","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"implementation","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"impl","type":"address"}],"name":"upgradeTo","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`
	logJSON := `
			{
                "logIndex": "0x0",
                "transactionIndex": "0x0",
                "transactionHash": "0x6a949727089705103e873c5dc9ebfaac79deb5fe5df0b9f02672988336130af9",
                "blockHash": "0xd80805f3b261f8dc9fd95a60030615c20ff1ca29ecb34101faf91512aedd9f2c",
                "blockNumber": "0x4b",
                "address": "0xf8afe0a06e27ddbd5ec8adbbd5cee5220c3d4d85",
                "data": "0x",
                "topics": [
                    "0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b",
                    "0x00000000000000000000000044f64ef4bc4952b133a9c4b07157770f048eebe9"
                ],
                "type": "mined"
            }
`
	var logObj types.Log
	err := logObj.UnmarshalJSON([]byte(logJSON))
	if err != nil {
		t.Errorf("UnmarshalJSON error: %v", err)
		return
	}

	abiInstance, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		t.Errorf("abi.JSON error: %v", err)
		return
	}

	rMap, name, rJSON, err := wm.DecodeReceiptLogResult(abiInstance, logObj)
	if err != nil {
		t.Errorf("DecodeReceiptLogResult error: %v", err)
		return
	}
	log.Infof("rMap: %+v", rMap)
	log.Infof("name: %+v", name)
	log.Infof("rJSON: %s", rJSON)
}
