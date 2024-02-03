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
	"github.com/web3password/satis/util"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/web3password/satis/log"
	"github.com/web3password/satis/model"
	pb "github.com/web3password/w3p-protobuf/user"
)

// GenerateID .
func (d *dao) GenerateID() int64 {
	requestId := time.Now().UnixNano()
	for i := 0; i < 2; i++ {
		rand.Seed(time.Now().UnixNano())
		requestId = requestId + rand.Int63n(time.Now().UnixNano())
	}

	return requestId
	//return d.snow.Generate().Int64()
}

// RegisterUser .
func (d *dao) RegisterUser(ctx context.Context, signature, params string) (*pb.RegisterRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.RegisterRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("RegisterUser start", log.String("trace_id", traceId), log.String("params", params))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDRegister,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: signature,
		Params:    params,
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("RegisterUser add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, err
	}

	timer := time.NewTimer(model.W3PTimeoutMax * time.Second)

	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.ResponseBytes{}
		log.Logger.Debug("RegisterUser response", log.String("trace_id", traceId))
		if err := bson.Unmarshal(res.GetData(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("RegisterUser bson unmarshal error", log.String("trace_id", traceId))
			return rsp, err
		}
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("RegisterUser response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("RegisterUser response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("RegisterUser timeout", log.String("trace_id", traceId))
		return rsp, fmt.Errorf("RegisterUser timeout")
	}

	log.Logger.Info("RegisterUser end", log.String("trace_id", traceId))
	return rsp, nil
}

func (d *dao) GetPersonalSignAddress(ctx context.Context, signature, params string) (model.PersonalSignListRsp, error) {
	requestID := d.GenerateID()
	rsp := model.PersonalSignListRsp{}
	rsp.List = make([]*model.PersonalSign, 0)
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("GetPersonalSignAddress start", log.String("trace_id", traceId), log.String("params", params))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDGetPersonalSignAddress,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: signature,
		Params:    params,
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("GetPersonalSignAddress add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, err
	}

	timer := time.NewTimer(model.W3PTimeoutMax * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.Response{}
		log.Logger.Debug("GetPersonalSignAddress response data", log.String("trace_id", traceId), log.String("res", res.GetParams()))
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("GetPersonalSignAddress json unmarshal fail err:%+v", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("GetPersonalSignAddress response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("GetPersonalSignAddress response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		if err := jsoniter.UnmarshalFromString(ret.Data, &rsp.List); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("GetPersonalSignAddress response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()))
			return rsp, err
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("GetPersonalSignAddress response timeout", log.String("trace_id", traceId))
		return rsp, fmt.Errorf("GetPersonalSignAddress timeout")
	}
	log.Logger.Info("GetPersonalSignAddress end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

// GetVIPInfo .
func (d *dao) GetVIPInfo(ctx context.Context, signature, params string) (model.VIPInfoRsp, error) {
	requestID := d.GenerateID()
	rsp := model.VIPInfoRsp{}
	traceId := util.GetTraceid(ctx)

	log.Logger.Info("GetVIPInfo start", log.String("trace_id", traceId), log.String("params", params))
	pbreq := &pb.StreamRsp{
		Cmd:       model.CMDGetVipInfo,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: signature,
		Params:    params,
		TraceId:   traceId,
	}

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(pbreq, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("GetVIPInfo add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, err
	}

	timer := time.NewTimer(model.W3PTimeoutMax * time.Second)

	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.ResponseBytes{}
		log.Logger.Debug("GetVIPInfo response", log.String("trace_id", traceId), log.Any("response", res.GetParams()))
		if err := bson.Unmarshal(res.GetData(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("GetVIPInfo bson unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()))
			return rsp, err
		}

		log.Logger.Debug("GetVIPInfo request result", log.String("trace_id", traceId), log.Any("ret.code", ret.Code))

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("GetVIPInfo response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("GetVIPInfo response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		if err := bson.Unmarshal(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("GetVIPInfo bson json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()))
			return rsp, err
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("GetVIPInfo timeout", log.String("trace_id", traceId))
		return rsp, fmt.Errorf("GetVIPInfo timeout")
		/*case <-time.After(3 * time.Second):
		d.addStreamRequest(pbreq, model.ARES_PROXY)
		log.Logger.Error("GetVIPInfo retry", log.String("trace_id", traceId), log.String("params", params))*/
	}

	log.Logger.Info("GetVIPInfo end", log.String("trace_id", traceId), log.Any("rsp.code", rsp.Code))
	//log.Logger.Debug("GetVIPInfo end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (d *dao) GetUserInfo(ctx context.Context, signature, params string) (model.UserInfoRsp, error) {
	requestID := d.GenerateID()
	rsp := model.UserInfoRsp{}
	traceId := util.GetTraceid(ctx)

	log.Logger.Debug("GetUserInfo start", log.String("trace_id", traceId), log.String("params", params))
	pbreq := &pb.StreamRsp{
		Cmd:       model.CMDGetUserInfo,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: signature,
		Params:    params,
		TraceId:   traceId,
	}

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(pbreq, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("GetUserInfo add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", params))
		return rsp, err
	}

	timer := time.NewTimer(model.W3PTimeoutMin * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.Response{}
		log.Logger.Debug("GetUserInfo response", log.String("trace_id", traceId), log.Any("response", res), log.String("params", params))
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("GetUserInfo json unmarshal fail", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", params))
			return rsp, err
		}
		log.Logger.Debug("GetUserInfo response data", log.Any("ret", ret), log.String("trace_id", traceId), log.Any("response", res.GetParams()))

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("GetUserInfo response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("GetUserInfo response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}
		//rsp.Data
		if err := jsoniter.UnmarshalFromString(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("GetUserInfo data json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", params))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		log.Logger.Debug("GetUserInfo response success", log.String("trace_id", traceId), log.Any("rsp", rsp), log.Any("response", res.GetParams()), log.String("params", params))
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("GetUserInfo timeout", log.String("trace_id", traceId), log.String("params", params))
		return rsp, fmt.Errorf("GetUserInfo timeout")
	}
	log.Logger.Debug("GetUserInfo end", log.String("trace_id", traceId))
	return rsp, nil
}

// Initialize
func (d *dao) Initialize(ctx context.Context, signature, params string) error {
	return nil
}

// GetVersionDesc
func (d *dao) GetVersionDesc(ctx context.Context, signature, params string) (model.VersionDescRsp, error) {
	traceId := util.GetTraceid(ctx)

	requestID := d.GenerateID()
	rsp := model.VersionDescRsp{}
	log.Logger.Debug("GetVersionDesc start", log.String("trace_id", traceId), log.String("params", params))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDGetVersionDesc,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: signature,
		Params:    params,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("GetVersionDesc add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", params))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &rsp); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("GetVersionDesc data json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", params))
			return rsp, err
		}

		if rsp.Code > model.StatusSystemErrorCode {
			log.Logger.Error("GetVersionDesc response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(rsp.Msg)
		}

	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("GetVersionDesc timeout", log.String("trace_id", traceId), log.String("params", params))
		rsp.VersionDesc = "mock"
	}
	log.Logger.Debug("GetVersionDesc end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (d *dao) StorageReport(ctx context.Context, signature, params string) (*pb.StorageReportRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.StorageReportRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("StorageReport start", log.String("trace_id", traceId), log.String("params", params))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDStorageReport,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: signature,
		Params:    params,
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("StorageReport request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", params))
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
			log.Logger.Error("StorageReport json unmarshal fail err:%+v", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("StorageReport response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("StorageReport response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("StorageReport timeout", log.String("trace_id", traceId), log.String("params", params))
		return rsp, fmt.Errorf("StorageReport response timeout")
	}
	log.Logger.Info("StorageReport end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (d *dao) StorageStat(ctx context.Context, signature, params string) (*pb.StorageStatRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.StorageStatRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Debug("StorageStat start", log.String("trace_id", traceId), log.String("params", params))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDStorageStat,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: signature,
		Params:    params,
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("StorageStat proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", params))
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
			log.Logger.Error("StorageStat json unmarshal fail", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("StorageStat response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("StorageStat response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		if err := jsoniter.UnmarshalFromString(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("StorageStat response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", params))
			return rsp, err
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("StorageStat timeout", log.String("trace_id", traceId), log.String("params", params))
		return rsp, fmt.Errorf("StorageStat response timeout")
	}
	log.Logger.Debug("StorageStat end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

// GetVersionConfig .
func (d *dao) GetVersionConfig(ctx context.Context, req *pb.GetVersionConfigReq) (*pb.GetVersionConfigRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.GetVersionConfigRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Debug("GetVersionConfig start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDGetVersionConfig,
		Token:     model.GetVersionConfigToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("GetVersionConfig proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.Response{}
		log.Logger.Debug("GetVersionConfig waitChan response ", log.Any("response", res))
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("GetVersionConfig response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("GetVersionConfig response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("GetVersionConfig response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		log.Logger.Debug("GetVersionConfig waitChan ret ", log.Any("ret", ret))

		if err := jsoniter.UnmarshalFromString(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("GetVersionConfig response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("GetVersionConfig timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("GetVersionConfig response timeout")
	}
	log.Logger.Debug("GetVersionConfig end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}
