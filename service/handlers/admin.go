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
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/web3password/jewel/encode"
	"github.com/web3password/satis/log"
	"github.com/web3password/satis/model"
	pb "github.com/web3password/w3p-protobuf/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"gopkg.in/mgo.v2/bson"
)

// AdminAddMember .
func AdminAddMember(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.AdminAddMemberReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Debug("AdminAddMember request", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.AdminAddMember(gtx, req)
	if err != nil {
		log.Logger.Error("AdminAddMember rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("AdminAddMember rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Debug("AdminAddMember end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("rsp", rsp))

	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

// AdminUpdateMember .
func AdminUpdateMember(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.AdminUpdateMemberReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Debug("AdminUpdateMember start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.AdminUpdateMember(gtx, req)
	if err != nil {
		log.Logger.Error("AdminUpdateMember rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("AdminUpdateMember rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Debug("AdminUpdateMember end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("rsp", rsp))

	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

// AdminBatchImportMember .
// batch add member use addMember api , this api is abandoned
func AdminBatchImportMember(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.AdminBatchImportMemberReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Debug("AdminBatchImportMember start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.AdminBatchImportMember(gtx, req)
	if err != nil {
		log.Logger.Error("AdminBatchImportMember rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("AdminBatchImportMember rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Debug("AdminBatchImportMember end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("rsp", rsp))

	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

// AdminRemoveMember .
func AdminRemoveMember(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.AdminRemoveMemberReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Debug("AdminRemoveMember start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.AdminRemoveMember(gtx, req)
	if err != nil {
		log.Logger.Error("AdminRemoveMember rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("AdminRemoveMember rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Debug("AdminRemoveMember end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("rsp", rsp))

	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

// AdminGetMemberList .
func AdminGetMemberList(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}
	obj := value.(*encode.Web3PasswordRequestBsonStruct)
	req := &pb.AdminGetMemberListReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	log.Logger.Debug("AdminGetMemberList start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.AdminGetMemberList(gtx, req)
	if err != nil {
		log.Logger.Error("AdminGetMemberList rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("AdminGetMemberList rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	rspData := make([]*model.AdminMemberInfo, 0)
	for _, item := range rsp.Data.GetList() {
		member := &model.AdminMemberInfo{
			MemberAddr:         item.MemberAddr,
			MemberData:         item.MemberData,
			AdminShareMnemonic: item.AdminShareMnemonic,
			Role:               item.Role,
			MemberSign:         item.MemberSign,
		}
		rspData = append(rspData, member)
	}

	wrapData := model.AdminMemberListRsp{
		List: rspData,
	}

	log.Logger.Debug("AdminGetMemberList end, datadata", log.String("trace_id", ctx.GetString("trace_id")), log.Any("data", wrapData))
	bytes, _ := bson.Marshal(wrapData)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

func AdminTransferSuperAdmin(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.AdminTransferSuperAdminReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	log.Logger.Debug("AdminTransferSuperAdmin start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	//var metadata runtime.ServerMetadata
	//rsp, err := userClient.AdminTransferSuperAdmin(context.Background(), req, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	rsp, err := userClient.AdminTransferSuperAdmin(gtx, req)
	if err != nil {
		log.Logger.Error("AdminTransferSuperAdmin error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("AdminTransferSuperAdmin rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	bytes, _ := bson.Marshal(rsp.GetData())
	log.Logger.Debug("AdminTransferSuperAdmin end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

func AdminOperationHistory(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.AdminOperationHistoryReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	log.Logger.Debug("AdminOperationHistory start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	//var metadata runtime.ServerMetadata
	//rsp, _ := userClient.AdminOperationHistory(context.Background(), req, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	rsp, err := userClient.AdminOperationHistory(gtx, req)
	if err != nil {
		log.Logger.Error("AdminOperationHistory error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("AdminOperationHistory rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	rspData := make([]model.AdminOperationHistoryBsonV2, 0)

	for _, item := range rsp.GetData() {
		var ct model.AdminOperationHistoryBsonV2
		json.Unmarshal([]byte(item.Content), &ct)
		rspData = append(rspData, ct)
	}

	wrapData := map[string]interface{}{
		"list": rspData,
	}
	log.Logger.Debug("AdminOperationHistory end, datadata", log.String("trace_id", ctx.GetString("trace_id")), log.Any("len", len(wrapData)))
	bytes, _ := bson.Marshal(wrapData)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

func AdminRegister(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "params error", []byte(""))
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.AdminRegisterReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	log.Logger.Debug("AdminRegister start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))

	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.AdminRegister(gtx, req)
	if err != nil {
		log.Logger.Error("AdminRegister error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("AdminRegister rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	orgRegisterRsp := model.VipRegister{}
	orgRegisterRsp.OrgId = rsp.GetData().GetOrgId()
	data, _ := bson.Marshal(orgRegisterRsp)
	log.Logger.Debug("AdminRegister end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", data))
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), data)
}

func Authorization(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "params error", []byte(""))
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.AdminAuthorizationReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
	}
	log.Logger.Debug("Authorization start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))

	var metadata runtime.ServerMetadata
	rsp, err := userClient.AdminAuthorization(context.Background(), req, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	if err != nil {
		log.Logger.Error("Authorization error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("Authorization rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	ret := &model.AdminAuthorizationRsp{
		Authorization: rsp.GetData().GetAuthorization(),
	}

	bytes, _ := bson.Marshal(ret)
	log.Logger.Debug("Authorization end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	return
}

// GetAdminMnemonic .
func GetAdminMnemonic(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.GetAdminMnemonicReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Debug("GetAdminMnemonic start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))

	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.GetAdminMnemonic(gtx, req)
	if err != nil {
		log.Logger.Error("GetAdminMnemonic error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("GetAdminMnemonic rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Debug("GetAdminMnemonic rsp", log.String("trace_id", ctx.GetString("trace_id")))
	if len(rsp.GetData().GetAdminShareMnemonic()) <= 0 && len(rsp.GetData().GetMemberShareMnemonic()) <= 0 {
		log.Logger.Error("GetAdminMnemonic error 2", log.String("trace_id", ctx.GetString("trace_id")))
		Response(ctx, model.StatusDataEmpty, model.NoAdminShareMnemonic, emptyByte)
		return
	}

	adminShareMnemonicInfo := &model.AdminShareMnemonicRsp{
		AdminShareMnemonic:  rsp.GetData().GetAdminShareMnemonic(),
		MemberShareMnemonic: rsp.GetData().GetMemberShareMnemonic(),
	}

	log.Logger.Debug("GetAdminMnemonic end", log.String("trace_id", ctx.GetString("trace_id")))
	bytes, _ := bson.Marshal(adminShareMnemonicInfo)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	return
}

func AdminUpdateOrgInfo(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}
	obj := value.(*encode.Web3PasswordRequestBsonStruct)
	req := &pb.AdminUpdateOrgInfoReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	log.Logger.Debug("AdminUpdateOrgInfo start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	var metadata runtime.ServerMetadata
	rsp, err := userClient.AdminUpdateOrgInfo(context.Background(), req, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	var data []byte
	if err != nil {
		log.Logger.Error("AdminUpdateOrgInfo error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("AdminUpdateOrgInfo rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	data, _ = bson.Marshal(rsp.Data)
	log.Logger.Debug("AdminUpdateOrgInfo end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), data)
}

func AdminGetOrgInfo(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}
	obj := value.(*encode.Web3PasswordRequestBsonStruct)
	req := &pb.AdminGetOrgInfoReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	log.Logger.Debug("AdminGetOrgInfo start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", req.GetParams()))
	var metadata runtime.ServerMetadata
	rsp, err := userClient.AdminGetOrgInfo(context.Background(), req, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	if err != nil {
		log.Logger.Error("AdminGetOrgInfo error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("AdminGetOrgInfo rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	orgData := model.OrgInfoData{}
	json.Unmarshal([]byte(rsp.GetData()), &orgData)
	data, _ := bson.Marshal(orgData)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), data)
	log.Logger.Debug("AdminGetOrgInfo end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp.code", rsp.GetCode()))
}
