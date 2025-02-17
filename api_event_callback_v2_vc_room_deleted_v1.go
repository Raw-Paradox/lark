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

// EventV2VCRoomDeletedV1 当删除会议室时, 会触发该事件。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=vc&version=v1&resource=room&event=deleted)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room/events/deleted
func (r *EventCallbackService) HandlerEventV2VCRoomDeletedV1(f EventV2VCRoomDeletedV1Handler) {
	r.cli.eventHandler.eventV2VCRoomDeletedV1Handler = f
}

// EventV2VCRoomDeletedV1Handler event EventV2VCRoomDeletedV1 handler
type EventV2VCRoomDeletedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2VCRoomDeletedV1) (string, error)

// EventV2VCRoomDeletedV1 ...
type EventV2VCRoomDeletedV1 struct {
	Room *EventV2VCRoomDeletedV1Room `json:"room,omitempty"` // 会议室信息
}

// EventV2VCRoomDeletedV1Room ...
type EventV2VCRoomDeletedV1Room struct {
	RoomID string `json:"room_id,omitempty"` // 会议室ID
}
