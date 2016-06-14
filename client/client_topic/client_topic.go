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
	"fmt"
	"flag"
	"encoding/json"
	"FishChatServer/log"
	"FishChatServer/libnet"
	"FishChatServer/protocol"
	"FishChatServer/common"
)

var InputConfFile = flag.String("conf_file", "client.json", "input conf file name")   

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

func heartBeat(cfg Config, msgServerClient *libnet.Session) {
	hb := common.NewHeartBeat("client", msgServerClient, cfg.HeartBeatTime, cfg.Expire, 10)
	hb.Beat()
}

var msgServerMap map[string]*libnet.Session

func init() {
	msgServerMap = make(map[string]*libnet.Session)
}

func main() {
	flag.Parse()
	cfg, err := LoadConfig(*InputConfFile)
	if err != nil {
		log.Error(err.Error())
		return
	}

	gatewayClient, err := libnet.Dial("tcp", cfg.GatewayServer)
	if err != nil {
		panic(err)
	}
	
	log.Info("req msg_server...")
	cmd := protocol.NewCmdSimple(protocol.REQ_MSG_SERVER_CMD)
	
	err = gatewayClient.Send(libnet.Json(cmd))
	if err != nil {
		log.Error(err.Error())
	}
	
	fmt.Println("input my id :")
	var input string
	if _, err := fmt.Scanf("%s\n", &input); err != nil {
		log.Error(err.Error())
	}
	var c protocol.CmdSimple
	err = gatewayClient.ProcessOnce(func(msg *libnet.InBuffer) error {
		log.Info(string(msg.Data))
		err = json.Unmarshal(msg.Data, &c)
		if err != nil {
			log.Error("error:", err)
		}
		return nil
	})
	if err != nil {
		log.Error(err.Error())
	}

	gatewayClient.Close()
	
	

	msgServerClient, err := libnet.Dial("tcp", string(c.GetArgs()[0]))
	if err != nil {
		panic(err)
	}
	
	msgServerMap[string(c.GetArgs()[0])] = msgServerClient
	
	log.Info("test.. send id...")
	cmd = protocol.NewCmdSimple(protocol.SEND_CLIENT_ID_CMD)
	cmd.AddArg(input)
	
	err = msgServerClient.Send(libnet.Json(cmd))
	if err != nil {
		log.Error(err.Error())
	}
	
	go heartBeat(cfg, msgServerClient)
	
	log.Info("test.. send topic msg...")
	cmd = protocol.NewCmdSimple(protocol.CREATE_TOPIC_CMD)
	fmt.Println("want to create a topic (y/n) :")
	if _, err := fmt.Scanf("%s\n", &input); err != nil {
		log.Error(err.Error())
	}
	if input == "y" {
		fmt.Println("CREATE_TOPIC_CMD | input topic name :")
		if _, err := fmt.Scanf("%s\n", &input); err != nil {
			log.Error(err.Error())
		}
		
		cmd.AddArg(input)
		
		err = msgServerClient.Send(libnet.Json(cmd))
		if err != nil {
			log.Error(err.Error())
		}
	}
	
	cmd = protocol.NewCmdSimple(protocol.JOIN_TOPIC_CMD)
	
	fmt.Println("JOIN_TOPIC_CMD | input topic name :")
	if _, err := fmt.Scanf("%s\n", &input); err != nil {
		log.Error(err.Error())
	}
	
	cmd.AddArg(input)
	fmt.Println("JOIN_TOPIC_CMD | input your ID :")
	if _, err := fmt.Scanf("%s\n", &input); err != nil {
		log.Error(err.Error())
	}
	
	cmd.AddArg(input)
	
	err = msgServerClient.Send(libnet.Json(cmd))
	if err != nil {
		log.Error(err.Error())
	}
	

	err = msgServerClient.ProcessOnce(func(msg *libnet.InBuffer) error {
		log.Info(string(msg.Data))
		err = json.Unmarshal(msg.Data, &c)
		if err != nil {
			log.Error("error:", err)
		}
		return nil
	})
	if err != nil {
		log.Error(err.Error())
	}
	
	topicMsgServerAddrstring := string(c.GetArgs()[0])
	topicName := string(c.GetArgs()[1])
	
	fmt.Println("topicMsgServerAddrstring :" + topicMsgServerAddrstring)
	fmt.Println("topicName :" + topicName)
	
	
	if _, f := msgServerMap[topicMsgServerAddrstring]; f == false {
		fmt.Println("to another server")
		newMsgServerClient, err := libnet.Dial("tcp", topicMsgServerAddrstring)
		if err != nil {
			panic(err)
		}
		
		log.Info("test.. send id...")
		cmd = protocol.NewCmdSimple(protocol.SEND_CLIENT_ID_CMD)
		cmd.AddArg(input)
		
		err = newMsgServerClient.Send(libnet.Json(cmd))
		if err != nil {
			log.Error(err.Error())
		}
		
		go heartBeat(cfg, newMsgServerClient)
		
		cmd = protocol.NewCmdSimple(protocol.JOIN_TOPIC_CMD)
	
		fmt.Println("JOIN_TOPIC_CMD | input topic name :")
		if _, err := fmt.Scanf("%s\n", &input); err != nil {
			log.Error(err.Error())
		}
		
		cmd.AddArg(input)
		fmt.Println("JOIN_TOPIC_CMD | input your ID :")
		if _, err := fmt.Scanf("%s\n", &input); err != nil {
			log.Error(err.Error())
		}
		
		cmd.AddArg(input)
		
		err = newMsgServerClient.Send(libnet.Json(cmd))
		if err != nil {
			log.Error(err.Error())
		}
		
		msgServerMap[topicMsgServerAddrstring] = newMsgServerClient
		
		cmd = protocol.NewCmdSimple(protocol.SEND_MESSAGE_TOPIC_CMD)

		fmt.Println("SEND_MESSAGE_TOPIC_CMD | input topic name :")
		if _, err := fmt.Scanf("%s\n", &input); err != nil {
			log.Error(err.Error())
		}
		
		cmd.AddArg(input)
		fmt.Println("SEND_MESSAGE_TOPIC_CMD | input message :")
		if _, err := fmt.Scanf("%s\n", &input); err != nil {
			log.Error(err.Error())
		}
		
		cmd.AddArg(input)
		
		err = newMsgServerClient.Send(libnet.Json(cmd))
		if err != nil {
			log.Error(err.Error())
		}
		
		newMsgServerClient.Process(func(msg *libnet.InBuffer) error {
			log.Info(string(msg.Data))
			return nil
		})
	} else {
		cmd = protocol.NewCmdSimple(protocol.SEND_MESSAGE_TOPIC_CMD)

		fmt.Println("SEND_MESSAGE_TOPIC_CMD | input topic name :")
		if _, err := fmt.Scanf("%s\n", &input); err != nil {
			log.Error(err.Error())
		}
		
		cmd.AddArg(input)
		fmt.Println("SEND_MESSAGE_TOPIC_CMD | input message :")
		if _, err := fmt.Scanf("%s\n", &input); err != nil {
			log.Error(err.Error())
		}
		
		cmd.AddArg(input)
		
		err = msgServerClient.Send(libnet.Json(cmd))
		if err != nil {
			log.Error(err.Error())
		}
		
		
		defer msgServerClient.Close()
		
		msgServerClient.Process(func(msg *libnet.InBuffer) error {
			log.Info(string(msg.Data))
			return nil
		})
	}
	log.Flush()
}
