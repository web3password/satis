/*
Copyright (C) 2024 Web3Password PTE. LTD.(Singapore UEN: 202333030C) - All Rights Reserved

Web3Password PTE. LTD.(Singapore UEN: 202333030C) holds the copyright of this file.

Unauthorized copying or redistribution of this file in binary forms via any medium is strictly prohibited.

For more information, please refer to https://www.web3password.com/web3password_license.txt
*/
package config

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v3"
)

var (
	config *Config
	lock   = sync.RWMutex{}
)

// Config .
type Config struct {
	RunningMode       string     `yaml:"running_mode"`
	Server            Server     `yaml:"server"`      // server start config
	HttpServer        HttpServer `yaml:"http_server"` // http server start config
	Node              Node       `yaml:"node"`        // node config
	MsgSize           Msg        `yaml:"msg"`         // msg size
	LogDir            string     `yaml:"log_dir"`
	Tls               Tls        `yaml:"tls"`
	OfficialDomain    string     `yaml:"official_domain"`  // official_domain
	OfficialDomains   []string   `yaml:"official_domains"` // official_domains
	PersonalWhiteList string     `yaml:"personal_white_list"`
	OrgWhiteList      string     `yaml:"org_white_list"`
}

type Tls struct {
	Ca        string `yaml:"ca"`
	ServerTls CrtKey `yaml:"server"`
	ClientTls CrtKey `yaml:"client"`
}

type CrtKey struct {
	Crt string `yaml:"crt"`
	Key string `yaml:"key"`
}

// Server .
type Server struct {
	EnableTLS bool   `yaml:"enable_tls"`
	TLSDomain string `yaml:"tls_domain"`
	IP        string `yaml:"ip"`
	Port      string `yaml:"port"`
	Proto     string `yaml:"proto"`
}

// HttpServer .
type HttpServer struct {
	IP          string `yaml:"ip"`
	Port        string `yaml:"port"`
	WithTraceID bool   `yaml:"with_trace_id"`
}

// Node .
type Node struct {
	Token string `yaml:"token"`
}

type Msg struct {
	Api  int `yaml:"api"`
	File int `yaml:"file"`
}

// GetServerProto .
func (c *Config) GetServerProto() string {
	return c.Server.Proto
}

// GetGRPCServerAddress .
func (c *Config) GetGRPCServerAddress() string {
	return fmt.Sprintf("%s:%s", c.Server.IP, c.Server.Port)
}

// GetHttpServerAddress .
func (c *Config) GetHttpServerAddress() string {
	return fmt.Sprintf("%s:%s", c.HttpServer.IP, c.HttpServer.Port)
}

func WatchConfig(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					ParseConfig(path)
					fmt.Println("modified file:", event.Name, config.MsgSize.File)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()
	err = watcher.Add(path)
	if err != nil {
		panic(err)
	}
	select {}
}
func ParseConfig(path string) {
	fmt.Println("load config", path)
	file, err := os.Open(path)
	if err != nil {
		panic("failed open config file err")
	}
	fileByte, err := io.ReadAll(file)
	if err != nil {
		panic("failed read config file err")
	}
	c := &Config{}
	if err = yaml.Unmarshal(fileByte, c); err != nil {
		panic("parse config file err")
	}

	lock.Lock()
	defer lock.Unlock()
	config = c
}
func GetConfig() *Config {
	lock.RLock()
	defer lock.RUnlock()
	return config
}
