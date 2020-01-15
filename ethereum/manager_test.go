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
	"github.com/blocktree/openwallet/log"
	"math/big"
	"path/filepath"
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
	absFile := filepath.Join("conf", "conf.ini")
	//log.Debug("absFile:", absFile)
	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		panic(err)
	}
	wm.LoadAssetsConfig(c)
	wm.WalletClient.Debug = true
	return wm
}


func TestWalletManager_GetErc20TokenEvent(t *testing.T) {
	wm := testNewWalletManager()
	txid := "0xa10ecccac1e3ee911fec660d5d789885e472262eadf13c372d6a2b30aca9454a"
	txevent, err := wm.GetErc20TokenEvent(txid)
	if err != nil {
		t.Errorf("GetErc20TokenEvent error: %v", err)
		return
	}
	log.Infof("txevent: %+v", txevent)
}


func TestFixGasLimit(t *testing.T) {
	fixGasLimitStr := "sfsd"
	fixGasLimit := new(big.Int)
	fixGasLimit.SetString(fixGasLimitStr, 10)
	fmt.Printf("fixGasLimit: %d\n", fixGasLimit.Int64())
}

func TestWalletManager_GetAddrBalance(t *testing.T) {
	wm := testNewWalletManager()
	balance, err := wm.WalletClient.GetAddrBalance2("", "pending")
	if err != nil {
		t.Errorf("GetAddrBalance2 error: %v", err)
		return
	}
	ethB, err := ConverWeiStringToEthDecimal(balance.String())
	log.Infof("ethB: %v", ethB)
}

func TestWalletManager_SetNetworkChainID(t *testing.T) {
	wm := testNewWalletManager()
	id, err := wm.SetNetworkChainID()
	if err != nil {
		t.Errorf("GetAddrBalance2 error: %v", err)
		return
	}
	log.Infof("chainID: %d", id)
}