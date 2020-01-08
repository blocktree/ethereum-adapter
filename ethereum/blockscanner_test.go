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
	"github.com/blocktree/openwallet/log"
	"testing"
)

func TestWalletManager_EthGetTransactionByHash(t *testing.T) {
	wm := testNewWalletManager()
	txid := "0xaf50436156307f7642449d60247417d214346e067dd6b5fa9b7ecc060e8a2206"
	tx, err := wm.WalletClient.EthGetTransactionByHash(txid)
	if err != nil {
		t.Errorf("get transaction by has failed, err=%v", err)
		return
	}
	log.Infof("tx: %+v", tx)
}

func TestWalletManager_ethGetTransactionReceipt(t *testing.T) {
	wm := testNewWalletManager()
	txid := "0x24f3accc0a71408a19f7a55b55db476cc94f3bcc0356b0d6e94c3ba5ae67608b"
	tx, err := wm.WalletClient.EthGetTransactionReceipt(txid)
	if err != nil {
		t.Errorf("get transaction by has failed, err=%v", err)
		return
	}
	log.Infof("tx: %+v", tx)
}

func TestWalletManager_EthGetBlockNumber(t *testing.T) {
	wm := testNewWalletManager()
	maxBlockHeight, err := wm.WalletClient.EthGetBlockNumber()
	if err != nil {
		t.Errorf("EthGetBlockNumber failed, err=%v", err)
		return
	}
	log.Infof("maxBlockHeight: %v", maxBlockHeight)
}