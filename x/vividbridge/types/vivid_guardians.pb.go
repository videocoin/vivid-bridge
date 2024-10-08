// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vividbridge/vividbridge/vivid_guardians.proto

package types

import (
	fmt "fmt"
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

type VividGuardians struct {
	Id             uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Keys           []string `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	ExpirationTime uint64   `protobuf:"varint,3,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	Creator        string   `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *VividGuardians) Reset()         { *m = VividGuardians{} }
func (m *VividGuardians) String() string { return proto.CompactTextString(m) }
func (*VividGuardians) ProtoMessage()    {}
func (*VividGuardians) Descriptor() ([]byte, []int) {
	return fileDescriptor_073391f1b97a9db4, []int{0}
}
func (m *VividGuardians) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VividGuardians) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VividGuardians.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VividGuardians) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VividGuardians.Merge(m, src)
}
func (m *VividGuardians) XXX_Size() int {
	return m.Size()
}
func (m *VividGuardians) XXX_DiscardUnknown() {
	xxx_messageInfo_VividGuardians.DiscardUnknown(m)
}

var xxx_messageInfo_VividGuardians proto.InternalMessageInfo

func (m *VividGuardians) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VividGuardians) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *VividGuardians) GetExpirationTime() uint64 {
	if m != nil {
		return m.ExpirationTime
	}
	return 0
}

func (m *VividGuardians) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*VividGuardians)(nil), "vividbridge.vividbridge.VividGuardians")
}

func init() {
	proto.RegisterFile("vividbridge/vividbridge/vivid_guardians.proto", fileDescriptor_073391f1b97a9db4)
}

var fileDescriptor_073391f1b97a9db4 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2d, 0xcb, 0x2c, 0xcb,
	0x4c, 0x49, 0x2a, 0xca, 0x4c, 0x49, 0x4f, 0xd5, 0xc7, 0x60, 0xc7, 0xa7, 0x97, 0x26, 0x16, 0xa5,
	0x64, 0x26, 0xe6, 0x15, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x23, 0x29, 0xd1, 0x43,
	0x62, 0x2b, 0x95, 0x71, 0xf1, 0x85, 0x81, 0xb8, 0xee, 0x30, 0x0d, 0x42, 0x7c, 0x5c, 0x4c, 0x99,
	0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x42, 0x5c, 0x2c, 0xd9,
	0xa9, 0x95, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x1a, 0x17, 0x5f,
	0x6a, 0x45, 0x41, 0x66, 0x51, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x48, 0x66, 0x6e, 0xaa, 0x04, 0x33,
	0x58, 0x3d, 0x9a, 0xa8, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x51, 0x6a, 0x62, 0x49, 0x7e, 0x91, 0x04,
	0x8b, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0xeb, 0x64, 0x75, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47,
	0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d,
	0xc7, 0x72, 0x0c, 0x51, 0x0a, 0x60, 0xe7, 0xe9, 0x42, 0xbd, 0x53, 0x81, 0xe2, 0xb9, 0x92, 0xca,
	0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x9f, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x36, 0x9d,
	0x70, 0x44, 0x04, 0x01, 0x00, 0x00,
}

func (m *VividGuardians) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VividGuardians) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VividGuardians) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintVividGuardians(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if m.ExpirationTime != 0 {
		i = encodeVarintVividGuardians(dAtA, i, uint64(m.ExpirationTime))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Keys) > 0 {
		for iNdEx := len(m.Keys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Keys[iNdEx])
			copy(dAtA[i:], m.Keys[iNdEx])
			i = encodeVarintVividGuardians(dAtA, i, uint64(len(m.Keys[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Id != 0 {
		i = encodeVarintVividGuardians(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVividGuardians(dAtA []byte, offset int, v uint64) int {
	offset -= sovVividGuardians(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VividGuardians) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovVividGuardians(uint64(m.Id))
	}
	if len(m.Keys) > 0 {
		for _, s := range m.Keys {
			l = len(s)
			n += 1 + l + sovVividGuardians(uint64(l))
		}
	}
	if m.ExpirationTime != 0 {
		n += 1 + sovVividGuardians(uint64(m.ExpirationTime))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovVividGuardians(uint64(l))
	}
	return n
}

func sovVividGuardians(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVividGuardians(x uint64) (n int) {
	return sovVividGuardians(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VividGuardians) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVividGuardians
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
			return fmt.Errorf("proto: VividGuardians: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VividGuardians: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVividGuardians
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVividGuardians
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
				return ErrInvalidLengthVividGuardians
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVividGuardians
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keys = append(m.Keys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationTime", wireType)
			}
			m.ExpirationTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVividGuardians
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpirationTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVividGuardians
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
				return ErrInvalidLengthVividGuardians
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVividGuardians
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVividGuardians(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVividGuardians
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
func skipVividGuardians(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVividGuardians
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
					return 0, ErrIntOverflowVividGuardians
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
					return 0, ErrIntOverflowVividGuardians
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
				return 0, ErrInvalidLengthVividGuardians
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVividGuardians
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVividGuardians
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVividGuardians        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVividGuardians          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVividGuardians = fmt.Errorf("proto: unexpected end of group")
)
