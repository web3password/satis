/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/

package dao

import (
	"context"
	"errors"
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/web3password/satis/log"
	"github.com/web3password/satis/model"
	"github.com/web3password/satis/util"
	storageProto "github.com/web3password/w3p-protobuf/storage"
	userProto "github.com/web3password/w3p-protobuf/user"
)

// FileUpload .
func (d *dao) FileUpload(ctx context.Context, req *userProto.FileUploadReq) (model.FileUploadRsp, error) {
	requestID := d.GenerateID()
	rsp := model.FileUploadRsp{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("FileUpload start", log.String("trace_id", trace_id), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&userProto.StreamRsp{
		Cmd:       model.CMDFileUpload,
		Token:     model.FileUploadToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   trace_id,
	}, model.STORAGE_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("FileUpload add proxy request error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	timer := time.NewTimer(model.W3PTimeoutFileAttachment * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := storageProto.UploadReply{}
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("FileUpload response json unmarshal error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()))
			return rsp, err
		}
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg

		if rsp.Code > model.StatusSystemErrorCode {
			log.Logger.Error("FileUpload response rsp error", log.String("trace_id", trace_id), log.Any("ret", ret))
			return rsp, errors.New(rsp.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("FileUpload response not good", log.String("trace_id", trace_id), log.Any("ret", ret))
			return rsp, nil
		}

		rsp.Cid = ret.GetData().GetCid()
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("FileUpload response timeout", log.String("trace_id", trace_id))
		return rsp, fmt.Errorf("fileupload response timeout")
	}

	log.Logger.Info("FileUpload end", log.String("trace_id", trace_id), log.Any("rsp", rsp))
	return rsp, nil
}

// FileDownload .
func (d *dao) FileDownload(ctx context.Context, req *userProto.FileDownloadReq) (model.FileDownLoadItemRsp, error) {
	requestID := d.GenerateID()
	rsp := model.FileDownLoadItemRsp{}
	trace_id := util.GetTraceid(ctx)

	log.Logger.Info("FileDownload start", log.String("trace_id", trace_id), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&userProto.StreamRsp{
		Cmd:       model.CMDFileDownload,
		Token:     model.FileDownloadToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		TraceId:   trace_id,
	}, model.STORAGE_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("FileDownload add proxy request error", log.String("errmsg", err.Error()))
		return rsp, err
	}

	timer := time.NewTimer(model.W3PTimeoutFileAttachment * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := storageProto.DownloadReply{}
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("FileDownload response json unmarshal error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg

		if rsp.Code > model.StatusSystemErrorCode {
			log.Logger.Error("FileDownload response rsp error", log.String("trace_id", trace_id), log.Any("rsp", rsp))
			return rsp, errors.New(rsp.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("FileDownload response not good", log.String("trace_id", trace_id), log.Any("rsp", rsp))
			return rsp, nil
		}

		rsp.Cid = ret.GetData().GetCid()
		rsp.Content = ret.GetData().GetContent()
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("FileDownload response timeout", log.String("trace_id", trace_id), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("FileDownload  response timeout")
	}

	log.Logger.Info("FileDownload end", log.String("trace_id", trace_id), log.Any("attach_length", len(rsp.Content)), log.Any("cid", rsp.Cid))
	return rsp, nil
}

// FileAttachment .
func (d *dao) FileAttachment(ctx context.Context, req *userProto.FileAttachmentReq) (model.FileAttachmentItem, error) {
	requestID := d.GenerateID()
	rsp := model.FileAttachmentItem{}
	trace_id := util.GetTraceid(ctx)
	log.Logger.Info("FileAttachment start", log.String("trace_id", trace_id), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&userProto.StreamRsp{
		Cmd:       model.CMDFileAttachment,
		Token:     model.FileAttachmentToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   trace_id,
	}, model.STORAGE_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("FileAttachment add proxy request error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()))
		return rsp, err
	}

	timer := time.NewTimer(model.W3PTimeoutFileAttachment * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := storageProto.DownloadReply{}
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("FileAttachment response json unmarshal error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()))
			return rsp, err
		}
		log.Logger.Debug("FileAttachment storage response", log.String("trace_id", trace_id), log.Any("ret_code", ret.Code), log.Any("ret_msg", ret.Msg))

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg

		if rsp.Code > model.StatusSystemErrorCode {
			log.Logger.Error("FileAttachment response rsp error", log.String("trace_id", trace_id), log.Any("rsp", rsp))
			return rsp, errors.New(rsp.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("FileAttachment response not good", log.String("trace_id", trace_id), log.Any("rsp", rsp))
			return rsp, nil
		}

		rsp.Success = true
		rsp.Message = "success"
		rsp.Cid = ret.GetData().GetCid()
		rsp.Content = ret.GetData().GetContent()
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("FileAttachment response timeout", log.String("trace_id", trace_id))
		return rsp, fmt.Errorf("FileAttachment  response timeout")
	}

	log.Logger.Info("FileAttachment end", log.String("trace_id", trace_id), log.Any("attach_length", len(rsp.Content)), log.Any("cid", rsp.Cid))
	return rsp, nil
}

func (d *dao) FileReport(ctx context.Context, req *userProto.FileReportReq) (model.FileReportRsp, error) {
	requestID := d.GenerateID()
	trace_id := util.GetTraceid(ctx)
	log.Logger.Debug("FileReport start", log.String("trace_id", trace_id), log.String("params", req.GetParams()))

	rsp := model.FileReportRsp{}

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&userProto.StreamRsp{
		Cmd:       model.CMDFileReport,
		Token:     model.FileReportToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   trace_id,
	}, model.STORAGE_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("FileReport  add proxy request error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := storageProto.ReportReply{}
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("FileReport response json unmarshal error", log.String("trace_id", trace_id), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		rsp.Code = model.StatusOK
		rsp.Msg = model.MsgOK

		if rsp.Code > model.StatusSystemErrorCode {
			log.Logger.Error("FileAttachment response rsp error", log.String("trace_id", trace_id), log.Any("ret", ret))
			return rsp, errors.New(rsp.Msg)
		}

		if !ret.Success {
			log.Logger.Error("FileReport response error status", log.String("trace_id", trace_id), log.Any("ret", ret))
			return rsp, errors.New("FileReport rpc error")
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("FileReport response timeout", log.String("trace_id", trace_id), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("FileReport response timeout")
	}

	log.Logger.Debug("FileReport end", log.String("trace_id", trace_id))
	return rsp, nil
}
