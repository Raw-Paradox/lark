// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetMeetingRecording 获取一个会议的录制文件。
//
// 会议结束后并且收到了"录制完成"的事件方可获取录制文件；只有会议owner（通过开放平台预约的会议即为预约人）有权限获取；录制时间太短(&lt;5s)有可能无法生成录制文件
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/get
func (r *VCAPI) GetMeetingRecording(ctx context.Context, request *GetMeetingRecordingReq, options ...MethodOptionFunc) (*GetMeetingRecordingResp, *Response, error) {
	if r.cli.mock.mockVCGetMeetingRecording != nil {
		r.cli.logDebug(ctx, "[lark] VC#GetMeetingRecording mock enable")
		return r.cli.mock.mockVCGetMeetingRecording(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] VC#GetMeetingRecording call api")
	r.cli.logDebug(ctx, "[lark] VC#GetMeetingRecording request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:              "GET",
		URL:                 "https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/recording",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(getMeetingRecordingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] VC#GetMeetingRecording GET https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/recording failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] VC#GetMeetingRecording GET https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/recording failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("VC", "GetMeetingRecording", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] VC#GetMeetingRecording request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockVCGetMeetingRecording(f func(ctx context.Context, request *GetMeetingRecordingReq, options ...MethodOptionFunc) (*GetMeetingRecordingResp, *Response, error)) {
	r.mockVCGetMeetingRecording = f
}

func (r *Mock) UnMockVCGetMeetingRecording() {
	r.mockVCGetMeetingRecording = nil
}

type GetMeetingRecordingReq struct {
	MeetingID string `path:"meeting_id" json:"-"` // 会议ID, 示例值："6911188411932033028"
}

type getMeetingRecordingResp struct {
	Code int                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *GetMeetingRecordingResp `json:"data,omitempty"` //
}

type GetMeetingRecordingResp struct {
	Recording *GetMeetingRecordingRespRecording `json:"recording,omitempty"` // 录制文件数据
}

type GetMeetingRecordingRespRecording struct {
	URL      string `json:"url,omitempty"`      // 录指文件URL
	Duration string `json:"duration,omitempty"` // 录制总时长（单位msec）
}
