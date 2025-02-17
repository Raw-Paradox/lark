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
	"io"
)

// DownloadBaikeImage 通过 file_token 下载原图片。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/file/download
func (r *BaikeService) DownloadBaikeImage(ctx context.Context, request *DownloadBaikeImageReq, options ...MethodOptionFunc) (*DownloadBaikeImageResp, *Response, error) {
	if r.cli.mock.mockBaikeDownloadBaikeImage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Baike#DownloadBaikeImage mock enable")
		return r.cli.mock.mockBaikeDownloadBaikeImage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Baike",
		API:                   "DownloadBaikeImage",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/baike/v1/files/:file_token/download",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(downloadBaikeImageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBaikeDownloadBaikeImage mock BaikeDownloadBaikeImage method
func (r *Mock) MockBaikeDownloadBaikeImage(f func(ctx context.Context, request *DownloadBaikeImageReq, options ...MethodOptionFunc) (*DownloadBaikeImageResp, *Response, error)) {
	r.mockBaikeDownloadBaikeImage = f
}

// UnMockBaikeDownloadBaikeImage un-mock BaikeDownloadBaikeImage method
func (r *Mock) UnMockBaikeDownloadBaikeImage() {
	r.mockBaikeDownloadBaikeImage = nil
}

// DownloadBaikeImageReq ...
type DownloadBaikeImageReq struct {
	FileToken string `path:"file_token" json:"-"` // 需要下载的文件 token, 示例值: "boxbcEcmKiD3***vgqWTpvdc7jc"
}

// downloadBaikeImageResp ...
type downloadBaikeImageResp struct {
	Code int64                   `json:"code,omitempty"`
	Msg  string                  `json:"msg,omitempty"`
	Data *DownloadBaikeImageResp `json:"data,omitempty"`
}

func (r *downloadBaikeImageResp) SetReader(file io.Reader) {
	if r.Data == nil {
		r.Data = &DownloadBaikeImageResp{}
	}
	r.Data.File = file
}

// DownloadBaikeImageResp ...
type DownloadBaikeImageResp struct {
	File io.Reader `json:"file,omitempty"`
}
