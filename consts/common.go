/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/
package consts

const (
	RunningModeOfficial = "official"
	RunningModeAudit    = "audit"
	RunningModeLocal    = "local"
)

const (
	PLATFORM_PAYPAL = "1"
	PLATFORM_STRIPE = "2"
	PLATFORM_GOOGLE = "3"
	PLATFORM_APPLE  = "4"
)

type AppType string

const (
	AppWeb     AppType = "1"
	AppIos     AppType = "2"
	AppAndroid AppType = "3"
)

func (a AppType) Check() bool {
	if a == AppAndroid || a == AppIos || a == AppWeb {
		return true
	}
	return false
}
