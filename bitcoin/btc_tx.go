package bitcoin

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"math"
)

func (t *TxTool) NewBTCTx(uos []UnspentOutputs, addresses []string, values []int64, opReturn string) (*wire.MsgTx, error) {
	if len(uos) == 0 || (len(addresses) != len(values)) {
		return nil, fmt.Errorf("param is invalid:%v,%v,%v", uos, addresses, values)
	}

	// get fee
	feeResult, err := t.RpcClientBTC.EstimateSmartFee(10, &btcjson.EstimateModeConservative)
	if err != nil {
		return nil, fmt.Errorf("EstimateFee err: %s", err.Error())
	}
	fee := *feeResult.FeeRate
	txFee := int64(math.Pow10(8) * fee)
	log.Warn("NewBTCTx fee:", fee, txFee)

	// new tx
	tx := wire.NewMsgTx(wire.TxVersion)
	var inTotal, outTotal int64

	// inputs
	for _, utxo := range uos {
		//fmt.Println("uos:", utxo.Hash, utxo.Address, utxo.Index)
		txIn, err := t.newTxIn(utxo.Hash, utxo.Index)
		if err != nil {
			return nil, fmt.Errorf("newTxIn err: %s", err.Error())
		}
		tx.AddTxIn(txIn)
		inTotal += utxo.Value
	}

	// output
	for i := range addresses {
		if values[i] < t.DustLimit {
			return nil, fmt.Errorf("the output value:%v is must bigger than:%v", values[i], t.DustLimit)
		}
		txOut, err := t.newTxOut(addresses[i], values[i])
		if err != nil {
			return nil, err
		}
		tx.AddTxOut(txOut)
		outTotal += values[i]
	}

	// fee and change
	feeValue := (txFee * int64(len(tx.TxIn)*148+len(tx.TxOut)*34+34+10)) / 1000
	charge := inTotal - outTotal - feeValue
	log.Warn("NewBTCTx:", inTotal, outTotal, feeValue, charge)
	if charge < 0 {
		return nil, InsufficientBalanceError
	}

	if charge >= t.DustLimit {
		outCharge, err := t.newTxOut(uos[0].Address, charge)
		if err != nil {
			return nil, err
		}
		tx.AddTxOut(outCharge)
	}

	// op_return
	if opReturn != "" {
		op, err := newOpReturn(opReturn)
		if err != nil {
			return nil, fmt.Errorf("newOpReturn err: %s", err.Error())
		}
		tx.AddTxOut(op)
	}

	return tx, nil
}

func (t *TxTool) SendBTCTx(tx *wire.MsgTx) (string, error) {
	if tx == nil {
		return "", fmt.Errorf("tx is nil")
	}
	buf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
	if err := tx.Serialize(buf); err != nil {
		return "", fmt.Errorf("tx.Serialize err: %s", err.Error())
	}
	txHex := hex.EncodeToString(buf.Bytes())
	maxFeeRate := int32(btcutil.SatoshiPerBitcoin / 10)
	cmd := btcjson.NewBitcoindSendRawTransactionCmd(txHex, maxFeeRate)
	var res rpcclient.FutureSendRawTransactionResult
	res = t.RpcClientBTC.SendCmd(cmd)
	hash, err := res.Receive()

	//hash, err := t.RpcClientBTC.SendRawTransaction(tx, true)
	if err != nil {
		return "", fmt.Errorf("SendRawTransaction err: %s", err.Error())
	}

	return hash.String(), nil
}