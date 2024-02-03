/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/
package service

import (
	pb "github.com/web3password/w3p-protobuf/user"
)

// Stream .
func (s *Service) Stream(server pb.User_StreamServer) error {
	return s.dao.Stream(server)
}
