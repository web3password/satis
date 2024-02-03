/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/

package model

type EmptyRsp struct {
}

type CheckTxRsp struct {
	Height int64  `json:"height"`
	Code   int32  `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
}

type BatchCheckTxRsp struct {
	Data struct {
		List []*BatchCheckTxRspDataListItem `json:"list" bson:"list"`
	} `json:"data"`
}

type BatchCheckTxRspDataListItem struct {
	Hash    string `json:"hash" bson:"hash"`
	Success int32  `json:"success" bson:"success"`
}

type AddOrDelCredentialRsp struct {
	TxHash string `json:"tx_hash"`
	Code   int32  `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
}

type BatchAddCredentialRsp struct {
	Data struct {
		List []*BatchAddCredentialRspDataListItem `json:"list" bson:"list"`
	} `json:"data"`
}

type BatchAddCredentialRspDataListItem struct {
	Id   string `json:"id" bson:"id"`
	Hash string `json:"hash" bson:"hash"`
}

type BatchDeleteCredentialRsp struct {
	Data struct {
		List []*BatchDeleteCredentialRspDataListItem `json:"list" bson:"list"`
	} `json:"data"`
}

type BatchDeleteCredentialRspDataListItem struct {
	Id   string `json:"id" bson:"id"`
	Hash string `json:"hash" bson:"hash"`
}

type GetCredentialRspData struct {
	List []*GetCredentialRsp `json:"list" bson:"list"`
}

type GetCredentialListRsp struct {
	Code int32               `json:"code,omitempty" bson:"code"`
	Msg  string              `json:"msg,omitempty" bson:"msg"`
	List []*GetCredentialRsp `json:"list" bson:"list"`
}
type GetCredentialRsp struct {
	Id          string `json:"id" bson:"id"`
	Credential  []byte `json:"credential" bson:"credential"`
	OpTimestamp int32  `json:"op_timestamp" bson:"op_timestamp"`
	Code        int32  `json:"code,omitempty" bson:"code"`
	Msg         string `json:"msg,omitempty" bson:"msg"`
}

type GetAllCredentialTimestampRspData struct {
	List []*GetAllCredentialTimestampRsp `json:"list" bson:"list"`
}

type GetAllCredentialTimestampListRsp struct {
	Code int32                           `json:"code,omitempty" bson:"code"`
	Msg  string                          `json:"msg,omitempty" bson:"msg"`
	List []*GetAllCredentialTimestampRsp `json:"list" bson:"list"`
}

type GetAllCredentialTimestampRsp struct {
	Id          string `json:"id" bson:"id"`
	OpTimestamp int32  `json:"op_timestamp" bson:"op_timestamp"`
}

type AddCredentialRsp struct {
	Hash string `bson:"hash"`
}

type NewCheckTxRsp struct {
	IsSuccess int32 `bson:"success"`
}
