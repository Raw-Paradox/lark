// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetFAQList 该接口用于获取服务台知识库详情。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/list
func (r *HelpdeskAPI) GetFAQList(ctx context.Context, request *GetFAQListReq, options ...MethodOptionFunc) (*GetFAQListResp, *Response, error) {
	if r.cli.mock.mockHelpdeskGetFAQList != nil {
		r.cli.logDebug(ctx, "[lark] Helpdesk#GetFAQList mock enable")
		return r.cli.mock.mockHelpdeskGetFAQList(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Helpdesk#GetFAQList call api")
	r.cli.logDebug(ctx, "[lark] Helpdesk#GetFAQList request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/faqs",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(getFAQListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Helpdesk#GetFAQList GET https://open.feishu.cn/open-apis/helpdesk/v1/faqs failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Helpdesk#GetFAQList GET https://open.feishu.cn/open-apis/helpdesk/v1/faqs failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Helpdesk", "GetFAQList", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Helpdesk#GetFAQList request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockHelpdeskGetFAQList(f func(ctx context.Context, request *GetFAQListReq, options ...MethodOptionFunc) (*GetFAQListResp, *Response, error)) {
	r.mockHelpdeskGetFAQList = f
}

func (r *Mock) UnMockHelpdeskGetFAQList() {
	r.mockHelpdeskGetFAQList = nil
}

type GetFAQListReq struct {
	CategoryID *string `query:"category_id" json:"-"` // 知识库分类ID, 示例值："6856395522433908739"
	Status     *string `query:"status" json:"-"`      // 搜索条件: 知识库状态 1:在线 0:删除，可恢复 2：删除，不可恢复	, 示例值："1"
	Search     *string `query:"search" json:"-"`      // 搜索条件: 关键词，匹配问题标题，问题关键字，用户姓名	, 示例值："点餐"
	PageToken  *string `query:"page_token" json:"-"`  // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："6856395634652479491"
	PageSize   *int    `query:"page_size" json:"-"`   // 分页大小, 示例值：10, 最大值：`100`
}

type getFAQListResp struct {
	Code int             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *GetFAQListResp `json:"data,omitempty"` //
}

type GetFAQListResp struct {
	HasMore   bool                  `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	PageSize  int                   `json:"page_size,omitempty"`  // 实际返回的FAQ数量
	Total     int                   `json:"total,omitempty"`      // 总数
	Items     []*GetFAQListRespItem `json:"items,omitempty"`      // 知识库列表
}

type GetFAQListRespItem struct {
	FaqID          string                        `json:"faq_id,omitempty"`          // 知识库ID
	ID             string                        `json:"id,omitempty"`              // 知识库旧版ID，请使用faq_id
	HelpdeskID     string                        `json:"helpdesk_id,omitempty"`     // 服务台ID
	Question       string                        `json:"question,omitempty"`        // 问题
	Answer         string                        `json:"answer,omitempty"`          // 答案
	AnswerRichtext string                        `json:"answer_richtext,omitempty"` // 富文本答案
	CreateTime     int                           `json:"create_time,omitempty"`     // 创建时间
	UpdateTime     int                           `json:"update_time,omitempty"`     // 修改时间
	Categories     []*HelpdeskCategory           `json:"categories,omitempty"`      // 分类
	Tags           []string                      `json:"tags,omitempty"`            // 关联词列表
	ExpireTime     int                           `json:"expire_time,omitempty"`     // 失效时间
	UpdateUser     *GetFAQListRespItemUpdateUser `json:"update_user,omitempty"`     // 更新用户
	CreateUser     *GetFAQListRespItemCreateUser `json:"create_user,omitempty"`     // 创建用户
}

type GetFAQListRespItemUpdateUser struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
}

type GetFAQListRespItemCreateUser struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
}
