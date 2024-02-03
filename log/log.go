/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/
package log

import (
	"context"
	"github.com/sirupsen/logrus"
)

// log config
type logConfig struct {
	log struct {
		level int `yaml:"level"`
	} `yaml:"log"`
}

var log = logrus.New()

func init() {
	//log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		CallerPrettyfier:  nil,
		PrettyPrint:       false,
	})
	log.SetLevel(logrus.DebugLevel)
}

func InfoContextf(ctx context.Context, format string, args ...any) {
	log.WithContext(ctx).Infof(format, args...)
}

func DebugContextf(ctx context.Context, format string, args ...any) {
	log.WithContext(ctx).Debugf(format, args...)
}

func ErrorContextf(ctx context.Context, format string, args ...any) {
	log.WithContext(ctx).Errorf(format, args...)
}

func TraceContextf(ctx context.Context, format string, args ...any) {
	log.WithContext(ctx).Tracef(format, args...)
}

func Infof(format string, args ...any) {
	log.Infof(format, args...)
}

func Errorf(format string, args ...any) {
	log.Errorf(format, args...)
}

func Debugf(format string, args ...any) {
	log.Debugf(format, args...)
}

func Fatalf(format string, args ...any) {
	log.Fatalf(format, args...)
}
