// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: HttpMsg_BJL.proto

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

type Http_BJL_GameInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableList []*Http_BJL_TableList `protobuf:"bytes,1,rep,name=TableList" json:"TableList,omitempty"` //桌子信息
}

func (x *Http_BJL_GameInfo) Reset() {
	*x = Http_BJL_GameInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HttpMsg_BJL_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http_BJL_GameInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http_BJL_GameInfo) ProtoMessage() {}

func (x *Http_BJL_GameInfo) ProtoReflect() protoreflect.Message {
	mi := &file_HttpMsg_BJL_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http_BJL_GameInfo.ProtoReflect.Descriptor instead.
func (*Http_BJL_GameInfo) Descriptor() ([]byte, []int) {
	return file_HttpMsg_BJL_proto_rawDescGZIP(), []int{0}
}

func (x *Http_BJL_GameInfo) GetTableList() []*Http_BJL_TableList {
	if x != nil {
		return x.TableList
	}
	return nil
}

type Http_BJL_TableList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableNo           *int32 `protobuf:"varint,1,req,name=TableNo" json:"TableNo,omitempty"`                     //桌子号
	Status            *int32 `protobuf:"varint,2,req,name=Status" json:"Status,omitempty"`                       //当前状态 0 下注中 1下注结束动画 2开奖中 3开奖动画 4结算动画 5等待开始游戏
	GameUserCount     *int32 `protobuf:"varint,3,req,name=GameUserCount" json:"GameUserCount,omitempty"`         //玩家人数(包含未下注玩家)
	RobotCount        *int32 `protobuf:"varint,4,req,name=RobotCount" json:"RobotCount,omitempty"`               //机器人数(包含未下注机器人)
	SumEarnAmount     *int64 `protobuf:"varint,5,req,name=SumEarnAmount" json:"SumEarnAmount,omitempty"`         //系统从开服到现在的总收益
	LastEarnAmount    *int32 `protobuf:"varint,6,req,name=LastEarnAmount" json:"LastEarnAmount,omitempty"`       //最后一局收益
	PoolCurrentAmount *int64 `protobuf:"varint,7,req,name=PoolCurrentAmount" json:"PoolCurrentAmount,omitempty"` //当前水池金币数
}

func (x *Http_BJL_TableList) Reset() {
	*x = Http_BJL_TableList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HttpMsg_BJL_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http_BJL_TableList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http_BJL_TableList) ProtoMessage() {}

func (x *Http_BJL_TableList) ProtoReflect() protoreflect.Message {
	mi := &file_HttpMsg_BJL_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http_BJL_TableList.ProtoReflect.Descriptor instead.
func (*Http_BJL_TableList) Descriptor() ([]byte, []int) {
	return file_HttpMsg_BJL_proto_rawDescGZIP(), []int{1}
}

func (x *Http_BJL_TableList) GetTableNo() int32 {
	if x != nil && x.TableNo != nil {
		return *x.TableNo
	}
	return 0
}

func (x *Http_BJL_TableList) GetStatus() int32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *Http_BJL_TableList) GetGameUserCount() int32 {
	if x != nil && x.GameUserCount != nil {
		return *x.GameUserCount
	}
	return 0
}

func (x *Http_BJL_TableList) GetRobotCount() int32 {
	if x != nil && x.RobotCount != nil {
		return *x.RobotCount
	}
	return 0
}

func (x *Http_BJL_TableList) GetSumEarnAmount() int64 {
	if x != nil && x.SumEarnAmount != nil {
		return *x.SumEarnAmount
	}
	return 0
}

func (x *Http_BJL_TableList) GetLastEarnAmount() int32 {
	if x != nil && x.LastEarnAmount != nil {
		return *x.LastEarnAmount
	}
	return 0
}

func (x *Http_BJL_TableList) GetPoolCurrentAmount() int64 {
	if x != nil && x.PoolCurrentAmount != nil {
		return *x.PoolCurrentAmount
	}
	return 0
}

type Http_BJL_TableDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableNo                  *int32               `protobuf:"varint,1,req,name=TableNo" json:"TableNo,omitempty"`                                    //桌子号
	Status                   *int32               `protobuf:"varint,2,req,name=Status" json:"Status,omitempty"`                                      //当前状态 0 下注中 1下注结束动画 2开奖中 3开奖动画 4结算动画 5等待开始游戏
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
	OpenResultZhuang         *int64               `protobuf:"varint,15,req,name=OpenResultZhuang" json:"OpenResultZhuang,omitempty"`                 //当前开奖结果为庄的收益数量
	OpenResultXian           *int64               `protobuf:"varint,16,req,name=OpenResultXian" json:"OpenResultXian,omitempty"`                     //当前开奖结果为闲的收益数量
	OpenResultHe             *int64               `protobuf:"varint,17,req,name=OpenResultHe" json:"OpenResultHe,omitempty"`                         //当前开奖结果为和的收益数量
	CurentOpenResult         *string              `protobuf:"bytes,18,req,name=CurentOpenResult" json:"CurentOpenResult,omitempty"`                  //当前开奖结果,只有在开奖动画状态下有值
	CurentOpenResultEarn     *int64               `protobuf:"varint,19,req,name=CurentOpenResultEarn" json:"CurentOpenResultEarn,omitempty"`         //当前开奖收益,只有在开奖动画状态下有值
	UserInfo                 []*Http_BJL_UserInfo `protobuf:"bytes,20,rep,name=UserInfo" json:"UserInfo,omitempty"`                                  //玩家信息
}

func (x *Http_BJL_TableDetail) Reset() {
	*x = Http_BJL_TableDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HttpMsg_BJL_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http_BJL_TableDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http_BJL_TableDetail) ProtoMessage() {}

func (x *Http_BJL_TableDetail) ProtoReflect() protoreflect.Message {
	mi := &file_HttpMsg_BJL_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http_BJL_TableDetail.ProtoReflect.Descriptor instead.
func (*Http_BJL_TableDetail) Descriptor() ([]byte, []int) {
	return file_HttpMsg_BJL_proto_rawDescGZIP(), []int{2}
}

func (x *Http_BJL_TableDetail) GetTableNo() int32 {
	if x != nil && x.TableNo != nil {
		return *x.TableNo
	}
	return 0
}

func (x *Http_BJL_TableDetail) GetStatus() int32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *Http_BJL_TableDetail) GetSumEarnAmount() int64 {
	if x != nil && x.SumEarnAmount != nil {
		return *x.SumEarnAmount
	}
	return 0
}

func (x *Http_BJL_TableDetail) GetLastEarnAmount() int32 {
	if x != nil && x.LastEarnAmount != nil {
		return *x.LastEarnAmount
	}
	return 0
}

func (x *Http_BJL_TableDetail) GetPoolCurrentAmount() int64 {
	if x != nil && x.PoolCurrentAmount != nil {
		return *x.PoolCurrentAmount
	}
	return 0
}

func (x *Http_BJL_TableDetail) GetGameUserCount() int32 {
	if x != nil && x.GameUserCount != nil {
		return *x.GameUserCount
	}
	return 0
}

func (x *Http_BJL_TableDetail) GetRobotCount() int32 {
	if x != nil && x.RobotCount != nil {
		return *x.RobotCount
	}
	return 0
}

func (x *Http_BJL_TableDetail) GetPopWaterRate() int32 {
	if x != nil && x.PopWaterRate != nil {
		return *x.PopWaterRate
	}
	return 0
}

func (x *Http_BJL_TableDetail) GetZhuangType() int32 {
	if x != nil && x.ZhuangType != nil {
		return *x.ZhuangType
	}
	return 0
}

func (x *Http_BJL_TableDetail) GetZhuangNickName() string {
	if x != nil && x.ZhuangNickName != nil {
		return *x.ZhuangNickName
	}
	return ""
}

func (x *Http_BJL_TableDetail) GetZhuangUserID() int32 {
	if x != nil && x.ZhuangUserID != nil {
		return *x.ZhuangUserID
	}
	return 0
}

func (x *Http_BJL_TableDetail) GetZhuangSurplusTotal() int32 {
	if x != nil && x.ZhuangSurplusTotal != nil {
		return *x.ZhuangSurplusTotal
	}
	return 0
}

func (x *Http_BJL_TableDetail) GetZhuangCurrentMoneyAmount() int64 {
	if x != nil && x.ZhuangCurrentMoneyAmount != nil {
		return *x.ZhuangCurrentMoneyAmount
	}
	return 0
}

func (x *Http_BJL_TableDetail) GetZhuangInitMoneyAmount() int64 {
	if x != nil && x.ZhuangInitMoneyAmount != nil {
		return *x.ZhuangInitMoneyAmount
	}
	return 0
}

func (x *Http_BJL_TableDetail) GetOpenResultZhuang() int64 {
	if x != nil && x.OpenResultZhuang != nil {
		return *x.OpenResultZhuang
	}
	return 0
}

func (x *Http_BJL_TableDetail) GetOpenResultXian() int64 {
	if x != nil && x.OpenResultXian != nil {
		return *x.OpenResultXian
	}
	return 0
}

func (x *Http_BJL_TableDetail) GetOpenResultHe() int64 {
	if x != nil && x.OpenResultHe != nil {
		return *x.OpenResultHe
	}
	return 0
}

func (x *Http_BJL_TableDetail) GetCurentOpenResult() string {
	if x != nil && x.CurentOpenResult != nil {
		return *x.CurentOpenResult
	}
	return ""
}

func (x *Http_BJL_TableDetail) GetCurentOpenResultEarn() int64 {
	if x != nil && x.CurentOpenResultEarn != nil {
		return *x.CurentOpenResultEarn
	}
	return 0
}

func (x *Http_BJL_TableDetail) GetUserInfo() []*Http_BJL_UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type Http_BJL_UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID          *int32  `protobuf:"varint,1,req,name=UserID" json:"UserID,omitempty"`                    //玩家ID
	UserType        *int32  `protobuf:"varint,2,req,name=UserType" json:"UserType,omitempty"`                //玩家类型 0机器人 1玩家
	NickName        *string `protobuf:"bytes,3,req,name=NickName" json:"NickName,omitempty"`                 //玩家昵称
	GameScore       *int64  `protobuf:"varint,4,req,name=GameScore" json:"GameScore,omitempty"`              //游戏总输赢
	InitMoney       *int64  `protobuf:"varint,5,req,name=InitMoney" json:"InitMoney,omitempty"`              //初始金币
	CurrentMoney    *int64  `protobuf:"varint,6,req,name=CurrentMoney" json:"CurrentMoney,omitempty"`        //当前金币
	BetInfoZhuang   *int64  `protobuf:"varint,7,req,name=BetInfoZhuang" json:"BetInfoZhuang,omitempty"`      //下注庄的金币数量
	BetInfoZhuangDZ *int64  `protobuf:"varint,8,req,name=BetInfoZhuangDZ" json:"BetInfoZhuangDZ,omitempty"`  //下注庄对子的金币数量
	BetInfoZhuangTW *int64  `protobuf:"varint,9,req,name=BetInfoZhuangTW" json:"BetInfoZhuangTW,omitempty"`  //下注庄天王的金币数量
	BetInfoZhuangSZ *int64  `protobuf:"varint,10,req,name=BetInfoZhuangSZ" json:"BetInfoZhuangSZ,omitempty"` //下注庄顺子的金币数量
	BetInfoXian     *int64  `protobuf:"varint,11,req,name=BetInfoXian" json:"BetInfoXian,omitempty"`         //下注闲的金币数量
	BetInfoXianDZ   *int64  `protobuf:"varint,12,req,name=BetInfoXianDZ" json:"BetInfoXianDZ,omitempty"`     //下注闲对子的金币数量
	BetInfoXianTW   *int64  `protobuf:"varint,13,req,name=BetInfoXianTW" json:"BetInfoXianTW,omitempty"`     //下注闲天王的金币数量
	BetInfoXianSZ   *int64  `protobuf:"varint,14,req,name=BetInfoXianSZ" json:"BetInfoXianSZ,omitempty"`     //下注闲顺子的金币数量
	BetInfoHe       *int64  `protobuf:"varint,15,req,name=BetInfoHe" json:"BetInfoHe,omitempty"`             //下注和的金币数量
	BetInfoHeTD     *int64  `protobuf:"varint,16,req,name=BetInfoHeTD" json:"BetInfoHeTD,omitempty"`         //下注同点和的金币数量
	BetInfoBig      *int64  `protobuf:"varint,17,opt,name=BetInfoBig" json:"BetInfoBig,omitempty"`           //下注大
	BetInfoSmall    *int64  `protobuf:"varint,18,opt,name=BetInfoSmall" json:"BetInfoSmall,omitempty"`       //下注小
}

func (x *Http_BJL_UserInfo) Reset() {
	*x = Http_BJL_UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HttpMsg_BJL_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http_BJL_UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http_BJL_UserInfo) ProtoMessage() {}

func (x *Http_BJL_UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_HttpMsg_BJL_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http_BJL_UserInfo.ProtoReflect.Descriptor instead.
func (*Http_BJL_UserInfo) Descriptor() ([]byte, []int) {
	return file_HttpMsg_BJL_proto_rawDescGZIP(), []int{3}
}

func (x *Http_BJL_UserInfo) GetUserID() int32 {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return 0
}

func (x *Http_BJL_UserInfo) GetUserType() int32 {
	if x != nil && x.UserType != nil {
		return *x.UserType
	}
	return 0
}

func (x *Http_BJL_UserInfo) GetNickName() string {
	if x != nil && x.NickName != nil {
		return *x.NickName
	}
	return ""
}

func (x *Http_BJL_UserInfo) GetGameScore() int64 {
	if x != nil && x.GameScore != nil {
		return *x.GameScore
	}
	return 0
}

func (x *Http_BJL_UserInfo) GetInitMoney() int64 {
	if x != nil && x.InitMoney != nil {
		return *x.InitMoney
	}
	return 0
}

func (x *Http_BJL_UserInfo) GetCurrentMoney() int64 {
	if x != nil && x.CurrentMoney != nil {
		return *x.CurrentMoney
	}
	return 0
}

func (x *Http_BJL_UserInfo) GetBetInfoZhuang() int64 {
	if x != nil && x.BetInfoZhuang != nil {
		return *x.BetInfoZhuang
	}
	return 0
}

func (x *Http_BJL_UserInfo) GetBetInfoZhuangDZ() int64 {
	if x != nil && x.BetInfoZhuangDZ != nil {
		return *x.BetInfoZhuangDZ
	}
	return 0
}

func (x *Http_BJL_UserInfo) GetBetInfoZhuangTW() int64 {
	if x != nil && x.BetInfoZhuangTW != nil {
		return *x.BetInfoZhuangTW
	}
	return 0
}

func (x *Http_BJL_UserInfo) GetBetInfoZhuangSZ() int64 {
	if x != nil && x.BetInfoZhuangSZ != nil {
		return *x.BetInfoZhuangSZ
	}
	return 0
}

func (x *Http_BJL_UserInfo) GetBetInfoXian() int64 {
	if x != nil && x.BetInfoXian != nil {
		return *x.BetInfoXian
	}
	return 0
}

func (x *Http_BJL_UserInfo) GetBetInfoXianDZ() int64 {
	if x != nil && x.BetInfoXianDZ != nil {
		return *x.BetInfoXianDZ
	}
	return 0
}

func (x *Http_BJL_UserInfo) GetBetInfoXianTW() int64 {
	if x != nil && x.BetInfoXianTW != nil {
		return *x.BetInfoXianTW
	}
	return 0
}

func (x *Http_BJL_UserInfo) GetBetInfoXianSZ() int64 {
	if x != nil && x.BetInfoXianSZ != nil {
		return *x.BetInfoXianSZ
	}
	return 0
}

func (x *Http_BJL_UserInfo) GetBetInfoHe() int64 {
	if x != nil && x.BetInfoHe != nil {
		return *x.BetInfoHe
	}
	return 0
}

func (x *Http_BJL_UserInfo) GetBetInfoHeTD() int64 {
	if x != nil && x.BetInfoHeTD != nil {
		return *x.BetInfoHeTD
	}
	return 0
}

func (x *Http_BJL_UserInfo) GetBetInfoBig() int64 {
	if x != nil && x.BetInfoBig != nil {
		return *x.BetInfoBig
	}
	return 0
}

func (x *Http_BJL_UserInfo) GetBetInfoSmall() int64 {
	if x != nil && x.BetInfoSmall != nil {
		return *x.BetInfoSmall
	}
	return 0
}

var File_HttpMsg_BJL_proto protoreflect.FileDescriptor

var file_HttpMsg_BJL_proto_rawDesc = []byte{
	0x0a, 0x11, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x73, 0x67, 0x5f, 0x42, 0x4a, 0x4c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a,
	0x11, 0x48, 0x74, 0x74, 0x70, 0x5f, 0x42, 0x4a, 0x4c, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x5f, 0x42, 0x4a, 0x4c, 0x5f, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0xaa, 0x01, 0x0a, 0x12, 0x48, 0x74, 0x74, 0x70, 0x5f, 0x42, 0x4a, 0x4c,
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
	0x22, 0x90, 0x04, 0x0a, 0x14, 0x48, 0x74, 0x74, 0x70, 0x5f, 0x42, 0x4a, 0x4c, 0x5f, 0x54, 0x61,
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
	0x6c, 0x74, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x18, 0x0f, 0x20, 0x02, 0x28, 0x03, 0x12, 0x16,
	0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x58, 0x69, 0x61, 0x6e,
	0x18, 0x10, 0x20, 0x02, 0x28, 0x03, 0x12, 0x14, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x48, 0x65, 0x18, 0x11, 0x20, 0x02, 0x28, 0x03, 0x12, 0x18, 0x0a, 0x10,
	0x43, 0x75, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x12, 0x20, 0x02, 0x28, 0x09, 0x12, 0x1c, 0x0a, 0x14, 0x43, 0x75, 0x72, 0x65, 0x6e, 0x74,
	0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x61, 0x72, 0x6e, 0x18, 0x13,
	0x20, 0x02, 0x28, 0x03, 0x12, 0x2d, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x5f, 0x42, 0x4a, 0x4c, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x91, 0x03, 0x0a, 0x11, 0x48, 0x74, 0x74, 0x70, 0x5f, 0x42, 0x4a, 0x4c,
	0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x12, 0x10, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x12, 0x10, 0x0a, 0x08, 0x4e,
	0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x12, 0x11, 0x0a,
	0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x02, 0x28, 0x03,
	0x12, 0x11, 0x0a, 0x09, 0x49, 0x6e, 0x69, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x05, 0x20,
	0x02, 0x28, 0x03, 0x12, 0x14, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x6f,
	0x6e, 0x65, 0x79, 0x18, 0x06, 0x20, 0x02, 0x28, 0x03, 0x12, 0x15, 0x0a, 0x0d, 0x42, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x02, 0x28, 0x03,
	0x12, 0x17, 0x0a, 0x0f, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x5a, 0x68, 0x75, 0x61, 0x6e,
	0x67, 0x44, 0x5a, 0x18, 0x08, 0x20, 0x02, 0x28, 0x03, 0x12, 0x17, 0x0a, 0x0f, 0x42, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x54, 0x57, 0x18, 0x09, 0x20, 0x02,
	0x28, 0x03, 0x12, 0x17, 0x0a, 0x0f, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x5a, 0x68, 0x75,
	0x61, 0x6e, 0x67, 0x53, 0x5a, 0x18, 0x0a, 0x20, 0x02, 0x28, 0x03, 0x12, 0x13, 0x0a, 0x0b, 0x42,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x58, 0x69, 0x61, 0x6e, 0x18, 0x0b, 0x20, 0x02, 0x28, 0x03,
	0x12, 0x15, 0x0a, 0x0d, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x58, 0x69, 0x61, 0x6e, 0x44,
	0x5a, 0x18, 0x0c, 0x20, 0x02, 0x28, 0x03, 0x12, 0x15, 0x0a, 0x0d, 0x42, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x58, 0x69, 0x61, 0x6e, 0x54, 0x57, 0x18, 0x0d, 0x20, 0x02, 0x28, 0x03, 0x12, 0x15,
	0x0a, 0x0d, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x58, 0x69, 0x61, 0x6e, 0x53, 0x5a, 0x18,
	0x0e, 0x20, 0x02, 0x28, 0x03, 0x12, 0x11, 0x0a, 0x09, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x48, 0x65, 0x18, 0x0f, 0x20, 0x02, 0x28, 0x03, 0x12, 0x13, 0x0a, 0x0b, 0x42, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x48, 0x65, 0x54, 0x44, 0x18, 0x10, 0x20, 0x02, 0x28, 0x03, 0x12, 0x12, 0x0a,
	0x0a, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x69, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x03, 0x12, 0x14, 0x0a, 0x0c, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x6d, 0x61, 0x6c,
	0x6c, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03,
}

var (
	file_HttpMsg_BJL_proto_rawDescOnce sync.Once
	file_HttpMsg_BJL_proto_rawDescData = file_HttpMsg_BJL_proto_rawDesc
)

func file_HttpMsg_BJL_proto_rawDescGZIP() []byte {
	file_HttpMsg_BJL_proto_rawDescOnce.Do(func() {
		file_HttpMsg_BJL_proto_rawDescData = protoimpl.X.CompressGZIP(file_HttpMsg_BJL_proto_rawDescData)
	})
	return file_HttpMsg_BJL_proto_rawDescData
}

var file_HttpMsg_BJL_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_HttpMsg_BJL_proto_goTypes = []interface{}{
	(*Http_BJL_GameInfo)(nil),    // 0: netproto.Http_BJL_GameInfo
	(*Http_BJL_TableList)(nil),   // 1: netproto.Http_BJL_TableList
	(*Http_BJL_TableDetail)(nil), // 2: netproto.Http_BJL_TableDetail
	(*Http_BJL_UserInfo)(nil),    // 3: netproto.Http_BJL_UserInfo
}
var file_HttpMsg_BJL_proto_depIdxs = []int32{
	1, // 0: netproto.Http_BJL_GameInfo.TableList:type_name -> netproto.Http_BJL_TableList
	3, // 1: netproto.Http_BJL_TableDetail.UserInfo:type_name -> netproto.Http_BJL_UserInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_HttpMsg_BJL_proto_init() }
func file_HttpMsg_BJL_proto_init() {
	if File_HttpMsg_BJL_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_HttpMsg_BJL_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http_BJL_GameInfo); i {
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
		file_HttpMsg_BJL_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http_BJL_TableList); i {
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
		file_HttpMsg_BJL_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http_BJL_TableDetail); i {
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
		file_HttpMsg_BJL_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http_BJL_UserInfo); i {
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
			RawDescriptor: file_HttpMsg_BJL_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HttpMsg_BJL_proto_goTypes,
		DependencyIndexes: file_HttpMsg_BJL_proto_depIdxs,
		MessageInfos:      file_HttpMsg_BJL_proto_msgTypes,
	}.Build()
	File_HttpMsg_BJL_proto = out.File
	file_HttpMsg_BJL_proto_rawDesc = nil
	file_HttpMsg_BJL_proto_goTypes = nil
	file_HttpMsg_BJL_proto_depIdxs = nil
}
