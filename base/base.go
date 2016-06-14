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

package base

import (
	"FishChatServer/libnet"
	"FishChatServer/protocol"
)

type ChannelMap map[string]*ChannelState
type SessionMap map[string]*libnet.Session

type AckMap map[string]map[string]string

const COMM_PREFIX = "IM"

var ChannleList []string

func init() {
	ChannleList = []string{protocol.SYSCTRL_CLIENT_STATUS, protocol.SYSCTRL_TOPIC_STATUS, protocol.SYSCTRL_TOPIC_SYNC,
		protocol.SYSCTRL_SEND, protocol.SYSCTRL_MONITOR, protocol.STORE_CLIENT_INFO, protocol.STORE_TOPIC_INFO}
}

type ChannelState struct {
	ChannelName  string
	Channel      *libnet.Channel
	ClientIDlist []string
}

func NewChannelState(channelName string, channel *libnet.Channel) *ChannelState {
	return &ChannelState{
		ChannelName:  channelName,
		Channel:      channel,
		ClientIDlist: make([]string, 0),
	}
}

type SessionState struct {
	ClientID   string
	Alive      bool
	ClientType string
}

func NewSessionState(alive bool, cid string, clienttype string) *SessionState {
	return &SessionState{
		ClientID:   cid,
		Alive:      alive,
		ClientType: clienttype,
	}
}

type Config interface {
	LoadConfig(configfile string) (*Config, error)
}
