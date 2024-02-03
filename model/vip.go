/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/

package model

import (
	"fmt"
	"github.com/web3password/jewel/tools"
	"github.com/web3password/satis/consts"
	"github.com/web3password/satis/util"
	"time"
)

type VipGetConfigParams struct {
	Address   string `json:"addr"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Token     string `json:"token"`
}

func (v VipGetConfigParams) Check() (string, bool) {
	if !tools.IsValidAddress(v.Address) {
		return fmt.Sprintf("invalid params address(%s)", v.Address), false
	}

	if v.Token != VipGetConfigToken {
		return fmt.Sprintf("invalid params token(%s)", v.Token), false
	}

	if !util.CheckTimestamp(v.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), v.Timestamp)
		return msg, false
	}

	return "", true
}

type VipSubscriptionListParams struct {
	Address   string         `json:"addr"`
	Timestamp int64          `json:"timestamp"`
	Nonce     string         `json:"nonce"`
	Token     string         `json:"token"`
	OrgId     string         `json:"org_id"`
	App       consts.AppType `json:"app"`
	Auth      string         `json:"personal_auth"`
}

func (v VipSubscriptionListParams) Check() (string, bool) {
	if !tools.IsValidAddress(v.Address) {
		return "invalid address " + v.Address, false
	}

	if v.Token != VipSubscriptionListToken {
		return "invalid token" + v.Token, false
	}

	if !util.CheckTimestamp(v.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), v.Timestamp)
		return msg, false
	}

	if !v.App.Check() {
		msg := fmt.Sprintf("invalid request params app=%s", v.App)
		return msg, false
	}

	return "", true
}

type VipSubscriptionListRspDataItem struct {
	Version             string `json:"version" bson:"version"`
	VipName             string `json:"vip_name" bson:"vip_name"`
	OrgID               string `json:"org_id" bson:"org_id"`
	OrgName             []byte `json:"org_name" bson:"org_name"`
	MemberShareMnemonic []byte `json:"member_share_mnemonic" bson:"member_share_mnemonic"`
	IsRenewing          int32  `json:"renewing" bson:"renewing"`
	VipType             string `json:"vip_type" bson:"vip_type"`
	StartTime           int32  `json:"start_time" bson:"start_time"`
	ExpiredTime         int32  `json:"expired_time" bson:"expired_time"`
	ActiveUser          int32  `json:"active_user" bson:"active_user"`
	BuyTime             int32  `json:"buy_time" bson:"buy_time"`
}

type VipSubscriptionListRspData struct {
	MyVip     []*VipSubscriptionListRspDataItem `json:"my_vip" bson:"my_vip"`
	JoinedVip []*VipSubscriptionListRspDataItem `json:"joined_vip" bson:"joined_vip"`
}

type VipGetConfigRspData struct {
	Paypal struct {
		ClientId string `json:"client_id" bson:"client_id"`
	} `json:"paypal" bson:"paypal"`
}

type VipPaymentListParams struct {
	Address   string         `json:"addr"`
	Timestamp int64          `json:"timestamp"`
	Nonce     string         `json:"nonce"`
	Token     string         `json:"token"`
	App       consts.AppType `json:"app"`
	Version   string         `json:"version"`
	Auth      string         `json:"personal_auth"`
}

func (v VipPaymentListParams) Check() (string, bool) {
	if !tools.IsValidAddress(v.Address) {
		return "invalid address " + v.Address, false
	}

	if v.Token != VipPaymentListToken {
		return "invalid token" + v.Token, false
	}

	if !util.CheckTimestamp(v.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), v.Timestamp)
		return msg, false
	}

	if !v.App.Check() {
		msg := fmt.Sprintf("invalid request params app=%s", v.App)
		return msg, false
	}

	return "", true
}

type VipCreteOrderRspData struct {
	OrderId         string `json:"order_id" bson:"order_id"`
	OutOrderId      string `json:"out_order_id" bson:"out_order_id"`
	PaymentVoucher  string `json:"payment_voucher" bson:"payment_voucher"`
	PaymentRedirect string `json:"payment_redirect" bson:"payment_redirect"`
}

type VipCreateOrderParams struct {
	Address   string         `json:"addr"`
	Timestamp int64          `json:"timestamp"`
	Nonce     string         `json:"nonce"`
	Token     string         `json:"token"`
	App       consts.AppType `json:"app"`
	Auth      string         `json:"personal_auth"`
}

func (v VipCreateOrderParams) Check() (string, bool) {
	if !tools.IsValidAddress(v.Address) {
		return "invalid address " + v.Address, false
	}

	if v.Token != VipCreateOrderToken && v.Token != VipPrice {
		return "invalid token " + v.Token, false
	}

	if !util.CheckTimestamp(v.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), v.Timestamp)
		return msg, false
	}

	if !v.App.Check() {
		msg := fmt.Sprintf("invalid request params app=%s", v.App)
		return msg, false
	}

	return "", true
}

type VipPaymentListRspData struct {
	IncludeUser string     `json:"include_user" bson:"include_user"`
	ShareUser   string     `json:"share_user" bson:"share_user"`
	ProductId   string     `json:"product_id" bson:"product_id"`
	Extend      string     `json:"extend" bson:"extend"`
	SkuList     []*SkuItem `json:"sku_list" bson:"sku_list"`
}

type SkuItem struct {
	ValidityPeriod    string `json:"validity_period" bson:"validity_period"`
	PriceDesc         string `json:"price_desc" bson:"price_desc"`
	Price             int32  `json:"price" bson:"price"`
	Currency          string `json:"currency" bson:"currency"`
	Recommend         string `json:"recommend" bson:"recommend"`
	ProductId         string `json:"product_id" bson:"product_id"`
	SkuId             string `json:"sku_id" bson:"sku_id"`
	SkuType           string `json:"sku_type" bson:"sku_type"`
	GoodId            int64  `json:"good_id" bson:"good_id"`
	OneMonthPrice     int32  `json:"one_month_price" bson:"one_month_price"`
	OneUserPrice      int32  `json:"one_user_price" bson:"one_user_price"`
	Save              string `json:"save" bson:"save"`
	FreeDays          string `json:"free_days" bson:"free_days"`
	PriceType         string `json:"price_type" bson:"price_type"`
	IsProbation       string `json:"probation" bson:"probation"`
	Platform          string `json:"platform" bson:"platform"`
	Discount          int32  `json:"discount" bson:"discount"`
	OriginalPrice     string `json:"original_price" bson:"original_price"`
	OriginalPriceShow string `json:"original_price_show" bson:"original_price_show"`
}

type VipCheckOrderParams struct {
	Address      string `json:"addr"`
	Timestamp    int64  `json:"timestamp"`
	Nonce        string `json:"nonce"`
	Token        string `json:"token"`
	OrgId        string `json:"org_id"`
	Platform     string `json:"platform"`
	OriginData   string `json:"origin_data"`
	PersonalAuth string `json:"personal_auth"`
}

func (v VipCheckOrderParams) Check() (string, bool) {
	if !tools.IsValidAddress(v.Address) {
		return fmt.Sprintf("invalid params address(%v)", v.Address), false
	}
	if !util.CheckTimestamp(v.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), v.Timestamp)
		return msg, false
	}

	if v.Token != VipCheckOrderToken {
		return fmt.Sprintf("invalid params token(%v)", v.Token), false
	}
	if v.OrgId == "" {
		return fmt.Sprintf("invalid params org_id(%v)", v.OrgId), false
	}
	if v.PersonalAuth == "" {
		return fmt.Sprintf("invalid params personal_auth(%v)", v.PersonalAuth), false
	}
	if v.Platform == "" {
		return fmt.Sprintf("invalid params platform(%v)", v.Platform), false
	}

	if v.Platform == consts.PLATFORM_GOOGLE {
		if v.OriginData == "" {
			msg := fmt.Sprintf("invalid params origin_data value=%s", v.OriginData)
			return msg, false
		}
	}

	return "", true
}

type VipCheckOrderRspData struct {
	OrderStatus string `json:"order_status" bson:"order_status"`
}

type VipAppleVerifyReceiptParams struct {
	OrgId         string `json:"org_id"`
	Address       string `json:"addr"`
	Timestamp     int64  `json:"timestamp"`
	Nonce         string `json:"nonce"`
	Token         string `json:"token"`
	PersonalAuth  string `json:"personal_auth"`
	OrderId       string `json:"order_id"`
	Receipt       string `json:"receipt"`
	TransactionId string `json:"transaction_id"`
	Restore       int32  `json:"restore"`
}

func (v VipAppleVerifyReceiptParams) Check() (string, bool) {
	if !tools.IsValidAddress(v.Address) {
		return "invalid address " + v.Address, false
	}

	if v.PersonalAuth == "" {
		return "invalid personal auth" + v.PersonalAuth, false
	}
	if v.Receipt == "" {
		return "invalid params receipt" + v.Receipt, false
	}
	if v.Restore == 0 && v.OrderId == "" {
		return fmt.Sprintf("invalid params restore = %d or order_id = %s", v.Restore, v.OrderId), false
	}

	if v.Token != VipAppleVerifyReceiptToken {
		return "invalid token" + v.Token, false
	}

	if !util.CheckTimestamp(v.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp = %d, client timestamp = %d", time.Now().Unix(), v.Timestamp)
		return msg, false
	}

	return "", true
}

type VipAppleVerifyReceiptRspData struct {
	OrderId string `json:"order_id" bson:"order_id"`
}

type GetDiscountCodeInfoParams struct {
	Address      string         `json:"addr"`
	Timestamp    int64          `json:"timestamp"`
	Nonce        string         `json:"nonce"`
	Token        string         `json:"token"`
	App          consts.AppType `json:"app"`
	DiscountCode string         `json:"discount_code"`
	Version      string         `json:"version"`
}

func (v GetDiscountCodeInfoParams) Check() (string, bool) {
	if !tools.IsValidAddress(v.Address) {
		return "invalid address " + v.Address, false
	}

	if v.Token != VipDiscountToken {
		return fmt.Sprintf("invalid token %q", v.Token) + v.Token, false
	}

	if !util.CheckTimestamp(v.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), v.Timestamp)
		return msg, false
	}
	if v.DiscountCode == "" {
		return "invalid params discount_code " + v.DiscountCode, false
	}

	if !v.App.Check() {
		msg := fmt.Sprintf("invalid request params app=%s", v.App)
		return msg, false
	}

	return "", true
}

type GetOrderListParams struct {
	Address      string `json:"addr"`
	Timestamp    int64  `json:"timestamp"`
	Nonce        string `json:"nonce"`
	Token        string `json:"token"`
	PersonalAuth string `json:"personal_auth"`
	Page         int32  `json:"page"`
	Pagesize     int32  `json:"pagesize"`
}

func (v GetOrderListParams) Check() (string, bool) {
	if !tools.IsValidAddress(v.Address) {
		return "invalid address " + v.Address, false
	}

	if v.Token != VipGetOrderList {
		return fmt.Sprintf("invalid token %q", v.Token) + v.Token, false
	}

	if !util.CheckTimestamp(v.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), v.Timestamp)
		return msg, false
	}
	if v.PersonalAuth == "" {
		return "invalid params personal_auth " + v.PersonalAuth, false
	}

	return "", true
}

type GetOrderListRspData struct {
	Total int32        `json:"total" bson:"total"`
	List  []*OrderList `json:"list" bson:"list"`
}

type OrderList struct {
	Addr        string `json:"addr" bson:"addr"`
	OrgId       string `json:"org_id" bson:"org_id"`
	OrderId     string `json:"order_id" bson:"order_id"`
	ProductDesc string `json:"product_desc" bson:"product_desc"`
	Status      int32  `json:"status" bson:"status"`
	PayPlatform string `json:"pay_platform" bson:"pay_platform"`
	Amount      int32  `json:"amount" bson:"amount"`
	PayTime     int32  `json:"pay_time" bson:"pay_time"`
}

type VipIOSPromotionSignParams struct {
	Address      string `json:"addr"`
	Timestamp    int64  `json:"timestamp"`
	Nonce        string `json:"nonce"`
	Token        string `json:"token"`
	PersonalAuth string `json:"personal_auth"`
}

func (v VipIOSPromotionSignParams) Check() (string, bool) {
	if !tools.IsValidAddress(v.Address) {
		return fmt.Sprintf("invalid params address(%v)", v.Address), false
	}
	if !util.CheckTimestamp(v.Timestamp) {
		msg := fmt.Sprintf("invalid timestamp server timestamp=%d, client timestamp=%d", time.Now().Unix(), v.Timestamp)
		return msg, false
	}

	if v.Token != VipIOSPromotionSign {
		return fmt.Sprintf("invalid params token(%v)", v.Token), false
	}

	if v.PersonalAuth == "" {
		return fmt.Sprintf("invalid params personal_auth(%v)", v.PersonalAuth), false
	}

	return "", true
}

type VipPriceRspData struct {
	TotalAmount     string `json:"total_amount" bson:"total_amount"`
	DiscountAmount  string `json:"discount_amount" bson:"discount_amount"`
	PromotionAmount string `json:"promotion_amount" bson:"promotion_amount"`
	PayAmount       string `json:"pay_amount" bson:"pay_amount"`
	RemainDays      string `json:"remain_days" bson:"remain_days"`
}
