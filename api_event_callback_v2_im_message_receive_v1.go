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

// EventV2IMMessageReceiveV1 机器人接收到用户发送的消息后触发此事件。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=im&version=v1&resource=message&event=receive)
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability), 并订阅 [消息与群组] 分类下的 [接收消息v2.0] 事件才可接收推送
// - 同时, 将根据应用具备的权限, 判断可推送的信息:
// --当具备[获取用户发给机器人的单聊消息]或者[读取用户发给机器人的单聊消息] 权限, 可接收与机器人单聊会话中用户发送的所有消息
// --当具备[获取群组中所有消息] 权限时, 可接收与机器人所在群聊会话中用户发送的所有消息
// --当具备[获取用户在群组中@机器人的消息]或者[接收群聊中@机器人消息事件] 权限, 可接收机器人所在群聊中用户 @ 机器人的消息
// - 特殊情况下可能会收到重复的推送, 如有幂等需求请使用 [message_id]去重, 不要依赖event_id
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/events/receive
func (r *EventCallbackService) HandlerEventV2IMMessageReceiveV1(f EventV2IMMessageReceiveV1Handler) {
	r.cli.eventHandler.eventV2IMMessageReceiveV1Handler = f
}

// EventV2IMMessageReceiveV1Handler event EventV2IMMessageReceiveV1 handler
type EventV2IMMessageReceiveV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2IMMessageReceiveV1) (string, error)

// EventV2IMMessageReceiveV1 ...
type EventV2IMMessageReceiveV1 struct {
	Sender  *EventV2IMMessageReceiveV1Sender  `json:"sender,omitempty"`  // 事件的发送者
	Message *EventV2IMMessageReceiveV1Message `json:"message,omitempty"` // 事件中包含的消息内容
}

// EventV2IMMessageReceiveV1Message ...
type EventV2IMMessageReceiveV1Message struct {
	MessageID   string                                     `json:"message_id,omitempty"`   // 消息的open_message_id, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	RootID      string                                     `json:"root_id,omitempty"`      // 根消息id, 用于回复消息场景, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	ParentID    string                                     `json:"parent_id,omitempty"`    // 父消息的id, 用于回复消息场景, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	CreateTime  string                                     `json:"create_time,omitempty"`  // 消息发送时间（毫秒）
	ChatID      string                                     `json:"chat_id,omitempty"`      // 消息所在的群组 ID
	ChatType    ChatMode                                   `json:"chat_type,omitempty"`    // 消息所在的群组类型, 可选值有: `p2p`: 单聊, `group`: 群组
	MessageType MsgType                                    `json:"message_type,omitempty"` // 消息类型
	Content     string                                     `json:"content,omitempty"`      // 消息内容, json 格式, [各类型消息Content](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/events/message_content)
	Mentions    []*EventV2IMMessageReceiveV1MessageMention `json:"mentions,omitempty"`     // 被提及用户的信息
}

// EventV2IMMessageReceiveV1MessageMention ...
type EventV2IMMessageReceiveV1MessageMention struct {
	Key       string                                     `json:"key,omitempty"`        // mention key
	ID        *EventV2IMMessageReceiveV1MessageMentionID `json:"id,omitempty"`         // 用户 ID
	Name      string                                     `json:"name,omitempty"`       // 用户姓名
	TenantKey string                                     `json:"tenant_key,omitempty"` // tenant key, 为租户在飞书上的唯一标识, 用来换取对应的tenant_access_token, 也可以用作租户在应用里面的唯一标识
}

// EventV2IMMessageReceiveV1MessageMentionID ...
type EventV2IMMessageReceiveV1MessageMentionID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

// EventV2IMMessageReceiveV1Sender ...
type EventV2IMMessageReceiveV1Sender struct {
	SenderID   *EventV2IMMessageReceiveV1SenderSenderID `json:"sender_id,omitempty"`   // 用户 ID
	SenderType string                                   `json:"sender_type,omitempty"` // 消息发送者类型。目前只支持用户(user)发送的消息。
	TenantKey  string                                   `json:"tenant_key,omitempty"`  // tenant key, 为租户在飞书上的唯一标识, 用来换取对应的tenant_access_token, 也可以用作租户在应用里面的唯一标识
}

// EventV2IMMessageReceiveV1SenderSenderID ...
type EventV2IMMessageReceiveV1SenderSenderID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
