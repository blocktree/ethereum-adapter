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
	txid := "0x957b588e4b52b8975017ad57ec094701fe1ac5037ee49ce6ddcea64e05d0da54"
	tx, err := wm.WalletClient.EthGetTransactionByHash(txid)
	if err != nil {
		t.Errorf("get transaction by has failed, err=%v", err)
		return
	}
	log.Infof("tx: %+v", tx)
}

func TestWalletManager_ethGetTransactionReceipt(t *testing.T) {
	wm := testNewWalletManager()
	txid := "0x957b588e4b52b8975017ad57ec094701fe1ac5037ee49ce6ddcea64e05d0da54"
	tx, err := wm.WalletClient.ethGetTransactionReceipt(txid)
	if err != nil {
		t.Errorf("get transaction by has failed, err=%v", err)
		return
	}
	log.Infof("tx: %+v", tx)
}