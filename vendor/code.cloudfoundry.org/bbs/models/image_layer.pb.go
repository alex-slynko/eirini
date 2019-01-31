// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: image_layer.proto

package models

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strconv "strconv"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ImageLayer_DigestAlgorithm int32

const (
	DigestAlgorithmInvalid ImageLayer_DigestAlgorithm = 0
	DigestAlgorithmSha256  ImageLayer_DigestAlgorithm = 1
	DigestAlgorithmSha512  ImageLayer_DigestAlgorithm = 2
)

var ImageLayer_DigestAlgorithm_name = map[int32]string{
	0: "DigestAlgorithmInvalid",
	1: "SHA256",
	2: "SHA512",
}
var ImageLayer_DigestAlgorithm_value = map[string]int32{
	"DigestAlgorithmInvalid": 0,
	"SHA256":                 1,
	"SHA512":                 2,
}

func (x ImageLayer_DigestAlgorithm) Enum() *ImageLayer_DigestAlgorithm {
	p := new(ImageLayer_DigestAlgorithm)
	*p = x
	return p
}
func (x ImageLayer_DigestAlgorithm) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(ImageLayer_DigestAlgorithm_name, int32(x))
}
func (x *ImageLayer_DigestAlgorithm) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ImageLayer_DigestAlgorithm_value, data, "ImageLayer_DigestAlgorithm")
	if err != nil {
		return err
	}
	*x = ImageLayer_DigestAlgorithm(value)
	return nil
}
func (ImageLayer_DigestAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_image_layer_c20242bb5a301fe7, []int{0, 0}
}

type ImageLayer_MediaType int32

const (
	MediaTypeInvalid ImageLayer_MediaType = 0
	MediaTypeTgz     ImageLayer_MediaType = 1
	MediaTypeTar     ImageLayer_MediaType = 2
	MediaTypeZip     ImageLayer_MediaType = 3
)

var ImageLayer_MediaType_name = map[int32]string{
	0: "MediaTypeInvalid",
	1: "TGZ",
	2: "TAR",
	3: "ZIP",
}
var ImageLayer_MediaType_value = map[string]int32{
	"MediaTypeInvalid": 0,
	"TGZ":              1,
	"TAR":              2,
	"ZIP":              3,
}

func (x ImageLayer_MediaType) Enum() *ImageLayer_MediaType {
	p := new(ImageLayer_MediaType)
	*p = x
	return p
}
func (x ImageLayer_MediaType) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(ImageLayer_MediaType_name, int32(x))
}
func (x *ImageLayer_MediaType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ImageLayer_MediaType_value, data, "ImageLayer_MediaType")
	if err != nil {
		return err
	}
	*x = ImageLayer_MediaType(value)
	return nil
}
func (ImageLayer_MediaType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_image_layer_c20242bb5a301fe7, []int{0, 1}
}

type ImageLayer_Type int32

const (
	LayerTypeInvalid   ImageLayer_Type = 0
	LayerTypeShared    ImageLayer_Type = 1
	LayerTypeExclusive ImageLayer_Type = 2
)

var ImageLayer_Type_name = map[int32]string{
	0: "LayerTypeInvalid",
	1: "SHARED",
	2: "EXCLUSIVE",
}
var ImageLayer_Type_value = map[string]int32{
	"LayerTypeInvalid": 0,
	"SHARED":           1,
	"EXCLUSIVE":        2,
}

func (x ImageLayer_Type) Enum() *ImageLayer_Type {
	p := new(ImageLayer_Type)
	*p = x
	return p
}
func (x ImageLayer_Type) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(ImageLayer_Type_name, int32(x))
}
func (x *ImageLayer_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ImageLayer_Type_value, data, "ImageLayer_Type")
	if err != nil {
		return err
	}
	*x = ImageLayer_Type(value)
	return nil
}
func (ImageLayer_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_image_layer_c20242bb5a301fe7, []int{0, 2}
}

type ImageLayer struct {
	Name            string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Url             string                     `protobuf:"bytes,2,opt,name=url" json:"url"`
	DestinationPath string                     `protobuf:"bytes,3,opt,name=destination_path,json=destinationPath" json:"destination_path"`
	LayerType       ImageLayer_Type            `protobuf:"varint,4,opt,name=layer_type,json=layerType,enum=models.ImageLayer_Type" json:"layer_type"`
	MediaType       ImageLayer_MediaType       `protobuf:"varint,5,opt,name=media_type,json=mediaType,enum=models.ImageLayer_MediaType" json:"media_type"`
	DigestAlgorithm ImageLayer_DigestAlgorithm `protobuf:"varint,6,opt,name=digest_algorithm,json=digestAlgorithm,enum=models.ImageLayer_DigestAlgorithm" json:"digest_algorithm,omitempty"`
	DigestValue     string                     `protobuf:"bytes,7,opt,name=digest_value,json=digestValue" json:"digest_value,omitempty"`
}

func (m *ImageLayer) Reset()      { *m = ImageLayer{} }
func (*ImageLayer) ProtoMessage() {}
func (*ImageLayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_layer_c20242bb5a301fe7, []int{0}
}
func (m *ImageLayer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ImageLayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ImageLayer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ImageLayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageLayer.Merge(dst, src)
}
func (m *ImageLayer) XXX_Size() int {
	return m.Size()
}
func (m *ImageLayer) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageLayer.DiscardUnknown(m)
}

var xxx_messageInfo_ImageLayer proto.InternalMessageInfo

func (m *ImageLayer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ImageLayer) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ImageLayer) GetDestinationPath() string {
	if m != nil {
		return m.DestinationPath
	}
	return ""
}

func (m *ImageLayer) GetLayerType() ImageLayer_Type {
	if m != nil {
		return m.LayerType
	}
	return LayerTypeInvalid
}

func (m *ImageLayer) GetMediaType() ImageLayer_MediaType {
	if m != nil {
		return m.MediaType
	}
	return MediaTypeInvalid
}

func (m *ImageLayer) GetDigestAlgorithm() ImageLayer_DigestAlgorithm {
	if m != nil {
		return m.DigestAlgorithm
	}
	return DigestAlgorithmInvalid
}

func (m *ImageLayer) GetDigestValue() string {
	if m != nil {
		return m.DigestValue
	}
	return ""
}

func init() {
	proto.RegisterType((*ImageLayer)(nil), "models.ImageLayer")
	proto.RegisterEnum("models.ImageLayer_DigestAlgorithm", ImageLayer_DigestAlgorithm_name, ImageLayer_DigestAlgorithm_value)
	proto.RegisterEnum("models.ImageLayer_MediaType", ImageLayer_MediaType_name, ImageLayer_MediaType_value)
	proto.RegisterEnum("models.ImageLayer_Type", ImageLayer_Type_name, ImageLayer_Type_value)
}
func (x ImageLayer_DigestAlgorithm) String() string {
	s, ok := ImageLayer_DigestAlgorithm_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ImageLayer_MediaType) String() string {
	s, ok := ImageLayer_MediaType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ImageLayer_Type) String() string {
	s, ok := ImageLayer_Type_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *ImageLayer) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ImageLayer)
	if !ok {
		that2, ok := that.(ImageLayer)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Url != that1.Url {
		return false
	}
	if this.DestinationPath != that1.DestinationPath {
		return false
	}
	if this.LayerType != that1.LayerType {
		return false
	}
	if this.MediaType != that1.MediaType {
		return false
	}
	if this.DigestAlgorithm != that1.DigestAlgorithm {
		return false
	}
	if this.DigestValue != that1.DigestValue {
		return false
	}
	return true
}
func (this *ImageLayer) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&models.ImageLayer{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Url: "+fmt.Sprintf("%#v", this.Url)+",\n")
	s = append(s, "DestinationPath: "+fmt.Sprintf("%#v", this.DestinationPath)+",\n")
	s = append(s, "LayerType: "+fmt.Sprintf("%#v", this.LayerType)+",\n")
	s = append(s, "MediaType: "+fmt.Sprintf("%#v", this.MediaType)+",\n")
	s = append(s, "DigestAlgorithm: "+fmt.Sprintf("%#v", this.DigestAlgorithm)+",\n")
	s = append(s, "DigestValue: "+fmt.Sprintf("%#v", this.DigestValue)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringImageLayer(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ImageLayer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ImageLayer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintImageLayer(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	dAtA[i] = 0x12
	i++
	i = encodeVarintImageLayer(dAtA, i, uint64(len(m.Url)))
	i += copy(dAtA[i:], m.Url)
	dAtA[i] = 0x1a
	i++
	i = encodeVarintImageLayer(dAtA, i, uint64(len(m.DestinationPath)))
	i += copy(dAtA[i:], m.DestinationPath)
	dAtA[i] = 0x20
	i++
	i = encodeVarintImageLayer(dAtA, i, uint64(m.LayerType))
	dAtA[i] = 0x28
	i++
	i = encodeVarintImageLayer(dAtA, i, uint64(m.MediaType))
	dAtA[i] = 0x30
	i++
	i = encodeVarintImageLayer(dAtA, i, uint64(m.DigestAlgorithm))
	dAtA[i] = 0x3a
	i++
	i = encodeVarintImageLayer(dAtA, i, uint64(len(m.DigestValue)))
	i += copy(dAtA[i:], m.DigestValue)
	return i, nil
}

func encodeVarintImageLayer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ImageLayer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovImageLayer(uint64(l))
	l = len(m.Url)
	n += 1 + l + sovImageLayer(uint64(l))
	l = len(m.DestinationPath)
	n += 1 + l + sovImageLayer(uint64(l))
	n += 1 + sovImageLayer(uint64(m.LayerType))
	n += 1 + sovImageLayer(uint64(m.MediaType))
	n += 1 + sovImageLayer(uint64(m.DigestAlgorithm))
	l = len(m.DigestValue)
	n += 1 + l + sovImageLayer(uint64(l))
	return n
}

func sovImageLayer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozImageLayer(x uint64) (n int) {
	return sovImageLayer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ImageLayer) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ImageLayer{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Url:` + fmt.Sprintf("%v", this.Url) + `,`,
		`DestinationPath:` + fmt.Sprintf("%v", this.DestinationPath) + `,`,
		`LayerType:` + fmt.Sprintf("%v", this.LayerType) + `,`,
		`MediaType:` + fmt.Sprintf("%v", this.MediaType) + `,`,
		`DigestAlgorithm:` + fmt.Sprintf("%v", this.DigestAlgorithm) + `,`,
		`DigestValue:` + fmt.Sprintf("%v", this.DigestValue) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringImageLayer(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ImageLayer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImageLayer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ImageLayer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImageLayer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageLayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthImageLayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageLayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthImageLayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageLayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthImageLayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LayerType", wireType)
			}
			m.LayerType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageLayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LayerType |= (ImageLayer_Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MediaType", wireType)
			}
			m.MediaType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageLayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MediaType |= (ImageLayer_MediaType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DigestAlgorithm", wireType)
			}
			m.DigestAlgorithm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageLayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DigestAlgorithm |= (ImageLayer_DigestAlgorithm(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DigestValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageLayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthImageLayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DigestValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImageLayer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImageLayer
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
func skipImageLayer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowImageLayer
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
					return 0, ErrIntOverflowImageLayer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowImageLayer
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthImageLayer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowImageLayer
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipImageLayer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthImageLayer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowImageLayer   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("image_layer.proto", fileDescriptor_image_layer_c20242bb5a301fe7) }

var fileDescriptor_image_layer_c20242bb5a301fe7 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x8f, 0xd2, 0x40,
	0x18, 0x87, 0x3b, 0x80, 0x18, 0xc6, 0xcd, 0x32, 0x8e, 0x8a, 0xdd, 0xc6, 0x0c, 0x4d, 0x93, 0x4d,
	0x36, 0x46, 0x21, 0x4b, 0x64, 0xef, 0xe0, 0x12, 0xc5, 0x60, 0xdc, 0x14, 0xdc, 0x18, 0x2e, 0x64,
	0x76, 0x3b, 0xb6, 0x4d, 0x5a, 0x5a, 0xcb, 0x40, 0x44, 0x2f, 0x9e, 0x39, 0xf9, 0x05, 0xb8, 0x9b,
	0xf8, 0x45, 0xf6, 0xc8, 0x71, 0x4f, 0x1b, 0x29, 0x17, 0xe3, 0x69, 0x3f, 0x82, 0xe9, 0xf0, 0xaf,
	0xdb, 0xec, 0xde, 0x3a, 0xef, 0xef, 0x79, 0x9f, 0x79, 0xe7, 0x4d, 0xe1, 0x43, 0xdb, 0xa5, 0x26,
	0xeb, 0x39, 0x74, 0xcc, 0x82, 0x92, 0x1f, 0x78, 0xdc, 0xc3, 0x59, 0xd7, 0x33, 0x98, 0x33, 0x50,
	0x5e, 0x9a, 0x36, 0xb7, 0x86, 0x67, 0xa5, 0x73, 0xcf, 0x2d, 0x9b, 0x9e, 0xe9, 0x95, 0x45, 0x7c,
	0x36, 0xfc, 0x2c, 0x4e, 0xe2, 0x20, 0xbe, 0x96, 0x6d, 0xda, 0xef, 0x2c, 0x84, 0xcd, 0x48, 0xd6,
	0x8a, 0x5c, 0xf8, 0x39, 0xcc, 0xf4, 0xa9, 0xcb, 0x64, 0xa0, 0x82, 0x83, 0x5c, 0xbd, 0x70, 0x71,
	0x55, 0x94, 0xfe, 0x5d, 0x15, 0x77, 0xa3, 0xda, 0x0b, 0xcf, 0xb5, 0x39, 0x73, 0x7d, 0x3e, 0xd6,
	0x05, 0x83, 0x0b, 0x30, 0x3d, 0x0c, 0x1c, 0x39, 0x25, 0xd0, 0x4c, 0x84, 0xea, 0x51, 0x01, 0x97,
	0x21, 0x32, 0xd8, 0x80, 0xdb, 0x7d, 0xca, 0x6d, 0xaf, 0xdf, 0xf3, 0x29, 0xb7, 0xe4, 0x74, 0x0c,
	0xca, 0xc7, 0xd2, 0x13, 0xca, 0x2d, 0xfc, 0x0e, 0x42, 0xf1, 0x92, 0x1e, 0x1f, 0xfb, 0x4c, 0xce,
	0xa8, 0xe0, 0x60, 0xb7, 0xf2, 0xb4, 0xb4, 0x7c, 0x4f, 0x69, 0x3b, 0x5c, 0xa9, 0x33, 0xf6, 0x59,
	0x1d, 0xaf, 0x66, 0x8a, 0xb5, 0xe8, 0x39, 0xf1, 0x1d, 0xc5, 0xf8, 0x03, 0x84, 0x2e, 0x33, 0x6c,
	0xba, 0x74, 0xdd, 0x13, 0xae, 0x67, 0xb7, 0xb8, 0xde, 0x47, 0xd0, 0x4d, 0xe1, 0xb6, 0x4f, 0xcf,
	0xb9, 0xeb, 0x18, 0x7f, 0x81, 0xc8, 0xb0, 0x4d, 0x36, 0xe0, 0x3d, 0xea, 0x98, 0x5e, 0x60, 0x73,
	0xcb, 0x95, 0xb3, 0x42, 0xab, 0xdd, 0xa2, 0x3d, 0x16, 0x68, 0x6d, 0x4d, 0xd6, 0xb5, 0x95, 0x5c,
	0x49, 0x3a, 0x62, 0xdb, 0xcc, 0x1b, 0x37, 0x9b, 0x70, 0x0d, 0xee, 0xac, 0xf0, 0x11, 0x75, 0x86,
	0x4c, 0xbe, 0x2f, 0x96, 0x47, 0x56, 0xaa, 0x42, 0x3c, 0x8b, 0x69, 0x1e, 0x2c, 0xeb, 0xa7, 0x51,
	0x59, 0xfb, 0x0e, 0xf3, 0x89, 0x51, 0xb0, 0x02, 0x0b, 0x89, 0x52, 0xb3, 0x3f, 0xa2, 0x8e, 0x6d,
	0x20, 0x09, 0xef, 0xc3, 0x6c, 0xfb, 0x6d, 0xad, 0x52, 0x3d, 0x42, 0x40, 0xd9, 0x9b, 0x4c, 0xd5,
	0x27, 0x09, 0xb2, 0x6d, 0xd1, 0x4a, 0xf5, 0x68, 0x85, 0x55, 0x0f, 0x2b, 0x28, 0x75, 0x17, 0x56,
	0x3d, 0xac, 0x68, 0x01, 0xcc, 0x6d, 0xd6, 0x8b, 0x1f, 0x43, 0xb4, 0x39, 0x6c, 0x2f, 0xdc, 0x83,
	0xe9, 0xce, 0x9b, 0x2e, 0x02, 0x0a, 0x9a, 0x4c, 0xd5, 0x9d, 0x0d, 0xd0, 0x31, 0xbf, 0x89, 0xa8,
	0xa6, 0xa3, 0x54, 0x32, 0xa2, 0x41, 0x14, 0x75, 0x9b, 0x27, 0x28, 0x9d, 0x88, 0xba, 0xb6, 0xaf,
	0x19, 0x30, 0xb3, 0xbe, 0xae, 0xb5, 0xfe, 0x19, 0xb6, 0xd7, 0x15, 0xc5, 0xe0, 0x7a, 0xe3, 0x18,
	0x01, 0xe5, 0xd1, 0x64, 0xaa, 0xe6, 0x37, 0x4c, 0xdb, 0xa2, 0x01, 0x33, 0xf0, 0x3e, 0xcc, 0x35,
	0x3e, 0xbd, 0x6e, 0x7d, 0x6c, 0x37, 0x4f, 0x1b, 0x28, 0xa5, 0x14, 0x26, 0x53, 0x15, 0x6f, 0x98,
	0xc6, 0xd7, 0x73, 0x67, 0x38, 0xb0, 0x47, 0xac, 0xfe, 0x6a, 0x36, 0x27, 0xd2, 0xe5, 0x9c, 0x48,
	0xd7, 0x73, 0x02, 0x7e, 0x84, 0x04, 0xfc, 0x0a, 0x09, 0xb8, 0x08, 0x09, 0x98, 0x85, 0x04, 0xfc,
	0x09, 0x09, 0xf8, 0x1b, 0x12, 0xe9, 0x3a, 0x24, 0xe0, 0xe7, 0x82, 0x48, 0xb3, 0x05, 0x91, 0x2e,
	0x17, 0x44, 0xfa, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x58, 0x37, 0x37, 0xae, 0x03, 0x00, 0x00,
}