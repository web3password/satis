/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/
package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/web3password/satis/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var Logger *zap.Logger

func SetLogger() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	encoderConfig.LevelKey = "level"
	encoderConfig.TimeKey = "time"
	encoderConfig.MessageKey = "msg"

	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)
	writer := getWriter(config.GetConfig().LogDir, "satis")
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(writer), level),
	)
	Logger = zap.New(core, zap.AddCaller())
	defer Logger.Sync()
}

func Any(k string, v interface{}) zapcore.Field {
	return zap.Any(k, v)
}
func String(k string, v string) zapcore.Field {
	return zap.String(k, v)
}
func Int64(k string, v int64) zapcore.Field {
	return zap.Int64(k, v)
}

func Error(err error) zapcore.Field {
	return zap.Error(err)
}

func getWriter(log_dir, prefix string) io.Writer {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	p := log_dir + prefix + "_%Y%m%d_" + hostname + ".log"

	hook, err := rotatelogs.New(
		p,
		rotatelogs.WithMaxAge(time.Hour*24*7),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
