/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/
package handlers

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/web3password/jewel/encode"
	"github.com/web3password/satis/log"
	"github.com/web3password/satis/model"
	pb "github.com/web3password/w3p-protobuf/user"
)

func Register(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		log.Logger.Info("satis handler register get request error", log.Any("request", value))
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.RegisterReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
	}

	log.Logger.Info("Register start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))

	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.Register(gtx, req)
	if err != nil {
		log.Logger.Error("Register rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("Register rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Info("Register end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

func GetPersonalSignAddress(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.GetPersonalSignAddressReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
	}

	log.Logger.Debug("GetPersonalSignAddress start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))

	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.GetPersonalSignAddress(gtx, req)
	if err != nil {
		log.Logger.Error("GetPersonalSignAddress rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err), log.Any("req", req))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("GetPersonalSignAddress rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	type Data struct {
		Signs []*model.PersonalSignResponse `bson:"signs"`
	}
	data := &Data{Signs: make([]*model.PersonalSignResponse, 0)}
	for _, v := range rsp.GetData().GetSigns() {
		sign := &model.PersonalSignResponse{
			Sign:   v.GetSign(),
			Params: v.GetParams(),
		}
		data.Signs = append(data.Signs, sign)
	}

	log.Logger.Debug("GetPersonalSignAddress end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("src.data", rsp.GetData()), log.Any("rsp.data", data))
	bytes, _ := bson.Marshal(data)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	return
}

func GetVIPInfo(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.GetVIPInfoReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
	}

	log.Logger.Debug("GetVIPInfo start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))

	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.GetVIPInfo(gtx, req)
	if err != nil {
		log.Logger.Error("GetVIPInfo error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err), log.Any("req", req))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("GetVIPInfo rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	orgInfoList := make([]model.UserOrgInfo, 0)

	for _, info := range rsp.GetData().OrgInfo {
		orgList := model.UserOrgInfo{
			OrgId:          info.OrgId,
			OrgName:        info.OrgName,
			Logo:           info.Logo,
			SelfHostUrl:    info.SelfHostUrl,
			SharedMnemonic: info.MemberShareMnemonic,
		}
		orgInfoList = append(orgInfoList, orgList)
	}
	vip := &model.VIPInfoResult{
		Signature: rsp.GetData().GetSignature(),
		Params:    rsp.GetData().GetParams(),
		OrgInfo:   orgInfoList,
	}

	log.Logger.Debug("GetVIPInfo end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("src.params", rsp.GetData().GetParams()), log.Any("rsp.params", vip.Params))
	bytes, _ := bson.Marshal(vip)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	return
}

func GetUserInfo(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.GetUserInfoReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
	}

	log.Logger.Debug("GetUserInfo start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))

	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.GetUserInfo(gtx, req)
	if err != nil {
		log.Logger.Error("GetUserInfo error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err), log.Any("req", req))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("GetUserInfo rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	userinfo := &model.UserInfoResponse{
		Address:     rsp.GetData().GetAddr(),
		ChainId:     rsp.GetData().GetChainId(),
		InviteCode:  rsp.GetData().GetInviteCode(),
		StorageType: rsp.GetData().GetStorageType(),
	}

	log.Logger.Debug("GetUserInfo end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("src.data", rsp.GetData()), log.Any("rsp.data", userinfo))
	bytes, _ := bson.Marshal(userinfo)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

func Response(ctx *gin.Context, code int, msg string, data []byte) {
	response, _ := encode.Web3PasswordResponseBsonEncode(code, msg, data)
	if IsHttpWithTraceID() {
		ctx.Header("X-Trace-id", ctx.GetString("trace_id"))
	}
	_, _ = ctx.Writer.Write(response)
}

func StorageReport(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.StorageReportReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
	}
	log.Logger.Debug("StorageReport start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	var metadata runtime.ServerMetadata
	rsp, err := userClient.StorageReport(context.Background(), req, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	if err != nil {
		log.Logger.Error("StorageReport rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("StorageReport rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Debug("StorageReport end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

func StorageStat(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.StorageStatReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
	}

	log.Logger.Debug("StorageStat start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))

	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.StorageStat(gtx, req)
	if err != nil {
		log.Logger.Error("StorageStat rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err), log.Any("req", req))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("StorageStat rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	stat := model.StorageStatBson{
		TotalReal:          rsp.GetData().GetTotalReal(),
		TotalHumanReadable: rsp.GetData().GetTotalHumanReadable(),
		Used:               rsp.GetData().GetUsed(),
		Left:               rsp.GetData().GetLeft(),
	}
	bytes, _ := bson.Marshal(stat)
	log.Logger.Debug("StorageStat end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("src.data", rsp.GetData()), log.Any("rsp.data", stat))
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

// GetVersionConfig .
func GetVersionConfig(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.GetVersionConfigReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Debug("GetVersionConfig start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))

	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.GetVersionConfig(gtx, req)
	if err != nil {
		log.Logger.Error("GetVersionConfig error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err), log.Any("req", req))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("GetVersionConfig rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	configInfo := &model.GetVersionConfigRsp{
		VersionValue:                  rsp.GetData().GetVersionValue(),
		VersionName:                   rsp.GetData().GetVersionName(),
		ShowRebate:                    rsp.GetData().GetShowRebate(),
		RebateRate:                    rsp.GetData().GetRebateRate(),
		RecordCountLimit:              rsp.GetData().GetRecordCountLimit(),
		OneRecordWebsiteLimit:         rsp.GetData().GetOneRecordWebsiteLimit(),
		OneRecordLinkwebsiteLimit:     rsp.GetData().OneRecordLinkwebsiteLimit,
		OneRecordAttachmentCountLimit: rsp.GetData().GetOneRecordAttachmentCountLimit(),
		OneUserAttachmemtSpaceLimit:   rsp.GetData().GetOneUserAttachmentSpaceLimit(),
		OneAttachmentSizeLimit:        rsp.GetData().GetOneAttachmentSizeLimit(),
		SharefolderCreateLimit:        rsp.GetData().GetSharefolderCreateLimit(),
		SharefolderRecordLimit:        rsp.GetData().GetSharefolderRecordLimit(),
		SharefolderMemberLimit:        rsp.GetData().GetSharefolderMemberLimit(),
		VersionUser:                   rsp.GetData().GetVersionUser(),
		AdminMemberLimit:              rsp.GetData().GetAdminMemberLimit(),
		OneRecordSizeLimit:            rsp.GetData().GetOneRecordSizeLimit(),
	}

	log.Logger.Debug("GetVersionConfig end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("src.data", rsp.GetData()), log.Any("rsp.data", configInfo))
	bytes, _ := bson.Marshal(configInfo)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	return
}
