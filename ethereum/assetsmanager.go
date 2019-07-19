package ethereum

import (
	"github.com/blocktree/openwallet/openwallet"
)

//GetAddressDecode 地址解析器
func (this *WalletManager) GetAddressDecode() openwallet.AddressDecoder {
	return this.Decoder
}

//GetTransactionDecoder 交易单解析器
func (this *WalletManager) GetTransactionDecoder() openwallet.TransactionDecoder {
	return this.TxDecoder
}

func (this *WalletManager) GetSmartContractDecoder() openwallet.SmartContractDecoder {
	return this.ContractDecoder
}

//GetBlockScanner 获取区块链
func (this *WalletManager) GetBlockScanner() openwallet.BlockScanner {
	//先加载是否有配置文件
	//err := this.loadConfig()
	//if err != nil {
	//	log.Errorf("load config failed, err=%v", err)
	//	return nil
	//}
	return this.Blockscanner
}

//ImportWatchOnlyAddress 导入观测地址
func (this *WalletManager) ImportWatchOnlyAddress(address ...*openwallet.Address) error {
	return nil
}

//CurveType 曲线类型
func (this *WalletManager) CurveType() uint32 {
	return this.Config.CurveType
}

//FullName 币种全名
func (this *WalletManager) FullName() string {
	return "Ethereum"
}

//SymbolID 币种标识
func (this *WalletManager) Symbol() string {
	return this.Config.Symbol
}

//小数位精度
func (this *WalletManager) Decimal() int32 {
	return 18
}
