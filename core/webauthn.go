package core

import (
	"fmt"
	"github.com/dotbitHQ/das-lib/common"
	"github.com/dotbitHQ/das-lib/witness"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
)

func (d *DasCore) GetKeyListCell(args []byte) (*indexer.LiveCell, error) {
	keyListCell, err := GetDasContractInfo(common.DasKeyListCellType)
	if err != nil {
		return nil, fmt.Errorf("GetDasContractInfo err: %s", err.Error())
	}
	dasLock, err := GetDasContractInfo(common.DasContractNameDispatchCellType)
	if err != nil {
		return nil, fmt.Errorf("GetDasContractInfo err: %s", err.Error())
	}
	searchKey := indexer.SearchKey{
		Script:     keyListCell.ToScript(nil),
		ScriptType: indexer.ScriptTypeType,
		ArgsLen:    0,
		Filter: &indexer.CellsFilter{
			Script: dasLock.ToScript(args),
		},
	}

	keyListCells, err := d.client.GetCells(d.ctx, &searchKey, indexer.SearchOrderDesc, 1, "")
	if err != nil {
		return nil, fmt.Errorf("GetCells err: %s", err.Error())
	}

	if subLen := len(keyListCells.Objects); subLen != 1 {
		return nil, nil
	}

	return keyListCells.Objects[0], nil
}

func (d *DasCore) GetIdxOfKeylist(loginAddr, signAddr DasAddressHex) (int, error) {
	var idx int
	if loginAddr.AddressHex == signAddr.AddressHex {
		return 255, nil
	}
	lockArgs, err := d.Daf().HexToArgs(loginAddr, loginAddr)
	KeyListCfgCell, err := d.GetKeyListCell(lockArgs)
	if err != nil {
		return 0, fmt.Errorf("GetKeyListCell(webauthn keyListCell) : %s", err.Error())
	}
	keyListConfigTx, err := d.Client().GetTransaction(d.ctx, KeyListCfgCell.OutPoint.TxHash)
	if err != nil {
		return 0, fmt.Errorf("GetTransaction err: " + err.Error())
	}
	webAuthnKeyListConfigBuilder, err := witness.WebAuthnKeyListDataBuilderFromTx(keyListConfigTx.Transaction, common.DataTypeNew)
	if err != nil {
		return 0, fmt.Errorf("WebAuthnKeyListDataBuilderFromTx err: " + err.Error())
	}
	dataBuilder := webAuthnKeyListConfigBuilder.DeviceKeyListCellData.AsBuilder()
	deviceKeyListCellDataBuilder := dataBuilder.Build()
	keyList := deviceKeyListCellDataBuilder.Keys()
	if keyList == nil {
		return 0, fmt.Errorf("login address status has not enable authorize")
	}

	for i := 0; i < int(keyList.Len()); i++ {
		mainAlgId := common.DasAlgorithmId(keyList.Get(uint(i)).MainAlgId().RawData()[0])
		subAlgId := common.DasSubAlgorithmId(keyList.Get(uint(i)).SubAlgId().RawData()[0])
		cid1 := keyList.Get(uint(i)).Cid().RawData()
		pk1 := keyList.Get(uint(i)).Pubkey().RawData()
		addressHex := common.Bytes2Hex(append(cid1, pk1...))
		if loginAddr.DasAlgorithmId == mainAlgId &&
			loginAddr.DasSubAlgorithmId == subAlgId &&
			addressHex == loginAddr.AddressHex {
			idx = i
			break
		}
	}
	if idx == -1 {
		return -1, nil
	}
	return idx, nil
}