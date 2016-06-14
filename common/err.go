//
// Copyright 2014 Hong Miao (miaohong@miaohong.org). All Rights Reserved.
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

package common

import (
	"errors"
)

var (
	NOTFOUNT        = errors.New("NOTFOUNT")
	NOTOPIC         = errors.New("NO TOPIC")
	SYNTAX_ERROR    = errors.New("SYNTAX ERROR")
	NOT_LOGIN       = errors.New("NOT LOGIN")
	DENY_ACCESS     = errors.New("DENY ACCESS")
	TOPIC_EXIST     = errors.New("TOPIC EXIST")
	TOPIC_NOT_EXIST = errors.New("TOPIC NOT EXIST")
	MONGO_ACCESS    = errors.New("MONGO NOT ACCESSABLE")
	NOT_ONLINE      = errors.New("NOT ONLINE")
	NOT_MEMBER      = errors.New("NOT MEMBER")
	MEMBER_EXIST    = errors.New("MEMBER EXIST")
)
