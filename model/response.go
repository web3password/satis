/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/

package model

type PersonalSignResponse struct {
	Sign   string `bson:"sign"`
	Params string `bson:"params"`
}

type UserInfoResponse struct {
	Address     string `bson:"addr"`
	ChainId     string `bson:"chain_id"`
	InviteCode  string `bson:"invite_code"`
	StorageType string `bson:"storage_type"`
}

//type AdminMemberListDataBson struct {
//	PrimaryAddress string `bson:"primary_address" json:"primary_address"`
//	EncryptData    string `bson:"encrypt_data" json:"encrypt_data"`
//}

type AdminMnemonicDataBson struct {
	AdminSharedMnemonic string `bson:"admin_shared_mnemonic" json:"admin_shared_mnemonic"`
}

type AdminOperationHistoryBson struct {
	Address   string `bson:"addr" json:"addr"`
	Action    string `bson:"action" json:"action"`
	Content   string `bson:"content" json:"content"`
	Timestamp int32  `bson:"op_timestamp" json:"op_timestamp"`
}

type AdminOperationHistoryBsonV2 struct {
	Address     string `json:"addr" bson:"addr"`
	OpUserInfo  string `json:"op_user_info" bson:"op_user_info"`
	Action      string `json:"action" bson:"action"`
	ObjAddress  string `json:"obj_addr" bson:"obj_addr"`
	ObjUserInfo string `json:"obj_user_info" bson:"obj_user_info"`
	Timestamp   int64  `json:"op_timestamp" bson:"op_timestamp"`
}

type StorageStatBson struct {
	TotalReal          int64  `bson:"total_real" json:"total_real"`
	TotalHumanReadable string `bson:"total_human_readable" json:"total_human_readable"`
	Used               int64  `bson:"used" json:"used"`
	Left               int64  `bson:"left" json:"left"`
}
