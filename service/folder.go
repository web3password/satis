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
	"errors"
	"time"

	"github.com/web3password/satis/util"

	jsoniter "github.com/json-iterator/go"
	"github.com/web3password/satis/log"
	"github.com/web3password/satis/model"
	pb "github.com/web3password/w3p-protobuf/user"
)

// ShareFolderCreate
func (s *Service) ShareFolderCreate(ctx context.Context, req *pb.ShareFolderCreateReq) (*pb.ShareFolderCreateRsp, error) {
	rsp := new(pb.ShareFolderCreateRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.EmptyData{}
	params := model.ShareFolderParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("ShareFolderCreate start", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))
	log.Logger.Debug("ShareFolderCreate start debug", log.String("trace_id", trace_id), log.String("data", base64.StdEncoding.EncodeToString(req.GetData())))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("ShareFolderCreate params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("ShareFolderCreate params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("ShareFolderCreate signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.ShareFolderCreate(ctx, req)
	if err != nil {
		log.Logger.Error("ShareFolderCreate request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("ShareFolderCreate request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderCreate request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	log.Logger.Info("ShareFolderCreate success", log.String("trace_id", trace_id), log.Any("rsp", rsp))
	return rsp, nil
}

// ShareFolderUpdate
func (s *Service) ShareFolderUpdate(ctx context.Context, req *pb.ShareFolderUpdateReq) (*pb.ShareFolderUpdateRsp, error) {
	rsp := new(pb.ShareFolderUpdateRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.EmptyData{}
	params := model.ShareFolderUpdateParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("ShareFolderUpdate start", log.String("trace_id", trace_id), log.String("params", req.GetParams()))
	log.Logger.Debug("ShareFolderUpdate start debug", log.String("trace_id", trace_id), log.String("data", base64.StdEncoding.EncodeToString(req.GetData())))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("ShareFolderUpdate params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(req.GetData()); !ok {
		rsp.Code = model.StatusTimestampErr
		rsp.Msg = errMsg
		log.Logger.Warn("ShareFolderUpdate params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Error("ShareFolderUpdate signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.ShareFolderUpdate(ctx, req)
	if err != nil {
		log.Logger.Error("ShareFolderUpdate request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("ShareFolderUpdate request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderUpdate request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	log.Logger.Info("ShareFolderUpdate success", log.String("trace_id", trace_id), log.Any("rsp", rsp), log.Any("ret", ret))
	return rsp, nil
}

// ShareFolderDestroy .
func (s *Service) ShareFolderDestroy(ctx context.Context, req *pb.ShareFolderDestroyReq) (*pb.ShareFolderDestroyRsp, error) {
	rsp := new(pb.ShareFolderDestroyRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.EmptyData{}
	params := model.ShareFolderCommonParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("ShareFolderDestroy start", log.String("trace_id", trace_id), log.String("params", req.GetParams()))
	log.Logger.Debug("ShareFolderDestroy start debug", log.String("trace_id", trace_id), log.String("data", base64.StdEncoding.EncodeToString(req.GetData())))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("ShareFolderDestroy params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("ShareFolderDestroy signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	if errMsg, ok := params.Check(model.ShareFolderDestroyToken, params.Timestamp, req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("ShareFolderDestroy timestamp fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.ShareFolderDestroy(ctx, req)
	if err != nil {
		log.Logger.Error("ShareFolderDestroy request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("ShareFolderDestroy request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderDestroy request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	log.Logger.Info("ShareFolderDestroy success", log.String("trace_id", trace_id), log.Any("rsp", rsp), log.Any("ret", ret))
	return rsp, nil
}

func (s *Service) ShareFolderAddMember(ctx context.Context, req *pb.ShareFolderAddMemberReq) (*pb.ShareFolderAddMemberRsp, error) {
	rsp := new(pb.ShareFolderAddMemberRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.EmptyData{}
	params := model.ShareFolderCommonParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("ShareFolderAddMember start", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))
	log.Logger.Debug("ShareFolderAddMember start debug", log.String("trace_id", trace_id), log.String("data", base64.StdEncoding.EncodeToString(req.GetData())))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("ShareFolderAddMember params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("ShareFolderDestroy signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}

	if errMsg, ok := params.Check(model.ShareFolderAddMemberToken, params.Timestamp, req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("ShareFolderDestroy timestamp fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.ShareFolderAddMember(ctx, req)
	log.Logger.Debug("ShareFolderAddMember ret", log.String("trace_id", trace_id), log.Any("ret", ret))
	if err != nil {
		log.Logger.Error("ShareFolderAddMember request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("ShareFolderAddMember request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderAddMember request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	log.Logger.Info("ShareFolderAddMember success debug", log.String("trace_id", trace_id), log.Any("rsp", rsp), log.Any("ret", ret))
	return rsp, nil
}

// ShareFolderUpdateMember .
func (s *Service) ShareFolderUpdateMember(ctx context.Context, req *pb.ShareFolderUpdateMemberReq) (*pb.ShareFolderUpdateMemberRsp, error) {
	rsp := new(pb.ShareFolderUpdateMemberRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.EmptyData{}
	params := model.ShareFolderCommonParams{}
	trace_id := util.GetTraceid(ctx)

	log.Logger.Info("ShareFolderUpdateMember start", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))
	log.Logger.Debug("ShareFolderUpdateMember start debug", log.String("trace_id", trace_id), log.String("data", base64.StdEncoding.EncodeToString(req.GetData())))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("ShareFolderUpdateMember params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("ShareFolderUpdateMember signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	if errMsg, ok := params.Check(model.ShareFolderUpdateMemberToken, params.Timestamp, req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("ShareFolderUpdateMember timestamp fail,", log.String("errMsg", errMsg), log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.ShareFolderUpdateMember(ctx, req)
	if err != nil {
		log.Logger.Error("ShareFolderUpdateMember request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("ShareFolderUpdateMember request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderUpdateMember request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	log.Logger.Info("ShareFolderUpdateMember success", log.String("trace_id", trace_id), log.Any("rsp", rsp), log.Any("ret", ret))
	return rsp, nil
}

// ShareFolderAddRecord .
func (s *Service) ShareFolderAddRecord(ctx context.Context, req *pb.ShareFolderAddRecordReq) (*pb.ShareFolderAddRecordRsp, error) {
	rsp := new(pb.ShareFolderAddRecordRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.EmptyData{}
	params := model.ShareFolderAddRecordParams{}
	trace_id := util.GetTraceid(ctx)

	log.Logger.Info("sharefolder add record body", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("sharefolder add record params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
	}
	if errMsg, ok := params.Check(req.GetData()); !ok {
		rsp.Code = model.StatusTimestampErr
		rsp.Msg = errMsg
		log.Logger.Warn("sharefolder add record params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("sharefolder add record signature fail", log.String("trace_id", trace_id), log.Any("now", time.Now().Unix()))
		return rsp, nil
	}
	ret, err := s.dao.ShareFolderAddRecord(ctx, req)
	if err != nil {
		log.Logger.Error("sharefolder add record request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("sharefolder add record request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("sharefolder add record request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	log.Logger.Info("sharefolder add record success", log.String("trace_id", trace_id), log.Any("rsp", rsp), log.Any("ret", ret))
	return rsp, nil
}

// ShareFolderDeleteRecord .
func (s *Service) ShareFolderDeleteRecord(ctx context.Context, req *pb.ShareFolderDeleteRecordReq) (*pb.ShareFolderDeleteRecordRsp, error) {
	rsp := new(pb.ShareFolderDeleteRecordRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.EmptyData{}
	params := model.ShareFolderDeleteRecordParams{}
	trace_id := util.GetTraceid(ctx)

	log.Logger.Info("sharefolder delete record body", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("sharefolder delete record params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return nil, errors.New("unmarshal json failed")
	}

	if errMsg, ok := params.Check(req.GetData()); !ok {
		rsp.Code = model.StatusTimestampErr
		rsp.Msg = errMsg
		log.Logger.Warn("sharefolder delete record params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}

	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("sharefolder delete record signature fail", log.String("trace_id", trace_id), log.Any("now", time.Now().Unix()))
		return rsp, nil
	}
	ret, err := s.dao.ShareFolderDeleteRecord(ctx, req)
	if err != nil {
		log.Logger.Error("sharefolder delete record request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("sharefolder delete record request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("sharefolder delete record request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}
	log.Logger.Info("sharefolder delete record success", log.String("trace_id", trace_id), log.Any("rsp", rsp), log.Any("ret", ret))
	return rsp, nil
}

// ShareFolderFolderList .
func (s *Service) ShareFolderFolderList(ctx context.Context, req *pb.ShareFolderFolderListReq) (*pb.ShareFolderFolderListRsp, error) {
	rsp := new(pb.ShareFolderFolderListRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	trace_id := util.GetTraceid(ctx)
	params := model.ShareFolderCommonParams{}

	log.Logger.Info("sharefolder folder list body", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("sharefolder folder list params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("sharefolder folder list signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	if errMsg, ok := params.Check(model.ShareFolderFolderListToken, params.Timestamp, []byte{}); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("sharefolder folder list params fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.ShareFolderFolderList(ctx, req)
	if err != nil {
		log.Logger.Error("sharefolder folder list get list error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("sharefolder folder list request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("sharefolder folder list request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.Data

	log.Logger.Info("sharefolder folder list success", log.String("trace_id", trace_id), log.Any("folder_count", len(rsp.GetData().GetList())))

	return rsp, nil
}

// ShareFolderRecordList .
func (s *Service) ShareFolderRecordList(ctx context.Context, req *pb.ShareFolderRecordListReq) (*pb.ShareFolderRecordListRsp, error) {
	rsp := new(pb.ShareFolderRecordListRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	params := model.ShareFolderCommonParams{}
	trace_id := util.GetTraceid(ctx)

	log.Logger.Info("sharefolder record list start", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("sharefolder record list params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("sharefolder record list signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	if errMsg, ok := params.Check(model.ShareFolderRecordListToken, params.Timestamp, req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("sharefolder record list params fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.ShareFolderRecordList(ctx, req)
	if err != nil {
		log.Logger.Error("sharefolder record list get list error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("sharefolder record list request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("sharefolder record list request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.Data

	log.Logger.Info("sharefolder record list end", log.String("trace_id", trace_id), log.Any("record_count", len(rsp.GetData().GetList())))
	return rsp, nil
}

// ShareFolderRecordList .
func (s *Service) ShareFolderRecordListByRid(ctx context.Context, req *pb.ShareFolderRecordListByRidReq) (*pb.ShareFolderRecordListByRidRsp, error) {
	rsp := new(pb.ShareFolderRecordListByRidRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	params := model.ShareFolderCommonParams{}
	trace_id := util.GetTraceid(ctx)

	log.Logger.Info("ShareFolderRecordListByRid start", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("ShareFolderRecordListByRid params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("ShareFolderRecordListByRid params fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	if errMsg, ok := params.Check(model.ShareFolderRecordListTokenByRid, params.Timestamp, req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("ShareFolderRecordListByRid params fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.ShareFolderRecordListByRid(ctx, req)
	if err != nil {
		log.Logger.Error("ShareFolderRecordListByRid error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("ShareFolderRecordListByRid request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderRecordListByRid request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.Data

	log.Logger.Info("ShareFolderRecordListByRid end", log.String("trace_id", trace_id), log.Any("count", len(rsp.GetData().GetList())))

	return rsp, nil
}

// ShareFolderMemberList .
func (s *Service) ShareFolderMemberList(ctx context.Context, req *pb.ShareFolderMemberListReq) (*pb.ShareFolderMemberListRsp, error) {
	rsp := new(pb.ShareFolderMemberListRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	params := model.ShareFolderCommonParams{}
	trace_id := util.GetTraceid(ctx)

	log.Logger.Info("ShareFolderMemberList start", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("ShareFolderMemberList params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("ShareFolderMemberList signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	if errMsg, ok := params.Check(model.ShareFolderMemberListToken, params.Timestamp, []byte{}); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("ShareFolderMemberList params fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.ShareFolderMemberList(ctx, req)
	if err != nil {
		log.Logger.Error("ShareFolderMemberList error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("ShareFolderMemberList request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderMemberList request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.Data

	log.Logger.Info("ShareFolderMemberList end", log.String("trace_id", trace_id), log.Any("count", len(rsp.GetData().GetList())))

	return rsp, nil
}

// ShareFolderDeleteMember .
func (s *Service) ShareFolderDeleteMember(ctx context.Context, req *pb.ShareFolderDeleteMemberReq) (*pb.ShareFolderDeleteMemberRsp, error) {
	rsp := new(pb.ShareFolderDeleteMemberRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.EmptyData{}
	params := model.ShareFolderCommonParams{}
	trace_id := util.GetTraceid(ctx)

	log.Logger.Info("ShareFolderDeleteMember start", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))
	log.Logger.Debug("ShareFolderDeleteMember start debug", log.String("trace_id", trace_id), log.String("data", base64.StdEncoding.EncodeToString(req.GetData())))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("ShareFolderDeleteMember params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("ShareFolderDeleteMember signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}

	if errMsg, ok := params.Check(model.ShareFolderDeleteMemberToken, params.Timestamp, req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("ShareFolderDeleteMember params fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.ShareFolderDeleteMember(ctx, req)
	if err != nil {
		log.Logger.Error("ShareFolderDeleteMember request service error", log.String("trace_id", trace_id), log.Any("rsp", rsp), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("ShareFolderDeleteMember request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderDeleteMember request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	log.Logger.Info("ShareFolderDeleteMember end", log.String("trace_id", trace_id))
	return rsp, nil
}

// ShareFolderMemberExit .
func (s *Service) ShareFolderMemberExit(ctx context.Context, req *pb.ShareFolderMemberExitReq) (*pb.ShareFolderMemberExitRsp, error) {
	rsp := new(pb.ShareFolderMemberExitRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.EmptyData{}
	params := model.ShareFolderCommonParams{}
	trace_id := util.GetTraceid(ctx)

	log.Logger.Info("ShareFolderMemberExit start", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("ShareFolderMemberExit params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("ShareFolderMemberExit signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}

	if errMsg, ok := params.Check(model.ShareFolderMemberExitToken, params.Timestamp, req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("ShareFolderMemberExit params fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.ShareFolderMemberExit(ctx, req)
	if err != nil {
		log.Logger.Error("ShareFolderMemberExit request service error", log.String("trace_id", trace_id), log.Any("rsp", rsp), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("ShareFolderMemberExit request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderMemberExit request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	log.Logger.Info("ShareFolderMemberExit end", log.String("trace_id", trace_id))
	log.Logger.Debug("ShareFolderMemberExit ret", log.String("trace_id", trace_id), log.Any("ret", ret))
	return rsp, nil
}

// ShareFolderBatchUpdate .
func (s *Service) ShareFolderBatchUpdate(ctx context.Context, req *pb.ShareFolderBatchUpdateReq) (*pb.ShareFolderBatchUpdateRsp, error) {
	rsp := new(pb.ShareFolderBatchUpdateRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.EmptyData{}
	params := model.ShareFolderCommonParams{}
	trace_id := util.GetTraceid(ctx)

	log.Logger.Info("ShareFolderBatchUpdate start", log.String("trace_id", trace_id), log.String("params", req.GetParams()))
	log.Logger.Debug("ShareFolderBatchUpdate start debug", log.String("trace_id", trace_id), log.String("data", base64.StdEncoding.EncodeToString(req.GetData())))

	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("ShareFolderBatchUpdate params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("ShareFolderBatchUpdate params fail", log.String("trace_id", trace_id))
		return rsp, nil
	}

	if errMsg, ok := params.Check(model.ShareFolderBatchUpdateToken, params.Timestamp, req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("ShareFolderBatchUpdate params fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.ShareFolderBatchUpdate(ctx, req)
	log.Logger.Debug("ShareFolderBatchUpdate ret", log.String("trace_id", trace_id), log.Any("ret", ret))
	if err != nil {
		log.Logger.Error("ShareFolderBatchUpdate request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg
	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("ShareFolderBatchUpdate request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderBatchUpdate request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	log.Logger.Info("ShareFolderBatchUpdate end", log.String("trace_id", trace_id))
	return rsp, nil
}
