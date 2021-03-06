// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: HttpMsg_LHD.proto

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

type Http_LHD_GameInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableList []*Http_LHD_TableList `protobuf:"bytes,1,rep,name=TableList" json:"TableList,omitempty"` //桌子信息
}

func (x *Http_LHD_GameInfo) Reset() {
	*x = Http_LHD_GameInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HttpMsg_LHD_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http_LHD_GameInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http_LHD_GameInfo) ProtoMessage() {}

func (x *Http_LHD_GameInfo) ProtoReflect() protoreflect.Message {
	mi := &file_HttpMsg_LHD_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http_LHD_GameInfo.ProtoReflect.Descriptor instead.
func (*Http_LHD_GameInfo) Descriptor() ([]byte, []int) {
	return file_HttpMsg_LHD_proto_rawDescGZIP(), []int{0}
}

func (x *Http_LHD_GameInfo) GetTableList() []*Http_LHD_TableList {
	if x != nil {
		return x.TableList
	}
	return nil
}

type Http_LHD_TableList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableNo           *int32 `protobuf:"varint,1,req,name=TableNo" json:"TableNo,omitempty"`                     //桌子号
	Status            *int32 `protobuf:"varint,2,req,name=Status" json:"Status,omitempty"`                       //当前状态 0 下注中 1 开奖中 2开奖动画 3结算动画 4等待开始游戏
	GameUserCount     *int32 `protobuf:"varint,3,req,name=GameUserCount" json:"GameUserCount,omitempty"`         //玩家人数(包含未下注玩家)
	RobotCount        *int32 `protobuf:"varint,4,req,name=RobotCount" json:"RobotCount,omitempty"`               //机器人数(包含未下注机器人)
	SumEarnAmount     *int64 `protobuf:"varint,5,req,name=SumEarnAmount" json:"SumEarnAmount,omitempty"`         //系统从开服到现在的总收益
	LastEarnAmount    *int32 `protobuf:"varint,6,req,name=LastEarnAmount" json:"LastEarnAmount,omitempty"`       //最后一局收益
	PoolCurrentAmount *int64 `protobuf:"varint,7,req,name=PoolCurrentAmount" json:"PoolCurrentAmount,omitempty"` //当前水池金币数
}

func (x *Http_LHD_TableList) Reset() {
	*x = Http_LHD_TableList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HttpMsg_LHD_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http_LHD_TableList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http_LHD_TableList) ProtoMessage() {}

func (x *Http_LHD_TableList) ProtoReflect() protoreflect.Message {
	mi := &file_HttpMsg_LHD_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http_LHD_TableList.ProtoReflect.Descriptor instead.
func (*Http_LHD_TableList) Descriptor() ([]byte, []int) {
	return file_HttpMsg_LHD_proto_rawDescGZIP(), []int{1}
}

func (x *Http_LHD_TableList) GetTableNo() int32 {
	if x != nil && x.TableNo != nil {
		return *x.TableNo
	}
	return 0
}

func (x *Http_LHD_TableList) GetStatus() int32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *Http_LHD_TableList) GetGameUserCount() int32 {
	if x != nil && x.GameUserCount != nil {
		return *x.GameUserCount
	}
	return 0
}

func (x *Http_LHD_TableList) GetRobotCount() int32 {
	if x != nil && x.RobotCount != nil {
		return *x.RobotCount
	}
	return 0
}

func (x *Http_LHD_TableList) GetSumEarnAmount() int64 {
	if x != nil && x.SumEarnAmount != nil {
		return *x.SumEarnAmount
	}
	return 0
}

func (x *Http_LHD_TableList) GetLastEarnAmount() int32 {
	if x != nil && x.LastEarnAmount != nil {
		return *x.LastEarnAmount
	}
	return 0
}

func (x *Http_LHD_TableList) GetPoolCurrentAmount() int64 {
	if x != nil && x.PoolCurrentAmount != nil {
		return *x.PoolCurrentAmount
	}
	return 0
}

type Http_LHD_TableDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableNo                  *int32               `protobuf:"varint,1,req,name=TableNo" json:"TableNo,omitempty"`                                    //桌子号
	Status                   *int32               `protobuf:"varint,2,req,name=Status" json:"Status,omitempty"`                                      //当前状态 0 下注中 1 开奖中 2开奖动画 3结算动画 4等待开始游戏
	SumEarnAmount            *int64               `protobuf:"varint,3,req,name=SumEarnAmount" json:"SumEarnAmount,omitempty"`                        //系统从开服到现在的总收益
	LastEarnAmount           *int32               `protobuf:"varint,4,req,name=LastEarnAmount" json:"LastEarnAmount,omitempty"`                      //最后一局收益
	PoolCurrentAmount        *int64               `protobuf:"varint,5,req,name=PoolCurrentAmount" json:"PoolCurrentAmount,omitempty"`                //当前水池金币数
	GameUserCount            *int32               `protobuf:"varint,6,req,name=GameUserCount" json:"GameUserCount,omitempty"`                        //游戏玩家人数(包含未下注玩家)
	RobotCount               *int32               `protobuf:"varint,7,req,name=RobotCount" json:"RobotCount,omitempty"`                              //机器人数(包含未下注机器人)
	PopWaterRate             *int32               `protobuf:"varint,8,req,name=PopWaterRate" json:"PopWaterRate,omitempty"`                          //当前注水率
	ZhuangType               *int32               `protobuf:"varint,9,req,name=ZhuangType" json:"ZhuangType,omitempty"`                              //庄家类型	0系统 1机器人当庄 2玩家上庄
	ZhuangNickName           *string              `protobuf:"bytes,10,req,name=ZhuangNickName" json:"ZhuangNickName,omitempty"`                      //庄家昵称
	ZhuangUserID             *int32               `protobuf:"varint,11,req,name=ZhuangUserID" json:"ZhuangUserID,omitempty"`                         //庄家用户ID
	ZhuangSurplusTotal       *int32               `protobuf:"varint,12,req,name=ZhuangSurplusTotal" json:"ZhuangSurplusTotal,omitempty"`             //庄家剩余局数
	ZhuangCurrentMoneyAmount *int64               `protobuf:"varint,13,req,name=ZhuangCurrentMoneyAmount" json:"ZhuangCurrentMoneyAmount,omitempty"` //庄家当前金币数 系统或机器人坐庄则表示水池金币数量
	ZhuangInitMoneyAmount    *int64               `protobuf:"varint,14,req,name=ZhuangInitMoneyAmount" json:"ZhuangInitMoneyAmount,omitempty"`       //庄家初始上庄金币数 系统或机器人坐庄则用0表示
	OpenResultDragon         *int64               `protobuf:"varint,15,req,name=OpenResultDragon" json:"OpenResultDragon,omitempty"`                 //当前开奖结果为龙的收益数量
	OpenResultTiger          *int64               `protobuf:"varint,16,req,name=OpenResultTiger" json:"OpenResultTiger,omitempty"`                   //当前开奖结果为虎的收益数量
	OpenResultDraw           *int64               `protobuf:"varint,17,req,name=OpenResultDraw" json:"OpenResultDraw,omitempty"`                     //当前开奖结果为和的收益数量
	CurentOpenResult         *string              `protobuf:"bytes,18,req,name=CurentOpenResult" json:"CurentOpenResult,omitempty"`                  //当前开奖结果,只有在开奖动画状态下有值
	CurentOpenResultEarn     *int64               `protobuf:"varint,19,req,name=CurentOpenResultEarn" json:"CurentOpenResultEarn,omitempty"`         //当前开奖收益,只有在开奖动画状态下有值
	UserInfo                 []*Http_LHD_UserInfo `protobuf:"bytes,20,rep,name=UserInfo" json:"UserInfo,omitempty"`                                  //玩家信息
}

func (x *Http_LHD_TableDetail) Reset() {
	*x = Http_LHD_TableDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HttpMsg_LHD_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http_LHD_TableDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http_LHD_TableDetail) ProtoMessage() {}

func (x *Http_LHD_TableDetail) ProtoReflect() protoreflect.Message {
	mi := &file_HttpMsg_LHD_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http_LHD_TableDetail.ProtoReflect.Descriptor instead.
func (*Http_LHD_TableDetail) Descriptor() ([]byte, []int) {
	return file_HttpMsg_LHD_proto_rawDescGZIP(), []int{2}
}

func (x *Http_LHD_TableDetail) GetTableNo() int32 {
	if x != nil && x.TableNo != nil {
		return *x.TableNo
	}
	return 0
}

func (x *Http_LHD_TableDetail) GetStatus() int32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *Http_LHD_TableDetail) GetSumEarnAmount() int64 {
	if x != nil && x.SumEarnAmount != nil {
		return *x.SumEarnAmount
	}
	return 0
}

func (x *Http_LHD_TableDetail) GetLastEarnAmount() int32 {
	if x != nil && x.LastEarnAmount != nil {
		return *x.LastEarnAmount
	}
	return 0
}

func (x *Http_LHD_TableDetail) GetPoolCurrentAmount() int64 {
	if x != nil && x.PoolCurrentAmount != nil {
		return *x.PoolCurrentAmount
	}
	return 0
}

func (x *Http_LHD_TableDetail) GetGameUserCount() int32 {
	if x != nil && x.GameUserCount != nil {
		return *x.GameUserCount
	}
	return 0
}

func (x *Http_LHD_TableDetail) GetRobotCount() int32 {
	if x != nil && x.RobotCount != nil {
		return *x.RobotCount
	}
	return 0
}

func (x *Http_LHD_TableDetail) GetPopWaterRate() int32 {
	if x != nil && x.PopWaterRate != nil {
		return *x.PopWaterRate
	}
	return 0
}

func (x *Http_LHD_TableDetail) GetZhuangType() int32 {
	if x != nil && x.ZhuangType != nil {
		return *x.ZhuangType
	}
	return 0
}

func (x *Http_LHD_TableDetail) GetZhuangNickName() string {
	if x != nil && x.ZhuangNickName != nil {
		return *x.ZhuangNickName
	}
	return ""
}

func (x *Http_LHD_TableDetail) GetZhuangUserID() int32 {
	if x != nil && x.ZhuangUserID != nil {
		return *x.ZhuangUserID
	}
	return 0
}

func (x *Http_LHD_TableDetail) GetZhuangSurplusTotal() int32 {
	if x != nil && x.ZhuangSurplusTotal != nil {
		return *x.ZhuangSurplusTotal
	}
	return 0
}

func (x *Http_LHD_TableDetail) GetZhuangCurrentMoneyAmount() int64 {
	if x != nil && x.ZhuangCurrentMoneyAmount != nil {
		return *x.ZhuangCurrentMoneyAmount
	}
	return 0
}

func (x *Http_LHD_TableDetail) GetZhuangInitMoneyAmount() int64 {
	if x != nil && x.ZhuangInitMoneyAmount != nil {
		return *x.ZhuangInitMoneyAmount
	}
	return 0
}

func (x *Http_LHD_TableDetail) GetOpenResultDragon() int64 {
	if x != nil && x.OpenResultDragon != nil {
		return *x.OpenResultDragon
	}
	return 0
}

func (x *Http_LHD_TableDetail) GetOpenResultTiger() int64 {
	if x != nil && x.OpenResultTiger != nil {
		return *x.OpenResultTiger
	}
	return 0
}

func (x *Http_LHD_TableDetail) GetOpenResultDraw() int64 {
	if x != nil && x.OpenResultDraw != nil {
		return *x.OpenResultDraw
	}
	return 0
}

func (x *Http_LHD_TableDetail) GetCurentOpenResult() string {
	if x != nil && x.CurentOpenResult != nil {
		return *x.CurentOpenResult
	}
	return ""
}

func (x *Http_LHD_TableDetail) GetCurentOpenResultEarn() int64 {
	if x != nil && x.CurentOpenResultEarn != nil {
		return *x.CurentOpenResultEarn
	}
	return 0
}

func (x *Http_LHD_TableDetail) GetUserInfo() []*Http_LHD_UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type Http_LHD_UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID        *int32  `protobuf:"varint,1,req,name=UserID" json:"UserID,omitempty"`               //玩家ID
	UserType      *int32  `protobuf:"varint,2,req,name=UserType" json:"UserType,omitempty"`           //玩家类型 0机器人 1玩家
	NickName      *string `protobuf:"bytes,3,req,name=NickName" json:"NickName,omitempty"`            //玩家昵称
	GameScore     *int64  `protobuf:"varint,4,req,name=GameScore" json:"GameScore,omitempty"`         //游戏总输赢
	InitMoney     *int64  `protobuf:"varint,5,req,name=InitMoney" json:"InitMoney,omitempty"`         //初始金币
	CurrentMoney  *int64  `protobuf:"varint,6,req,name=CurrentMoney" json:"CurrentMoney,omitempty"`   //当前金币
	BetInfoDragon *int64  `protobuf:"varint,7,req,name=BetInfoDragon" json:"BetInfoDragon,omitempty"` //下注龙的金币数量
	BetInfoTiger  *int64  `protobuf:"varint,8,req,name=BetInfoTiger" json:"BetInfoTiger,omitempty"`   //下注虎的金币数量
	BetInfoDraw   *int64  `protobuf:"varint,9,req,name=BetInfoDraw" json:"BetInfoDraw,omitempty"`     //下注和的金币数量
}

func (x *Http_LHD_UserInfo) Reset() {
	*x = Http_LHD_UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HttpMsg_LHD_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http_LHD_UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http_LHD_UserInfo) ProtoMessage() {}

func (x *Http_LHD_UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_HttpMsg_LHD_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http_LHD_UserInfo.ProtoReflect.Descriptor instead.
func (*Http_LHD_UserInfo) Descriptor() ([]byte, []int) {
	return file_HttpMsg_LHD_proto_rawDescGZIP(), []int{3}
}

func (x *Http_LHD_UserInfo) GetUserID() int32 {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return 0
}

func (x *Http_LHD_UserInfo) GetUserType() int32 {
	if x != nil && x.UserType != nil {
		return *x.UserType
	}
	return 0
}

func (x *Http_LHD_UserInfo) GetNickName() string {
	if x != nil && x.NickName != nil {
		return *x.NickName
	}
	return ""
}

func (x *Http_LHD_UserInfo) GetGameScore() int64 {
	if x != nil && x.GameScore != nil {
		return *x.GameScore
	}
	return 0
}

func (x *Http_LHD_UserInfo) GetInitMoney() int64 {
	if x != nil && x.InitMoney != nil {
		return *x.InitMoney
	}
	return 0
}

func (x *Http_LHD_UserInfo) GetCurrentMoney() int64 {
	if x != nil && x.CurrentMoney != nil {
		return *x.CurrentMoney
	}
	return 0
}

func (x *Http_LHD_UserInfo) GetBetInfoDragon() int64 {
	if x != nil && x.BetInfoDragon != nil {
		return *x.BetInfoDragon
	}
	return 0
}

func (x *Http_LHD_UserInfo) GetBetInfoTiger() int64 {
	if x != nil && x.BetInfoTiger != nil {
		return *x.BetInfoTiger
	}
	return 0
}

func (x *Http_LHD_UserInfo) GetBetInfoDraw() int64 {
	if x != nil && x.BetInfoDraw != nil {
		return *x.BetInfoDraw
	}
	return 0
}

var File_HttpMsg_LHD_proto protoreflect.FileDescriptor

var file_HttpMsg_LHD_proto_rawDesc = []byte{
	0x0a, 0x11, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x73, 0x67, 0x5f, 0x4c, 0x48, 0x44, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a,
	0x11, 0x48, 0x74, 0x74, 0x70, 0x5f, 0x4c, 0x48, 0x44, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x5f, 0x4c, 0x48, 0x44, 0x5f, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0xaa, 0x01, 0x0a, 0x12, 0x48, 0x74, 0x74, 0x70, 0x5f, 0x4c, 0x48, 0x44,
	0x5f, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x0a, 0x07, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x12, 0x0e, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x12, 0x15, 0x0a, 0x0d, 0x47,
	0x61, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x02,
	0x28, 0x05, 0x12, 0x12, 0x0a, 0x0a, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x02, 0x28, 0x05, 0x12, 0x15, 0x0a, 0x0d, 0x53, 0x75, 0x6d, 0x45, 0x61, 0x72,
	0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x02, 0x28, 0x03, 0x12, 0x16, 0x0a,
	0x0e, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x61, 0x72, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x02, 0x28, 0x05, 0x12, 0x19, 0x0a, 0x11, 0x50, 0x6f, 0x6f, 0x6c, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x02, 0x28, 0x03,
	0x22, 0x93, 0x04, 0x0a, 0x14, 0x48, 0x74, 0x74, 0x70, 0x5f, 0x4c, 0x48, 0x44, 0x5f, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0f, 0x0a, 0x07, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x12, 0x0e, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x12, 0x15, 0x0a, 0x0d, 0x53, 0x75,
	0x6d, 0x45, 0x61, 0x72, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x02, 0x28,
	0x03, 0x12, 0x16, 0x0a, 0x0e, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x61, 0x72, 0x6e, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x02, 0x28, 0x05, 0x12, 0x19, 0x0a, 0x11, 0x50, 0x6f, 0x6f,
	0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x02, 0x28, 0x03, 0x12, 0x15, 0x0a, 0x0d, 0x47, 0x61, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x02, 0x28, 0x05, 0x12, 0x12, 0x0a, 0x0a, 0x52,
	0x6f, 0x62, 0x6f, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x02, 0x28, 0x05, 0x12,
	0x14, 0x0a, 0x0c, 0x50, 0x6f, 0x70, 0x57, 0x61, 0x74, 0x65, 0x72, 0x52, 0x61, 0x74, 0x65, 0x18,
	0x08, 0x20, 0x02, 0x28, 0x05, 0x12, 0x12, 0x0a, 0x0a, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x02, 0x28, 0x05, 0x12, 0x16, 0x0a, 0x0e, 0x5a, 0x68, 0x75,
	0x61, 0x6e, 0x67, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x02, 0x28,
	0x09, 0x12, 0x14, 0x0a, 0x0c, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x0b, 0x20, 0x02, 0x28, 0x05, 0x12, 0x1a, 0x0a, 0x12, 0x5a, 0x68, 0x75, 0x61, 0x6e,
	0x67, 0x53, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x0c, 0x20,
	0x02, 0x28, 0x05, 0x12, 0x20, 0x0a, 0x18, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0d, 0x20, 0x02, 0x28, 0x03, 0x12, 0x1d, 0x0a, 0x15, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x49,
	0x6e, 0x69, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e,
	0x20, 0x02, 0x28, 0x03, 0x12, 0x18, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x44, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x02, 0x28, 0x03, 0x12, 0x17,
	0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x69, 0x67, 0x65,
	0x72, 0x18, 0x10, 0x20, 0x02, 0x28, 0x03, 0x12, 0x16, 0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x72, 0x61, 0x77, 0x18, 0x11, 0x20, 0x02, 0x28, 0x03, 0x12,
	0x18, 0x0a, 0x10, 0x43, 0x75, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x12, 0x20, 0x02, 0x28, 0x09, 0x12, 0x1c, 0x0a, 0x14, 0x43, 0x75, 0x72,
	0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x61, 0x72,
	0x6e, 0x18, 0x13, 0x20, 0x02, 0x28, 0x03, 0x12, 0x2d, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x65, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x5f, 0x4c, 0x48, 0x44, 0x5f, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x11, 0x48, 0x74, 0x74, 0x70, 0x5f,
	0x4c, 0x48, 0x44, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x12, 0x10, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x12, 0x10,
	0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09,
	0x12, 0x11, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20,
	0x02, 0x28, 0x03, 0x12, 0x11, 0x0a, 0x09, 0x49, 0x6e, 0x69, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x18, 0x05, 0x20, 0x02, 0x28, 0x03, 0x12, 0x14, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x06, 0x20, 0x02, 0x28, 0x03, 0x12, 0x15, 0x0a, 0x0d,
	0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x02, 0x28, 0x03, 0x12, 0x14, 0x0a, 0x0c, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x54, 0x69,
	0x67, 0x65, 0x72, 0x18, 0x08, 0x20, 0x02, 0x28, 0x03, 0x12, 0x13, 0x0a, 0x0b, 0x42, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x44, 0x72, 0x61, 0x77, 0x18, 0x09, 0x20, 0x02, 0x28, 0x03,
}

var (
	file_HttpMsg_LHD_proto_rawDescOnce sync.Once
	file_HttpMsg_LHD_proto_rawDescData = file_HttpMsg_LHD_proto_rawDesc
)

func file_HttpMsg_LHD_proto_rawDescGZIP() []byte {
	file_HttpMsg_LHD_proto_rawDescOnce.Do(func() {
		file_HttpMsg_LHD_proto_rawDescData = protoimpl.X.CompressGZIP(file_HttpMsg_LHD_proto_rawDescData)
	})
	return file_HttpMsg_LHD_proto_rawDescData
}

var file_HttpMsg_LHD_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_HttpMsg_LHD_proto_goTypes = []interface{}{
	(*Http_LHD_GameInfo)(nil),    // 0: netproto.Http_LHD_GameInfo
	(*Http_LHD_TableList)(nil),   // 1: netproto.Http_LHD_TableList
	(*Http_LHD_TableDetail)(nil), // 2: netproto.Http_LHD_TableDetail
	(*Http_LHD_UserInfo)(nil),    // 3: netproto.Http_LHD_UserInfo
}
var file_HttpMsg_LHD_proto_depIdxs = []int32{
	1, // 0: netproto.Http_LHD_GameInfo.TableList:type_name -> netproto.Http_LHD_TableList
	3, // 1: netproto.Http_LHD_TableDetail.UserInfo:type_name -> netproto.Http_LHD_UserInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_HttpMsg_LHD_proto_init() }
func file_HttpMsg_LHD_proto_init() {
	if File_HttpMsg_LHD_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_HttpMsg_LHD_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http_LHD_GameInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_HttpMsg_LHD_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http_LHD_TableList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_HttpMsg_LHD_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http_LHD_TableDetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_HttpMsg_LHD_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http_LHD_UserInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_HttpMsg_LHD_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HttpMsg_LHD_proto_goTypes,
		DependencyIndexes: file_HttpMsg_LHD_proto_depIdxs,
		MessageInfos:      file_HttpMsg_LHD_proto_msgTypes,
	}.Build()
	File_HttpMsg_LHD_proto = out.File
	file_HttpMsg_LHD_proto_rawDesc = nil
	file_HttpMsg_LHD_proto_goTypes = nil
	file_HttpMsg_LHD_proto_depIdxs = nil
}
