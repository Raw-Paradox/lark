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

// GetVerification 获取企业主体名称、是否认证等信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/verification-v1/verification/get
func (r *VerificationService) GetVerification(ctx context.Context, request *GetVerificationReq, options ...MethodOptionFunc) (*GetVerificationResp, *Response, error) {
	if r.cli.mock.mockVerificationGetVerification != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Verification#GetVerification mock enable")
		return r.cli.mock.mockVerificationGetVerification(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Verification",
		API:                   "GetVerification",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/verification/v1/verification",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getVerificationResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVerificationGetVerification mock VerificationGetVerification method
func (r *Mock) MockVerificationGetVerification(f func(ctx context.Context, request *GetVerificationReq, options ...MethodOptionFunc) (*GetVerificationResp, *Response, error)) {
	r.mockVerificationGetVerification = f
}

// UnMockVerificationGetVerification un-mock VerificationGetVerification method
func (r *Mock) UnMockVerificationGetVerification() {
	r.mockVerificationGetVerification = nil
}

// GetVerificationReq ...
type GetVerificationReq struct {
}

// GetVerificationResp ...
type GetVerificationResp struct {
	Verification *GetVerificationRespVerification `json:"verification,omitempty"` // 认证信息
}

// GetVerificationRespVerification ...
type GetVerificationRespVerification struct {
	Name            string `json:"name,omitempty"`             // 企业主体名称
	HasVerification bool   `json:"has_verification,omitempty"` // 企业是否完成认证； true 表示已经完成认证, false 表示未认证
}

// getVerificationResp ...
type getVerificationResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *GetVerificationResp `json:"data,omitempty"`
}
