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

// GetContactJobFamily 该接口用于获取单个序列信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/job_family/get
func (r *ContactService) GetContactJobFamily(ctx context.Context, request *GetContactJobFamilyReq, options ...MethodOptionFunc) (*GetContactJobFamilyResp, *Response, error) {
	if r.cli.mock.mockContactGetContactJobFamily != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#GetContactJobFamily mock enable")
		return r.cli.mock.mockContactGetContactJobFamily(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "GetContactJobFamily",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/job_families/:job_family_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getContactJobFamilyResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactGetContactJobFamily mock ContactGetContactJobFamily method
func (r *Mock) MockContactGetContactJobFamily(f func(ctx context.Context, request *GetContactJobFamilyReq, options ...MethodOptionFunc) (*GetContactJobFamilyResp, *Response, error)) {
	r.mockContactGetContactJobFamily = f
}

// UnMockContactGetContactJobFamily un-mock ContactGetContactJobFamily method
func (r *Mock) UnMockContactGetContactJobFamily() {
	r.mockContactGetContactJobFamily = nil
}

// GetContactJobFamilyReq ...
type GetContactJobFamilyReq struct {
	JobFamilyID string `path:"job_family_id" json:"-"` // 序列ID, 示例值: "mga5oa8ayjlp9rb"
}

// GetContactJobFamilyResp ...
type GetContactJobFamilyResp struct {
	JobFamily *GetContactJobFamilyRespJobFamily `json:"job_family,omitempty"` // 序列信息
}

// GetContactJobFamilyRespJobFamily ...
type GetContactJobFamilyRespJobFamily struct {
	Name              string                                             `json:"name,omitempty"`                 // 序列名称。1-100字符, 支持中、英文及符号
	Description       string                                             `json:"description,omitempty"`          // 序列描述, 描述序列详情信息
	ParentJobFamilyID string                                             `json:"parent_job_family_id,omitempty"` // 上级序列ID。需是该租户的序列ID列表中的值, 对应唯一的序列名称。
	Status            bool                                               `json:"status,omitempty"`               // 是否启用
	I18nName          []*GetContactJobFamilyRespJobFamilyI18nName        `json:"i18n_name,omitempty"`            // 多语言序列名称
	I18nDescription   []*GetContactJobFamilyRespJobFamilyI18nDescription `json:"i18n_description,omitempty"`     // 多语言描述
	JobFamilyID       string                                             `json:"job_family_id,omitempty"`        // 职级序列ID
}

// GetContactJobFamilyRespJobFamilyI18nDescription ...
type GetContactJobFamilyRespJobFamilyI18nDescription struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}

// GetContactJobFamilyRespJobFamilyI18nName ...
type GetContactJobFamilyRespJobFamilyI18nName struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}

// getContactJobFamilyResp ...
type getContactJobFamilyResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *GetContactJobFamilyResp `json:"data,omitempty"`
}
