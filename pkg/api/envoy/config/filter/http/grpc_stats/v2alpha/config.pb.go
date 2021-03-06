// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/grpc_stats/v2alpha/config.proto

package envoy_config_filter_http_grpc_stats_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// gRPC statistics filter configuration
type FilterConfig struct {
	// If true, the filter maintains a filter state object with the request and response message
	// counts.
	EmitFilterState      bool     `protobuf:"varint,1,opt,name=emit_filter_state,json=emitFilterState,proto3" json:"emit_filter_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_1419fab6b23f453d, []int{0}
}
func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(m, src)
}
func (m *FilterConfig) XXX_Size() int {
	return m.Size()
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetEmitFilterState() bool {
	if m != nil {
		return m.EmitFilterState
	}
	return false
}

// gRPC statistics filter state object in protobuf form.
type FilterObject struct {
	// Count of request messages in the request stream.
	RequestMessageCount uint64 `protobuf:"varint,1,opt,name=request_message_count,json=requestMessageCount,proto3" json:"request_message_count,omitempty"`
	// Count of response messages in the response stream.
	ResponseMessageCount uint64   `protobuf:"varint,2,opt,name=response_message_count,json=responseMessageCount,proto3" json:"response_message_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterObject) Reset()         { *m = FilterObject{} }
func (m *FilterObject) String() string { return proto.CompactTextString(m) }
func (*FilterObject) ProtoMessage()    {}
func (*FilterObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_1419fab6b23f453d, []int{1}
}
func (m *FilterObject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FilterObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FilterObject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FilterObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterObject.Merge(m, src)
}
func (m *FilterObject) XXX_Size() int {
	return m.Size()
}
func (m *FilterObject) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterObject.DiscardUnknown(m)
}

var xxx_messageInfo_FilterObject proto.InternalMessageInfo

func (m *FilterObject) GetRequestMessageCount() uint64 {
	if m != nil {
		return m.RequestMessageCount
	}
	return 0
}

func (m *FilterObject) GetResponseMessageCount() uint64 {
	if m != nil {
		return m.ResponseMessageCount
	}
	return 0
}

func init() {
	proto.RegisterType((*FilterConfig)(nil), "envoy.config.filter.http.grpc_stats.v2alpha.FilterConfig")
	proto.RegisterType((*FilterObject)(nil), "envoy.config.filter.http.grpc_stats.v2alpha.FilterObject")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/grpc_stats/v2alpha/config.proto", fileDescriptor_1419fab6b23f453d)
}

var fileDescriptor_1419fab6b23f453d = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x4d, 0x4a, 0x03, 0x31,
	0x14, 0x80, 0x89, 0xa8, 0x48, 0x14, 0xc4, 0xf1, 0x97, 0x2e, 0x06, 0xe9, 0x4a, 0x14, 0x13, 0x68,
	0x5d, 0xa8, 0xcb, 0x16, 0xdc, 0x89, 0xa5, 0x1e, 0x60, 0x48, 0xa7, 0xaf, 0xd3, 0x48, 0x9b, 0xc4,
	0xe4, 0xb5, 0xb4, 0x47, 0x70, 0xe5, 0xd6, 0x23, 0xb9, 0xf4, 0x08, 0x32, 0x47, 0xf0, 0x00, 0x22,
	0x49, 0x66, 0x50, 0x74, 0xe5, 0x2e, 0xe4, 0x7b, 0xdf, 0x7b, 0xf0, 0xd1, 0x4b, 0x50, 0x73, 0xbd,
	0xe4, 0xb9, 0x56, 0x23, 0x59, 0xf0, 0x91, 0x9c, 0x20, 0x58, 0x3e, 0x46, 0x34, 0xbc, 0xb0, 0x26,
	0xcf, 0x1c, 0x0a, 0x74, 0x7c, 0xde, 0x12, 0x13, 0x33, 0x16, 0xd5, 0x14, 0x33, 0x56, 0xa3, 0x4e,
	0xce, 0x82, 0xc9, 0xaa, 0xbf, 0x68, 0x32, 0x6f, 0xb2, 0x6f, 0x93, 0x55, 0x66, 0x23, 0x9d, 0x0d,
	0x8d, 0xe0, 0x42, 0x29, 0x8d, 0x02, 0xa5, 0x56, 0x8e, 0x4f, 0x65, 0x61, 0x05, 0x42, 0x5c, 0xd6,
	0x38, 0x9c, 0x8b, 0x89, 0x1c, 0x0a, 0x04, 0x5e, 0x3f, 0x22, 0x68, 0x5e, 0xd3, 0xad, 0x9b, 0xb0,
	0xba, 0x1b, 0xee, 0x24, 0xa7, 0x74, 0x07, 0xa6, 0x12, 0xb3, 0x78, 0x2f, 0x5c, 0x81, 0x23, 0x72,
	0x4c, 0x4e, 0x36, 0xfa, 0xdb, 0x1e, 0xc4, 0xe1, 0x7b, 0xff, 0xdd, 0x5c, 0xd4, 0xee, 0xdd, 0xe0,
	0x01, 0x72, 0x4c, 0x5a, 0x74, 0xdf, 0xc2, 0xe3, 0x0c, 0x1c, 0x66, 0x53, 0x70, 0x4e, 0x14, 0x90,
	0xe5, 0x7a, 0xa6, 0x30, 0xf8, 0xab, 0xfd, 0xdd, 0x0a, 0xde, 0x46, 0xd6, 0xf5, 0x28, 0xb9, 0xa0,
	0x07, 0x16, 0x9c, 0xd1, 0xca, 0xc1, 0x2f, 0x69, 0x25, 0x48, 0x7b, 0x35, 0xfd, 0x69, 0x75, 0x9e,
	0xc8, 0x6b, 0x99, 0x92, 0xb7, 0x32, 0x25, 0xef, 0x65, 0x4a, 0x3e, 0x5e, 0x3e, 0x9f, 0xd7, 0xce,
	0xeb, 0x60, 0xb0, 0x40, 0x50, 0xce, 0x37, 0xa8, 0xa2, 0xb9, 0xbf, 0xd5, 0xda, 0xf4, 0x4a, 0x6a,
	0x16, 0xe6, 0x8d, 0xd5, 0x8b, 0x25, 0xfb, 0x47, 0xeb, 0xce, 0x66, 0x4c, 0xd5, 0xf3, 0xfd, 0x7a,
	0x64, 0xb0, 0x1e, 0x42, 0xb6, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x8a, 0x9c, 0x36, 0xea,
	0x01, 0x00, 0x00,
}

func (m *FilterConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilterConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FilterConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.EmitFilterState {
		i--
		if m.EmitFilterState {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FilterObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilterObject) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FilterObject) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ResponseMessageCount != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.ResponseMessageCount))
		i--
		dAtA[i] = 0x10
	}
	if m.RequestMessageCount != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.RequestMessageCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FilterConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EmitFilterState {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FilterObject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestMessageCount != 0 {
		n += 1 + sovConfig(uint64(m.RequestMessageCount))
	}
	if m.ResponseMessageCount != 0 {
		n += 1 + sovConfig(uint64(m.ResponseMessageCount))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FilterConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: FilterConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilterConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmitFilterState", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
			m.EmitFilterState = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FilterObject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: FilterObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilterObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestMessageCount", wireType)
			}
			m.RequestMessageCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestMessageCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseMessageCount", wireType)
			}
			m.ResponseMessageCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResponseMessageCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)
