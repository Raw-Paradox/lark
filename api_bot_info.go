// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// GetBotInfo 获取机器人的基本信息。
//
// 需要启用机器人能力（前往[开发者后台](https://open.feishu.cn/app) - 选择你要获取信息的应用 - 导航栏点击应用功能 - 机器人，开启机器人能力并发布后即可。）
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uAjMxEjLwITMx4CMyETM
func (r *BotService) GetBotInfo(ctx context.Context, request *GetBotInfoReq, options ...MethodOptionFunc) (*GetBotInfoResp, *Response, error) {
	if r.cli.mock.mockBotGetBotInfo != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bot#GetBotInfo mock enable")
		return r.cli.mock.mockBotGetBotInfo(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bot",
		API:                   "GetBotInfo",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/bot/v3/info",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getBotInfoResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBotGetBotInfo mock BotGetBotInfo method
func (r *Mock) MockBotGetBotInfo(f func(ctx context.Context, request *GetBotInfoReq, options ...MethodOptionFunc) (*GetBotInfoResp, *Response, error)) {
	r.mockBotGetBotInfo = f
}

// UnMockBotGetBotInfo un-mock BotGetBotInfo method
func (r *Mock) UnMockBotGetBotInfo() {
	r.mockBotGetBotInfo = nil
}

// GetBotInfoReq ...
type GetBotInfoReq struct {
}

// getBotInfoResp ...
type getBotInfoResp struct {
	Code int64           `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *GetBotInfoResp `json:"bot,omitempty"`  // 机器人信息
}

// GetBotInfoResp ...
type GetBotInfoResp struct {
	ActivateStatus int64    `json:"activate_status,omitempty"` // app 当前状态。,0: 初始化，租户待安装,1: 租户停用,2: 租户启用,3: 安装后待启用,4: 升级待启用,5: license过期停用,6: Lark套餐到期或降级停用
	AppName        string   `json:"app_name,omitempty"`        // app 名称
	AvatarURL      string   `json:"avatar_url,omitempty"`      // app 图像地址
	IpWhiteList    []string `json:"ip_white_list,omitempty"`   // app 的 IP 白名单地址
	OpenID         string   `json:"open_id,omitempty"`         // 机器人的open_id
}