package ethereum_txsigner

import (
	"fmt"
	"github.com/blocktree/go-owcrypt"
)

var Default = &TransactionSigner{}

type TransactionSigner struct {

}

// SignTransactionHash 交易哈希签名算法
// required
func (singer *TransactionSigner) SignTransactionHash(msg []byte, privateKey []byte, eccType uint32) ([]byte, error) {

	signature, v, sigErr := owcrypt.Signature(privateKey, nil, msg, eccType)
	if sigErr != owcrypt.SUCCESS {
		return nil, fmt.Errorf("transaction hash sign failed")
	}
	signature = append(signature, v)
	return signature, nil
}