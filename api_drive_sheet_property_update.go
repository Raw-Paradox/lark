// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateSheetProperty 该接口用于根据 spreadsheetToken 更新表格属性，如更新表格标题。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ucTMzUjL3EzM14yNxMTN
func (r *DriveAPI) UpdateSheetProperty(ctx context.Context, request *UpdateSheetPropertyReq, options ...MethodOptionFunc) (*UpdateSheetPropertyResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateSheetProperty != nil {
		r.cli.logDebug(ctx, "[lark] Drive#UpdateSheetProperty mock enable")
		return r.cli.mock.mockDriveUpdateSheetProperty(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Drive#UpdateSheetProperty call api")
	r.cli.logDebug(ctx, "[lark] Drive#UpdateSheetProperty request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "PUT",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/properties",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateSheetPropertyResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Drive#UpdateSheetProperty PUT https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/properties failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Drive#UpdateSheetProperty PUT https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/properties failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Drive", "UpdateSheetProperty", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Drive#UpdateSheetProperty request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockDriveUpdateSheetProperty(f func(ctx context.Context, request *UpdateSheetPropertyReq, options ...MethodOptionFunc) (*UpdateSheetPropertyResp, *Response, error)) {
	r.mockDriveUpdateSheetProperty = f
}

func (r *Mock) UnMockDriveUpdateSheetProperty() {
	r.mockDriveUpdateSheetProperty = nil
}

type UpdateSheetPropertyReq struct {
	SpreadsheetToken string                            `path:"spreadsheetToken" json:"-"` // spreadsheet 的 token，获取方式见[ 对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN) 的第 4 项
	Properties       *UpdateSheetPropertyReqProperties `json:"properties,omitempty"`      // spreadsheet 的属性
}

type UpdateSheetPropertyReqProperties struct {
	Title string `json:"title,omitempty"` // spreadsheet 的标题，最大长度100个字符
}

type updateSheetPropertyResp struct {
	Code int                      `json:"code,omitempty"`
	Msg  string                   `json:"msg,omitempty"`
	Data *UpdateSheetPropertyResp `json:"data,omitempty"`
}

type UpdateSheetPropertyResp struct {
	SpreadsheetToken string `json:"spreadsheetToken,omitempty"` // spreadsheet 的 token
	Title            string `json:"title,omitempty"`            // spreadsheet 的标题
}
