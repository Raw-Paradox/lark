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
	"encoding/json"
)

// UpdateMessageDelay 用户对消息卡片完成交互操作后, 对指定用户更新卡片内容, 给与及时反馈。
//
// ### 使用场景
// 1. 用户点击卡片后业务方需要处理较长时间, 无法在3s内及时返回需要展示的卡片内容。
// 2. 只更新一部分收到这张卡片成员（同一个`message_id`）看到的卡片内容。延迟更新使用交互回传事件中的`token`来指定目标更新的消息, 无需额外关注原消息的`message_id`。
// <b>注意事项</b>:
// - 需要用户主动交互触发, 不支持无条件更新
// - 延迟更新使用的token有效期为30分钟, 超时则无法更新卡片
// - 调用延迟更新接口需要晚于同步返回, 否则会出现不可预测行为
// 服务端处理时, 可先立即 return 空串, 再在30分钟内调用延迟更新接口更新卡片
// - 只能更新用户交互对应卡片, 不允许更新其他卡片
// - 卡片内容经转换后不能超过100KB
// - 同一token仅能使用3次
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMDO1YjLzgTN24yM4UjN
func (r *MessageService) UpdateMessageDelay(ctx context.Context, request *UpdateMessageDelayReq, options ...MethodOptionFunc) (*UpdateMessageDelayResp, *Response, error) {
	if r.cli.mock.mockMessageUpdateMessageDelay != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#UpdateMessageDelay mock enable")
		return r.cli.mock.mockMessageUpdateMessageDelay(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "UpdateMessageDelay",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/interactive/v1/card/update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateMessageDelayResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageUpdateMessageDelay mock MessageUpdateMessageDelay method
func (r *Mock) MockMessageUpdateMessageDelay(f func(ctx context.Context, request *UpdateMessageDelayReq, options ...MethodOptionFunc) (*UpdateMessageDelayResp, *Response, error)) {
	r.mockMessageUpdateMessageDelay = f
}

// UnMockMessageUpdateMessageDelay un-mock MessageUpdateMessageDelay method
func (r *Mock) UnMockMessageUpdateMessageDelay() {
	r.mockMessageUpdateMessageDelay = nil
}

// UpdateMessageDelayReq ...
type UpdateMessageDelayReq struct {
	Token string                     `json:"token,omitempty"` // 用于更新卡片的token, 不是tenant_access_token（可通过[卡片交互返回内容](https://open.feishu.cn/document/ukTMukTMukTM/uEzNwUjLxcDM14SM3ATN)获取）
	Card  *UpdateMessageDelayReqCard `json:"card,omitempty"`  // 消息卡片的描述内容, 具体参考[卡片结构](https://open.feishu.cn/document/ukTMukTMukTM/uEjNwUjLxYDM14SM2ATN)
}

// UpdateMessageDelayReqCard ...
type UpdateMessageDelayReqCard struct {
	Card    interface{} `json:"card,omitempty"`     // 消息卡片的描述内容, 具体参考[卡片结构](https://open.feishu.cn/document/ukTMukTMukTM/uEjNwUjLxYDM14SM2ATN)
	OpenIDs []string    `json:"open_ids,omitempty"` // 指定需要更新的用户, 共享卡片默认更新所有人卡片, 无需填写该字段。推荐使用 OpenID, 获取方式可参考文档[如何获取 Open ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid)
}

func (r UpdateMessageDelayReqCard) MarshalJSON() ([]byte, error) {
	bs, err := json.Marshal(r.Card)
	if err != nil {
		return nil, err
	}
	m := map[string]interface{}{}
	if err = json.Unmarshal(bs, &m); err != nil {
		return nil, err
	}
	m["open_ids"] = r.OpenIDs
	return json.Marshal(m)
}

// UpdateMessageDelayResp ...
type UpdateMessageDelayResp struct {
}

// updateMessageDelayResp ...
type updateMessageDelayResp struct {
	Code int64                   `json:"code,omitempty"` // 返回码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 返回码描述
	Data *UpdateMessageDelayResp `json:"data,omitempty"`
}
