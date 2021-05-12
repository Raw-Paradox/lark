// Code generated by lark_sdk_gen. DO NOT EDIT.

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Approval_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Approval()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetInstanceList(ctx, &lark.GetInstanceListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		moduleCli := cli.Approval()

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApprovalGetInstanceList(func(ctx context.Context, request *lark.GetInstanceListReq, options ...lark.MethodOptionFunc) (*lark.GetInstanceListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalGetInstanceList()

			_, _, err := moduleCli.GetInstanceList(ctx, &lark.GetInstanceListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Approval()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetInstanceList(ctx, &lark.GetInstanceListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})
}
