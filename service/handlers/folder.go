/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/
package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/web3password/jewel/encode"
	"github.com/web3password/satis/log"
	"github.com/web3password/satis/model"
	pb "github.com/web3password/w3p-protobuf/user"
	"google.golang.org/grpc/metadata"
	"gopkg.in/mgo.v2/bson"
)

func ShareFolderCreate(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.ShareFolderCreateReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Info("ShareFolderCreate start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))

	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.ShareFolderCreate(gtx, req)
	if err != nil {
		log.Logger.Error("ShareFolderCreate rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderCreate rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Info("ShareFolderCreate end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("src.data", rsp.GetData()), log.Any("rsp.data", rsp.GetData()))
	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	return
}

func ShareFolderUpdate(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.ShareFolderUpdateReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Info("ShareFolderUpdate start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))
	log.Logger.Debug("ShareFolderUpdate start", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("Signature", obj.SignatureStr),
		log.Any("params", obj.ParamsStr),
		log.Any("Data", obj.AppendData),
	)

	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.ShareFolderUpdate(gtx, req)
	if err != nil {
		log.Logger.Error("ShareFolderUpdate rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderUpdate rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Info("ShareFolderUpdate end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("src.data", rsp.GetData()), log.Any("rsp.data", rsp.GetData()))
	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	return
}

// ShareFolderDestroy .
func ShareFolderDestroy(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.ShareFolderDestroyReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Info("ShareFolderDestroy start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))
	log.Logger.Debug("ShareFolderDestroy start", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("Signature", obj.SignatureStr),
		log.Any("params", obj.ParamsStr),
		log.Any("Data", obj.AppendData),
	)
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.ShareFolderDestroy(gtx, req)
	if err != nil {
		log.Logger.Error("ShareFolderDestroy rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderDestroy rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Info("ShareFolderDestroy end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("src.data", rsp.GetData()), log.Any("rsp.data", rsp.GetData()))
	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

// ShareFolderAddMember .
func ShareFolderAddMember(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.ShareFolderAddMemberReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Info("ShareFolderAddMember start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))
	log.Logger.Debug("ShareFolderAddMember start debug", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.ShareFolderAddMember(gtx, req)
	if err != nil {
		log.Logger.Error("ShareFolderAddMember rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderAddMember rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Info("ShareFolderAddMember end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))

	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

// ShareFolderDeleteMember .
func ShareFolderDeleteMember(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.ShareFolderDeleteMemberReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Info("ShareFolderDeleteMember start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))
	log.Logger.Debug("ShareFolderDeleteMember start debug", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.ShareFolderDeleteMember(gtx, req)
	if err != nil {
		log.Logger.Error("ShareFolderDeleteMember rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderDeleteMember rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Info("ShareFolderDeleteMember end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))

	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

// ShareFolderUpdateMember .
func ShareFolderUpdateMember(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.ShareFolderUpdateMemberReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Info("ShareFolderUpdateMember start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))
	log.Logger.Debug("ShareFolderUpdateMember start debug", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.ShareFolderUpdateMember(gtx, req)
	if err != nil {
		log.Logger.Error("ShareFolderUpdateMember rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderUpdateMember rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Info("ShareFolderUpdateMember end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp.data", rsp.GetData()))

	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

// ShareFolderAddRecord .
func ShareFolderAddRecord(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.ShareFolderAddRecordReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Info("ShareFolderAddRecord start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.ShareFolderAddRecord(gtx, req)
	if err != nil {
		log.Logger.Error("ShareFolderAddRecord rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderAddRecord rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Info("ShareFolderAddRecord end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp.data", rsp.GetData()))
	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

// ShareFolderDeleteRecord .
func ShareFolderDeleteRecord(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.ShareFolderDeleteRecordReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Info("ShareFolderDeleteRecord start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))
	//log.Logger.Debug("ShareFolderDeleteRecord start debug", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.ShareFolderDeleteRecord(gtx, req)
	if err != nil {
		log.Logger.Error("ShareFolderDeleteRecord rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderDeleteRecord rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Info("ShareFolderDeleteRecord end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp.data", rsp.GetData()))
	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

func ShareFolderFolderList(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}
	obj := value.(*encode.Web3PasswordRequestBsonStruct)
	req := &pb.ShareFolderFolderListReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
	}
	log.Logger.Info("ShareFolderFolderList start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))
	//log.Logger.Debug("ShareFolderFolderList start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.ShareFolderFolderList(gtx, req)
	if err != nil {
		log.Logger.Error("ShareFolderFolderList rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderFolderList rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	//handle bson
	rspData := make([]*model.ShareFolder, 0)
	for _, item := range rsp.Data.GetList() {
		rspData = append(rspData, &model.ShareFolder{
			FolderId:       item.FolderId,
			FolderName:     item.FolderName,
			FolderOwner:    item.FolderOwner,
			FolderMnemonic: item.FolderMnemonic,
			Timestamp:      item.GetTimestamp(),
			FolderAuth:     item.FolderAuth,
		})
	}
	wrapData := model.ShareFolderListRsp{
		List: rspData,
	}
	bytes, _ := bson.Marshal(wrapData)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	log.Logger.Info("ShareFolderFolderList end", log.String("trace_id", ctx.GetString("trace_id")))
	//log.Logger.Debug("ShareFolderFolderList end debug", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp.data", rsp.GetData()))
}

func ShareFolderRecordList(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}
	obj := value.(*encode.Web3PasswordRequestBsonStruct)
	req := &pb.ShareFolderRecordListReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	log.Logger.Info("ShareFolderRecordList start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))
	log.Logger.Debug("ShareFolderRecordList start debug", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.ShareFolderRecordList(gtx, req)
	if err != nil {
		log.Logger.Error("ShareFolderRecordList rsp error", log.Error(err), log.String("trace_id", ctx.GetString("trace_id")))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderRecordList rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	//handle bson
	rspData := make([]*model.ShareFolderRecord, 0)
	for _, item := range rsp.Data.GetList() {
		record := &model.ShareFolderRecord{
			Id:         item.Id,
			RecordId:   item.RecordId,
			FolderId:   item.FolderId,
			RecordData: item.RecordData,
			OwnerAddr:  item.OwnerAddr,
		}
		rspData = append(rspData, record)
	}

	wrapData := model.ShareFolderRecordRsp{
		List: rspData,
	}
	bytes, _ := bson.Marshal(wrapData)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	log.Logger.Info("ShareFolderRecordList end", log.String("trace_id", ctx.GetString("trace_id")))
	//rspListBson, _ := bson.Marshal(rspData)
	//log.Logger.Debug("ShareFolderRecordList end debug", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", base64.StdEncoding.EncodeToString(rspListBson)))
}

func ShareFolderRecordListByRid(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}
	obj := value.(*encode.Web3PasswordRequestBsonStruct)
	req := &pb.ShareFolderRecordListByRidReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	log.Logger.Info("ShareFolderRecordListByRid start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))
	log.Logger.Debug("ShareFolderRecordListByRid start debug", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.ShareFolderRecordListByRid(gtx, req)
	if err != nil {
		log.Logger.Error("ShareFolderRecordListByRid rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderRecordListByRid rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	rspData := make([]*model.ShareFolderRecordByRid, 0)
	for _, item := range rsp.Data.GetList() {
		record := &model.ShareFolderRecordByRid{
			Id:             item.Id,
			RecordId:       item.RecordId,
			FolderId:       item.FolderId,
			RecordData:     item.RecordData,
			FolderMnemonic: item.FolderMnemonic,
		}
		rspData = append(rspData, record)
	}

	wrapData := model.ShareFolderRecordByRidRsp{
		List: rspData,
	}
	bytes, _ := bson.Marshal(wrapData)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	log.Logger.Info("ShareFolderRecordListByRid end", log.String("trace_id", ctx.GetString("trace_id")))
	//log.Logger.Debug("ShareFolderRecordListByRid end debug", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp.data", rsp.GetData()))
}

// ShareFolderMemberExit .
func ShareFolderMemberExit(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.ShareFolderMemberExitReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Info("ShareFolderMemberExit start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))
	log.Logger.Debug("ShareFolderMemberExit start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.ShareFolderMemberExit(gtx, req)
	if err != nil {
		log.Logger.Error("ShareFolderMemberExit rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}
	if rsp.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderMemberExit rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Info("ShareFolderMemberExit end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))

	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

// ShareFolderBatchUpdate .
func ShareFolderBatchUpdate(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.ShareFolderBatchUpdateReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Info("ShareFolderBatchUpdate start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))
	log.Logger.Debug("ShareFolderBatchUpdate start debug", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.ShareFolderBatchUpdate(gtx, req)
	if err != nil {
		log.Logger.Error("ShareFolderBatchUpdate rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderBatchUpdate rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Info("ShareFolderBatchUpdate end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))

	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

// ShareFolderMemberList .
func ShareFolderMemberList(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, "request error", emptyByte)
		ctx.Abort()
		return
	}
	obj := value.(*encode.Web3PasswordRequestBsonStruct)
	req := &pb.ShareFolderMemberListReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	log.Logger.Info("ShareFolderMemberList start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))
	log.Logger.Debug("ShareFolderMemberList start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.ShareFolderMemberList(gtx, req)
	if err != nil {
		log.Logger.Error("ShareFolderMemberList error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("ShareFolderMemberList rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	rspData := make([]*model.ShareFolderMemberInfo, 0)
	for _, item := range rsp.Data.GetList() {
		member := &model.ShareFolderMemberInfo{
			MemberAddr:     item.MemberAddr,
			MemberName:     item.MemberName,
			MemberSign:     item.MemberSign,
			FolderId:       item.FolderId,
			FolderMnemonic: item.FolderMnemonic,
		}
		rspData = append(rspData, member)
	}

	wrapData := model.ShareFolderMemberListRsp{
		List: rspData,
	}

	log.Logger.Info("ShareFolderMemberList end", log.String("trace_id", ctx.GetString("trace_id")))
	log.Logger.Debug("ShareFolderMemberList end debug", log.String("trace_id", ctx.GetString("trace_id")), log.Any("data", wrapData))
	bytes, _ := bson.Marshal(wrapData)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}
