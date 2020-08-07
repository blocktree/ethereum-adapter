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
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openwallet"
	"testing"
)

func TestWalletManager_EthGetTransactionByHash(t *testing.T) {
	wm := testNewWalletManager()
	txid := "0x4e1b5e68dffcee0f10f089c8302e0b4ccf9dd309562fe0e2861592e9a5220ebf"
	tx, err := wm.GetTransactionByHash(txid)
	if err != nil {
		t.Errorf("get transaction by has failed, err=%v", err)
		return
	}
	log.Infof("tx: %+v", tx)
}

func TestWalletManager_ethGetTransactionReceipt(t *testing.T) {
	wm := testNewWalletManager()
	txid := "0x4755c2d59e698d85cd5c77dee54a4aef52bcb9aa536d684a57c025227f3dbd74"
	tx, err := wm.GetTransactionReceipt(txid)
	if err != nil {
		t.Errorf("get transaction by has failed, err=%v", err)
		return
	}
	log.Infof("tx: %+v", tx)
}

func TestWalletManager_EthGetBlockNumber(t *testing.T) {
	wm := testNewWalletManager()
	maxBlockHeight, err := wm.GetBlockNumber()
	if err != nil {
		t.Errorf("EthGetBlockNumber failed, err=%v", err)
		return
	}
	log.Infof("maxBlockHeight: %v", maxBlockHeight)
}

func TestBlockScanner_ExtractTransactionAndReceiptData(t *testing.T) {
	wm := testNewWalletManager()

	addrs := map[string]openwallet.ScanTargetResult{
		"0x3440f720862aa7dfd4f86ecc78542b3ded900c02": openwallet.ScanTargetResult{
			SourceKey: "receiver",
			Exist:     true,
		},

		"0xbff77bb15fec867b7db7b18a34fca6d20712ce2b": openwallet.ScanTargetResult{
			SourceKey: "GOOD",
			Exist:     true,
			TargetInfo: &openwallet.SmartContract{
				ContractID: "GOOD",
				Symbol:     "ETH",
				Address:    "0xbff77bb15fec867b7db7b18a34fca6d20712ce2b",
				Token:      "S",
				Protocol:   "",
				Name:       "S",
				Decimals:   2,
			},
		},
	}
	txid := "0xda660592894dd357eedadbb69c82d7ff57859d6fb6269d2ea2eab0dce1dfd8e1"
	scanTargetFunc := func(target openwallet.ScanTargetParam) openwallet.ScanTargetResult {
		return addrs[target.ScanTarget]
	}
	result, contractResult, err := wm.GetBlockScanner().ExtractTransactionAndReceiptData(txid, scanTargetFunc)
	if err != nil {
		t.Errorf("ExtractTransactionAndReceiptData failed, err=%v", err)
		return
	}

	for sourceKey, keyData := range result {
		log.Notice("account:", sourceKey)
		for _, data := range keyData {

			for i, input := range data.TxInputs {
				log.Std.Notice("data.TxInputs[%d]: %+v", i, input)
			}

			for i, output := range data.TxOutputs {
				log.Std.Notice("data.TxOutputs[%d]: %+v", i, output)
			}

			log.Std.Notice("data.Transaction: %+v", data.Transaction)
		}
	}

	for sourceKey, keyData := range contractResult {
		log.Notice("contractID:", sourceKey)
		log.Std.Notice("data.ContractTransaction: %+v", keyData)
	}
}
