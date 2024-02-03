/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/
package service

import (
	"github.com/web3password/satis/config"
	"github.com/web3password/satis/dao"
	pb "github.com/web3password/w3p-protobuf/user"
)

// Service .
type Service struct {
	*pb.UnimplementedUserServer
	dao dao.DAO
}

// NewService .
func NewService(conf *config.Config) *Service {
	service := &Service{
		dao: dao.NewDAO(conf),
	}

	return service
}
