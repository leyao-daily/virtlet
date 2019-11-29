// +build !ignore_autogenerated

/*
Copyright 2019 Mirantis

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageTranslation) DeepCopyInto(out *ImageTranslation) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]TranslationRule, len(*in))
		copy(*out, *in)
	}
	if in.Transports != nil {
		in, out := &in.Transports, &out.Transports
		*out = make(map[string]TransportProfile, len(*in))
		for key, val := range *in {
			newVal := new(TransportProfile)
			val.DeepCopyInto(newVal)
			(*out)[key] = *newVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageTranslation.
func (in *ImageTranslation) DeepCopy() *ImageTranslation {
	if in == nil {
		return nil
	}
	out := new(ImageTranslation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSCertificate) DeepCopyInto(out *TLSCertificate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSCertificate.
func (in *TLSCertificate) DeepCopy() *TLSCertificate {
	if in == nil {
		return nil
	}
	out := new(TLSCertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	if in.Certificates != nil {
		in, out := &in.Certificates, &out.Certificates
		*out = make([]TLSCertificate, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TranslationRule) DeepCopyInto(out *TranslationRule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TranslationRule.
func (in *TranslationRule) DeepCopy() *TranslationRule {
	if in == nil {
		return nil
	}
	out := new(TranslationRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransportProfile) DeepCopyInto(out *TransportProfile) {
	*out = *in
	if in.MaxRedirects != nil {
		in, out := &in.MaxRedirects, &out.MaxRedirects
		if *in == nil {
			*out = nil
		} else {
			*out = new(int)
			**out = **in
		}
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		if *in == nil {
			*out = nil
		} else {
			*out = new(TLSConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransportProfile.
func (in *TransportProfile) DeepCopy() *TransportProfile {
	if in == nil {
		return nil
	}
	out := new(TransportProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtletConfig) DeepCopyInto(out *VirtletConfig) {
	*out = *in
	if in.FDServerSocketPath != nil {
		in, out := &in.FDServerSocketPath, &out.FDServerSocketPath
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.DatabasePath != nil {
		in, out := &in.DatabasePath, &out.DatabasePath
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.DownloadProtocol != nil {
		in, out := &in.DownloadProtocol, &out.DownloadProtocol
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.ImageDir != nil {
		in, out := &in.ImageDir, &out.ImageDir
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.ImageTranslationConfigsDir != nil {
		in, out := &in.ImageTranslationConfigsDir, &out.ImageTranslationConfigsDir
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.SkipImageTranslation != nil {
		in, out := &in.SkipImageTranslation, &out.SkipImageTranslation
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.LibvirtURI != nil {
		in, out := &in.LibvirtURI, &out.LibvirtURI
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.RawDevices != nil {
		in, out := &in.RawDevices, &out.RawDevices
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.CRISocketPath != nil {
		in, out := &in.CRISocketPath, &out.CRISocketPath
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.DisableLogging != nil {
		in, out := &in.DisableLogging, &out.DisableLogging
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.DisableKVM != nil {
		in, out := &in.DisableKVM, &out.DisableKVM
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.EnableSriov != nil {
		in, out := &in.EnableSriov, &out.EnableSriov
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.CNIPluginDir != nil {
		in, out := &in.CNIPluginDir, &out.CNIPluginDir
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.CNIConfigDir != nil {
		in, out := &in.CNIConfigDir, &out.CNIConfigDir
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.EnableRegexpImageTranslation != nil {
		in, out := &in.EnableRegexpImageTranslation, &out.EnableRegexpImageTranslation
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.CPUModel != nil {
		in, out := &in.CPUModel, &out.CPUModel
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.StreamPort != nil {
		in, out := &in.StreamPort, &out.StreamPort
		if *in == nil {
			*out = nil
		} else {
			*out = new(int)
			**out = **in
		}
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		if *in == nil {
			*out = nil
		} else {
			*out = new(int)
			**out = **in
		}
	}
	if in.KubeletRootDir != nil {
		in, out := &in.KubeletRootDir, &out.KubeletRootDir
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtletConfig.
func (in *VirtletConfig) DeepCopy() *VirtletConfig {
	if in == nil {
		return nil
	}
	out := new(VirtletConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtletConfigMapping) DeepCopyInto(out *VirtletConfigMapping) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtletConfigMapping.
func (in *VirtletConfigMapping) DeepCopy() *VirtletConfigMapping {
	if in == nil {
		return nil
	}
	out := new(VirtletConfigMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtletConfigMapping) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtletConfigMappingList) DeepCopyInto(out *VirtletConfigMappingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtletConfigMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtletConfigMappingList.
func (in *VirtletConfigMappingList) DeepCopy() *VirtletConfigMappingList {
	if in == nil {
		return nil
	}
	out := new(VirtletConfigMappingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtletConfigMappingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtletConfigMappingSpec) DeepCopyInto(out *VirtletConfigMappingSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		if *in == nil {
			*out = nil
		} else {
			*out = new(VirtletConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtletConfigMappingSpec.
func (in *VirtletConfigMappingSpec) DeepCopy() *VirtletConfigMappingSpec {
	if in == nil {
		return nil
	}
	out := new(VirtletConfigMappingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtletImageMapping) DeepCopyInto(out *VirtletImageMapping) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtletImageMapping.
func (in *VirtletImageMapping) DeepCopy() *VirtletImageMapping {
	if in == nil {
		return nil
	}
	out := new(VirtletImageMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtletImageMapping) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtletImageMappingList) DeepCopyInto(out *VirtletImageMappingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtletImageMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtletImageMappingList.
func (in *VirtletImageMappingList) DeepCopy() *VirtletImageMappingList {
	if in == nil {
		return nil
	}
	out := new(VirtletImageMappingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtletImageMappingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}