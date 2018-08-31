/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubeadm/vendor/k8s.io/api/storage/v1alpha1/generated.proto
// DO NOT EDIT!

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubeadm/vendor/k8s.io/api/storage/v1alpha1/generated.proto

	It has these top-level messages:
		VolumeAttachment
		VolumeAttachmentList
		VolumeAttachmentSource
		VolumeAttachmentSpec
		VolumeAttachmentStatus
		VolumeError
*/
package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

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

func (m *VolumeAttachment) Reset()                    { *m = VolumeAttachment{} }
func (*VolumeAttachment) ProtoMessage()               {}
func (*VolumeAttachment) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *VolumeAttachmentList) Reset()                    { *m = VolumeAttachmentList{} }
func (*VolumeAttachmentList) ProtoMessage()               {}
func (*VolumeAttachmentList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *VolumeAttachmentSource) Reset()                    { *m = VolumeAttachmentSource{} }
func (*VolumeAttachmentSource) ProtoMessage()               {}
func (*VolumeAttachmentSource) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *VolumeAttachmentSpec) Reset()                    { *m = VolumeAttachmentSpec{} }
func (*VolumeAttachmentSpec) ProtoMessage()               {}
func (*VolumeAttachmentSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func (m *VolumeAttachmentStatus) Reset()                    { *m = VolumeAttachmentStatus{} }
func (*VolumeAttachmentStatus) ProtoMessage()               {}
func (*VolumeAttachmentStatus) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{4} }

func (m *VolumeError) Reset()                    { *m = VolumeError{} }
func (*VolumeError) ProtoMessage()               {}
func (*VolumeError) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{5} }

func init() {
	proto.RegisterType((*VolumeAttachment)(nil), "k8s.io.api.storage.v1alpha1.VolumeAttachment")
	proto.RegisterType((*VolumeAttachmentList)(nil), "k8s.io.api.storage.v1alpha1.VolumeAttachmentList")
	proto.RegisterType((*VolumeAttachmentSource)(nil), "k8s.io.api.storage.v1alpha1.VolumeAttachmentSource")
	proto.RegisterType((*VolumeAttachmentSpec)(nil), "k8s.io.api.storage.v1alpha1.VolumeAttachmentSpec")
	proto.RegisterType((*VolumeAttachmentStatus)(nil), "k8s.io.api.storage.v1alpha1.VolumeAttachmentStatus")
	proto.RegisterType((*VolumeError)(nil), "k8s.io.api.storage.v1alpha1.VolumeError")
}
func (m *VolumeAttachment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VolumeAttachment) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
	n2, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Status.Size()))
	n3, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *VolumeAttachmentList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VolumeAttachmentList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n4, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *VolumeAttachmentSource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VolumeAttachmentSource) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.PersistentVolumeName != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.PersistentVolumeName)))
		i += copy(dAtA[i:], *m.PersistentVolumeName)
	}
	return i, nil
}

func (m *VolumeAttachmentSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VolumeAttachmentSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Attacher)))
	i += copy(dAtA[i:], m.Attacher)
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Source.Size()))
	n5, err := m.Source.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.NodeName)))
	i += copy(dAtA[i:], m.NodeName)
	return i, nil
}

func (m *VolumeAttachmentStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VolumeAttachmentStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	if m.Attached {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	if len(m.AttachmentMetadata) > 0 {
		keysForAttachmentMetadata := make([]string, 0, len(m.AttachmentMetadata))
		for k := range m.AttachmentMetadata {
			keysForAttachmentMetadata = append(keysForAttachmentMetadata, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForAttachmentMetadata)
		for _, k := range keysForAttachmentMetadata {
			dAtA[i] = 0x12
			i++
			v := m.AttachmentMetadata[string(k)]
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			i = encodeVarintGenerated(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.AttachError != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.AttachError.Size()))
		n6, err := m.AttachError.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.DetachError != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.DetachError.Size()))
		n7, err := m.DetachError.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}

func (m *VolumeError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VolumeError) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Time.Size()))
	n8, err := m.Time.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Message)))
	i += copy(dAtA[i:], m.Message)
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *VolumeAttachment) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *VolumeAttachmentList) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *VolumeAttachmentSource) Size() (n int) {
	var l int
	_ = l
	if m.PersistentVolumeName != nil {
		l = len(*m.PersistentVolumeName)
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *VolumeAttachmentSpec) Size() (n int) {
	var l int
	_ = l
	l = len(m.Attacher)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Source.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.NodeName)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *VolumeAttachmentStatus) Size() (n int) {
	var l int
	_ = l
	n += 2
	if len(m.AttachmentMetadata) > 0 {
		for k, v := range m.AttachmentMetadata {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	if m.AttachError != nil {
		l = m.AttachError.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.DetachError != nil {
		l = m.DetachError.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *VolumeError) Size() (n int) {
	var l int
	_ = l
	l = m.Time.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Message)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *VolumeAttachment) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VolumeAttachment{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "VolumeAttachmentSpec", "VolumeAttachmentSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "VolumeAttachmentStatus", "VolumeAttachmentStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VolumeAttachmentList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VolumeAttachmentList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "VolumeAttachment", "VolumeAttachment", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VolumeAttachmentSource) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VolumeAttachmentSource{`,
		`PersistentVolumeName:` + valueToStringGenerated(this.PersistentVolumeName) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VolumeAttachmentSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VolumeAttachmentSpec{`,
		`Attacher:` + fmt.Sprintf("%v", this.Attacher) + `,`,
		`Source:` + strings.Replace(strings.Replace(this.Source.String(), "VolumeAttachmentSource", "VolumeAttachmentSource", 1), `&`, ``, 1) + `,`,
		`NodeName:` + fmt.Sprintf("%v", this.NodeName) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VolumeAttachmentStatus) String() string {
	if this == nil {
		return "nil"
	}
	keysForAttachmentMetadata := make([]string, 0, len(this.AttachmentMetadata))
	for k := range this.AttachmentMetadata {
		keysForAttachmentMetadata = append(keysForAttachmentMetadata, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForAttachmentMetadata)
	mapStringForAttachmentMetadata := "map[string]string{"
	for _, k := range keysForAttachmentMetadata {
		mapStringForAttachmentMetadata += fmt.Sprintf("%v: %v,", k, this.AttachmentMetadata[k])
	}
	mapStringForAttachmentMetadata += "}"
	s := strings.Join([]string{`&VolumeAttachmentStatus{`,
		`Attached:` + fmt.Sprintf("%v", this.Attached) + `,`,
		`AttachmentMetadata:` + mapStringForAttachmentMetadata + `,`,
		`AttachError:` + strings.Replace(fmt.Sprintf("%v", this.AttachError), "VolumeError", "VolumeError", 1) + `,`,
		`DetachError:` + strings.Replace(fmt.Sprintf("%v", this.DetachError), "VolumeError", "VolumeError", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VolumeError) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VolumeError{`,
		`Time:` + strings.Replace(strings.Replace(this.Time.String(), "Time", "k8s_io_apimachinery_pkg_apis_meta_v1.Time", 1), `&`, ``, 1) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *VolumeAttachment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: VolumeAttachment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VolumeAttachment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *VolumeAttachmentList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: VolumeAttachmentList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VolumeAttachmentList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, VolumeAttachment{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *VolumeAttachmentSource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: VolumeAttachmentSource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VolumeAttachmentSource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PersistentVolumeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.PersistentVolumeName = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *VolumeAttachmentSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: VolumeAttachmentSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VolumeAttachmentSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attacher", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attacher = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Source.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *VolumeAttachmentStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: VolumeAttachmentStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VolumeAttachmentStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attached", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Attached = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthGenerated
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.AttachmentMetadata == nil {
				m.AttachmentMetadata = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthGenerated
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(dAtA[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.AttachmentMetadata[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.AttachmentMetadata[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachError", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AttachError == nil {
				m.AttachError = &VolumeError{}
			}
			if err := m.AttachError.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DetachError", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DetachError == nil {
				m.DetachError = &VolumeError{}
			}
			if err := m.DetachError.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *VolumeError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: VolumeError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VolumeError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Time.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/kubeadm/vendor/k8s.io/api/storage/v1alpha1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 704 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xc7, 0xe3, 0x24, 0x6d, 0xd3, 0xcd, 0xf3, 0x52, 0xad, 0xa2, 0xe7, 0x89, 0x82, 0xe4, 0x54,
	0x39, 0x15, 0x44, 0xd7, 0xa4, 0x20, 0x54, 0x71, 0x8b, 0xd5, 0x1e, 0x10, 0x6d, 0x41, 0x5b, 0xc4,
	0x01, 0x38, 0xb0, 0xb1, 0xa7, 0x8e, 0x9b, 0xfa, 0x45, 0xbb, 0xeb, 0x48, 0xbd, 0x71, 0xe2, 0xcc,
	0x8d, 0x6f, 0xc0, 0x67, 0xc9, 0x8d, 0x1e, 0x7b, 0x8a, 0xa8, 0xf9, 0x16, 0x5c, 0x40, 0x5e, 0x6f,
	0x5e, 0x68, 0x52, 0x68, 0x7b, 0xf3, 0xcc, 0xce, 0xfc, 0x66, 0xe6, 0xbf, 0xb3, 0x46, 0x3b, 0xfd,
	0x6d, 0x41, 0xfc, 0xc8, 0xea, 0x27, 0x5d, 0xe0, 0x21, 0x48, 0x10, 0xd6, 0x00, 0x42, 0x37, 0xe2,
	0x96, 0x3e, 0x60, 0xb1, 0x6f, 0x09, 0x19, 0x71, 0xe6, 0x81, 0x35, 0x68, 0xb3, 0x93, 0xb8, 0xc7,
	0xda, 0x96, 0x07, 0x21, 0x70, 0x26, 0xc1, 0x25, 0x31, 0x8f, 0x64, 0x84, 0xef, 0xe4, 0xc1, 0x84,
	0xc5, 0x3e, 0xd1, 0xc1, 0x64, 0x1c, 0xdc, 0xd8, 0xf4, 0x7c, 0xd9, 0x4b, 0xba, 0xc4, 0x89, 0x02,
	0xcb, 0x8b, 0xbc, 0xc8, 0x52, 0x39, 0xdd, 0xe4, 0x48, 0x59, 0xca, 0x50, 0x5f, 0x39, 0xab, 0xf1,
	0x68, 0x5a, 0x38, 0x60, 0x4e, 0xcf, 0x0f, 0x81, 0x9f, 0x5a, 0x71, 0xdf, 0xcb, 0x1c, 0xc2, 0x0a,
	0x40, 0x32, 0x6b, 0x30, 0xd7, 0x41, 0xc3, 0xba, 0x2a, 0x8b, 0x27, 0xa1, 0xf4, 0x03, 0x98, 0x4b,
	0x78, 0xfc, 0xa7, 0x04, 0xe1, 0xf4, 0x20, 0x60, 0x97, 0xf3, 0x5a, 0x9f, 0x8b, 0x68, 0xed, 0x55,
	0x74, 0x92, 0x04, 0xd0, 0x91, 0x92, 0x39, 0xbd, 0x00, 0x42, 0x89, 0xdf, 0xa1, 0x4a, 0xd6, 0x98,
	0xcb, 0x24, 0xab, 0x1b, 0xeb, 0xc6, 0x46, 0x75, 0xeb, 0x01, 0x99, 0x4a, 0x32, 0xe1, 0x93, 0xb8,
	0xef, 0x65, 0x0e, 0x41, 0xb2, 0x68, 0x32, 0x68, 0x93, 0xe7, 0xdd, 0x63, 0x70, 0xe4, 0x3e, 0x48,
	0x66, 0xe3, 0xe1, 0xa8, 0x59, 0x48, 0x47, 0x4d, 0x34, 0xf5, 0xd1, 0x09, 0x15, 0x1f, 0xa2, 0xb2,
	0x88, 0xc1, 0xa9, 0x17, 0x15, 0xbd, 0x4d, 0x7e, 0x23, 0x38, 0xb9, 0xdc, 0xde, 0x61, 0x0c, 0x8e,
	0xfd, 0x97, 0xc6, 0x97, 0x33, 0x8b, 0x2a, 0x18, 0x7e, 0x83, 0x96, 0x85, 0x64, 0x32, 0x11, 0xf5,
	0x92, 0xc2, 0x3e, 0xbc, 0x19, 0x56, 0xa5, 0xda, 0xff, 0x68, 0xf0, 0x72, 0x6e, 0x53, 0x8d, 0x6c,
	0x0d, 0x0d, 0x54, 0xbb, 0x9c, 0xb2, 0xe7, 0x0b, 0x89, 0xdf, 0xce, 0x89, 0x45, 0xae, 0x27, 0x56,
	0x96, 0xad, 0xa4, 0x5a, 0xd3, 0x25, 0x2b, 0x63, 0xcf, 0x8c, 0x50, 0x14, 0x2d, 0xf9, 0x12, 0x02,
	0x51, 0x2f, 0xae, 0x97, 0x36, 0xaa, 0x5b, 0x9b, 0x37, 0x1a, 0xc9, 0xfe, 0x5b, 0x93, 0x97, 0x9e,
	0x66, 0x0c, 0x9a, 0xa3, 0x5a, 0x47, 0xe8, 0xbf, 0xb9, 0xe1, 0xa3, 0x84, 0x3b, 0x80, 0xf7, 0x50,
	0x2d, 0x06, 0x2e, 0x7c, 0x21, 0x21, 0x94, 0x79, 0xcc, 0x01, 0x0b, 0x40, 0xcd, 0xb5, 0x6a, 0xd7,
	0xd3, 0x51, 0xb3, 0xf6, 0x62, 0xc1, 0x39, 0x5d, 0x98, 0xd5, 0xfa, 0xb2, 0x40, 0xb2, 0xec, 0xba,
	0xf0, 0x7d, 0x54, 0x61, 0xca, 0x03, 0x5c, 0xa3, 0x27, 0x12, 0x74, 0xb4, 0x9f, 0x4e, 0x22, 0xd4,
	0xb5, 0xaa, 0xf6, 0xf4, 0xb6, 0xdc, 0xf0, 0x5a, 0x55, 0xea, 0xcc, 0xb5, 0x2a, 0x9b, 0x6a, 0x64,
	0xd6, 0x4a, 0x18, 0xb9, 0xf9, 0x94, 0xa5, 0x5f, 0x5b, 0x39, 0xd0, 0x7e, 0x3a, 0x89, 0x68, 0xfd,
	0x28, 0x2d, 0x90, 0x4e, 0xed, 0xc7, 0xcc, 0x4c, 0xae, 0x9a, 0xa9, 0x32, 0x37, 0x93, 0x3b, 0x99,
	0xc9, 0xc5, 0x9f, 0x0c, 0x84, 0xd9, 0x04, 0xb1, 0x3f, 0xde, 0x9f, 0xfc, 0x92, 0x9f, 0xdd, 0x62,
	0x6f, 0x49, 0x67, 0x8e, 0xb6, 0x1b, 0x4a, 0x7e, 0x6a, 0x37, 0x74, 0x17, 0x78, 0x3e, 0x80, 0x2e,
	0x68, 0x01, 0x1f, 0xa3, 0x6a, 0xee, 0xdd, 0xe5, 0x3c, 0xe2, 0xfa, 0x25, 0x6d, 0x5c, 0xa3, 0x23,
	0x15, 0x6f, 0x9b, 0xe9, 0xa8, 0x59, 0xed, 0x4c, 0x01, 0xdf, 0x47, 0xcd, 0xea, 0xcc, 0x39, 0x9d,
	0x85, 0x67, 0xb5, 0x5c, 0x98, 0xd6, 0x2a, 0xdf, 0xa6, 0xd6, 0x0e, 0x5c, 0x5d, 0x6b, 0x06, 0xde,
	0xd8, 0x45, 0xff, 0x5f, 0x21, 0x11, 0x5e, 0x43, 0xa5, 0x3e, 0x9c, 0xe6, 0x9b, 0x48, 0xb3, 0x4f,
	0x5c, 0x43, 0x4b, 0x03, 0x76, 0x92, 0xe4, 0x1b, 0xb7, 0x4a, 0x73, 0xe3, 0x49, 0x71, 0xdb, 0x68,
	0x7d, 0x30, 0xd0, 0x6c, 0x0d, 0xbc, 0x87, 0xca, 0xd9, 0xef, 0x55, 0xbf, 0xfc, 0x7b, 0xd7, 0x7b,
	0xf9, 0x2f, 0xfd, 0x00, 0xa6, 0x7f, 0xb0, 0xcc, 0xa2, 0x8a, 0x82, 0xef, 0xa2, 0x95, 0x00, 0x84,
	0x60, 0x9e, 0xae, 0x6c, 0xff, 0xab, 0x83, 0x56, 0xf6, 0x73, 0x37, 0x1d, 0x9f, 0xdb, 0x64, 0x78,
	0x61, 0x16, 0xce, 0x2e, 0xcc, 0xc2, 0xf9, 0x85, 0x59, 0x78, 0x9f, 0x9a, 0xc6, 0x30, 0x35, 0x8d,
	0xb3, 0xd4, 0x34, 0xce, 0x53, 0xd3, 0xf8, 0x9a, 0x9a, 0xc6, 0xc7, 0x6f, 0x66, 0xe1, 0x75, 0x65,
	0x2c, 0xdc, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0xba, 0xdb, 0x12, 0x1a, 0x07, 0x00, 0x00,
}
