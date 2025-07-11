//go:build !ignore_autogenerated

/*
Copyright 2023.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentSource) DeepCopyInto(out *ContentSource) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(ContentSourceRef)
		**out = **in
	}
	if in.ConfigMapRef != nil {
		in, out := &in.ConfigMapRef, &out.ConfigMapRef
		*out = new(ContentSourceRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentSource.
func (in *ContentSource) DeepCopy() *ContentSource {
	if in == nil {
		return nil
	}
	out := new(ContentSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentSourceRef) DeepCopyInto(out *ContentSourceRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentSourceRef.
func (in *ContentSourceRef) DeepCopy() *ContentSourceRef {
	if in == nil {
		return nil
	}
	out := new(ContentSourceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *File) DeepCopyInto(out *File) {
	*out = *in
	out.File = in.File
	if in.ContentFrom != nil {
		in, out := &in.ContentFrom, &out.ContentFrom
		*out = new(ContentSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new File.
func (in *File) DeepCopy() *File {
	if in == nil {
		return nil
	}
	out := new(File)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JoinTokenSecretRef) DeepCopyInto(out *JoinTokenSecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JoinTokenSecretRef.
func (in *JoinTokenSecretRef) DeepCopy() *JoinTokenSecretRef {
	if in == nil {
		return nil
	}
	out := new(JoinTokenSecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K0sConfigSpec) DeepCopyInto(out *K0sConfigSpec) {
	*out = *in
	if in.K0s != nil {
		in, out := &in.K0s, &out.K0s
		*out = (*in).DeepCopy()
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]File, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PreStartCommands != nil {
		in, out := &in.PreStartCommands, &out.PreStartCommands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PostStartCommands != nil {
		in, out := &in.PostStartCommands, &out.PostStartCommands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Tunneling = in.Tunneling
	if in.CustomUserDataRef != nil {
		in, out := &in.CustomUserDataRef, &out.CustomUserDataRef
		*out = new(ContentSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K0sConfigSpec.
func (in *K0sConfigSpec) DeepCopy() *K0sConfigSpec {
	if in == nil {
		return nil
	}
	out := new(K0sConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K0sControllerConfig) DeepCopyInto(out *K0sControllerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K0sControllerConfig.
func (in *K0sControllerConfig) DeepCopy() *K0sControllerConfig {
	if in == nil {
		return nil
	}
	out := new(K0sControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *K0sControllerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K0sControllerConfigList) DeepCopyInto(out *K0sControllerConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]K0sControllerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K0sControllerConfigList.
func (in *K0sControllerConfigList) DeepCopy() *K0sControllerConfigList {
	if in == nil {
		return nil
	}
	out := new(K0sControllerConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *K0sControllerConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K0sControllerConfigSpec) DeepCopyInto(out *K0sControllerConfigSpec) {
	*out = *in
	if in.K0sConfigSpec != nil {
		in, out := &in.K0sConfigSpec, &out.K0sConfigSpec
		*out = new(K0sConfigSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K0sControllerConfigSpec.
func (in *K0sControllerConfigSpec) DeepCopy() *K0sControllerConfigSpec {
	if in == nil {
		return nil
	}
	out := new(K0sControllerConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K0sControllerConfigStatus) DeepCopyInto(out *K0sControllerConfigStatus) {
	*out = *in
	if in.DataSecretName != nil {
		in, out := &in.DataSecretName, &out.DataSecretName
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K0sControllerConfigStatus.
func (in *K0sControllerConfigStatus) DeepCopy() *K0sControllerConfigStatus {
	if in == nil {
		return nil
	}
	out := new(K0sControllerConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K0sWorkerConfig) DeepCopyInto(out *K0sWorkerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K0sWorkerConfig.
func (in *K0sWorkerConfig) DeepCopy() *K0sWorkerConfig {
	if in == nil {
		return nil
	}
	out := new(K0sWorkerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *K0sWorkerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K0sWorkerConfigList) DeepCopyInto(out *K0sWorkerConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]K0sWorkerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K0sWorkerConfigList.
func (in *K0sWorkerConfigList) DeepCopy() *K0sWorkerConfigList {
	if in == nil {
		return nil
	}
	out := new(K0sWorkerConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *K0sWorkerConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K0sWorkerConfigSpec) DeepCopyInto(out *K0sWorkerConfigSpec) {
	*out = *in
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]File, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PreStartCommands != nil {
		in, out := &in.PreStartCommands, &out.PreStartCommands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PostStartCommands != nil {
		in, out := &in.PostStartCommands, &out.PostStartCommands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomUserDataRef != nil {
		in, out := &in.CustomUserDataRef, &out.CustomUserDataRef
		*out = new(ContentSource)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretMetadata != nil {
		in, out := &in.SecretMetadata, &out.SecretMetadata
		*out = new(SecretMetadata)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K0sWorkerConfigSpec.
func (in *K0sWorkerConfigSpec) DeepCopy() *K0sWorkerConfigSpec {
	if in == nil {
		return nil
	}
	out := new(K0sWorkerConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K0sWorkerConfigStatus) DeepCopyInto(out *K0sWorkerConfigStatus) {
	*out = *in
	if in.DataSecretName != nil {
		in, out := &in.DataSecretName, &out.DataSecretName
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K0sWorkerConfigStatus.
func (in *K0sWorkerConfigStatus) DeepCopy() *K0sWorkerConfigStatus {
	if in == nil {
		return nil
	}
	out := new(K0sWorkerConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K0sWorkerConfigTemplate) DeepCopyInto(out *K0sWorkerConfigTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K0sWorkerConfigTemplate.
func (in *K0sWorkerConfigTemplate) DeepCopy() *K0sWorkerConfigTemplate {
	if in == nil {
		return nil
	}
	out := new(K0sWorkerConfigTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *K0sWorkerConfigTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K0sWorkerConfigTemplateList) DeepCopyInto(out *K0sWorkerConfigTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]K0sWorkerConfigTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K0sWorkerConfigTemplateList.
func (in *K0sWorkerConfigTemplateList) DeepCopy() *K0sWorkerConfigTemplateList {
	if in == nil {
		return nil
	}
	out := new(K0sWorkerConfigTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *K0sWorkerConfigTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K0sWorkerConfigTemplateResource) DeepCopyInto(out *K0sWorkerConfigTemplateResource) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K0sWorkerConfigTemplateResource.
func (in *K0sWorkerConfigTemplateResource) DeepCopy() *K0sWorkerConfigTemplateResource {
	if in == nil {
		return nil
	}
	out := new(K0sWorkerConfigTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K0sWorkerConfigTemplateSpec) DeepCopyInto(out *K0sWorkerConfigTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K0sWorkerConfigTemplateSpec.
func (in *K0sWorkerConfigTemplateSpec) DeepCopy() *K0sWorkerConfigTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(K0sWorkerConfigTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretMetadata) DeepCopyInto(out *SecretMetadata) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretMetadata.
func (in *SecretMetadata) DeepCopy() *SecretMetadata {
	if in == nil {
		return nil
	}
	out := new(SecretMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelingSpec) DeepCopyInto(out *TunnelingSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelingSpec.
func (in *TunnelingSpec) DeepCopy() *TunnelingSpec {
	if in == nil {
		return nil
	}
	out := new(TunnelingSpec)
	in.DeepCopyInto(out)
	return out
}
