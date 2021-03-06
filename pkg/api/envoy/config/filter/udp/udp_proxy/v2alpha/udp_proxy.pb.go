// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/udp/udp_proxy/v2alpha/udp_proxy.proto

package envoy_config_filter_udp_udp_proxy_v2alpha

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// Configuration for the UDP proxy filter.
type UdpProxyConfig struct {
	// The stat prefix used when emitting UDP proxy filter stats.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Types that are valid to be assigned to RouteSpecifier:
	//	*UdpProxyConfig_Cluster
	RouteSpecifier isUdpProxyConfig_RouteSpecifier `protobuf_oneof:"route_specifier"`
	// The idle timeout for sessions. Idle is defined as no datagrams between received or sent by
	// the session. The default if not specified is 1 minute.
	IdleTimeout          *types.Duration `protobuf:"bytes,3,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UdpProxyConfig) Reset()         { *m = UdpProxyConfig{} }
func (m *UdpProxyConfig) String() string { return proto.CompactTextString(m) }
func (*UdpProxyConfig) ProtoMessage()    {}
func (*UdpProxyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_82ee653a514c2977, []int{0}
}
func (m *UdpProxyConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UdpProxyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UdpProxyConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UdpProxyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UdpProxyConfig.Merge(m, src)
}
func (m *UdpProxyConfig) XXX_Size() int {
	return m.Size()
}
func (m *UdpProxyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_UdpProxyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_UdpProxyConfig proto.InternalMessageInfo

type isUdpProxyConfig_RouteSpecifier interface {
	isUdpProxyConfig_RouteSpecifier()
	MarshalTo([]byte) (int, error)
	Size() int
}

type UdpProxyConfig_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof" json:"cluster,omitempty"`
}

func (*UdpProxyConfig_Cluster) isUdpProxyConfig_RouteSpecifier() {}

func (m *UdpProxyConfig) GetRouteSpecifier() isUdpProxyConfig_RouteSpecifier {
	if m != nil {
		return m.RouteSpecifier
	}
	return nil
}

func (m *UdpProxyConfig) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *UdpProxyConfig) GetCluster() string {
	if x, ok := m.GetRouteSpecifier().(*UdpProxyConfig_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *UdpProxyConfig) GetIdleTimeout() *types.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UdpProxyConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UdpProxyConfig_Cluster)(nil),
	}
}

func init() {
	proto.RegisterType((*UdpProxyConfig)(nil), "envoy.config.filter.udp.udp_proxy.v2alpha.UdpProxyConfig")
}

func init() {
	proto.RegisterFile("envoy/config/filter/udp/udp_proxy/v2alpha/udp_proxy.proto", fileDescriptor_82ee653a514c2977)
}

var fileDescriptor_82ee653a514c2977 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4a, 0xec, 0x30,
	0x18, 0xc5, 0x6f, 0x66, 0x2e, 0x0e, 0x66, 0xfc, 0x03, 0x5d, 0xe8, 0x38, 0x8b, 0x32, 0xe8, 0xa6,
	0x6e, 0x12, 0x18, 0x17, 0x22, 0xb8, 0x8a, 0x2e, 0x5c, 0x96, 0xa2, 0x6e, 0x4b, 0xa6, 0x49, 0x6b,
	0x20, 0x36, 0x21, 0x4d, 0xca, 0xcc, 0x83, 0xf9, 0x0e, 0x2e, 0x7d, 0x04, 0xe9, 0x53, 0x88, 0x2b,
	0x69, 0x32, 0x65, 0x70, 0xe7, 0x22, 0x10, 0x72, 0xce, 0xf9, 0xf2, 0xfb, 0x0e, 0xbc, 0xe1, 0x75,
	0xab, 0x36, 0xb8, 0x50, 0x75, 0x29, 0x2a, 0x5c, 0x0a, 0x69, 0xb9, 0xc1, 0x8e, 0xe9, 0xfe, 0xe4,
	0xda, 0xa8, 0xf5, 0x06, 0xb7, 0x4b, 0x2a, 0xf5, 0x0b, 0xdd, 0xbd, 0x20, 0x6d, 0x94, 0x55, 0xd1,
	0xa5, 0x8f, 0xa2, 0x10, 0x45, 0x21, 0x8a, 0x1c, 0xd3, 0x68, 0x67, 0xdc, 0x46, 0xe7, 0x71, 0xa5,
	0x54, 0x25, 0x39, 0xf6, 0xc1, 0x95, 0x2b, 0x31, 0x73, 0x86, 0x5a, 0xa1, 0xea, 0x30, 0x6a, 0x7e,
	0xda, 0x52, 0x29, 0x18, 0xb5, 0x1c, 0x0f, 0x97, 0x20, 0x9c, 0xbf, 0x01, 0x78, 0xf4, 0xc4, 0x74,
	0xda, 0x4f, 0xbb, 0xf3, 0x1f, 0x45, 0x09, 0x9c, 0x36, 0x96, 0xda, 0x5c, 0x1b, 0x5e, 0x8a, 0xf5,
	0x0c, 0x2c, 0x40, 0xb2, 0x4f, 0x26, 0xdf, 0xe4, 0xbf, 0x19, 0x2d, 0x40, 0x06, 0x7b, 0x2d, 0xf5,
	0x52, 0x74, 0x01, 0x27, 0x85, 0x74, 0x8d, 0xe5, 0x66, 0x36, 0xfa, 0xe5, 0x7a, 0xf8, 0x97, 0x0d,
	0x4a, 0x74, 0x0b, 0x0f, 0x04, 0x93, 0x3c, 0xb7, 0xe2, 0x95, 0x2b, 0x67, 0x67, 0xe3, 0x05, 0x48,
	0xa6, 0xcb, 0x33, 0x14, 0x88, 0xd1, 0x40, 0x8c, 0xee, 0xb7, 0xc4, 0xd9, 0xb4, 0xb7, 0x3f, 0x06,
	0x37, 0x39, 0x81, 0xc7, 0x46, 0x39, 0xcb, 0xf3, 0x46, 0xf3, 0x42, 0x94, 0x82, 0x9b, 0x68, 0xfc,
	0x45, 0x00, 0x79, 0x7e, 0xef, 0x62, 0xf0, 0xd1, 0xc5, 0xe0, 0xb3, 0x8b, 0x01, 0xbc, 0x16, 0x0a,
	0xf9, 0xb2, 0x42, 0x2b, 0x7f, 0xee, 0x8d, 0x1c, 0x0e, 0xbb, 0xa7, 0x3d, 0x46, 0x0a, 0x56, 0x7b,
	0x9e, 0xe7, 0xea, 0x27, 0x00, 0x00, 0xff, 0xff, 0x37, 0xbb, 0x14, 0xd2, 0xb7, 0x01, 0x00, 0x00,
}

func (m *UdpProxyConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UdpProxyConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UdpProxyConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.IdleTimeout != nil {
		{
			size, err := m.IdleTimeout.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUdpProxy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.RouteSpecifier != nil {
		{
			size := m.RouteSpecifier.Size()
			i -= size
			if _, err := m.RouteSpecifier.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.StatPrefix) > 0 {
		i -= len(m.StatPrefix)
		copy(dAtA[i:], m.StatPrefix)
		i = encodeVarintUdpProxy(dAtA, i, uint64(len(m.StatPrefix)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UdpProxyConfig_Cluster) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UdpProxyConfig_Cluster) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Cluster)
	copy(dAtA[i:], m.Cluster)
	i = encodeVarintUdpProxy(dAtA, i, uint64(len(m.Cluster)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func encodeVarintUdpProxy(dAtA []byte, offset int, v uint64) int {
	offset -= sovUdpProxy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UdpProxyConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovUdpProxy(uint64(l))
	}
	if m.RouteSpecifier != nil {
		n += m.RouteSpecifier.Size()
	}
	if m.IdleTimeout != nil {
		l = m.IdleTimeout.Size()
		n += 1 + l + sovUdpProxy(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UdpProxyConfig_Cluster) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cluster)
	n += 1 + l + sovUdpProxy(uint64(l))
	return n
}

func sovUdpProxy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUdpProxy(x uint64) (n int) {
	return sovUdpProxy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UdpProxyConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUdpProxy
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
			return fmt.Errorf("proto: UdpProxyConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UdpProxyConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUdpProxy
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
				return ErrInvalidLengthUdpProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUdpProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUdpProxy
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
				return ErrInvalidLengthUdpProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUdpProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RouteSpecifier = &UdpProxyConfig_Cluster{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdleTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUdpProxy
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
				return ErrInvalidLengthUdpProxy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUdpProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IdleTimeout == nil {
				m.IdleTimeout = &types.Duration{}
			}
			if err := m.IdleTimeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUdpProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUdpProxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUdpProxy
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
func skipUdpProxy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUdpProxy
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
					return 0, ErrIntOverflowUdpProxy
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
					return 0, ErrIntOverflowUdpProxy
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
				return 0, ErrInvalidLengthUdpProxy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUdpProxy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUdpProxy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUdpProxy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUdpProxy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUdpProxy = fmt.Errorf("proto: unexpected end of group")
)
