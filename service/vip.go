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

func (s *Service) VipGetConfig(ctx context.Context, req *pb.VipGetConfigReq) (*pb.VipGetConfigRsp, error) {
	rsp := new(pb.VipGetConfigRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"

	params := model.VipGetConfigParams{}
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("VipGetConfig start", log.String("trace_id", traceId), log.Any("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("VipGetConfig params parse fail", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("VipGetConfig params fail", log.String("trace_id", traceId), log.Any("errMsg", errMsg))
		return rsp, nil
	}

	ret, err := s.dao.VipGetConfig(ctx, req)
	if err != nil {
		log.Logger.Error("VipGetConfig request service error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		rsp.Msg = err.Error()
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("VipGetConfig request error", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("VipGetConfig request failed", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.GetData()
	log.Logger.Info("VipGetConfig success", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (s *Service) VipSubscriptionList(ctx context.Context, req *pb.VipSubscriptionListReq) (*pb.VipSubscriptionListRsp, error) {
	rsp := new(pb.VipSubscriptionListRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"

	params := model.VipSubscriptionListParams{}
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("VipSubscriptionList start", log.String("trace_id", traceId), log.Any("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("VipSubscriptionList params parse fail", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("VipSubscriptionList params fail", log.String("trace_id", traceId), log.Any("errMsg", errMsg))
		return rsp, nil
	}

	_, err := util.ParseAuthParams(params.Auth)
	if err != nil {
		rsp.Code = model.StatusAuthErr
		rsp.Msg = err.Error()
		log.Logger.Warn("VipSubscriptionList ParseAuthParams personal auth fail", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	ret, err := s.dao.VipSubscriptionList(ctx, req)
	if err != nil {
		log.Logger.Error("VipSubscriptionList request service error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		rsp.Msg = err.Error()
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("VipSubscriptionList request error", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("VipSubscriptionList request failed", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.GetData()
	log.Logger.Info("VipSubscriptionList success", log.String("trace_id", traceId))
	return rsp, nil
}

func (s *Service) VipPaymentList(ctx context.Context, req *pb.VipPaymentListReq) (*pb.VipPaymentListRsp, error) {
	rsp := new(pb.VipPaymentListRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"

	params := model.VipPaymentListParams{}
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("VipPaymentList start", log.String("trace_id", traceId), log.Any("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("VipPaymentList params parse fail", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("VipPaymentList params fail", log.String("trace_id", traceId), log.Any("errMsg", errMsg))
		return rsp, nil
	}

	_, err := util.ParseAuthParams(params.Auth)
	if err != nil {
		rsp.Code = model.StatusAuthErr
		rsp.Msg = err.Error()
		log.Logger.Warn("VipPaymentList ParseAuthParams personal auth fail", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	ret, err := s.dao.VipPaymentList(ctx, req)
	if err != nil {
		log.Logger.Error("VipPaymentList request service error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		rsp.Msg = err.Error()
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("VipPaymentList request error", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("VipPaymentList request failed", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.GetData()
	log.Logger.Info("VipPaymentList success", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (s *Service) VipCreateOrder(ctx context.Context, req *pb.VipCreateOrderReq) (*pb.VipCreateOrderRsp, error) {
	rsp := new(pb.VipCreateOrderRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"

	params := model.VipCreateOrderParams{}
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("VipCreateOrder start", log.String("trace_id", traceId), log.Any("params", req.GetParams()), log.Any("data", base64.StdEncoding.EncodeToString(req.GetData())))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("VipCreateOrder params parse fail", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("VipCreateOrder params fail", log.String("trace_id", traceId), log.Any("errMsg", errMsg))
		return rsp, nil
	}

	_, err := util.ParseAuthParams(params.Auth)
	if err != nil {
		rsp.Code = model.StatusAuthErr
		rsp.Msg = err.Error()
		log.Logger.Warn("VipSubscriptionList ParseAuthParams personal auth fail", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	ret, err := s.dao.VipCreateOrder(ctx, req)
	if err != nil {
		log.Logger.Error("VipCreateOrder request service error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		rsp.Msg = err.Error()
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("VipCreateOrder request error", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("VipCreateOrder request failed", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}
	rsp.Data = ret.GetData()
	log.Logger.Info("VipCreateOrder success", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (s *Service) VipCheckOrder(ctx context.Context, req *pb.VipCheckOrderReq) (*pb.VipCheckOrderRsp, error) {
	rsp := new(pb.VipCheckOrderRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"

	params := model.VipCheckOrderParams{}
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("VipCheckOrder start", log.String("trace_id", traceId), log.Any("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("VipCheckOrder params parse fail", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("VipCheckOrder params fail", log.String("trace_id", traceId), log.Any("errMsg", errMsg))
		return rsp, nil
	}

	_, err := util.ParseAuthParams(params.PersonalAuth)
	if err != nil {
		rsp.Code = model.StatusAuthErr
		rsp.Msg = err.Error()
		log.Logger.Warn("VipCheckOrder ParseAuthParams personal auth fail", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	ret, err := s.dao.VipCheckOrder(ctx, req)
	if err != nil {
		log.Logger.Error("VipCheckOrder request service error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		rsp.Msg = err.Error()
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("VipCheckOrder request error", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("VipCheckOrder request failed", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}
	rsp.Data = ret.GetData()
	log.Logger.Info("VipCheckOrder success", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (s *Service) VipAppleVerifyReceipt(ctx context.Context, req *pb.VipAppleVerifyReceiptReq) (*pb.VipAppleVerifyReceiptRsp, error) {
	rsp := new(pb.VipAppleVerifyReceiptRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"

	params := model.VipAppleVerifyReceiptParams{}
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("VipAppleVerifyReceipt start", log.String("trace_id", traceId), log.Any("params", req.GetParams()), log.Any("data", base64.StdEncoding.EncodeToString(req.GetData())))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("VipAppleVerifyReceipt params parse fail", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("VipAppleVerifyReceipt params fail", log.String("trace_id", traceId), log.Any("errMsg", errMsg))
		return rsp, nil
	}

	_, err := util.ParseAuthParams(params.PersonalAuth)
	if err != nil {
		rsp.Code = model.StatusAuthErr
		rsp.Msg = err.Error()
		log.Logger.Warn("VipAppleVerifyReceipt ParseAuthParams personal auth fail", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	ret, err := s.dao.VipAppleVerifyReceipt(ctx, req)
	if err != nil {
		log.Logger.Error("VipAppleVerifyReceipt request service error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		rsp.Msg = err.Error()
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("VipAppleVerifyReceipt request error", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("VipAppleVerifyReceipt request failed", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}
	rsp.Data = ret.GetData()
	log.Logger.Info("VipAppleVerifyReceipt success", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (s *Service) GetDiscountCodeInfo(ctx context.Context, req *pb.GetDiscountCodeInfoReq) (*pb.GetDiscountCodeInfoRsp, error) {
	rsp := new(pb.GetDiscountCodeInfoRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"

	params := model.GetDiscountCodeInfoParams{}
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("GetDiscountCodeInfo start", log.String("trace_id", traceId), log.Any("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("GetDiscountCodeInfo params parse fail", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("GetDiscountCodeInfo params fail", log.String("trace_id", traceId), log.Any("errMsg", errMsg))
		return rsp, nil
	}

	ret, err := s.dao.GetDiscountCodeInfo(ctx, req)
	if err != nil {
		log.Logger.Error("GetDiscountCodeInfo request service error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		rsp.Msg = err.Error()
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("GetDiscountCodeInfo request error", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("GetDiscountCodeInfo request failed", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}
	rsp.Data = ret.GetData()
	log.Logger.Info("GetDiscountCodeInfo success", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (s *Service) GetOrderList(ctx context.Context, req *pb.GetOrderListReq) (*pb.GetOrderListRsp, error) {
	rsp := new(pb.GetOrderListRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"

	params := model.GetOrderListParams{}
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("GetOrderList start", log.String("trace_id", traceId), log.Any("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("GetOrderList params parse fail", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("GetOrderList params fail", log.String("trace_id", traceId), log.Any("errMsg", errMsg))
		return rsp, nil
	}

	ret, err := s.dao.GetOrderList(ctx, req)
	if err != nil {
		log.Logger.Error("GetOrderList request service error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		rsp.Msg = err.Error()
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("GetOrderList request error", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("GetOrderList request failed", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}
	rsp.Data = ret.GetData()
	log.Logger.Info("GetOrderList success", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (s *Service) VipIOSPromotionSign(ctx context.Context, req *pb.GetVipIOSPromotionSignReq) (*pb.GetVipIOSPromotionSignRsp, error) {
	rsp := new(pb.GetVipIOSPromotionSignRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"

	params := model.VipIOSPromotionSignParams{}
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("VipIOSPromotionSign start", log.String("trace_id", traceId), log.Any("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = model.MsgParamsErr
		log.Logger.Warn("VipIOSPromotionSign params parse fail", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("VipIOSPromotionSign params fail", log.String("trace_id", traceId), log.Any("errMsg", errMsg))
		return rsp, nil
	}

	ret, err := s.dao.VipIOSPromotionSign(ctx, req)
	if err != nil {
		//rsp.Code = daoResp.GetCode() // panic
		rsp.Msg = err.Error()
		log.Logger.Error("VipIOSPromotionSign request service error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("VipIOSPromotionSign request error", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("VipIOSPromotionSign request failed", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}

	rsp.Data = ret.GetData()
	log.Logger.Info("VipIOSPromotionSign success", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}

func (s *Service) VipPrice(ctx context.Context, req *pb.VipPriceReq) (*pb.VipPriceRsp, error) {
	rsp := new(pb.VipPriceRsp)
	rsp.Code = model.StatusServiceCheckErr
	rsp.Msg = "system error"

	params := model.VipCreateOrderParams{}
	traceId := util.GetTraceid(ctx)
	log.Logger.Info("VipPrice start", log.String("trace_id", traceId), log.Any("params", req.GetParams()))
	if err := jsoniter.UnmarshalFromString(req.GetParams(), &params); err != nil {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = "params parse fail"
		log.Logger.Warn("VipPrice params parse fail", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		return rsp, nil
	}

	if errMsg, ok := params.Check(); !ok {
		rsp.Code = model.StatusParamsErr
		rsp.Msg = errMsg
		log.Logger.Warn("VipPrice params error", log.String("trace_id", traceId), log.Any("errMsg", errMsg))
		return rsp, nil
	}

	ret, err := s.dao.VipPrice(ctx, req)
	if err != nil {
		log.Logger.Error("VipPrice request service error", log.String("trace_id", traceId), log.String("errmsg", err.Error()))
		rsp.Msg = err.Error()
		return rsp, nil
	}

	rsp.Code = ret.Code
	rsp.Msg = ret.Msg

	if ret.Code > model.StatusSystemErrorCode {
		log.Logger.Error("VipPrice request error", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}

	if ret.Code != model.StatusOK {
		log.Logger.Warn("VipPrice request failed", log.String("trace_id", traceId), log.Any("ret", ret))
		return rsp, nil
	}
	rsp.Data = ret.GetData()
	log.Logger.Info("VipPrice success", log.String("trace_id", traceId), log.Any("rsp", rsp))
	return rsp, nil
}
