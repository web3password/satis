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

const (
	CMDPong                    = "1"
	CMDRegister                = "2"
	CMDGetUserInfo             = "3"
	CMDGetPersonalSignAddress  = "5"
	CMDGetVipInfo              = "6"
	CMDGetVersionDesc          = "7"
	CMDInitialize              = "8"
	CMDAdminRegister           = "9"
	CMDAdminAddMember          = "10"
	CMDAdminBatchImportMember  = "11"
	CMDAdminUpdateMember       = "12"
	CMDAdminRemoveMember       = "13"
	CMDAdminTransferSuperAdmin = "14"
	CMDAdminGetMemberList      = "15"
	CMDAdminGetOrgInfo         = "16"
	CMDAdminUpdateOrgInfo      = "17"
	CMDAdminOperationHistory   = "18"
	CMDAdminGetAdminMnemonic   = "19"

	CMDIndexCheckTx                   = "20"
	CMDIndexAddOrDelCredential        = "21"
	CMDIndexGetPrimaryAddrIndexDetail = "22"
	CMDIndexDelPrimaryAddrIndex       = "23"
	CMDIndexGetAllCredentialTimestamp = "24"
	CMDIndexGetPrimaryAddrIndexList   = "25"
	CMDIndexBatchCheckTx              = "100"
	CMDIndexBatchAddCredential        = "101"
	CMDIndexBatchDeleteCredential     = "102"

	CMDShareFolderCreate       = "26"
	CMDShareFolderUpdate       = "27"
	CMDShareFolderDestroy      = "28"
	CMDShareFolderAddRecord    = "29"
	CMDShareFolderDeleteRecord = "30"
	CMDShareFolderAddMember    = "31"
	CMDShareFolderUpdateMember = "32"
	CMDShareFolderDeleteMember = "33"
	CMDShareFolderMemberExit   = "34"
	CMDShareFolderBatchUpdate  = "41"
	CMDShareFolderFolderList   = "35"
	CMDShareFolderRecordList   = "36"
	CMDShareFolderAttachement  = "37"
	CMDShareFolderMemberList   = "40"

	CMDFileUpload     = "41"
	CMDFileReport     = "42"
	CMDFileDownload   = "43"
	CMDFileAttachment = "44"

	CMDStorageReport = "38"
	CMDStorageStat   = "39"

	CMDAdminAuthorization = "50"
	CMDGetVersionConfig   = "51"

	CMDShareFolderRecordListByRid = "54"

	CMDVipSubscriptionList    = "100"
	CMDVipCreateOrder         = "101"
	CMDVipCheckOrder          = "102"
	CMDVipPaymentList         = "103"
	CMDVipAppleVerifyReceipt  = "104"
	CMDGetDiscountCode        = "105"
	CMDVipGetConfig           = "106"
	CMDGetOrderList           = "107"
	CMDGracefulRestartSignal  = "404"
	CMDGetVipIOSPromotionSign = "109"
	CMDGetVipPrice            = "110"

	RegisterToken               = "userRegister"
	GetVIPInfoToken             = "getVipInfo"
	GetUserInfoToken            = "userInfo"
	GetPersonalSignAddressToken = "getPersonalSignAddress"

	GetCredentialToken              = "getCredential"
	AddCredentialToken              = "addCredential"
	DelCredentialToken              = "delCredential"
	DeleteAllCredentialToken        = "deleteAllCredential"
	GetAllCredentialTimestampToken  = "getAllCredentialTimestamp"
	GetCredentialListToken          = "getCredentialList"
	IndexBatchCheckTxToken          = "batchCheckTx"
	IndexBatchAddCredentialToken    = "batchAddCredential"
	IndexBatchDeleteCredentialToken = "batchDeleteCredential"

	InitializeToken              = "userInit"
	VersionDescToken             = "versionDesc"
	AdminRegisterToken           = "adminRegister"
	AdminAddMemberToken          = "adminAddMember"
	AdminBatchImportMemberToken  = "adminBatchImportMember"
	AdminUpdateMemberToken       = "adminUpdateMember"
	AdminRemoveMemberToken       = "adminRemoveMember"
	AdminTransferSuperAdminToken = "adminTransferSuperAdmin"
	AdminGetMemberListToken      = "adminGetMemberList"
	AdminGetOrgInfoToken         = "adminGetOrgInfo"
	AdminUpdateOrgInfoToken      = "adminUpdateOrgInfo"
	AdminOperationHistoryToken   = "adminOperationHistory"
	AadminAuthorizationToken     = "adminAuthorization"
	AdminGetAdminMnemonicToken   = "getAdminMnemonic"
	StorageReportToken           = "storageReport"
	StorageStatToken             = "storageStat"

	ShareFolderCreateToken          = "shareFolderCreate"
	ShareFolderUpdateToken          = "shareFolderUpdate"
	ShareFolderDestroyToken         = "shareFolderDestroy"
	ShareFolderAddRecordToken       = "shareFolderAddRecord"
	ShareFolderDeleteRecordToken    = "shareFolderDeleteRecord"
	ShareFolderAddMemberToken       = "shareFolderAddMember"
	ShareFolderUpdateMemberToken    = "shareFolderUpdateMember"
	ShareFolderDeleteMemberToken    = "shareFolderDeleteMember"
	ShareFolderMemberExitToken      = "shareFolderMemberExit"
	ShareFolderBatchUpdateToken     = "shareFolderBatchUpdate"
	ShareFolderFolderListToken      = "shareFolderFolderList"
	ShareFolderRecordListToken      = "shareFolderRecordList"
	ShareFolderRecordListTokenByRid = "shareFolderRecordListByRid"
	ShareFolderAttachementToken     = "shareFolderAttachement"
	ShareFolderMemberListToken      = "shareFolderMemberList"
	FileUploadToken                 = "fileUpload"
	FileDownloadToken               = "fileDownload"
	FileAttachmentToken             = "fileAttachment"
	FileReportToken                 = "fileReport"
	GetVersionConfigToken           = "getVersionConfig"

	VipGetConfigToken          = "vip-getConfig"
	VipSubscriptionListToken   = "vip-subscriptionList"
	VipCreateOrderToken        = "vip-createOrder"
	VipCheckOrderToken         = "vip-checkOrder"
	VipAppleVerifyReceiptToken = "vip-apple-verifyReceipt"
	VipPaymentListToken        = "vip-paymentList"
	VipDiscountToken           = "vip-discountcode"
	VipGetOrderList            = "vip-getOrderList"
	VipIOSPromotionSign        = "vip-apple-getVipIOSPromotionSign"

	VipPrice = "vip-price"
)

type AdminAuthorizationRsp struct {
	Authorization string `json:"authorization"`
	Code          int32  `json:"code,omitempty"`
	Msg           string `json:"msg,omitempty"`
}

type UserInfoRsp struct {
	Code int32           `json:"code,omitempty"`
	Msg  string          `json:"msg,omitempty"`
	Data *UserInfoResult `json:"data,omitempty"`
}

type UserInfoResult struct {
	Address     string `json:"addr"`
	ChainId     string `json:"chain_id"`
	InviteCode  string `json:"invite_code"`
	StorageType string `json:"storage_type"`
}

type VersionDescRsp struct {
	VersionDesc string `json:"version_desc"`
	Code        int32  `json:"code,omitempty"`
	Msg         string `json:"msg,omitempty"`
}

// VIPInfo .
type VIPInfo struct {
	ExpireTime  int64  `json:"expireTime"`
	Version     string `json:"version"`
	VersionDesc string `json:"version_desc"`
}

// VIPInfoRsp .
type VIPInfoRsp struct {
	Code int32          `json:"code,omitempty" bson:"code"`
	Msg  string         `json:"msg,omitempty" bson:"msg"`
	Data *VIPInfoResult `json:"data,omitempty" bson:"data"`
}

type VIPInfoResult struct {
	Signature string        `json:"signature" bson:"signature"`
	Params    string        `json:"params" bson:"params"`
	OrgInfo   []UserOrgInfo `json:"org_info" bson:"org_info"`
}

type UserOrgInfo struct {
	OrgId          string `json:"org_id" bson:"org_id"`
	OrgName        []byte `json:"org_name" bson:"org_name"`
	Logo           []byte `json:"logo" bson:"logo"`
	SelfHostUrl    []byte `json:"self_host_url" bson:"self_host_url"`
	SharedMnemonic []byte `json:"member_share_mnemonic" bson:"member_share_mnemonic"`
}

type PersonalSignListRsp struct {
	Code int32           `json:"code,omitempty" bson:"code"`
	Msg  string          `json:"msg,omitempty" bson:"msg"`
	List []*PersonalSign `json:"list" bson:"list"`
}

// PersonalSign .
type PersonalSign struct {
	Sign   string `json:"sign"`
	Params string `json:"params"`
}

// AdminMember .
type AdminMember struct {
	PrimaryAddress string `json:"primary_address" bson:"primary_address"`
	EncryptData    string `json:"encrypt_data" bson:"encrypt_data"`
}

// AdminShareMnemonic .
type AdminShareMnemonic struct {
	AdminShareMnemonic  string `json:"admin_share_mnemonic" bson:"admin_share_mnemonic"`
	MemberShareMnemonic string `json:"member_share_mnemonic" bson:"member_share_mnemonic"`
}

//type AdminMemberListRsp struct {
//	Code int    `json:"code"`
//	Msg  string `json:"msg"`
//	Data struct {
//		Members []struct {
//			PrimaryAddress string `json:"primary_address"`
//			EncryptData    string `json:"encrypt_data"`
//		} `json:"members"`
//	} `json:"data"`
//}

type PersonalSignParams struct {
	// primary address
	Address string `json:"addr"`
	// nonce
	Nonce string `json:"nonce"`
	//server_timestamp
	ServerTimestamp int64 `json:"server_timestamp"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// token
	PersonalAddress string `json:"personal_address"`
}

type RegisterParams struct {
	// primary address
	Address string `json:"addr"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// token
	Token string `json:"token"`
}

type InitializeParams struct {
	// primary address
	Address string `json:"addr"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// token
	Token string `json:"token"`
}

type GetPersonalSignAddressParams struct {
	// primary address
	Address string `json:"addr"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// token
	Token string `json:"token"`
	// official_addrs
	OfficialAddrs []string `json:"official_addrs"`
}

type GetVIPInfoParams struct {
	// primary address
	Address string `json:"addr"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// token
	Token string `json:"token"`
}

type GetUserInfoParams struct {
	// primary address
	Address string `json:"addr"`
	// primary address
	TagAddr string `json:"org_id"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// token
	Token string `json:"token"`
}

type CheckTxParams struct {
	// primary address
	Address string `json:"addr"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// tx_hash
	TxHash string `json:"hash"`
	OrgId  string `json:"org_id"`
}

type BatchCheckTxParams struct {
	Addr      string `json:"addr"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Token     string `json:"token"`
	Hash      string `json:"hash"`
	OrgId     string `json:"org_id"`
}

type AddCredentialParams struct {
	// primary address
	Address string `json:"addr"`
	// op_timestamp
	OpTimestamp int64 `json:"op_timestamp"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// id
	ID string `json:"id"`
	// credential
	Hash  string `json:"hash"`
	OrgId string `json:"org_id"`
}

type BatchAddCredentialParams struct {
	Addr      string `json:"addr"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Token     string `json:"token"`
	Hash      string `json:"hash"`
	OrgId     string `json:"org_id"`
}

type DeleteCredentialParams struct {
	// primary address
	Address string `json:"addr"`
	// op_timestamp
	OpTimestamp int64 `json:"op_timestamp"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// id
	ID string `json:"id"`
	// credential
	Hash  string `json:"hash"`
	OrgId string `json:"org_id"`
}

type BatchDeleteCredentialParams struct {
	Addr      string `json:"addr"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Token     string `json:"token"`
	Hash      string `json:"hash"`
	OrgId     string `json:"org_id"`
}

type GetCredentialParams struct {
	// primary address
	Address string `json:"addr"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// id
	ID string `json:"id"`
	// token
	Token string `json:"token"`
	OrgId string `json:"org_id"`
}

type DeleteAllCredentialParams struct {
	// primary address
	Address string `json:"addr"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// token
	Token string `json:"token"`
	OrgId string `json:"org_id"`
}

type GetAllCredentialTimestampParams struct {
	// primary address
	Address string `json:"addr"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// token
	Token string `json:"token"`
	OrgId string `json:"org_id"`
}

type GetCredentialListParams struct {
	// primary address
	Address string `json:"addr"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// ids
	IDs []string `json:"ids"`
	// token
	Token string `json:"token"`
	OrgId string `json:"org_id"`
}

type GetVersionDescParams struct {
	// primary address
	Address string `json:"addr"`
	// version
	Version string `json:"version"`
	// language
	Language string `json:"language"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// token
	Token string `json:"token"`
}

type AdminRegisterParams struct {
	// primary address
	Address string `json:"addr"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// token
	Token string `json:"token"`
	//auth
	Auth string `json:"personal_auth"`
}

type AdminRemoveMemberParams struct {
	Address       string `json:"addr"`
	Timestamp     int64  `json:"timestamp"`
	Nonce         string `json:"nonce"`
	Token         string `json:"token"`
	TagAddress    string `json:"tag_address"`
	MemberAddress string `json:"member_address"`
}

type AdminAddOrUpdateMemberParams struct {
	// primary address
	Address string `json:"addr"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// token
	Token string `json:"token"`
	// tag_address
	TagAddress string `json:"tag_address"`
	// member_address
	MemberAddress string `json:"member_address"`
	// member_data
	MemberData string `json:"member_data"`
	// member_share_mnemonic
	MemberShareMnemonic string `json:"member_share_mnemonic"`
	// admin_share_mnemonic
	AdminShareMnemonic string `json:"admin_share_mnemonic"`
}

type AdminGetMemberListParams struct {
	// primary address
	Address string `json:"addr"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// token
	Token string `json:"token"`
	// tag_address
	TagAddress string `json:"tag_address"`
}

type AdminAuthorizationParams struct {
	// primary address
	Address string `json:"addr"`
	//tag_address
	TagAddress string `json:"tag_address"`
	// nonce
	Nonce string `json:"nonce"`
	// token
	Token string `json:"token"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
}

type TransferSuperAdminParams struct {
	// primary address
	Address string `json:"addr"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// token
	Token string `json:"token"`
	// credential
	Hash string `json:"hash"`
	//// sign
	//Sign string `json:"sign"`
	//// member_addr
	//TagAddress string `json:"tag_address"`
	//// new_super_admin_addr
	//NewSuperAdminAddr string `json:"super_admin"`
	//// new_super_admin_data
	//NewSuperAdminData string `json:"super_admin_data"`

}

type OperationHistoryParams struct {
	// primary address
	Address string `json:"addr"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// token
	Token string `json:"token"`
	// view_addr
	ViewAddress string `json:"view_addr"`
	// credential
	Hash string `json:"hash"`
}

type StorageReportParams struct {
	InitializeParams
	// action
	Action string `json:"action"`
	// amount
	Amount int64 `json:"amount"`
}

type StorageStatParams struct {
	InitializeParams
}

type AdminGetOrgInfoParams struct {
	// primary address
	Address string `json:"addr"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// token
	Token string `json:"token"`
}

type GetAdminMnemonicParams struct {
	// primary address
	Address string `json:"addr"`
	// tag_address
	TagAddress string `json:"tag_address"`
	// timestamp
	Timestamp int64 `json:"timestamp"`
	// nonce
	Nonce string `json:"nonce"`
	// token
	Token string `json:"token"`
}

type AdminBatchImportMemberParams struct {
	Addr       string `json:"addr"`
	Timestamp  int64  `json:"timestamp"`
	Nonce      string `json:"nonce"`
	Token      string `json:"token"`
	MemberList []struct {
		TagAddress          string `json:"tag_address"`
		MemberAddress       string `json:"member_address"`
		MemberData          string `json:"member_data"`
		MemberShareMnemonic string `json:"member_share_mnemonic"`
		AdminShareMnemonic  string `json:"admin_share_mnemonic"`
	} `json:"member_list"`
}

type GetVersionConfigParams struct {
	Address   string `json:"addr"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Token     string `json:"token"`
	Hash      string `json:"hash"`
}

func (s GetVersionConfigParams) Check(token string, timestamp int64, data []byte) (string, bool) {
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

type GetVersionData struct {
	Version string `json:"version" bson:"version"`
}

type GetVersionConfigRsp struct {
	VersionName                   string `json:"version_name" bson:"version_name"`
	VersionValue                  string `json:"version_value" bson:"version_value"`
	ShowRebate                    int32  `json:"show_rebate" bson:"show_rebate"`
	RebateRate                    string `json:"rebate_rate" bson:"rebate_rate"`
	RecordCountLimit              int32  `json:"record_count_limit" bson:"record_count_limit"`
	OneRecordWebsiteLimit         int32  `json:"one_record_website_limit"  bson:"one_record_website_limit"`
	OneRecordLinkwebsiteLimit     int32  `json:"one_record_linkwebsite_limit"  bson:"one_record_linkwebsite_limit"`
	OneRecordAttachmentCountLimit int32  `json:"one_record_attachment_count_limit" bson:"one_record_attachment_count_limit"`
	OneUserAttachmemtSpaceLimit   int64  `json:"one_user_attachment_space_limit" bson:"one_user_attachment_space_limit"`
	OneAttachmentSizeLimit        int64  `json:"one_attachment_size_limit" bson:"one_attachment_size_limit"`
	SharefolderCreateLimit        int32  `json:"sharefolder_create_limit" bson:"sharefolder_create_limit"`
	SharefolderRecordLimit        int32  `json:"sharefolder_record_limit" bson:"sharefolder_record_limit"`
	SharefolderMemberLimit        int32  `json:"sharefolder_member_limit" bson:"sharefolder_member_limit"`
	VersionUser                   int32  `json:"version_user" bson:"version_user"`
	AdminMemberLimit              int32  `json:"admin_member_limit" bson:"admin_member_limit"`
	OneRecordSizeLimit            int64  `json:"one_record_size_limit" bson:"one_record_size_limit"`
}

func (r AdminGetOrgInfoParams) Check() bool {
	if !tools.IsValidAddress(r.Address) {
		return false
	}
	if r.Timestamp == 0 {
		return false
	}
	return true
}

func (r RegisterParams) Check() (string, bool) {
	if !tools.IsValidAddress(r.Address) {
		return "invalid address " + r.Address, false
	}
	if !util.CheckTimestamp(r.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), r.Timestamp)
		return msg, false
	}
	if r.Token != RegisterToken {
		return "invalid token " + r.Token, false
	}
	return "", true
}

func (r InitializeParams) Check() bool {
	if !tools.IsValidAddress(r.Address) {
		return false
	}
	if r.Timestamp == 0 {
		return false
	}
	if r.Token != InitializeToken {
		return false
	}
	return true
}

func (r GetPersonalSignAddressParams) Check() (string, bool) {
	if !tools.IsValidAddress(r.Address) {
		return "invalid address " + r.Address, false
	}
	if !util.CheckTimestamp(r.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), r.Timestamp)
		return msg, false
	}
	if r.Token != GetPersonalSignAddressToken {
		return "invalid token " + r.Token, false
	}
	if len(r.OfficialAddrs) == 0 {
		return "empty official_addrs", false
	}
	return "", true
}

func (r GetVIPInfoParams) Check() (string, bool) {
	if !tools.IsValidAddress(r.Address) {
		return "invalid address " + r.Address, false
	}
	if !util.CheckTimestamp(r.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), r.Timestamp)
		return msg, false
	}
	if r.Token != GetVIPInfoToken {
		return "invalid token " + r.Token, false
	}
	return "", true
}

func (r GetUserInfoParams) Check() (string, bool) {
	if !tools.IsValidAddress(r.Address) {
		return "invalid address " + r.Address, false
	}
	if !util.CheckTimestamp(r.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), r.Timestamp)
		return msg, false
	}
	if r.Token != GetUserInfoToken {
		return "invalid token " + r.Token, false
	}
	return "", true
}

func (c CheckTxParams) Check() (string, bool) {
	if !tools.IsValidAddress(c.Address) {
		return "invalid address " + c.Address, false
	}

	if !util.CheckTimestamp(c.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), c.Timestamp)
		return msg, false
	}
	if c.TxHash == "" {
		return "empty hash", false
	}
	return "", true
}

func (c BatchCheckTxParams) Check(data []byte) (string, bool) {
	if !tools.IsValidAddress(c.Addr) {
		return "invalid address " + c.Addr, false
	}
	if !util.CheckTimestamp(c.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), c.Timestamp)
		return msg, false
	}

	if isEqual, serverHash := tools.CompareHash(data, c.Hash); !isEqual {
		return fmt.Sprintf("invalid hash, server hash=%s, client hash=%s", serverHash, c.Hash), false
	}

	if c.Token != IndexBatchCheckTxToken {
		return "invalid token " + c.Token, false
	}
	return "", true
}

func (a AddCredentialParams) Check(data []byte) (string, bool) {
	if !tools.IsValidAddress(a.Address) {
		return "invalid address " + a.Address, false
	}
	if a.OpTimestamp == 0 {
		return "invalid op_timestamp " + a.Address, false
	}
	if !util.CheckTimestamp(a.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), a.Timestamp)
		return msg, false
	}
	if isEqual, serverHash := tools.CompareHash(data, a.Hash); !isEqual {
		return fmt.Sprintf("invalid hash, server hash=%s, client hash=%s", serverHash, a.Hash), false
	}
	if a.ID == "" {
		return "invalid id", false
	}
	return "", true
}

func (c BatchAddCredentialParams) Check(data []byte) (string, bool) {
	if !tools.IsValidAddress(c.Addr) {
		return "invalid address " + c.Addr, false
	}
	if !util.CheckTimestamp(c.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), c.Timestamp)
		return msg, false
	}
	if isEqual, serverHash := tools.CompareHash(data, c.Hash); !isEqual {
		return fmt.Sprintf("invalid hash, server hash=%s, client hash=%s", serverHash, c.Hash), false
	}
	if c.Token != IndexBatchAddCredentialToken {
		return "invalid token " + c.Hash, false
	}
	return "", true
}

func (d DeleteCredentialParams) Check(data []byte) (string, bool) {
	if !tools.IsValidAddress(d.Address) {
		return "invalid address " + d.Address, false
	}
	if !util.CheckTimestamp(d.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), d.Timestamp)
		return msg, false
	}
	if isEqual, serverHash := tools.CompareHash(data, d.Hash); !isEqual {
		return fmt.Sprintf("invalid hash, server hash=%s, client hash=%s", serverHash, d.Hash), false
	}
	if d.ID == "" {
		return "invalid id ", false
	}
	return "", true
}

func (c BatchDeleteCredentialParams) Check(data []byte) (string, bool) {
	if !tools.IsValidAddress(c.Addr) {
		return "invalid address " + c.Addr, false
	}
	if !util.CheckTimestamp(c.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), c.Timestamp)
		return msg, false
	}
	if isEqual, serverHash := tools.CompareHash(data, c.Hash); !isEqual {
		return fmt.Sprintf("invalid hash, server hash=%s, client hash=%s", serverHash, c.Hash), false
	}
	if c.Token != IndexBatchDeleteCredentialToken {
		return "invalid token " + c.Token, false
	}
	return "", true
}

func (g GetCredentialParams) Check() (string, bool) {
	if !tools.IsValidAddress(g.Address) {
		return "invalid address " + g.Address, false
	}
	if !util.CheckTimestamp(g.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), g.Timestamp)
		return msg, false
	}
	if g.ID == "" {
		return "invalid id ", false
	}
	if g.Token != GetCredentialToken {
		return "invalid token " + g.Token, false
	}
	return "", true
}

func (d DeleteAllCredentialParams) Check() (string, bool) {
	if !tools.IsValidAddress(d.Address) {
		return "invalid address " + d.Address, false
	}
	if !util.CheckTimestamp(d.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), d.Timestamp)
		return msg, false
	}
	if d.Token != DeleteAllCredentialToken {
		return "invalid token " + d.Token, false
	}
	return "", true
}

func (g GetAllCredentialTimestampParams) Check() (string, bool) {
	if !tools.IsValidAddress(g.Address) {
		return "invalid address " + g.Address, false
	}
	if !util.CheckTimestamp(g.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), g.Timestamp)
		return msg, false
	}
	if g.Token != GetAllCredentialTimestampToken {
		return "invalid token " + g.Token, false
	}
	return "", true
}

func (g GetCredentialListParams) Check() (string, bool) {
	if !tools.IsValidAddress(g.Address) {
		return "invalid address " + g.Address, false
	}
	if !util.CheckTimestamp(g.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), g.Timestamp)
		return msg, false
	}
	if len(g.IDs) == 0 {
		return "invalid ids", false
	}
	if len(g.IDs) > 500 {
		return "max length ids is 500", false
	}
	if g.Token != GetCredentialListToken {
		return "invalid token " + g.Token, false
	}
	return "", true
}

func (g GetVersionDescParams) Check() bool {
	if !tools.IsValidAddress(g.Address) {
		return false
	}
	if g.Timestamp == 0 {
		return false
	}
	if g.Version == "" {
		return false
	}
	if g.Token != VersionDescToken {
		return false
	}
	return true
}

func (a AdminAddOrUpdateMemberParams) Check() bool {
	if !tools.IsValidAddress(a.Address) {
		return false
	}
	if a.Timestamp == 0 {
		return false
	}
	if a.MemberAddress == "" {
		return false
	}
	if a.TagAddress == "" {
		return false
	}
	if a.MemberData == "" {
		return false
	}
	//if a.Token != AdminAddMemberToken {
	//	return false
	//}
	return true
}

func (r AdminRegisterParams) Check() (string, bool) {
	if !tools.IsValidAddress(r.Address) {
		return "invalid address " + r.Address, false
	}
	if r.Token != AdminRegisterToken {
		return "invalid token " + r.Token, false
	}
	if !util.CheckTimestamp(r.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), r.Timestamp)
		return msg, false
	}
	if r.Auth == "" {
		return "invalid auth " + r.Auth, false
	}
	return "", true
}

func (t TransferSuperAdminParams) Check() bool {
	if !tools.IsValidAddress(t.Address) {
		return false
	}
	if !util.CheckTimestamp(t.Timestamp) {
		return false
	}
	if t.Token != AdminTransferSuperAdminToken {
		return false
	}
	return true
}

func (t OperationHistoryParams) Check() bool {
	if !tools.IsValidAddress(t.Address) {
		return false
	}
	if t.Timestamp == 0 {
		return false
	}
	if t.Token != AdminOperationHistoryToken {
		return false
	}
	return true
}

func (t AdminAuthorizationParams) Check() bool {
	if !tools.IsValidAddress(t.Address) || !tools.IsValidAddress(t.TagAddress) {
		return false
	}

	if t.Timestamp == 0 {
		return false
	}
	if t.Token != AadminAuthorizationToken {
		return false
	}
	return true
}

func (a AdminGetMemberListParams) Check() bool {
	if !tools.IsValidAddress(a.Address) {
		return false
	}
	if a.Timestamp == 0 {
		return false
	}
	if a.TagAddress == "" {
		return false
	}
	if a.Token != AdminGetMemberListToken {
		return false
	}
	return true
}

func (g GetAdminMnemonicParams) Check() bool {
	if !tools.IsValidAddress(g.Address) {
		return false
	}
	if g.Timestamp == 0 {
		return false
	}
	if g.TagAddress == "" {
		return false
	}
	if g.Token != AdminGetAdminMnemonicToken {
		return false
	}
	return true
}

func (a AdminBatchImportMemberParams) Check() bool {
	if !tools.IsValidAddress(a.Addr) {
		return false
	}
	if a.Timestamp == 0 {
		return false
	}
	if a.Token != AdminBatchImportMemberToken {
		return false
	}
	if len(a.MemberList) < 0 {
		return false
	}
	return true
}

func (a AdminRemoveMemberParams) Check() bool {
	if !tools.IsValidAddress(a.Address) {
		return false
	}
	if a.Timestamp == 0 {
		return false
	}
	if a.Token != AdminRemoveMemberToken {
		return false
	}
	if a.TagAddress == "" {
		return false
	}
	if a.MemberAddress == "" {
		return false
	}
	return true
}

func (s StorageReportParams) Check() (string, bool) {
	if !tools.IsValidAddress(s.Address) {
		return "invalid address " + s.Address, false
	}
	if s.Action != "incr" && s.Action != "decr" {
		return "invalid action " + s.Action, false
	}

	if s.Amount == 0 {
		return "invalid amount", false
	}
	if s.Token != StorageReportToken {
		return "invalid token " + s.Token, false
	}

	if !util.CheckTimestamp(s.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), s.Timestamp)
		return msg, false
	}
	return "", true
}

func (s StorageStatParams) Check() (string, bool) {
	if !tools.IsValidAddress(s.Address) {
		return "invalid address " + s.Address, false
	}

	if s.Token != StorageStatToken {
		return "invalid token " + s.Token, false
	}

	if !util.CheckTimestamp(s.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), s.Timestamp)
		return msg, false
	}
	return "", true
}
