// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// EventV2ContactEmployeeTypeEnumActivedV3
//
// 启用人员类型会发出对应事件，需要权限为`以应用身份访问通讯录`,推送方式为`Webhook`。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=contact&version=v3&resource=employee_type_enum&event=actived)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/events/actived
func (r *EventCallbackService) HandlerEventV2ContactEmployeeTypeEnumActivedV3(f eventV2ContactEmployeeTypeEnumActivedV3Handler) {
	r.cli.eventHandler.eventV2ContactEmployeeTypeEnumActivedV3Handler = f
}

type eventV2ContactEmployeeTypeEnumActivedV3Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2ContactEmployeeTypeEnumActivedV3) (string, error)

type EventV2ContactEmployeeTypeEnumActivedV3 struct {
	OldEnum *EventV2ContactEmployeeTypeEnumActivedV3OldEnum `json:"old_enum,omitempty"` // 旧枚举类型
	NewEnum *EventV2ContactEmployeeTypeEnumActivedV3NewEnum `json:"new_enum,omitempty"` // 新枚举类型
}

type EventV2ContactEmployeeTypeEnumActivedV3OldEnum struct {
	EnumID      string                                                     `json:"enum_id,omitempty"`      // 枚举值id
	EnumValue   string                                                     `json:"enum_value,omitempty"`   // 枚举值
	Content     string                                                     `json:"content,omitempty"`      // 枚举内容, 长度范围：`1` ～ `100` 字符
	EnumType    int64                                                      `json:"enum_type,omitempty"`    // 类型, 可选值有: `1`：内置类型, `2`：自定义
	EnumStatus  int64                                                      `json:"enum_status,omitempty"`  // 类型, 可选值有: `1`：激活, `2`：未激活
	I18nContent *EventV2ContactEmployeeTypeEnumActivedV3OldEnumI18nContent `json:"i18n_content,omitempty"` // i18n定义
}

type EventV2ContactEmployeeTypeEnumActivedV3OldEnumI18nContent struct {
	Locale string `json:"locale,omitempty"` // 语言
	Value  string `json:"value,omitempty"`  // i18n内容
}

type EventV2ContactEmployeeTypeEnumActivedV3NewEnum struct {
	EnumID      string                                                     `json:"enum_id,omitempty"`      // 枚举值id
	EnumValue   string                                                     `json:"enum_value,omitempty"`   // 枚举值
	Content     string                                                     `json:"content,omitempty"`      // 枚举内容, 长度范围：`1` ～ `100` 字符
	EnumType    int64                                                      `json:"enum_type,omitempty"`    // 类型, 可选值有: `1`：内置类型, `2`：自定义
	EnumStatus  int64                                                      `json:"enum_status,omitempty"`  // 类型, 可选值有: `1`：激活, `2`：未激活
	I18nContent *EventV2ContactEmployeeTypeEnumActivedV3NewEnumI18nContent `json:"i18n_content,omitempty"` // i18n定义
}

type EventV2ContactEmployeeTypeEnumActivedV3NewEnumI18nContent struct {
	Locale string `json:"locale,omitempty"` // 语言
	Value  string `json:"value,omitempty"`  // i18n内容
}
