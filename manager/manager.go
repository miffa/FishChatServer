//
// Copyright 2014 Hong Miao. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"

	"FishChatServer/libnet"
	"FishChatServer/log"
)

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
const char* build_time(void) {
	static const char* psz_build_time = "["__DATE__ " " __TIME__ "]";
	return psz_build_time;
}
*/
import "C"

var (
	buildTime = C.GoString(C.build_time())
)

func BuildTime() string {
	return buildTime
}

func init() {
	flag.Set("alsologtostderr", "false")
	flag.Set("log_dir", "false")
}

const VERSION string = "0.15"

func version() {
	fmt.Printf("manager version %s Copyright (c) 2014-2015 Harold Miao (miaohong@miaohong.org)  \n", VERSION)
}

var InputConfFile = flag.String("conf_file", "manager.json", "input conf file name")

func main() {
	version()
	fmt.Printf("built on %s\n", BuildTime())
	flag.Parse()
	cfg := NewManagerConfig(*InputConfFile)
	err := cfg.LoadConfig()
	if err != nil {
		log.Error(err.Error())
		return
	}

	server, err := libnet.Listen(cfg.TransportProtocols, cfg.Listen)
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("server start:", server.Listener().Addr().String())

	sm := NewManager(cfg)
	go sm.subscribeChannels()

	server.Serve(func(session *libnet.Session) {

	})
}
