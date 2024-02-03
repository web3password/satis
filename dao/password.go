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
	jsoniter "github.com/json-iterator/go"
	"github.com/web3password/satis/log"
	"github.com/web3password/satis/util"
	pb "github.com/web3password/w3p-protobuf/user"
	pbindex "github.com/web3password/w3p-protobuf/user_data_index"
	"gopkg.in/mgo.v2/bson"
	"time"

	"github.com/web3password/satis/model"
)

func (d *dao) CheckTx(ctx context.Context, req *pb.CheckTxReq) (model.CheckTxRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := model.CheckTxRsp{}

	log.Logger.Info("CheckTx start", log.String("trace_id", traceId), log.Int64("requestID", requestID), log.Any("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDIndexCheckTx,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("CheckTx add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()

		var ret pbindex.CheckTxResponse
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("CheckTx response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("CheckTx response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("CheckTx response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		rsp.Height = ret.GetData().Height
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("CheckTx response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("CheckTx response timeout")
	}

	log.Logger.Info("CheckTx end", log.String("trace_id", traceId), log.Int64("requestID", requestID), log.Any("rsp", rsp))
	return rsp, nil
}

func (d *dao) BatchCheckTx(ctx context.Context, req *pb.BatchCheckTxReq) (*pb.BatchCheckTxRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := &pb.BatchCheckTxRsp{
		Data: &pb.BatchCheckTxRsp_Data{},
	}

	log.Logger.Info("BatchCheckTx start", log.String("trace_id", traceId), log.Int64("requestID", requestID), log.Any("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDIndexBatchCheckTx,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("BatchCheckTx add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(model.W3PTimeoutMax * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()

		var ret pbindex.BatchCheckTxResponse
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("BatchCheckTx response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("BatchCheckTx response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("BatchCheckTx response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		list := make([]*pb.BatchCheckTxRsp_Data_ListItem, 0, len(ret.GetData().List))
		for _, item := range ret.GetData().List {
			tmpData := &pb.BatchCheckTxRsp_Data_ListItem{
				Hash:    item.Hash,
				Success: item.Success,
			}
			list = append(list, tmpData)
		}
		rsp.Data.List = list

		log.Logger.Info("BatchCheckTx response success", log.String("trace_id", traceId), log.Any("rsp", rsp), log.String("params", req.GetParams()))
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("BatchCheckTx handleResponse-Timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("BatchCheckTx timeout")
	}

	log.Logger.Info("BatchCheckTx end", log.String("trace_id", traceId), log.Int64("requestID", requestID), log.Any("rsp", rsp))
	return rsp, nil
}

// AddOrDelCredential .
func (d *dao) AddOrDelCredential(ctx context.Context, signature, params string, data []byte) (model.AddOrDelCredentialRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := model.AddOrDelCredentialRsp{}

	log.Logger.Info("AddOrDelCredential start", log.String("trace_id", traceId), log.Int64("requestID", requestID), log.Any("params", params))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDIndexAddOrDelCredential,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: signature,
		Params:    params,
		Data:      data,
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("AddOrDelCredential add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()

		var ret pbindex.AddOrDelCredentialResponse
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("AddOrDelCredential response json unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()), log.String("params", params))
			return rsp, err
		}

		log.Logger.Info("AddOrDelCredential success ", log.Any("ret", ret), log.String("trace_id", traceId))
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("AddOrDelCredential response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("AddOrDelCredential response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		rsp.TxHash = ret.GetData().GetTxhash()
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("AddOrDelCredential-handleResponse-Timeout", log.String("trace_id", traceId))
		return rsp, fmt.Errorf("AddOrDelCredential timeout")
	}

	log.Logger.Info("AddOrDelCredential end", log.String("trace_id", traceId), log.Int64("requestID", requestID))
	log.Logger.Debug("AddOrDelCredential end debug", log.String("trace_id", traceId), log.Int64("requestID", requestID), log.Any("rsp", rsp))
	return rsp, nil
}

func (d *dao) BatchAddCredential(ctx context.Context, req *pb.BatchAddCredentialReq) (*pb.BatchAddCredentialRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := &pb.BatchAddCredentialRsp{
		Data: &pb.BatchAddCredentialRsp_Data{},
	}

	log.Logger.Info("BatchAddCredential start", log.String("trace_id", traceId), log.Int64("requestID", requestID), log.Any("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDIndexBatchAddCredential,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("BatchAddCredential add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, err
	}

	timer := time.NewTimer(model.W3PTimeoutMax * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()

		var ret pbindex.BatchAddCredentialResponse
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("BatchAddCredential response json unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()))
			return rsp, err
		}

		log.Logger.Info("BatchAddCredential response success ", log.String("trace_id", traceId))
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("BatchAddCredential response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("BatchAddCredential response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		list := make([]*pb.BatchAddCredentialRsp_Data_ListItem, 0, len(ret.GetData().List))
		for _, item := range ret.GetData().List {
			tmpData := &pb.BatchAddCredentialRsp_Data_ListItem{
				Id:   item.Id,
				Hash: item.Hash,
			}
			list = append(list, tmpData)
		}
		rsp.Data.List = list
		log.Logger.Info("BatchAddCredential response ret", log.String("trace_id", traceId))
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("BatchAddCredential-handleResponse-Timeout", log.String("trace_id", traceId))
		return rsp, fmt.Errorf("BatchAddCredential timeout")
	}

	log.Logger.Info("BatchAddCredential end", log.String("trace_id", traceId), log.Int64("requestID", requestID))
	return rsp, nil
}

func (d *dao) BatchDeleteCredential(ctx context.Context, req *pb.BatchDeleteCredentialReq) (*pb.BatchDeleteCredentialRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := &pb.BatchDeleteCredentialRsp{
		Data: &pb.BatchDeleteCredentialRsp_Data{},
	}

	log.Logger.Info("BatchDeleteCredential start", log.String("trace_id", traceId), log.Int64("requestID", requestID), log.Any("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDIndexBatchDeleteCredential,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("BatchDeleteCredential add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(model.W3PTimeoutMax * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		var ret pbindex.BatchDeleteCredentialResponse
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("BatchDeleteCredential response json unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		log.Logger.Info("BatchDeleteCredential response success ", log.String("trace_id", traceId))
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("BatchDeleteCredential response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("BatchDeleteCredential response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		list := make([]*pb.BatchDeleteCredentialRsp_Data_ListItem, 0, len(ret.GetData().List))
		for _, item := range ret.GetData().List {
			tmpData := &pb.BatchDeleteCredentialRsp_Data_ListItem{
				Id:   item.Id,
				Hash: item.Hash,
			}
			list = append(list, tmpData)
		}
		rsp.Data.List = list
		log.Logger.Info("BatchDeleteCredential response success", log.String("trace_id", traceId))
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("BatchDeleteCredential-handleResponse-Timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("BatchDeleteCredential timeout")
	}

	log.Logger.Info("BatchDeleteCredential end", log.String("trace_id", traceId), log.Int64("requestID", requestID))
	return rsp, nil
}

// GetPrimaryAddrIndexDetail .
func (d *dao) GetPrimaryAddrIndexDetail(ctx context.Context, signature, params string) (model.GetCredentialRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := model.GetCredentialRsp{}

	log.Logger.Info("GetPrimaryAddrIndexDetail start", log.String("trace_id", traceId), log.Int64("requestID", requestID))
	log.Logger.Debug("GetPrimaryAddrIndexDetail start debug", log.String("trace_id", traceId), log.Int64("requestID", requestID), log.Any("params", params))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDIndexGetPrimaryAddrIndexDetail,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: signature,
		Params:    params,
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("GetPrimaryAddrIndexDetail add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", params))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()

		var ret pbindex.GetPrimaryAddrIndexDetailResponse
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("GetPrimaryAddrIndexDetail response json unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()), log.String("params", params))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("GetPrimaryAddrIndexDetail response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("GetPrimaryAddrIndexDetail response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		rsp.Id = ret.GetData().Id
		rsp.OpTimestamp = int32(ret.GetData().OpTimestamp)
		rsp.Credential = ret.GetData().Credential
		log.Logger.Debug("GetPrimaryAddrIndexDetail response success", log.String("trace_id", traceId))
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("GetPrimaryAddrIndexDetail-handleResponse-Timeout", log.String("trace_id", traceId), log.String("params", params))
		return rsp, fmt.Errorf("GetPrimaryAddrIndexDetail timeout")
	}

	log.Logger.Info("GetPrimaryAddrIndexDetail end", log.String("trace_id", traceId), log.Int64("requestID", requestID))
	return rsp, nil
}

// DelPrimaryAddrIndex .
func (d *dao) DeleteAllCredential(ctx context.Context, signature, params string) (model.AddOrDelCredentialRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := model.AddOrDelCredentialRsp{}

	log.Logger.Info("DeleteAllCredential start", log.String("trace_id", traceId), log.Int64("requestID", requestID), log.Any("params", params))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	/*if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDIndexDelPrimaryAddrIndex,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: signature,
		Params:    params,
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("DeleteAllCredential add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", params))
		return rsp, err
	}*/

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDIndexDelPrimaryAddrIndex,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: signature,
		Params:    params,
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("DeleteAllCredential add proxy1 request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", params))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.ResponseBytes{}
		log.Logger.Debug("DeleteAllCredential response", log.String("trace_id", traceId), log.Any("response", res), log.String("params", params))
		if err := bson.Unmarshal(res.GetData(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("DeleteAllCredential response unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()), log.String("params", params))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("DeleteAllCredential response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("DeleteAllCredential response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}
		log.Logger.Info("DeleteAllCredential response success", log.Any("rsp", rsp), log.String("trace_id", traceId), log.Any("response", res.GetParams()), log.String("params", params))
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("dao-DeleteAllCredential-handleResponse-Timeout", log.String("trace_id", traceId), log.String("params", params))
		return rsp, fmt.Errorf("DeleteAllCredential timeout")
	}

	log.Logger.Info("DeleteAllCredential end", log.String("trace_id", traceId), log.Int64("requestID", requestID))
	return rsp, nil
}

func (d *dao) GetAllCredentialTimestamp(ctx context.Context, req *pb.GetAllCredentialTimestampReq) (model.GetAllCredentialTimestampListRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := model.GetAllCredentialTimestampListRsp{}
	rsp.List = make([]*model.GetAllCredentialTimestampRsp, 0)

	log.Logger.Info("GetAllCredentialTimestamp start", log.String("trace_id", traceId), log.Int64("requestID", requestID), log.Any("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDIndexGetAllCredentialTimestamp,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("GetAllCredentialTimestamp add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(model.W3PTimeoutMax * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()

		var ret pbindex.GetAllCredentialTimestampResponse
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("GetAllCredentialTimestamp response json unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("GetAllCredentialTimestamp response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("GetAllCredentialTimestamp response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		for _, item := range ret.GetData() {
			tmpData := &model.GetAllCredentialTimestampRsp{
				Id:          item.Id,
				OpTimestamp: int32(item.OpTimestamp),
			}

			rsp.List = append(rsp.List, tmpData)
		}
		log.Logger.Info("GetAllCredentialTimestamp response success", log.String("trace_id", traceId))
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("dao-GetAllCredentialTimestamp-handleResponse-Timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("GetAllCredentialTimestamp timeout")
	}

	log.Logger.Info("GetAllCredentialTimestamp end", log.String("trace_id", traceId), log.Int64("requestID", requestID))
	return rsp, nil
}

// GetPrimaryAddrIndexList .
func (d *dao) GetPrimaryAddrIndexList(ctx context.Context, req *pb.GetCredentialListReq) (model.GetCredentialListRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := model.GetCredentialListRsp{}
	rsp.List = make([]*model.GetCredentialRsp, 0)

	log.Logger.Info("GetPrimaryAddrIndexList start", log.String("trace_id", traceId), log.Int64("requestID", requestID), log.Any("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDIndexGetPrimaryAddrIndexList,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("GetPrimaryAddrIndexList add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(model.W3PTimeoutMax * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()

		var ret pbindex.GetPrimaryAddrIndexListResponse
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("GetPrimaryAddrIndexList response json unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("GetPrimaryAddrIndexList response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("GetPrimaryAddrIndexList response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		for _, item := range ret.GetData() {
			tmpData := &model.GetCredentialRsp{
				Id:          item.Id,
				Credential:  item.Credential,
				OpTimestamp: int32(item.OpTimestamp),
			}
			rsp.List = append(rsp.List, tmpData)
		}
		log.Logger.Info("GetPrimaryAddrIndexList response success", log.String("trace_id", traceId), log.Int64("requestID", requestID))
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("GetPrimaryAddrIndexList response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("GetPrimaryAddrIndexList timeout")
	}

	log.Logger.Info("GetPrimaryAddrIndexList end", log.String("trace_id", traceId), log.Int64("requestID", requestID))

	return rsp, nil
}
