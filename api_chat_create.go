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

// CreateChat 创建群并设置群头像、群名、群描述等。
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 本接口支持在创建群的同时拉用户或机器人进群；如果仅需要拉用户或者机器人入群参考 [将用户或机器人拉入群聊](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/create)接口
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/create
func (r *ChatService) CreateChat(ctx context.Context, request *CreateChatReq, options ...MethodOptionFunc) (*CreateChatResp, *Response, error) {
	if r.cli.mock.mockChatCreateChat != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#CreateChat mock enable")
		return r.cli.mock.mockChatCreateChat(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "CreateChat",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createChatResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatCreateChat mock ChatCreateChat method
func (r *Mock) MockChatCreateChat(f func(ctx context.Context, request *CreateChatReq, options ...MethodOptionFunc) (*CreateChatResp, *Response, error)) {
	r.mockChatCreateChat = f
}

// UnMockChatCreateChat un-mock ChatCreateChat method
func (r *Mock) UnMockChatCreateChat() {
	r.mockChatCreateChat = nil
}

// CreateChatReq ...
type CreateChatReq struct {
	UserIDType             *IDType                             `query:"user_id_type" json:"-"`             // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	SetBotManager          *bool                               `query:"set_bot_manager" json:"-"`          // 如果在请求体的 [owner_id] 字段指定了某个用户为群主, 可以选择是否同时设置创建此群的机器人为管理员, 此标志位用于标记是否设置创建群的机器人为管理员, 示例值: false, 默认值: `false`
	UUID                   *string                             `query:"uuid" json:"-"`                     // 由开发者生成的唯一字符串序列, 用于创建群组请求去重；持有相同uuid的请求10小时内只可成功创建1个群聊, 示例值: "b13g2t38-1jd2-458b-8djf-dtbca5104204", 最大长度: `50` 字符
	Avatar                 *string                             `json:"avatar,omitempty"`                   // 群头像对应的 Image Key, 可通过[上传图片](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create)获取（注意: 上传图片的 [image_type] 需要指定为 [avatar]）, 示例值: "default-avatar_44ae0ca3-e140-494b-956f-78091e348435"
	Name                   *string                             `json:"name,omitempty"`                     // 群名称, 注意: 公开群名称的长度不得少于2个字符, 私有群若未填写群名称, 群名称默认设置为 ”`(无主题)`“, 示例值: "测试群名称"
	Description            *string                             `json:"description,omitempty"`              // 群描述, 示例值: "测试群描述"
	I18nNames              *I18nNames                          `json:"i18n_names,omitempty"`               // 群国际化名称
	OwnerID                *string                             `json:"owner_id,omitempty"`                 // 创建群时指定的群主, 不填时指定建群的机器人为群主。群主 ID值应与查询参数中的 [user_id_type] 对应；推荐使用 OpenID, 获取方式可参考文档[如何获取 Open ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), 示例值: "ou_7d8a6e6df7621556ce0d21922b676706ccs"
	UserIDList             []string                            `json:"user_id_list,omitempty"`             // 创建群时邀请的群成员, ID 类型在查询参数 [user_id_type] 中指定；推荐使用 OpenID, 获取方式可参考文档[如何获取 Open ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), 注意: 最多同时邀请 50 个用户, 为便于在客户端查看效果, 建议调试接口时加入开发者自身ID, 示例值: ["ou_7d8a6e6df7621556ce0d21922b676706ccs"], 最大长度: `50`
	BotIDList              []string                            `json:"bot_id_list,omitempty"`              // 创建群时邀请的群机器人；可参考[如何获取应用的 App ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-app-id)来获取应用的App ID, 注意: 操作此接口的机器人会自动入群, 无需重复填写, 拉机器人入群请使用`app_id`, 最多同时邀请5个机器人, 并且群组最多容纳 15 个机器人, 示例值: ["cli_a10fbf7e94b8d01d"], 最大长度: `5`
	ChatMode               *ChatMode                           `json:"chat_mode,omitempty"`                // 群模式, 可选值有: `group`: 群组, 示例值: "group", 默认值: `group`
	ChatType               *ChatType                           `json:"chat_type,omitempty"`                // 群类型, 可选值有: `private`: 私有群, `public`: 公开群, 示例值: "private", 默认值: `private`
	External               *bool                               `json:"external,omitempty"`                 // 是否是外部群；若群组需要邀请不同租户的用户或机器人, 请指定为外部群；, 示例值: false, 默认值: `false`
	JoinMessageVisibility  *MessageVisibility                  `json:"join_message_visibility,omitempty"`  // 入群消息可见性, 可选值有: `only_owner`: 仅群主和管理员可见, `all_members`: 所有成员可见, `not_anyone`: 任何人均不可见, 示例值: "all_members", 默认值: `all_members`
	LeaveMessageVisibility *MessageVisibility                  `json:"leave_message_visibility,omitempty"` // 退群消息可见性, 可选值有: `only_owner`: 仅群主和管理员可见, `all_members`: 所有成员可见, `not_anyone`: 任何人均不可见, 示例值: "all_members", 默认值: `all_members`
	MembershipApproval     *MembershipApproval                 `json:"membership_approval,omitempty"`      // 加群审批, 可选值有: `no_approval_required`: 无需审批, `approval_required`: 需要审批, 示例值: "no_approval_required", 默认值: `no_approval_required`
	RestrictedModeSetting  *CreateChatReqRestrictedModeSetting `json:"restricted_mode_setting,omitempty"`  // 保密模式设置
}

// CreateChatReqRestrictedModeSetting ...
type CreateChatReqRestrictedModeSetting struct {
	Status                         *bool   `json:"status,omitempty"`                            // 保密模式是否开启, 注意: status为true时, screenshot_has_permission_setting、download_has_permission_setting、message_has_permission_setting不能全为all_members, status为false时, screenshot_has_permission_setting、download_has_permission_setting、message_has_permission_setting不能存在not_anyone, 示例值: false
	ScreenshotHasPermissionSetting *string `json:"screenshot_has_permission_setting,omitempty"` // 允许截屏录屏, 示例值: "all_members", 可选值有: all_members: 所有成员允许截屏录屏, not_anyone: 所有成员禁止截屏录屏
	DownloadHasPermissionSetting   *string `json:"download_has_permission_setting,omitempty"`   // 允许下载消息中图片、视频和文件, 示例值: "all_members", 可选值有: all_members: 所有成员允许下载资源, not_anyone: 所有成员禁止下载资源
	MessageHasPermissionSetting    *string `json:"message_has_permission_setting,omitempty"`    // 允许复制和转发消息, 示例值: "all_members", 可选值有: all_members: 所有成员允许复制和转发消息, not_anyone: 所有成员禁止复制和转发消息
}

// CreateChatResp ...
type CreateChatResp struct {
	ChatID                 string                               `json:"chat_id,omitempty"`                  // 群 ID, 详情参见: [群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description)
	Avatar                 string                               `json:"avatar,omitempty"`                   // 群头像 URL
	Name                   string                               `json:"name,omitempty"`                     // 群名称
	Description            string                               `json:"description,omitempty"`              // 群描述
	I18nNames              *I18nNames                           `json:"i18n_names,omitempty"`               // 群国际化名称
	OwnerID                string                               `json:"owner_id,omitempty"`                 // 群主 ID, ID值与查询参数中的 [user_id_type] 对应；不同 ID 的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 注意: 当群主是机器人时, 该字段不返回
	OwnerIDType            IDType                               `json:"owner_id_type,omitempty"`            // 群主 ID 对应的ID类型, 与查询参数中的 [user_id_type] 相同。取值为: `open_id`、`user_id`、`union_id`其中之一, 注意: 当群主是机器人时, 该字段不返回
	AddMemberPermission    AddMemberPermission                  `json:"add_member_permission,omitempty"`    // 拉 用户或机器人 入群权限, 可选值有: `only_owner`: 仅群主和管理员, `all_members`: 所有成员
	ShareCardPermission    ShareCardPermission                  `json:"share_card_permission,omitempty"`    // 群分享权限, 可选值有: `allowed`: 允许, `not_allowed`: 不允许
	AtAllPermission        AtAllPermission                      `json:"at_all_permission,omitempty"`        // at 所有人权限, 可选值有: `only_owner`: 仅群主和管理员, `all_members`: 所有成员
	EditPermission         EditPermission                       `json:"edit_permission,omitempty"`          // 群编辑权限, 可选值有: `only_owner`: 仅群主和管理员, `all_members`: 所有成员
	ChatMode               ChatMode                             `json:"chat_mode,omitempty"`                // 群模式, 可选值有: `group`: 群组
	ChatType               ChatType                             `json:"chat_type,omitempty"`                // 群类型, 可选值有: `private`: 私有群, `public`: 公开群
	ChatTag                string                               `json:"chat_tag,omitempty"`                 // 群标签, 如有多个, 则按照下列顺序返回第一个, 可选值有: `inner`: 内部群, `tenant`: 公司群, `department`: 部门群, `edu`: 教育群, `meeting`: 会议群, `customer_service`: 客服群
	External               bool                                 `json:"external,omitempty"`                 // 是否是外部群
	TenantKey              string                               `json:"tenant_key,omitempty"`               // 租户在飞书上的唯一标识, 用来换取对应的tenant_access_token, 也可以用作租户在应用里面的唯一标识
	JoinMessageVisibility  MessageVisibility                    `json:"join_message_visibility,omitempty"`  // 入群消息可见性, 可选值有: `only_owner`: 仅群主和管理员可见, `all_members`: 所有成员可见, `not_anyone`: 任何人均不可见
	LeaveMessageVisibility MessageVisibility                    `json:"leave_message_visibility,omitempty"` // 出群消息可见性, 可选值有: `only_owner`: 仅群主和管理员可见, `all_members`: 所有成员可见, `not_anyone`: 任何人均不可见
	MembershipApproval     MembershipApproval                   `json:"membership_approval,omitempty"`      // 加群审批, 可选值有: `no_approval_required`: 无需审批, `approval_required`: 需要审批
	ModerationPermission   ModerationPermission                 `json:"moderation_permission,omitempty"`    // 发言权限, 可选值有: `only_owner`: 仅群主和管理员, `all_members`: 所有成员, `moderator_list`: 指定群成员
	RestrictedModeSetting  *CreateChatRespRestrictedModeSetting `json:"restricted_mode_setting,omitempty"`  // 保密模式设置
}

// CreateChatRespRestrictedModeSetting ...
type CreateChatRespRestrictedModeSetting struct {
	Status                         bool   `json:"status,omitempty"`                            // 保密模式是否开启
	ScreenshotHasPermissionSetting string `json:"screenshot_has_permission_setting,omitempty"` // 允许截屏录屏, 可选值有: all_members: 所有成员允许截屏录屏, not_anyone: 所有成员禁止截屏录屏
	DownloadHasPermissionSetting   string `json:"download_has_permission_setting,omitempty"`   // 允许下载消息中图片、视频和文件, 可选值有: all_members: 所有成员允许下载资源, not_anyone: 所有成员禁止下载资源
	MessageHasPermissionSetting    string `json:"message_has_permission_setting,omitempty"`    // 允许复制和转发消息, 可选值有: all_members: 所有成员允许复制和转发消息, not_anyone: 所有成员禁止复制和转发消息
}

// createChatResp ...
type createChatResp struct {
	Code int64           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *CreateChatResp `json:"data,omitempty"`
}
