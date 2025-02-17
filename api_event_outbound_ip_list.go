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

// GetEventOutboundIpList 飞书开放平台向应用配置的回调地址推送事件时, 是通过特定的 IP 发送出去的, 应用可以通过本接口获取所有相关的 IP 地址。
//
// IP 地址有变更可能, 建议应用每隔 6 小时定时拉取最新的 IP 地址, 以免由于企业防火墙设置, 导致应用无法及时接收到飞书开放平台推送的事件。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/event-v1/outbound_ip/list
func (r *EventService) GetEventOutboundIpList(ctx context.Context, request *GetEventOutboundIpListReq, options ...MethodOptionFunc) (*GetEventOutboundIpListResp, *Response, error) {
	if r.cli.mock.mockEventGetEventOutboundIpList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Event#GetEventOutboundIpList mock enable")
		return r.cli.mock.mockEventGetEventOutboundIpList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Event",
		API:                   "GetEventOutboundIpList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/event/v1/outbound_ip",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getEventOutboundIpListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockEventGetEventOutboundIpList mock EventGetEventOutboundIpList method
func (r *Mock) MockEventGetEventOutboundIpList(f func(ctx context.Context, request *GetEventOutboundIpListReq, options ...MethodOptionFunc) (*GetEventOutboundIpListResp, *Response, error)) {
	r.mockEventGetEventOutboundIpList = f
}

// UnMockEventGetEventOutboundIpList un-mock EventGetEventOutboundIpList method
func (r *Mock) UnMockEventGetEventOutboundIpList() {
	r.mockEventGetEventOutboundIpList = nil
}

// GetEventOutboundIpListReq ...
type GetEventOutboundIpListReq struct {
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 10, 最大值: `50`
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "xxx"
}

// GetEventOutboundIpListResp ...
type GetEventOutboundIpListResp struct {
	IpList    []string `json:"ip_list,omitempty"`    // outbound ip
	PageToken string   `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool     `json:"has_more,omitempty"`   // 是否还有更多项
}

// getEventOutboundIpListResp ...
type getEventOutboundIpListResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *GetEventOutboundIpListResp `json:"data,omitempty"`
}
