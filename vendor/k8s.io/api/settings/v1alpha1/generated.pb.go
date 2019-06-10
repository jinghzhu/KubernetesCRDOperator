/*
Copyright 2019 The Kubernetes Authors.

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
// source: k8s.io/kubernetes/vendor/k8s.io/api/settings/v1alpha1/generated.proto
// DO NOT EDIT!

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/api/settings/v1alpha1/generated.proto

	It has these top-level messages:
		PodPreset
		PodPresetList
		PodPresetSpec
*/
package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_api_core_v1 "k8s.io/api/core/v1"

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

func (m *PodPreset) Reset()                    { *m = PodPreset{} }
func (*PodPreset) ProtoMessage()               {}
func (*PodPreset) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *PodPresetList) Reset()                    { *m = PodPresetList{} }
func (*PodPresetList) ProtoMessage()               {}
func (*PodPresetList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *PodPresetSpec) Reset()                    { *m = PodPresetSpec{} }
func (*PodPresetSpec) ProtoMessage()               {}
func (*PodPresetSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func init() {
	proto.RegisterType((*PodPreset)(nil), "k8s.io.api.settings.v1alpha1.PodPreset")
	proto.RegisterType((*PodPresetList)(nil), "k8s.io.api.settings.v1alpha1.PodPresetList")
	proto.RegisterType((*PodPresetSpec)(nil), "k8s.io.api.settings.v1alpha1.PodPresetSpec")
}
func (m *PodPreset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodPreset) MarshalTo(dAtA []byte) (int, error) {
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
	return i, nil
}

func (m *PodPresetList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodPresetList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n3, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
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

func (m *PodPresetSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodPresetSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Selector.Size()))
	n4, err := m.Selector.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.Env) > 0 {
		for _, msg := range m.Env {
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
	if len(m.EnvFrom) > 0 {
		for _, msg := range m.EnvFrom {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Volumes) > 0 {
		for _, msg := range m.Volumes {
			dAtA[i] = 0x22
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.VolumeMounts) > 0 {
		for _, msg := range m.VolumeMounts {
			dAtA[i] = 0x2a
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
func (m *PodPreset) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *PodPresetList) Size() (n int) {
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

func (m *PodPresetSpec) Size() (n int) {
	var l int
	_ = l
	l = m.Selector.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Env) > 0 {
		for _, e := range m.Env {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.EnvFrom) > 0 {
		for _, e := range m.EnvFrom {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.Volumes) > 0 {
		for _, e := range m.Volumes {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.VolumeMounts) > 0 {
		for _, e := range m.VolumeMounts {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
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
func (this *PodPreset) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PodPreset{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "PodPresetSpec", "PodPresetSpec", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PodPresetList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PodPresetList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "PodPreset", "PodPreset", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PodPresetSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PodPresetSpec{`,
		`Selector:` + strings.Replace(strings.Replace(this.Selector.String(), "LabelSelector", "k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector", 1), `&`, ``, 1) + `,`,
		`Env:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Env), "EnvVar", "k8s_io_api_core_v1.EnvVar", 1), `&`, ``, 1) + `,`,
		`EnvFrom:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.EnvFrom), "EnvFromSource", "k8s_io_api_core_v1.EnvFromSource", 1), `&`, ``, 1) + `,`,
		`Volumes:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Volumes), "Volume", "k8s_io_api_core_v1.Volume", 1), `&`, ``, 1) + `,`,
		`VolumeMounts:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.VolumeMounts), "VolumeMount", "k8s_io_api_core_v1.VolumeMount", 1), `&`, ``, 1) + `,`,
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
func (m *PodPreset) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PodPreset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodPreset: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *PodPresetList) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PodPresetList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodPresetList: illegal tag %d (wire type %d)", fieldNum, wire)
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
			m.Items = append(m.Items, PodPreset{})
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
func (m *PodPresetSpec) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PodPresetSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodPresetSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
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
			if err := m.Selector.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Env", wireType)
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
			m.Env = append(m.Env, k8s_io_api_core_v1.EnvVar{})
			if err := m.Env[len(m.Env)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnvFrom", wireType)
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
			m.EnvFrom = append(m.EnvFrom, k8s_io_api_core_v1.EnvFromSource{})
			if err := m.EnvFrom[len(m.EnvFrom)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Volumes", wireType)
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
			m.Volumes = append(m.Volumes, k8s_io_api_core_v1.Volume{})
			if err := m.Volumes[len(m.Volumes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VolumeMounts", wireType)
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
			m.VolumeMounts = append(m.VolumeMounts, k8s_io_api_core_v1.VolumeMount{})
			if err := m.VolumeMounts[len(m.VolumeMounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/settings/v1alpha1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x8e, 0xd2, 0x4e,
	0x1c, 0xc7, 0xe9, 0xb2, 0x04, 0xfe, 0x03, 0x9b, 0xbf, 0x69, 0x3c, 0x34, 0xc4, 0x94, 0x95, 0x8b,
	0x9b, 0x18, 0x67, 0x64, 0xd7, 0x18, 0xbd, 0x36, 0xc1, 0xc4, 0x04, 0xe2, 0xa6, 0x24, 0x9b, 0x68,
	0x3c, 0x38, 0x94, 0x9f, 0xa5, 0x42, 0x67, 0x9a, 0x99, 0x69, 0x13, 0x6f, 0x3e, 0x82, 0x2f, 0xe0,
	0x93, 0xe8, 0x03, 0x70, 0xdc, 0xe3, 0x9e, 0x36, 0x52, 0x5f, 0xc4, 0x4c, 0x99, 0x02, 0x8a, 0x28,
	0xb7, 0xce, 0x8f, 0xef, 0xe7, 0x33, 0xbf, 0x2f, 0x2d, 0xea, 0xcf, 0x9e, 0x49, 0x1c, 0x71, 0x32,
	0x4b, 0xc7, 0x20, 0x18, 0x28, 0x90, 0x24, 0x03, 0x36, 0xe1, 0x82, 0x98, 0x1f, 0x68, 0x12, 0x11,
	0x09, 0x4a, 0x45, 0x2c, 0x94, 0x24, 0xeb, 0xd1, 0x79, 0x32, 0xa5, 0x3d, 0x12, 0x02, 0x03, 0x41,
	0x15, 0x4c, 0x70, 0x22, 0xb8, 0xe2, 0xf6, 0xbd, 0x55, 0x1a, 0xd3, 0x24, 0xc2, 0x65, 0x1a, 0x97,
	0xe9, 0xf6, 0xa3, 0x30, 0x52, 0xd3, 0x74, 0x8c, 0x03, 0x1e, 0x93, 0x90, 0x87, 0x9c, 0x14, 0xd0,
	0x38, 0x7d, 0x5f, 0x9c, 0x8a, 0x43, 0xf1, 0xb4, 0x92, 0xb5, 0xbb, 0x5b, 0x57, 0x07, 0x5c, 0x00,
	0xc9, 0x76, 0x2e, 0x6c, 0x3f, 0xd9, 0x64, 0x62, 0x1a, 0x4c, 0x23, 0x06, 0xe2, 0x23, 0x49, 0x66,
	0xa1, 0x1e, 0x48, 0x12, 0x83, 0xa2, 0x7f, 0xa2, 0xc8, 0x3e, 0x4a, 0xa4, 0x4c, 0x45, 0x31, 0xec,
	0x00, 0x4f, 0xff, 0x05, 0xc8, 0x60, 0x0a, 0x31, 0xdd, 0xe1, 0x2e, 0xf6, 0x71, 0xa9, 0x8a, 0xe6,
	0x24, 0x62, 0x4a, 0x2a, 0xf1, 0x3b, 0xd4, 0xfd, 0x66, 0xa1, 0xff, 0x2e, 0xf9, 0xe4, 0x52, 0x80,
	0x04, 0x65, 0xbf, 0x43, 0x0d, 0x5d, 0x63, 0x42, 0x15, 0x75, 0xac, 0x53, 0xeb, 0xac, 0x79, 0xfe,
	0x18, 0x6f, 0xfe, 0xe5, 0xb5, 0x15, 0x27, 0xb3, 0x50, 0x0f, 0x24, 0xd6, 0x69, 0x9c, 0xf5, 0xf0,
	0xab, 0xf1, 0x07, 0x08, 0xd4, 0x10, 0x14, 0xf5, 0xec, 0xc5, 0x6d, 0xa7, 0x92, 0xdf, 0x76, 0xd0,
	0x66, 0xe6, 0xaf, 0xad, 0xf6, 0x10, 0x1d, 0xcb, 0x04, 0x02, 0xe7, 0xa8, 0xb0, 0x3f, 0xc4, 0x7f,
	0x7b, 0x87, 0x78, 0xbd, 0xd8, 0x28, 0x81, 0xc0, 0x6b, 0x19, 0xf1, 0xb1, 0x3e, 0xf9, 0x85, 0xa6,
	0xfb, 0xd5, 0x42, 0x27, 0xeb, 0xd4, 0x20, 0x92, 0xca, 0x7e, 0xbb, 0x53, 0x01, 0x1f, 0x56, 0x41,
	0xd3, 0x45, 0x81, 0x3b, 0xe6, 0x9e, 0x46, 0x39, 0xd9, 0x5a, 0x7f, 0x80, 0x6a, 0x91, 0x82, 0x58,
	0x3a, 0x47, 0xa7, 0xd5, 0xb3, 0xe6, 0xf9, 0x83, 0x03, 0xf7, 0xf7, 0x4e, 0x8c, 0xb3, 0xf6, 0x52,
	0xd3, 0xfe, 0x4a, 0xd2, 0xfd, 0x52, 0xdd, 0xda, 0x5e, 0xb7, 0xb2, 0x29, 0x6a, 0x48, 0x98, 0x43,
	0xa0, 0xb8, 0x30, 0xdb, 0x5f, 0x1c, 0xb8, 0x3d, 0x1d, 0xc3, 0x7c, 0x64, 0xd0, 0x4d, 0x85, 0x72,
	0xe2, 0xaf, 0xb5, 0xf6, 0x73, 0x54, 0x05, 0x96, 0x99, 0x02, 0xed, 0xed, 0x02, 0xfa, 0xbb, 0xd7,
	0xae, 0x3e, 0xcb, 0xae, 0xa8, 0xf0, 0x9a, 0x46, 0x52, 0xed, 0xb3, 0xcc, 0xd7, 0x8c, 0x3d, 0x40,
	0x75, 0x60, 0xd9, 0x0b, 0xc1, 0x63, 0xa7, 0x5a, 0xe0, 0xf7, 0xf7, 0xe0, 0x3a, 0x32, 0xe2, 0xa9,
	0x08, 0xc0, 0xfb, 0xdf, 0x58, 0xea, 0x66, 0xec, 0x97, 0x0a, 0xbb, 0x8f, 0xea, 0x19, 0x9f, 0xa7,
	0x31, 0x48, 0xe7, 0x78, 0xff, 0x32, 0x57, 0x45, 0x64, 0xa3, 0x59, 0x9d, 0xa5, 0x5f, 0xb2, 0xf6,
	0x6b, 0xd4, 0x5a, 0x3d, 0x0e, 0x79, 0xca, 0x94, 0x74, 0x6a, 0x85, 0xab, 0xb3, 0xdf, 0x55, 0xe4,
	0xbc, 0xbb, 0x46, 0xd8, 0xda, 0x1a, 0x4a, 0xff, 0x17, 0x95, 0x87, 0x17, 0x4b, 0xb7, 0x72, 0xbd,
	0x74, 0x2b, 0x37, 0x4b, 0xb7, 0xf2, 0x29, 0x77, 0xad, 0x45, 0xee, 0x5a, 0xd7, 0xb9, 0x6b, 0xdd,
	0xe4, 0xae, 0xf5, 0x3d, 0x77, 0xad, 0xcf, 0x3f, 0xdc, 0xca, 0x9b, 0x46, 0xf9, 0xbe, 0x7f, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x0e, 0x35, 0x98, 0xb5, 0xd9, 0x04, 0x00, 0x00,
}
