/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/
package service

import (
	"context"
	"encoding/base64"
	"github.com/web3password/satis/util"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/web3password/satis/log"
	"github.com/web3password/satis/model"
	pb "github.com/web3password/w3p-protobuf/user"
)

// Register register user
func (s *Service) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRsp, error) {
	rsp := new(pb.RegisterRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = model.MsgSystemErr
	rsp.Data = &pb.EmptyData{}
	params := model.RegisterParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("user register start", log.String("trace_id", trace_id), log.Any("req", req))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("user register params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("user register params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("user register signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}

	ret, err := s.dao.RegisterUser(ctx, req.GetSignature(), req.GetParams())
	if err != nil {
		log.Logger.Error("user register request service error", log.String("trace_id", trace_id), log.Any("ret", ret), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("user register request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("user register request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	log.Logger.Info("user register success", log.String("trace_id", trace_id), log.Any("rsp", rsp))
	return rsp, nil
}

// GetPersonalSignAddress get personal sign address
func (s *Service) GetPersonalSignAddress(ctx context.Context, req *pb.GetPersonalSignAddressReq) (*pb.GetPersonalSignAddressRsp, error) {
	rsp := new(pb.GetPersonalSignAddressRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = model.MsgSystemErr
	rsp.Data = &pb.GetPersonalSignAddressRsp_Data{}
	params := model.GetPersonalSignAddressParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("GetPersonalSignAddress start", log.String("trace_id", trace_id), log.Any("req", req))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("GetPersonalSignAddress params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("GetPersonalSignAddress params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("GetPersonalSignAddress signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.GetPersonalSignAddress(ctx, req.GetSignature(), req.GetParams())
	if err != nil {
		log.Logger.Error("GetPersonalSignAddress request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("GetPersonalSignAddress request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("GetPersonalSignAddress request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	for _, p := range ret.List {
		rsp.Data.Signs = append(rsp.Data.Signs, &pb.GetPersonalSignAddressRsp_Sign{
			Sign:   p.Sign,
			Params: p.Params,
		})
	}

	log.Logger.Info("GetPersonalSignAddress success", log.String("trace_id", trace_id), log.Any("rsp", rsp))
	return rsp, nil
}

// GetVIPInfo get vip info
func (s *Service) GetVIPInfo(ctx context.Context, req *pb.GetVIPInfoReq) (*pb.GetVIPInfoRsp, error) {
	rsp := new(pb.GetVIPInfoRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = model.MsgSystemErr
	rsp.Data = &pb.GetVIPInfoRsp_Data{}
	params := model.GetVIPInfoParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("GetVIPInfo start", log.String("trace_id", trace_id), log.Any("req", req))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("GetVIPInfo params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("GetVIPInfo params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("GetVIPInfo signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.GetVIPInfo(ctx, req.GetSignature(), req.GetParams())
	if err != nil {
		log.Logger.Error("GetVIPInfo request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	log.Logger.Debug("GetVIPInfo request service result", log.String("trace_id", trace_id), log.Any("ret.code", ret.Code))

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("GetVIPInfo request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("GetVIPInfo request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data.Signature = ret.Data.Signature
	rsp.Data.Params = ret.Data.Params
	list := make([]*pb.GetVIPInfoRsp_Data_List, 0)
	for _, info := range ret.Data.OrgInfo {
		orgList := &pb.GetVIPInfoRsp_Data_List{
			OrgId:               info.OrgId,
			OrgName:             info.OrgName,
			Logo:                info.Logo,
			SelfHostUrl:         info.SelfHostUrl,
			MemberShareMnemonic: info.SharedMnemonic,
		}
		list = append(list, orgList)
	}
	rsp.Data.OrgInfo = list
	log.Logger.Info("GetVIPInfo success", log.String("trace_id", trace_id), log.Any("org_count", len(rsp.GetData().GetOrgInfo())))
	return rsp, nil
}

// GetUserInfo GetUserInfo
func (s *Service) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (*pb.GetUserInfoRsp, error) {
	rsp := new(pb.GetUserInfoRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = model.MsgSystemErr
	rsp.Data = &pb.GetUserInfoRsp_Data{}
	params := model.GetUserInfoParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("GetUserInfo start", log.String("trace_id", trace_id), log.Any("req", req))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("GetUserInfo params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("GetUserInfo params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("GetUserInfo signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.GetUserInfo(ctx, req.GetSignature(), req.GetParams())
	if err != nil {
		log.Logger.Error("GetUserInfo request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("GetUserInfo request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("GetUserInfo request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data.Addr = ret.Data.Address
	rsp.Data.ChainId = ret.Data.ChainId
	rsp.Data.InviteCode = ret.Data.InviteCode
	rsp.Data.StorageType = ret.Data.StorageType
	log.Logger.Info("GetUserInfo success", log.String("trace_id", trace_id))
	log.Logger.Debug("GetUserInfo success debug", log.String("trace_id", trace_id), log.Any("rsp", rsp))

	return rsp, nil
}

// GetLatestBlockTimestamp get block time
func (s *Service) GetLatestBlockTimestamp(ctx context.Context, req *pb.GetLatestBlockTimestampReq) (
	*pb.GetLatestBlockTimestampRsp, error) {
	rsp := new(pb.GetLatestBlockTimestampRsp)
	rsp.Code = model.StatusOK
	rsp.Msg = model.MsgOK
	rsp.Data = &pb.GetLatestBlockTimestampRsp_Data{
		Timestamp: int32(time.Now().Unix()),
	}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Debug("GetLatestBlockTimestamp response", log.String("trace_id", trace_id), log.Any("req", req), log.Any("rsp", rsp))
	return rsp, nil
}

func (s *Service) StorageReport(ctx context.Context, req *pb.StorageReportReq) (*pb.StorageReportRsp, error) {
	rsp := new(pb.StorageReportRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = model.MsgSystemErr
	params := model.StorageReportParams{}
	trace_id := util.GetTraceid(ctx)

	log.Logger.Info("StorageReport start", log.String("trace_id", trace_id), log.Any("req", req))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("StorageReport params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("StorageReport params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("StorageReport signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}

	ret, err := s.dao.StorageReport(ctx, req.GetSignature(), req.GetParams())
	if err != nil {
		log.Logger.Error("StorageReport error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("StorageReport request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("StorageReport request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	log.Logger.Info("StorageReport success", log.String("trace_id", trace_id), log.Any("rsp", rsp))
	return rsp, nil
}

func (s *Service) StorageStat(ctx context.Context, req *pb.StorageStatReq) (*pb.StorageStatRsp, error) {
	rsp := new(pb.StorageStatRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = model.MsgSystemErr
	params := model.StorageStatParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("StorageStat start", log.String("trace_id", trace_id), log.Any("req", req))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr

		log.Logger.Warn("StorageStat parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("StorageStat params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("StorageStat signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}

	ret, err := s.dao.StorageStat(ctx, req.GetSignature(), req.GetParams())
	if err != nil {
		log.Logger.Error("StorageStat request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("StorageStat request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("StorageStat request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.Data
	log.Logger.Info("StorageStat success", log.String("trace_id", trace_id), log.Any("rsp", rsp))
	return rsp, nil
}

// Register register user
func (s *Service) GetVersionDesc(ctx context.Context, req *pb.GetVersionDescReq) (*pb.GetVersionDescRsp, error) {
	rsp := new(pb.GetVersionDescRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = model.MsgSystemErr
	rsp.Data = &pb.GetVersionDescRsp_Data{}
	params := model.GetVersionDescParams{}
	trace_id := util.GetTraceid(ctx)

	log.Logger.Info("GetVersionDesc start", log.String("trace_id", trace_id), log.Any("req", req))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("StorageStat parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !params.Check() {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("GetVersionDesc params fail", log.String("trace_id", trace_id))
		return rsp, nil
	}

	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("GetVersionDesc signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}

	ret, err := s.dao.GetVersionDesc(ctx, req.GetSignature(), req.GetParams())
	if err != nil {
		log.Logger.Error("GetVersionDesc error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("GetVersionDesc request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("GetVersionDesc request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data.VersionDesc = ret.VersionDesc

	log.Logger.Info("GetVersionDesc success", log.String("trace_id", trace_id), log.Any("rsp", rsp))

	return rsp, nil
}

func (s *Service) AdminRegister(ctx context.Context, req *pb.AdminRegisterReq) (*pb.AdminRegisterRsp, error) {
	rsp := new(pb.AdminRegisterRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = model.MsgSystemErr
	params := model.AdminRegisterParams{}
	trace_id := util.GetTraceid(ctx)

	log.Logger.Info("AdminRegister start", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))
	log.Logger.Debug("AdminRegister start debug", log.String("trace_id", trace_id), log.Any("data", base64.StdEncoding.EncodeToString(req.GetData())))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("AdminRegister parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("AdminRegister params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	/*if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Error("AdminRegister signature error", log.String("trace_id", trace_id), log.Any("req", req))
		return rsp, nil
	}*/
	ret, err := s.dao.AdminRegister(ctx, req)
	if err != nil {
		log.Logger.Error("AdminRegister request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("AdminRegister request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("AdminRegister request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.Data
	log.Logger.Info("AdminRegister success", log.String("trace_id", trace_id), log.Any("rsp", rsp))
	return rsp, nil
}

func (s *Service) AdminTransferSuperAdmin(ctx context.Context, req *pb.AdminTransferSuperAdminReq) (*pb.AdminTransferSuperAdminRsp, error) {
	rsp := new(pb.AdminTransferSuperAdminRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = model.MsgSystemErr
	params := model.TransferSuperAdminParams{}
	trace_id := util.GetTraceid(ctx)

	log.Logger.Info("AdminTransferSuperAdmin start", log.String("trace_id", trace_id), log.Any("params", req.GetParams()),
		log.Any("data", base64.StdEncoding.EncodeToString(req.GetData())))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("AdminTransferSuperAdmin parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !params.Check() {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("AdminTransferSuperAdmin params fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("AdminTransferSuperAdmin signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.AdminTransferSuperAdmin(ctx, req)
	if err != nil {
		log.Logger.Error("AdminTransferSuperAdmin transfer error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("AdminTransferSuperAdmin request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("AdminTransferSuperAdmin request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.Data
	log.Logger.Info("AdminTransferSuperAdmin success", log.String("trace_id", trace_id), log.Any("rsp", rsp))
	return rsp, nil
}

func (s *Service) AdminOperationHistory(ctx context.Context, req *pb.AdminOperationHistoryReq) (*pb.AdminOperationHistoryRsp, error) {
	rsp := new(pb.AdminOperationHistoryRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = model.MsgSystemErr
	params := model.OperationHistoryParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("AdminOperationHistory start", log.String("trace_id", trace_id), log.String("params", req.GetParams()))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("AdminOperationHistory parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !params.Check() {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("AdminOperationHistory params fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	if !util.CheckTimestamp(params.Timestamp) {
		rsp.Code = model.StatusTimestampErr
		rsp.Msg = model.MsgTimestamp
		log.Logger.Warn("AdminOperationHistory timestamp fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("AdminOperationHistory signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.AdminOperationHistory(ctx, req)
	if err != nil {
		log.Logger.Error("AdminOperationHistory get error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("AdminOperationHistory request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("AdminOperationHistory request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.Data
	log.Logger.Info("AdminOperationHistory success", log.String("trace_id", trace_id))
	return rsp, nil
}

func (s *Service) AdminGetOrgInfo(ctx context.Context, req *pb.AdminGetOrgInfoReq) (*pb.AdminGetOrgInfoRsp, error) {
	rsp := new(pb.AdminGetOrgInfoRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = model.MsgSystemErr
	params := model.AdminGetOrgInfoParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("AdminGetOrgInfo start", log.String("trace_id", trace_id), log.String("params", req.GetParams()))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("AdminGetOrgInfo parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	if !params.Check() {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("AdminGetOrgInfo params fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	if !util.CheckTimestamp(params.Timestamp) {
		rsp.Code = model.StatusTimestampErr
		rsp.Msg = model.MsgTimestamp
		log.Logger.Warn("AdminGetOrgInfo timestamp fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("AdminGetOrgInfo signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.AdminGetOrgInfo(ctx, req)
	if err != nil {
		log.Logger.Error("AdminGetOrgInfo get error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("AdminGetOrgInfo request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("AdminGetOrgInfo request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.Data
	log.Logger.Info("AdminGetOrgInfo success", log.String("trace_id", trace_id))
	return rsp, nil
}

// GetVersionConfig get version config info
func (s *Service) GetVersionConfig(ctx context.Context, req *pb.GetVersionConfigReq) (*pb.GetVersionConfigRsp, error) {
	rsp := new(pb.GetVersionConfigRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.GetVersionConfigRsp_Data{}
	params := model.GetVersionConfigParams{}
	trace_id := util.GetTraceid(ctx)

	log.Logger.Info("GetVersionConfig start", log.String("trace_id", trace_id), log.String("params", req.GetParams()))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("GetVersionConfig params parse fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	if errMsg, ok := params.Check(model.GetVersionConfigToken, params.Timestamp, req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("GetVersionConfig params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("GetVersionConfig signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.GetVersionConfig(ctx, req)
	if err != nil {
		log.Logger.Error("GetVersionConfig service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("GetVersionConfig request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("GetVersionConfig request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.Data

	log.Logger.Info("GetVersionConfig success", log.String("trace_id", trace_id))

	return rsp, nil
}
