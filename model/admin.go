/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/

package model

import (
	"fmt"
	"github.com/web3password/jewel/tools"
	"github.com/web3password/satis/util"
	"time"
)

type AdminRsp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type AdminCommonParams struct {
	Address   string `json:"addr"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Token     string `json:"token"`
	Hash      string `json:"hash"`
}

func (s AdminCommonParams) Check(token string, timestamp int64, data []byte) (string, bool) {
	nowTimestamp := time.Now().Unix()

	if !tools.IsValidAddress(s.Address) {
		return "invalid address " + s.Address, false
	}

	if len(s.Nonce) > util.W3PMaxNonceLength || len(s.Hash) > util.W3PMaxNonceLength {
		return "invalid nonce or hash", false
	}

	if s.Timestamp == 0 {
		msg := fmt.Sprintf("invalid timestamp , req timestamp is zero : %d", s.Timestamp)
		return msg, false
	}
	if timestamp > nowTimestamp {
		msg := fmt.Sprintf("invalid timestamp(client time>system time) , server timestamp=%d, client timestamp=%d", nowTimestamp, s.Timestamp)
		return msg, false
	}
	if nowTimestamp-timestamp > util.W3PTimeout {
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

type AdminMemberInfo struct {
	//MemberId           string `json:"member_id" bson:"member_id"`
	MemberAddr string `json:"member_addr" bson:"member_addr"`
	//MemberName         []byte `json:"member_name" bson:"member_name"`
	MemberData         []byte `json:"member_data" bson:"member_data"`
	AdminShareMnemonic []byte `json:"admin_share_mnemonic" bson:"admin_share_mnemonic"`
	Role               string `json:"role" bson:"role"`
	MemberSign         string `json:"member_sign" bson:"member_sign"`
}

type AdminMemberListRsp struct {
	List []*AdminMemberInfo `json:"list" bson:"list"`
}

type AdminShareMnemonicRsp struct {
	AdminShareMnemonic  []byte `json:"admin_share_mnemonic" bson:"admin_share_mnemonic"`
	MemberShareMnemonic []byte `json:"member_share_mnemonic" bson:"member_share_mnemonic"`
}

type MemberShareMnemonicRsp struct {
	MemberShareMnemonic string `json:"member_share_mnemonic" bson:"member_share_mnemonic"`
}

type OrgInfoData struct {
	OrgId       string `json:"org_id" bson:"org_id"`
	OrgName     []byte `json:"org_name" bson:"org_name"`
	Logo        []byte `json:"logo" bson:"logo"`
	SelfHostUrl []byte `json:"self_host_url" bson:"self_host_url"`
	Version     string `json:"version" bson:"version"`
	VipType     string `json:"vip_type" bson:"vip_type"`
	UsedSeats   string `json:"used_seats" bson:"used_seats"`
	Seats       string `json:"seats" bson:"seats"`
	ExpiredTime string `json:"expired_time" bson:"expired_time"`
}

type VipRegister struct {
	OrgId string `bson:"org_id"`
}
