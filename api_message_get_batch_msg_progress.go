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

// GetBatchSentMessageProgress 该接口在[查询批量消息推送和阅读人数](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/batch_message/read_user)查询结果的基础上, 增加了批量请求中有效的userid数量以及消息撤回进度数据。
//
// 注意事项:
// - 只能查询30天内通过[批量发送消息](https://open.feishu.cn/document/ukTMukTMukTM/ucDO1EjL3gTNx4yN4UTM)接口产生的消息
// - 该接口返回的数据为查询时刻的快照数据
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/batch_message/get_progress
func (r *MessageService) GetBatchSentMessageProgress(ctx context.Context, request *GetBatchSentMessageProgressReq, options ...MethodOptionFunc) (*GetBatchSentMessageProgressResp, *Response, error) {
	if r.cli.mock.mockMessageGetBatchSentMessageProgress != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#GetBatchSentMessageProgress mock enable")
		return r.cli.mock.mockMessageGetBatchSentMessageProgress(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "GetBatchSentMessageProgress",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/batch_messages/:batch_message_id/get_progress",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getBatchSentMessageProgressResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageGetBatchSentMessageProgress mock MessageGetBatchSentMessageProgress method
func (r *Mock) MockMessageGetBatchSentMessageProgress(f func(ctx context.Context, request *GetBatchSentMessageProgressReq, options ...MethodOptionFunc) (*GetBatchSentMessageProgressResp, *Response, error)) {
	r.mockMessageGetBatchSentMessageProgress = f
}

// UnMockMessageGetBatchSentMessageProgress un-mock MessageGetBatchSentMessageProgress method
func (r *Mock) UnMockMessageGetBatchSentMessageProgress() {
	r.mockMessageGetBatchSentMessageProgress = nil
}

// GetBatchSentMessageProgressReq ...
type GetBatchSentMessageProgressReq struct {
	BatchMessageID string `path:"batch_message_id" json:"-"` // 待查询的批量消息的ID, 通过调用[批量发送消息接口](	/ssl:ttdoc/ukTMukTMukTM/ucDO1EjL3gTNx4yN4UTM)的返回值中得到, 示例值: "bm-0b3d5d1b2df7c6d5dbd1abe2c91e2217"
}

// GetBatchSentMessageProgressResp ...
type GetBatchSentMessageProgressResp struct {
	BatchMessageSendProgress   []*GetBatchSentMessageProgressRespBatchMessageSendProgres   `json:"batch_message_send_progress,omitempty"`   // 消息发送进度
	BatchMessageRecallProgress []*GetBatchSentMessageProgressRespBatchMessageRecallProgres `json:"batch_message_recall_progress,omitempty"` // 消息撤回进度
}

// GetBatchSentMessageProgressRespBatchMessageRecallProgres ...
type GetBatchSentMessageProgressRespBatchMessageRecallProgres struct {
	Recall      bool  `json:"recall,omitempty"`       // 该条批量消息是否被执行过撤回操作
	RecallCount int64 `json:"recall_count,omitempty"` // 已经成功撤回的消息数量
}

// GetBatchSentMessageProgressRespBatchMessageSendProgres ...
type GetBatchSentMessageProgressRespBatchMessageSendProgres struct {
	ValidUserIDsCount   int64 `json:"valid_user_ids_count,omitempty"`   // 批量请求中有效的userid数量(包含机器人不可见用户), 注意: 当valid_user_ids_count为0有两种情况: * 批量任务还没有开始被调度（请等待一会再调用该接口）, * 批量发送消息时传入的所有openIDs、employeID、departmentiIDs都不包含有效的用户
	SuccessUserIDsCount int64 `json:"success_user_ids_count,omitempty"` // 已经成功给用户发送成功的消息数量, 注意: 最终success_user_ids_count不一定等于valid_user_ids_count, 因为valid_user_ids_count包含了对机器人不可见的用户
	ReadUserIDsCount    int64 `json:"read_user_ids_count,omitempty"`    // 已读信息用户数量
}

// getBatchSentMessageProgressResp ...
type getBatchSentMessageProgressResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *GetBatchSentMessageProgressResp `json:"data,omitempty"`
}
