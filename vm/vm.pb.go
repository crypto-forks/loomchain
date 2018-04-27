// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/loom/vm/vm.proto

/*
Package vm is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/loom/vm/vm.proto

It has these top-level messages:
	MessageTx
	DeployTx
	CallTx
	DeployResponse
*/
package vm

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type VMType int32

const (
	VMType_PLUGIN VMType = 0
	VMType_EVM    VMType = 1
)

var VMType_name = map[int32]string{
	0: "PLUGIN",
	1: "EVM",
}
var VMType_value = map[string]int32{
	"PLUGIN": 0,
	"EVM":    1,
}

func (x VMType) String() string {
	return proto.EnumName(VMType_name, int32(x))
}
func (VMType) EnumDescriptor() ([]byte, []int) { return fileDescriptorVm, []int{0} }

type MessageTx struct {
	To   *types.Address `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
	From *types.Address `protobuf:"bytes,2,opt,name=from" json:"from,omitempty"`
	Data []byte         `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MessageTx) Reset()                    { *m = MessageTx{} }
func (m *MessageTx) String() string            { return proto.CompactTextString(m) }
func (*MessageTx) ProtoMessage()               {}
func (*MessageTx) Descriptor() ([]byte, []int) { return fileDescriptorVm, []int{0} }

func (m *MessageTx) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *MessageTx) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *MessageTx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeployTx struct {
	VmType VMType `protobuf:"varint,1,opt,name=vm_type,json=vmType,proto3,enum=VMType" json:"vm_type,omitempty"`
	Code   []byte `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (m *DeployTx) Reset()                    { *m = DeployTx{} }
func (m *DeployTx) String() string            { return proto.CompactTextString(m) }
func (*DeployTx) ProtoMessage()               {}
func (*DeployTx) Descriptor() ([]byte, []int) { return fileDescriptorVm, []int{1} }

func (m *DeployTx) GetVmType() VMType {
	if m != nil {
		return m.VmType
	}
	return VMType_PLUGIN
}

func (m *DeployTx) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

type CallTx struct {
	VmType VMType `protobuf:"varint,1,opt,name=vm_type,json=vmType,proto3,enum=VMType" json:"vm_type,omitempty"`
	Input  []byte `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
}

func (m *CallTx) Reset()                    { *m = CallTx{} }
func (m *CallTx) String() string            { return proto.CompactTextString(m) }
func (*CallTx) ProtoMessage()               {}
func (*CallTx) Descriptor() ([]byte, []int) { return fileDescriptorVm, []int{2} }

func (m *CallTx) GetVmType() VMType {
	if m != nil {
		return m.VmType
	}
	return VMType_PLUGIN
}

func (m *CallTx) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

type DeployResponse struct {
	Contract *types.Address `protobuf:"bytes,1,opt,name=contract" json:"contract,omitempty"`
	Output   []byte         `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (m *DeployResponse) Reset()                    { *m = DeployResponse{} }
func (m *DeployResponse) String() string            { return proto.CompactTextString(m) }
func (*DeployResponse) ProtoMessage()               {}
func (*DeployResponse) Descriptor() ([]byte, []int) { return fileDescriptorVm, []int{3} }

func (m *DeployResponse) GetContract() *types.Address {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *DeployResponse) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageTx)(nil), "MessageTx")
	proto.RegisterType((*DeployTx)(nil), "DeployTx")
	proto.RegisterType((*CallTx)(nil), "CallTx")
	proto.RegisterType((*DeployResponse)(nil), "DeployResponse")
	proto.RegisterEnum("VMType", VMType_name, VMType_value)
}

func init() { proto.RegisterFile("github.com/loomnetwork/loom/vm/vm.proto", fileDescriptorVm) }

var fileDescriptorVm = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x71, 0x4b, 0xf3, 0x30,
	0x10, 0xc6, 0xdf, 0x6e, 0x7b, 0xd3, 0x7a, 0x8e, 0x31, 0x82, 0x48, 0x11, 0x85, 0x52, 0x04, 0x87,
	0x60, 0x27, 0xf3, 0x0b, 0x4c, 0x54, 0x44, 0xb0, 0x43, 0x4a, 0x9d, 0x7f, 0x4a, 0xd7, 0xc6, 0x3a,
	0x6c, 0x7b, 0xa1, 0xb9, 0xd6, 0xf5, 0xdb, 0x4b, 0xd3, 0xa1, 0xa0, 0x08, 0x42, 0x48, 0x9e, 0xcb,
	0xdd, 0xfd, 0x9e, 0xe4, 0xe0, 0x24, 0x5d, 0xd3, 0x6b, 0xb5, 0xf2, 0x62, 0xcc, 0xa7, 0x19, 0x62,
	0x5e, 0x08, 0x7a, 0xc7, 0xf2, 0x4d, 0xeb, 0x69, 0xdd, 0x2e, 0x4f, 0x96, 0x48, 0x78, 0x70, 0xfe,
	0x4b, 0x61, 0x8a, 0x67, 0xba, 0x96, 0x1a, 0x29, 0x54, 0xb7, 0x77, 0x1d, 0xee, 0x13, 0xec, 0xf8,
	0x42, 0xa9, 0x28, 0x15, 0xe1, 0x86, 0xdb, 0xd0, 0x23, 0xb4, 0x0d, 0xc7, 0x98, 0xec, 0xce, 0x2c,
	0xef, 0x32, 0x49, 0x4a, 0xa1, 0x54, 0xd0, 0x23, 0xe4, 0x87, 0x30, 0x78, 0x29, 0x31, 0xb7, 0x7b,
	0xdf, 0x72, 0xfa, 0x96, 0x73, 0x18, 0x24, 0x11, 0x45, 0x76, 0xdf, 0x31, 0x26, 0xc3, 0x40, 0x6b,
	0x77, 0x0e, 0xd6, 0xb5, 0x90, 0x19, 0x36, 0xe1, 0x86, 0x3b, 0x60, 0xd6, 0xf9, 0x73, 0x6b, 0xab,
	0xe1, 0xa3, 0x99, 0xe9, 0x2d, 0xfd, 0xb0, 0x91, 0x22, 0x60, 0x75, 0xde, 0x9e, 0x2d, 0x21, 0xc6,
	0x44, 0x68, 0xfe, 0x30, 0xd0, 0xda, 0x9d, 0x03, 0xbb, 0x8a, 0xb2, 0xec, 0x4f, 0xfd, 0x7b, 0xf0,
	0x7f, 0x5d, 0xc8, 0x8a, 0xb6, 0x80, 0x2e, 0x70, 0x17, 0x30, 0xea, 0xde, 0x10, 0x08, 0x25, 0xb1,
	0x50, 0x82, 0x1f, 0x83, 0x15, 0x63, 0x41, 0x65, 0x14, 0xd3, 0x8f, 0x7f, 0x7e, 0x66, 0xf8, 0x3e,
	0x30, 0xac, 0xe8, 0x0b, 0xb7, 0x8d, 0x4e, 0x8f, 0x80, 0x75, 0xbe, 0x1c, 0x80, 0x3d, 0xdc, 0x3f,
	0xde, 0xde, 0x2d, 0xc6, 0xff, 0xb8, 0x09, 0xfd, 0x9b, 0xa5, 0x3f, 0x36, 0x56, 0x4c, 0x8f, 0xf4,
	0xe2, 0x23, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x02, 0xa3, 0xce, 0xaf, 0x01, 0x00, 0x00,
}
