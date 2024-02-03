/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/
package main

import (
	"flag"
	"net"
	"time"

	"google.golang.org/grpc/credentials"

	"github.com/fvbock/endless"
	"github.com/web3password/jewel/tools"
	"github.com/web3password/satis/config"
	"github.com/web3password/satis/log"
	"github.com/web3password/satis/service"
	"github.com/web3password/satis/service/handlers"
	pb "github.com/web3password/w3p-protobuf/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var (
	confPath = flag.String("conf", "-conf ../config.yaml", "-conf config.yaml")
	kaep     = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: true,            // Allow pings even when there are no active streams
	}

	kasp = keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
		MaxConnectionAge:      0 * time.Second,  // If any connection is alive for more than 30 seconds, send a GOAWAY
		MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
		Time:                  60 * time.Second, // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout:               90 * time.Second, // Wait 1 second for the ping ack before assuming the connection is dead
	}
)

func main() {
	flag.Parse()
	go config.WatchConfig(*confPath)
	config.ParseConfig(*confPath)
	conf := config.GetConfig()
	listen, err := net.Listen(conf.GetServerProto(), conf.GetGRPCServerAddress())
	if err != nil {
		log.Fatalf("failed to listen err:%+v", err)
	}
	log.SetLogger()
	options := []grpc.ServerOption{
		grpc.KeepaliveEnforcementPolicy(kaep),
		grpc.KeepaliveParams(kasp),
		grpc.MaxRecvMsgSize(conf.MsgSize.File),
		grpc.MaxSendMsgSize(conf.MsgSize.File),
	}
	if conf.Server.EnableTLS {
		tlsConfig, err := tools.TLSServerConfig(conf.Tls.Ca, conf.Tls.ServerTls.Crt, conf.Tls.ServerTls.Key)
		if err != nil {
			panic(err)
		}
		options = append(options, grpc.Creds(credentials.NewTLS(tlsConfig)))
	}
	s := grpc.NewServer(options...)
	pb.RegisterUserServer(s, service.NewService(conf))
	log.Logger.Info("grpc server listening at", log.Any("port", listen.Addr()))
	go func() {
		if err = s.Serve(listen); err != nil {
			log.Logger.Error("failed to grpc serve err", log.Error(err))
		}
	}()
	handlers.Init(conf)
	err = endless.ListenAndServe(conf.GetHttpServerAddress(), service.Routers())
	if err != nil {
		log.Logger.Error("server run error", log.Error(err))
	}
}
