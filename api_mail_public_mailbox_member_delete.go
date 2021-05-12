// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeletePublicMailboxMember 删除公共邮箱单个成员
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/delete
func (r *MailAPI) DeletePublicMailboxMember(ctx context.Context, request *DeletePublicMailboxMemberReq, options ...MethodOptionFunc) (*DeletePublicMailboxMemberResp, *Response, error) {
	if r.cli.mock.mockMailDeletePublicMailboxMember != nil {
		r.cli.logDebug(ctx, "[lark] Mail#DeletePublicMailboxMember mock enable")
		return r.cli.mock.mockMailDeletePublicMailboxMember(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Mail#DeletePublicMailboxMember call api")
	r.cli.logDebug(ctx, "[lark] Mail#DeletePublicMailboxMember request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/:member_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deletePublicMailboxMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Mail#DeletePublicMailboxMember DELETE https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/:member_id failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Mail#DeletePublicMailboxMember DELETE https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/:member_id failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Mail", "DeletePublicMailboxMember", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Mail#DeletePublicMailboxMember request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockMailDeletePublicMailboxMember(f func(ctx context.Context, request *DeletePublicMailboxMemberReq, options ...MethodOptionFunc) (*DeletePublicMailboxMemberResp, *Response, error)) {
	r.mockMailDeletePublicMailboxMember = f
}

func (r *Mock) UnMockMailDeletePublicMailboxMember() {
	r.mockMailDeletePublicMailboxMember = nil
}

type DeletePublicMailboxMemberReq struct {
	PublicMailboxID string `path:"public_mailbox_id" json:"-"` // 公共邮箱唯一标识或公共邮箱地址, 示例值："xxxxxxxxxxxxxxx 或 test_public_mailbox@xxx.xx"
	MemberID        string `path:"member_id" json:"-"`         // 公共邮箱内成员唯一标识, 示例值："xxxxxxxxxxxxxxx"
}

type deletePublicMailboxMemberResp struct {
	Code int                            `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                         `json:"msg,omitempty"`  // 错误描述
	Data *DeletePublicMailboxMemberResp `json:"data,omitempty"`
}

type DeletePublicMailboxMemberResp struct{}
