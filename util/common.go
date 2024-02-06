/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/

package util

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/web3password/jewel/tools"
	"github.com/web3password/satis/log"
	"google.golang.org/grpc/metadata"
)

const (
	W3PTimeout              = 12
	W3PMaxNonceLength       = 100
	W3PMaxGeneralLenth      = 1024
	W3PMax2048Lenth         = 2048
	W3PMaxRecordLength      = 2500
	W3PMax4096Length        = 4096
	W3PMaxBatchRecordNumber = 500
	W3PMaxBodyLength        = 2 * 1024 * 1024
	W3PMaxAttachmentLength  = 60 * 1024 * 1024
)

// CheckTimestamp check timestamp
func CheckTimestamp(timestamp int64, optionalArgs ...int64) bool {
	allowDiff := int64(W3PTimeout)
	if len(optionalArgs) > 0 && optionalArgs[0] > 0 {
		allowDiff = optionalArgs[0]
	}
	nowTimestamp := time.Now().Unix()
	if timestamp == 0 {
		return false
	}
	if timestamp > nowTimestamp {
		return false
	}
	if (nowTimestamp - timestamp) > allowDiff {
		return false
	}
	return true
}

// CheckSignature check signature
func CheckSignature(addr, signature, params string) bool {
	if err := tools.BizVerifySignature(signature, []byte(params), addr); err != nil {
		log.Logger.Warn("checkSign fail",
			log.String("addr", addr),
			log.String("signature", signature),
			log.String("params", params),
			log.String("errmsg", err.Error()))
		return false
	}
	return true
}

func GetTraceid(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	trace, ok := md["trace_id"]
	if !ok {
		return ""
	}
	return trace[0]
}

type PersonalAuth struct {
	Signature string `json:"signature" bson:"signature"`
	Params    string `json:"params" bson:"params"`
}

type AuthParams struct {
	PrimaryAddress string      `json:"addr" bson:"addr"`
	Timestamp      int64       `json:"timestamp" bson:"timestamp"`
	PublicKeyList  []PublicKey `json:"keys" bson:"keys"`
	Nonce          string      `json:"nonce" bson:"nonce"`
	Type           string      `json:"type" bson:"type"`
}

type PublicKey struct {
	PublicKey string `json:"key" bson:"key"`
	ID        int32  `json:"id" bson:"id"`
}

func ParseAuthParams(auth string) (AuthParams, error) {
	var personalAuth PersonalAuth
	var authParams AuthParams

	err := json.Unmarshal([]byte(auth), &personalAuth)
	if err != nil {
		return authParams, errors.New("unmarshal personal-auth failed")
	}

	err = json.Unmarshal([]byte(personalAuth.Params), &authParams)
	if err != nil {
		return authParams, errors.New("unmarshal personal-auth params failed")
	}

	if !CheckSignature(authParams.PrimaryAddress, personalAuth.Signature, personalAuth.Params) {
		return authParams, errors.New("invalid signautre")
	}

	nowTimestamp := time.Now().Unix()
	if nowTimestamp > authParams.Timestamp {
		return authParams, errors.New("invalid timestamp")
	}

	return authParams, nil
}
