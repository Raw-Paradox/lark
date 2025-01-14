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

// GetHelpdeskFAQ 该接口用于获取服务台知识库详情。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/get
func (r *HelpdeskService) GetHelpdeskFAQ(ctx context.Context, request *GetHelpdeskFAQReq, options ...MethodOptionFunc) (*GetHelpdeskFAQResp, *Response, error) {
	if r.cli.mock.mockHelpdeskGetHelpdeskFAQ != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#GetHelpdeskFAQ mock enable")
		return r.cli.mock.mockHelpdeskGetHelpdeskFAQ(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "GetHelpdeskFAQ",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/helpdesk/v1/faqs/:id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(getHelpdeskFAQResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskGetHelpdeskFAQ mock HelpdeskGetHelpdeskFAQ method
func (r *Mock) MockHelpdeskGetHelpdeskFAQ(f func(ctx context.Context, request *GetHelpdeskFAQReq, options ...MethodOptionFunc) (*GetHelpdeskFAQResp, *Response, error)) {
	r.mockHelpdeskGetHelpdeskFAQ = f
}

// UnMockHelpdeskGetHelpdeskFAQ un-mock HelpdeskGetHelpdeskFAQ method
func (r *Mock) UnMockHelpdeskGetHelpdeskFAQ() {
	r.mockHelpdeskGetHelpdeskFAQ = nil
}

// GetHelpdeskFAQReq ...
type GetHelpdeskFAQReq struct {
	ID string `path:"id" json:"-"` // 知识库ID, 示例值: "6856395634652479491"
}

// GetHelpdeskFAQResp ...
type GetHelpdeskFAQResp struct {
	FAQ *GetHelpdeskFAQRespFAQ `json:"faq,omitempty"` // 知识库详情
}

// GetHelpdeskFAQRespFAQ ...
type GetHelpdeskFAQRespFAQ struct {
	FAQID          string                                 `json:"faq_id,omitempty"`          // 知识库ID
	ID             string                                 `json:"id,omitempty"`              // 知识库旧版ID, 请使用faq_id
	HelpdeskID     string                                 `json:"helpdesk_id,omitempty"`     // 服务台ID
	Question       string                                 `json:"question,omitempty"`        // 问题
	Answer         string                                 `json:"answer,omitempty"`          // 答案
	AnswerRichtext []*GetHelpdeskFAQRespFAQAnswerRichtext `json:"answer_richtext,omitempty"` // 富文本答案
	CreateTime     int64                                  `json:"create_time,omitempty"`     // 创建时间
	UpdateTime     int64                                  `json:"update_time,omitempty"`     // 修改时间
	Categories     []*HelpdeskCategory                    `json:"categories,omitempty"`      // 分类
	Tags           []string                               `json:"tags,omitempty"`            // 相似问题列表
	ExpireTime     int64                                  `json:"expire_time,omitempty"`     // 失效时间
	UpdateUser     *GetHelpdeskFAQRespFAQUpdateUser       `json:"update_user,omitempty"`     // 更新用户
	CreateUser     *GetHelpdeskFAQRespFAQCreateUser       `json:"create_user,omitempty"`     // 创建用户
}

// GetHelpdeskFAQRespFAQAnswerRichtext ...
type GetHelpdeskFAQRespFAQAnswerRichtext struct {
	Content string `json:"content,omitempty"` // 内容
	Type    string `json:"type,omitempty"`    // 类型
}

// GetHelpdeskFAQRespFAQCreateUser ...
type GetHelpdeskFAQRespFAQCreateUser struct {
	ID         string `json:"id,omitempty"`         // 用户ID
	AvatarURL  string `json:"avatar_url,omitempty"` // 用户头像url
	Name       string `json:"name,omitempty"`       // 用户名
	Department string `json:"department,omitempty"` // 所在部门名称
	City       string `json:"city,omitempty"`       // 城市
	Country    string `json:"country,omitempty"`    // 国家代号(CountryCode), 参考: http://www.mamicode.com/info-detail-2186501.html
}

// GetHelpdeskFAQRespFAQUpdateUser ...
type GetHelpdeskFAQRespFAQUpdateUser struct {
	ID         string `json:"id,omitempty"`         // 用户ID
	AvatarURL  string `json:"avatar_url,omitempty"` // 用户头像url
	Name       string `json:"name,omitempty"`       // 用户名
	Department string `json:"department,omitempty"` // 所在部门名称
	City       string `json:"city,omitempty"`       // 城市
	Country    string `json:"country,omitempty"`    // 国家代号(CountryCode), 参考: http://www.mamicode.com/info-detail-2186501.html
}

// getHelpdeskFAQResp ...
type getHelpdeskFAQResp struct {
	Code int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *GetHelpdeskFAQResp `json:"data,omitempty"`
}
