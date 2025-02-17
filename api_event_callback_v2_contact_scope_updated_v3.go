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

// EventV2ContactScopeUpdatedV3 当应用通讯录范围权限发生变更时, 订阅这个事件的应用会收到事件。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=contact&version=v3&resource=scope&event=updated)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/scope/events/updated
func (r *EventCallbackService) HandlerEventV2ContactScopeUpdatedV3(f EventV2ContactScopeUpdatedV3Handler) {
	r.cli.eventHandler.eventV2ContactScopeUpdatedV3Handler = f
}

// EventV2ContactScopeUpdatedV3Handler event EventV2ContactScopeUpdatedV3 handler
type EventV2ContactScopeUpdatedV3Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2ContactScopeUpdatedV3) (string, error)

// EventV2ContactScopeUpdatedV3 ...
type EventV2ContactScopeUpdatedV3 struct {
	Added   *EventV2ContactScopeUpdatedV3Added   `json:"added,omitempty"`   // 当通讯录范围权限变更时, 新增的对象
	Removed *EventV2ContactScopeUpdatedV3Removed `json:"removed,omitempty"` // 当通讯录范围权限发生变更时, 移除的对象
}

// EventV2ContactScopeUpdatedV3Added ...
type EventV2ContactScopeUpdatedV3Added struct {
	Departments []*EventV2ContactScopeUpdatedV3AddedDepartment `json:"departments,omitempty"` // 部门对象
	Users       []*EventV2ContactScopeUpdatedV3AddedUser       `json:"users,omitempty"`       // 用户对象
	UserGroups  []*EventV2ContactScopeUpdatedV3AddedUserGroup  `json:"user_groups,omitempty"` // 用户组对象
}

// EventV2ContactScopeUpdatedV3AddedDepartment ...
type EventV2ContactScopeUpdatedV3AddedDepartment struct {
	Name               string                                               `json:"name,omitempty"`                 // 部门名称, 注意: 不可包含斜杠, 最小长度: `1` 字符, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	I18nName           *EventV2ContactScopeUpdatedV3AddedDepartmentI18nName `json:"i18n_name,omitempty"`            // 国际化的部门名称, 注意: 不可包含斜杠, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	ParentDepartmentID string                                               `json:"parent_department_id,omitempty"` // 父部门的ID, * 在根部门下创建新部门, 该参数值为 “0”, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	DepartmentID       string                                               `json:"department_id,omitempty"`        // 本部门的自定义部门ID, 注意: 除需要满足正则规则外, 同时不能以`od-`开头, 最大长度: `64` 字符, 正则校验: `^0|[^od][A-Za-z0-9]*`, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	OpenDepartmentID   string                                               `json:"open_department_id,omitempty"`   // 部门的open_id, 类型与通过请求的查询参数传入的department_id_type相同
	LeaderUserID       string                                               `json:"leader_user_id,omitempty"`       // 部门主管用户ID, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	ChatID             string                                               `json:"chat_id,omitempty"`              // 部门群ID, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	Order              string                                               `json:"order,omitempty"`                // 部门的排序, 即部门在其同级部门的展示顺序, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	UnitIDs            []string                                             `json:"unit_ids,omitempty"`             // 部门单位自定义ID列表, 当前只支持一个, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	MemberCount        int64                                                `json:"member_count,omitempty"`         // 部门下用户的个数, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	Status             *EventV2ContactScopeUpdatedV3AddedDepartmentStatus   `json:"status,omitempty"`               // 部门状态, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	Leaders            []*EventV2ContactScopeUpdatedV3AddedDepartmentLeader `json:"leaders,omitempty"`              // 部门负责人
}

// EventV2ContactScopeUpdatedV3AddedDepartmentI18nName ...
type EventV2ContactScopeUpdatedV3AddedDepartmentI18nName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 部门的中文名
	JaJp string `json:"ja_jp,omitempty"` // 部门的日文名
	EnUs string `json:"en_us,omitempty"` // 部门的英文名
}

// EventV2ContactScopeUpdatedV3AddedDepartmentLeader ...
type EventV2ContactScopeUpdatedV3AddedDepartmentLeader struct {
	LeaderType int64  `json:"leaderType,omitempty"` // 负责人类型, 可选值有: 1: 主负责人, 2: 副负责人
	LeaderID   string `json:"leaderID,omitempty"`   // 负责人ID, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
}

// EventV2ContactScopeUpdatedV3AddedDepartmentStatus ...
type EventV2ContactScopeUpdatedV3AddedDepartmentStatus struct {
	IsDeleted bool `json:"is_deleted,omitempty"` // 是否被删除
}

// EventV2ContactScopeUpdatedV3AddedUser ...
type EventV2ContactScopeUpdatedV3AddedUser struct {
	UnionID         string                                             `json:"union_id,omitempty"`         // 用户的union_id, 应用开发商发布的不同应用中同一用户的标识, 不同ID的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	UserID          string                                             `json:"user_id,omitempty"`          // 用户的user_id, 租户内用户的唯一标识, 不同ID的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 字段权限要求: 获取用户 user ID
	OpenID          string                                             `json:"open_id,omitempty"`          // 用户的open_id, 应用内用户的唯一标识, 不同ID的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	Name            string                                             `json:"name,omitempty"`             // 用户名, 最小长度: `1` 字符, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
	EnName          string                                             `json:"en_name,omitempty"`          // 英文名, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
	Nickname        string                                             `json:"nickname,omitempty"`         // 别名, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
	Email           string                                             `json:"email,omitempty"`            // 邮箱, 注意: 1. 非中国大陆手机号成员必须同时添加邮箱, 2. 邮箱不可重复, 字段权限要求: 获取用户邮箱信息
	Mobile          string                                             `json:"mobile,omitempty"`           // 手机号, 注意: 1. 在本企业内不可重复, 2. 未认证企业仅支持添加中国大陆手机号, 通过飞书认证的企业允许添加海外手机号, 3. 国际电话区号前缀中必须包含加号 +, 4. 该 mobile 字段在海外版飞书非必填, 字段权限要求: 获取用户手机号
	Gender          int64                                              `json:"gender,omitempty"`           // 性别, 可选值有: 0: 保密, 1: 男, 2: 女, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户性别, 以应用身份访问通讯录, 读取通讯录
	Avatar          *EventV2ContactScopeUpdatedV3AddedUserAvatar       `json:"avatar,omitempty"`           // 用户头像信息, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
	Status          *EventV2ContactScopeUpdatedV3AddedUserStatus       `json:"status,omitempty"`           // 用户状态, 枚举类型, 包括is_frozen、is_resigned、is_activated、is_exited, 用户状态转移参见: [用户状态图](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/field-overview#4302b5a1), 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	LeaderUserID    string                                             `json:"leader_user_id,omitempty"`   // 用户的直接主管的用户ID, ID值与查询参数中的user_id_type 对应, 不同 ID 的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 获取方式参见[如何获取user_id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户组织架构信息, 以应用身份访问通讯录, 读取通讯录
	City            string                                             `json:"city,omitempty"`             // 工作城市, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	Country         string                                             `json:"country,omitempty"`          // 国家或地区Code缩写, 具体写入格式请参考 [国家/地区码表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/country-code-description), 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	WorkStation     string                                             `json:"work_station,omitempty"`     // 工位, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	JoinTime        int64                                              `json:"join_time,omitempty"`        // 入职时间, 时间戳格式, 表示从1970年1月1日开始所经过的秒数, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	EmployeeNo      string                                             `json:"employee_no,omitempty"`      // 工号, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	EmployeeType    int64                                              `json:"employee_type,omitempty"`    // 员工类型, 可选值有: `1`: 正式员工, `2`: 实习生, `3`: 外包, `4`: 劳务, `5`: 顾问, 同时可读取到自定义员工类型的 int 值, 可通过下方接口获取到该租户的自定义员工类型的名称, 参见[获取人员类型](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/list), 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	CustomAttrs     []*EventV2ContactScopeUpdatedV3AddedUserCustomAttr `json:"custom_attrs,omitempty"`     // 自定义字段, 请确保你的组织管理员已在管理后台/组织架构/成员字段管理/自定义字段管理/全局设置中开启了“允许开放平台 API 调用“, 否则该字段不会生效/返回, 更多详情参见[用户接口相关问题](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN#77061525), 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	EnterpriseEmail string                                             `json:"enterprise_email,omitempty"` // 企业邮箱, 请先确保已在管理后台启用飞书邮箱服务, 创建用户时, 企业邮箱的使用方式参见[用户接口相关问题](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN#77061525), 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	JobTitle        string                                             `json:"job_title,omitempty"`        // 职务, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	IsFrozen        bool                                               `json:"is_frozen,omitempty"`        // 是否暂停用户
	Geo             string                                             `json:"geo,omitempty"`              // 数据驻留地
}

// EventV2ContactScopeUpdatedV3AddedUserAvatar ...
type EventV2ContactScopeUpdatedV3AddedUserAvatar struct {
	Avatar72     string `json:"avatar_72,omitempty"`     // 72*72像素头像链接
	Avatar240    string `json:"avatar_240,omitempty"`    // 240*240像素头像链接
	Avatar640    string `json:"avatar_640,omitempty"`    // 640*640像素头像链接
	AvatarOrigin string `json:"avatar_origin,omitempty"` // 原始头像链接
}

// EventV2ContactScopeUpdatedV3AddedUserCustomAttr ...
type EventV2ContactScopeUpdatedV3AddedUserCustomAttr struct {
	Type  string                                                `json:"type,omitempty"`  // 自定义字段类型, `TEXT`: 文本, `HREF`: 网页, `ENUMERATION`: 枚举, `PICTURE_ENUM`: 图片, `GENERIC_USER`: 用户, 具体说明参见常见问题的[用户接口相关问题](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN#77061525)
	ID    string                                                `json:"id,omitempty"`    // 自定义字段ID
	Value *EventV2ContactScopeUpdatedV3AddedUserCustomAttrValue `json:"value,omitempty"` // 自定义字段取值
}

// EventV2ContactScopeUpdatedV3AddedUserCustomAttrValue ...
type EventV2ContactScopeUpdatedV3AddedUserCustomAttrValue struct {
	Text        string                                                           `json:"text,omitempty"`         // 字段类型为`TEXT`时该参数定义字段值, 必填；字段类型为`HREF`时该参数定义网页标题, 必填
	URL         string                                                           `json:"url,omitempty"`          // 字段类型为 HREF 时, 该参数定义默认 URL, 例如手机端跳转小程序, PC端跳转网页
	PcURL       string                                                           `json:"pc_url,omitempty"`       // 字段类型为 HREF 时, 该参数定义PC端 URL
	OptionID    string                                                           `json:"option_id,omitempty"`    // 字段类型为 ENUMERATION 或 PICTURE_ENUM 时, 该参数定义选项值
	OptionValue string                                                           `json:"option_value,omitempty"` // 选项类型的值, 表示 成员详情/自定义字段 中选项选中的值
	Name        string                                                           `json:"name,omitempty"`         // 选项类型为图片时, 表示图片的名称
	PictureURL  string                                                           `json:"picture_url,omitempty"`  // 图片链接
	GenericUser *EventV2ContactScopeUpdatedV3AddedUserCustomAttrValueGenericUser `json:"generic_user,omitempty"` // 字段类型为 GENERIC_USER 时, 该参数定义引用人员
}

// EventV2ContactScopeUpdatedV3AddedUserCustomAttrValueGenericUser ...
type EventV2ContactScopeUpdatedV3AddedUserCustomAttrValueGenericUser struct {
	ID   string `json:"id,omitempty"`   // 用户的user_id, 具体参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	Type int64  `json:"type,omitempty"` // 用户类型: 1: 用户, 目前固定为1, 表示用户类型
}

// EventV2ContactScopeUpdatedV3AddedUserGroup ...
type EventV2ContactScopeUpdatedV3AddedUserGroup struct {
	UserGroupID string `json:"user_group_id,omitempty"` // 用户组的自定义ID, 长度范围: `1` ～ `64` 字符
	Name        string `json:"name,omitempty"`          // 用户组的名称, 长度范围: `1` ～ `100` 字符
	Type        int64  `json:"type,omitempty"`          // 用户组的类型, 可选值有: 1: 普通用户组, 2: 动态用户组
	MemberCount int64  `json:"member_count,omitempty"`  // 成员数量
	Status      int64  `json:"status,omitempty"`        // 用户组状态, 可选值有: 0: 未知, 1: 计算完毕, 2: 计算中, 3: 计算失败
}

// EventV2ContactScopeUpdatedV3AddedUserStatus ...
type EventV2ContactScopeUpdatedV3AddedUserStatus struct {
	IsFrozen    bool `json:"is_frozen,omitempty"`    // 是否暂停
	IsResigned  bool `json:"is_resigned,omitempty"`  // 是否离职
	IsActivated bool `json:"is_activated,omitempty"` // 是否激活
	IsExited    bool `json:"is_exited,omitempty"`    // 是否主动退出, 主动退出一段时间后用户会自动转为已离职
	IsUnjoin    bool `json:"is_unjoin,omitempty"`    // 是否未加入, 需要用户自主确认才能加入团队
}

// EventV2ContactScopeUpdatedV3Removed ...
type EventV2ContactScopeUpdatedV3Removed struct {
	Departments []*EventV2ContactScopeUpdatedV3RemovedDepartment `json:"departments,omitempty"` // 部门对象
	Users       []*EventV2ContactScopeUpdatedV3RemovedUser       `json:"users,omitempty"`       // 用户对象
	UserGroups  []*EventV2ContactScopeUpdatedV3RemovedUserGroup  `json:"user_groups,omitempty"` // 用户组对象
}

// EventV2ContactScopeUpdatedV3RemovedDepartment ...
type EventV2ContactScopeUpdatedV3RemovedDepartment struct {
	Name               string                                                 `json:"name,omitempty"`                 // 部门名称, 注意: 不可包含斜杠, 最小长度: `1` 字符, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	I18nName           *EventV2ContactScopeUpdatedV3RemovedDepartmentI18nName `json:"i18n_name,omitempty"`            // 国际化的部门名称, 注意: 不可包含斜杠, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	ParentDepartmentID string                                                 `json:"parent_department_id,omitempty"` // 父部门的ID, * 在根部门下创建新部门, 该参数值为 “0”, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	DepartmentID       string                                                 `json:"department_id,omitempty"`        // 本部门的自定义部门ID, 注意: 除需要满足正则规则外, 同时不能以`od-`开头, 最大长度: `64` 字符, 正则校验: `^0|[^od][A-Za-z0-9]*`, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	OpenDepartmentID   string                                                 `json:"open_department_id,omitempty"`   // 部门的open_id, 类型与通过请求的查询参数传入的department_id_type相同
	LeaderUserID       string                                                 `json:"leader_user_id,omitempty"`       // 部门主管用户ID, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	ChatID             string                                                 `json:"chat_id,omitempty"`              // 部门群ID, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	Order              string                                                 `json:"order,omitempty"`                // 部门的排序, 即部门在其同级部门的展示顺序, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	UnitIDs            []string                                               `json:"unit_ids,omitempty"`             // 部门单位自定义ID列表, 当前只支持一个, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	MemberCount        int64                                                  `json:"member_count,omitempty"`         // 部门下用户的个数, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	Status             *EventV2ContactScopeUpdatedV3RemovedDepartmentStatus   `json:"status,omitempty"`               // 部门状态, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	Leaders            []*EventV2ContactScopeUpdatedV3RemovedDepartmentLeader `json:"leaders,omitempty"`              // 部门负责人
}

// EventV2ContactScopeUpdatedV3RemovedDepartmentI18nName ...
type EventV2ContactScopeUpdatedV3RemovedDepartmentI18nName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 部门的中文名
	JaJp string `json:"ja_jp,omitempty"` // 部门的日文名
	EnUs string `json:"en_us,omitempty"` // 部门的英文名
}

// EventV2ContactScopeUpdatedV3RemovedDepartmentLeader ...
type EventV2ContactScopeUpdatedV3RemovedDepartmentLeader struct {
	LeaderType int64  `json:"leaderType,omitempty"` // 负责人类型, 可选值有: 1: 主负责人, 2: 副负责人
	LeaderID   string `json:"leaderID,omitempty"`   // 负责人ID, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
}

// EventV2ContactScopeUpdatedV3RemovedDepartmentStatus ...
type EventV2ContactScopeUpdatedV3RemovedDepartmentStatus struct {
	IsDeleted bool `json:"is_deleted,omitempty"` // 是否被删除
}

// EventV2ContactScopeUpdatedV3RemovedUser ...
type EventV2ContactScopeUpdatedV3RemovedUser struct {
	UnionID         string                                               `json:"union_id,omitempty"`         // 用户的union_id, 应用开发商发布的不同应用中同一用户的标识, 不同ID的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	UserID          string                                               `json:"user_id,omitempty"`          // 用户的user_id, 租户内用户的唯一标识, 不同ID的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 字段权限要求: 获取用户 user ID
	OpenID          string                                               `json:"open_id,omitempty"`          // 用户的open_id, 应用内用户的唯一标识, 不同ID的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	Name            string                                               `json:"name,omitempty"`             // 用户名, 最小长度: `1` 字符, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
	EnName          string                                               `json:"en_name,omitempty"`          // 英文名, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
	Nickname        string                                               `json:"nickname,omitempty"`         // 别名, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
	Email           string                                               `json:"email,omitempty"`            // 邮箱, 注意: 1. 非中国大陆手机号成员必须同时添加邮箱, 2. 邮箱不可重复, 字段权限要求: 获取用户邮箱信息
	Mobile          string                                               `json:"mobile,omitempty"`           // 手机号, 注意: 1. 在本企业内不可重复, 2. 未认证企业仅支持添加中国大陆手机号, 通过飞书认证的企业允许添加海外手机号, 3. 国际电话区号前缀中必须包含加号 +, 4. 该 mobile 字段在海外版飞书非必填, 字段权限要求: 获取用户手机号
	Gender          int64                                                `json:"gender,omitempty"`           // 性别, 可选值有: 0: 保密, 1: 男, 2: 女, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户性别, 以应用身份访问通讯录, 读取通讯录
	Avatar          *EventV2ContactScopeUpdatedV3RemovedUserAvatar       `json:"avatar,omitempty"`           // 用户头像信息, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
	Status          *EventV2ContactScopeUpdatedV3RemovedUserStatus       `json:"status,omitempty"`           // 用户状态, 枚举类型, 包括is_frozen、is_resigned、is_activated、is_exited, 用户状态转移参见: [用户状态图](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/field-overview#4302b5a1), 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	LeaderUserID    string                                               `json:"leader_user_id,omitempty"`   // 用户的直接主管的用户ID, ID值与查询参数中的user_id_type 对应, 不同 ID 的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 获取方式参见[如何获取user_id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户组织架构信息, 以应用身份访问通讯录, 读取通讯录
	City            string                                               `json:"city,omitempty"`             // 工作城市, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	Country         string                                               `json:"country,omitempty"`          // 国家或地区Code缩写, 具体写入格式请参考 [国家/地区码表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/country-code-description), 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	WorkStation     string                                               `json:"work_station,omitempty"`     // 工位, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	JoinTime        int64                                                `json:"join_time,omitempty"`        // 入职时间, 时间戳格式, 表示从1970年1月1日开始所经过的秒数, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	EmployeeNo      string                                               `json:"employee_no,omitempty"`      // 工号, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	EmployeeType    int64                                                `json:"employee_type,omitempty"`    // 员工类型, 可选值有: `1`: 正式员工, `2`: 实习生, `3`: 外包, `4`: 劳务, `5`: 顾问, 同时可读取到自定义员工类型的 int 值, 可通过下方接口获取到该租户的自定义员工类型的名称, 参见[获取人员类型](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/list), 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	CustomAttrs     []*EventV2ContactScopeUpdatedV3RemovedUserCustomAttr `json:"custom_attrs,omitempty"`     // 自定义字段, 请确保你的组织管理员已在管理后台/组织架构/成员字段管理/自定义字段管理/全局设置中开启了“允许开放平台 API 调用“, 否则该字段不会生效/返回, 更多详情参见[用户接口相关问题](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN#77061525), 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	EnterpriseEmail string                                               `json:"enterprise_email,omitempty"` // 企业邮箱, 请先确保已在管理后台启用飞书邮箱服务, 创建用户时, 企业邮箱的使用方式参见[用户接口相关问题](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN#77061525), 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	JobTitle        string                                               `json:"job_title,omitempty"`        // 职务, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录
	IsFrozen        bool                                                 `json:"is_frozen,omitempty"`        // 是否暂停用户
	Geo             string                                               `json:"geo,omitempty"`              // 数据驻留地
}

// EventV2ContactScopeUpdatedV3RemovedUserAvatar ...
type EventV2ContactScopeUpdatedV3RemovedUserAvatar struct {
	Avatar72     string `json:"avatar_72,omitempty"`     // 72*72像素头像链接
	Avatar240    string `json:"avatar_240,omitempty"`    // 240*240像素头像链接
	Avatar640    string `json:"avatar_640,omitempty"`    // 640*640像素头像链接
	AvatarOrigin string `json:"avatar_origin,omitempty"` // 原始头像链接
}

// EventV2ContactScopeUpdatedV3RemovedUserCustomAttr ...
type EventV2ContactScopeUpdatedV3RemovedUserCustomAttr struct {
	Type  string                                                  `json:"type,omitempty"`  // 自定义字段类型, `TEXT`: 文本, `HREF`: 网页, `ENUMERATION`: 枚举, `PICTURE_ENUM`: 图片, `GENERIC_USER`: 用户, 具体说明参见常见问题的[用户接口相关问题](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN#77061525)
	ID    string                                                  `json:"id,omitempty"`    // 自定义字段ID
	Value *EventV2ContactScopeUpdatedV3RemovedUserCustomAttrValue `json:"value,omitempty"` // 自定义字段取值
}

// EventV2ContactScopeUpdatedV3RemovedUserCustomAttrValue ...
type EventV2ContactScopeUpdatedV3RemovedUserCustomAttrValue struct {
	Text        string                                                             `json:"text,omitempty"`         // 字段类型为`TEXT`时该参数定义字段值, 必填；字段类型为`HREF`时该参数定义网页标题, 必填
	URL         string                                                             `json:"url,omitempty"`          // 字段类型为 HREF 时, 该参数定义默认 URL, 例如手机端跳转小程序, PC端跳转网页
	PcURL       string                                                             `json:"pc_url,omitempty"`       // 字段类型为 HREF 时, 该参数定义PC端 URL
	OptionID    string                                                             `json:"option_id,omitempty"`    // 字段类型为 ENUMERATION 或 PICTURE_ENUM 时, 该参数定义选项值
	OptionValue string                                                             `json:"option_value,omitempty"` // 选项类型的值, 表示 成员详情/自定义字段 中选项选中的值
	Name        string                                                             `json:"name,omitempty"`         // 选项类型为图片时, 表示图片的名称
	PictureURL  string                                                             `json:"picture_url,omitempty"`  // 图片链接
	GenericUser *EventV2ContactScopeUpdatedV3RemovedUserCustomAttrValueGenericUser `json:"generic_user,omitempty"` // 字段类型为 GENERIC_USER 时, 该参数定义引用人员
}

// EventV2ContactScopeUpdatedV3RemovedUserCustomAttrValueGenericUser ...
type EventV2ContactScopeUpdatedV3RemovedUserCustomAttrValueGenericUser struct {
	ID   string `json:"id,omitempty"`   // 用户的user_id, 具体参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	Type int64  `json:"type,omitempty"` // 用户类型: 1: 用户, 目前固定为1, 表示用户类型
}

// EventV2ContactScopeUpdatedV3RemovedUserGroup ...
type EventV2ContactScopeUpdatedV3RemovedUserGroup struct {
	UserGroupID string `json:"user_group_id,omitempty"` // 用户组的自定义ID, 长度范围: `1` ～ `64` 字符
	Name        string `json:"name,omitempty"`          // 用户组的名称, 长度范围: `1` ～ `100` 字符
	Type        int64  `json:"type,omitempty"`          // 用户组的类型, 可选值有: 1: 普通用户组, 2: 动态用户组
	MemberCount int64  `json:"member_count,omitempty"`  // 成员数量
	Status      int64  `json:"status,omitempty"`        // 用户组状态, 可选值有: 0: 未知, 1: 计算完毕, 2: 计算中, 3: 计算失败
}

// EventV2ContactScopeUpdatedV3RemovedUserStatus ...
type EventV2ContactScopeUpdatedV3RemovedUserStatus struct {
	IsFrozen    bool `json:"is_frozen,omitempty"`    // 是否暂停
	IsResigned  bool `json:"is_resigned,omitempty"`  // 是否离职
	IsActivated bool `json:"is_activated,omitempty"` // 是否激活
	IsExited    bool `json:"is_exited,omitempty"`    // 是否主动退出, 主动退出一段时间后用户会自动转为已离职
	IsUnjoin    bool `json:"is_unjoin,omitempty"`    // 是否未加入, 需要用户自主确认才能加入团队
}
