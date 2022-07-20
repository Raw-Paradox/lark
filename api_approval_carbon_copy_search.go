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

// SearchApprovalCarbonCopy 该接口通过不同条件查询审批系统中符合条件的审批抄送列表。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/search_cc
func (r *ApprovalService) SearchApprovalCarbonCopy(ctx context.Context, request *SearchApprovalCarbonCopyReq, options ...MethodOptionFunc) (*SearchApprovalCarbonCopyResp, *Response, error) {
	if r.cli.mock.mockApprovalSearchApprovalCarbonCopy != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#SearchApprovalCarbonCopy mock enable")
		return r.cli.mock.mockApprovalSearchApprovalCarbonCopy(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "SearchApprovalCarbonCopy",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/approval/v4/instances/search_cc",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(searchApprovalCarbonCopyResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalSearchApprovalCarbonCopy mock ApprovalSearchApprovalCarbonCopy method
func (r *Mock) MockApprovalSearchApprovalCarbonCopy(f func(ctx context.Context, request *SearchApprovalCarbonCopyReq, options ...MethodOptionFunc) (*SearchApprovalCarbonCopyResp, *Response, error)) {
	r.mockApprovalSearchApprovalCarbonCopy = f
}

// UnMockApprovalSearchApprovalCarbonCopy un-mock ApprovalSearchApprovalCarbonCopy method
func (r *Mock) UnMockApprovalSearchApprovalCarbonCopy() {
	r.mockApprovalSearchApprovalCarbonCopy = nil
}

// SearchApprovalCarbonCopyReq ...
type SearchApprovalCarbonCopyReq struct {
	PageSize           *int64  `query:"page_size" json:"-"`            // 分页大小, 示例值: 10, 默认值: `10`, 取值范围: `5` ～ `200`
	PageToken          *string `query:"page_token" json:"-"`           // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "nF1ZXJ5VGhlbkZldGNoCgAAAAAA6PZwFmUzSldvTC1yU"
	UserIDType         *IDType `query:"user_id_type" json:"-"`         // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	UserID             string  `json:"user_id,omitempty"`              // 根据x_user_type填写用户 id, 示例值: "lwiu098wj"
	ApprovalCode       *string `json:"approval_code,omitempty"`        // 审批定义 code, 示例值: "EB828003-9FFE-4B3F-AA50-2E199E2ED942"
	InstanceCode       *string `json:"instance_code,omitempty"`        // 审批实例 code, 示例值: "EB828003-9FFE-4B3F-AA50-2E199E2ED943"
	InstanceExternalID *string `json:"instance_external_id,omitempty"` // 审批实例第三方 id 注: 和 approval_code 取并集, 示例值: "EB828003-9FFE-4B3F-AA50-2E199E2ED976"
	GroupExternalID    *string `json:"group_external_id,omitempty"`    // 审批定义分组第三方 id 注: 和 instance_code 取并集, 示例值: "1234567"
	CcTitle            *string `json:"cc_title,omitempty"`             // 审批实例标题（只有第三方审批有）, 示例值: "test"
	ReadStatus         *string `json:"read_status,omitempty"`          // 审批抄送状态, 注: 若不设置, 查询全部状态 若不在集合中, 报错, 示例值: "read", 可选值有: READ: 已读, UNREAD: 未读, ALL: 所有状态
	CcCreateTimeFrom   *string `json:"cc_create_time_from,omitempty"`  // 实例查询开始时间（unix毫秒时间戳）, 示例值: "1547654251506"
	CcCreateTimeTo     *string `json:"cc_create_time_to,omitempty"`    // 实例查询结束时间 (unix毫秒时间戳), 示例值: "1547654251506"
	Locale             *string `json:"locale,omitempty"`               // 地区, 示例值: "zh-CN", 可选值有: zh-CN: 中文, en-US: 英文, ja-JP: 日文
}

// SearchApprovalCarbonCopyResp ...
type SearchApprovalCarbonCopyResp struct {
	Count     int64                             `json:"count,omitempty"`      // 查询返回条数
	CcList    []*SearchApprovalCarbonCopyRespCc `json:"cc_list,omitempty"`    // 审批实例列表
	PageToken string                            `json:"page_token,omitempty"` // 翻页 Token
	HasMore   bool                              `json:"has_more,omitempty"`   // 是否有更多任务可供拉取
}

// SearchApprovalCarbonCopyRespCc ...
type SearchApprovalCarbonCopyRespCc struct {
	Approval *SearchApprovalCarbonCopyRespCcApproval `json:"approval,omitempty"` // 审批定义
	Group    *SearchApprovalCarbonCopyRespCcGroup    `json:"group,omitempty"`    // 审批定义分组
	Instance *SearchApprovalCarbonCopyRespCcInstance `json:"instance,omitempty"` // 审批实例信息
	Cc       *SearchApprovalCarbonCopyRespCcCc       `json:"cc,omitempty"`       // 审批任务
}

// SearchApprovalCarbonCopyRespCcApproval ...
type SearchApprovalCarbonCopyRespCcApproval struct {
	Code       string                                          `json:"code,omitempty"`        // 审批定义 code
	Name       string                                          `json:"name,omitempty"`        // 审批定义名称
	IsExternal bool                                            `json:"is_external,omitempty"` // 是否为第三方审批
	External   *SearchApprovalCarbonCopyRespCcApprovalExternal `json:"external,omitempty"`    // 第三方审批信息
}

// SearchApprovalCarbonCopyRespCcApprovalExternal ...
type SearchApprovalCarbonCopyRespCcApprovalExternal struct {
	BatchCcRead bool `json:"batch_cc_read,omitempty"` // 是否支持批量读
}

// SearchApprovalCarbonCopyRespCcCc ...
type SearchApprovalCarbonCopyRespCcCc struct {
	UserID     string                                `json:"user_id,omitempty"`     // 审批抄送发起人 id
	CreateTime string                                `json:"create_time,omitempty"` // 审批实例开始时间
	ReadStatus string                                `json:"read_status,omitempty"` // 审批实例状态, 可选值有: read: 已读, unread: 未读
	Title      string                                `json:"title,omitempty"`       // 审批实例名称（只有第三方审批有）
	Extra      string                                `json:"extra,omitempty"`       // 审批实例扩展字段, string型json
	Link       *SearchApprovalCarbonCopyRespCcCcLink `json:"link,omitempty"`        // 审批实例链接（只有第三方审批有）
}

// SearchApprovalCarbonCopyRespCcCcLink ...
type SearchApprovalCarbonCopyRespCcCcLink struct {
	PcLink     string `json:"pc_link,omitempty"`     // 审批实例 pc 端链接
	MobileLink string `json:"mobile_link,omitempty"` // 审批实例移动端链接
}

// SearchApprovalCarbonCopyRespCcGroup ...
type SearchApprovalCarbonCopyRespCcGroup struct {
	ExternalID string `json:"external_id,omitempty"` // 审批定义分组外部 id
	Name       string `json:"name,omitempty"`        // 审批定义分组名称
}

// SearchApprovalCarbonCopyRespCcInstance ...
type SearchApprovalCarbonCopyRespCcInstance struct {
	Code       string                                      `json:"code,omitempty"`        // 审批实例 code
	ExternalID string                                      `json:"external_id,omitempty"` // 审批实例外部 id
	UserID     string                                      `json:"user_id,omitempty"`     // 审批实例发起人 id
	StartTime  string                                      `json:"start_time,omitempty"`  // 审批实例开始时间
	EndTime    string                                      `json:"end_time,omitempty"`    // 审批实例结束时间
	Status     string                                      `json:"status,omitempty"`      // 审批实例状态, 可选值有: reject: 拒绝, pending: 审批中, recall: 撤回, deleted: 已删除, approved: 通过
	Title      string                                      `json:"title,omitempty"`       // 审批实例名称（只有第三方审批有）
	Extra      string                                      `json:"extra,omitempty"`       // 审批实例扩展字段, string型json
	SerialID   string                                      `json:"serial_id,omitempty"`   // 审批流水号
	Link       *SearchApprovalCarbonCopyRespCcInstanceLink `json:"link,omitempty"`        // 审批实例链接（只有第三方审批有）
}

// SearchApprovalCarbonCopyRespCcInstanceLink ...
type SearchApprovalCarbonCopyRespCcInstanceLink struct {
	PcLink     string `json:"pc_link,omitempty"`     // 审批实例 pc 端链接
	MobileLink string `json:"mobile_link,omitempty"` // 审批实例移动端链接
}

// searchApprovalCarbonCopyResp ...
type searchApprovalCarbonCopyResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *SearchApprovalCarbonCopyResp `json:"data,omitempty"`
}
