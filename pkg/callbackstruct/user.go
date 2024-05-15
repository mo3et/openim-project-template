// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package callbackstruct

import (
	pbuser "github.com/openimsdk/openim-project-template/pkg/protocol/user"
)

type CallbackBeforeUserRegisterReq struct {
	CallbackCommand `json:"callbackCommand"`
	Secret          string             `json:"secret"`
	Users           []*pbuser.UserInfo `json:"users"`
}

type CallbackBeforeUserRegisterResp struct {
	CommonCallbackResp
	Users []*pbuser.UserInfo `json:"users"`
}

type CallbackAfterUserRegisterReq struct {
	CallbackCommand `json:"callbackCommand"`
	Secret          string             `json:"secret"`
	Users           []*pbuser.UserInfo `json:"users"`
}

type CallbackAfterUserRegisterResp struct {
	CommonCallbackResp
}