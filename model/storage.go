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

type FileUploadRsp struct {
	Cid  string `bson:"cid" json:"cid"`
	Code int32  `json:"code,omitempty" bson:"code"`
	Msg  string `json:"msg,omitempty" bson:"msg"`
}

type FileReportRsp struct {
	Code int32  `json:"code,omitempty" bson:"code"`
	Msg  string `json:"msg,omitempty" bson:"msg"`
}

type FileDownloadRsp struct {
	Files []FileDownLoadItemRsp `bson:"files" json:"files"`
}

type FileDownLoadItemRsp struct {
	Cid     string `bson:"cid"`
	Content []byte `bson:"content"`
	Code    int32  `json:"code,omitempty" bson:"code"`
	Msg     string `json:"msg,omitempty" bson:"msg"`
}

type FileAttachmentRsp struct {
	Files []FileAttachmentItem `bson:"files" json:"files"`
}

type FileAttachmentItem struct {
	Cid     string `bson:"cid"`
	Content []byte `bson:"content"`
	Success bool   `bson:"success"`
	Message string `bson:"message"`
	Code    int32  `json:"code,omitempty" bson:"code"`
	Msg     string `json:"msg,omitempty" bson:"msg"`
}

type FileUploadReqParams struct {
	Addr      string `json:"addr"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Token     string `json:"token"`
	Sha256    string `json:"sha256"`
	Rid       string `json:"rid"`
	OrgId     string `json:"org_id"`
}

type FileDownloadReqParams struct {
	Addr      string `json:"addr"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Token     string `json:"token"`
	Cid       string `json:"cid"`
	Rid       string `json:"rid"`
	OrgId     string `json:"org_id"`
}

type FileReportReqParams struct {
	Addr      string `json:"addr"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Token     string `json:"token"`
	Rid       string `json:"rid"`
	FlowId    int32  `json:"flow_id"`
	Hash      string `json:"hash"`
	OrgId     string `json:"org_id"`
}

type FileAttachmentReqParams struct {
	Addr      string `json:"addr"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Token     string `json:"token"`
	Hash      string `json:"hash"`
	OrgId     string `json:"org_id"`
}

func (f FileUploadReqParams) Check(data []byte) (string, bool) {
	if !tools.IsValidAddress(f.Addr) {
		return "invalid address " + f.Addr, false
	}

	if f.Token != FileUploadToken {
		return "invalid token " + f.Token, false
	}

	if !util.CheckTimestamp(int64(f.Timestamp), 120) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), f.Timestamp)
		return msg, false
	}

	if len(data) == 0 {
		return "invalid empty data ", false
	}

	if len(f.Nonce) > util.W3PMaxNonceLength || len(f.Rid) > util.W3PMaxNonceLength || len(f.OrgId) > util.W3PMaxNonceLength || len(f.Sha256) > util.W3PMaxNonceLength {
		return "invalid nonce or rid or orgid or hash", false
	}

	if len(data) > util.W3PMaxAttachmentLength {
		return "invalid attachment size", false
	}

	if isEqual, serverHash := tools.CompareHash(data, f.Sha256); !isEqual {
		return fmt.Sprintf("invalid hash, server hash=%s, client hash(sha256)=%s", serverHash, f.Sha256), false
	}

	if len(f.OrgId) == 0 {
		return "invalid org_id", false
	}

	return "", true
}

func (f FileDownloadReqParams) Check() (string, bool) {
	if !tools.IsValidAddress(f.Addr) {
		return "invalid address " + f.Addr, false
	}

	if f.Token != FileDownloadToken {
		return "invalid token " + f.Token, false
	}

	if !util.CheckTimestamp(int64(f.Timestamp)) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), f.Timestamp)
		return msg, false
	}

	if len(f.OrgId) == 0 {
		return "invalid org_id", false
	}

	if len(f.Nonce) > util.W3PMaxNonceLength || len(f.Rid) > util.W3PMaxNonceLength || len(f.OrgId) > util.W3PMaxNonceLength || len(f.Cid) > util.W3PMaxNonceLength {
		return "invalid nonce or rid or orgid or cid", false
	}

	return "", true
}

func (f FileReportReqParams) Check() (string, bool) {
	if !tools.IsValidAddress(f.Addr) {
		return "invalid addr " + f.Addr, false
	}
	if !util.CheckTimestamp(int64(f.Timestamp)) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), f.Timestamp)
		return msg, false
	}
	if len(f.Nonce) == 0 {
		return "invalid nonce " + f.Nonce, false
	}
	if f.Token != FileReportToken {
		return "invalid token " + f.Token, false
	}
	if len(f.Rid) == 0 {
		return "invalid rid " + f.Rid, false
	}
	if f.FlowId == 0 {
		return "invalid flow_id " + string(f.FlowId), false
	}
	if len(f.Hash) == 0 {
		return "invalid hash " + f.Hash, false
	}
	if len(f.OrgId) == 0 {
		return "invalid org_id", false
	}

	if len(f.Nonce) > util.W3PMaxNonceLength || len(f.Rid) > util.W3PMaxNonceLength || len(f.OrgId) > util.W3PMaxNonceLength || len(f.Hash) > util.W3PMaxNonceLength {
		return "invalid nonce or rid or orgid or hash", false
	}

	return "", true
}

func (f FileAttachmentReqParams) Check() (string, bool) {
	if !tools.IsValidAddress(f.Addr) {
		return "invalid address " + f.Addr, false
	}

	if f.Token != FileAttachmentToken {
		return "invalid token " + f.Token, false
	}

	if !util.CheckTimestamp(int64(f.Timestamp), 60) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), f.Timestamp)
		return msg, false
	}
	if len(f.OrgId) == 0 {
		return "invalid org_id", false
	}

	if len(f.Nonce) > util.W3PMaxNonceLength || len(f.OrgId) > util.W3PMaxNonceLength || len(f.Hash) > util.W3PMaxNonceLength {
		return "invalid nonce or orgid or hash", false
	}

	return "", true
}
