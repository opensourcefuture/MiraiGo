// Code generated by yaprotoc. DO NOT EDIT.
// source: oidb/0x88d.proto

package oidb0x88d

import (
	oidb0xef0 "github.com/Mrs4s/MiraiGo/internal/protobuf/data/oidb/oidb0xef0"
	"github.com/pkg/errors"
	"github.com/segmentio/encoding/proto"
)

type GroupCardPrefix struct {
	Introduction []byte   `protobuf:"bytes,1,opt"`
	Prefix       [][]byte `protobuf:"bytes,2,rep"`
}

func (x *GroupCardPrefix) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type GroupExInfoOnly struct {
	TribeId          *uint32 `protobuf:"varint,1,opt"`
	MoneyForAddGroup *uint32 `protobuf:"varint,2,opt"`
}

func (x *GroupExInfoOnly) GetTribeId() uint32 {
	if x != nil && x.TribeId != nil {
		return *x.TribeId
	}
	return 0
}

func (x *GroupExInfoOnly) GetMoneyForAddGroup() uint32 {
	if x != nil && x.MoneyForAddGroup != nil {
		return *x.MoneyForAddGroup
	}
	return 0
}

func (x *GroupExInfoOnly) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type GroupGeoInfo struct {
	Owneruin   *uint64 `protobuf:"varint,1,opt"`
	Settime    *uint32 `protobuf:"varint,2,opt"`
	Cityid     *uint32 `protobuf:"varint,3,opt"`
	Longitude  *int64  `protobuf:"varint,4,opt"`
	Latitude   *int64  `protobuf:"varint,5,opt"`
	Geocontent []byte  `protobuf:"bytes,6,opt"`
	PoiId      *uint64 `protobuf:"varint,7,opt"`
}

func (x *GroupGeoInfo) GetOwneruin() uint64 {
	if x != nil && x.Owneruin != nil {
		return *x.Owneruin
	}
	return 0
}

func (x *GroupGeoInfo) GetSettime() uint32 {
	if x != nil && x.Settime != nil {
		return *x.Settime
	}
	return 0
}

func (x *GroupGeoInfo) GetCityid() uint32 {
	if x != nil && x.Cityid != nil {
		return *x.Cityid
	}
	return 0
}

func (x *GroupGeoInfo) GetLongitude() int64 {
	if x != nil && x.Longitude != nil {
		return *x.Longitude
	}
	return 0
}

func (x *GroupGeoInfo) GetLatitude() int64 {
	if x != nil && x.Latitude != nil {
		return *x.Latitude
	}
	return 0
}

func (x *GroupGeoInfo) GetPoiId() uint64 {
	if x != nil && x.PoiId != nil {
		return *x.PoiId
	}
	return 0
}

func (x *GroupGeoInfo) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type GroupHeadPortrait struct {
	PicCnt           *uint32                  `protobuf:"varint,1,opt"`
	Info             []*GroupHeadPortraitInfo `protobuf:"bytes,2,rep"`
	DefaultId        *uint32                  `protobuf:"varint,3,opt"`
	VerifyingPicCnt  *uint32                  `protobuf:"varint,4,opt"`
	VerifyingpicInfo []*GroupHeadPortraitInfo `protobuf:"bytes,5,rep"`
}

func (x *GroupHeadPortrait) GetPicCnt() uint32 {
	if x != nil && x.PicCnt != nil {
		return *x.PicCnt
	}
	return 0
}

func (x *GroupHeadPortrait) GetDefaultId() uint32 {
	if x != nil && x.DefaultId != nil {
		return *x.DefaultId
	}
	return 0
}

func (x *GroupHeadPortrait) GetVerifyingPicCnt() uint32 {
	if x != nil && x.VerifyingPicCnt != nil {
		return *x.VerifyingPicCnt
	}
	return 0
}

func (x *GroupHeadPortrait) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type GroupHeadPortraitInfo struct {
	PicId  *uint32 `protobuf:"varint,1,opt"`
	LeftX  *uint32 `protobuf:"varint,2,opt"`
	LeftY  *uint32 `protobuf:"varint,3,opt"`
	RightX *uint32 `protobuf:"varint,4,opt"`
	RightY *uint32 `protobuf:"varint,5,opt"`
}

func (x *GroupHeadPortraitInfo) GetPicId() uint32 {
	if x != nil && x.PicId != nil {
		return *x.PicId
	}
	return 0
}

func (x *GroupHeadPortraitInfo) GetLeftX() uint32 {
	if x != nil && x.LeftX != nil {
		return *x.LeftX
	}
	return 0
}

func (x *GroupHeadPortraitInfo) GetLeftY() uint32 {
	if x != nil && x.LeftY != nil {
		return *x.LeftY
	}
	return 0
}

func (x *GroupHeadPortraitInfo) GetRightX() uint32 {
	if x != nil && x.RightX != nil {
		return *x.RightX
	}
	return 0
}

func (x *GroupHeadPortraitInfo) GetRightY() uint32 {
	if x != nil && x.RightY != nil {
		return *x.RightY
	}
	return 0
}

func (x *GroupHeadPortraitInfo) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type GroupInfo struct {
	GroupOwner                              *uint64                 `protobuf:"varint,1,opt"`
	GroupCreateTime                         *uint32                 `protobuf:"varint,2,opt"`
	GroupFlag                               *uint32                 `protobuf:"varint,3,opt"`
	GroupFlagExt                            *uint32                 `protobuf:"varint,4,opt"`
	GroupMemberMaxNum                       *uint32                 `protobuf:"varint,5,opt"`
	GroupMemberNum                          *uint32                 `protobuf:"varint,6,opt"`
	GroupOption                             *uint32                 `protobuf:"varint,7,opt"`
	GroupClassExt                           *uint32                 `protobuf:"varint,8,opt"`
	GroupSpecialClass                       *uint32                 `protobuf:"varint,9,opt"`
	GroupLevel                              *uint32                 `protobuf:"varint,10,opt"`
	GroupFace                               *uint32                 `protobuf:"varint,11,opt"`
	GroupDefaultPage                        *uint32                 `protobuf:"varint,12,opt"`
	GroupInfoSeq                            *uint32                 `protobuf:"varint,13,opt"`
	GroupRoamingTime                        *uint32                 `protobuf:"varint,14,opt"`
	GroupName                               []byte                  `protobuf:"bytes,15,opt"`
	GroupMemo                               []byte                  `protobuf:"bytes,16,opt"`
	GroupFingerMemo                         []byte                  `protobuf:"bytes,17,opt"`
	GroupClassText                          []byte                  `protobuf:"bytes,18,opt"`
	GroupAllianceCode                       []uint32                `protobuf:"varint,19,rep"`
	GroupExtraAdmNum                        *uint32                 `protobuf:"varint,20,opt"`
	GroupUin                                *uint64                 `protobuf:"varint,21,opt"`
	GroupCurMsgSeq                          *uint32                 `protobuf:"varint,22,opt"`
	GroupLastMsgTime                        *uint32                 `protobuf:"varint,23,opt"`
	GroupQuestion                           []byte                  `protobuf:"bytes,24,opt"`
	GroupAnswer                             []byte                  `protobuf:"bytes,25,opt"`
	GroupVisitorMaxNum                      *uint32                 `protobuf:"varint,26,opt"`
	GroupVisitorCurNum                      *uint32                 `protobuf:"varint,27,opt"`
	LevelNameSeq                            *uint32                 `protobuf:"varint,28,opt"`
	GroupAdminMaxNum                        *uint32                 `protobuf:"varint,29,opt"`
	GroupAioSkinTimestamp                   *uint32                 `protobuf:"varint,30,opt"`
	GroupBoardSkinTimestamp                 *uint32                 `protobuf:"varint,31,opt"`
	GroupAioSkinUrl                         []byte                  `protobuf:"bytes,32,opt"`
	GroupBoardSkinUrl                       []byte                  `protobuf:"bytes,33,opt"`
	GroupCoverSkinTimestamp                 *uint32                 `protobuf:"varint,34,opt"`
	GroupCoverSkinUrl                       []byte                  `protobuf:"bytes,35,opt"`
	GroupGrade                              *uint32                 `protobuf:"varint,36,opt"`
	ActiveMemberNum                         *uint32                 `protobuf:"varint,37,opt"`
	CertificationType                       *uint32                 `protobuf:"varint,38,opt"`
	CertificationText                       []byte                  `protobuf:"bytes,39,opt"`
	GroupRichFingerMemo                     []byte                  `protobuf:"bytes,40,opt"`
	TagRecord                               []*TagRecord            `protobuf:"bytes,41,rep"`
	GroupGeoInfo                            *GroupGeoInfo           `protobuf:"bytes,42,opt"`
	HeadPortraitSeq                         *uint32                 `protobuf:"varint,43,opt"`
	HeadPortrait                            *GroupHeadPortrait      `protobuf:"bytes,44,opt"`
	ShutupTimestamp                         *uint32                 `protobuf:"varint,45,opt"`
	ShutupTimestampMe                       *uint32                 `protobuf:"varint,46,opt"`
	CreateSourceFlag                        *uint32                 `protobuf:"varint,47,opt"`
	CmduinMsgSeq                            *uint32                 `protobuf:"varint,48,opt"`
	CmduinJoinTime                          *uint32                 `protobuf:"varint,49,opt"`
	CmduinUinFlag                           *uint32                 `protobuf:"varint,50,opt"`
	CmduinFlagEx                            *uint32                 `protobuf:"varint,51,opt"`
	CmduinNewMobileFlag                     *uint32                 `protobuf:"varint,52,opt"`
	CmduinReadMsgSeq                        *uint32                 `protobuf:"varint,53,opt"`
	CmduinLastMsgTime                       *uint32                 `protobuf:"varint,54,opt"`
	GroupTypeFlag                           *uint32                 `protobuf:"varint,55,opt"`
	AppPrivilegeFlag                        *uint32                 `protobuf:"varint,56,opt"`
	StGroupExInfo                           *GroupExInfoOnly        `protobuf:"bytes,57,opt"`
	GroupSecLevel                           *uint32                 `protobuf:"varint,58,opt"`
	GroupSecLevelInfo                       *uint32                 `protobuf:"varint,59,opt"`
	CmduinPrivilege                         *uint32                 `protobuf:"varint,60,opt"`
	PoidInfo                                []byte                  `protobuf:"bytes,61,opt"`
	CmduinFlagEx2                           *uint32                 `protobuf:"varint,62,opt"`
	ConfUin                                 *uint64                 `protobuf:"varint,63,opt"`
	ConfMaxMsgSeq                           *uint32                 `protobuf:"varint,64,opt"`
	ConfToGroupTime                         *uint32                 `protobuf:"varint,65,opt"`
	PasswordRedbagTime                      *uint32                 `protobuf:"varint,66,opt"`
	SubscriptionUin                         *uint64                 `protobuf:"varint,67,opt"`
	MemberListChangeSeq                     *uint32                 `protobuf:"varint,68,opt"`
	MembercardSeq                           *uint32                 `protobuf:"varint,69,opt"`
	RootId                                  *uint64                 `protobuf:"varint,70,opt"`
	ParentId                                *uint64                 `protobuf:"varint,71,opt"`
	TeamSeq                                 *uint32                 `protobuf:"varint,72,opt"`
	HistoryMsgBeginTime                     *uint64                 `protobuf:"varint,73,opt"`
	InviteNoAuthNumLimit                    *uint64                 `protobuf:"varint,74,opt"`
	CmduinHistoryMsgSeq                     *uint32                 `protobuf:"varint,75,opt"`
	CmduinJoinMsgSeq                        *uint32                 `protobuf:"varint,76,opt"`
	GroupFlagext3                           *uint32                 `protobuf:"varint,77,opt"`
	GroupOpenAppid                          *uint32                 `protobuf:"varint,78,opt"`
	IsConfGroup                             *uint32                 `protobuf:"varint,79,opt"`
	IsModifyConfGroupFace                   *uint32                 `protobuf:"varint,80,opt"`
	IsModifyConfGroupName                   *uint32                 `protobuf:"varint,81,opt"`
	NoFingerOpenFlag                        *uint32                 `protobuf:"varint,82,opt"`
	NoCodeFingerOpenFlag                    *uint32                 `protobuf:"varint,83,opt"`
	AutoAgreeJoinGroupUserNumForNormalGroup *uint32                 `protobuf:"varint,84,opt"`
	AutoAgreeJoinGroupUserNumForConfGroup   *uint32                 `protobuf:"varint,85,opt"`
	IsAllowConfGroupMemberNick              *uint32                 `protobuf:"varint,86,opt"`
	IsAllowConfGroupMemberAtAll             *uint32                 `protobuf:"varint,87,opt"`
	IsAllowConfGroupMemberModifyGroupName   *uint32                 `protobuf:"varint,88,opt"`
	LongGroupName                           []byte                  `protobuf:"bytes,89,opt"`
	CmduinJoinRealMsgSeq                    *uint32                 `protobuf:"varint,90,opt"`
	IsGroupFreeze                           *uint32                 `protobuf:"varint,91,opt"`
	MsgLimitFrequency                       *uint32                 `protobuf:"varint,92,opt"`
	JoinGroupAuth                           []byte                  `protobuf:"bytes,93,opt"`
	HlGuildAppid                            *uint32                 `protobuf:"varint,94,opt"`
	HlGuildSubType                          *uint32                 `protobuf:"varint,95,opt"`
	HlGuildOrgid                            *uint32                 `protobuf:"varint,96,opt"`
	IsAllowHlGuildBinary                    *uint32                 `protobuf:"varint,97,opt"`
	CmduinRingtoneId                        *uint32                 `protobuf:"varint,98,opt"`
	GroupFlagext4                           *uint32                 `protobuf:"varint,99,opt"`
	GroupFreezeReason                       *uint32                 `protobuf:"varint,100,opt"`
	IsAllowRecallMsg                        *uint32                 `protobuf:"varint,101,opt"`
	ImportantMsgLatestSeq                   *uint32                 `protobuf:"varint,102,opt"`
	GroupSchoolInfo                         []byte                  `protobuf:"bytes,103,opt"`
	AppealDeadline                          *uint32                 `protobuf:"varint,104,opt"`
	StGroupCardPrefix                       *GroupCardPrefix        `protobuf:"bytes,105,opt"`
	AllianceId                              *uint64                 `protobuf:"varint,106,opt"`
	CmduinFlagEx3Grocery                    *uint32                 `protobuf:"varint,107,opt"`
	GroupInfoExtSeq                         *uint32                 `protobuf:"varint,108,opt"`
	StGroupInfoExt                          *oidb0xef0.GroupInfoExt `protobuf:"bytes,109,opt"`
	CmduinGroupRemarkName                   []byte                  `protobuf:"bytes,110,opt"`
}

func (x *GroupInfo) GetGroupOwner() uint64 {
	if x != nil && x.GroupOwner != nil {
		return *x.GroupOwner
	}
	return 0
}

func (x *GroupInfo) GetGroupCreateTime() uint32 {
	if x != nil && x.GroupCreateTime != nil {
		return *x.GroupCreateTime
	}
	return 0
}

func (x *GroupInfo) GetGroupFlag() uint32 {
	if x != nil && x.GroupFlag != nil {
		return *x.GroupFlag
	}
	return 0
}

func (x *GroupInfo) GetGroupFlagExt() uint32 {
	if x != nil && x.GroupFlagExt != nil {
		return *x.GroupFlagExt
	}
	return 0
}

func (x *GroupInfo) GetGroupMemberMaxNum() uint32 {
	if x != nil && x.GroupMemberMaxNum != nil {
		return *x.GroupMemberMaxNum
	}
	return 0
}

func (x *GroupInfo) GetGroupMemberNum() uint32 {
	if x != nil && x.GroupMemberNum != nil {
		return *x.GroupMemberNum
	}
	return 0
}

func (x *GroupInfo) GetGroupOption() uint32 {
	if x != nil && x.GroupOption != nil {
		return *x.GroupOption
	}
	return 0
}

func (x *GroupInfo) GetGroupClassExt() uint32 {
	if x != nil && x.GroupClassExt != nil {
		return *x.GroupClassExt
	}
	return 0
}

func (x *GroupInfo) GetGroupSpecialClass() uint32 {
	if x != nil && x.GroupSpecialClass != nil {
		return *x.GroupSpecialClass
	}
	return 0
}

func (x *GroupInfo) GetGroupLevel() uint32 {
	if x != nil && x.GroupLevel != nil {
		return *x.GroupLevel
	}
	return 0
}

func (x *GroupInfo) GetGroupFace() uint32 {
	if x != nil && x.GroupFace != nil {
		return *x.GroupFace
	}
	return 0
}

func (x *GroupInfo) GetGroupDefaultPage() uint32 {
	if x != nil && x.GroupDefaultPage != nil {
		return *x.GroupDefaultPage
	}
	return 0
}

func (x *GroupInfo) GetGroupInfoSeq() uint32 {
	if x != nil && x.GroupInfoSeq != nil {
		return *x.GroupInfoSeq
	}
	return 0
}

func (x *GroupInfo) GetGroupRoamingTime() uint32 {
	if x != nil && x.GroupRoamingTime != nil {
		return *x.GroupRoamingTime
	}
	return 0
}

func (x *GroupInfo) GetGroupExtraAdmNum() uint32 {
	if x != nil && x.GroupExtraAdmNum != nil {
		return *x.GroupExtraAdmNum
	}
	return 0
}

func (x *GroupInfo) GetGroupUin() uint64 {
	if x != nil && x.GroupUin != nil {
		return *x.GroupUin
	}
	return 0
}

func (x *GroupInfo) GetGroupCurMsgSeq() uint32 {
	if x != nil && x.GroupCurMsgSeq != nil {
		return *x.GroupCurMsgSeq
	}
	return 0
}

func (x *GroupInfo) GetGroupLastMsgTime() uint32 {
	if x != nil && x.GroupLastMsgTime != nil {
		return *x.GroupLastMsgTime
	}
	return 0
}

func (x *GroupInfo) GetGroupVisitorMaxNum() uint32 {
	if x != nil && x.GroupVisitorMaxNum != nil {
		return *x.GroupVisitorMaxNum
	}
	return 0
}

func (x *GroupInfo) GetGroupVisitorCurNum() uint32 {
	if x != nil && x.GroupVisitorCurNum != nil {
		return *x.GroupVisitorCurNum
	}
	return 0
}

func (x *GroupInfo) GetLevelNameSeq() uint32 {
	if x != nil && x.LevelNameSeq != nil {
		return *x.LevelNameSeq
	}
	return 0
}

func (x *GroupInfo) GetGroupAdminMaxNum() uint32 {
	if x != nil && x.GroupAdminMaxNum != nil {
		return *x.GroupAdminMaxNum
	}
	return 0
}

func (x *GroupInfo) GetGroupAioSkinTimestamp() uint32 {
	if x != nil && x.GroupAioSkinTimestamp != nil {
		return *x.GroupAioSkinTimestamp
	}
	return 0
}

func (x *GroupInfo) GetGroupBoardSkinTimestamp() uint32 {
	if x != nil && x.GroupBoardSkinTimestamp != nil {
		return *x.GroupBoardSkinTimestamp
	}
	return 0
}

func (x *GroupInfo) GetGroupCoverSkinTimestamp() uint32 {
	if x != nil && x.GroupCoverSkinTimestamp != nil {
		return *x.GroupCoverSkinTimestamp
	}
	return 0
}

func (x *GroupInfo) GetGroupGrade() uint32 {
	if x != nil && x.GroupGrade != nil {
		return *x.GroupGrade
	}
	return 0
}

func (x *GroupInfo) GetActiveMemberNum() uint32 {
	if x != nil && x.ActiveMemberNum != nil {
		return *x.ActiveMemberNum
	}
	return 0
}

func (x *GroupInfo) GetCertificationType() uint32 {
	if x != nil && x.CertificationType != nil {
		return *x.CertificationType
	}
	return 0
}

func (x *GroupInfo) GetHeadPortraitSeq() uint32 {
	if x != nil && x.HeadPortraitSeq != nil {
		return *x.HeadPortraitSeq
	}
	return 0
}

func (x *GroupInfo) GetShutupTimestamp() uint32 {
	if x != nil && x.ShutupTimestamp != nil {
		return *x.ShutupTimestamp
	}
	return 0
}

func (x *GroupInfo) GetShutupTimestampMe() uint32 {
	if x != nil && x.ShutupTimestampMe != nil {
		return *x.ShutupTimestampMe
	}
	return 0
}

func (x *GroupInfo) GetCreateSourceFlag() uint32 {
	if x != nil && x.CreateSourceFlag != nil {
		return *x.CreateSourceFlag
	}
	return 0
}

func (x *GroupInfo) GetCmduinMsgSeq() uint32 {
	if x != nil && x.CmduinMsgSeq != nil {
		return *x.CmduinMsgSeq
	}
	return 0
}

func (x *GroupInfo) GetCmduinJoinTime() uint32 {
	if x != nil && x.CmduinJoinTime != nil {
		return *x.CmduinJoinTime
	}
	return 0
}

func (x *GroupInfo) GetCmduinUinFlag() uint32 {
	if x != nil && x.CmduinUinFlag != nil {
		return *x.CmduinUinFlag
	}
	return 0
}

func (x *GroupInfo) GetCmduinFlagEx() uint32 {
	if x != nil && x.CmduinFlagEx != nil {
		return *x.CmduinFlagEx
	}
	return 0
}

func (x *GroupInfo) GetCmduinNewMobileFlag() uint32 {
	if x != nil && x.CmduinNewMobileFlag != nil {
		return *x.CmduinNewMobileFlag
	}
	return 0
}

func (x *GroupInfo) GetCmduinReadMsgSeq() uint32 {
	if x != nil && x.CmduinReadMsgSeq != nil {
		return *x.CmduinReadMsgSeq
	}
	return 0
}

func (x *GroupInfo) GetCmduinLastMsgTime() uint32 {
	if x != nil && x.CmduinLastMsgTime != nil {
		return *x.CmduinLastMsgTime
	}
	return 0
}

func (x *GroupInfo) GetGroupTypeFlag() uint32 {
	if x != nil && x.GroupTypeFlag != nil {
		return *x.GroupTypeFlag
	}
	return 0
}

func (x *GroupInfo) GetAppPrivilegeFlag() uint32 {
	if x != nil && x.AppPrivilegeFlag != nil {
		return *x.AppPrivilegeFlag
	}
	return 0
}

func (x *GroupInfo) GetGroupSecLevel() uint32 {
	if x != nil && x.GroupSecLevel != nil {
		return *x.GroupSecLevel
	}
	return 0
}

func (x *GroupInfo) GetGroupSecLevelInfo() uint32 {
	if x != nil && x.GroupSecLevelInfo != nil {
		return *x.GroupSecLevelInfo
	}
	return 0
}

func (x *GroupInfo) GetCmduinPrivilege() uint32 {
	if x != nil && x.CmduinPrivilege != nil {
		return *x.CmduinPrivilege
	}
	return 0
}

func (x *GroupInfo) GetCmduinFlagEx2() uint32 {
	if x != nil && x.CmduinFlagEx2 != nil {
		return *x.CmduinFlagEx2
	}
	return 0
}

func (x *GroupInfo) GetConfUin() uint64 {
	if x != nil && x.ConfUin != nil {
		return *x.ConfUin
	}
	return 0
}

func (x *GroupInfo) GetConfMaxMsgSeq() uint32 {
	if x != nil && x.ConfMaxMsgSeq != nil {
		return *x.ConfMaxMsgSeq
	}
	return 0
}

func (x *GroupInfo) GetConfToGroupTime() uint32 {
	if x != nil && x.ConfToGroupTime != nil {
		return *x.ConfToGroupTime
	}
	return 0
}

func (x *GroupInfo) GetPasswordRedbagTime() uint32 {
	if x != nil && x.PasswordRedbagTime != nil {
		return *x.PasswordRedbagTime
	}
	return 0
}

func (x *GroupInfo) GetSubscriptionUin() uint64 {
	if x != nil && x.SubscriptionUin != nil {
		return *x.SubscriptionUin
	}
	return 0
}

func (x *GroupInfo) GetMemberListChangeSeq() uint32 {
	if x != nil && x.MemberListChangeSeq != nil {
		return *x.MemberListChangeSeq
	}
	return 0
}

func (x *GroupInfo) GetMembercardSeq() uint32 {
	if x != nil && x.MembercardSeq != nil {
		return *x.MembercardSeq
	}
	return 0
}

func (x *GroupInfo) GetRootId() uint64 {
	if x != nil && x.RootId != nil {
		return *x.RootId
	}
	return 0
}

func (x *GroupInfo) GetParentId() uint64 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

func (x *GroupInfo) GetTeamSeq() uint32 {
	if x != nil && x.TeamSeq != nil {
		return *x.TeamSeq
	}
	return 0
}

func (x *GroupInfo) GetHistoryMsgBeginTime() uint64 {
	if x != nil && x.HistoryMsgBeginTime != nil {
		return *x.HistoryMsgBeginTime
	}
	return 0
}

func (x *GroupInfo) GetInviteNoAuthNumLimit() uint64 {
	if x != nil && x.InviteNoAuthNumLimit != nil {
		return *x.InviteNoAuthNumLimit
	}
	return 0
}

func (x *GroupInfo) GetCmduinHistoryMsgSeq() uint32 {
	if x != nil && x.CmduinHistoryMsgSeq != nil {
		return *x.CmduinHistoryMsgSeq
	}
	return 0
}

func (x *GroupInfo) GetCmduinJoinMsgSeq() uint32 {
	if x != nil && x.CmduinJoinMsgSeq != nil {
		return *x.CmduinJoinMsgSeq
	}
	return 0
}

func (x *GroupInfo) GetGroupFlagext3() uint32 {
	if x != nil && x.GroupFlagext3 != nil {
		return *x.GroupFlagext3
	}
	return 0
}

func (x *GroupInfo) GetGroupOpenAppid() uint32 {
	if x != nil && x.GroupOpenAppid != nil {
		return *x.GroupOpenAppid
	}
	return 0
}

func (x *GroupInfo) GetIsConfGroup() uint32 {
	if x != nil && x.IsConfGroup != nil {
		return *x.IsConfGroup
	}
	return 0
}

func (x *GroupInfo) GetIsModifyConfGroupFace() uint32 {
	if x != nil && x.IsModifyConfGroupFace != nil {
		return *x.IsModifyConfGroupFace
	}
	return 0
}

func (x *GroupInfo) GetIsModifyConfGroupName() uint32 {
	if x != nil && x.IsModifyConfGroupName != nil {
		return *x.IsModifyConfGroupName
	}
	return 0
}

func (x *GroupInfo) GetNoFingerOpenFlag() uint32 {
	if x != nil && x.NoFingerOpenFlag != nil {
		return *x.NoFingerOpenFlag
	}
	return 0
}

func (x *GroupInfo) GetNoCodeFingerOpenFlag() uint32 {
	if x != nil && x.NoCodeFingerOpenFlag != nil {
		return *x.NoCodeFingerOpenFlag
	}
	return 0
}

func (x *GroupInfo) GetAutoAgreeJoinGroupUserNumForNormalGroup() uint32 {
	if x != nil && x.AutoAgreeJoinGroupUserNumForNormalGroup != nil {
		return *x.AutoAgreeJoinGroupUserNumForNormalGroup
	}
	return 0
}

func (x *GroupInfo) GetAutoAgreeJoinGroupUserNumForConfGroup() uint32 {
	if x != nil && x.AutoAgreeJoinGroupUserNumForConfGroup != nil {
		return *x.AutoAgreeJoinGroupUserNumForConfGroup
	}
	return 0
}

func (x *GroupInfo) GetIsAllowConfGroupMemberNick() uint32 {
	if x != nil && x.IsAllowConfGroupMemberNick != nil {
		return *x.IsAllowConfGroupMemberNick
	}
	return 0
}

func (x *GroupInfo) GetIsAllowConfGroupMemberAtAll() uint32 {
	if x != nil && x.IsAllowConfGroupMemberAtAll != nil {
		return *x.IsAllowConfGroupMemberAtAll
	}
	return 0
}

func (x *GroupInfo) GetIsAllowConfGroupMemberModifyGroupName() uint32 {
	if x != nil && x.IsAllowConfGroupMemberModifyGroupName != nil {
		return *x.IsAllowConfGroupMemberModifyGroupName
	}
	return 0
}

func (x *GroupInfo) GetCmduinJoinRealMsgSeq() uint32 {
	if x != nil && x.CmduinJoinRealMsgSeq != nil {
		return *x.CmduinJoinRealMsgSeq
	}
	return 0
}

func (x *GroupInfo) GetIsGroupFreeze() uint32 {
	if x != nil && x.IsGroupFreeze != nil {
		return *x.IsGroupFreeze
	}
	return 0
}

func (x *GroupInfo) GetMsgLimitFrequency() uint32 {
	if x != nil && x.MsgLimitFrequency != nil {
		return *x.MsgLimitFrequency
	}
	return 0
}

func (x *GroupInfo) GetHlGuildAppid() uint32 {
	if x != nil && x.HlGuildAppid != nil {
		return *x.HlGuildAppid
	}
	return 0
}

func (x *GroupInfo) GetHlGuildSubType() uint32 {
	if x != nil && x.HlGuildSubType != nil {
		return *x.HlGuildSubType
	}
	return 0
}

func (x *GroupInfo) GetHlGuildOrgid() uint32 {
	if x != nil && x.HlGuildOrgid != nil {
		return *x.HlGuildOrgid
	}
	return 0
}

func (x *GroupInfo) GetIsAllowHlGuildBinary() uint32 {
	if x != nil && x.IsAllowHlGuildBinary != nil {
		return *x.IsAllowHlGuildBinary
	}
	return 0
}

func (x *GroupInfo) GetCmduinRingtoneId() uint32 {
	if x != nil && x.CmduinRingtoneId != nil {
		return *x.CmduinRingtoneId
	}
	return 0
}

func (x *GroupInfo) GetGroupFlagext4() uint32 {
	if x != nil && x.GroupFlagext4 != nil {
		return *x.GroupFlagext4
	}
	return 0
}

func (x *GroupInfo) GetGroupFreezeReason() uint32 {
	if x != nil && x.GroupFreezeReason != nil {
		return *x.GroupFreezeReason
	}
	return 0
}

func (x *GroupInfo) GetIsAllowRecallMsg() uint32 {
	if x != nil && x.IsAllowRecallMsg != nil {
		return *x.IsAllowRecallMsg
	}
	return 0
}

func (x *GroupInfo) GetImportantMsgLatestSeq() uint32 {
	if x != nil && x.ImportantMsgLatestSeq != nil {
		return *x.ImportantMsgLatestSeq
	}
	return 0
}

func (x *GroupInfo) GetAppealDeadline() uint32 {
	if x != nil && x.AppealDeadline != nil {
		return *x.AppealDeadline
	}
	return 0
}

func (x *GroupInfo) GetAllianceId() uint64 {
	if x != nil && x.AllianceId != nil {
		return *x.AllianceId
	}
	return 0
}

func (x *GroupInfo) GetCmduinFlagEx3Grocery() uint32 {
	if x != nil && x.CmduinFlagEx3Grocery != nil {
		return *x.CmduinFlagEx3Grocery
	}
	return 0
}

func (x *GroupInfo) GetGroupInfoExtSeq() uint32 {
	if x != nil && x.GroupInfoExtSeq != nil {
		return *x.GroupInfoExtSeq
	}
	return 0
}

func (x *GroupInfo) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type ReqBody struct {
	Appid           *uint32         `protobuf:"varint,1,opt"`
	Stzreqgroupinfo []*ReqGroupInfo `protobuf:"bytes,2,rep"`
	PcClientVersion *uint32         `protobuf:"varint,3,opt"`
}

func (x *ReqBody) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *ReqBody) GetPcClientVersion() uint32 {
	if x != nil && x.PcClientVersion != nil {
		return *x.PcClientVersion
	}
	return 0
}

func (x *ReqBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type ReqGroupInfo struct {
	GroupCode            *uint64    `protobuf:"varint,1,opt"`
	Stgroupinfo          *GroupInfo `protobuf:"bytes,2,opt"`
	LastGetGroupNameTime *uint32    `protobuf:"varint,3,opt"`
}

func (x *ReqGroupInfo) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *ReqGroupInfo) GetLastGetGroupNameTime() uint32 {
	if x != nil && x.LastGetGroupNameTime != nil {
		return *x.LastGetGroupNameTime
	}
	return 0
}

func (x *ReqGroupInfo) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type RspBody struct {
	Stzrspgroupinfo []*RspGroupInfo `protobuf:"bytes,1,rep"`
	Errorinfo       []byte          `protobuf:"bytes,2,opt"`
}

func (x *RspBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type RspGroupInfo struct {
	GroupCode   *uint64    `protobuf:"varint,1,opt"`
	Result      *uint32    `protobuf:"varint,2,opt"`
	Stgroupinfo *GroupInfo `protobuf:"bytes,3,opt"`
}

func (x *RspGroupInfo) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *RspGroupInfo) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *RspGroupInfo) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type TagRecord struct {
	FromUin   *uint64 `protobuf:"varint,1,opt"`
	GroupCode *uint64 `protobuf:"varint,2,opt"`
	TagId     []byte  `protobuf:"bytes,3,opt"`
	SetTime   *uint64 `protobuf:"varint,4,opt"`
	GoodNum   *uint32 `protobuf:"varint,5,opt"`
	BadNum    *uint32 `protobuf:"varint,6,opt"`
	TagLen    *uint32 `protobuf:"varint,7,opt"`
	TagValue  []byte  `protobuf:"bytes,8,opt"`
}

func (x *TagRecord) GetFromUin() uint64 {
	if x != nil && x.FromUin != nil {
		return *x.FromUin
	}
	return 0
}

func (x *TagRecord) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *TagRecord) GetSetTime() uint64 {
	if x != nil && x.SetTime != nil {
		return *x.SetTime
	}
	return 0
}

func (x *TagRecord) GetGoodNum() uint32 {
	if x != nil && x.GoodNum != nil {
		return *x.GoodNum
	}
	return 0
}

func (x *TagRecord) GetBadNum() uint32 {
	if x != nil && x.BadNum != nil {
		return *x.BadNum
	}
	return 0
}

func (x *TagRecord) GetTagLen() uint32 {
	if x != nil && x.TagLen != nil {
		return *x.TagLen
	}
	return 0
}

func (x *TagRecord) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}
