/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/
package model

const (
	// StatusOK request success
	StatusOK     = 111111
	StatusFAILED = 222226

	// StatusServiceCheckErr request system error
	StatusServiceCheckErr = 222222
	// StatusParamsErr request params error
	StatusParamsErr = 222223
	// StatusSignatureErr signature error
	StatusSignatureErr = 222224
	// StatusTimestampErr timestamp error
	StatusTimestampErr  = 222225
	StatusAuthErr       = 222226
	StatusDataEmpty     = 222227
	StatusLimitCheckErr = 222228
	StatusLogicCheckErr = 222229
	StatusForbiddenErr  = 222403

	StatusSystemError     = 333333
	StatusSystemErrorCode = 300000

	ARES_PROXY    = "ares"
	INDEX_PROXY   = "index"
	STORAGE_PROXY = "storage"

	MsgOK                = "success"
	MsgParamsErr         = "params fail"
	MsgSystemErr         = "system fail"
	MsgSignatureErr      = "signature fail"
	MsgTimestamp         = "timestamp fail"
	NoAdminShareMnemonic = "no admin share mnemonic"
	MsgLimit             = "You have reached the item limit and cannot add any more."
	MsgRepeat            = "You have already added this member, please do not add again."
	MsgTimeoutErr        = "service timeout error"

	W3PTimeoutMin            = 12
	W3PTimeoutMax            = 15
	W3PTimeoutFileAttachment = 60
)

type Response struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type ResponseBytes struct {
	Code int32  `json:"code" bson:"code"`
	Msg  string `json:"msg" bson:"msg"`
	Data []byte `json:"data" bson:"data"`
}
