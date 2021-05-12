// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateMailGroup 创建一个邮件组
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/create
func (r *MailAPI) CreateMailGroup(ctx context.Context, request *CreateMailGroupReq, options ...MethodOptionFunc) (*CreateMailGroupResp, *Response, error) {
	if r.cli.mock.mockMailCreateMailGroup != nil {
		r.cli.logDebug(ctx, "[lark] Mail#CreateMailGroup mock enable")
		return r.cli.mock.mockMailCreateMailGroup(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Mail#CreateMailGroup call api")
	r.cli.logDebug(ctx, "[lark] Mail#CreateMailGroup request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/mailgroups",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createMailGroupResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Mail#CreateMailGroup POST https://open.feishu.cn/open-apis/mail/v1/mailgroups failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Mail#CreateMailGroup POST https://open.feishu.cn/open-apis/mail/v1/mailgroups failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Mail", "CreateMailGroup", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Mail#CreateMailGroup request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockMailCreateMailGroup(f func(ctx context.Context, request *CreateMailGroupReq, options ...MethodOptionFunc) (*CreateMailGroupResp, *Response, error)) {
	r.mockMailCreateMailGroup = f
}

func (r *Mock) UnMockMailCreateMailGroup() {
	r.mockMailCreateMailGroup = nil
}

type CreateMailGroupReq struct {
	Email          *string `json:"email,omitempty"`             // 邮件组地址, 示例值："test_mail_group@xxx.xx"
	Name           *string `json:"name,omitempty"`              // 邮件组名称, 示例值："test mail group"
	Description    *string `json:"description,omitempty"`       // 邮件组描述, 示例值："mail group for testing"
	WhoCanSendMail *string `json:"who_can_send_mail,omitempty"` // 谁可发送邮件到此邮件组, 示例值："ALL_INTERNAL_USERS", 可选值有: `ANYONE`：任何人, `ALL_INTERNAL_USERS`：仅组织内部成员, `ALL_GROUP_MEMBERS`：仅邮件组成员, `CUSTOM_MEMBERS`：自定义成员
}

type createMailGroupResp struct {
	Code int                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *CreateMailGroupResp `json:"data,omitempty"` //
}

type CreateMailGroupResp struct {
	MailGroupID             string `json:"mailgroup_id,omitempty"`               // 邮件组ID
	Email                   string `json:"email,omitempty"`                      // 邮件组地址
	Name                    string `json:"name,omitempty"`                       // 邮件组名称
	Description             string `json:"description,omitempty"`                // 邮件组描述
	DirectMembersCount      string `json:"direct_members_count,omitempty"`       // 邮件组成员数量
	IncludeExternalMember   bool   `json:"include_external_member,omitempty"`    // 是否包含外部成员
	IncludeAllCompanyMember bool   `json:"include_all_company_member,omitempty"` // 是否是全员邮件组
	WhoCanSendMail          string `json:"who_can_send_mail,omitempty"`          // 谁可发送邮件到此邮件组, 可选值有: `ANYONE`：任何人, `ALL_INTERNAL_USERS`：仅组织内部成员, `ALL_GROUP_MEMBERS`：仅邮件组成员, `CUSTOM_MEMBERS`：自定义成员
}
