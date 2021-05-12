// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetStatisticsData
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/query-statistics-data
func (r *AttendanceAPI) GetStatisticsData(ctx context.Context, request *GetStatisticsDataReq, options ...MethodOptionFunc) (*GetStatisticsDataResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetStatisticsData != nil {
		r.cli.logDebug(ctx, "[lark] Attendance#GetStatisticsData mock enable")
		return r.cli.mock.mockAttendanceGetStatisticsData(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Attendance#GetStatisticsData call api")
	r.cli.logDebug(ctx, "[lark] Attendance#GetStatisticsData request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/user_stats_datas/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getStatisticsDataResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Attendance#GetStatisticsData POST https://open.feishu.cn/open-apis/attendance/v1/user_stats_datas/query failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Attendance#GetStatisticsData POST https://open.feishu.cn/open-apis/attendance/v1/user_stats_datas/query failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Attendance", "GetStatisticsData", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Attendance#GetStatisticsData request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockAttendanceGetStatisticsData(f func(ctx context.Context, request *GetStatisticsDataReq, options ...MethodOptionFunc) (*GetStatisticsDataResp, *Response, error)) {
	r.mockAttendanceGetStatisticsData = f
}

func (r *Mock) UnMockAttendanceGetStatisticsData() {
	r.mockAttendanceGetStatisticsData = nil
}

type GetStatisticsDataReq struct {
	EmployeeType     EmployeeType `query:"employee_type" json:"-"`      // 用户 ID 类型, 可选值有: `employee_id`, `employee_no`
	Locale           string       `json:"locale,omitempty"`             // 语言类型, 可选值有: `en`：英文, `ja`：日文, `zh`：中文
	StatsType        string       `json:"stats_type,omitempty"`         // 统计类型,      , 可选值有: `daily`：日度统计, `month`：月度统计
	StartDate        int          `json:"start_date,omitempty"`         // 开始时间, 示例值：20210316
	EndDate          int          `json:"end_date,omitempty"`           // 结束时间, 示例值：20210323,      ,      （时间间隔不超过 40 天）
	UserIDs          []string     `json:"user_ids,omitempty"`           // 查询的用户 ID 列表,      ,      （用户数量不超过 20）
	NeedHistory      *bool        `json:"need_history,omitempty"`       // 是否包含历史数据, 示例值：true
	CurrentGroupOnly *bool        `json:"current_group_only,omitempty"` // 是否只包含当前考勤组, 示例值：true
}

type getStatisticsDataResp struct {
	Code int                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *GetStatisticsDataResp `json:"data,omitempty"` //
}

type GetStatisticsDataResp struct {
	UserDatas []*GetStatisticsDataRespUserData `json:"user_datas,omitempty"` // 用户统计数据
}

type GetStatisticsDataRespUserData struct {
	Name   string                               `json:"name,omitempty"`    // 姓名
	UserID string                               `json:"user_id,omitempty"` // 用户 ID
	Datas  []*GetStatisticsDataRespUserDataData `json:"datas,omitempty"`   // 数据
}

type GetStatisticsDataRespUserDataData struct {
	Code     int                                         `json:"code,omitempty"`     // 字段编号
	Value    string                                      `json:"value,omitempty"`    // 数据值
	Features []*GetStatisticsDataRespUserDataDataFeature `json:"features,omitempty"` // 数据属性
}

type GetStatisticsDataRespUserDataDataFeature struct {
	Key   string `json:"key,omitempty"`   // 属性名
	Value string `json:"value,omitempty"` // 属性值
}
