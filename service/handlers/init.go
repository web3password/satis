/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/
package handlers

import (
	"context"
	"gopkg.in/mgo.v2/bson"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/web3password/jewel/tools"
	"github.com/web3password/satis/config"
	"github.com/web3password/satis/log"
	pb "github.com/web3password/w3p-protobuf/user"
)

var (
	userClient pb.UserClient
	emptyByte  []byte
)

func Init(conf *config.Config) {
	type empty struct {
	}
	emptyByte, _ = bson.Marshal(empty{})
	userClient = NewUserClient(conf)
}

func NewUserClient(conf *config.Config) pb.UserClient {
	options := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(conf.MsgSize.File)),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(conf.MsgSize.File)),
	}
	if conf.Server.EnableTLS {
		tlsConfig, err := tools.TLSClientConfig(conf.Tls.Ca, conf.Tls.ClientTls.Crt, conf.Tls.ClientTls.Key)
		if err != nil {
			panic(err)
		}
		tlsConfig.ServerName = conf.Server.TLSDomain
		options = append(options, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	} else {
		options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	conn, err := grpc.DialContext(
		context.Background(),
		conf.GetGRPCServerAddress(),
		options...,
	)
	if err != nil {
		log.Fatalf("failed to dial grpc server:%+v", err)
		panic(err)
	}
	userGrpcClient := pb.NewUserClient(conn)
	return userGrpcClient
}

func GetUserClient() pb.UserClient {
	return userClient
}

func GetDefaultApiMaxSize() int {
	return config.GetConfig().MsgSize.Api
}

func GetFileMaxSize() int {
	return config.GetConfig().MsgSize.File
}

func GetRunningMode() string {
	return config.GetConfig().RunningMode
}

func IsHttpWithTraceID() bool {
	return config.GetConfig().HttpServer.WithTraceID
}
