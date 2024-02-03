/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/
package dao

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/web3password/satis/log"
	"github.com/web3password/satis/model"
	pb "github.com/web3password/w3p-protobuf/user"
	"google.golang.org/grpc/metadata"
)

// map["addr"] = map[key]server
var streamMap = sync.Map{}
var requestChanLength = 20000

func (d *dao) DeleteItem(group, nodeID, connKey string) map[string]string {
	d.lock.Lock()
	defer d.lock.Unlock()
	// clear
	if _, ok := d.clients[group]; ok {
		if _, ok := d.clients[group][nodeID]; ok {
			v, _ := streamMap.Load(nodeID)
			oldConnKey := v.(string)
			if oldConnKey == "" || oldConnKey == connKey {
				delete(d.clients[group], nodeID)
			} else {
				return d.clients[group]
			}
		}
	}
	log.Logger.Warn("stream connect delete group client", log.String("group", group), log.String("connKey", connKey), log.String("nodeID", nodeID), log.Any("groups", d.clients))
	return d.clients[group]
}

func (d *dao) UpdateGroup(group, clientID, connKey string) {
	d.lock.Lock()
	defer d.lock.Unlock()
	if _, ok := d.clients[group]; ok {
		d.clients[group][clientID] = clientID
		streamMap.Store(clientID, connKey)
		log.Logger.Warn("stream connect add group client", log.String("group", group), log.String("connKey", connKey), log.String("client_id", clientID), log.Any("groups", d.clients))
		return
	}

	streamMap.Store(clientID, connKey)
	d.clients[group] = map[string]string{
		clientID: clientID,
	}
	log.Logger.Warn("stream connect init group client", log.String("group", group), log.String("connKey", connKey), log.String("client_id", clientID), log.Any("groups", d.clients))
	return
}

func (d *dao) Stream(server pb.User_StreamServer) error {
	md, ok := metadata.FromIncomingContext(server.Context())
	nodeID := "default_node"
	nodeGroup := "default_group"
	nodeConn := "default_conn"

	if ok {
		mds := md.Get("client_id")
		if len(mds) > 0 {
			nodeID = mds[0]
		}
		groups := md.Get("group")
		if len(groups) > 0 {
			nodeGroup = groups[0]
		}
		conns := md.Get("conn")
		if len(conns) > 0 {
			nodeConn = conns[0]
		}
	}

	streamMap.Store(nodeID, nodeConn)
	log.Logger.Info(fmt.Sprintf("stream connect start nodeGroup:%s nodeID:%s connKey:%s", nodeGroup, nodeID, nodeConn))
	if _, ok := d.requestChan[nodeID]; !ok {
		d.requestChan[nodeID] = make(chan *pb.StreamRsp, requestChanLength)
		log.Logger.Info(fmt.Sprintf("stream connect start init requestChan success nodeGroup:%s nodeID:%s connKey:%s", nodeGroup, nodeID, nodeConn))
	}
	d.UpdateGroup(nodeGroup, nodeID, nodeConn)
	log.Logger.Info(fmt.Sprintf("stream connect success nodeGroup:%s nodeID:%s connKey:%s", nodeGroup, nodeID, nodeConn))
	clientClose := make(chan any, requestChanLength)

	go func() {
		defer func() {
			clientClose <- nil
			log.Logger.Warn("server send msg defer close", log.Any("connkey", nodeConn), log.String("node", nodeID))
		}()
		requestChan := d.requestChan[nodeID]
		for r := range requestChan {
			traceId := r.GetTraceId()
			log.Logger.Debug("server send msg start", log.Any("connkey", nodeConn), log.Any("send length", len(d.requestChan[nodeID])), log.String("cmd", r.GetCmd()), log.String("req", r.GetParams()), log.String("node", nodeID), log.String("trace_id", traceId))
			if err := server.Send(r); err != nil {
				log.Logger.Error("server send msg error", log.Any("connkey", nodeConn), log.String("cmd", r.GetCmd()), log.Any("req", r.GetParams()), log.Error(err), log.String("node", nodeID), log.String("trace_id", traceId))
				break
			}
			log.Logger.Debug("server send msg success", log.Any("connkey", nodeConn), log.String("node", nodeID), log.String("trace_id", traceId))
		}
	}()

	go func() {
		defer func() {
			req := &pb.StreamRsp{
				Cmd:       model.CMDPong,
				RequestId: 0,
				Token:     d.conf.Node.Token,
				Signature: "",
				Params:    "ping",
				TraceId:   uuid.NewString(),
			}

			d.requestChan[nodeID] <- req

			clientClose <- nil
			log.Logger.Warn("server recv msg defer close", log.Any("connkey", nodeConn), log.String("node", nodeID))
		}()
		for {
			recvId := uuid.NewString()
			log.Logger.Debug("server recv msg start", log.Any("connkey", nodeConn), log.String("node", nodeID), log.String("recvId", recvId))
			res, err := server.Recv()
			if err != nil {
				log.Logger.Error("server recv msg error", log.Any("connkey", nodeConn), log.Error(err), log.String("node", nodeID), log.String("recvId", recvId))
				break
			}
			traceId := res.GetTraceId()
			log.Logger.Debug("server recv msg data", log.Any("connkey", nodeConn), log.Any("res", res.GetCmd()), log.String("node", nodeID), log.String("recvId", recvId), log.String("trace_id", traceId))
			if res.GetCmd() == model.CMDFileDownload || res.GetCmd() == model.CMDFileAttachment {
				log.Logger.Debug("server recv msg 1", log.Any("connkey", nodeConn), log.Any("res", res.GetCmd()), log.String("node", nodeID), log.String("trace_id", traceId))
			} else if res.GetCmd() == model.CMDGracefulRestartSignal {
				log.Logger.Info("server recv msg data", log.Any("connkey", nodeConn), log.Any("res", res), log.String("node", nodeID), log.String("recvId", recvId), log.String("trace_id", traceId))
				d.DeleteItem(nodeGroup, nodeID, nodeConn)
				//delete(d.requestChan, nodeID)
			} else {
				//log.Logger.Debug("server recv msg 2", log.Any("connkey", nodeConn), log.Any("res", res.GetParams()), log.String("node", nodeID), log.String("trace_id", traceId))
			}

			/*if res.GetToken() != config.GetConfig().Node.Token {
				log.Logger.Error("server RecvMsg invalid token", log.Error(err), log.String("node", client_id))
				break
			}*/
			requestID := res.GetRequestId()
			log.Logger.Debug("server recv msg success for requestid", log.Any("connkey", nodeConn), log.String("nodeID", nodeID), log.String("recvId", recvId), log.String("trace_id", traceId), log.Int64("requestID", requestID))

			if requestID == 0 {
				continue
			}
			log.Logger.Debug("loadStreamResponseWaitChan will go...", log.Any("connkey", nodeConn), log.String("nodeID", nodeID), log.Int64("requestID", requestID), log.String("recvId", recvId), log.String("trace_id", traceId))

			ch, err := d.loadStreamResponseWaitChan(requestID)
			if err != nil {
				log.Logger.Debug("loadStreamResponseWaitChan continue...", log.Any("connkey", nodeConn), log.String("nodeID", nodeID), log.Int64("requestID", requestID), log.String("recvId", recvId), log.String("trace_id", traceId))
				continue
			}
			log.Logger.Debug("loadStreamResponseWaitChan channel start...", log.Any("connkey", nodeConn), log.String("nodeID", nodeID), log.Int64("requestID", requestID), log.String("recvId", recvId), log.String("trace_id", traceId))

			ch <- res
			log.Logger.Debug("loadStreamResponseWaitChan channel end...", log.Any("connkey", nodeConn), log.String("nodeID", nodeID), log.Int64("requestID", requestID), log.String("recvId", recvId), log.String("trace_id", traceId))

		}
	}()
	for {
		select {
		case <-clientClose:
			log.Logger.Warn("stream connect closed start", log.Any("connkey", nodeConn), log.String("nodeID", nodeID), log.Any("groups", d.clients))
			d.DeleteItem(nodeGroup, nodeID, nodeConn)
			log.Logger.Warn("stream connect closed success", log.Any("connkey", nodeConn), log.String("nodeID", nodeID), log.Any("groups", d.clients))
			return fmt.Errorf("close client")
		}
	}
	return nil
}

func (d *dao) heartBeat() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			for _, ch := range d.requestChan {
				ch <- &pb.StreamRsp{
					Cmd:       model.CMDPong,
					RequestId: 0,
					Token:     d.conf.Node.Token,
					Signature: "",
					Params:    "ping",
					TraceId:   uuid.NewString(),
				}
			}
		}
	}
}

func (d *dao) addStreamRequest(req *pb.StreamRsp, group string) error {
	if _, ok := d.clients[group]; !ok {
		return fmt.Errorf(group + " not found")
	}

	if len(d.clients[group]) == 0 {
		return fmt.Errorf(group + " there is no client")
	}

	isSuccess := false
	for i := 0; i < len(d.clients[group]); i++ {
		var nodeIDs []string
		for _, v := range d.clients[group] {
			nodeIDs = append(nodeIDs, v)
		}
		if len(nodeIDs) == 0 {
			log.Logger.Error(fmt.Sprintf("group:%s has empty nodeIDs", group))
			break
		}

		index := 0
		if len(nodeIDs) != 1 {
			r := rand.New(rand.NewSource(time.Now().UnixMilli()))
			index = r.Intn(len(nodeIDs) - 1)
		}

		nodeID := nodeIDs[index]
		if _, ok := d.requestChan[nodeID]; ok {
			d.requestChan[nodeID] <- req
			isSuccess = true
			break
		}
		log.Logger.Warn(fmt.Sprintf(group + " there is no client " + nodeID))
		time.Sleep(10 * time.Millisecond)
	}
	if !isSuccess {
		log.Logger.Error(fmt.Sprintf("addStreamRequest fail clients:%+v group:%s req:%+v", d.clients, group, req))
		return fmt.Errorf("addStreamRequest fail")
	}

	return nil
}

func (d *dao) addStreamResponseWaitChan(requestID int64) chan *pb.StreamReq {
	waitChan := make(chan *pb.StreamReq, 1)
	d.responseWait.Store(requestID, waitChan)
	return waitChan
}

func (d *dao) loadStreamResponseWaitChan(requestID int64) (chan *pb.StreamReq, error) {
	v, ok := d.responseWait.Load(requestID)
	if !ok {
		log.Logger.Warn("server requestID not found", log.Int64("requestID", requestID))
		return nil, fmt.Errorf("requestID not found")
	}
	ch, ok := v.(chan *pb.StreamReq)
	if !ok {
		log.Logger.Error("server waitChan type error", log.Int64("requestID", requestID))
		return nil, fmt.Errorf("waitChan type error")
	}
	return ch, nil
}

func (d *dao) delStreamResponseWaitChan(requestID int64) {
	d.responseWait.Delete(requestID)
}
