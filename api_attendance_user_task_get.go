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

// GetAttendanceUserTask 获取企业内员工的实际打卡结果, 包括上班打卡结果和下班打卡结果。
//
// - 如果企业给一个员工设定的班次是上午 9 点和下午 6 点各打一次上下班卡, 即使员工在这期间打了多次卡, 该接口也只会返回 1 条记录。
// - 如果要获取打卡的详细数据, 如打卡位置等信息, 可使用“获取打卡流水记录”或“批量查询打卡流水记录”的接口。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task/query
func (r *AttendanceService) GetAttendanceUserTask(ctx context.Context, request *GetAttendanceUserTaskReq, options ...MethodOptionFunc) (*GetAttendanceUserTaskResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetAttendanceUserTask != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#GetAttendanceUserTask mock enable")
		return r.cli.mock.mockAttendanceGetAttendanceUserTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "GetAttendanceUserTask",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/user_tasks/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAttendanceUserTaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceGetAttendanceUserTask mock AttendanceGetAttendanceUserTask method
func (r *Mock) MockAttendanceGetAttendanceUserTask(f func(ctx context.Context, request *GetAttendanceUserTaskReq, options ...MethodOptionFunc) (*GetAttendanceUserTaskResp, *Response, error)) {
	r.mockAttendanceGetAttendanceUserTask = f
}

// UnMockAttendanceGetAttendanceUserTask un-mock AttendanceGetAttendanceUserTask method
func (r *Mock) UnMockAttendanceGetAttendanceUserTask() {
	r.mockAttendanceGetAttendanceUserTask = nil
}

// GetAttendanceUserTaskReq ...
type GetAttendanceUserTaskReq struct {
	EmployeeType          EmployeeType `query:"employee_type" json:"-"`           // 员工工号类型, 示例值: "employee_id", 可选值有: employee_id: 员工 employee ID, 即[飞书管理后台](https://bytedance.feishu.cn/admin/contacts/departmentanduser) > 组织架构 > 成员与部门 > 成员详情中的用户 ID, employee_no: 员工工号, 即[飞书管理后台](https://bytedance.feishu.cn/admin/contacts/departmentanduser) > 组织架构 > 成员与部门 > 成员详情中的工号
	IgnoreInvalidUsers    *bool        `query:"ignore_invalid_users" json:"-"`    // 是否忽略无效和没有权限的用户。如果 true, 则返回有效用户的信息, 并告知无效和没有权限的用户信息；如果 false, 且 user_ids 中存在无效或没有权限的用户, 则返回错误, 示例值: true
	IncludeTerminatedUser *bool        `query:"include_terminated_user" json:"-"` // 由于新入职员工可以复用已离职员工的 employee_no/employee_id, 如果 true, 则返回 employee_no/employee_id 对应的所有在职 + 离职员工的数据；如果 false, 则只返回 employee_no/employee_id 对应的在职或最近一个离职员工的数据, 示例值: true
	UserIDs               []string     `json:"user_ids,omitempty"`                // employee_no 或 employee_id 列表, 长度不超过 50, 示例值: abd754f7
	CheckDateFrom         int64        `json:"check_date_from,omitempty"`         // 查询的起始工作日, 示例值: 20190817
	CheckDateTo           int64        `json:"check_date_to,omitempty"`           // 查询的结束工作日, 示例值: 20190820
}

// GetAttendanceUserTaskResp ...
type GetAttendanceUserTaskResp struct {
	UserTaskResults     []*GetAttendanceUserTaskRespUserTaskResult `json:"user_task_results,omitempty"`     // 打卡任务列表
	InvalidUserIDs      []string                                   `json:"invalid_user_ids,omitempty"`      // 无效用户 ID 列表
	UnauthorizedUserIDs []string                                   `json:"unauthorized_user_ids,omitempty"` // 没有权限用户 ID 列表
}

// GetAttendanceUserTaskRespUserTaskResult ...
type GetAttendanceUserTaskRespUserTaskResult struct {
	ResultID     string                                           `json:"result_id,omitempty"`     // 打卡记录 ID
	UserID       string                                           `json:"user_id,omitempty"`       // 用户 ID
	EmployeeName string                                           `json:"employee_name,omitempty"` // 用户姓名
	Day          int64                                            `json:"day,omitempty"`           // 日期
	GroupID      string                                           `json:"group_id,omitempty"`      // 考勤组 ID
	ShiftID      string                                           `json:"shift_id,omitempty"`      // 班次 ID
	Records      []*GetAttendanceUserTaskRespUserTaskResultRecord `json:"records,omitempty"`       // 用户考勤记录
}

// GetAttendanceUserTaskRespUserTaskResultRecord ...
type GetAttendanceUserTaskRespUserTaskResultRecord struct {
	CheckInRecordID          string                                                       `json:"check_in_record_id,omitempty"`          // 上班打卡记录 ID
	CheckInRecord            *GetAttendanceUserTaskRespUserTaskResultRecordCheckInRecord  `json:"check_in_record,omitempty"`             // 上班打卡记录
	CheckOutRecordID         string                                                       `json:"check_out_record_id,omitempty"`         // 下班打卡记录 ID
	CheckOutRecord           *GetAttendanceUserTaskRespUserTaskResultRecordCheckOutRecord `json:"check_out_record,omitempty"`            // 下班打卡记录
	CheckInResult            string                                                       `json:"check_in_result,omitempty"`             // 上班打卡结果, 可选值有: NoNeedCheck: 无需打卡, SystemCheck: 系统打卡, Normal: 正常, Early: 早退, Late: 迟到, Lack: 缺卡
	CheckOutResult           string                                                       `json:"check_out_result,omitempty"`            // 下班打卡结果, 可选值有: NoNeedCheck: 无需打卡, SystemCheck: 系统打卡, Normal: 正常, Early: 早退, Late: 迟到, Lack: 缺卡
	CheckInResultSupplement  string                                                       `json:"check_in_result_supplement,omitempty"`  // 上班打卡结果补充, 可选值有: None: 无, ManagerModification: 管理员修改, CardReplacement: 补卡通过, ShiftChange: 换班, Travel: 出差, Leave: 请假, GoOut: 外出, CardReplacementApplication: 补卡申请中, FieldPunch: 外勤打卡
	CheckOutResultSupplement string                                                       `json:"check_out_result_supplement,omitempty"` // 下班打卡结果补充, 可选值有: None: 无, ManagerModification: 管理员修改, CardReplacement: 补卡通过, ShiftChange: 换班, Travel: 出差, Leave: 请假, GoOut: 外出, CardReplacementApplication: 补卡申请中, FieldPunch: 外勤打卡
	CheckInShiftTime         string                                                       `json:"check_in_shift_time,omitempty"`         // 上班打卡时间
	CheckOutShiftTime        string                                                       `json:"check_out_shift_time,omitempty"`        // 下班打卡时间
}

// GetAttendanceUserTaskRespUserTaskResultRecordCheckInRecord ...
type GetAttendanceUserTaskRespUserTaskResultRecordCheckInRecord struct {
	UserID       string   `json:"user_id,omitempty"`       // 用户 ID
	CreatorID    string   `json:"creator_id,omitempty"`    // 记录创建者 ID
	LocationName string   `json:"location_name,omitempty"` // 打卡位置名称信息
	CheckTime    string   `json:"check_time,omitempty"`    // 打卡时间, 精确到秒的时间戳
	Comment      string   `json:"comment,omitempty"`       // 打卡备注
	RecordID     string   `json:"record_id,omitempty"`     // 打卡记录 ID
	Ssid         string   `json:"ssid,omitempty"`          // 打卡 Wi-Fi 的 SSID
	Bssid        string   `json:"bssid,omitempty"`         // 打卡 Wi-Fi 的 MAC 地址
	IsField      bool     `json:"is_field,omitempty"`      // 是否为外勤打卡
	IsWifi       bool     `json:"is_wifi,omitempty"`       // 是否为 Wi-Fi 打卡
	Type         int64    `json:"type,omitempty"`          // 记录生成方式, 可选值有: 0: 用户打卡, 1: 管理员修改, 2: 用户补卡, 3: 系统自动生成, 4: 下班免打卡, 5: 考勤机, 6: 极速打卡, 7: 考勤开放平台导入
	PhotoURLs    []string `json:"photo_urls,omitempty"`    // 打卡照片列表
	CheckResult  string   `json:"check_result,omitempty"`  // 打卡结果, 可选值有: NoNeedCheck: 无需打卡, SystemCheck: 系统打卡, Normal: 正常, Early: 早退, Late: 迟到, SeriousLate: 严重迟到, Lack: 缺卡, Invalid: 无效, None: 无状态, Todo: 尚未打卡
}

// GetAttendanceUserTaskRespUserTaskResultRecordCheckOutRecord ...
type GetAttendanceUserTaskRespUserTaskResultRecordCheckOutRecord struct {
	UserID       string   `json:"user_id,omitempty"`       // 用户 ID
	CreatorID    string   `json:"creator_id,omitempty"`    // 记录创建者 ID
	LocationName string   `json:"location_name,omitempty"` // 打卡位置名称信息
	CheckTime    string   `json:"check_time,omitempty"`    // 打卡时间, 精确到秒的时间戳
	Comment      string   `json:"comment,omitempty"`       // 打卡备注
	RecordID     string   `json:"record_id,omitempty"`     // 打卡记录 ID
	Ssid         string   `json:"ssid,omitempty"`          // 打卡 Wi-Fi 的 SSID
	Bssid        string   `json:"bssid,omitempty"`         // 打卡 Wi-Fi 的 MAC 地址
	IsField      bool     `json:"is_field,omitempty"`      // 是否为外勤打卡
	IsWifi       bool     `json:"is_wifi,omitempty"`       // 是否为 Wi-Fi 打卡
	Type         int64    `json:"type,omitempty"`          // 记录生成方式, 可选值有: 0: 用户打卡, 1: 管理员修改, 2: 用户补卡, 3: 系统自动生成, 4: 下班免打卡, 5: 考勤机, 6: 极速打卡, 7: 考勤开放平台导入
	PhotoURLs    []string `json:"photo_urls,omitempty"`    // 打卡照片列表
	CheckResult  string   `json:"check_result,omitempty"`  // 打卡结果, 可选值有: NoNeedCheck: 无需打卡, SystemCheck: 系统打卡, Normal: 正常, Early: 早退, Late: 迟到, SeriousLate: 严重迟到, Lack: 缺卡, Invalid: 无效, None: 无状态, Todo: 尚未打卡
}

// getAttendanceUserTaskResp ...
type getAttendanceUserTaskResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *GetAttendanceUserTaskResp `json:"data,omitempty"`
}
