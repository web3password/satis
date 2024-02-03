/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/

package dao

import (
	"context"
	"math/big"
	"math/rand"
	"net"
	"sync"

	"github.com/bwmarrin/snowflake"
	"github.com/web3password/satis/config"
	"github.com/web3password/satis/log"
	"github.com/web3password/satis/model"
	pb "github.com/web3password/w3p-protobuf/user"
)

type DAO interface {
	GenerateID() int64

	Stream(server pb.User_StreamServer) error

	RegisterUser(ctx context.Context, signature, params string) (*pb.RegisterRsp, error)
	Initialize(ctx context.Context, signature, params string) error
	GetVersionDesc(ctx context.Context, signature, params string) (model.VersionDescRsp, error)
	//GetPersonalSignAddress(ctx context.Context, signature, params string) ([]*model.PersonalSign, error)
	GetPersonalSignAddress(ctx context.Context, signature, params string) (model.PersonalSignListRsp, error)
	GetVIPInfo(ctx context.Context, signature, params string) (model.VIPInfoRsp, error)
	GetUserInfo(ctx context.Context, signature, params string) (model.UserInfoRsp, error)
	StorageReport(ctx context.Context, signature, params string) (*pb.StorageReportRsp, error)
	StorageStat(ctx context.Context, signature, params string) (*pb.StorageStatRsp, error)
	GetVersionConfig(ctx context.Context, req *pb.GetVersionConfigReq) (*pb.GetVersionConfigRsp, error)

	CheckTx(ctx context.Context, req *pb.CheckTxReq) (model.CheckTxRsp, error)
	BatchCheckTx(ctx context.Context, req *pb.BatchCheckTxReq) (*pb.BatchCheckTxRsp, error)
	AddOrDelCredential(ctx context.Context, signature, params string, data []byte) (model.AddOrDelCredentialRsp, error)
	BatchAddCredential(ctx context.Context, req *pb.BatchAddCredentialReq) (*pb.BatchAddCredentialRsp, error)
	BatchDeleteCredential(ctx context.Context, req *pb.BatchDeleteCredentialReq) (*pb.BatchDeleteCredentialRsp, error)
	GetPrimaryAddrIndexDetail(ctx context.Context, signature, params string) (model.GetCredentialRsp, error)
	DeleteAllCredential(ctx context.Context, signature, params string) (model.AddOrDelCredentialRsp, error)
	//GetAllCredentialTimestamp(ctx context.Context, req *pb.GetAllCredentialTimestampReq) ([]*model.GetAllCredentialTimestampRsp, error)
	GetAllCredentialTimestamp(ctx context.Context, req *pb.GetAllCredentialTimestampReq) (model.GetAllCredentialTimestampListRsp, error)
	//GetPrimaryAddrIndexList(ctx context.Context, req *pb.GetCredentialListReq) ([]*model.GetCredentialRsp, error)
	GetPrimaryAddrIndexList(ctx context.Context, req *pb.GetCredentialListReq) (model.GetCredentialListRsp, error)

	AdminRegister(ctx context.Context, req *pb.AdminRegisterReq) (*pb.AdminRegisterRsp, error)
	AdminAddMember(ctx context.Context, req *pb.AdminAddMemberReq) (model.AdminRsp, error)
	AdminTransferSuperAdmin(ctx context.Context, req *pb.AdminTransferSuperAdminReq) (*pb.AdminTransferSuperAdminRsp, error)
	AdminOperationHistory(ctx context.Context, req *pb.AdminOperationHistoryReq) (*pb.AdminOperationHistoryRsp, error)
	AdminGetOrgInfo(ctx context.Context, req *pb.AdminGetOrgInfoReq) (*pb.AdminGetOrgInfoRsp, error)
	AdminBatchImportMember(ctx context.Context, rep *pb.AdminBatchImportMemberReq) (*pb.AdminBatchImportMemberRsp, error)
	AdminUpdateMember(ctx context.Context, req *pb.AdminUpdateMemberReq) (model.AdminRsp, error)
	AdminRemoveMember(ctx context.Context, req *pb.AdminRemoveMemberReq) (model.AdminRsp, error)
	AdminGetMemberList(ctx context.Context, req *pb.AdminGetMemberListReq) (*pb.AdminGetMemberListRsp, error)
	AdminAuthorization(ctx context.Context, signature, params string) (model.AdminAuthorizationRsp, error)
	GetAdminMnemonic(ctx context.Context, req *pb.GetAdminMnemonicReq) (*pb.GetAdminMnemonicRsp, error)
	AdminUpdateOrgInfo(ctx context.Context, req *pb.AdminUpdateOrgInfoReq) (*pb.AdminUpdateOrgInfoRsp, error)

	ShareFolderCreate(ctx context.Context, req *pb.ShareFolderCreateReq) (model.ShareFolderRsp, error)
	ShareFolderUpdate(ctx context.Context, req *pb.ShareFolderUpdateReq) (model.ShareFolderRsp, error)
	ShareFolderDestroy(ctx context.Context, req *pb.ShareFolderDestroyReq) (*pb.ShareFolderDestroyRsp, error)
	ShareFolderAddMember(ctx context.Context, req *pb.ShareFolderAddMemberReq) (*pb.ShareFolderAddMemberRsp, error)
	ShareFolderUpdateMember(ctx context.Context, req *pb.ShareFolderUpdateMemberReq) (*pb.ShareFolderUpdateMemberRsp, error)
	ShareFolderFolderList(ctx context.Context, req *pb.ShareFolderFolderListReq) (*pb.ShareFolderFolderListRsp, error)
	ShareFolderRecordList(ctx context.Context, req *pb.ShareFolderRecordListReq) (*pb.ShareFolderRecordListRsp, error)
	ShareFolderRecordListByRid(ctx context.Context, req *pb.ShareFolderRecordListByRidReq) (*pb.ShareFolderRecordListByRidRsp, error)
	ShareFolderAddRecord(ctx context.Context, req *pb.ShareFolderAddRecordReq) (*pb.ShareFolderAddRecordRsp, error)
	ShareFolderDeleteRecord(ctx context.Context, req *pb.ShareFolderDeleteRecordReq) (*pb.ShareFolderDeleteRecordRsp, error)
	ShareFolderMemberList(ctx context.Context, req *pb.ShareFolderMemberListReq) (*pb.ShareFolderMemberListRsp, error)
	ShareFolderDeleteMember(ctx context.Context, req *pb.ShareFolderDeleteMemberReq) (*pb.ShareFolderDeleteMemberRsp, error)
	ShareFolderMemberExit(ctx context.Context, req *pb.ShareFolderMemberExitReq) (*pb.ShareFolderMemberExitRsp, error)
	ShareFolderBatchUpdate(ctx context.Context, req *pb.ShareFolderBatchUpdateReq) (*pb.ShareFolderBatchUpdateRsp, error)
	FileUpload(ctx context.Context, req *pb.FileUploadReq) (model.FileUploadRsp, error)
	FileDownload(ctx context.Context, req *pb.FileDownloadReq) (model.FileDownLoadItemRsp, error)
	FileAttachment(ctx context.Context, req *pb.FileAttachmentReq) (model.FileAttachmentItem, error)
	FileReport(ctx context.Context, req *pb.FileReportReq) (model.FileReportRsp, error)

	VipGetConfig(ctx context.Context, req *pb.VipGetConfigReq) (*pb.VipGetConfigRsp, error)
	VipSubscriptionList(ctx context.Context, req *pb.VipSubscriptionListReq) (*pb.VipSubscriptionListRsp, error)
	VipPaymentList(ctx context.Context, req *pb.VipPaymentListReq) (*pb.VipPaymentListRsp, error)
	VipCreateOrder(ctx context.Context, req *pb.VipCreateOrderReq) (*pb.VipCreateOrderRsp, error)
	VipCheckOrder(ctx context.Context, req *pb.VipCheckOrderReq) (*pb.VipCheckOrderRsp, error)
	VipAppleVerifyReceipt(ctx context.Context, req *pb.VipAppleVerifyReceiptReq) (*pb.VipAppleVerifyReceiptRsp, error)

	GetDiscountCodeInfo(ctx context.Context, req *pb.GetDiscountCodeInfoReq) (*pb.GetDiscountCodeInfoRsp, error)
	GetOrderList(ctx context.Context, req *pb.GetOrderListReq) (*pb.GetOrderListRsp, error)
	VipIOSPromotionSign(ctx context.Context, req *pb.GetVipIOSPromotionSignReq) (*pb.GetVipIOSPromotionSignRsp, error)

	VipPrice(ctx context.Context, req *pb.VipPriceReq) (*pb.VipPriceRsp, error)
}

type dao struct {
	conf         *config.Config
	snow         *snowflake.Node
	responseWait sync.Map
	requestChan  map[string]chan *pb.StreamRsp
	clients      map[string]map[string]string //sign、storage、index
	lock         sync.RWMutex
}

func NewDAO(conf *config.Config) DAO {
	ip := getClientIP()
	long := big.NewInt(0).SetBytes(net.ParseIP(ip).To4()).Int64()
	nodeNum := rand.New(rand.NewSource(long)).Int63n(1023)
	node, err := snowflake.NewNode(nodeNum)
	if err != nil {
		log.Fatalf("GenerateSeq ID ip:%s nodeNum:%d err:%+v", ip, nodeNum, err)
	}
	log.Infof("GenerateSeq ID ip:%s nodeNum:%d", ip, nodeNum)
	d := &dao{
		conf:         conf,
		snow:         node,
		responseWait: sync.Map{},
		requestChan:  make(map[string]chan *pb.StreamRsp),
		clients:      make(map[string]map[string]string),
	}
	go d.heartBeat()
	return d
}

func getClientIP() string {
	ip := "127.0.0.1"
	address, err := net.InterfaceAddrs()
	if err != nil {
		return ip
	}
	for _, addr := range address {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				break
			}

		}
	}
	return ip
}
