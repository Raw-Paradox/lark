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

// UpdateSearchSchema 修改数据范式。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/schema/patch
func (r *SearchService) UpdateSearchSchema(ctx context.Context, request *UpdateSearchSchemaReq, options ...MethodOptionFunc) (*UpdateSearchSchemaResp, *Response, error) {
	if r.cli.mock.mockSearchUpdateSearchSchema != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Search#UpdateSearchSchema mock enable")
		return r.cli.mock.mockSearchUpdateSearchSchema(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Search",
		API:                   "UpdateSearchSchema",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/search/v2/schemas/:schema_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateSearchSchemaResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockSearchUpdateSearchSchema mock SearchUpdateSearchSchema method
func (r *Mock) MockSearchUpdateSearchSchema(f func(ctx context.Context, request *UpdateSearchSchemaReq, options ...MethodOptionFunc) (*UpdateSearchSchemaResp, *Response, error)) {
	r.mockSearchUpdateSearchSchema = f
}

// UnMockSearchUpdateSearchSchema un-mock SearchUpdateSearchSchema method
func (r *Mock) UnMockSearchUpdateSearchSchema() {
	r.mockSearchUpdateSearchSchema = nil
}

// UpdateSearchSchemaReq ...
type UpdateSearchSchemaReq struct {
	SchemaID string                        `path:"schema_id" json:"-"` // 用户自定义数据范式的唯一标识, 示例值: "custom_schema_id", 最大长度: `40` 字符, 正则校验: `^[a-zA-Z][a-zA-Z0-9-_].*$`
	Display  *UpdateSearchSchemaReqDisplay `json:"display,omitempty"`  // 数据展示相关配置
}

// UpdateSearchSchemaReqDisplay ...
type UpdateSearchSchemaReqDisplay struct {
	CardKey       string                                       `json:"card_key,omitempty"`       // 搜索数据的展示卡片, 卡片详细信息请参考 [通用模块接入指南](/document/uAjLw4CM/ukTMukTMukTM/search-v2/common-template-intergration-handbook)  "请求创建数据范式"部分, 示例值: "search_common_card", 可选值有: search_common_card: 普通 common 卡片
	FieldsMapping []*UpdateSearchSchemaReqDisplayFieldsMapping `json:"fields_mapping,omitempty"` // 数据字段名称和展示字段名称的映射关系。如果没有设置, 则只会展示 与展示字段名称同名的 数据字段
}

// UpdateSearchSchemaReqDisplayFieldsMapping ...
type UpdateSearchSchemaReqDisplayFieldsMapping struct {
	DisplayField string `json:"display_field,omitempty"` // 展示字段名称, 与 card_key 有关, 每个模版能展示的字段不同。该字段不能重复, 示例值: "summary"
	DataField    string `json:"data_field,omitempty"`    // 数据字段的名称。需要确保该字段对应在 schema 属性定义中的 is_returnable 为 true, 否则无法展示。需要使用 ${xxx} 的规则来描述, 示例值: "${description}"
}

// UpdateSearchSchemaResp ...
type UpdateSearchSchemaResp struct {
	Schema *UpdateSearchSchemaRespSchema `json:"schema,omitempty"` // 数据范式实例
}

// UpdateSearchSchemaRespSchema ...
type UpdateSearchSchemaRespSchema struct {
	Properties []*UpdateSearchSchemaRespSchemaPropertie `json:"properties,omitempty"` // 数据范式的属性定义
	Display    *UpdateSearchSchemaRespSchemaDisplay     `json:"display,omitempty"`    // 数据展示相关配置
	SchemaID   string                                   `json:"schema_id,omitempty"`  // 用户自定义数据范式的唯一标识
}

// UpdateSearchSchemaRespSchemaDisplay ...
type UpdateSearchSchemaRespSchemaDisplay struct {
	CardKey       string                                              `json:"card_key,omitempty"`       // 搜索数据的展示卡片, 卡片详细信息请参考 [通用模块接入指南](/document/uAjLw4CM/ukTMukTMukTM/search-v2/common-template-intergration-handbook)  "请求创建数据范式"部分, 可选值有: search_common_card: 普通 common 卡片
	FieldsMapping []*UpdateSearchSchemaRespSchemaDisplayFieldsMapping `json:"fields_mapping,omitempty"` // 数据字段名称和展示字段名称的映射关系。如果没有设置, 则只会展示 与展示字段名称同名的 数据字段
}

// UpdateSearchSchemaRespSchemaDisplayFieldsMapping ...
type UpdateSearchSchemaRespSchemaDisplayFieldsMapping struct {
	DisplayField string `json:"display_field,omitempty"` // 展示字段名称, 与 card_key 有关, 每个模版能展示的字段不同。该字段不能重复
	DataField    string `json:"data_field,omitempty"`    // 数据字段的名称。需要确保该字段对应在 schema 属性定义中的 is_returnable 为 true, 否则无法展示。需要使用 ${xxx} 的规则来描述
}

// UpdateSearchSchemaRespSchemaPropertie ...
type UpdateSearchSchemaRespSchemaPropertie struct {
	Name            string                                                `json:"name,omitempty"`             // 属性名
	Type            string                                                `json:"type,omitempty"`             // 属性类型, 可选值有: text: 长文本类型, int: 64位整数类型, tag: 标签类型, timestamp: Unix 时间戳类型（单位为秒）, double: 浮点数类型（小数）, tinytext: 短文本类型, （utf8 编码）长度小于 140 的文本。在设置 search_options 时, 与 text 类型有区别, 支持更多召回策略
	IsSearchable    bool                                                  `json:"is_searchable,omitempty"`    // 该属性是否可用作搜索, 默认为 false
	IsSortable      bool                                                  `json:"is_sortable,omitempty"`      // 该属性是否可用作搜索结果排序, 默认为 false。如果为 true, 需要再配置 sortOptions
	IsReturnable    bool                                                  `json:"is_returnable,omitempty"`    // 该属性是否可用作返回字段, 为 false 时, 该字段不会被召回和展示。默认为 false
	SortOptions     *UpdateSearchSchemaRespSchemaPropertieSortOptions     `json:"sort_options,omitempty"`     // 属性排序的可选配置, 当 is_sortable 为 true 时, 该字段为必填字段
	TypeDefinitions *UpdateSearchSchemaRespSchemaPropertieTypeDefinitions `json:"type_definitions,omitempty"` // 相关类型数据的定义和约束
	SearchOptions   *UpdateSearchSchemaRespSchemaPropertieSearchOptions   `json:"search_options,omitempty"`   // 属性搜索的可选配置, 当 is_searchable 为 true 时, 该字段为必填参数
}

// UpdateSearchSchemaRespSchemaPropertieSearchOptions ...
type UpdateSearchSchemaRespSchemaPropertieSearchOptions struct {
	EnableSemanticMatch     bool `json:"enable_semantic_match,omitempty"`      // 是否支持语义切词召回。默认不支持（推荐使用在长文本的场景）
	EnableExactMatch        bool `json:"enable_exact_match,omitempty"`         // 是否支持精确匹配。默认不支持（推荐使用在短文本、需要精确查找的场景）
	EnablePrefixMatch       bool `json:"enable_prefix_match,omitempty"`        // 是否支持前缀匹配（短文本的默认的分词/召回策略。前缀长度为 1-12）
	EnableNumberSuffixMatch bool `json:"enable_number_suffix_match,omitempty"` // 是否支持数据后缀匹配。默认不支持（推荐使用在短文本、有数字后缀查找的场景。后缀长度为3-12）
	EnableCamelMatch        bool `json:"enable_camel_match,omitempty"`         // 是否支持驼峰英文匹配。默认不支持（推荐使用在短文本, 且包含驼峰形式英文的查找场景）
}

// UpdateSearchSchemaRespSchemaPropertieSortOptions ...
type UpdateSearchSchemaRespSchemaPropertieSortOptions struct {
	Priority int64  `json:"priority,omitempty"` // 排序的优先级, 可选范围为 0~4, 0为最高优先级。如果优先级相同, 则随机进行排序。默认为0, 可选值有: 0: 最高优先级, 1: 次高优先级, 2: 次次高优先级, 3: 次低优先级, 4: 最低优先级
	Order    string `json:"order,omitempty"`    // 排序的顺序。默认为 desc, 可选值有: asc: 升序, desc: 降序
}

// UpdateSearchSchemaRespSchemaPropertieTypeDefinitions ...
type UpdateSearchSchemaRespSchemaPropertieTypeDefinitions struct {
	Tag []*UpdateSearchSchemaRespSchemaPropertieTypeDefinitionsTag `json:"tag,omitempty"` // 标签类型的定义
}

// UpdateSearchSchemaRespSchemaPropertieTypeDefinitionsTag ...
type UpdateSearchSchemaRespSchemaPropertieTypeDefinitionsTag struct {
	Name  string `json:"name,omitempty"`  // tag 对应的枚举值名称
	Color string `json:"color,omitempty"` // 标签对应的颜色, 可选值有: red: 含警示性、敏感性的提示信息, green: 表示成功、完成、完毕的提示信息, blue: 组件架构、职能等中性信息, grey: 中立系统提示信息（慎重使用）, yellow: 焦点信息、推广性信息
	Text  string `json:"text,omitempty"`  // 标签中展示的文本
}

// updateSearchSchemaResp ...
type updateSearchSchemaResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *UpdateSearchSchemaResp `json:"data,omitempty"`
}
