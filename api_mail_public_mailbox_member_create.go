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

// CreatePublicMailboxMember 向公共邮箱添加单个成员
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/create
func (r *MailService) CreatePublicMailboxMember(ctx context.Context, request *CreatePublicMailboxMemberReq, options ...MethodOptionFunc) (*CreatePublicMailboxMemberResp, *Response, error) {
	if r.cli.mock.mockMailCreatePublicMailboxMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#CreatePublicMailboxMember mock enable")
		return r.cli.mock.mockMailCreatePublicMailboxMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "CreatePublicMailboxMember",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createPublicMailboxMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailCreatePublicMailboxMember mock MailCreatePublicMailboxMember method
func (r *Mock) MockMailCreatePublicMailboxMember(f func(ctx context.Context, request *CreatePublicMailboxMemberReq, options ...MethodOptionFunc) (*CreatePublicMailboxMemberResp, *Response, error)) {
	r.mockMailCreatePublicMailboxMember = f
}

// UnMockMailCreatePublicMailboxMember un-mock MailCreatePublicMailboxMember method
func (r *Mock) UnMockMailCreatePublicMailboxMember() {
	r.mockMailCreatePublicMailboxMember = nil
}

// CreatePublicMailboxMemberReq ...
type CreatePublicMailboxMemberReq struct {
	PublicMailboxID string        `path:"public_mailbox_id" json:"-"` // 公共邮箱唯一标识或公共邮箱地址, 示例值: "xxxxxxxxxxxxxxx 或 test_public_mailbox@xxx.xx"
	UserIDType      *IDType       `query:"user_id_type" json:"-"`     // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	UserID          *string       `json:"user_id,omitempty"`          // 租户内用户的唯一标识（当成员类型是USER时有值）, 示例值: "xxxxxxxxxx"
	Type            *MailUserType `json:"type,omitempty"`             // 成员类型, 示例值: "USER", 可选值有: USER: 内部用户
}

// CreatePublicMailboxMemberResp ...
type CreatePublicMailboxMemberResp struct {
	MemberID string       `json:"member_id,omitempty"` // 公共邮箱内成员唯一标识
	UserID   string       `json:"user_id,omitempty"`   // 租户内用户的唯一标识（当成员类型是USER时有值）
	Type     MailUserType `json:"type,omitempty"`      // 成员类型, 可选值有: USER: 内部用户
}

// createPublicMailboxMemberResp ...
type createPublicMailboxMemberResp struct {
	Code int64                          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                         `json:"msg,omitempty"`  // 错误描述
	Data *CreatePublicMailboxMemberResp `json:"data,omitempty"`
}
