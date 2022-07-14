// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.1
// source: GameRoom.proto

package netproto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GameRoomClassID int32

const (
	GameRoomClassID_LoginRoomID              GameRoomClassID = 1
	GameRoomClassID_LoginRoomRetID           GameRoomClassID = 2
	GameRoomClassID_UserQueueID              GameRoomClassID = 3
	GameRoomClassID_GameReadyID              GameRoomClassID = 4
	GameRoomClassID_RequestGameVerID         GameRoomClassID = 5
	GameRoomClassID_GameVerInfoID            GameRoomClassID = 6
	GameRoomClassID_UserQueueAfterQuitID     GameRoomClassID = 11
	GameRoomClassID_UserForceLeaveRoomID     GameRoomClassID = 12
	GameRoomClassID_SlotGetGameProgressID    GameRoomClassID = 13
	GameRoomClassID_SlotGetGameProgressRetID GameRoomClassID = 14
	GameRoomClassID_SlotSaveGameProgressID   GameRoomClassID = 15
	GameRoomClassID_SlotGetJackpotID         GameRoomClassID = 16
	GameRoomClassID_SlotGetJackpotRetID      GameRoomClassID = 17
	GameRoomClassID_SlotUpdateJackpotID      GameRoomClassID = 18
	GameRoomClassID_UserSitTableID           GameRoomClassID = 19
	GameRoomClassID_ShowTableListID          GameRoomClassID = 20
	GameRoomClassID_UserLeaveTableID         GameRoomClassID = 21
	GameRoomClassID_UserLeaveTableRetID      GameRoomClassID = 22
	GameRoomClassID_SlotGetJackpotGroupID    GameRoomClassID = 23
	GameRoomClassID_SlotGetJackpotGroupRetID GameRoomClassID = 24
	GameRoomClassID_SlotUpdateJackpotGroupID GameRoomClassID = 25
	GameRoomClassID_UserSendVipBrowReqID     GameRoomClassID = 26
	GameRoomClassID_UserSendVipBrowRetID     GameRoomClassID = 27
	GameRoomClassID_BroadcastUserVipBrowID   GameRoomClassID = 28
	GameRoomClassID_AddUserTaskGameEvent     GameRoomClassID = 29
	GameRoomClassID_GetUserSingleControl     GameRoomClassID = 30
	GameRoomClassID_GetSinglePool            GameRoomClassID = 31
	GameRoomClassID_UpdateSinglePool         GameRoomClassID = 32
	GameRoomClassID_GetUserReelsControl      GameRoomClassID = 33
	GameRoomClassID_GetReelsControlCYZS      GameRoomClassID = 34
	GameRoomClassID_GetReelsControlCSFFF     GameRoomClassID = 35
	GameRoomClassID_GameRoomEmoji            GameRoomClassID = 36
	GameRoomClassID_GameRoomEmojiRet         GameRoomClassID = 37
	GameRoomClassID_DzpkChoiceTable          GameRoomClassID = 38 // 選桌
	GameRoomClassID_DzpkChoiceTableRet       GameRoomClassID = 39
	GameRoomClassID_DzpkChoiceSeat           GameRoomClassID = 40 // 選座位
	GameRoomClassID_DzpkChoiceSeatRet        GameRoomClassID = 41
	GameRoomClassID_DzpkSitSeatTimeout       GameRoomClassID = 42 // 座位被取消
	GameRoomClassID_DzpkBuyChip              GameRoomClassID = 43 // 買籌碼
)

// Enum value maps for GameRoomClassID.
var (
	GameRoomClassID_name = map[int32]string{
		1:  "LoginRoomID",
		2:  "LoginRoomRetID",
		3:  "UserQueueID",
		4:  "GameReadyID",
		5:  "RequestGameVerID",
		6:  "GameVerInfoID",
		11: "UserQueueAfterQuitID",
		12: "UserForceLeaveRoomID",
		13: "SlotGetGameProgressID",
		14: "SlotGetGameProgressRetID",
		15: "SlotSaveGameProgressID",
		16: "SlotGetJackpotID",
		17: "SlotGetJackpotRetID",
		18: "SlotUpdateJackpotID",
		19: "UserSitTableID",
		20: "ShowTableListID",
		21: "UserLeaveTableID",
		22: "UserLeaveTableRetID",
		23: "SlotGetJackpotGroupID",
		24: "SlotGetJackpotGroupRetID",
		25: "SlotUpdateJackpotGroupID",
		26: "UserSendVipBrowReqID",
		27: "UserSendVipBrowRetID",
		28: "BroadcastUserVipBrowID",
		29: "AddUserTaskGameEvent",
		30: "GetUserSingleControl",
		31: "GetSinglePool",
		32: "UpdateSinglePool",
		33: "GetUserReelsControl",
		34: "GetReelsControlCYZS",
		35: "GetReelsControlCSFFF",
		36: "GameRoomEmoji",
		37: "GameRoomEmojiRet",
		38: "DzpkChoiceTable",
		39: "DzpkChoiceTableRet",
		40: "DzpkChoiceSeat",
		41: "DzpkChoiceSeatRet",
		42: "DzpkSitSeatTimeout",
		43: "DzpkBuyChip",
	}
	GameRoomClassID_value = map[string]int32{
		"LoginRoomID":              1,
		"LoginRoomRetID":           2,
		"UserQueueID":              3,
		"GameReadyID":              4,
		"RequestGameVerID":         5,
		"GameVerInfoID":            6,
		"UserQueueAfterQuitID":     11,
		"UserForceLeaveRoomID":     12,
		"SlotGetGameProgressID":    13,
		"SlotGetGameProgressRetID": 14,
		"SlotSaveGameProgressID":   15,
		"SlotGetJackpotID":         16,
		"SlotGetJackpotRetID":      17,
		"SlotUpdateJackpotID":      18,
		"UserSitTableID":           19,
		"ShowTableListID":          20,
		"UserLeaveTableID":         21,
		"UserLeaveTableRetID":      22,
		"SlotGetJackpotGroupID":    23,
		"SlotGetJackpotGroupRetID": 24,
		"SlotUpdateJackpotGroupID": 25,
		"UserSendVipBrowReqID":     26,
		"UserSendVipBrowRetID":     27,
		"BroadcastUserVipBrowID":   28,
		"AddUserTaskGameEvent":     29,
		"GetUserSingleControl":     30,
		"GetSinglePool":            31,
		"UpdateSinglePool":         32,
		"GetUserReelsControl":      33,
		"GetReelsControlCYZS":      34,
		"GetReelsControlCSFFF":     35,
		"GameRoomEmoji":            36,
		"GameRoomEmojiRet":         37,
		"DzpkChoiceTable":          38,
		"DzpkChoiceTableRet":       39,
		"DzpkChoiceSeat":           40,
		"DzpkChoiceSeatRet":        41,
		"DzpkSitSeatTimeout":       42,
		"DzpkBuyChip":              43,
	}
)

func (x GameRoomClassID) Enum() *GameRoomClassID {
	p := new(GameRoomClassID)
	*p = x
	return p
}

func (x GameRoomClassID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameRoomClassID) Descriptor() protoreflect.EnumDescriptor {
	return file_GameRoom_proto_enumTypes[0].Descriptor()
}

func (GameRoomClassID) Type() protoreflect.EnumType {
	return &file_GameRoom_proto_enumTypes[0]
}

func (x GameRoomClassID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *GameRoomClassID) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = GameRoomClassID(num)
	return nil
}

// Deprecated: Use GameRoomClassID.Descriptor instead.
func (GameRoomClassID) EnumDescriptor() ([]byte, []int) {
	return file_GameRoom_proto_rawDescGZIP(), []int{0}
}

var File_GameRoom_proto protoreflect.FileDescriptor

var file_GameRoom_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xa4, 0x07, 0x0a, 0x0f, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x12, 0x0f,
	0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x10, 0x01, 0x12,
	0x12, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x74, 0x49,
	0x44, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x49, 0x44, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x61, 0x64,
	0x79, 0x49, 0x44, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x47, 0x61, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x49, 0x44, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x47,
	0x61, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x10, 0x06, 0x12, 0x18,
	0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x75, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72,
	0x51, 0x75, 0x69, 0x74, 0x49, 0x44, 0x10, 0x0b, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72,
	0x46, 0x6f, 0x72, 0x63, 0x65, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44,
	0x10, 0x0c, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x6c, 0x6f, 0x74, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d,
	0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x44, 0x10, 0x0d, 0x12, 0x1c, 0x0a,
	0x18, 0x53, 0x6c, 0x6f, 0x74, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x74, 0x49, 0x44, 0x10, 0x0e, 0x12, 0x1a, 0x0a, 0x16, 0x53,
	0x6c, 0x6f, 0x74, 0x53, 0x61, 0x76, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x49, 0x44, 0x10, 0x0f, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x6c, 0x6f, 0x74, 0x47,
	0x65, 0x74, 0x4a, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x74, 0x49, 0x44, 0x10, 0x10, 0x12, 0x17, 0x0a,
	0x13, 0x53, 0x6c, 0x6f, 0x74, 0x47, 0x65, 0x74, 0x4a, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x74, 0x52,
	0x65, 0x74, 0x49, 0x44, 0x10, 0x11, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x6c, 0x6f, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4a, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x74, 0x49, 0x44, 0x10, 0x12, 0x12,
	0x12, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49,
	0x44, 0x10, 0x13, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x77, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x44, 0x10, 0x14, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x65, 0x61, 0x76, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x44, 0x10, 0x15, 0x12, 0x17,
	0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x74, 0x49, 0x44, 0x10, 0x16, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x6c, 0x6f, 0x74, 0x47,
	0x65, 0x74, 0x4a, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44,
	0x10, 0x17, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x6c, 0x6f, 0x74, 0x47, 0x65, 0x74, 0x4a, 0x61, 0x63,
	0x6b, 0x70, 0x6f, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x74, 0x49, 0x44, 0x10, 0x18,
	0x12, 0x1c, 0x0a, 0x18, 0x53, 0x6c, 0x6f, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x61,
	0x63, 0x6b, 0x70, 0x6f, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x10, 0x19, 0x12, 0x18,
	0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x69, 0x70, 0x42, 0x72, 0x6f,
	0x77, 0x52, 0x65, 0x71, 0x49, 0x44, 0x10, 0x1a, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x6e, 0x64, 0x56, 0x69, 0x70, 0x42, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x74, 0x49, 0x44,
	0x10, 0x1b, 0x12, 0x1a, 0x0a, 0x16, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x56, 0x69, 0x70, 0x42, 0x72, 0x6f, 0x77, 0x49, 0x44, 0x10, 0x1c, 0x12, 0x18,
	0x0a, 0x14, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x47, 0x61, 0x6d,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x1d, 0x12, 0x18, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x10, 0x1e, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50,
	0x6f, 0x6f, 0x6c, 0x10, 0x1f, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x10, 0x20, 0x12, 0x17, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x65, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x10, 0x21, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x65, 0x6c, 0x73,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x43, 0x59, 0x5a, 0x53, 0x10, 0x22, 0x12, 0x18, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x65, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x43, 0x53, 0x46, 0x46, 0x46, 0x10, 0x23, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x61, 0x6d, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x10, 0x24, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65, 0x74, 0x10, 0x25,
	0x12, 0x13, 0x0a, 0x0f, 0x44, 0x7a, 0x70, 0x6b, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x10, 0x26, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x7a, 0x70, 0x6b, 0x43, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x74, 0x10, 0x27, 0x12, 0x12, 0x0a,
	0x0e, 0x44, 0x7a, 0x70, 0x6b, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x61, 0x74, 0x10,
	0x28, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x7a, 0x70, 0x6b, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x53,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x74, 0x10, 0x29, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x7a, 0x70, 0x6b,
	0x53, 0x69, 0x74, 0x53, 0x65, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x10, 0x2a,
	0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x7a, 0x70, 0x6b, 0x42, 0x75, 0x79, 0x43, 0x68, 0x69, 0x70, 0x10,
	0x2b,
}

var (
	file_GameRoom_proto_rawDescOnce sync.Once
	file_GameRoom_proto_rawDescData = file_GameRoom_proto_rawDesc
)

func file_GameRoom_proto_rawDescGZIP() []byte {
	file_GameRoom_proto_rawDescOnce.Do(func() {
		file_GameRoom_proto_rawDescData = protoimpl.X.CompressGZIP(file_GameRoom_proto_rawDescData)
	})
	return file_GameRoom_proto_rawDescData
}

var file_GameRoom_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GameRoom_proto_goTypes = []interface{}{
	(GameRoomClassID)(0), // 0: netproto.GameRoomClassID
}
var file_GameRoom_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GameRoom_proto_init() }
func file_GameRoom_proto_init() {
	if File_GameRoom_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GameRoom_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GameRoom_proto_goTypes,
		DependencyIndexes: file_GameRoom_proto_depIdxs,
		EnumInfos:         file_GameRoom_proto_enumTypes,
	}.Build()
	File_GameRoom_proto = out.File
	file_GameRoom_proto_rawDesc = nil
	file_GameRoom_proto_goTypes = nil
	file_GameRoom_proto_depIdxs = nil
}
