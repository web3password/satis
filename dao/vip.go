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
	"github.com/web3password/satis/log"
	"github.com/web3password/satis/model"
	"github.com/web3password/satis/util"
	pb "github.com/web3password/w3p-protobuf/user"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (d *dao) VipGetConfig(ctx context.Context, req *pb.VipGetConfigReq) (*pb.VipGetConfigRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := new(pb.VipGetConfigRsp)
	log.Logger.Debug("VipGetConfig start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDVipGetConfig,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("VipGetConfig add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
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
			log.Logger.Error("VipGetConfig response bson unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("VipGetConfig response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("VipGetConfig response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		if err := bson.Unmarshal(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("VipGetConfig response bson unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		log.Logger.Debug("VipGetConfig response success", log.String("trace_id", traceId))
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("VipGetConfig get vip subscription list timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("VipGetConfig timeout")
	}

	log.Logger.Debug("VipGetConfig end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (d *dao) VipSubscriptionList(ctx context.Context, req *pb.VipSubscriptionListReq) (*pb.VipSubscriptionListRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := new(pb.VipSubscriptionListRsp)
	log.Logger.Info("VipSubscriptionList start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDVipSubscriptionList,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("VipSubscriptionList add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
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
			log.Logger.Error("VipSubscriptionList response bson unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("VipSubscriptionList response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("VipSubscriptionList response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		if err := bson.Unmarshal(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("VipSubscriptionList response bson unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()))
			return rsp, err
		}
		log.Logger.Debug("VipSubscriptionList response success", log.String("trace_id", traceId), log.Any("response", res.GetParams()))
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("VipSubscriptionList get vip subscription list timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("VipSubscriptionList timeout")
	}
	log.Logger.Info("VipSubscriptionList end", log.String("trace_id", traceId))
	//log.Logger.Debug("VipSubscriptionList end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (d *dao) VipPaymentList(ctx context.Context, req *pb.VipPaymentListReq) (*pb.VipPaymentListRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := new(pb.VipPaymentListRsp)
	log.Logger.Info("VipPaymentList start", log.String("trace_id", traceId))
	log.Logger.Debug("VipPaymentList start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDVipPaymentList,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("VipPaymentList add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
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
			log.Logger.Error("VipPaymentList response bson unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("VipPaymentList response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("VipPaymentList response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		if err := bson.Unmarshal(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("VipPaymentList response bson unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		log.Logger.Debug("VipPaymentList response success", log.Any("rsp", rsp), log.String("trace_id", traceId), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("VipPaymentList get vip payment list list timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("VipPaymentList timeout")
	}

	log.Logger.Info("VipPaymentList end", log.String("trace_id", traceId))
	log.Logger.Debug("VipPaymentList end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (d *dao) VipCreateOrder(ctx context.Context, req *pb.VipCreateOrderReq) (*pb.VipCreateOrderRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := new(pb.VipCreateOrderRsp)
	//log.Logger.Info("VipCreateOrder start", log.String("trace_id", traceId))
	log.Logger.Info("VipCreateOrder start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDVipCreateOrder,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("VipCreateOrder add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
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
			log.Logger.Error("VipCreateOrder response bson unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("VipCreateOrder response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("VipCreateOrder response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		if err := bson.Unmarshal(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("VipCreateOrder response bson unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()))
			return rsp, err
		}
		log.Logger.Info("VipCreateOrder response success", log.String("trace_id", traceId))
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("VipCreateOrder get vip create order timeout", log.String("trace_id", traceId))
		return rsp, fmt.Errorf("VipCreateOrder timeout")
	}

	//log.Logger.Info("VipCreateOrder end", log.String("trace_id", traceId))
	log.Logger.Info("VipCreateOrder end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (d *dao) VipCheckOrder(ctx context.Context, req *pb.VipCheckOrderReq) (*pb.VipCheckOrderRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := new(pb.VipCheckOrderRsp)

	log.Logger.Info("VipCheckOrder start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDVipCheckOrder,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("VipCheckOrder add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
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
			log.Logger.Error("VipCheckOrder response bson unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("VipCheckOrder response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("VipCheckOrder response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		if err := bson.Unmarshal(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("VipCheckOrder response bson unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()))
			return rsp, err
		}
		log.Logger.Info("VipCheckOrder response success", log.String("trace_id", traceId))
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("VipCheckOrder check order status timeout", log.String("trace_id", traceId))
		return rsp, fmt.Errorf("VipCheckOrder timeout")
	}

	//log.Logger.Info("VipCheckOrder end", log.String("trace_id", traceId))
	log.Logger.Info("VipCheckOrder end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (d *dao) VipAppleVerifyReceipt(ctx context.Context, req *pb.VipAppleVerifyReceiptReq) (*pb.VipAppleVerifyReceiptRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := new(pb.VipAppleVerifyReceiptRsp)

	log.Logger.Info("VipAppleVerifyReceipt start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDVipAppleVerifyReceipt,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("VipAppleVerifyReceipt add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
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
			log.Logger.Error("VipAppleVerifyReceipt response bson unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()), log.String("errMsg", err.Error()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("VipAppleVerifyReceipt response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("VipAppleVerifyReceipt response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		if err := bson.Unmarshal(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("VipAppleVerifyReceipt response bson unmarshal error", log.String("trace_id", traceId), log.String("errMsg", err.Error()), log.Any("response", res.GetParams()))
			return rsp, err
		}
		log.Logger.Debug("VipAppleVerifyReceipt response success", log.String("trace_id", traceId), log.Any("response", res.GetParams()))
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("VipAppleVerifyReceipt timeout", log.String("trace_id", traceId))
		return rsp, fmt.Errorf("VipAppleVerifyReceipt timeout")
	}

	log.Logger.Info("VipAppleVerifyReceipt end", log.String("trace_id", traceId))
	return rsp, nil
}

func (d *dao) GetDiscountCodeInfo(ctx context.Context, req *pb.GetDiscountCodeInfoReq) (*pb.GetDiscountCodeInfoRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := new(pb.GetDiscountCodeInfoRsp)

	log.Logger.Info("GetDiscountCodeInfo start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDGetDiscountCode,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("GetDiscountCodeInfo add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
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
			log.Logger.Error("GetDiscountCodeInfo response bson unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("GetDiscountCodeInfo response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("GetDiscountCodeInfo response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		if err := bson.Unmarshal(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("GetDiscountCodeInfo response bson unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()))
			return rsp, err
		}
		log.Logger.Info("GetDiscountCodeInfo response success", log.Any("rsp", rsp), log.String("trace_id", traceId), log.Any("response", res.GetParams()))
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("GetDiscountCodeInfo timeout", log.String("trace_id", traceId))
		return rsp, fmt.Errorf("GetDiscountCodeInfo timeout")
	}

	log.Logger.Info("GetDiscountCodeInfo end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (d *dao) GetOrderList(ctx context.Context, req *pb.GetOrderListReq) (*pb.GetOrderListRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := new(pb.GetOrderListRsp)

	log.Logger.Info("GetOrderList start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDGetOrderList,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("GetOrderList add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
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
			log.Logger.Error("GetOrderList response bson unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("GetOrderList response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("GetOrderList response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		if err := bson.Unmarshal(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusSystemError
			rsp.Msg = model.MsgParamsErr
			log.Logger.Error("GetOrderList response bson unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		log.Logger.Debug("GetOrderList response success", log.Any("rsp", rsp), log.String("trace_id", traceId), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("GetOrderList timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("GetOrderList timeout")
	}

	log.Logger.Info("GetOrderList end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (d *dao) VipIOSPromotionSign(ctx context.Context, req *pb.GetVipIOSPromotionSignReq) (*pb.GetVipIOSPromotionSignRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := new(pb.GetVipIOSPromotionSignRsp)

	log.Logger.Info("VipIOSPromotionSign start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)

	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDGetVipIOSPromotionSign,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		log.Logger.Error("VipIOSPromotionSign add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()

		ret := model.ResponseBytes{}
		if err := bson.Unmarshal(res.GetData(), &ret); err != nil {
			log.Logger.Error("VipIOSPromotionSign response bson unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("VipIOSPromotionSign response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("VipIOSPromotionSign response result failed", log.String("trace_id", traceId), log.Any("result", ret), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, nil
		}

		if err := bson.Unmarshal(ret.Data, &rsp.Data); err != nil {
			log.Logger.Error("VipIOSPromotionSign response bson unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
			return rsp, err
		}
		log.Logger.Debug("VipIOSPromotionSign response success", log.Any("rsp", rsp), log.String("trace_id", traceId), log.Any("response", res.GetParams()), log.String("params", req.GetParams()))
	case <-timer.C:
		log.Logger.Error("VipIOSPromotionSign timeout", log.String("trace_id", traceId), log.String("params", req.GetParams()))
		return rsp, fmt.Errorf("GetOrderList timeout")
	}

	log.Logger.Info("VipIOSPromotionSign end", log.String("trace_id", traceId))
	log.Logger.Debug("VipIOSPromotionSign end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (d *dao) VipPrice(ctx context.Context, req *pb.VipPriceReq) (*pb.VipPriceRsp, error) {
	requestID := d.GenerateID()
	traceId := util.GetTraceid(ctx)
	rsp := new(pb.VipPriceRsp)

	log.Logger.Info("VipPrice start", log.String("trace_id", traceId), log.String("params", req.GetParams()))

	waitChan := d.addStreamResponseWaitChan(requestID)
	defer d.delStreamResponseWaitChan(requestID)
	if err := d.addStreamRequest(&pb.StreamRsp{
		Cmd:       model.CMDGetVipPrice,
		Token:     d.conf.Node.Token,
		RequestId: requestID,
		Signature: req.GetSignature(),
		Params:    req.GetParams(),
		Data:      req.GetData(),
		TraceId:   traceId,
	}, model.ARES_PROXY); err != nil {
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgSystemErr
		log.Logger.Error("VipPrice add proxy request error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.String("params", req.GetParams()))
		return rsp, err
	}

	timer := time.NewTimer(util.W3PTimeout * time.Second)
	select {
	case res := <-waitChan:
		timer.Stop()

		ret := model.ResponseBytes{}
		if err := bson.Unmarshal(res.GetData(), &ret); err != nil {
			rsp.Code = model.StatusServiceCheckErr
			rsp.Msg = model.MsgParamsErr
			log.Logger.Warn("VipPrice response bson unmarshal error", log.String("trace_id", traceId), log.Any("response", res.GetParams()))
			return rsp, err
		}

		rsp.Code = ret.Code
		rsp.Msg = ret.Msg
		if ret.Code > model.StatusSystemErrorCode {
			log.Logger.Error("VipPrice response rsp error", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, errors.New(ret.Msg)
		}

		if ret.Code != model.StatusOK {
			log.Logger.Warn("VipPrice response not good", log.String("trace_id", traceId), log.Any("rsp", rsp))
			return rsp, nil
		}

		if err := bson.Unmarshal(ret.Data, &rsp.Data); err != nil {
			rsp.Code = model.StatusServiceCheckErr
			rsp.Msg = model.MsgParamsErr
			log.Logger.Warn("VipPrice response bson unmarshal error", log.String("trace_id", traceId), log.String("errmsg", err.Error()), log.Any("response", res.GetParams()))
			return rsp, err
		}
		log.Logger.Info("VipPrice response success", log.Any("rsp", rsp), log.String("trace_id", traceId), log.Any("response", res.GetParams()))
	case <-timer.C:
		rsp.Code = model.StatusSystemError
		rsp.Msg = model.MsgTimeoutErr
		log.Logger.Error("VipPrice timeout", log.String("trace_id", traceId))
		return rsp, fmt.Errorf("VipPrice timeout")
	}

	log.Logger.Info("VipPrice end", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}
