/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/
package handlers

import (
	"context"
	"github.com/web3password/satis/log"
	"google.golang.org/grpc/metadata"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/web3password/jewel/encode"
	"github.com/web3password/satis/model"
	pb "github.com/web3password/w3p-protobuf/user"
	"gopkg.in/mgo.v2/bson"
)

func GetLatestBlockTimestamp(ctx *gin.Context) {
	type GetLatestBlockTimestampRsp struct {
		Timestamp int32 `bson:"timestamp"`
	}

	data := &GetLatestBlockTimestampRsp{
		Timestamp: int32(time.Now().Unix()),
	}
	log.Logger.Debug("GetLatestBlockTimestamp success", log.String("trace_id", ctx.GetString("trace_id")), log.Any("data", data))
	bytes, _ := bson.Marshal(data)
	Response(ctx, model.StatusOK, model.MsgOK, bytes)
}

func CheckTx(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.CheckTxReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
	}

	log.Logger.Debug("checktx start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.CheckTx(gtx, req)
	if err != nil {
		log.Logger.Error("checktx rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err), log.Any("params", obj.ParamsStr))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("checktx rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	is_success := model.NewCheckTxRsp{IsSuccess: rsp.GetData().GetIsSuccess()}
	log.Logger.Debug("checktx end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
	bytes, _ := bson.Marshal(is_success)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	return
}

func BatchCheckTx(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.BatchCheckTxReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Debug("BatchCheckTx start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.BatchCheckTx(gtx, req)
	if err != nil {
		log.Logger.Error("BatchCheckTx error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err), log.Any("params", obj.ParamsStr))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("BatchCheckTx rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Debug("BatchCheckTx end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))

	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	return
}

func AddCredential(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.AddCredentialReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Debug("AddCredential start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req.GetParams()))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.AddCredential(gtx, req)
	if err != nil {
		log.Logger.Error("AddCredential error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err), log.Any("req", req.GetParams()))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("AddCredential rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	hash := model.AddCredentialRsp{Hash: rsp.GetData().GetTxHash()}

	log.Logger.Debug("AddCredential end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
	bytes, _ := bson.Marshal(hash)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	return
}

func BatchAddCredential(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.BatchAddCredentialReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	log.Logger.Debug("BatchAddCredential start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req.GetParams()))

	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.BatchAddCredential(gtx, req)
	if err != nil {
		log.Logger.Error("BatchAddCredential error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err), log.Any("req", req.GetParams()))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("BatchAddCredential rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Debug("BatchAddCredential end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	return
}

func GetCredential(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.GetCredentialReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
	}
	log.Logger.Debug("GetCredential start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req.GetParams()))

	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.GetCredential(gtx, req)
	if err != nil {
		log.Logger.Error("GetCredential error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err), log.Any("req", req.GetParams()))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("GetCredential rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	rspData := model.GetCredentialRsp{}
	rspData.Id = rsp.GetData().Id
	rspData.OpTimestamp = rsp.GetData().OpTimestamp
	rspData.Credential = rsp.GetData().Credential
	bytes, _ := bson.Marshal(rspData)

	log.Logger.Debug("GetCredential end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	return
}
func DeleteCredential(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.DeleteCredentialReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Debug("DeleteCredential start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req.GetParams()))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.DeleteCredential(gtx, req)
	if err != nil {
		log.Logger.Error("DeleteCredential error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err), log.Any("req", req.GetParams()))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("DeleteCredential rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	hash := model.AddCredentialRsp{Hash: rsp.GetData().GetTxHash()}
	bytes, _ := bson.Marshal(hash)

	log.Logger.Debug("DeleteCredential end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	return
}

func BatchDeleteCredential(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.BatchDeleteCredentialReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	log.Logger.Debug("BatchDeleteCredential start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req.GetParams()))

	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.BatchDeleteCredential(gtx, req)
	if err != nil {
		log.Logger.Error("BatchDeleteCredential error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err), log.Any("req", req.GetParams()))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("BatchDeleteCredential rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Debug("BatchDeleteCredential end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
	bytes, _ := bson.Marshal(rsp.GetData())
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	return
}

func DeleteAllCredential(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.DeleteAllCredentialReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
	}

	log.Logger.Debug("DeleteAllCredential start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req.GetParams()))
	empty := make([]byte, 0)
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.DeleteAllCredential(gtx, req)
	if err != nil {
		log.Logger.Error("DeleteAllCredential error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err), log.Any("req", req.GetParams()))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), empty)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("DeleteAllCredential rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Debug("DeleteAllCredential end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
	data := model.EmptyRsp{}
	bytes, _ := bson.Marshal(data)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	return
}

func GetAllCredentialTimestamp(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.GetAllCredentialTimestampReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
	}
	log.Logger.Debug("GetAllCredentialTimestamp start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req.GetParams()))
	empty := make([]byte, 0)
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.GetAllCredentialTimestamp(gtx, req)
	if err != nil {
		log.Logger.Error("GetAllCredentialTimestamp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), empty)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("GetAllCredentialTimestamp rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	rspData := make([]*model.GetAllCredentialTimestampRsp, 0)
	for _, item := range rsp.GetData() {
		tmpData := &model.GetAllCredentialTimestampRsp{
			Id:          item.Id,
			OpTimestamp: item.OpTimestamp,
		}
		rspData = append(rspData, tmpData)
	}
	wrapData := model.GetAllCredentialTimestampRspData{
		List: rspData,
	}

	bytes, _ := bson.Marshal(wrapData)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	//log.Logger.Debug("GetAllCredentialTimestamp end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
	log.Logger.Debug("GetAllCredentialTimestamp end", log.String("trace_id", ctx.GetString("trace_id")))
	return
}

func GetCredentialList(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.GetCredentialListReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
	}
	log.Logger.Debug("GetCredentialList start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("req", req.GetParams()))
	empty := make([]byte, 0)
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.GetCredentialList(gtx, req)
	if err != nil {
		log.Logger.Error("GetCredentialList error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err), log.Any("req", req.GetParams()))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), empty)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("GetCredentialList rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	rspData := make([]*model.GetCredentialRsp, 0)
	for _, item := range rsp.GetData() {
		rspData = append(rspData, &model.GetCredentialRsp{
			Id:          item.Id,
			OpTimestamp: item.OpTimestamp,
			Credential:  item.Credential,
		})
	}
	log.Logger.Debug("GetCredentialList end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("count", len(rsp.GetData())))
	wrapData := model.GetCredentialRspData{
		List: rspData,
	}
	bytes, _ := bson.Marshal(wrapData)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
	return
}
