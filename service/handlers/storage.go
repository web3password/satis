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

func FileUpload(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.FileUploadReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}
	log.Logger.Info("FileUpload start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr), log.Int64("attach_length", int64(len(obj.AppendData))))

	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.FileUpload(gtx, req)
	if err != nil {
		log.Logger.Error("FileUpload rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("FileUpload rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	cid := &model.FileUploadRsp{
		Cid: rsp.GetData().GetCid(),
	}

	log.Logger.Info("FileUpload end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp.GetData()))
	bytes, _ := bson.Marshal(cid)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

func FileDownload(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.FileDownloadReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
	}
	log.Logger.Info("FileDownload start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.FileDownload(gtx, req)
	if err != nil {
		log.Logger.Error("FileDownload rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("FileDownload rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	fileDownloadRsp := &model.FileDownloadRsp{
		Files: nil,
	}
	fileDownloadItem := model.FileDownLoadItemRsp{
		Cid:     rsp.GetData().GetCid(),
		Content: rsp.GetData().GetContent(),
	}
	fileDownloadRsp.Files = append(fileDownloadRsp.Files, fileDownloadItem)
	log.Logger.Info("FileDownload end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("response cid", rsp.GetData().GetCid()),
		log.Any("response data length in byte", len(rsp.GetData().GetContent())),
		log.Any("download params", obj.ParamsStr),
	)
	bytes, _ := bson.Marshal(fileDownloadRsp)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

func FileAttachment(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	obj := value.(*encode.Web3PasswordRequestBsonStruct)

	req := &pb.FileAttachmentReq{
		Signature: obj.SignatureStr,
		Params:    obj.ParamsStr,
		Data:      obj.AppendData,
	}

	log.Logger.Info("FileAttachment start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", obj.ParamsStr))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.FileAttachment(gtx, req)
	if err != nil {
		log.Logger.Error("FileAttachment rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err), log.Any("params", obj.ParamsStr))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("FileAttachment rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	fileAttachmentRsp := &model.FileAttachmentRsp{
		Files: nil,
	}
	fileAttachmentItem := model.FileAttachmentItem{
		Success: rsp.GetData().GetSuccess(),
		Message: rsp.GetData().GetMessage(),
		Cid:     rsp.GetData().GetCid(),
		Content: rsp.GetData().GetContent(),
	}
	fileAttachmentRsp.Files = append(fileAttachmentRsp.Files, fileAttachmentItem)
	log.Logger.Info("FileAttachment end", log.String("trace_id", ctx.GetString("trace_id")),
		log.Any("response attach length", len(rsp.GetData().GetContent())),
		log.Any("download params", obj.ParamsStr),
	)
	bytes, _ := bson.Marshal(fileAttachmentRsp)
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), bytes)
}

func FileReport(ctx *gin.Context) {
	value, ok := ctx.Get("request")
	if !ok {
		Response(ctx, model.StatusParamsErr, model.MsgParamsErr, emptyByte)
		ctx.Abort()
		return
	}

	req := value.(*encode.Web3PasswordRequestBsonStruct)

	rpcReq := &pb.FileReportReq{
		Signature: req.SignatureStr,
		Params:    req.ParamsStr,
		Data:      req.AppendData,
	}

	log.Logger.Info("FileReport start", log.String("trace_id", ctx.GetString("trace_id")), log.Any("params", rpcReq.GetParams()))
	gtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"trace_id", ctx.GetString("trace_id"),
	))
	rsp, err := userClient.FileReport(gtx, rpcReq)
	if err != nil {
		log.Logger.Error("FileReport rsp error", log.String("trace_id", ctx.GetString("trace_id")), log.Error(err))
		Response(ctx, model.StatusServiceCheckErr, model.MsgSystemErr, emptyByte)
		return
	}

	if rsp.Code != model.StatusOK {
		log.Logger.Warn("FileReport rsp warning", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
		Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
		return
	}

	log.Logger.Info("FileReport end", log.String("trace_id", ctx.GetString("trace_id")), log.Any("rsp", rsp))
	Response(ctx, int(rsp.GetCode()), rsp.GetMsg(), emptyByte)
	return
}
