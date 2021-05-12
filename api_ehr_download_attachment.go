// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
	"io"
)

// DownloadAttachments
//
// 根据文件 token 下载文件。
// 调用「[批量获取员工花名册信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/ehr/employees)」接口的返回值中，「文件」类型的字段 id，即是文件 token
// ![image.png](//sf1-ttcdn-tos.pstatp.com/obj/open-platform-opendoc/bed391d2a8ce6ed2d5985ea69bf92850_9GY1mnuDXP.png)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/ehr/ehr-v1/attachment/get
func (r *EHRAPI) DownloadAttachments(ctx context.Context, request *DownloadAttachmentsReq, options ...MethodOptionFunc) (*DownloadAttachmentsResp, *Response, error) {
	if r.cli.mock.mockEHRDownloadAttachments != nil {
		r.cli.logDebug(ctx, "[lark] EHR#DownloadAttachments mock enable")
		return r.cli.mock.mockEHRDownloadAttachments(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] EHR#DownloadAttachments call api")
	r.cli.logDebug(ctx, "[lark] EHR#DownloadAttachments request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/ehr/v1/attachments/:token",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(downloadAttachmentsResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] EHR#DownloadAttachments GET https://open.feishu.cn/open-apis/ehr/v1/attachments/:token failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] EHR#DownloadAttachments GET https://open.feishu.cn/open-apis/ehr/v1/attachments/:token failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("EHR", "DownloadAttachments", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] EHR#DownloadAttachments request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockEHRDownloadAttachments(f func(ctx context.Context, request *DownloadAttachmentsReq, options ...MethodOptionFunc) (*DownloadAttachmentsResp, *Response, error)) {
	r.mockEHRDownloadAttachments = f
}

func (r *Mock) UnMockEHRDownloadAttachments() {
	r.mockEHRDownloadAttachments = nil
}

type DownloadAttachmentsReq struct {
	Token string `path:"token" json:"-"` // 文件 token, 示例值："09bf7b924f9a4a69875788891b5970d8"
}

type downloadAttachmentsResp struct {
	IsFile bool                     `json:"is_file,omitempty"`
	Code   int                      `json:"code,omitempty"`
	Msg    string                   `json:"msg,omitempty"`
	Data   *DownloadAttachmentsResp `json:"data,omitempty"`
}

func (r *downloadAttachmentsResp) IsFileType() bool {
	return r.IsFile
}

func (r *downloadAttachmentsResp) SetFile(file io.Reader) {
	r.Data = &DownloadAttachmentsResp{
		File: file,
	}
}

type DownloadAttachmentsResp struct {
	File io.Reader `json:"file,omitempty"`
}
