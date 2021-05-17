// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetBotInfo 获取机器人的基本信息。
//
// 需要启用机器人能力
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uAjMxEjLwITMx4CMyETM
func (r *BotService) GetBotInfo(ctx context.Context, request *GetBotInfoReq, options ...MethodOptionFunc) (*GetBotInfoResp, *Response, error) {
	r.cli.logInfo(ctx, "[lark][Bot][GetBotInfo] call api")
	r.cli.logDebug(ctx, "[lark][Bot][GetBotInfo] request: %s", jsonString(request))

	if r.cli.mock.mockBotGetBotInfo != nil {
		r.cli.logDebug(ctx, "[lark][Bot][GetBotInfo] mock enable")
		return r.cli.mock.mockBotGetBotInfo(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/bot/v3/info",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getBotInfoResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark][Bot][GetBotInfo] GET https://open.feishu.cn/open-apis/bot/v3/info failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark][Bot][GetBotInfo] GET https://open.feishu.cn/open-apis/bot/v3/info failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Bot", "GetBotInfo", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark][Bot][GetBotInfo] request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockBotGetBotInfo(f func(ctx context.Context, request *GetBotInfoReq, options ...MethodOptionFunc) (*GetBotInfoResp, *Response, error)) {
	r.mockBotGetBotInfo = f
}

func (r *Mock) UnMockBotGetBotInfo() {
	r.mockBotGetBotInfo = nil
}

type GetBotInfoReq struct{}

type getBotInfoResp struct {
	Code int             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *GetBotInfoResp `json:"bot,omitempty"`  // 机器人信息
}

type GetBotInfoResp struct {
	ActivateStatus int      `json:"activate_status,omitempty"` // app 当前状态。,0: 初始化，租户待安装,1: 租户停用,2: 租户启用,3: 安装后待启用,4: 升级待启用,5: license过期停用,6: Lark套餐到期或降级停用
	AppName        string   `json:"app_name,omitempty"`        // app 名称
	AvatarUrl      string   `json:"avatar_url,omitempty"`      // app 图像地址
	IpWhiteList    []string `json:"ip_white_list,omitempty"`   // app 的 IP 白名单地址
	OpenID         string   `json:"open_id,omitempty"`         // 机器人的open_id
}