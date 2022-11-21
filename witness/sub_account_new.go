package witness

import (
	"fmt"
	"github.com/dotbitHQ/das-lib/common"
	"github.com/dotbitHQ/das-lib/molecule"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type SubAccountBuilderNew struct{}

// === SubAccountMintSign ===

type SubAccountMintSignVersion = uint32

const (
	SubAccountMintSignVersion1 SubAccountMintSignVersion = 1
)

type SubAccountMintSign struct {
	versionBys          []byte
	expiredTimestampBys []byte

	Version            SubAccountMintSignVersion
	Signature          []byte
	ExpiredTimestamp   uint32
	AccountListSmtRoot []byte
}

func (s *SubAccountBuilderNew) ConvertSubAccountMintSignFromBytes(dataBys []byte) (*SubAccountMintSign, error) {
	var res SubAccountMintSign
	index, indexLen, dataLen := uint32(0), uint32(4), uint32(0)

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.versionBys = dataBys[index+indexLen : index+indexLen+dataLen]
	res.Version, _ = molecule.Bytes2GoU32(res.versionBys)
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.Signature = dataBys[index+indexLen : index+indexLen+dataLen]
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.expiredTimestampBys = dataBys[index+indexLen : index+indexLen+dataLen]
	res.ExpiredTimestamp, _ = molecule.Bytes2GoU32(res.expiredTimestampBys)
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.AccountListSmtRoot = dataBys[index+indexLen : index+indexLen+dataLen]
	index = index + indexLen + dataLen

	return &res, nil
}
func (s *SubAccountBuilderNew) GenSubAccountMintSignBytes(p SubAccountMintSign) (dataBys []byte) {
	versionBys := molecule.GoU32ToMoleculeU32(p.Version)
	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(versionBys.RawData())))...)
	dataBys = append(dataBys, versionBys.RawData()...)

	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(p.Signature)))...)
	dataBys = append(dataBys, p.Signature...)

	expiredTimestampBys := molecule.GoU32ToMoleculeU32(p.ExpiredTimestamp)
	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(expiredTimestampBys.RawData())))...)
	dataBys = append(dataBys, expiredTimestampBys.RawData()...)

	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(p.AccountListSmtRoot)))...)
	dataBys = append(dataBys, p.Signature...)

	return
}

// === SubAccountNew ===

type SubAccountNewVersion = uint32

const (
	SubAccountNewVersion1 SubAccountNewVersion = 1
	SubAccountNewVersion2 SubAccountNewVersion = 2
)

type SubAccountNew struct {
	// v2
	Version           SubAccountNewVersion
	versionBys        []byte
	Signature         []byte
	SignRole          []byte
	NewRoot           []byte
	Proof             []byte
	Action            string
	actionBys         []byte
	SubAccountData    *SubAccountData
	subAccountDataBys []byte
	EditKey           string
	editKeyBys        []byte
	EditValue         []byte
	//
	EditLockArgs          []byte
	EditRecords           []Record
	RenewExpiredAt        uint64
	CurrentSubAccountData *SubAccountData
	// v1
	PrevRoot    []byte
	CurrentRoot []byte
}

func (s *SubAccountBuilderNew) genSubAccountNewBytesV1(p SubAccountNew) (dataBys []byte, err error) {
	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(p.Signature)))...)
	dataBys = append(dataBys, p.Signature...)

	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(p.SignRole)))...)
	dataBys = append(dataBys, p.SignRole...)

	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(p.PrevRoot)))...)
	dataBys = append(dataBys, p.PrevRoot...)

	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(p.CurrentRoot)))...)
	dataBys = append(dataBys, p.CurrentRoot...)

	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(p.Proof)))...)
	dataBys = append(dataBys, p.Proof...)

	versionBys := molecule.GoU32ToMoleculeU32(SubAccountCurrentVersion)
	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(versionBys.RawData())))...)
	dataBys = append(dataBys, versionBys.RawData()...)

	if p.SubAccountData == nil {
		return nil, fmt.Errorf("SubAccountData is nil")
	}
	subAccountData, err := p.SubAccountData.ConvertToMoleculeSubAccount()
	if err != nil {
		return nil, fmt.Errorf("ConvertToMoleculeSubAccount err: %s", err.Error())
	}
	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(subAccountData.AsSlice())))...)
	dataBys = append(dataBys, subAccountData.AsSlice()...)

	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len([]byte(p.EditKey))))...)
	dataBys = append(dataBys, p.EditKey...)

	var editValue []byte
	switch p.EditKey {
	case common.EditKeyOwner, common.EditKeyManager:
		editValue = p.EditLockArgs
	case common.EditKeyRecords:
		records := ConvertToCellRecords(p.EditRecords)
		editValue = records.AsSlice()
	case common.EditKeyExpiredAt:
		expiredAt := molecule.GoU64ToMoleculeU64(p.RenewExpiredAt)
		editValue = expiredAt.AsSlice()
	}

	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(editValue)))...)
	dataBys = append(dataBys, editValue...)
	return
}
func (s *SubAccountBuilderNew) genSubAccountNewBytesV2(p SubAccountNew) (dataBys []byte, err error) {
	versionBys := molecule.GoU32ToMoleculeU32(p.Version)
	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(versionBys.RawData())))...)
	dataBys = append(dataBys, versionBys.RawData()...)

	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(p.Signature)))...)
	dataBys = append(dataBys, p.Signature...)

	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(p.SignRole)))...)
	dataBys = append(dataBys, p.SignRole...)

	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(p.NewRoot)))...)
	dataBys = append(dataBys, p.NewRoot...)

	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(p.Proof)))...)
	dataBys = append(dataBys, p.Proof...)

	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len([]byte(p.Action))))...)
	dataBys = append(dataBys, p.Action...)

	if p.SubAccountData == nil {
		return nil, fmt.Errorf("SubAccountData is nil")
	}
	subAccountData, err := p.SubAccountData.ConvertToMoleculeSubAccount()
	if err != nil {
		return nil, fmt.Errorf("ConvertToMoleculeSubAccount err: %s", err.Error())
	}
	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(subAccountData.AsSlice())))...)
	dataBys = append(dataBys, subAccountData.AsSlice()...)

	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len([]byte(p.EditKey))))...)
	dataBys = append(dataBys, p.EditKey...)

	var editValue []byte
	switch p.EditKey {
	case common.EditKeyOwner, common.EditKeyManager:
		editValue = p.EditLockArgs
	case common.EditKeyRecords:
		records := ConvertToCellRecords(p.EditRecords)
		editValue = records.AsSlice()
	case common.EditKeyExpiredAt:
		expiredAt := molecule.GoU64ToMoleculeU64(p.RenewExpiredAt)
		editValue = expiredAt.AsSlice()
	}

	dataBys = append(dataBys, molecule.GoU32ToBytes(uint32(len(editValue)))...)
	dataBys = append(dataBys, editValue...)

	return
}
func (s *SubAccountBuilderNew) GenSubAccountNewBytes(p SubAccountNew) (dataBys []byte, err error) {
	if p.Version == SubAccountNewVersion2 {
		return s.genSubAccountNewBytesV2(p)
	}
	return s.genSubAccountNewBytesV1(p)
}
func (s *SubAccountBuilderNew) convertSubAccountNewFromBytesV1(dataBys []byte) (*SubAccountNew, error) {
	var res SubAccountNew
	index, indexLen, dataLen := uint32(0), uint32(4), uint32(0)

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.Signature = dataBys[index+indexLen : index+indexLen+dataLen]
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.SignRole = dataBys[index+indexLen : index+indexLen+dataLen]
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.PrevRoot = dataBys[index+indexLen : index+indexLen+dataLen]
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.CurrentRoot = dataBys[index+indexLen : index+indexLen+dataLen]
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.Proof = dataBys[index+indexLen : index+indexLen+dataLen]
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.versionBys = dataBys[index+indexLen : index+indexLen+dataLen]
	res.Version, _ = molecule.Bytes2GoU32(res.versionBys)
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.subAccountDataBys = dataBys[index+indexLen : index+indexLen+dataLen]
	switch res.Version {
	default:
		subAccount, err := s.ConvertSubAccountDataFromBytes(res.subAccountDataBys)
		if err != nil {
			return nil, fmt.Errorf("ConvertToSubAccount err: %s", err.Error())
		}
		res.SubAccountData = subAccount
	}
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.editKeyBys = dataBys[index+indexLen : index+indexLen+dataLen]
	res.EditKey = string(res.editKeyBys)
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.EditValue = dataBys[index+indexLen : index+indexLen+dataLen]
	res.ConvertCurrentSubAccountData()
	index = index + indexLen + dataLen

	return &res, nil
}
func (s *SubAccountBuilderNew) convertSubAccountNewFromBytesV2(dataBys []byte) (*SubAccountNew, error) {
	var res SubAccountNew
	index, indexLen, dataLen := uint32(0), uint32(4), uint32(0)

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.versionBys = dataBys[index+indexLen : index+indexLen+dataLen]
	res.Version, _ = molecule.Bytes2GoU32(res.versionBys)
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.Signature = dataBys[index+indexLen : index+indexLen+dataLen]
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.SignRole = dataBys[index+indexLen : index+indexLen+dataLen]
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.NewRoot = dataBys[index+indexLen : index+indexLen+dataLen]
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.Proof = dataBys[index+indexLen : index+indexLen+dataLen]
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.actionBys = dataBys[index+indexLen : index+indexLen+dataLen]
	res.Action = string(res.actionBys)
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.subAccountDataBys = dataBys[index+indexLen : index+indexLen+dataLen]
	switch res.Version {
	default:
		subAccount, err := s.ConvertSubAccountDataFromBytes(res.subAccountDataBys)
		if err != nil {
			return nil, fmt.Errorf("ConvertToSubAccount err: %s", err.Error())
		}
		res.SubAccountData = subAccount
	}
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.editKeyBys = dataBys[index+indexLen : index+indexLen+dataLen]
	res.EditKey = string(res.editKeyBys)
	index = index + indexLen + dataLen

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	res.EditValue = dataBys[index+indexLen : index+indexLen+dataLen]
	res.ConvertCurrentSubAccountData()
	index = index + indexLen + dataLen

	return &res, nil
}
func (s *SubAccountBuilderNew) ConvertSubAccountNewFromBytes(dataBys []byte) (*SubAccountNew, error) {
	index, indexLen, dataLen := uint32(0), uint32(4), uint32(0)

	dataLen, _ = molecule.Bytes2GoU32(dataBys[index : index+indexLen])
	if dataLen == 4 {
		return s.convertSubAccountNewFromBytesV2(dataBys)
	} else {
		return s.convertSubAccountNewFromBytesV1(dataBys)
	}
}
func (s *SubAccountBuilderNew) SubAccountNewMapFromTx(tx *types.Transaction) (map[string]*SubAccountNew, error) {
	var respMap = make(map[string]*SubAccountNew)

	err := GetWitnessDataFromTx(tx, func(actionDataType common.ActionDataType, dataBys []byte) (bool, error) {
		switch actionDataType {
		case common.ActionDataTypeSubAccount:
			subAccountNew, err := s.ConvertSubAccountNewFromBytes(dataBys)
			if err != nil {
				return false, err
			}
			respMap[subAccountNew.SubAccountData.AccountId] = subAccountNew
		}
		return true, nil
	})

	if err != nil {
		return nil, fmt.Errorf("GetWitnessDataFromTx err: %s", err.Error())
	}
	if len(respMap) == 0 {
		return nil, fmt.Errorf("not exist sub account")
	}
	return respMap, nil
}

// === EditValue ===
func (s *SubAccountNew) ConvertCurrentSubAccountData() {
	currentSubAccountData := *s.SubAccountData
	s.CurrentSubAccountData = &currentSubAccountData

	if s.EditKey != "" {
		s.CurrentSubAccountData.Nonce++
	}
	switch s.EditKey {
	case common.EditKeyOwner:
		s.CurrentSubAccountData.Lock = &types.Script{
			CodeHash: s.CurrentSubAccountData.Lock.CodeHash,
			HashType: s.CurrentSubAccountData.Lock.HashType,
			Args:     s.EditValue,
		}
		s.EditLockArgs = s.EditValue
		s.CurrentSubAccountData.Records = nil
	case common.EditKeyManager:
		s.CurrentSubAccountData.Lock = &types.Script{
			CodeHash: s.CurrentSubAccountData.Lock.CodeHash,
			HashType: s.CurrentSubAccountData.Lock.HashType,
			Args:     s.EditValue,
		}
		s.EditLockArgs = s.EditValue
	case common.EditKeyRecords:
		records, _ := molecule.RecordsFromSlice(s.EditValue, true)
		s.EditRecords = ConvertToRecords(records)
		s.CurrentSubAccountData.Records = s.EditRecords
	case common.EditKeyExpiredAt:
		expiredAt, _ := molecule.Uint64FromSlice(s.EditValue, true)
		s.RenewExpiredAt, _ = molecule.Bytes2GoU64(expiredAt.RawData())
		s.CurrentSubAccountData.ExpiredAt = s.RenewExpiredAt
	}
}

// === SubAccountData ===
type SubAccountData struct {
	Lock                 *types.Script           `json:"lock"`
	AccountId            string                  `json:"account_id"`
	AccountCharSet       []common.AccountCharSet `json:"account_char_set"`
	Suffix               string                  `json:"suffix"`
	RegisteredAt         uint64                  `json:"registered_at"`
	ExpiredAt            uint64                  `json:"expired_at"`
	Status               uint8                   `json:"status"`
	Records              []Record                `json:"records"`
	Nonce                uint64                  `json:"nonce"`
	EnableSubAccount     uint8                   `json:"enable_sub_account"`
	RenewSubAccountPrice uint64                  `json:"renew_sub_account_price"`
}

func (s *SubAccountBuilderNew) ConvertSubAccountDataFromBytes(dataBys []byte) (*SubAccountData, error) {
	subAccount, err := molecule.SubAccountFromSlice(dataBys, true)
	if err != nil {
		return nil, fmt.Errorf("SubAccountDataFromSlice err: %s", err.Error())
	}
	var tmp SubAccountData
	tmp.Lock = molecule.MoleculeScript2CkbScript(subAccount.Lock())
	tmp.AccountId = common.Bytes2Hex(subAccount.Id().RawData())
	tmp.AccountCharSet = common.ConvertToAccountCharSets(subAccount.Account())
	tmp.Suffix = string(subAccount.Suffix().RawData())
	tmp.RegisteredAt, _ = molecule.Bytes2GoU64(subAccount.RegisteredAt().RawData())
	tmp.ExpiredAt, _ = molecule.Bytes2GoU64(subAccount.ExpiredAt().RawData())
	tmp.Status, _ = molecule.Bytes2GoU8(subAccount.Status().RawData())
	tmp.Records = ConvertToRecords(subAccount.Records())
	tmp.Nonce, _ = molecule.Bytes2GoU64(subAccount.Nonce().RawData())
	tmp.EnableSubAccount, _ = molecule.Bytes2GoU8(subAccount.EnableSubAccount().RawData())
	tmp.RenewSubAccountPrice, _ = molecule.Bytes2GoU64(subAccount.RenewSubAccountPrice().RawData())

	return &tmp, nil
}
func (s *SubAccountData) ConvertToMoleculeSubAccount() (*molecule.SubAccount, error) {
	if s.Lock == nil {
		return nil, fmt.Errorf("lock is nil")
	}
	lock := molecule.CkbScript2MoleculeScript(s.Lock)
	accountChars := common.ConvertToAccountChars(s.AccountCharSet)
	accountId, err := molecule.AccountIdFromSlice(common.Hex2Bytes(s.AccountId), true)
	if err != nil {
		return nil, fmt.Errorf("AccountIdFromSlice err: %s", err.Error())
	}
	suffix := molecule.GoBytes2MoleculeBytes([]byte(s.Suffix))
	registeredAt := molecule.GoU64ToMoleculeU64(s.RegisteredAt)
	expiredAt := molecule.GoU64ToMoleculeU64(s.ExpiredAt)
	status := molecule.GoU8ToMoleculeU8(s.Status)
	records := ConvertToCellRecords(s.Records)
	nonce := molecule.GoU64ToMoleculeU64(s.Nonce)
	enableSubAccount := molecule.GoU8ToMoleculeU8(s.EnableSubAccount)
	renewSubAccountPrice := molecule.GoU64ToMoleculeU64(s.RenewSubAccountPrice)

	moleculeSubAccount := molecule.NewSubAccountBuilder().
		Lock(lock).
		Id(*accountId).
		Account(*accountChars).
		Suffix(suffix).
		RegisteredAt(registeredAt).
		ExpiredAt(expiredAt).
		Status(status).
		Records(*records).
		Nonce(nonce).
		EnableSubAccount(enableSubAccount).
		RenewSubAccountPrice(renewSubAccountPrice).
		Build()
	return &moleculeSubAccount, nil
}
func (s *SubAccountData) Account() string {
	var account string
	for _, v := range s.AccountCharSet {
		account += v.Char
	}
	return account + s.Suffix
}
func (s *SubAccountData) ToH256() ([]byte, error) {
	moleculeSubAccount, err := s.ConvertToMoleculeSubAccount()
	if err != nil {
		return nil, fmt.Errorf("ConvertToMoleculeSubAccount err: %s", err.Error())
	}
	return blake2b.Blake256(moleculeSubAccount.AsSlice())
}
