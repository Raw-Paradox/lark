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

// DeleteCalendarACL
//
// 该接口用于以当前身份（应用 / 用户）删除日历的控制权限，即日历成员。
// 身份由 Header Authorization 的 Token 类型决定。
// 当前身份需要有日历的 owner 权限，并且日历的类型只能为 primary 或 shared。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/delete
func (r *CalendarService) DeleteCalendarACL(ctx context.Context, request *DeleteCalendarACLReq, options ...MethodOptionFunc) (*DeleteCalendarACLResp, *Response, error) {
	if r.cli.mock.mockCalendarDeleteCalendarACL != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#DeleteCalendarACL mock enable")
		return r.cli.mock.mockCalendarDeleteCalendarACL(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "DeleteCalendarACL",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/acls/:acl_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteCalendarACLResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarDeleteCalendarACL mock CalendarDeleteCalendarACL method
func (r *Mock) MockCalendarDeleteCalendarACL(f func(ctx context.Context, request *DeleteCalendarACLReq, options ...MethodOptionFunc) (*DeleteCalendarACLResp, *Response, error)) {
	r.mockCalendarDeleteCalendarACL = f
}

// UnMockCalendarDeleteCalendarACL un-mock CalendarDeleteCalendarACL method
func (r *Mock) UnMockCalendarDeleteCalendarACL() {
	r.mockCalendarDeleteCalendarACL = nil
}

// DeleteCalendarACLReq ...
type DeleteCalendarACLReq struct {
	CalendarID string `path:"calendar_id" json:"-"` // 日历ID。参见[日历ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction), 示例值："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	ACLID      string `path:"acl_id" json:"-"`      // acl资源ID。参见[ACL ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/introduction), 示例值："user_xxxxxx"
}

// deleteCalendarACLResp ...
type deleteCalendarACLResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *DeleteCalendarACLResp `json:"data,omitempty"`
}

// DeleteCalendarACLResp ...
type DeleteCalendarACLResp struct {
}