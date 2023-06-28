// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: router/v1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the router genesis state
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// key - information about forwarded packet: src_channel
	// (parsedReceiver.Channel), src_port (parsedReceiver.Port), sequence value -
	// information about original packet for refunding if necessary: retries,
	// srcPacketSender, srcPacket.DestinationChannel, srcPacket.DestinationPort
	InFlightPackets map[string]InFlightPacket `protobuf:"bytes,2,rep,name=in_flight_packets,json=inFlightPackets,proto3" json:"in_flight_packets" yaml:"in_flight_packets" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4940b763c55c4e0b, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetInFlightPackets() map[string]InFlightPacket {
	if m != nil {
		return m.InFlightPackets
	}
	return nil
}

// Params defines the set of IBC router parameters.
type Params struct {
	FeePercentage github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=fee_percentage,json=feePercentage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fee_percentage" yaml:"fee_percentage"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_4940b763c55c4e0b, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// InFlightPacket contains information about original packet for
// writing the acknowledgement and refunding if necessary.
type InFlightPacket struct {
	OriginalSenderAddress  string `protobuf:"bytes,1,opt,name=original_sender_address,json=originalSenderAddress,proto3" json:"original_sender_address,omitempty"`
	RefundChannelId        string `protobuf:"bytes,2,opt,name=refund_channel_id,json=refundChannelId,proto3" json:"refund_channel_id,omitempty"`
	RefundPortId           string `protobuf:"bytes,3,opt,name=refund_port_id,json=refundPortId,proto3" json:"refund_port_id,omitempty"`
	PacketSrcChannelId     string `protobuf:"bytes,4,opt,name=packet_src_channel_id,json=packetSrcChannelId,proto3" json:"packet_src_channel_id,omitempty"`
	PacketSrcPortId        string `protobuf:"bytes,5,opt,name=packet_src_port_id,json=packetSrcPortId,proto3" json:"packet_src_port_id,omitempty"`
	PacketTimeoutTimestamp uint64 `protobuf:"varint,6,opt,name=packet_timeout_timestamp,json=packetTimeoutTimestamp,proto3" json:"packet_timeout_timestamp,omitempty"`
	PacketTimeoutHeight    string `protobuf:"bytes,7,opt,name=packet_timeout_height,json=packetTimeoutHeight,proto3" json:"packet_timeout_height,omitempty"`
	PacketData             []byte `protobuf:"bytes,8,opt,name=packet_data,json=packetData,proto3" json:"packet_data,omitempty"`
	RefundSequence         uint64 `protobuf:"varint,9,opt,name=refund_sequence,json=refundSequence,proto3" json:"refund_sequence,omitempty"`
	RetriesRemaining       int32  `protobuf:"varint,10,opt,name=retries_remaining,json=retriesRemaining,proto3" json:"retries_remaining,omitempty"`
	Timeout                uint64 `protobuf:"varint,11,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Nonrefundable          bool   `protobuf:"varint,12,opt,name=nonrefundable,proto3" json:"nonrefundable,omitempty"`
}

func (m *InFlightPacket) Reset()         { *m = InFlightPacket{} }
func (m *InFlightPacket) String() string { return proto.CompactTextString(m) }
func (*InFlightPacket) ProtoMessage()    {}
func (*InFlightPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_4940b763c55c4e0b, []int{2}
}
func (m *InFlightPacket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InFlightPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InFlightPacket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InFlightPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InFlightPacket.Merge(m, src)
}
func (m *InFlightPacket) XXX_Size() int {
	return m.Size()
}
func (m *InFlightPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_InFlightPacket.DiscardUnknown(m)
}

var xxx_messageInfo_InFlightPacket proto.InternalMessageInfo

func (m *InFlightPacket) GetOriginalSenderAddress() string {
	if m != nil {
		return m.OriginalSenderAddress
	}
	return ""
}

func (m *InFlightPacket) GetRefundChannelId() string {
	if m != nil {
		return m.RefundChannelId
	}
	return ""
}

func (m *InFlightPacket) GetRefundPortId() string {
	if m != nil {
		return m.RefundPortId
	}
	return ""
}

func (m *InFlightPacket) GetPacketSrcChannelId() string {
	if m != nil {
		return m.PacketSrcChannelId
	}
	return ""
}

func (m *InFlightPacket) GetPacketSrcPortId() string {
	if m != nil {
		return m.PacketSrcPortId
	}
	return ""
}

func (m *InFlightPacket) GetPacketTimeoutTimestamp() uint64 {
	if m != nil {
		return m.PacketTimeoutTimestamp
	}
	return 0
}

func (m *InFlightPacket) GetPacketTimeoutHeight() string {
	if m != nil {
		return m.PacketTimeoutHeight
	}
	return ""
}

func (m *InFlightPacket) GetPacketData() []byte {
	if m != nil {
		return m.PacketData
	}
	return nil
}

func (m *InFlightPacket) GetRefundSequence() uint64 {
	if m != nil {
		return m.RefundSequence
	}
	return 0
}

func (m *InFlightPacket) GetRetriesRemaining() int32 {
	if m != nil {
		return m.RetriesRemaining
	}
	return 0
}

func (m *InFlightPacket) GetTimeout() uint64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *InFlightPacket) GetNonrefundable() bool {
	if m != nil {
		return m.Nonrefundable
	}
	return false
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "router.v1.GenesisState")
	proto.RegisterMapType((map[string]InFlightPacket)(nil), "router.v1.GenesisState.InFlightPacketsEntry")
	proto.RegisterType((*Params)(nil), "router.v1.Params")
	proto.RegisterType((*InFlightPacket)(nil), "router.v1.InFlightPacket")
}

func init() { proto.RegisterFile("router/v1/genesis.proto", fileDescriptor_4940b763c55c4e0b) }

var fileDescriptor_4940b763c55c4e0b = []byte{
	// 643 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xcd, 0x4e, 0x1b, 0x3b,
	0x14, 0xce, 0x84, 0x10, 0x88, 0x13, 0xfe, 0x7c, 0xe1, 0x32, 0x97, 0x45, 0x12, 0x45, 0xe8, 0xde,
	0xe8, 0xd2, 0x64, 0x04, 0x55, 0x2b, 0xc4, 0xae, 0x29, 0x2d, 0x65, 0xd5, 0x68, 0xc2, 0xaa, 0x52,
	0x35, 0x72, 0x66, 0x4e, 0x26, 0x16, 0x33, 0xf6, 0xd4, 0x76, 0x42, 0xd3, 0xa7, 0xe8, 0xba, 0x9b,
	0xbe, 0x0e, 0x4b, 0x96, 0x55, 0x17, 0x51, 0x05, 0x6f, 0xc0, 0x13, 0x54, 0x63, 0x4f, 0x20, 0x11,
	0x5d, 0x8d, 0x7d, 0xbe, 0x1f, 0x7f, 0xe7, 0x8c, 0x65, 0xb4, 0x2b, 0xf8, 0x48, 0x81, 0x70, 0xc6,
	0x87, 0x4e, 0x08, 0x0c, 0x24, 0x95, 0xed, 0x44, 0x70, 0xc5, 0x71, 0xc9, 0x00, 0xed, 0xf1, 0xe1,
	0xde, 0x76, 0xc8, 0x43, 0xae, 0xab, 0x4e, 0xba, 0x32, 0x84, 0xc6, 0xb7, 0x3c, 0xaa, 0x9c, 0x19,
	0x49, 0x4f, 0x11, 0x05, 0xd8, 0x41, 0xc5, 0x84, 0x08, 0x12, 0x4b, 0xdb, 0xaa, 0x5b, 0xcd, 0xf2,
	0xd1, 0x56, 0xfb, 0xc1, 0xa2, 0xdd, 0xd5, 0x40, 0xa7, 0x70, 0x3d, 0xad, 0xe5, 0xdc, 0x8c, 0x86,
	0xbf, 0xa0, 0x2d, 0xca, 0xbc, 0x41, 0x44, 0xc3, 0xa1, 0xf2, 0x12, 0xe2, 0x5f, 0x82, 0x92, 0x76,
	0xbe, 0xbe, 0xd4, 0x2c, 0x1f, 0x3d, 0x9b, 0xd3, 0xce, 0x1f, 0xd2, 0x3e, 0x67, 0x6f, 0x35, 0xbf,
	0x6b, 0xe8, 0x6f, 0x98, 0x12, 0x93, 0x4e, 0x3d, 0xb5, 0xbd, 0x9f, 0xd6, 0xec, 0x09, 0x89, 0xa3,
	0x93, 0xc6, 0x13, 0xd3, 0x86, 0xbb, 0x41, 0x17, 0x75, 0x7b, 0x1f, 0xd1, 0xf6, 0x9f, 0xac, 0xf0,
	0x26, 0x5a, 0xba, 0x84, 0x89, 0xee, 0xa0, 0xe4, 0xa6, 0x4b, 0xec, 0xa0, 0xe5, 0x31, 0x89, 0x46,
	0x60, 0xe7, 0x75, 0x57, 0xff, 0xcc, 0x25, 0x5b, 0x74, 0x70, 0x0d, 0xef, 0x24, 0x7f, 0x6c, 0x35,
	0x3e, 0xa3, 0xa2, 0x69, 0x19, 0x33, 0xb4, 0x3e, 0x00, 0xf0, 0x12, 0x10, 0x3e, 0x30, 0x45, 0x42,
	0x30, 0xde, 0x9d, 0xb3, 0x34, 0xf3, 0xcf, 0x69, 0xed, 0xdf, 0x90, 0xaa, 0xe1, 0xa8, 0xdf, 0xf6,
	0x79, 0xec, 0xf8, 0x5c, 0xc6, 0x5c, 0x66, 0x9f, 0x96, 0x0c, 0x2e, 0x1d, 0x35, 0x49, 0x40, 0xb6,
	0x4f, 0xc1, 0xbf, 0x9f, 0xd6, 0x76, 0x4c, 0x77, 0x8b, 0x6e, 0x0d, 0x77, 0x6d, 0x00, 0xd0, 0x7d,
	0xdc, 0x7f, 0x2f, 0xa0, 0xf5, 0xc5, 0x5c, 0xf8, 0x25, 0xda, 0xe5, 0x82, 0x86, 0x94, 0x91, 0xc8,
	0x93, 0xc0, 0x02, 0x10, 0x1e, 0x09, 0x02, 0x01, 0x52, 0x66, 0x7d, 0xee, 0xcc, 0xe0, 0x9e, 0x46,
	0x5f, 0x19, 0x10, 0xff, 0x8f, 0xb6, 0x04, 0x0c, 0x46, 0x2c, 0xf0, 0xfc, 0x21, 0x61, 0x0c, 0x22,
	0x8f, 0x06, 0x7a, 0x0a, 0x25, 0x77, 0xc3, 0x00, 0xaf, 0x4d, 0xfd, 0x3c, 0xc0, 0xfb, 0x68, 0x3d,
	0xe3, 0x26, 0x5c, 0xa8, 0x94, 0xb8, 0xa4, 0x89, 0x15, 0x53, 0xed, 0x72, 0xa1, 0xce, 0x03, 0x7c,
	0x88, 0x76, 0xcc, 0x2f, 0xf1, 0xa4, 0xf0, 0xe7, 0x5d, 0x0b, 0x9a, 0x8c, 0x0d, 0xd8, 0x13, 0xfe,
	0xa3, 0xf1, 0x01, 0xc2, 0x73, 0x92, 0x99, 0xf9, 0xb2, 0x49, 0xf1, 0xc0, 0xcf, 0xfc, 0x8f, 0x91,
	0x9d, 0x91, 0x15, 0x8d, 0x81, 0x8f, 0xcc, 0x57, 0x2a, 0x12, 0x27, 0x76, 0xb1, 0x6e, 0x35, 0x0b,
	0xee, 0xdf, 0x06, 0xbf, 0x30, 0xf0, 0xc5, 0x0c, 0xc5, 0x47, 0x0f, 0xc9, 0x66, 0xca, 0x21, 0xa4,
	0x23, 0xb4, 0x57, 0xf4, 0x49, 0x7f, 0x2d, 0xc8, 0xde, 0x69, 0x08, 0xd7, 0x50, 0x39, 0xd3, 0x04,
	0x44, 0x11, 0x7b, 0xb5, 0x6e, 0x35, 0x2b, 0x2e, 0x32, 0xa5, 0x53, 0xa2, 0x08, 0xfe, 0x0f, 0x65,
	0x73, 0xf2, 0x24, 0x7c, 0x1a, 0x01, 0xf3, 0xc1, 0x2e, 0xe9, 0x14, 0xd9, 0xac, 0x7a, 0x59, 0x15,
	0x1f, 0xa4, 0x93, 0x56, 0x82, 0x82, 0xf4, 0x04, 0xc4, 0x84, 0x32, 0xca, 0x42, 0x1b, 0xd5, 0xad,
	0xe6, 0xb2, 0xbb, 0x99, 0x01, 0xee, 0xac, 0x8e, 0x6d, 0xb4, 0x92, 0x65, 0xb4, 0xcb, 0xda, 0x6d,
	0xb6, 0xc5, 0xfb, 0x68, 0x8d, 0x71, 0x66, 0xbc, 0x49, 0x3f, 0x02, 0xbb, 0x52, 0xb7, 0x9a, 0xab,
	0xee, 0x62, 0xb1, 0x43, 0xaf, 0x6f, 0xab, 0xd6, 0xcd, 0x6d, 0xd5, 0xfa, 0x75, 0x5b, 0xb5, 0xbe,
	0xde, 0x55, 0x73, 0x37, 0x77, 0xd5, 0xdc, 0x8f, 0xbb, 0x6a, 0xee, 0xc3, 0xfb, 0xa7, 0x77, 0x91,
	0xf6, 0xfd, 0x16, 0x49, 0x12, 0xe9, 0xc4, 0x34, 0x08, 0x22, 0xb8, 0x22, 0x02, 0x1c, 0xd3, 0x61,
	0x6b, 0xc0, 0xc5, 0x15, 0x11, 0x41, 0x6b, 0x0e, 0x19, 0xbf, 0x70, 0xb2, 0xf7, 0x44, 0x5f, 0xdc,
	0x7e, 0x51, 0x3f, 0x15, 0xcf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xed, 0x2e, 0x73, 0xcd, 0x66,
	0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InFlightPackets) > 0 {
		for k := range m.InFlightPackets {
			v := m.InFlightPackets[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintGenesis(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenesis(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.FeePercentage.Size()
		i -= size
		if _, err := m.FeePercentage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *InFlightPacket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InFlightPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InFlightPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Nonrefundable {
		i--
		if m.Nonrefundable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x60
	}
	if m.Timeout != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Timeout))
		i--
		dAtA[i] = 0x58
	}
	if m.RetriesRemaining != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.RetriesRemaining))
		i--
		dAtA[i] = 0x50
	}
	if m.RefundSequence != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.RefundSequence))
		i--
		dAtA[i] = 0x48
	}
	if len(m.PacketData) > 0 {
		i -= len(m.PacketData)
		copy(dAtA[i:], m.PacketData)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PacketData)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.PacketTimeoutHeight) > 0 {
		i -= len(m.PacketTimeoutHeight)
		copy(dAtA[i:], m.PacketTimeoutHeight)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PacketTimeoutHeight)))
		i--
		dAtA[i] = 0x3a
	}
	if m.PacketTimeoutTimestamp != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PacketTimeoutTimestamp))
		i--
		dAtA[i] = 0x30
	}
	if len(m.PacketSrcPortId) > 0 {
		i -= len(m.PacketSrcPortId)
		copy(dAtA[i:], m.PacketSrcPortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PacketSrcPortId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PacketSrcChannelId) > 0 {
		i -= len(m.PacketSrcChannelId)
		copy(dAtA[i:], m.PacketSrcChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PacketSrcChannelId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RefundPortId) > 0 {
		i -= len(m.RefundPortId)
		copy(dAtA[i:], m.RefundPortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.RefundPortId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RefundChannelId) > 0 {
		i -= len(m.RefundChannelId)
		copy(dAtA[i:], m.RefundChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.RefundChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OriginalSenderAddress) > 0 {
		i -= len(m.OriginalSenderAddress)
		copy(dAtA[i:], m.OriginalSenderAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.OriginalSenderAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.InFlightPackets) > 0 {
		for k, v := range m.InFlightPackets {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovGenesis(uint64(len(k))) + 1 + l + sovGenesis(uint64(l))
			n += mapEntrySize + 1 + sovGenesis(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.FeePercentage.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *InFlightPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OriginalSenderAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.RefundChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.RefundPortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.PacketSrcChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.PacketSrcPortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.PacketTimeoutTimestamp != 0 {
		n += 1 + sovGenesis(uint64(m.PacketTimeoutTimestamp))
	}
	l = len(m.PacketTimeoutHeight)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.PacketData)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.RefundSequence != 0 {
		n += 1 + sovGenesis(uint64(m.RefundSequence))
	}
	if m.RetriesRemaining != 0 {
		n += 1 + sovGenesis(uint64(m.RetriesRemaining))
	}
	if m.Timeout != 0 {
		n += 1 + sovGenesis(uint64(m.Timeout))
	}
	if m.Nonrefundable {
		n += 2
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InFlightPackets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InFlightPackets == nil {
				m.InFlightPackets = make(map[string]InFlightPacket)
			}
			var mapkey string
			mapvalue := &InFlightPacket{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthGenesis
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthGenesis
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &InFlightPacket{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenesis(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenesis
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.InFlightPackets[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeePercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeePercentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InFlightPacket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InFlightPacket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InFlightPacket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginalSenderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginalSenderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefundChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefundChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefundPortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefundPortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketSrcChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PacketSrcChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketSrcPortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PacketSrcPortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketTimeoutTimestamp", wireType)
			}
			m.PacketTimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PacketTimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketTimeoutHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PacketTimeoutHeight = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PacketData = append(m.PacketData[:0], dAtA[iNdEx:postIndex]...)
			if m.PacketData == nil {
				m.PacketData = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefundSequence", wireType)
			}
			m.RefundSequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RefundSequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetriesRemaining", wireType)
			}
			m.RetriesRemaining = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RetriesRemaining |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonrefundable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Nonrefundable = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)