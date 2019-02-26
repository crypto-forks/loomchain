//+build !evm

package gateway

import (
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/fnConsensus"
)

type BatchSignWithdrawalFn struct {
}

func (b *BatchSignWithdrawalFn) PrepareContext() (bool, []byte, error) {
	return false, nil, nil
}

func (b *BatchSignWithdrawalFn) SubmitMultiSignedMessage(ctx []byte, key []byte, signatures [][]byte) {

}

func (b *BatchSignWithdrawalFn) GetMessageAndSignature(ctx []byte) ([]byte, []byte, error) {
	return nil, nil, nil
}

func (b *BatchSignWithdrawalFn) MapMessage(ctx, key, message []byte) error {
	return nil
}

func CreateBatchSignWithdrawalFn(isLoomcoinFn bool, chainID string, fnRegistry fnConsensus.FnRegistry, tgConfig *TransferGatewayConfig) (*BatchSignWithdrawalFn, error) {
	return nil, errors.New("not implemented in non-EVM build")
}
