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

// RecognizeBasicImage 可识别图片中的文字, 按图片中的区域划分, 分段返回文本列表。
//
// 单租户限流: 20QPS, 同租户下的应用没有限流, 共享本租户的 20QPS 限流
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/optical_char_recognition-v1/image/basic_recognize
func (r *AIService) RecognizeBasicImage(ctx context.Context, request *RecognizeBasicImageReq, options ...MethodOptionFunc) (*RecognizeBasicImageResp, *Response, error) {
	if r.cli.mock.mockAIRecognizeBasicImage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] AI#RecognizeBasicImage mock enable")
		return r.cli.mock.mockAIRecognizeBasicImage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "AI",
		API:                   "RecognizeBasicImage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/optical_char_recognition/v1/image/basic_recognize",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(recognizeBasicImageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAIRecognizeBasicImage mock AIRecognizeBasicImage method
func (r *Mock) MockAIRecognizeBasicImage(f func(ctx context.Context, request *RecognizeBasicImageReq, options ...MethodOptionFunc) (*RecognizeBasicImageResp, *Response, error)) {
	r.mockAIRecognizeBasicImage = f
}

// UnMockAIRecognizeBasicImage un-mock AIRecognizeBasicImage method
func (r *Mock) UnMockAIRecognizeBasicImage() {
	r.mockAIRecognizeBasicImage = nil
}

// RecognizeBasicImageReq ...
type RecognizeBasicImageReq struct {
	Image *string `json:"image,omitempty"` // base64 后的图片数据, 示例值: "base64后的图片二进制数据"
}

// RecognizeBasicImageResp ...
type RecognizeBasicImageResp struct {
	TextList []string `json:"text_list,omitempty"` // 按区域识别, 返回文本列表
}

// recognizeBasicImageResp ...
type recognizeBasicImageResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *RecognizeBasicImageResp `json:"data,omitempty"`
}
