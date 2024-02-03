/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/

package middleware

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/web3password/jewel/encode"
	"github.com/web3password/satis/config"
	"github.com/web3password/satis/consts"
	"github.com/web3password/satis/log"
	"github.com/web3password/satis/model"
	"github.com/web3password/satis/service/handlers"
	"io"
	"math/big"
	"net/http"
	"strings"
)

func Agent(runningMode string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyBytes, err := io.ReadAll(ctx.Request.Body)
		ctx.Request.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		if runningMode == consts.RunningModeAudit {
			log.Logger.Info(ctx.Request.RequestURI, log.String("audit_log", base64.StdEncoding.EncodeToString(bodyBytes)))
		}
		if err != nil {
			handlers.Response(ctx, model.StatusParamsErr, "params error", []byte(""))
			ctx.Abort()
			return
		}
		personalWhiteList := config.GetConfig().PersonalWhiteList
		orgWhiteList := config.GetConfig().OrgWhiteList
		params := model.CommonParams{}
		if len(personalWhiteList) > 0 || len(orgWhiteList) > 0 {
			request, _ := encode.Web3PasswordRequestBsonDecode(bodyBytes)
			if err := jsoniter.UnmarshalFromString(request.ParamsStr, &params); err != nil {
				handlers.Response(ctx, model.StatusParamsErr, "params error", []byte(""))
				ctx.Abort()
				return
			}
		}
		if runningMode == consts.RunningModeLocal &&
			(strings.Contains(ctx.Request.RequestURI, "/file") ||
				strings.Contains(ctx.Request.RequestURI, "/sharefolder") ||
				strings.Contains(ctx.Request.RequestURI, "/getLatestBlockTimestamp") ||
				strings.Contains(ctx.Request.RequestURI, "/checkTx") ||
				strings.Contains(ctx.Request.RequestURI, "/addCredential") ||
				strings.Contains(ctx.Request.RequestURI, "/deleteCredential") ||
				strings.Contains(ctx.Request.RequestURI, "/getCredential") ||
				strings.Contains(ctx.Request.RequestURI, "/deleteAllCredential") ||
				strings.Contains(ctx.Request.RequestURI, "/getAllCredentialTimestamp") ||
				strings.Contains(ctx.Request.RequestURI, "/getCredentialList")) {
			if (len(personalWhiteList) > 0 || len(orgWhiteList) > 0) && !strings.Contains(personalWhiteList, params.Address) && !strings.Contains(orgWhiteList, params.OrgId) {
				handlers.Response(ctx, model.StatusForbiddenErr, "current node forbid error", []byte(""))
				ctx.Abort()
				return
			}
			ctx.Next()
			return
		}
		// admin
		if runningMode == consts.RunningModeLocal && (strings.Contains(ctx.Request.RequestURI, "/admin")) {
			if strings.Contains(ctx.Request.RequestURI, "/register") ||
				strings.Contains(ctx.Request.RequestURI, "/addMember") ||
				strings.Contains(ctx.Request.RequestURI, "/removeMember") ||
				strings.Contains(ctx.Request.RequestURI, "/transferSuperAdmin") {
				defer ReportOfficial(ctx)
			}
			ctx.Next()
			return
		}
		officialURL := randomOfficeUrl() + ctx.Request.RequestURI
		res := DoRequestBinary(officialURL, bodyBytes, "POST")
		handlers.Response(ctx, res.Code, res.Msg, res.Data)
		ctx.Abort()
		return
	}
}

func randomOfficeUrl() string {
	domains := config.GetConfig().OfficialDomains
	randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(domains))))
	if err != nil {
		return domains[0]
	}
	return domains[randomIndex.Int64()]
}

func ReportOfficial(ctx *gin.Context) {
	officialURL := randomOfficeUrl() + ctx.Request.RequestURI
	log.Logger.Info("default official URL:", log.String("official", officialURL))
	bytes, err := io.ReadAll(ctx.Request.Body)
	log.Infof("officialURL: %+v", officialURL)
	if err != nil {
		ctx.Abort()
		return
	}
	DoRequestBinary(officialURL, bytes, "POST")
	return
}

func DoRequestBinary(url string, body []byte, method string) *encode.Web3PasswordResponseBsonStruct {
	if method == "" {
		method = "POST"
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/octet-stream")
	if err != nil {
		//log.Logger.Error("NewRequest url:%s, err:%+v\n", url, err)
		return nil
	}
	var client = &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		//log.Logger.Error("request failed, url:%s err:%+v\n", url, err)
		return nil
	}
	defer func() { _ = rsp.Body.Close() }()
	statusCode := rsp.StatusCode
	if statusCode != 200 {
		log.Infof("request failed, statusCode:%d url:%s\n", statusCode, url)
		return nil
	}
	r, err := io.ReadAll(rsp.Body)
	if err != nil {
		//log.Logger.Error("read body err:%+v\n", err)
		return nil
	}
	decode, err := encode.Web3PasswordResponseBsonDecode(r)
	if err != nil {
		//log.Logger.Error("response decode failed, url: %s, err: %s\n", url, err)
		return nil
	}
	return decode
}
