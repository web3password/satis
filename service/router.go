/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/
package service

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/web3password/jewel/encode"
	"github.com/web3password/satis/consts"
	"github.com/web3password/satis/log"
	"github.com/web3password/satis/middleware"
	"github.com/web3password/satis/model"
	"github.com/web3password/satis/service/handlers"
	"io"
	"net/http"
	"strings"
)

func Routers() (engine *gin.Engine) {
	router := gin.Default()
	router.NoRoute(Handle404)
	runningMode := handlers.GetRunningMode()
	if consts.RunningModeOfficial != runningMode {
		router.Use(middleware.Agent(runningMode))
	}
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	router.Use(cors.New(corsConfig))
	router.Use(gin.Recovery())
	router.Use(ParamsCheck())
	user := router.Group("/web3password")
	user.POST("/userRegister", handlers.Register)
	user.POST("/getPersonalSignAddress", handlers.GetPersonalSignAddress)
	user.POST("/getVipInfo", handlers.GetVIPInfo)
	user.POST("/userInfo", handlers.GetUserInfo)
	user.POST("/getLatestBlockTimestamp", handlers.GetLatestBlockTimestamp)
	user.POST("/checkTx", handlers.CheckTx)
	user.POST("/batchCheckTx", handlers.BatchCheckTx)
	user.POST("/addCredential", handlers.AddCredential)
	user.POST("/batchAddCredential", handlers.BatchAddCredential)
	user.POST("/getCredential", handlers.GetCredential)
	user.POST("/deleteCredential", handlers.DeleteCredential)
	user.POST("/batchDeleteCredential", handlers.BatchDeleteCredential)
	user.POST("/deleteAllCredential", handlers.DeleteAllCredential)
	user.POST("/getAllCredentialTimestamp", handlers.GetAllCredentialTimestamp)
	user.POST("/getCredentialList", handlers.GetCredentialList)
	user.POST("/storageStat", handlers.StorageStat)
	user.POST("/getVersionConfig", handlers.GetVersionConfig)

	admin := router.Group("/web3password/admin")
	admin.POST("/authorization", handlers.Authorization)
	admin.POST("/addMember", handlers.AdminAddMember)
	admin.POST("/batchImportMember", handlers.AdminBatchImportMember)
	admin.POST("/updateMember", handlers.AdminUpdateMember)
	admin.POST("/removeMember", handlers.AdminRemoveMember)
	admin.POST("/transferSuperAdmin", handlers.AdminTransferSuperAdmin)
	admin.POST("/getMemberList", handlers.AdminGetMemberList)
	admin.POST("/getOrgInfo", handlers.AdminGetOrgInfo)
	admin.POST("/updateOrgInfo", handlers.AdminUpdateOrgInfo)
	admin.POST("/operationHistory", handlers.AdminOperationHistory)
	admin.POST("/getAdminShareMnemonic", handlers.GetAdminMnemonic)

	storage := router.Group("/web3password/file")
	storage.POST("/upload", handlers.FileUpload)
	storage.POST("/uploadIocopy", handlers.FileUpload)
	storage.POST("/uploadBufio", handlers.FileUpload)
	storage.POST("/download", handlers.FileDownload)
	storage.POST("/attachment", handlers.FileAttachment)
	storage.POST("/report", handlers.FileReport)

	folder := router.Group("/web3password/sharefolder")
	folder.POST("/create", handlers.ShareFolderCreate)
	folder.POST("/update", handlers.ShareFolderUpdate)
	folder.POST("/destroy", handlers.ShareFolderDestroy)
	folder.POST("/addrecord", handlers.ShareFolderAddRecord)
	folder.POST("/deleterecord", handlers.ShareFolderDeleteRecord)
	folder.POST("/addmember", handlers.ShareFolderAddMember)
	folder.POST("/updatemember", handlers.ShareFolderUpdateMember)
	folder.POST("/memberlist", handlers.ShareFolderMemberList)
	folder.POST("/memberexit", handlers.ShareFolderMemberExit)
	folder.POST("/deletemember", handlers.ShareFolderDeleteMember)
	folder.POST("/batchUpdate", handlers.ShareFolderBatchUpdate)
	folder.POST("/folderlist", handlers.ShareFolderFolderList)
	folder.POST("/recordlist", handlers.ShareFolderRecordList)
	folder.POST("/recordlistbyrid", handlers.ShareFolderRecordListByRid)
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	vip := router.Group("/web3password/vip")
	vip.POST("/getConfig", handlers.VipGetConfig)
	vip.POST("/subscriptionList", handlers.SubscriptionList)
	vip.POST("/createOrder", handlers.VipCreateOrder)
	vip.POST("/checkOrder", handlers.VipCheckOrder)
	vip.POST("/apple/in-app-purchase/verifyReceipt", handlers.VipAppleVerifyReceipt)
	vip.POST("/register", handlers.AdminRegister)
	vip.POST("/paymentList", handlers.VipPaymentList)
	vip.POST("/discount", handlers.VipDiscount)
	vip.POST("/getOrderList", handlers.GetOrderList)
	vip.POST("/getVipIOSPromotionSign", handlers.GetVipIOSPromotionSign)
	vip.POST("/price", handlers.VipPrice)

	return router
}

func Handle404(c *gin.Context) {
	c.String(http.StatusNotFound, "not found")
	return
}

func ParamsCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		nonce := uuid.NewString()
		ctx.Set("trace_id", nonce)

		log.Logger.Info(ctx.Request.RequestURI,
			log.String("method", ctx.Request.Method),
			log.String("ip", ctx.ClientIP()),
			log.Any("header", ctx.Request.Header),
			log.String("user-agent", ctx.Request.UserAgent()),
			log.String("trace_id", nonce))

		var err error
		var bytes []byte
		if ctx.Request.RequestURI == "/web3password/file/upload" {
			bytes, err = ctx.GetRawData()
		} else if ctx.Request.RequestURI == "/web3password/file/uploadIocopy" {
			bytes, err = ReadHttpBodyIocopy(ctx)
		} else if ctx.Request.RequestURI == "/web3password/file/uploadBufio" {
			bytes, err = ReadHttpBodyBufio(ctx)
		} else {
			bytes, err = ctx.GetRawData()
		}

		if err != nil {
			log.Logger.Warn("read req error", log.Error(err), log.String("trace_id", nonce))
			handlers.Response(ctx, model.StatusParamsErr, "params error %s", []byte(""))
			ctx.Abort()
			return
		}

		limit := handlers.GetDefaultApiMaxSize()
		if strings.Contains(ctx.Request.RequestURI, "file") {
			limit = handlers.GetFileMaxSize()
		}

		if len(bytes) > limit {
			log.Logger.Warn("params check request body size is limited", log.String("trace_id", nonce), log.Any("size", len(bytes)), log.Any("limit", limit))
			handlers.Response(ctx, model.StatusParamsErr, "request body size is limited", []byte(""))
			ctx.Abort()
			return
		}

		request, err := encode.Web3PasswordRequestBsonDecode(bytes)
		if err != nil {
			log.Logger.Warn("params check req decode error", log.Error(err))
			handlers.Response(ctx, model.StatusParamsErr, "params error", []byte(""))
			ctx.Abort()
			return
		}

		log.Logger.Info("satis handler request", log.String("params", request.ParamsStr), log.String("trace_id", nonce), log.Any("body-length", len(bytes)))
		ctx.Set("request", request)
		ctx.Next()
	}
}

func ReadHttpBodyIocopy(ctx *gin.Context) ([]byte, error) {
	// Create a buffer to store the request body
	var buf bytes.Buffer

	// Copy the request body to the buffer
	_, err := io.Copy(&buf, ctx.Request.Body)
	if err != nil {
		return []byte{}, errors.New("failed to copy request body")
	}
	return buf.Bytes(), nil
}

func ReadHttpBodyBufio(ctx *gin.Context) ([]byte, error) {
	// Create a bufio.Reader from the request's body
	bodyReader := bufio.NewReader(ctx.Request.Body)

	// Create a buffer to store the request body
	var requestBody []byte
	buffer := make([]byte, 1024)

	for {
		n, err := bodyReader.Read(buffer)
		if err != nil {
			break
		}
		requestBody = append(requestBody, buffer[:n]...)
	}
	return requestBody, nil
}
