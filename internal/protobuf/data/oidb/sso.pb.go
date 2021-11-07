// Code generated by yaprotoc. DO NOT EDIT.
// source: oidb/sso.proto

package oidb

import (
	"github.com/pkg/errors"
	"github.com/segmentio/encoding/proto"
)

type OIDBSSOPkg struct {
	Command       *uint32 `protobuf:"varint,1,opt"`
	ServiceType   *uint32 `protobuf:"varint,2,opt"`
	Result        *uint32 `protobuf:"varint,3,opt"`
	Bodybuffer    []byte  `protobuf:"bytes,4,opt"`
	ErrorMsg      *string `protobuf:"bytes,5,opt"`
	ClientVersion *string `protobuf:"bytes,6,opt"`
}

func (x *OIDBSSOPkg) GetCommand() uint32 {
	if x != nil && x.Command != nil {
		return *x.Command
	}
	return 0
}

func (x *OIDBSSOPkg) GetServiceType() uint32 {
	if x != nil && x.ServiceType != nil {
		return *x.ServiceType
	}
	return 0
}

func (x *OIDBSSOPkg) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *OIDBSSOPkg) GetErrorMsg() string {
	if x != nil && x.ErrorMsg != nil {
		return *x.ErrorMsg
	}
	return ""
}

func (x *OIDBSSOPkg) GetClientVersion() string {
	if x != nil && x.ClientVersion != nil {
		return *x.ClientVersion
	}
	return ""
}

func (x *OIDBSSOPkg) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}
