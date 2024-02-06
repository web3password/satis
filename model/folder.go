/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/

package model

import (
	"fmt"
	"time"

	"github.com/web3password/jewel/tools"
	"github.com/web3password/satis/util"
)

type ShareFolderRsp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}
type ShareFolderParams struct {
	Address   string `json:"addr"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Token     string `json:"token"`
	Hash      string `json:"hash"`
}

type ShareFolderUpdateParams struct {
	ShareFolderParams
}

//// MemberList .
//type MemberList struct {
//	List []MemberInfo `json:"list" bson:"list"`
//}

// MemberInfo .
type MemberInfo []struct {
	MemberAddr     string `json:"member_addr" bson:"member_addr"`
	MemberName     []byte `json:"member_name" bson:"member_name"`
	MemberSign     string `json:"member_sign" bson:"member_sign"`
	FolderId       string `json:"folder_id" bson:"folder_id"`
	FolderMnemonic []byte `json:"folder_mnemonic" bson:"folder_mnemonic"`
}

// MemberSign .
//type MemberSign struct {
//	Signature        string `json:"signature" bson:"signature"`
//	MemberSignParams string `json:"member_sign_params" bson:"member_sign_params"`
//}
//
//// MemberSignParams .
//type MemberSignParams struct {
//	PrimaryAddress string      `json:"primary_address" bson:"primary_address"`
//	PublicKeyList  []PublicKey `json:"public_key_list" bson:"public_key_list"`
//	Timestamp      int64       `json:"timestamp" bson:"timestamp"`
//	Nonce          string      `json:"nonce" bson:"nonce"`
//}

// PublicKey .
type PublicKey struct {
	PublicKey string `json:"public_key" bson:"public_key"`
	ID        int32  `json:"id" bson:"id"`
}

func (f ShareFolderParams) Check(data []byte) (string, bool) {
	if !tools.IsValidAddress(f.Address) {
		return "invalid address " + f.Address, false
	}

	if f.Token != ShareFolderCreateToken {
		return "invalid token " + f.Token, false
	}

	if !util.CheckTimestamp(f.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), f.Timestamp)
		return msg, false
	}

	if len(data) == 0 {
		return "invalid empty data ", false
	}

	if isEqual, serverHash := tools.CompareHash(data, f.Hash); !isEqual {
		return fmt.Sprintf("invalid hash, server hash=%s, client hash=%s", serverHash, f.Hash), false
	}

	return "", true
}

func (f ShareFolderUpdateParams) Check(data []byte) (string, bool) {
	if f.Token != ShareFolderUpdateToken {
		return "invalid token " + f.Token, false
	}
	return "", true
}

type ShareFolderCommonParams struct {
	Address   string `json:"addr"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Token     string `json:"token"`
	Hash      string `json:"hash"`
}

func (s ShareFolderCommonParams) Check(token string, timestamp int64, data []byte) (string, bool) {
	nowTimestamp := time.Now().Unix()

	if !tools.IsValidAddress(s.Address) {
		return "invalid address " + s.Address, false
	}
	if s.Timestamp == 0 {
		msg := fmt.Sprintf("invalid timestamp , req timestamp is zero : %d", s.Timestamp)
		return msg, false
	}
	if timestamp > nowTimestamp {
		msg := fmt.Sprintf("invalid timestamp(client time>system time) , server timestamp=%d, client timestamp=%d", nowTimestamp, s.Timestamp)
		return msg, false
	}
	if nowTimestamp-timestamp > 12 {
		msg := fmt.Sprintf("invalid timestamp(client timeout, above 5s) , server timestamp=%d, client timestamp=%d", nowTimestamp, s.Timestamp)
		return msg, false
	}
	if s.Token != token {
		return "invalid token " + s.Token, false
	}

	if len(data) > 0 {
		if isEqual, serverHash := tools.CompareHash(data, s.Hash); !isEqual {
			return fmt.Sprintf("invalid hash, server hash=%s, client hash=%s", serverHash, s.Hash), false
		}
	}

	return "", true
}

type ShareFolderAddMemberDataReq struct {
	MemberAddr     string `bson:"member_addr"`
	MemberSign     string `bson:"member_sign"`
	FolderId       string `bson:"folder_id"`
	FolderMnemonic []byte `bson:"folder_mnemonic"`
	MemberName     []byte `bson:"member_name"`
}

func (s ShareFolderAddMemberDataReq) Check() bool {
	if s.MemberAddr == "" {
		return false
	}
	if s.FolderId == "" {
		return false
	}
	if s.FolderId == "" {
		return false
	}
	if len(s.FolderMnemonic) <= 0 {
		return false
	}
	if len(s.MemberName) <= 0 {
		return false
	}
	return true
}

type ShareFolderAddRecordParams struct {
	Address   string `json:"addr"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Token     string `json:"token"`
	Hash      string `json:"hash"`
}

func (f ShareFolderAddRecordParams) Check(data []byte) (string, bool) {
	if !tools.IsValidAddress(f.Address) {
		return "invalid address " + f.Address, false
	}

	if f.Token != ShareFolderAddRecordToken {
		return "invalid token " + f.Token, false
	}

	if !util.CheckTimestamp(f.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), f.Timestamp)
		return msg, false
	}

	if len(data) == 0 {
		return "invalid empty data ", false
	}

	if isEqual, serverHash := tools.CompareHash(data, f.Hash); !isEqual {
		return fmt.Sprintf("invalid hash, server hash=%s, client hash=%s", serverHash, f.Hash), false
	}

	return "", true
}

type ShareFolderDeleteRecordParams struct {
	Address   string `json:"addr"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Token     string `json:"token"`
	Hash      string `json:"hash"`
	FolderId  string `json:"folder_id"`
	RecordId  string `json:"record_id"`
}

func (f ShareFolderDeleteRecordParams) Check(data []byte) (string, bool) {
	if !tools.IsValidAddress(f.Address) {
		return "invalid address " + f.Address, false
	}

	if f.Token != ShareFolderDeleteRecordToken {
		return "invalid token " + f.Token, false
	}

	if !util.CheckTimestamp(f.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), f.Timestamp)
		return msg, false
	}

	if len(data) == 0 {
		return "invalid empty data ", false
	}

	if isEqual, serverHash := tools.CompareHash(data, f.Hash); !isEqual {
		return fmt.Sprintf("invalid hash, server hash=%s, client hash=%s", serverHash, f.Hash), false
	}

	return "", true
}

type ShareFolder struct {
	FolderId       string `json:"folder_id" bson:"folder_id"`
	FolderName     []byte `json:"folder_name" bson:"folder_name"`
	FolderOwner    string `json:"folder_owner" bson:"folder_owner"`
	FolderMnemonic []byte `json:"folder_mnemonic" bson:"folder_mnemonic"`
	FolderAuth     string `json:"folder_auth" bson:"folder_auth"`
	Timestamp      int32  `json:"timestamp"  bson:"timestamp"`
}

type ShareFolderListRsp struct {
	List []*ShareFolder `json:"list" bson:"list"`
}

type ShareFolderRecord struct {
	FolderId   string `json:"folder_id" bson:"folder_id"`
	RecordId   string `json:"record_id" bson:"record_id"`
	Id         string `json:"id" bson:"id"`
	RecordData []byte `json:"record_data" bson:"record_data"`
	OwnerAddr  string `json:"owner_addr" bson:"owner_addr"`
}

type ShareFolderRecordRsp struct {
	List []*ShareFolderRecord `json:"list" bson:"list"`
}

type ShareFolderRecordByRid struct {
	FolderId       string `json:"folder_id" bson:"folder_id"`
	RecordId       string `json:"record_id" bson:"record_id"`
	Id             string `json:"id" bson:"id"`
	RecordData     []byte `json:"record_data" bson:"record_data"`
	OwnerAddr      string `json:"owner_addr" bson:"owner_addr"`
	FolderMnemonic []byte `json:"folder_mnemonic" bson:"folder_mnemonic"`
}

type ShareFolderRecordByRidRsp struct {
	List []*ShareFolderRecordByRid `json:"list" bson:"list"`
}

type ShareFolderMemberInfo struct {
	MemberAddr     string `json:"member_addr" bson:"member_addr"`
	MemberName     []byte `json:"member_name" bson:"member_name"`
	MemberSign     string `json:"member_sign" bson:"member_sign"`
	FolderId       string `json:"folder_id" bson:"folder_id"`
	FolderMnemonic []byte `json:"folder_mnemonic" bson:"folder_mnemonic"`
}
type ShareFolderMemberListRsp struct {
	List []*ShareFolderMemberInfo `json:"list" bson:"list"`
}
