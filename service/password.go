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

func (s *Service) CheckTx(ctx context.Context, req *pb.CheckTxReq) (*pb.CheckTxRsp, error) {
	rsp := new(pb.CheckTxRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.CheckTxRsp_Data{}
	params := model.CheckTxParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("checktx start", log.String("trace_id", trace_id), log.String("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("checktx params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("checktx params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("checktx signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.CheckTx(ctx, req)
	if err != nil {
		log.Logger.Error("checktx request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("checktx request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("checktx request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Height > 0 {
		rsp.Data.IsSuccess = model.StatusOK
	} else {
		rsp.Data.IsSuccess = model.StatusFAILED
	}

	log.Logger.Info("checktx success", log.String("trace_id", trace_id), log.Any("rsp", rsp))

	return rsp, nil
}

func (s *Service) BatchCheckTx(ctx context.Context, req *pb.BatchCheckTxReq) (*pb.BatchCheckTxRsp, error) {
	rsp := new(pb.BatchCheckTxRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.BatchCheckTxRsp_Data{}
	params := model.BatchCheckTxParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("batch checktx start", log.String("trace_id", trace_id), log.String("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("batch checktx params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(req.GetData()); !ok {
		rsp.Code = model.StatusTimestampErr
		rsp.Msg = errMsg
		log.Logger.Warn("batch checktx  params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Addr, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("batch checktx signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.BatchCheckTx(ctx, req)
	if err != nil {
		log.Logger.Error("batch checktx request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("batch checktx request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("batch checktx request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.Data
	log.Logger.Info("batch checktx end", log.String("trace_id", trace_id), log.Any("rsp", rsp))

	return rsp, nil
}

// AddCredential add/update credential
func (s *Service) AddCredential(ctx context.Context, req *pb.AddCredentialReq) (*pb.AddCredentialRsp, error) {
	rsp := new(pb.AddCredentialRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.AddCredentialRsp_Data{}
	params := model.AddCredentialParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("AddCredential start", log.String("trace_id", trace_id), log.String("params", req.GetParams()))
	log.Logger.Debug("AddCredential start debug", log.String("trace_id", trace_id), log.String("data", base64.StdEncoding.EncodeToString(req.GetData())))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("AddCredential params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("AddCredential params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}

	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("AddCredential signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.AddOrDelCredential(ctx, req.GetSignature(), req.GetParams(), req.GetData())
	if err != nil {
		log.Logger.Error("AddCredential request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("AddCredential request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("AddCredential request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = &pb.AddCredentialRsp_Data{TxHash: ret.TxHash}
	log.Logger.Info("AddCredential end", log.String("trace_id", trace_id), log.Any("rsp", rsp))
	return rsp, nil
}

func (s *Service) BatchAddCredential(ctx context.Context, req *pb.BatchAddCredentialReq) (*pb.BatchAddCredentialRsp, error) {
	rsp := new(pb.BatchAddCredentialRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.BatchAddCredentialRsp_Data{}
	params := model.BatchAddCredentialParams{}

	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("batch add credential start", log.String("trace_id", trace_id), log.String("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("batch add credential params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("batch add credential checn params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Addr, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("batch add credential signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.BatchAddCredential(ctx, req)
	if err != nil {
		log.Logger.Error("batch add credential request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("batch add credential request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("batch add credential request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.Data
	log.Logger.Info("batch add credential end", log.String("trace_id", trace_id), log.Any("rsp", rsp))

	return rsp, nil
}

// DeleteCredential delete credential
func (s *Service) DeleteCredential(ctx context.Context, req *pb.DeleteCredentialReq) (*pb.DeleteCredentialRsp, error) {
	rsp := new(pb.DeleteCredentialRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.DeleteCredentialRsp_Data{}
	params := model.DeleteCredentialParams{}

	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("delete credential body", log.String("trace_id", trace_id), log.String("params", req.GetParams()))
	log.Logger.Debug("delete credential body debug", log.String("trace_id", trace_id), log.String("data", base64.StdEncoding.EncodeToString(req.GetData())))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("delete credential params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("delete credential params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}

	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("delete credential signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.AddOrDelCredential(ctx, req.GetSignature(), req.GetParams(), req.GetData())
	if err != nil {
		log.Logger.Error("delete credential request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("delete credential request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("delete credential request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data.TxHash = ret.TxHash
	log.Logger.Info("delete credential end", log.String("trace_id", trace_id), log.Any("rsp", rsp))

	return rsp, nil
}

func (s *Service) BatchDeleteCredential(ctx context.Context, req *pb.BatchDeleteCredentialReq) (*pb.BatchDeleteCredentialRsp, error) {
	rsp := new(pb.BatchDeleteCredentialRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.BatchDeleteCredentialRsp_Data{}
	params := model.BatchDeleteCredentialParams{}

	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("batch delete credential start", log.String("trace_id", trace_id), log.String("params", req.GetParams()))
	log.Logger.Debug("batch delete credential start debug", log.String("trace_id", trace_id), log.String("data", base64.StdEncoding.EncodeToString(req.GetData())))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("batch delete credential params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("batch delete credential check params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Addr, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("batch delete credential signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.BatchDeleteCredential(ctx, req)
	if err != nil {
		log.Logger.Error("batch delete credential request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("batch delete credential request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("batch delete credential request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.Data
	log.Logger.Info("batch delete credential end", log.String("trace_id", trace_id), log.Any("rsp", rsp))

	return rsp, nil
}

// GetCredential get credential
func (s *Service) GetCredential(ctx context.Context, req *pb.GetCredentialReq) (*pb.GetCredentialRsp, error) {
	rsp := new(pb.GetCredentialRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.GetCredentialRsp_Item{}
	params := model.GetCredentialParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("get credential start", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("get credential params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("get credential check params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("get credential signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.GetPrimaryAddrIndexDetail(ctx, req.GetSignature(), req.GetParams())
	if err != nil {
		log.Logger.Error("get credential request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("get credential request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("get credential request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data.Id = ret.Id
	rsp.Data.OpTimestamp = ret.OpTimestamp
	rsp.Data.Credential = ret.Credential
	log.Logger.Info("get credential success", log.String("trace_id", trace_id))

	return rsp, nil
}

// DeleteAllCredential delete account
func (s *Service) DeleteAllCredential(ctx context.Context, req *pb.DeleteAllCredentialReq) (
	*pb.DeleteAllCredentialRsp, error) {
	rsp := new(pb.DeleteAllCredentialRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.EmptyData{}
	params := model.DeleteAllCredentialParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("delete all credential start", log.String("trace_id", trace_id), log.String("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("delete all credential params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("delete all credential check params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("delete all credential signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.DeleteAllCredential(ctx, req.GetSignature(), req.GetParams())
	if err != nil {
		log.Logger.Error("delete all credential request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("delete all credential request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("delete all credential request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	log.Logger.Info("delete all credential success", log.String("trace_id", trace_id), log.Any("rsp", rsp))

	return rsp, nil
}

// GetAllCredentialTimestamp
func (s *Service) GetAllCredentialTimestamp(ctx context.Context, req *pb.GetAllCredentialTimestampReq) (
	*pb.GetAllCredentialTimestampRsp, error) {
	rsp := new(pb.GetAllCredentialTimestampRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = make([]*pb.GetAllCredentialTimestampRsp_Item, 0)
	params := model.GetAllCredentialTimestampParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("GetAllCredentialTimestamp start", log.String("trace_id", trace_id), log.String("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("GetAllCredentialTimestamp params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("GetAllCredentialTimestamp check params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("GetAllCredentialTimestamp signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.GetAllCredentialTimestamp(ctx, req)
	if err != nil {
		log.Logger.Error("GetAllCredentialTimestamp request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("GetAllCredentialTimestampl request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("GetAllCredentialTimestamp request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	for _, item := range ret.List {
		rsp.Data = append(rsp.Data, &pb.GetAllCredentialTimestampRsp_Item{
			Id:          item.Id,
			OpTimestamp: item.OpTimestamp,
		})
	}

	log.Logger.Info("GetAllCredentialTimestamp success", log.String("trace_id", trace_id), log.Any("count", len(rsp.GetData())))

	return rsp, nil
}

// GetCredentialList .
func (s *Service) GetCredentialList(ctx context.Context, req *pb.GetCredentialListReq) (
	*pb.GetCredentialListRsp, error) {
	rsp := new(pb.GetCredentialListRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = make([]*pb.GetCredentialListRsp_Item, 0)
	params := model.GetCredentialListParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("GetCredentialList start", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("GetCredentialList params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("GetCredentialList params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}

	if !util.CheckSignature(params.Address, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("GetCredentialList  signature fail", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.GetPrimaryAddrIndexList(ctx, req)
	if err != nil {
		log.Logger.Error("GetCredentialList request service error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("GetCredentialList request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("GetCredentialList request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	for _, item := range ret.List {
		rsp.Data = append(rsp.Data, &pb.GetCredentialListRsp_Item{
			Id:          item.Id,
			OpTimestamp: item.OpTimestamp,
			Credential:  item.Credential,
		})
	}
	log.Logger.Debug("GetCredentialList success debug", log.String("trace_id", trace_id), log.Any("count", len(rsp.GetData())))
	return rsp, nil
}
