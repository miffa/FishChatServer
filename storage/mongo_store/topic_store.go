//
// Copyright 2014-2015 Hong Miao (miaohong@miaohong.org). All Rights Reserved.
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

package mongo_store

type Member struct {
	ID   string
	Name string
	Type string
}

func NewMember(ID string, Name string, Type string) *Member {
	return &Member{
		ID:   ID,
		Name: Name,
		Type: Type,
	}
}

type TopicStoreData struct {
	TopicName  string    `bson:"TopicName"`
	CreaterID  string    `bson:"CreaterID"`
	MemberList []*Member `bson:"MemberList"`
	//MsgServerAddr  string      `bson:"MsgServerAddr"`
}

func NewTopicStoreData(topicName string, createrID string /*, msgServerAddr string*/) *TopicStoreData {
	return &TopicStoreData{
		TopicName:  topicName,
		CreaterID:  createrID,
		MemberList: make([]*Member, 0),
		//MsgServerAddr: msgServerAddr,
	}
}

func (self *TopicStoreData) AddMember(m *Member) {
	self.MemberList = append(self.MemberList, m)
}

func (self *TopicStoreData) MemberExist(t string) bool {
	for _, m := range self.MemberList {
		if m.ID == t {
			return true
		}
	}
	return false
}

func (self *TopicStoreData) RemoveMember(t string) {
	for i, m := range self.MemberList {
		if m.ID == t {
			self.MemberList = append(self.MemberList[:i], self.MemberList[i+1:]...)
			break
		}
	}
}
