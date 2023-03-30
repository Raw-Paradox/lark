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

// UpdateAdminBadgeGrant 通过该接口可以修改特定授予名单的相关信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/badge-grant/update
func (r *AdminService) UpdateAdminBadgeGrant(ctx context.Context, request *UpdateAdminBadgeGrantReq, options ...MethodOptionFunc) (*UpdateAdminBadgeGrantResp, *Response, error) {
	if r.cli.mock.mockAdminUpdateAdminBadgeGrant != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Admin#UpdateAdminBadgeGrant mock enable")
		return r.cli.mock.mockAdminUpdateAdminBadgeGrant(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Admin",
		API:                   "UpdateAdminBadgeGrant",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/admin/v1/badges/:badge_id/grants/:grant_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateAdminBadgeGrantResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAdminUpdateAdminBadgeGrant mock AdminUpdateAdminBadgeGrant method
func (r *Mock) MockAdminUpdateAdminBadgeGrant(f func(ctx context.Context, request *UpdateAdminBadgeGrantReq, options ...MethodOptionFunc) (*UpdateAdminBadgeGrantResp, *Response, error)) {
	r.mockAdminUpdateAdminBadgeGrant = f
}

// UnMockAdminUpdateAdminBadgeGrant un-mock AdminUpdateAdminBadgeGrant method
func (r *Mock) UnMockAdminUpdateAdminBadgeGrant() {
	r.mockAdminUpdateAdminBadgeGrant = nil
}

// UpdateAdminBadgeGrantReq ...
type UpdateAdminBadgeGrantReq struct {
	BadgeID          string                              `path:"badge_id" json:"-"`            // 勋章ID, 示例值: "m_DjMzaK", 长度范围: `1` ～ `64` 字符
	GrantID          string                              `path:"grant_id" json:"-"`            // 授予名单ID, 示例值: "g_uS4yux", 长度范围: `1` ～ `64` 字符
	UserIDType       *IDType                             `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType                   `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 示例值: "open_department_id", 可选值有: department_id: 以自定义department_id来标识部门, open_department_id: 以open_department_id来标识部门, 默认值: `open_department_id`
	Name             string                              `json:"name,omitempty"`               // 勋章下唯一的授予事项, 最多100个字符, 示例值: "激励勋章的授予名单"
	GrantType        int64                               `json:"grant_type,omitempty"`         // 授予名单类型, 示例值: 0, 可选值有: 0: 手动选择有效期, 1: 匹配系统入职时间, 默认值: `0`, 取值范围: `0` ～ `1`
	TimeZone         string                              `json:"time_zone,omitempty"`          // 授予名单的生效时间对应的时区, 用于检查RuleDetail的时间戳的取值是否规范, 取值范围为TZ database name, 示例值: "Asia/Shanghai", 最小长度: `1` 字符
	RuleDetail       *UpdateAdminBadgeGrantReqRuleDetail `json:"rule_detail,omitempty"`        // 规则详情
	IsGrantAll       bool                                `json:"is_grant_all,omitempty"`       // 是否授予给全员。1.为false时, 需要关联1~500个用户群体。2.为true时, 不可关联用户、用户组、部门, 示例值: false, 默认值: `false`
	UserIDs          []string                            `json:"user_ids,omitempty"`           // 授予的用户ID列表, 授予名单列表接口返回结果中不返回该字段, 只在详情接口返回, 示例值: ["u273y71"]
	DepartmentIDs    []string                            `json:"department_ids,omitempty"`     // 授予的部门ID列表, 授予名单列表接口返回结果中不返回该字段, 只在详情接口返回, 示例值: ["h121921"]
	GroupIDs         []string                            `json:"group_ids,omitempty"`          // 授予的用户组ID列表, 授予名单列表接口返回结果中不返回该字段, 只在详情接口返回, 示例值: ["g122817"]
}

// UpdateAdminBadgeGrantReqRuleDetail ...
type UpdateAdminBadgeGrantReqRuleDetail struct {
	EffectiveTime   *string `json:"effective_time,omitempty"`   // 开始生效的时间戳。1.手动设置有效期类型勋章, 配置有效期限需要配置该字段；2.时间戳必须是所在时区当天的零点时间戳, 如时区为Asia/Shanghai时区时的1649606400, 示例值: "1649606400"
	ExpirationTime  *string `json:"expiration_time,omitempty"`  // 结束生效的时间戳。1.手动设置有效期类型勋章, 配置有效期限需要配置该字段；2.最大值: 不得超过effective_time+100 年；3.非永久有效: 时间戳必须是所在时区当天的23:59:59时间戳, 如时区为Asia/Shanghai时区时的1649692799；4.永久有效: 传值为0即可, 示例值: "1649692799"
	Anniversary     *int64  `json:"anniversary,omitempty"`      // 入职周年日。根据入职时间发放类型勋章, 需要配置该字段, 示例值: 1, 默认值: `1`, 取值范围: `1` ～ `60`
	EffectivePeriod *int64  `json:"effective_period,omitempty"` // 有效期限。根据入职时间发放类型勋章, 需要配置该字段, 示例值: 1, 可选值有: 1: 有效期为一年, 2: 永久有效, 默认值: `1`, 取值范围: `1` ～ `2`
}

// UpdateAdminBadgeGrantResp ...
type UpdateAdminBadgeGrantResp struct {
	Grant *UpdateAdminBadgeGrantRespGrant `json:"grant,omitempty"` // 授予名单
}

// UpdateAdminBadgeGrantRespGrant ...
type UpdateAdminBadgeGrantRespGrant struct {
	ID            string                                    `json:"id,omitempty"`             // 租户内授予名单的唯一标识, 该值由系统随机生成。
	BadgeID       string                                    `json:"badge_id,omitempty"`       // 企业勋章的唯一ID
	Name          string                                    `json:"name,omitempty"`           // 勋章下唯一的授予事项, 最多100个字符。
	GrantType     int64                                     `json:"grant_type,omitempty"`     // 授予名单类型, 可选值有: 0: 手动选择有效期, 1: 匹配系统入职时间
	TimeZone      string                                    `json:"time_zone,omitempty"`      // 授予名单的生效时间对应的时区, 用于检查RuleDetail的时间戳的取值是否规范, 取值范围为TZ database name
	RuleDetail    *UpdateAdminBadgeGrantRespGrantRuleDetail `json:"rule_detail,omitempty"`    // 规则详情
	IsGrantAll    bool                                      `json:"is_grant_all,omitempty"`   // 是否授予给全员。1.为false时, 需要关联1~500个用户群体。2.为true时, 不可关联用户、用户组、部门。
	UserIDs       []string                                  `json:"user_ids,omitempty"`       // 授予的用户ID列表, 授予名单列表接口返回结果中不返回该字段, 只在详情接口返回
	DepartmentIDs []string                                  `json:"department_ids,omitempty"` // 授予的部门ID列表, 授予名单列表接口返回结果中不返回该字段, 只在详情接口返回
	GroupIDs      []string                                  `json:"group_ids,omitempty"`      // 授予的用户组ID列表, 授予名单列表接口返回结果中不返回该字段, 只在详情接口返回
}

// UpdateAdminBadgeGrantRespGrantRuleDetail ...
type UpdateAdminBadgeGrantRespGrantRuleDetail struct {
	EffectiveTime   string `json:"effective_time,omitempty"`   // 开始生效的时间戳。1.手动设置有效期类型勋章, 配置有效期限需要配置该字段；2.时间戳必须是所在时区当天的零点时间戳, 如时区为Asia/Shanghai时区时的1649606400
	ExpirationTime  string `json:"expiration_time,omitempty"`  // 结束生效的时间戳。1.手动设置有效期类型勋章, 配置有效期限需要配置该字段；2.最大值: 不得超过effective_time+100 年；3.非永久有效: 时间戳必须是所在时区当天的23:59:59时间戳, 如时区为Asia/Shanghai时区时的1649692799；4.永久有效: 传值为0即可
	Anniversary     int64  `json:"anniversary,omitempty"`      // 入职周年日。根据入职时间发放类型勋章, 需要配置该字段。
	EffectivePeriod int64  `json:"effective_period,omitempty"` // 有效期限。根据入职时间发放类型勋章, 需要配置该字段, 可选值有: 1: 有效期为一年, 2: 永久有效
}

// updateAdminBadgeGrantResp ...
type updateAdminBadgeGrantResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *UpdateAdminBadgeGrantResp `json:"data,omitempty"`
}
