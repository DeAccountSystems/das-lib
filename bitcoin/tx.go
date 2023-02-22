package bitcoin

import (
	"context"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"math"
)

type TxTool struct {
	RpcClient        *BaseRequest
	Ctx              context.Context
	RemoteSignClient rpc.Client

	DustLimit DustLimit
	Params    chaincfg.Params
}

var (
	InsufficientBalanceError = errors.New("InsufficientBalanceError")
)

type DustLimit = int64

const (
	signSize                = 107
	outSize                 = 35 //P2PKH
	DustLimitBtc  DustLimit = 546
	DustLimitBch  DustLimit = 546
	DustLimitLtc  DustLimit = 5460
	DustLimitDoge DustLimit = 100000000
)

func (t *TxTool) NewTx(uos []UnspentOutputs, addresses []string, values []int64) (*wire.MsgTx, error) {
	if len(uos) == 0 || (len(addresses) != len(values)) {
		return nil, fmt.Errorf("param is invalid:%v,%v,%v", uos, addresses, values)
	}

	// get fee
	var fee float64
	err := t.RpcClient.Request(RpcMethodEstimateFee, []interface{}{10}, &fee)
	if err != nil {
		return nil, fmt.Errorf("req RpcMethodEstimateFee err: %s", err.Error())
	}
	txFee := int64(math.Pow10(8) * fee)
	log.Warn("NewTx fee:", fee, txFee)

	// new tx
	tx := wire.NewMsgTx(wire.TxVersion)
	var inTotal, outTotal int64

	// inputs
	for _, utxo := range uos {
		in, err := t.newTxIn(utxo.Hash, utxo.Index)
		if err != nil {
			return nil, fmt.Errorf("newTxIn err: %s", err.Error())
		}
		tx.AddTxIn(in)
		inTotal += utxo.Value
	}

	// output
	for i := range addresses {
		if values[i] < t.DustLimit {
			return nil, fmt.Errorf("the output value:%v is must bigger than:%v", values[i], t.DustLimit)
		}
		out1, err := t.newTxOut(addresses[i], values[i])
		if err != nil {
			return nil, err
		}
		tx.AddTxOut(out1)
		outTotal += values[i]
	}

	// change
	signSizeTmp := len(tx.TxIn) * signSize
	feeValue := (txFee * int64(tx.SerializeSize()+signSizeTmp)) / 1000
	charge := inTotal - outTotal - feeValue
	log.Warn("NewTx:", inTotal, outTotal, feeValue, charge)
	if charge < 0 {
		return nil, InsufficientBalanceError
	}

	feeValue = (txFee * int64(tx.SerializeSize()+signSizeTmp+outSize)) / 1000
	charge = inTotal - outTotal - feeValue
	if charge >= t.DustLimit {
		outCharge, err := t.newTxOut(uos[0].Address, charge)
		if err != nil {
			return nil, err
		}
		tx.AddTxOut(outCharge)
	}

	return tx, nil
}

type UnspentOutputs struct {
	Private string `json:"private"`
	Address string `json:"address"`
	Hash    string `json:"hash"`
	Index   uint32 `json:"index"`
	Value   int64  `json:"value"`
}

func (t *TxTool) newTxIn(hashStr string, index uint32) (*wire.TxIn, error) {
	hash, err := chainhash.NewHashFromStr(hashStr)
	if err != nil {
		return nil, fmt.Errorf("NewHashFromStr err: %s", err.Error())
	}
	outPoint := wire.NewOutPoint(hash, index)
	return wire.NewTxIn(outPoint, nil, nil), nil
}

func (t *TxTool) newTxOut(addr string, value int64) (*wire.TxOut, error) {
	decodeAddress, err := btcutil.DecodeAddress(addr, &t.Params)
	if err != nil {
		return nil, fmt.Errorf("DecodeAddress err: %s", err)
	}
	script, err := txscript.PayToAddrScript(decodeAddress)
	if err != nil {
		return nil, fmt.Errorf("PayToAddrScript err: %v", err)
	}
	return wire.NewTxOut(value, script), nil
}

func (t *TxTool) SendTx(tx *wire.MsgTx) (hash string, err error) {
	if tx == nil {
		return "", fmt.Errorf("tx is nil")
	}
	raw, err := txToString(tx)
	if err != nil {
		return "", fmt.Errorf("txToString err: %s", err.Error())
	}

	params := []interface{}{raw, false}
	err = t.RpcClient.Request(RpcMethodSendRawTransaction, params, &hash)
	if err != nil {
		return "", fmt.Errorf("send tx req err: %s", err.Error())
	}

	return hash, nil
}
