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
	"github.com/blocktree/openwallet/v2/openwallet"
	"testing"
)

func TestWalletManager_GetTokenBalanceByAddress(t *testing.T) {
	wm := testNewWalletManager()

	contract := openwallet.SmartContract{
		Address:  "0xc1b3702b8b2778cd8315efd40ca6c1addf094c80",
		Symbol:   "ETH",
		Name:     "OCoin",
		Token:    "OCN",
		Decimals: 18,
	}

	tokenBalances, err := wm.ContractDecoder.GetTokenBalanceByAddress(contract, "0x373944b86Bc07887f2cdcC5EF5E055ee33AC2d3C", "0xfa23640bd91618d7fe79934d0eaaa747fb801fae")
	if err != nil {
		t.Errorf("GetTokenBalanceByAddress unexpected error: %v", err)
		return
	}
	for _, b := range tokenBalances {
		t.Logf("token balance: %+v", b.Balance)
	}
}
