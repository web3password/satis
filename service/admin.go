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
	jsoniter "github.com/json-iterator/go"
	"github.com/web3password/satis/log"
	"github.com/web3password/satis/model"
	"github.com/web3password/satis/util"
	pb "github.com/web3password/w3p-protobuf/user"
)

// AdminAddMember add admin member
func (s *Service) AdminAddMember(ctx context.Context, req *pb.AdminAddMemberReq) (*pb.AdminAddMemberRsp, error) {
	rsp := new(pb.AdminAddMemberRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.EmptyData{}
	params := model.AdminCommonParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("admin add member body", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("admin add member params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("admin add member signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}

	if errMsg, ok := params.Check(model.AdminAddMemberToken, params.Timestamp, req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("admin add member params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	ret, err := s.dao.AdminAddMember(ctx, req)
	log.Logger.Info("admin add member service return", log.String("trace_id", trace_id), log.Any("ret", ret))
	if err != nil {
		log.Logger.Error("admin add member request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()), log.Any("rsp", rsp))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("admin add member request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("admin add member request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	log.Logger.Info("admin add member success", log.String("trace_id", trace_id), log.Any("ret", ret))

	return rsp, nil
}

// AdminAuthorization
func (s *Service) AdminAuthorization(ctx context.Context, req *pb.AdminAuthorizationReq) (*pb.AdminAuthorizationRsp, error) {
	rsp := new(pb.AdminAuthorizationRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	params := model.AdminAuthorizationParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("AdminAuthorization body", log.String("trace_id", trace_id), log.String("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("AdminAuthorizationReq params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !params.Check() {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = "params error"
		log.Logger.Warn("AdminAuthorizationReq params fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	if !util.CheckTimestamp(params.Timestamp) {
		rsp.Code = model.StatusTimestampErr
		rsp.Msg = model.MsgTimestamp
		log.Logger.Warn("AdminAuthorizationReq timestamp fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("AdminAuthorizationReq signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}

	ret, err := s.dao.AdminAuthorization(ctx, req.GetSignature(), req.GetParams())
	if err != nil {
		log.Logger.Error("AdminAuthorizationReq AdminAuthorization error", log.String("trace_id", trace_id))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("AdminAuthorizationReq request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("AdminAuthorizationReq request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = &pb.AdminAuthorizationRsp_Data{Authorization: ret.Authorization}
	log.Logger.Info("AdminAuthorization success", log.String("trace_id", trace_id), log.Any("rsp", rsp))
	return rsp, nil
}

// AdminUpdateMember admin update member
func (s *Service) AdminUpdateMember(ctx context.Context, req *pb.AdminUpdateMemberReq) (*pb.AdminUpdateMemberRsp, error) {
	rsp := new(pb.AdminUpdateMemberRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.EmptyData{}
	params := model.AdminCommonParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("AdminUpdateMember body", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("AdminUpdateMember params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("AdminUpdateMember signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}

	if errMsg, ok := params.Check(model.AdminUpdateMemberToken, params.Timestamp, req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("AdminUpdateMember params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	ret, err := s.dao.AdminUpdateMember(ctx, req)
	if err != nil {
		log.Logger.Error("AdminUpdateMember request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("AdminUpdateMember request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("AdminUpdateMember request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	log.Logger.Info("AdminUpdateMember success debug", log.String("trace_id", trace_id), log.Any("ret", ret))

	return rsp, nil
}

// AdminRemoveMember admin remove member
func (s *Service) AdminRemoveMember(ctx context.Context, req *pb.AdminRemoveMemberReq) (*pb.AdminRemoveMemberRsp, error) {
	rsp := new(pb.AdminRemoveMemberRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.EmptyData{}
	params := model.AdminCommonParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("AdminRemoveMember body", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))
	log.Logger.Debug("AdminRemoveMember debug", log.String("trace_id", trace_id), log.Any("params", req.GetParams()), log.String("data", base64.StdEncoding.EncodeToString(req.GetData())))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("AdminRemoveMember params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("AdminRemoveMember signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}

	log.Logger.Debug("AdminRemoveMember data", log.String("trace_id", trace_id))

	if errMsg, ok := params.Check(model.AdminRemoveMemberToken, params.Timestamp, req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("AdminRemoveMember params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	ret, err := s.dao.AdminRemoveMember(ctx, req)
	log.Logger.Debug("AdminRemoveMember ret", log.String("trace_id", trace_id), log.Any("ret", ret))
	if err != nil {
		log.Logger.Error("AdminRemoveMember request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("AdminRemoveMember request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("AdminRemoveMember request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	log.Logger.Info("AdminRemoveMember success", log.String("trace_id", trace_id), log.Any("ret", ret))
	return rsp, nil
}

// AdminGetMemberList admin get member list
func (s *Service) AdminGetMemberList(ctx context.Context, req *pb.AdminGetMemberListReq) (*pb.AdminGetMemberListRsp, error) {
	rsp := new(pb.AdminGetMemberListRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	params := model.AdminCommonParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("AdminGetMemberList body", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("AdminGetMemberList params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("AdminGetMemberList signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	if errMsg, ok := params.Check(model.AdminGetMemberListToken, params.Timestamp, []byte{}); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("AdminGetMemberList params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	ret, err := s.dao.AdminGetMemberList(ctx, req)
	if err != nil {
		log.Logger.Error("AdminGetMemberList request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("AdminGetMemberList request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("AdminGetMemberList request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.Data

	log.Logger.Info("AdminGetMemberList success", log.String("trace_id", trace_id))
	return rsp, nil
}

// GetAdminMnemonic get admin shared mnemonic
func (s *Service) GetAdminMnemonic(ctx context.Context, req *pb.GetAdminMnemonicReq) (*pb.GetAdminMnemonicRsp, error) {
	rsp := new(pb.GetAdminMnemonicRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.GetAdminMnemonicRsp_Data{}
	params := model.AdminCommonParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("GetAdminMnemonic body", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("GetAdminMnemonic params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("GetAdminMnemonic signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	if errMsg, ok := params.Check(model.AdminGetAdminMnemonicToken, params.Timestamp, []byte{}); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("GetAdminMnemonic params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	ret, err := s.dao.GetAdminMnemonic(ctx, req)
	if err != nil {
		log.Logger.Error("GetAdminMnemonic service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("GetAdminMnemonic request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("GetAdminMnemonic request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.Data
	log.Logger.Info("GetAdminMnemonic success", log.String("trace_id", trace_id))
	return rsp, nil
}

// AdminBatchImportMember admin batch import member
// batch add member use addMember api , this api is abandoned
func (s *Service) AdminBatchImportMember(ctx context.Context, req *pb.AdminBatchImportMemberReq) (*pb.AdminBatchImportMemberRsp, error) {
	rsp := new(pb.AdminBatchImportMemberRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	params := model.AdminCommonParams{}
	trace_id := util.GetTraceid(ctx)

	log.Logger.Info("AdminBatchImportMember body", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("AdminBatchImportMember params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("AdminBatchImportMember signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	if errMsg, ok := params.Check(model.AdminBatchImportMemberToken, params.Timestamp, []byte{}); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("AdminBatchImportMember params fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.AdminBatchImportMember(ctx, req)
	if err != nil {
		log.Logger.Error("AdminBatchImportMember service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("AdminBatchImportMember request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("AdminBatchImportMember request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.Data
	log.Logger.Info("AdminBatchImportMember success", log.String("trace_id", trace_id))
	return rsp, nil
}

func (s *Service) AdminUpdateOrgInfo(ctx context.Context, req *pb.AdminUpdateOrgInfoReq) (*pb.AdminUpdateOrgInfoRsp, error) {
	rsp := new(pb.AdminUpdateOrgInfoRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	params := model.AdminGetOrgInfoParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("AdminUpdateOrgInfo body", log.String("trace_id", trace_id), log.String("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("AdminUpdateOrgInfo params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !params.Check() {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("AdminUpdateOrgInfo params fail", log.String("trace_id", trace_id))

		return rsp, nil
	}
	if !util.CheckTimestamp(params.Timestamp) {
		rsp.Code = model.StatusTimestampErr
		rsp.Msg = model.MsgTimestamp
		log.Logger.Warn("AdminUpdateOrgInfo timestamp fail", log.String("trace_id", trace_id))

		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("AdminUpdateOrgInfo signature fail", log.String("trace_id", trace_id))

		return rsp, nil
	}
	ret, err := s.dao.AdminUpdateOrgInfo(ctx, req)
	if err != nil {
		log.Logger.Error("AdminUpdateOrgInfo service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("AdminBatchImportMember request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("AdminBatchImportMember request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.Data

	log.Logger.Info("AdminUpdateOrgInfo success", log.String("trace_id", trace_id))
	log.Logger.Debug("AdminUpdateOrgInfo success debug", log.String("trace_id", trace_id), log.Any("rsp", rsp))

	return rsp, nil
}
