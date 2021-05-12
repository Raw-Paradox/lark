// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// BatchGetOKR 根据OKR id批量获取OKR
//
// 使用<md-tag mode="inline" type="token-tenant">tenant_access_token</md-tag>需要额外申请权限以应用身份访问OKR信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/okr/batch_get
func (r *OKRAPI) BatchGetOKR(ctx context.Context, request *BatchGetOKRReq, options ...MethodOptionFunc) (*BatchGetOKRResp, *Response, error) {
	if r.cli.mock.mockOKRBatchGetOKR != nil {
		r.cli.logDebug(ctx, "[lark] OKR#BatchGetOKR mock enable")
		return r.cli.mock.mockOKRBatchGetOKR(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] OKR#BatchGetOKR call api")
	r.cli.logDebug(ctx, "[lark] OKR#BatchGetOKR request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/okr/v1/okrs/batch_get",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(batchGetOKRResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] OKR#BatchGetOKR GET https://open.feishu.cn/open-apis/okr/v1/okrs/batch_get failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] OKR#BatchGetOKR GET https://open.feishu.cn/open-apis/okr/v1/okrs/batch_get failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("OKR", "BatchGetOKR", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] OKR#BatchGetOKR request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockOKRBatchGetOKR(f func(ctx context.Context, request *BatchGetOKRReq, options ...MethodOptionFunc) (*BatchGetOKRResp, *Response, error)) {
	r.mockOKRBatchGetOKR = f
}

func (r *Mock) UnMockOKRBatchGetOKR() {
	r.mockOKRBatchGetOKR = nil
}

type BatchGetOKRReq struct {
	UserIDType *IDType  `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	OkrIDs     []string `query:"okr_ids" json:"-"`      // OKR ID 列表，最多10个
	Lang       *string  `query:"lang" json:"-"`         // 请求OKR的语言版本（比如@的人名），lang=en_us/zh_cn，请求 Query中, 示例值："zh_cn", 默认值: `zh_cn`
}

type batchGetOKRResp struct {
	Code int              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string           `json:"msg,omitempty"`  // 错误描述
	Data *BatchGetOKRResp `json:"data,omitempty"` //
}

type BatchGetOKRResp struct {
	OkrList []*BatchGetOKRRespOkr `json:"okr_list,omitempty"` // OKR 列表
}

type BatchGetOKRRespOkr struct {
	ID            string                         `json:"id,omitempty"`             // id
	Permission    int                            `json:"permission,omitempty"`     // OKR的访问权限, 可选值有: `0`：此时OKR只返回id, `1`：返回OKR的其他具体字段
	Name          string                         `json:"name,omitempty"`           // 名称
	ObjectiveList []*BatchGetOKRRespOkrObjective `json:"objective_list,omitempty"` // Objective列表
}

type BatchGetOKRRespOkrObjective struct {
	ID                    string                                          `json:"id,omitempty"`                      // Objective ID
	Permission            int                                             `json:"permission,omitempty"`              // 权限, 可选值有: `0`：此时OKR只返回id, `1`：返回OKR的其他具体字段
	Content               string                                          `json:"content,omitempty"`                 // Objective 内容
	ProgressReport        string                                          `json:"progress_report,omitempty"`         // Objective 进度记录内容
	Score                 int                                             `json:"score,omitempty"`                   // Objective 分数（0 - 100）
	Weight                float64                                         `json:"weight,omitempty"`                  // Objective的权重（0 - 100）
	ProgressRate          *BatchGetOKRRespOkrObjectiveProgressRate        `json:"progress_rate,omitempty"`           // Objective进度
	KrList                []*BatchGetOKRRespOkrObjectiveKr                `json:"kr_list,omitempty"`                 // Objective KeyResult 列表
	AlignedObjectiveList  []*BatchGetOKRRespOkrObjectiveAlignedObjective  `json:"aligned_objective_list,omitempty"`  // 对齐到该Objective的Objective列表
	AligningObjectiveList []*BatchGetOKRRespOkrObjectiveAligningObjective `json:"aligning_objective_list,omitempty"` // 该Objective对齐到的Objective列表
}

type BatchGetOKRRespOkrObjectiveProgressRate struct {
	Percent int    `json:"percent,omitempty"` // Objective 进度百分比 >= 0
	Status  string `json:"status,omitempty"`  // Objective 进度状态, 可选值有: `-1`：未更新, `0`：正常, `1`：有风险, `2`：已延期
}

type BatchGetOKRRespOkrObjectiveKr struct {
	ID           string                                     `json:"id,omitempty"`            // Key Result ID
	Content      string                                     `json:"content,omitempty"`       // KeyResult 内容
	Score        int                                        `json:"score,omitempty"`         // KeyResult打分（0 - 100）
	Weight       int                                        `json:"weight,omitempty"`        // KeyResult权重（0 - 100）（废弃）
	KrWeight     float64                                    `json:"kr_weight,omitempty"`     // KeyResult的权重（0 - 100）
	ProgressRate *BatchGetOKRRespOkrObjectiveKrProgressRate `json:"progress_rate,omitempty"` // KR进度
}

type BatchGetOKRRespOkrObjectiveKrProgressRate struct {
	Percent int    `json:"percent,omitempty"` // Objective 进度百分比 >= 0
	Status  string `json:"status,omitempty"`  // Objective 进度状态, 可选值有: `-1`：未更新, `0`：正常, `1`：有风险, `2`：已延期
}

type BatchGetOKRRespOkrObjectiveAlignedObjective struct {
	ID    string                                            `json:"id,omitempty"`     // Objective的ID
	OkrID string                                            `json:"okr_id,omitempty"` // OKR的ID
	Owner *BatchGetOKRRespOkrObjectiveAlignedObjectiveOwner `json:"owner,omitempty"`  // 该Objective的Owner
}

type BatchGetOKRRespOkrObjectiveAlignedObjectiveOwner struct {
	OpenID string `json:"open_id,omitempty"` // 用户的 open_id
}

type BatchGetOKRRespOkrObjectiveAligningObjective struct {
	ID    string                                             `json:"id,omitempty"`     // Objective的ID
	OkrID string                                             `json:"okr_id,omitempty"` // OKR的ID
	Owner *BatchGetOKRRespOkrObjectiveAligningObjectiveOwner `json:"owner,omitempty"`  // 该Objective的Owner
}

type BatchGetOKRRespOkrObjectiveAligningObjectiveOwner struct {
	OpenID string `json:"open_id,omitempty"` // 用户的 open_id
}
