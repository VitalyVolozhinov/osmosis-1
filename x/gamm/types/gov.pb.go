// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/gamm/v1beta1/gov.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	migration "github.com/osmosis-labs/osmosis/v17/x/gamm/types/migration"
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

// ReplaceMigrationRecordsProposal is a gov Content type for updating the
// migration records. If a ReplaceMigrationRecordsProposal passes, the
// proposal’s records override the existing MigrationRecords set in the module.
// Each record specifies a single connection between a single balancer pool and
// a single concentrated pool.
type ReplaceMigrationRecordsProposal struct {
	Title       string                                     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                                     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Records     []migration.BalancerToConcentratedPoolLink `protobuf:"bytes,3,rep,name=records,proto3" json:"records"`
}

func (m *ReplaceMigrationRecordsProposal) Reset()      { *m = ReplaceMigrationRecordsProposal{} }
func (*ReplaceMigrationRecordsProposal) ProtoMessage() {}
func (*ReplaceMigrationRecordsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_f31b9a6c0dbbdfa3, []int{0}
}
func (m *ReplaceMigrationRecordsProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReplaceMigrationRecordsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReplaceMigrationRecordsProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplaceMigrationRecordsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplaceMigrationRecordsProposal.Merge(m, src)
}
func (m *ReplaceMigrationRecordsProposal) XXX_Size() int {
	return m.Size()
}
func (m *ReplaceMigrationRecordsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplaceMigrationRecordsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_ReplaceMigrationRecordsProposal proto.InternalMessageInfo

// For example: if the existing DistrRecords were:
// [(Balancer 1, CL 5), (Balancer 2, CL 6), (Balancer 3, CL 7)]
// And an UpdateMigrationRecordsProposal includes
// [(Balancer 2, CL 0), (Balancer 3, CL 4), (Balancer 4, CL 10)]
// This would leave Balancer 1 record, delete Balancer 2 record,
// Edit Balancer 3 record, and Add Balancer 4 record
// The result MigrationRecords in state would be:
// [(Balancer 1, CL 5), (Balancer 3, CL 4), (Balancer 4, CL 10)]
type UpdateMigrationRecordsProposal struct {
	Title       string                                     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                                     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Records     []migration.BalancerToConcentratedPoolLink `protobuf:"bytes,3,rep,name=records,proto3" json:"records"`
}

func (m *UpdateMigrationRecordsProposal) Reset()      { *m = UpdateMigrationRecordsProposal{} }
func (*UpdateMigrationRecordsProposal) ProtoMessage() {}
func (*UpdateMigrationRecordsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_f31b9a6c0dbbdfa3, []int{1}
}
func (m *UpdateMigrationRecordsProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateMigrationRecordsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateMigrationRecordsProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateMigrationRecordsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMigrationRecordsProposal.Merge(m, src)
}
func (m *UpdateMigrationRecordsProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateMigrationRecordsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMigrationRecordsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMigrationRecordsProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ReplaceMigrationRecordsProposal)(nil), "osmosis.gamm.v1beta1.ReplaceMigrationRecordsProposal")
	proto.RegisterType((*UpdateMigrationRecordsProposal)(nil), "osmosis.gamm.v1beta1.UpdateMigrationRecordsProposal")
}

func init() { proto.RegisterFile("osmosis/gamm/v1beta1/gov.proto", fileDescriptor_f31b9a6c0dbbdfa3) }

var fileDescriptor_f31b9a6c0dbbdfa3 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0x2f, 0xce, 0xcd,
	0x2f, 0xce, 0x2c, 0xd6, 0x4f, 0x4f, 0xcc, 0xcd, 0xd5, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34,
	0xd4, 0x4f, 0xcf, 0x2f, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x81, 0xca, 0xeb, 0x81,
	0xe4, 0xf5, 0xa0, 0xf2, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x05, 0xfa, 0x20, 0x16, 0x44,
	0xad, 0x94, 0x12, 0x76, 0xb3, 0x52, 0xf3, 0x52, 0x41, 0x06, 0x40, 0xd4, 0x28, 0x62, 0x55, 0x53,
	0x9c, 0x91, 0x58, 0x94, 0x9a, 0x02, 0x55, 0x22, 0x99, 0x0c, 0x56, 0x13, 0x0f, 0x31, 0x1f, 0xc2,
	0x81, 0x4a, 0x09, 0x26, 0xe6, 0x66, 0xe6, 0xe5, 0xeb, 0x83, 0x49, 0x88, 0x90, 0x52, 0x3b, 0x13,
	0x97, 0x7c, 0x50, 0x6a, 0x41, 0x4e, 0x62, 0x72, 0xaa, 0x6f, 0x66, 0x7a, 0x51, 0x62, 0x49, 0x66,
	0x7e, 0x5e, 0x50, 0x6a, 0x72, 0x7e, 0x51, 0x4a, 0x71, 0x40, 0x51, 0x7e, 0x41, 0x7e, 0x71, 0x62,
	0x8e, 0x90, 0x08, 0x17, 0x6b, 0x49, 0x66, 0x49, 0x4e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x84, 0x23, 0xa4, 0xc0, 0xc5, 0x9d, 0x92, 0x5a, 0x9c, 0x5c, 0x94, 0x59, 0x00, 0xd2, 0x23,
	0xc1, 0x04, 0x96, 0x43, 0x16, 0x12, 0x0a, 0xe1, 0x62, 0x2f, 0x82, 0x18, 0x25, 0xc1, 0xac, 0xc0,
	0xac, 0xc1, 0x6d, 0x64, 0xa2, 0x87, 0x2d, 0x38, 0xf4, 0x9c, 0x12, 0x73, 0x12, 0xf3, 0x92, 0x53,
	0x8b, 0x42, 0xf2, 0x9d, 0xf3, 0xf3, 0x92, 0x53, 0xf3, 0x4a, 0x8a, 0x12, 0x4b, 0x52, 0x53, 0x02,
	0xf2, 0xf3, 0x73, 0x7c, 0x32, 0xf3, 0xb2, 0x9d, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0x82, 0x19,
	0x65, 0x15, 0xd6, 0xb1, 0x40, 0x9e, 0x61, 0xc6, 0x02, 0x79, 0x86, 0x17, 0x0b, 0xe4, 0x19, 0x4f,
	0x6d, 0xd1, 0x95, 0x82, 0x7a, 0x11, 0x14, 0xe8, 0x30, 0x13, 0x9d, 0xf3, 0xf3, 0x4a, 0x52, 0xf3,
	0x4a, 0xba, 0x9e, 0x6f, 0xd0, 0x52, 0x87, 0x85, 0x18, 0x01, 0x5f, 0x2a, 0xb5, 0x32, 0x71, 0xc9,
	0x85, 0x16, 0xa4, 0x24, 0x96, 0x0c, 0x95, 0x80, 0x08, 0x25, 0x2d, 0x20, 0xd4, 0x60, 0x01, 0x81,
	0xdf, 0x93, 0x4e, 0x5e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c,
	0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x90,
	0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x35, 0x4c, 0x37, 0x27, 0x31,
	0xa9, 0x18, 0xc6, 0xd1, 0x2f, 0x33, 0x34, 0xd7, 0xaf, 0x80, 0x24, 0xcd, 0x92, 0xca, 0x82, 0xd4,
	0xe2, 0x24, 0x36, 0x70, 0x22, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xe5, 0x3d, 0x03,
	0x27, 0x03, 0x00, 0x00,
}

func (this *ReplaceMigrationRecordsProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReplaceMigrationRecordsProposal)
	if !ok {
		that2, ok := that.(ReplaceMigrationRecordsProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.Records) != len(that1.Records) {
		return false
	}
	for i := range this.Records {
		if !this.Records[i].Equal(&that1.Records[i]) {
			return false
		}
	}
	return true
}
func (this *UpdateMigrationRecordsProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateMigrationRecordsProposal)
	if !ok {
		that2, ok := that.(UpdateMigrationRecordsProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.Records) != len(that1.Records) {
		return false
	}
	for i := range this.Records {
		if !this.Records[i].Equal(&that1.Records[i]) {
			return false
		}
	}
	return true
}
func (m *ReplaceMigrationRecordsProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplaceMigrationRecordsProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReplaceMigrationRecordsProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateMigrationRecordsProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateMigrationRecordsProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateMigrationRecordsProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReplaceMigrationRecordsProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func (m *UpdateMigrationRecordsProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReplaceMigrationRecordsProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: ReplaceMigrationRecordsProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplaceMigrationRecordsProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Records = append(m.Records, migration.BalancerToConcentratedPoolLink{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *UpdateMigrationRecordsProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: UpdateMigrationRecordsProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateMigrationRecordsProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Records = append(m.Records, migration.BalancerToConcentratedPoolLink{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
