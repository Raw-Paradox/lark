// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// SubscribeEvent 用于订阅服务台事件
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/event/subscribe
func (r *HelpdeskAPI) SubscribeEvent(ctx context.Context, request *SubscribeEventReq, options ...MethodOptionFunc) (*SubscribeEventResp, *Response, error) {
	if r.cli.mock.mockHelpdeskSubscribeEvent != nil {
		r.cli.logDebug(ctx, "[lark] Helpdesk#SubscribeEvent mock enable")
		return r.cli.mock.mockHelpdeskSubscribeEvent(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Helpdesk#SubscribeEvent call api")
	r.cli.logDebug(ctx, "[lark] Helpdesk#SubscribeEvent request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/events/subscribe",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(subscribeEventResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Helpdesk#SubscribeEvent POST https://open.feishu.cn/open-apis/helpdesk/v1/events/subscribe failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Helpdesk#SubscribeEvent POST https://open.feishu.cn/open-apis/helpdesk/v1/events/subscribe failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Helpdesk", "SubscribeEvent", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Helpdesk#SubscribeEvent request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockHelpdeskSubscribeEvent(f func(ctx context.Context, request *SubscribeEventReq, options ...MethodOptionFunc) (*SubscribeEventResp, *Response, error)) {
	r.mockHelpdeskSubscribeEvent = f
}

func (r *Mock) UnMockHelpdeskSubscribeEvent() {
	r.mockHelpdeskSubscribeEvent = nil
}

type SubscribeEventReq struct {
	Events []*SubscribeEventReqEvent `json:"events,omitempty"` // 可订阅的事件列表
}

type SubscribeEventReqEvent struct {
	Type    string `json:"type,omitempty"`    // 事件类型, 示例值："helpdesk.ticket_message"
	Subtype string `json:"subtype,omitempty"` // 事件子类型, 示例值："ticket_message.created_v1"
}

type subscribeEventResp struct {
	Code int                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *SubscribeEventResp `json:"data,omitempty"`
}

type SubscribeEventResp struct{}
