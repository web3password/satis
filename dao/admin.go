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
	"github.com/web3password/satis/model"
	"github.com/web3password/satis/util"
	pb "github.com/web3password/w3p-protobuf/user"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// AdminRegister admin register
func (d *dao) AdminRegister(ctx context.Context, req *pb.AdminRegisterReq) (*pb.AdminRegisterRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.AdminRegisterRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Debug("AdminRegister start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDAdminRegister,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		log.Logger.Error("AdminRegister add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		ret := model.ResponseBytes{}
		if err := bson.Unmarshal(res.GetData(), &ret); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("AdminRegister response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("AdminRegister response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("AdminRegister response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		vipOrgRegister := model.VipRegister{}
		err := bson.Unmarshal(ret.Data, &vipOrgRegister)
		if err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("AdminRegister Unmarshal failed", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(model.MsgSystemErr)
		}
		rsp.Data = &pb.AdminRegisterRsp_Data{OrgId: vipOrgRegister.OrgId}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("AdminRegister timeout", log.String("trace_id", traceId), log.Any("rsp", rsp))
	}

	log.Logger.Debug("AdminRegister end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

// AdminTransferSuperAdmin .
func (d *dao) AdminTransferSuperAdmin(ctx context.Context, req *pb.AdminTransferSuperAdminReq) (*pb.AdminTransferSuperAdminRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.AdminTransferSuperAdminRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("AdminTransferSuperAdmin start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDAdminTransferSuperAdmin,
		Token:     model.AdminTransferSuperAdminToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		log.Logger.Error("AdminTransferSuperAdmin add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
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
			log.Logger.Error("AdminTransferSuperAdmin response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		log.Logger.Info("AdminTransferSuperAdmin response", log.String("trace_id", traceId), log.Any("ret", ret))

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("AdminTransferSuperAdmin response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("AdminTransferSuperAdmin response timeout", log.String("trace_id", traceId), log.Any("rsp", rsp))
		return rsp, fmt.Errorf("AdminTransferSuperAdmin response timeout")
	}

	log.Logger.Info("AdminTransferSuperAdmin end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

// AdminOperationHistory .
func (d *dao) AdminOperationHistory(ctx context.Context, req *pb.AdminOperationHistoryReq) (*pb.AdminOperationHistoryRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.AdminOperationHistoryRsp)
	rsp.Code = model.StatusOK
	traceId := util.GetTraceid(ctx)
	log.Logger.Debug("AdminOperationHistory start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDAdminOperationHistory,
		Token:     model.AdminOperationHistoryToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("AdminOperationHistory addStreamRequest error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
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
			log.Logger.Error("AdminOperationHistory response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		log.Logger.Debug("AdminOperationHistory ret", log.String("trace_id", traceId), log.Any("ret.code", ret.Code))

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("AdminOperationHistory response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("AdminOperationHistory response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		//areas val to satis
		if err := jsoniter.UnmarshalFromString(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("AdminOperationHistory json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		return rsp, nil
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("AdminOperationHistory timeout", log.String("trace_id", traceId), log.Any("rsp", rsp))
		return rsp, fmt.Errorf("admin operation history response timeout")
	}
	log.Logger.Debug("AdminOperationHistory end, list rsp", log.String("trace_id", traceId), log.Any("rsp.code", rsp.GetCode()), log.Any("len", len(rsp.GetData())))
	return rsp, nil
}

// AdminGetOrgInfo .
func (d *dao) AdminGetOrgInfo(ctx context.Context, req *pb.AdminGetOrgInfoReq) (*pb.AdminGetOrgInfoRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.AdminGetOrgInfoRsp)
	rsp.Code = model.StatusOK
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("AdminGetOrgInfo start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDAdminGetOrgInfo,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("AdminGetOrgInfo addStreamRequest error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
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
			log.Logger.Error("AdminGetOrgInfo json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("AdminGetOrgInfo response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("AdminGetOrgInfo response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		rsp.Data = ret.Data
		return rsp, nil
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("AdminGetOrgInfo timeout", log.String("trace_id", traceId))
		return rsp, fmt.Errorf("AdminGetOrgInfo response timeout")
	}
	log.Logger.Info("AdminGetOrgInfo end", log.String("trace_id", traceId), log.Any("rsp.code", rsp.GetCode()))
	return rsp, nil
}

// AdminUpdateOrgInfo .
func (d *dao) AdminUpdateOrgInfo(ctx context.Context, req *pb.AdminUpdateOrgInfoReq) (*pb.AdminUpdateOrgInfoRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.AdminUpdateOrgInfoRsp)
	rsp.Code = model.StatusOK
	traceId := util.GetTraceid(ctx)
	log.Logger.Debug("AdminUpdateOrgInfo start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDAdminUpdateOrgInfo,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		TraceId:   traceId,
		Data:      req.GetData(),
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("AdminUpdateOrgInfo addStreamRequest error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
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
			log.Logger.Error("AdminUpdateOrgInfo json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("AdminUpdateOrgInfo response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("AdminUpdateOrgInfo timeout", log.String("trace_id", traceId))
		return rsp, fmt.Errorf("AdminUpdateOrgInfo response timeout")
	}
	log.Logger.Debug("AdminUpdateOrgInfo end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

// AdminAuthorization .
func (d *dao) AdminAuthorization(ctx context.Context, signature, params string) (model.AdminAuthorizationRsp, error) {
	requestID := d.GenerateID()
	rsp := model.AdminAuthorizationRsp{}
	traceId := util.GetTraceid(ctx)
	log.Logger.Debug("AdminAuthorization start", log.String("trace_id", traceId), log.String("params", params))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDAdminAuthorization,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: signature,
		Params:    params,
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("AdminAuthorization addStreamRequest error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()
		if err := jsoniter.UnmarshalFromString(res.GetParams(), &rsp); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("AdminAuthorization json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
			return rsp, err
		}
		if rsp.Code > model.StatusSystemErrorCode {
			log.Logger.Error("AdminAuthorization response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(rsp.Msg)
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("AdminAuthorization timeout", log.String("trace_id", traceId), log.String("params", params))
		return rsp, fmt.Errorf("AdminAuthorization timeout")
	}
	log.Logger.Debug("AdminAuthorization end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

// AdminAddMember add admin member
func (d *dao) AdminAddMember(ctx context.Context, req *pb.AdminAddMemberReq) (model.AdminRsp, error) {
	requestID := d.GenerateID()
	//rsp := new(pb.AdminAddMemberRsp)
	rsp := model.AdminRsp{}
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("AdminAddMember start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDAdminAddMember,
		Token:     model.AdminAddMemberToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("AdminAddMember addStreamRequest error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
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
			log.Logger.Error("AdminAddMember response params json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		log.Logger.Info("AdminAddMember response ret", log.String("trace_id", traceId), log.Any("ret", ret))

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("AdminAddMember response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("AdminAddMember response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("admin add member response timeout")
	}

	log.Logger.Info("AdminAddMember end", log.String("trace_id", traceId), log.Any("rsp", rsp))

	return rsp, nil
}

// AdminGetMemberList .
func (d *dao) AdminGetMemberList(ctx context.Context, req *pb.AdminGetMemberListReq) (*pb.AdminGetMemberListRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.AdminGetMemberListRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Debug("AdminGetMemberList start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDAdminGetMemberList,
		Token:     model.AdminGetMemberListToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("AdminGetMemberList addStreamRequest error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
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
			log.Logger.Error("AdminGetMemberList response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("res.params", res.GetParams()))
			return rsp, err
		}
		log.Logger.Debug("AdminGetMemberList response ret", log.String("trace_id", traceId), log.Any("ret.code", ret.Code))
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("AdminGetMemberList response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("AdminGetMemberList response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		if err := jsoniter.UnmarshalFromString(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("AdminGetMemberList response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
			return rsp, err
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("AdminGetMemberList response timeout", log.String("trace_id", traceId))
		return rsp, fmt.Errorf("admin get member list response timeout")
	}

	log.Logger.Info("AdminGetMemberList end", log.String("trace_id", traceId), log.Any("rsp.code", rsp.GetCode()), log.Any("len-memberlist", len(rsp.GetData().GetList())))

	return rsp, nil
}

// AdminUpdateMember admin update member
func (d *dao) AdminUpdateMember(ctx context.Context, req *pb.AdminUpdateMemberReq) (model.AdminRsp, error) {
	requestID := d.GenerateID()
	//rsp := new(pb.AdminUpdateMemberRsp)
	rsp := model.AdminRsp{}
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("AdminUpdateMember start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDAdminUpdateMember,
		Token:     model.AdminUpdateMemberToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("AdminUpdateMember addStreamRequest error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
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
			log.Logger.Error("AdminUpdateMember response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		log.Logger.Info("AdminUpdateMember result", log.String("trace_id", traceId), log.Any("ret", ret))
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("AdminUpdateMember response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("AdminUpdateMember response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("AdminUpdateMember response timeout")
	}

	log.Logger.Info("AdminUpdateMember end", log.String("trace_id", traceId), log.Any("rsp", rsp))

	return rsp, nil
}

// AdminRemoveMember admin remove member
func (d *dao) AdminRemoveMember(ctx context.Context, req *pb.AdminRemoveMemberReq) (model.AdminRsp, error) {
	requestID := d.GenerateID()
	//rsp := new(pb.AdminRemoveMemberRsp)
	rsp := model.AdminRsp{}
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("AdminRemoveMember start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDAdminRemoveMember,
		Token:     model.AdminRemoveMemberToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("AdminRemoveMember addStreamRequest error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
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
			log.Logger.Error("admin remove member response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		log.Logger.Info("AdminRemoveMember response", log.String("trace_id", traceId), log.Any("ret", ret))
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("AdminRemoveMember response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("AdminRemoveMember response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("AdminRemoveMember response timeout")
	}

	log.Logger.Info("AdminRemoveMember end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

// GetAdminMnemonic  get admin mnemonic
func (d *dao) GetAdminMnemonic(ctx context.Context, req *pb.GetAdminMnemonicReq) (*pb.GetAdminMnemonicRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.GetAdminMnemonicRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Debug("GetAdminMnemonic start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDAdminGetAdminMnemonic,
		Token:     model.AdminGetAdminMnemonicToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("GetAdminMnemonic addStreamRequest error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
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
			log.Logger.Error("GetAdminMnemonic response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		log.Logger.Debug("GetAdminMnemonic response ret", log.String("trace_id", traceId), log.Any("ret", ret))
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("GetAdminMnemonic response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("GetAdminMnemonic response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		log.Logger.Debug("GetAdminMnemonic response ret.Data", log.String("trace_id", traceId), log.Any("response", ret.Data))
		if err := jsoniter.UnmarshalFromString(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("GetAdminMnemonic response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("GetAdminMnemonic response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("admin get admin mnemonic response timeout")
	}

	log.Logger.Debug("GetAdminMnemonic end",
		log.String("trace_id", traceId),
		log.Any("rsp.code", rsp.GetCode()),
		log.Any("rsp.msg", rsp.GetMsg()),
	)
	return rsp, nil
}

// AdminBatchImportMember batch import members
func (d *dao) AdminBatchImportMember(ctx context.Context, req *pb.AdminBatchImportMemberReq) (*pb.AdminBatchImportMemberRsp, error) {
	requestID := d.GenerateID()
	rsp := new(pb.AdminBatchImportMemberRsp)
	traceId := util.GetTraceid(ctx)
	log.Logger.Debug("AdminBatchImportMember start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDAdminBatchImportMember,
		Token:     model.AdminBatchImportMemberToken,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("AdminBatchImportMember addStreamRequest error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
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
			log.Logger.Error("AdminBatchImportMember response json unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		log.Logger.Debug("AdminBatchImportMember response ret", log.String("trace_id", traceId), log.Any("ret", ret))
		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("AdminBatchImportMember response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("AdminBatchImportMember response timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("AdminBatchImportMember response timeout")
	}

	log.Logger.Debug("AdminBatchImportMember end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}
