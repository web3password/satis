/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/
package handlers

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/web3password/jewel/encode"
	"github.com/web3password/satis/log"
	"github.com/web3password/satis/model"
	pb "github.com/web3password/w3p-protobuf/user"
	"google.golang.org/grpc/metadata"
	"gopkg.in/mgo.v2/bson"
)

func VipGetConfig(ctx *gin.Context) {
	var (
		traceId = ctx.GetString("trace_id")
	)
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.VipGetConfigReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", traceId,
	))
	rsp, err := userClient.VipGetConfig(gtx, req)
	if err != nil {
		log.Logger.Error("VipGetConfig get error", log.String("trace_id", traceId), log.Any("rsp", rsp))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("VipGetConfig rsp warning", log.String("trace_id", traceId), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	data, err := fmtResponseDataForVipGetConfig(rsp.GetData())
	if err != nil {
		log.Logger.Error("VipGetConfig error", log.String("trace_id", traceId))
		Response(ctx, int(rsp.GetCode()), err.Error(), emptyByte)
		return
	}

	log.Logger.Debug("VipGetConfig end", log.String("trace_id", traceId),
		log.Any("src.data", rsp.GetData()), log.Any("rsp.data", data))
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), data)
	return
}

func fmtResponseDataForVipGetConfig(source *pb.VipGetConfigRsp_Data) ([]byte, error) {
	sourceJSON, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}

	var target model.VipGetConfigRspData
	err = json.Unmarshal(sourceJSON, &target)
	if err != nil {
		return nil, err
	}

	bytes, err := bson.Marshal(target)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func SubscriptionList(ctx *gin.Context) {
	var (
		traceId = ctx.GetString("trace_id")
	)
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.VipSubscriptionListReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	log.Logger.Info("SubscriptionList start", log.String("trace_id", ctx.GetString("trace_id")), log.String("params", obj.ParamsStr))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", traceId,
	))
	rsp, err := userClient.VipSubscriptionList(gtx, req)
	if err != nil {
		log.Logger.Error("VipSubscriptionList error", log.String("trace_id", traceId), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("VipSubscriptionList rsp warning", log.String("trace_id", traceId), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	data, err := fmtResponseDataForVipSubscriptionList(rsp.GetData())
	if err != nil {
		log.Logger.Error("SubscriptionList error", log.String("trace_id", traceId), log.Error(err))
		Response(ctx, int(rsp.GetCode()), err.Error(), emptyByte)
		return
	}

	log.Logger.Info("SubscriptionList end", log.String("trace_id", ctx.GetString("trace_id")))
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), data)
	return
}

func fmtResponseDataForVipSubscriptionList(source *pb.VipSubscriptionListRsp_Data) ([]byte, error) {
	sourceJSON, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}

	var target model.VipSubscriptionListRspData
	err = json.Unmarshal(sourceJSON, &target)
	if err != nil {
		return nil, err
	}

	bytes, err := bson.Marshal(target)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func VipPaymentList(ctx *gin.Context) {
	var (
		traceId = ctx.GetString("trace_id")
	)
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.VipPaymentListReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", traceId,
	))
	rsp, err := userClient.VipPaymentList(gtx, req)
	if err != nil {
		log.Logger.Error("VipPaymentList error", log.String("trace_id", traceId), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("VipPaymentList rsp warning", log.String("trace_id", traceId), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	data, err := fmtResponseDataForVipPaymentList(rsp.GetData())
	if err != nil {
		log.Logger.Error("VipPaymentList error", log.String("trace_id", ctx.GetString("trace_id")))
		Response(ctx, int(rsp.GetCode()), err.Error(), emptyByte)
		return
	}

	log.Logger.Debug("VipPaymentList end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("src.data", rsp.GetData()), log.Any("rsp.data", data))
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), data)
	return
}

func fmtResponseDataForVipPaymentList(source *pb.VipPaymentListRsp_Data) ([]byte, error) {
	sourceJSON, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}

	var target model.VipPaymentListRspData
	err = json.Unmarshal(sourceJSON, &target)
	if err != nil {
		return nil, err
	}
	log.Logger.Debug("VipPaymentList end2", log.String("trace_id", string(sourceJSON)))
	bytes, err := bson.Marshal(target)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func VipCreateOrder(ctx *gin.Context) {
	var (
		traceId = ctx.GetString("trace_id")
	)
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.VipCreateOrderReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", traceId,
	))
	rsp, err := userClient.VipCreateOrder(gtx, req)
	if err != nil {
		log.Logger.Error("VipCreateOrder error", log.String("trace_id", traceId), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("VipCreateOrder rsp warning", log.String("trace_id", traceId), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	data, err := fmtResponseDataForVipCreateOrder(rsp.GetData())
	if err != nil {
		log.Logger.Error("VipCreateOrder error", log.String("trace_id", traceId), log.Error(err))
		Response(ctx, int(rsp.GetCode()), err.Error(), emptyByte)
		return
	}

	log.Logger.Debug("VipCreateOrder end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("src.data", rsp.GetData()), log.Any("rsp.data", data))
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), data)
	return
}

func fmtResponseDataForVipCreateOrder(source *pb.VipCreateOrderRsp_Data) ([]byte, error) {
	sourceJSON, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}

	var target model.VipCreteOrderRspData
	err = json.Unmarshal(sourceJSON, &target)
	if err != nil {
		return nil, err
	}

	bytes, err := bson.Marshal(target)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func VipCheckOrder(ctx *gin.Context) {
	var (
		traceId = ctx.GetString("trace_id")
	)
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.VipCheckOrderReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", traceId,
	))
	rsp, err := userClient.VipCheckOrder(gtx, req)
	if err != nil {
		log.Logger.Error("VipCheckOrder userClient.VipCheckOrder error",
			log.String("trace_id", traceId),
			log.Any("req", req),
			log.Error(err),
		)
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("VipCheckOrder rsp warning", log.String("trace_id", traceId), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	data, err := fmtResponseDataForVipCheckOrder(rsp.GetData())
	if err != nil {
		log.Logger.Error("VipCheckOrder fmtResponseDataForVipCheckOrder error",
			log.String("trace_id", ctx.GetString("trace_id")),
			log.Any("req", req),
			log.String("errMsg", err.Error()),
		)
		Response(ctx, int(rsp.GetCode()), err.Error(), emptyByte)
		return
	}

	log.Logger.Debug("VipCheckOrder resp",
		log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("req", req),
		log.Any("rsp", rsp),
		log.Any("data", data),
	)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), data)
	return
}

func fmtResponseDataForVipCheckOrder(source *pb.VipCheckOrderRsp_Data) ([]byte, error) {
	sourceJSON, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}

	var target model.VipCheckOrderRspData
	err = json.Unmarshal(sourceJSON, &target)
	if err != nil {
		return nil, err
	}

	bytes, err := bson.Marshal(target)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func VipAppleVerifyReceipt(ctx *gin.Context) {
	var (
		traceId = ctx.GetString("trace_id")
	)
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.VipAppleVerifyReceiptReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", traceId,
	))
	rsp, err := userClient.VipAppleVerifyReceipt(gtx, req)
	if err != nil {
		log.Logger.Error("VipAppleVerifyReceipt error", log.String("trace_id", traceId), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("VipAppleVerifyReceipt rsp warning", log.String("trace_id", traceId), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	data, err := fmtResponseDataForVipAppleVerifyReceipt(rsp.GetData())
	if err != nil {
		log.Logger.Error("VipAppleVerifyReceipt fmtResponseDataForVipCheckOrder error",
			log.String("trace_id", ctx.GetString("trace_id")),
			log.Any("req", req),
			log.String("errMsg", err.Error()),
		)
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	log.Logger.Debug("VipAppleVerifyReceipt resp",
		log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("req", req),
		log.Any("rsp", rsp),
		log.Any("data", data),
	)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), data)
	return
}

func fmtResponseDataForVipAppleVerifyReceipt(source *pb.VipAppleVerifyReceiptRsp_Data) ([]byte, error) {
	sourceJSON, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}

	var target model.VipAppleVerifyReceiptRspData
	err = json.Unmarshal(sourceJSON, &target)
	if err != nil {
		return nil, err
	}

	bytes, err := bson.Marshal(target)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func VipDiscount(ctx *gin.Context) {
	var (
		traceId = ctx.GetString("trace_id")
	)
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.GetDiscountCodeInfoReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
	}
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", traceId,
	))
	rsp, err := userClient.GetDiscountCodeInfo(gtx, req)
	if err != nil {
		log.Logger.Error("GetDiscountCodeInfo error", log.String("trace_id", traceId), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("GetDiscountCodeInfo rsp warning", log.String("trace_id", traceId), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	data, _ := bson.Marshal(rsp.GetData())
	log.Logger.Debug("VipDiscount end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("src.data", rsp.GetData()), log.Any("rsp.data", data))
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), data)
	return
}

func GetOrderList(ctx *gin.Context) {
	var (
		traceId = ctx.GetString("trace_id")
	)
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.GetOrderListReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
	}
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", traceId,
	))
	rsp, err := userClient.GetOrderList(gtx, req)
	if err != nil {
		log.Logger.Error("userClient.GetOrderList error", log.String("trace_id", traceId), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}
	if rsp.Code != model.StatusOK {
		log.Logger.Warn("GetOrderList rsp warning", log.String("trace_id", traceId), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	data, err := fmtResponseDataForGetOrderList(rsp.GetData())
	if err != nil {
		log.Logger.Error("GetOrderList error", log.String("trace_id", ctx.GetString("trace_id")))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	log.Logger.Debug("GetOrderList end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("src.data", rsp.GetData()), log.Any("rsp.data", data))
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), data)
	return
}

func fmtResponseDataForGetOrderList(source *pb.GetOrderListRsp_Data) ([]byte, error) {
	sourceJSON, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}

	var target model.GetOrderListRspData
	err = json.Unmarshal(sourceJSON, &target)
	if err != nil {
		return nil, err
	}
	log.Logger.Debug("GetOrderList end2", log.String("trace_id", string(sourceJSON)))
	bytes, err := bson.Marshal(target)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func GetVipIOSPromotionSign(ctx *gin.Context) {
	var (
		traceId = ctx.GetString("trace_id")
	)
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.GetVipIOSPromotionSignReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", traceId,
	))
	rsp, err := userClient.VipIOSPromotionSign(gtx, req)
	if err != nil {
		log.Logger.Error("userClient.GetVipIOSPromotionSign error", log.String("trace_id", traceId), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	data, err := bson.Marshal(rsp.GetData())
	if err != nil {
		log.Logger.Error("GetVipIOSPromotionSign error", log.String("trace_id", ctx.GetString("trace_id")))
		Response(ctx, int(rsp.GetCode()), err.Error(), emptyByte)
		return
	}

	log.Logger.Debug("GetVipIOSPromotionSign end", log.String("trace_id", ctx.GetString("trace_id")))

	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), data)
	return
}

func VipPrice(ctx *gin.Context) {
	var (
		traceId = ctx.GetString("trace_id")
	)
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.VipPriceReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", traceId,
	))
	rsp, err := userClient.VipPrice(gtx, req)
	if err != nil {
		log.Logger.Error("VipPrice error", log.String("trace_id", traceId), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("VipPrice rsp warning", log.String("trace_id", traceId), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	data, err := fmtVipPrice(rsp.GetData())
	if err != nil {
		log.Logger.Error("VipPrice fmtVipPrice error", log.String("trace_id", ctx.GetString("trace_id")))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	log.Logger.Debug("VipPrice end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("src.data", rsp.GetData()), log.Any("rsp.data", data))
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), data)
	return
}

func fmtVipPrice(source *pb.VipPriceRsp_Data) ([]byte, error) {
	sourceJSON, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}

	var target model.VipPriceRspData
	err = json.Unmarshal(sourceJSON, &target)
	if err != nil {
		return nil, err
	}
	log.Logger.Debug("VipPrice end2", log.String("trace_id", string(sourceJSON)))
	bytes, err := bson.Marshal(target)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
