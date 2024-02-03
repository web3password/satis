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
	pb "github.com/web3password/w3p-protobuf/user"
)

// ShareFolderCreate .
func (d *dao) ShareFolderCreate(ctx context.Context, req *pb.ShareFolderCreateReq) (model.ShareFolderRsp, error) {
	requestID := d.GenerateID()
	rsp := model.ShareFolderRsp{}
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("ShareFolderCreate start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDShareFolderCreate,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("ShareFolderCreate add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.Response{}
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderCreate response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		log.Logger.Info("ShareFolderCreate response rsp", log.String("trace_id", traceId), log.Any("rsp", ret))
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("ShareFolderCreate response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("ShareFolderCreate response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("ShareFolderCreate response timeout")
	}
	log.Logger.Info("ShareFolderCreate end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

// ShareFolderDestroy .
func (d *dao) ShareFolderDestroy(ctx context.Context, req *pb.ShareFolderDestroyReq) (*pb.ShareFolderDestroyRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.ShareFolderDestroyRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("ShareFolderDestroy start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDShareFolderDestroy,
		Token:     model.ShareFolderDestroyToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("ShareFolderDestroy add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.Response{}
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderDestroy response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		log.Logger.Info("ShareFolderDestroy response rsp", log.String("trace_id", traceId), log.Any("rsp", ret))
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("ShareFolderDestroy response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("ShareFolderDestroy response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("ShareFolderDestroy response timeout")
	}

	log.Logger.Info("ShareFolderDestroy end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

// ShareFolderAddMember .
func (d *dao) ShareFolderAddMember(ctx context.Context, req *pb.ShareFolderAddMemberReq) (*pb.ShareFolderAddMemberRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.ShareFolderAddMemberRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("ShareFolderAddMember start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDShareFolderAddMember,
		Token:     model.ShareFolderAddMemberToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("ShareFolderAddMember add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}
	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.Response{}
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderAddMember response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		log.Logger.Info("ShareFolderAddMember response rsp", log.String("trace_id", traceId), log.Any("rsp", ret))
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("ShareFolderAddMember response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("ShareFolderAddMember response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("ShareFolderAddMember response timeout")
	}

	log.Logger.Info("ShareFolderAddMember end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

// ShareFolderUpdateMember .
func (d *dao) ShareFolderUpdateMember(ctx context.Context, req *pb.ShareFolderUpdateMemberReq) (*pb.ShareFolderUpdateMemberRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.ShareFolderUpdateMemberRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("ShareFolderUpdateMember start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDShareFolderUpdateMember,
		Token:     model.ShareFolderUpdateMemberToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("ShareFolderUpdateMember add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}
	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.Response{}
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderUpdateMember response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		log.Logger.Info("ShareFolderUpdateMember response rsp", log.String("trace_id", traceId), log.Any("rsp", ret))
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("ShareFolderUpdateMember response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("ShareFolderUpdateMember response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("ShareFolderUpdateMember response timeout")
	}
	log.Logger.Info("ShareFolderUpdateMember end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

// ShareFolderUpdate .
func (d *dao) ShareFolderUpdate(ctx context.Context, req *pb.ShareFolderUpdateReq) (model.ShareFolderRsp, error) {
	requestID := d.GenerateID()
	rsp := model.ShareFolderRsp{}
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("ShareFolderUpdate start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDShareFolderUpdate,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("ShareFolderUpdate add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.Response{}
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderUpdate response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		log.Logger.Info("ShareFolderUpdate response rsp", log.String("trace_id", traceId), log.Any("rsp", ret))
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("ShareFolderUpdate response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("ShareFolderUpdate response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("ShareFolderUpdate response timeout")
	}
	log.Logger.Info("ShareFolderUpdate end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

// ShareFolderFolderList .
func (d *dao) ShareFolderFolderList(ctx context.Context, req *pb.ShareFolderFolderListReq) (*pb.ShareFolderFolderListRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.ShareFolderFolderListRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Debug("ShareFolderFolderList start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDShareFolderFolderList,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("ShareFolderFolderList add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}
	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.Response{}
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderFolderList response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("ShareFolderFolderList response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("ShareFolderFolderList response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		if err := jsoniter.UnmarshalFromString(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderFolderList response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("ShareFolderFolderList response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("ShareFolderFolderList response timeout")
	}
	//rspBson, _ := bson.Marshal(rsp)
	//log.Logger.Debug("ShareFolderFolderList end", log.String("trace_id", traceId), log.Any("rsp", base64.StdEncoding.EncodeToString(rspBson)))
	log.Logger.Debug("ShareFolderFolderList end", log.String("trace_id", traceId))
	return rsp, nil
}

// ShareFolderRecordList .
func (d *dao) ShareFolderRecordList(ctx context.Context, req *pb.ShareFolderRecordListReq) (*pb.ShareFolderRecordListRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.ShareFolderRecordListRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Debug("ShareFolderRecordList start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDShareFolderRecordList,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("ShareFolderRecordList add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}
	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.Response{}
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderRecordList response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("ShareFolderFolderList response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("ShareFolderRecordList response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		if err := jsoniter.UnmarshalFromString(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderRecordList response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("ShareFolderRecordList response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("ShareFolderRecordList response timeout")
	}
	log.Logger.Debug("ShareFolderFolderList end", log.String("trace_id", traceId))
	//rspBson, _ := bson.Marshal(rsp)
	//log.Logger.Debug("ShareFolderFolderList end", log.String("trace_id", traceId), log.Any("rsp", base64.StdEncoding.EncodeToString(rspBson)))
	return rsp, nil
}

func (d *dao) ShareFolderRecordListByRid(ctx context.Context, req *pb.ShareFolderRecordListByRidReq) (*pb.ShareFolderRecordListByRidRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.ShareFolderRecordListByRidRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Debug("ShareFolderRecordListByRid start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDShareFolderRecordListByRid,
		Token:     model.ShareFolderRecordListTokenByRid,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("ShareFolderRecordListByRid add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.Response{}
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderRecordListByRid response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("ShareFolderRecordListByRid response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("ShareFolderRecordListByRid response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		if err := jsoniter.UnmarshalFromString(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderRecordListByRid response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("ShareFolderRecordListByRid response timeout", log.String("trace_id", traceId))
		return rsp, fmt.Errorf("ShareFolderRecordListByRid response timeout")
	}
	rsp.Code = 111111
	rsp.Msg = "success"

	//rspBson, _ := bson.Marshal(rsp)
	log.Logger.Debug("ShareFolderRecordListByRid end", log.String("trace_id", traceId))
	return rsp, nil
}

// ShareFolderAddRecord .
func (d *dao) ShareFolderAddRecord(ctx context.Context, req *pb.ShareFolderAddRecordReq) (*pb.ShareFolderAddRecordRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.ShareFolderAddRecordRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("ShareFolderAddRecord start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDShareFolderAddRecord,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("ShareFolderAddRecord add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.Response{}
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderAddRecord response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("ShareFolderAddRecord response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("ShareFolderAddRecord response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("ShareFolderAddRecord response timeout")
	}
	log.Logger.Info("ShareFolderAddRecord end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

// ShareFolderDeleteRecord .
func (d *dao) ShareFolderDeleteRecord(ctx context.Context, req *pb.ShareFolderDeleteRecordReq) (*pb.ShareFolderDeleteRecordRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.ShareFolderDeleteRecordRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Debug("ShareFolderDeleteRecord start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDShareFolderDeleteRecord,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("ShareFolderDeleteRecord add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.Response{}
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderDeleteRecord response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("ShareFolderDeleteRecord response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("ShareFolderDeleteRecord response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("ShareFolderDeleteRecord response timeout")
	}
	log.Logger.Info("ShareFolderDeleteRecord end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

// ShareFolderMemberList .
func (d *dao) ShareFolderMemberList(ctx context.Context, req *pb.ShareFolderMemberListReq) (*pb.ShareFolderMemberListRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.ShareFolderMemberListRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Debug("ShareFolderMemberList start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDShareFolderMemberList,
		Token:     model.ShareFolderMemberListToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("ShareFolderMemberList add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.Response{}
		log.Logger.Debug("ShareFolderMemberList waitChan response ", log.String("trace_id", traceId), log.Any("response", res))
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderMemberList response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("ShareFolderMemberList response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("ShareFolderMemberList response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		log.Logger.Debug("ShareFolderMemberList waitChan ret ", log.String("trace_id", traceId), log.Any("ret", ret))

		if err := jsoniter.UnmarshalFromString(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderMemberList response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("ShareFolderMemberList response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("ShareFolderMemberList response timeout")
	}
	//rspBson, _ := bson.Marshal(rsp)
	//log.Logger.Debug("ShareFolderMemberList end", log.String("trace_id", traceId), log.Any("rsp", base64.StdEncoding.EncodeToString(rspBson)))
	log.Logger.Debug("ShareFolderMemberList end", log.String("trace_id", traceId))
	return rsp, nil
}

// ShareFolderDeleteMember .
func (d *dao) ShareFolderDeleteMember(ctx context.Context, req *pb.ShareFolderDeleteMemberReq) (*pb.ShareFolderDeleteMemberRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.ShareFolderDeleteMemberRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("ShareFolderDeleteMember start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDShareFolderDeleteMember,
		Token:     model.ShareFolderDeleteMemberToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("ShareFolderDeleteMember add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.Response{}
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderDeleteMember response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		log.Logger.Info("ShareFolderDeleteMember response ret", log.String("trace_id", traceId), log.Any("ret", ret))

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("ShareFolderDeleteMember response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("ShareFolderDeleteMember response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("ShareFolderDeleteMember response timeout")
	}

	log.Logger.Debug("ShareFolderDeleteMember end", log.String("trace_id", traceId), log.Any("rsp", rsp))

	return rsp, nil
}

// ShareFolderMemberExit .
func (d *dao) ShareFolderMemberExit(ctx context.Context, req *pb.ShareFolderMemberExitReq) (*pb.ShareFolderMemberExitRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.ShareFolderMemberExitRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("ShareFolderMemberExit start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDShareFolderMemberExit,
		Token:     model.ShareFolderMemberExitToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("ShareFolderMemberExit add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.Response{}
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderMemberExit response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		log.Logger.Info("ShareFolderMemberExit response ret", log.String("trace_id", traceId), log.Any("ret", ret))

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("ShareFolderMemberExit response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("ShareFolderMemberExit response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("ShareFolderMemberExit response timeout")
	}

	log.Logger.Info("ShareFolderMemberExit end", log.String("trace_id", traceId), log.Any("rsp", rsp))

	return rsp, nil
}

// ShareFolderBatchUpdate .
func (d *dao) ShareFolderBatchUpdate(ctx context.Context, req *pb.ShareFolderBatchUpdateReq) (*pb.ShareFolderBatchUpdateRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.ShareFolderBatchUpdateRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("ShareFolderBatchUpdate start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDShareFolderBatchUpdate,
		Token:     model.ShareFolderBatchUpdateToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.INDEX_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("ShareFolderBatchUpdate add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.Response{}
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("ShareFolderBatchUpdate response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		log.Logger.Info("ShareFolderBatchUpdate response ret", log.String("trace_id", traceId), log.Any("ret", ret))
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("ShareFolderMemberExit response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("ShareFolderBatchUpdate response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("ShareFolderBatchUpdate response timeout")
	}

	log.Logger.Info("ShareFolderBatchUpdate end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}
