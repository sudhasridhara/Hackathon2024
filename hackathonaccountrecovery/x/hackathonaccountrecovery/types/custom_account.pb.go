// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hackathonaccountrecovery/hackathonaccountrecovery/custom_account.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
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

type CustomAccount struct {
	BaseAccount      string `protobuf:"bytes,1,opt,name=baseAccount,proto3" json:"baseAccount,omitempty"`
	RecoveryMultisig string `protobuf:"bytes,2,opt,name=recoveryMultisig,proto3" json:"recoveryMultisig,omitempty"`
}

func (m *CustomAccount) Reset()         { *m = CustomAccount{} }
func (m *CustomAccount) String() string { return proto.CompactTextString(m) }
func (*CustomAccount) ProtoMessage()    {}
func (*CustomAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebbf9cb3dec323a2, []int{0}
}
func (m *CustomAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CustomAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CustomAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CustomAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomAccount.Merge(m, src)
}
func (m *CustomAccount) XXX_Size() int {
	return m.Size()
}
func (m *CustomAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomAccount.DiscardUnknown(m)
}

var xxx_messageInfo_CustomAccount proto.InternalMessageInfo

func (m *CustomAccount) GetBaseAccount() string {
	if m != nil {
		return m.BaseAccount
	}
	return ""
}

func (m *CustomAccount) GetRecoveryMultisig() string {
	if m != nil {
		return m.RecoveryMultisig
	}
	return ""
}

func init() {
	proto.RegisterType((*CustomAccount)(nil), "hackathonaccountrecovery.hackathonaccountrecovery.CustomAccount")
}

func init() {
	proto.RegisterFile("hackathonaccountrecovery/hackathonaccountrecovery/custom_account.proto", fileDescriptor_ebbf9cb3dec323a2)
}

var fileDescriptor_ebbf9cb3dec323a2 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0xcb, 0x48, 0x4c, 0xce,
	0x4e, 0x2c, 0xc9, 0xc8, 0xcf, 0x4b, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x29, 0x4a, 0x4d, 0xce,
	0x2f, 0x4b, 0x2d, 0xaa, 0xd4, 0xc7, 0x29, 0x91, 0x5c, 0x5a, 0x5c, 0x92, 0x9f, 0x1b, 0x0f, 0x15,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x32, 0xc4, 0xa5, 0x5c, 0x0f, 0x97, 0x84, 0x52, 0x2c,
	0x17, 0xaf, 0x33, 0xd8, 0x28, 0x47, 0x88, 0x84, 0x90, 0x02, 0x17, 0x77, 0x52, 0x62, 0x71, 0x2a,
	0x94, 0x2b, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x2c, 0x24, 0xa4, 0xc5, 0x25, 0x00, 0xd3,
	0xee, 0x5b, 0x9a, 0x53, 0x92, 0x59, 0x9c, 0x99, 0x2e, 0xc1, 0x04, 0x56, 0x86, 0x21, 0xee, 0x14,
	0x7c, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c,
	0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x96, 0x38, 0xbd, 0x56, 0x81,
	0xdb, 0xd7, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0xdf, 0x1a, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x04, 0xff, 0x67, 0xd1, 0x37, 0x01, 0x00, 0x00,
}

func (m *CustomAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CustomAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CustomAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RecoveryMultisig) > 0 {
		i -= len(m.RecoveryMultisig)
		copy(dAtA[i:], m.RecoveryMultisig)
		i = encodeVarintCustomAccount(dAtA, i, uint64(len(m.RecoveryMultisig)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BaseAccount) > 0 {
		i -= len(m.BaseAccount)
		copy(dAtA[i:], m.BaseAccount)
		i = encodeVarintCustomAccount(dAtA, i, uint64(len(m.BaseAccount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCustomAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovCustomAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CustomAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BaseAccount)
	if l > 0 {
		n += 1 + l + sovCustomAccount(uint64(l))
	}
	l = len(m.RecoveryMultisig)
	if l > 0 {
		n += 1 + l + sovCustomAccount(uint64(l))
	}
	return n
}

func sovCustomAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCustomAccount(x uint64) (n int) {
	return sovCustomAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CustomAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCustomAccount
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
			return fmt.Errorf("proto: CustomAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CustomAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomAccount
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
				return ErrInvalidLengthCustomAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCustomAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecoveryMultisig", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomAccount
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
				return ErrInvalidLengthCustomAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCustomAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecoveryMultisig = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCustomAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCustomAccount
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
func skipCustomAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCustomAccount
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
					return 0, ErrIntOverflowCustomAccount
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
					return 0, ErrIntOverflowCustomAccount
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
				return 0, ErrInvalidLengthCustomAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCustomAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCustomAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCustomAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCustomAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCustomAccount = fmt.Errorf("proto: unexpected end of group")
)
