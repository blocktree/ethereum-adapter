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

package openwtester

import (
	"github.com/blocktree/openwallet/v2/openw"
	"testing"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openwallet"
)

func testGetAssetsAccountBalance(tm *openw.WalletManager, walletID, accountID string) {
	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func testGetAssetsAccountTokenBalance(tm *openw.WalletManager, walletID, accountID string, contract openwallet.SmartContract) {
	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("token balance:", balance.Balance)
}

func testCreateTransactionStep(tm *openw.WalletManager, walletID, accountID, to, amount, feeRate string, contract *openwallet.SmartContract, extParam map[string]interface{}) (*openwallet.RawTransaction, error) {

	//err := tm.RefreshAssetsAccountBalance(testApp, accountID)
	//if err != nil {
	//	log.Error("RefreshAssetsAccountBalance failed, unexpected error:", err)
	//	return nil, err
	//}

	rawTx, err := tm.CreateTransaction(testApp, walletID, accountID, amount, to, feeRate, "", contract, extParam)

	if err != nil {
		log.Error("CreateTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTx, nil
}

func testCreateSummaryTransactionStep(
	tm *openw.WalletManager,
	walletID, accountID, summaryAddress, minTransfer, retainedBalance, feeRate string,
	start, limit int,
	contract *openwallet.SmartContract,
	feeSupportAccount *openwallet.FeesSupportAccount) ([]*openwallet.RawTransactionWithError, error) {

	rawTxArray, err := tm.CreateSummaryRawTransactionWithError(testApp, walletID, accountID, summaryAddress, minTransfer,
		retainedBalance, feeRate, start, limit, contract, feeSupportAccount)

	if err != nil {
		log.Error("CreateSummaryTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTxArray, nil
}

func testSignTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	_, err := tm.SignTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, "12345678", rawTx)
	if err != nil {
		log.Error("SignTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testVerifyTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	//log.Info("rawTx.Signatures:", rawTx.Signatures)

	_, err := tm.VerifyTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("VerifyTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testSubmitTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	tx, err := tm.SubmitTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("SubmitTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Std.Info("tx: %+v", tx)
	//log.Info("wxID:", tx.WxID)
	log.Info("txID:", rawTx.TxID)

	return rawTx, nil
}

func TestTransfer_ETH(t *testing.T) {

	addrs := []string{
		"0x1d47898162fdb807c3143078658788bc468af394",
		//"0x6c3e80ec7fb07a79bf610311bd5a72f7955e41cd",
		//"0x80e7ed8adba18afb5d434961db5bbb9c6ca097f6",
		//"0x81f66231abe6047ba70f1c8efb001d6ae31b75eb",
		//"0x932eb86cad71e02af5df3873f5c6f4c57e9a6eeb",
		//"0x938db33857b30b0878d1593d51961468c1ac4d6f",
		//"0x9501d338e3aeda069ef1adb769d5f8eab4d19c28",
		//"0x982c2d7eb62e121f253437dd9eab4441f3498565",
		//"0xb99b2e29ad98f8547114abc09ff53f03cd939ad5",
		//"0xdbbaeb5d88edbdb20e6c95904fb5ffa731373c3c",
		//"0xf43d89450ae388169b8b29ce2991a07f6d747eb8",
		//"0xf650fed6738b573bccdd47360cb03621bba15684",
		//"0xf75f4e7fa0ac8e774276d52428c5892a531af62a",

		//"",
	}

	tm := testInitWalletManager()
	walletID := "WBGYxZ6yEX582Mx8mGvygXevdLVc7NQnLM"
	accountID := "9EfTQiMEaKSMd1CjxMXRMMxukrwckxdBZpiEkS2B3avD"
	//accountID := "AfF8aoW2M2bQwVc2aJ38cCGEcnXF3WCsma1Day7zGA4C"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	//extParam := map[string]interface{}{
	//	"nonce": 20,
	//}

	for _, to := range addrs {
		rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "0.001", "", nil, nil)
		if err != nil {
			return
		}

		log.Std.Info("rawTx: %+v", rawTx)

		_, err = testSignTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTx)
		if err != nil {
			return
		}
	}
}

func TestTransfer_ERC20(t *testing.T) {

	addrs := []string{
		"0x1d47898162fdb807c3143078658788bc468af394",
		"0x6c3e80ec7fb07a79bf610311bd5a72f7955e41cd",
		"0x80e7ed8adba18afb5d434961db5bbb9c6ca097f6",
		"0x81f66231abe6047ba70f1c8efb001d6ae31b75eb",
		"0x932eb86cad71e02af5df3873f5c6f4c57e9a6eeb",
		"0x938db33857b30b0878d1593d51961468c1ac4d6f",
		"0x9501d338e3aeda069ef1adb769d5f8eab4d19c28",
		"0x982c2d7eb62e121f253437dd9eab4441f3498565",
		"0xb99b2e29ad98f8547114abc09ff53f03cd939ad5",
		"0xdbbaeb5d88edbdb20e6c95904fb5ffa731373c3c",
		"0xf43d89450ae388169b8b29ce2991a07f6d747eb8",
		"0xf650fed6738b573bccdd47360cb03621bba15684",
		"0xf75f4e7fa0ac8e774276d52428c5892a531af62a",

		//"0xd35f9ea14d063af9b3567064fab567275b09f03d",
	}

	tm := testInitWalletManager()
	walletID := "WBGYxZ6yEX582Mx8mGvygXevdLVc7NQnLM"
	accountID := "9EfTQiMEaKSMd1CjxMXRMMxukrwckxdBZpiEkS2B3avD"
	//accountID := "AfF8aoW2M2bQwVc2aJ38cCGEcnXF3WCsma1Day7zGA4C"

	contract := openwallet.SmartContract{
		Address:  "0x4092678e4e78230f46a1534c0fbc8fa39780892b",
		Symbol:   "ETH",
		Name:     "OCoin",
		Token:    "OCN",
		Decimals: 18,
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	for _, to := range addrs {
		rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "0.2", "", &contract, nil)
		if err != nil {
			return
		}

		_, err = testSignTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTx)
		if err != nil {
			return
		}
	}
}

func TestSummary_ETH(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WBGYxZ6yEX582Mx8mGvygXevdLVc7NQnLM"
	//accountID := "9EfTQiMEaKSMd1CjxMXRMMxukrwckxdBZpiEkS2B3avD"
	accountID := "CxE3ds4JdTHXV1f2xSsE6qahgfReKR9iPmFPcBmTfaKP"
	summaryAddress := "0xbb0d592280f170069821bcae6c861a5686b77c43"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "",
		0, 100, nil, nil)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//执行汇总交易
	for _, rawTxWithErr := range rawTxArray {

		if rawTxWithErr.Error != nil {
			log.Error(rawTxWithErr.Error.Error())
			continue
		}

		_, err = testSignTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}
	}

}

func TestSummary_ERC20(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WBGYxZ6yEX582Mx8mGvygXevdLVc7NQnLM"
	accountID := "CxE3ds4JdTHXV1f2xSsE6qahgfReKR9iPmFPcBmTfaKP"
	summaryAddress := "0xbb0d592280f170069821bcae6c861a5686b77c43"

	feesSupport := openwallet.FeesSupportAccount{
		AccountID: "9EfTQiMEaKSMd1CjxMXRMMxukrwckxdBZpiEkS2B3avD",
		//FixSupportAmount: "0.01",
		FeesSupportScale: "2",
	}

	contract := openwallet.SmartContract{
		Address:  "0x4092678e4e78230f46a1534c0fbc8fa39780892b",
		Symbol:   "ETH",
		Name:     "OCoin",
		Token:    "OCN",
		Decimals: 18,
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	list, err := tm.GetAddressList(testApp, walletID, accountID, 0, -1, false)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}

	addressLimit := 3

	//分页汇总交易
	for i := 0; i < len(list); i = i + addressLimit {
		rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
			summaryAddress, "", "", "",
			i, addressLimit, &contract, &feesSupport)
		if err != nil {
			log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
			return
		}

		//执行汇总交易
		for _, rawTxWithErr := range rawTxArray {

			if rawTxWithErr.Error != nil {
				log.Error(rawTxWithErr.Error.Error())
				continue
			}

			_, err = testSignTransactionStep(tm, rawTxWithErr.RawTx)
			if err != nil {
				return
			}

			_, err = testVerifyTransactionStep(tm, rawTxWithErr.RawTx)
			if err != nil {
				return
			}

			_, err = testSubmitTransactionStep(tm, rawTxWithErr.RawTx)
			if err != nil {
				return
			}
		}
	}

}
