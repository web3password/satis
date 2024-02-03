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

func (s *Service) FileUpload(ctx context.Context, req *pb.FileUploadReq) (*pb.FileUploadRsp, error) {
	rsp := new(pb.FileUploadRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.FileUploadRsp_Data{}
	params := model.FileUploadReqParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("FileUpload start", log.String("trace_id", trace_id), log.Any("params", req.GetParams()), log.Any("attach_len", len(req.GetData())))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("FileUpload params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(req.GetData()); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("FileUpload params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Addr, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("FileUpload checksignature failed", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.FileUpload(ctx, req)
	if err != nil {
		log.Logger.Error("FileUpload error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("FileUpload request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("FileUpload request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = &pb.FileUploadRsp_Data{Cid: ret.Cid}
	log.Logger.Info("FileUpload success", log.String("trace_id", trace_id), log.Any("rsp", rsp))

	return rsp, nil
}

func (s *Service) FileDownload(ctx context.Context, req *pb.FileDownloadReq) (*pb.FileDownloadRsp, error) {
	rsp := new(pb.FileDownloadRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.FileDownloadRsp_Data{}
	params := model.FileDownloadReqParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("FileDownload start", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("FileDownload params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("FileDownload params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Addr, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("FileDownload checksignature failed", log.String("trace_id", trace_id))

		return rsp, nil
	}

	ret, err := s.dao.FileDownload(ctx, req)
	if err != nil {
		log.Logger.Error("FileDownload server failed", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("FileDownload request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("FileDownload request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = &pb.FileDownloadRsp_Data{
		Cid:     ret.Cid,
		Content: ret.Content,
	}

	log.Logger.Info("FileDownload success", log.String("trace_id", trace_id), log.Any("attach_length", len(rsp.GetData().GetContent())))
	return rsp, nil
}

func (s *Service) FileAttachment(ctx context.Context, req *pb.FileAttachmentReq) (*pb.FileAttachmentRsp, error) {
	rsp := new(pb.FileAttachmentRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.FileAttachmentRsp_Data{}
	params := model.FileAttachmentReqParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("FileAttachment start", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("FileAttachment params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("FileAttachment params fail", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Addr, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("FileAttachment checksignature failed", log.String("trace_id", trace_id))
		return rsp, nil
	}

	ret, err := s.dao.FileAttachment(ctx, req)
	if err != nil {
		log.Logger.Error("FileAttachment result", log.String("trace_id", trace_id), log.Any("err", err))
		rsp.Data.Message = ret.Message
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("FileAttachment request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("FileAttachment request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = &pb.FileAttachmentRsp_Data{
		Success: ret.Success,
		Message: ret.Message,
		Cid:     ret.Cid,
		Content: ret.Content,
	}

	log.Logger.Info("FileAttachment success", log.String("trace_id", trace_id), log.Any("content_length", len(rsp.GetData().GetContent())))
	return rsp, nil
}

func (s *Service) FileReport(ctx context.Context, req *pb.FileReportReq) (*pb.FileReportRsp, error) {
	rsp := new(pb.FileReportRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"
	rsp.Data = &pb.EmptyData{}
	params := model.FileReportReqParams{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("FileReport start", log.String("trace_id", trace_id), log.Any("params", req.GetParams()))
	log.Logger.Debug("FileReport start debug", log.String("trace_id", trace_id), log.Any("data", base64.StdEncoding.EncodeToString(req.GetData())))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("FileReport params parse fail", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}
	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("FileReport params error", log.String("trace_id", trace_id), log.Any("errMsg", errMsg))
		return rsp, nil
	}
	if !util.CheckSignature(params.Addr, req.GetSignature(), req.GetParams()) {
		rsp.Code = model.StatusSignatureErr
		rsp.Msg = model.MsgSignatureErr
		log.Logger.Warn("FileReport checksignature failed", log.String("trace_id", trace_id))
		return rsp, nil
	}
	ret, err := s.dao.FileReport(ctx, req)
	if err != nil {
		log.Logger.Error("FileReport error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("FileReport request error", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("FileReport request failed", log.String("trace_id", trace_id), log.Any("ret", ret))
		return rsp, nil
	}

	log.Logger.Info("FileReport success", log.String("trace_id", trace_id), log.Any("rsp", rsp))
	return rsp, nil
}
