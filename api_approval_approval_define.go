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

// CreateApproval 用于通过接口创建简单的审批定义, 可以灵活指定定义的基础信息、表单和流程等。创建成功后, 不支持从审批管理后台删除该定义。不推荐企业自建应用使用, 如有需要尽量联系管理员在审批管理后台创建定义。
//
// 接口谨慎调用, 创建后的审批定义无法停用/删除
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/create
func (r *ApprovalService) CreateApproval(ctx context.Context, request *CreateApprovalReq, options ...MethodOptionFunc) (*CreateApprovalResp, *Response, error) {
	if r.cli.mock.mockApprovalCreateApproval != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#CreateApproval mock enable")
		return r.cli.mock.mockApprovalCreateApproval(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "CreateApproval",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/approval/v4/approvals",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createApprovalResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalCreateApproval mock ApprovalCreateApproval method
func (r *Mock) MockApprovalCreateApproval(f func(ctx context.Context, request *CreateApprovalReq, options ...MethodOptionFunc) (*CreateApprovalResp, *Response, error)) {
	r.mockApprovalCreateApproval = f
}

// UnMockApprovalCreateApproval un-mock ApprovalCreateApproval method
func (r *Mock) UnMockApprovalCreateApproval() {
	r.mockApprovalCreateApproval = nil
}

// CreateApprovalReq ...
type CreateApprovalReq struct {
	DepartmentIDType  *DepartmentIDType                `query:"department_id_type" json:"-"`  // 此次调用中使用的部门ID的类型, 示例值: "open_department_id", 可选值有: department_id: 以自定义department_id来标识部门, open_department_id: 以open_department_id来标识部门
	UserIDType        *IDType                          `query:"user_id_type" json:"-"`        // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ApprovalName      string                           `json:"approval_name,omitempty"`       // 审批名称的国际化文案 Key, 以 @i18n@ 开头, 长度不得少于 9 个字符, 示例值: "@i18n@approval_name"
	ApprovalCode      *string                          `json:"approval_code,omitempty"`       // 传空表示新建, 示例值: "7C468A54-8745-2245-9675-08B7C63E7A85"
	Description       *string                          `json:"description,omitempty"`         // 审批描述的国际化文案 Key, 以 @i18n@ 开头, 长度不得少于 9 个字符, 示例值: "@i18n@description"
	Viewers           []*CreateApprovalReqViewer       `json:"viewers,omitempty"`             // viewers 字段指定了哪些人能从审批应用的前台发起该审批, 1. 当 view_type 为 USER, 需要填写viewer_user_id；, 2. 当 view_type 为DEPARTMENT, 需要填写viewer_department_id；, 3. 当 view_type 为TENANT或NONE时, viewer_user_id和viewer_department_id无需填写
	Form              *CreateApprovalReqForm           `json:"form,omitempty"`                // 审批定义表单
	NodeList          []*CreateApprovalReqNode         `json:"node_list,omitempty"`           // 审批定义节点, 需要将开始节点作为 list 第一个元素, 结束节点作为最后一个元素
	Settings          *CreateApprovalReqSettings       `json:"settings,omitempty"`            // 审批定义其他设置
	Config            *CreateApprovalReqConfig         `json:"config,omitempty"`              // 审批定义配置项, 用于配置对应审批定义是否可以由用户在审批后台进行修改
	Icon              *int64                           `json:"icon,omitempty"`                // 审批图标枚举, 详见下方说明, 默认为 0, 示例值: 0, 默认值: `0`
	I18nResources     []*CreateApprovalReqI18nResource `json:"i18n_resources,omitempty"`      // 国际化文案
	ProcessManagerIDs []string                         `json:"process_manager_ids,omitempty"` // 根据user_id_type填写流程管理员的用户id, 示例值: ["1c5ea995"]
}

// CreateApprovalReqConfig ...
type CreateApprovalReqConfig struct {
	CanUpdateViewer  bool   `json:"can_update_viewer,omitempty"`  // 允许用户修改可见范围, 示例值: false
	CanUpdateForm    bool   `json:"can_update_form,omitempty"`    // 允许用户更新表单, 示例值: false
	CanUpdateProcess bool   `json:"can_update_process,omitempty"` // 允许用户更新流程定义, 示例值: false
	CanUpdateRevert  bool   `json:"can_update_revert,omitempty"`  // 允许用户更新撤回设置, 示例值: false
	HelpURL          string `json:"help_url,omitempty"`           // 帮助文档链接, 示例值: "https://www.baidu.com"
}

// CreateApprovalReqForm ...
type CreateApprovalReqForm struct {
	FormContent string `json:"form_content,omitempty"` // 审批定义表单, json 数组, 见下方form_content字段说明, 示例值: "[{\"id\":\"user_name\", \"type\": \"input\", \"required\":true, \"name\":\"@i18n@widget1\"}]"
}

// CreateApprovalReqI18nResource ...
type CreateApprovalReqI18nResource struct {
	Locale    string                               `json:"locale,omitempty"`     // 语言可选值有: zh-CN: 中文 en-US: 英文 ja-JP: 日文, 示例值: "zh-CN", 可选值有: zh-CN: 中文, en-US: 英文, ja-JP: 日文
	Texts     []*CreateApprovalReqI18nResourceText `json:"texts,omitempty"`      // 文案 key, value, i18n key 以 @i18n@ 开头； 该字段主要用于做国际化, 语序用户同时传多个语言的文案, 审批中心会根据用户当前的语音环境使用对应的文案, 如果没有传用户当前的语音环境文案, 则会使用默认的语言文案。
	IsDefault bool                                 `json:"is_default,omitempty"` // 是否默认语言, 默认语言需要包含所有key, 非默认语言如果key不存在会使用默认语言代替, 示例值: true
}

// CreateApprovalReqI18nResourceText ...
type CreateApprovalReqI18nResourceText struct {
	Key   string `json:"key,omitempty"`   // 文案key, 示例值: "@i18n@1"
	Value string `json:"value,omitempty"` // 文案, 示例值: "people"
}

// CreateApprovalReqNode ...
type CreateApprovalReqNode struct {
	ID             string                               `json:"id,omitempty"`              // 节点 ID, 开始节点的 ID 为 START, 结束节点的 ID 为 END, 开始和结束节点不需要指定 name、node_type 以及 approver, 示例值: "START"
	Name           *string                              `json:"name,omitempty"`            // 节点名称的国际化文案 Key, 以 @i18n@ 开头, 长度不得少于 9 个字符, 示例值: "@i18n@node_name"
	NodeType       *string                              `json:"node_type,omitempty"`       // 审批类型枚举, 当 node_type 为依次审批时, 审批人必须为『发起人自选』, 示例值: "AND", 可选值有: AND: 会签, OR: 或签, SEQUENTIAL: 依次审批
	Approver       []*CreateApprovalReqNodeApprover     `json:"approver,omitempty"`        // 审批人列表
	Ccer           []*CreateApprovalReqNodeCcer         `json:"ccer,omitempty"`            // 抄送人列表
	PrivilegeField *CreateApprovalReqNodePrivilegeField `json:"privilege_field,omitempty"` // 表单项的控件权限
}

// CreateApprovalReqNodeApprover ...
type CreateApprovalReqNodeApprover struct {
	Type   string  `json:"type,omitempty"`    // 审批/抄送人类型, 1. 当 type 为 Supervisor、SupervisorTopDown、DepartmentManager 、DepartmentManagerTopDown 这 4 种时, 需要在 level 中填写对应的级数, 例如: 由下往上三级主管审批, level = 3；, 2. 当 type 为 Personal 时, 需要填写对应的user_id, 用于指定用户；, 3. 当 approver 为 Free 发起人自选时, 不需要指定 user_id 和level；, 4. ccer不支持 Free 发起人自选, 示例值: "Supervisor", 可选值有: Supervisor: 主管审批（由下往上）, SupervisorTopDown: 主管审批（从上往下）, DepartmentManager: 部门负责人审批（由下往上）, DepartmentManagerTopDown: 部门负责人审批（从上往下）, Personal: 指定成员, Free: 发起人自选
	UserID *string `json:"user_id,omitempty"` // 用户id, 根据user_id_type填写, 示例值: "f7cb567e"
	Level  *string `json:"level,omitempty"`   // 审批级数, 当 type 为 Supervisor、SupervisorTopDown、DepartmentManager 、DepartmentManagerTopDown 这 4 种时, 需要在 level 中填写对应的级数, 例如: 由下往上三级主管审批, level = 3, 示例值: "3"
}

// CreateApprovalReqNodeCcer ...
type CreateApprovalReqNodeCcer struct {
	Type   string  `json:"type,omitempty"`    // 审批/抄送人类型, 1. 当 type 为 Supervisor、SupervisorTopDown、DepartmentManager 、DepartmentManagerTopDown 这 4 种时, 需要在 level 中填写对应的级数, 例如: 由下往上三级主管审批, level = 3；, 2. 当 type 为 Personal 时, 需要填写对应的user_id, 用于指定用户；, 3. 当 approver 为 Free 发起人自选时, 不需要指定 user_id 和level；, 4. ccer不支持 Free 发起人自选, 示例值: "Supervisor", 可选值有: Supervisor: 主管审批（由下往上）, SupervisorTopDown: 主管审批（从上往下）, DepartmentManager: 部门负责人审批（由下往上）, DepartmentManagerTopDown: 部门负责人审批（从上往下）, Personal: 指定成员, Free: 发起人自选
	UserID *string `json:"user_id,omitempty"` // 用户id, 根据user_id_type填写, 示例值: "f7cb567e"
	Level  *string `json:"level,omitempty"`   // 审批级数, 当 type 为 Supervisor、SupervisorTopDown、DepartmentManager 、DepartmentManagerTopDown 这 4 种时, 需要在 level 中填写对应的级数, 例如: 由下往上三级主管审批, level = 3, 示例值: "3"
}

// CreateApprovalReqNodePrivilegeField ...
type CreateApprovalReqNodePrivilegeField struct {
	Writable []string `json:"writable,omitempty"` // 可写权限的表单项的 id列表, 示例值: 9293493
	Readable []string `json:"readable,omitempty"` // 可读权限的表单项的 id列表, 示例值: 9293493
}

// CreateApprovalReqSettings ...
type CreateApprovalReqSettings struct {
	RevertInterval *int64 `json:"revert_interval,omitempty"` // 审批实例通过后允许撤回的时间, 以秒为单位, 默认 31 天, 0 为不可撤回, 示例值: 0
	RevertOption   *int64 `json:"revert_option,omitempty"`   // 是否支持审批通过第一个节点后撤回, 默认为1, 0为不支持, 示例值: 0
}

// CreateApprovalReqViewer ...
type CreateApprovalReqViewer struct {
	ViewerType         *string `json:"viewer_type,omitempty"`          // 可见人类型, 示例值: "USER", 可选值有: TENANT: 租户内可见, DEPARTMENT: 指定部门, USER: 指定用户, NONE: 任何人都不可见
	ViewerUserID       *string `json:"viewer_user_id,omitempty"`       // 当 view_type 是 USER, 根据user_id_type填写用户id, 示例值: "19a294c2"
	ViewerDepartmentID *string `json:"viewer_department_id,omitempty"` // 当 view_type 为DEPARTMENT, 根据department_id_type填写部门id, 示例值: "od-ac9d697abfa990b715dcc33d58a62a9d"
}

// CreateApprovalResp ...
type CreateApprovalResp struct {
	ApprovalCode string `json:"approval_code,omitempty"` // 审批定义 Code
	ApprovalID   string `json:"approval_id,omitempty"`   // 审批定义 id
}

// createApprovalResp ...
type createApprovalResp struct {
	Code int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *CreateApprovalResp `json:"data,omitempty"`
}
