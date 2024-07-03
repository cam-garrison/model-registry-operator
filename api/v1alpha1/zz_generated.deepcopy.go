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

package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayConfig) DeepCopyInto(out *GatewayConfig) {
	*out = *in
	if in.IstioIngress != nil {
		in, out := &in.IstioIngress, &out.IstioIngress
		*out = new(string)
		**out = **in
	}
	if in.ControlPlane != nil {
		in, out := &in.ControlPlane, &out.ControlPlane
		*out = new(string)
		**out = **in
	}
	in.Rest.DeepCopyInto(&out.Rest)
	in.Grpc.DeepCopyInto(&out.Grpc)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayConfig.
func (in *GatewayConfig) DeepCopy() *GatewayConfig {
	if in == nil {
		return nil
	}
	out := new(GatewayConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrpcSpec) DeepCopyInto(out *GrpcSpec) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrpcSpec.
func (in *GrpcSpec) DeepCopy() *GrpcSpec {
	if in == nil {
		return nil
	}
	out := new(GrpcSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioConfig) DeepCopyInto(out *IstioConfig) {
	*out = *in
	if in.AuthConfigLabels != nil {
		in, out := &in.AuthConfigLabels, &out.AuthConfigLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(GatewayConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Audiences != nil {
		in, out := &in.Audiences, &out.Audiences
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioConfig.
func (in *IstioConfig) DeepCopy() *IstioConfig {
	if in == nil {
		return nil
	}
	out := new(IstioConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelRegistry) DeepCopyInto(out *ModelRegistry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelRegistry.
func (in *ModelRegistry) DeepCopy() *ModelRegistry {
	if in == nil {
		return nil
	}
	out := new(ModelRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelRegistry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelRegistryList) DeepCopyInto(out *ModelRegistryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ModelRegistry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelRegistryList.
func (in *ModelRegistryList) DeepCopy() *ModelRegistryList {
	if in == nil {
		return nil
	}
	out := new(ModelRegistryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelRegistryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelRegistrySpec) DeepCopyInto(out *ModelRegistrySpec) {
	*out = *in
	in.Rest.DeepCopyInto(&out.Rest)
	in.Grpc.DeepCopyInto(&out.Grpc)
	if in.Postgres != nil {
		in, out := &in.Postgres, &out.Postgres
		*out = new(PostgresConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.MySQL != nil {
		in, out := &in.MySQL, &out.MySQL
		*out = new(MySQLConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.EnableDatabaseUpgrade != nil {
		in, out := &in.EnableDatabaseUpgrade, &out.EnableDatabaseUpgrade
		*out = new(bool)
		**out = **in
	}
	if in.DowngradeDbSchemaVersion != nil {
		in, out := &in.DowngradeDbSchemaVersion, &out.DowngradeDbSchemaVersion
		*out = new(int64)
		**out = **in
	}
	if in.Istio != nil {
		in, out := &in.Istio, &out.Istio
		*out = new(IstioConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelRegistrySpec.
func (in *ModelRegistrySpec) DeepCopy() *ModelRegistrySpec {
	if in == nil {
		return nil
	}
	out := new(ModelRegistrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelRegistryStatus) DeepCopyInto(out *ModelRegistryStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelRegistryStatus.
func (in *ModelRegistryStatus) DeepCopy() *ModelRegistryStatus {
	if in == nil {
		return nil
	}
	out := new(ModelRegistryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLConfig) DeepCopyInto(out *MySQLConfig) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.PasswordSecret != nil {
		in, out := &in.PasswordSecret, &out.PasswordSecret
		*out = new(SecretKeyValue)
		**out = **in
	}
	if in.SSLCertificateSecret != nil {
		in, out := &in.SSLCertificateSecret, &out.SSLCertificateSecret
		*out = new(SecretKeyValue)
		**out = **in
	}
	if in.SSLKeySecret != nil {
		in, out := &in.SSLKeySecret, &out.SSLKeySecret
		*out = new(SecretKeyValue)
		**out = **in
	}
	if in.SSLRootCertificateSecret != nil {
		in, out := &in.SSLRootCertificateSecret, &out.SSLRootCertificateSecret
		*out = new(SecretKeyValue)
		**out = **in
	}
	if in.SSLRootCertificatesSecretName != nil {
		in, out := &in.SSLRootCertificatesSecretName, &out.SSLRootCertificatesSecretName
		*out = new(string)
		**out = **in
	}
	if in.SSLCipher != nil {
		in, out := &in.SSLCipher, &out.SSLCipher
		*out = new(string)
		**out = **in
	}
	if in.VerifyServerCert != nil {
		in, out := &in.VerifyServerCert, &out.VerifyServerCert
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLConfig.
func (in *MySQLConfig) DeepCopy() *MySQLConfig {
	if in == nil {
		return nil
	}
	out := new(MySQLConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresConfig) DeepCopyInto(out *PostgresConfig) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.PasswordSecret != nil {
		in, out := &in.PasswordSecret, &out.PasswordSecret
		*out = new(SecretKeyValue)
		**out = **in
	}
	if in.SSLCertificateSecret != nil {
		in, out := &in.SSLCertificateSecret, &out.SSLCertificateSecret
		*out = new(SecretKeyValue)
		**out = **in
	}
	if in.SSLKeySecret != nil {
		in, out := &in.SSLKeySecret, &out.SSLKeySecret
		*out = new(SecretKeyValue)
		**out = **in
	}
	if in.SSLPasswordSecret != nil {
		in, out := &in.SSLPasswordSecret, &out.SSLPasswordSecret
		*out = new(SecretKeyValue)
		**out = **in
	}
	if in.SSLRootCertificateSecret != nil {
		in, out := &in.SSLRootCertificateSecret, &out.SSLRootCertificateSecret
		*out = new(SecretKeyValue)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresConfig.
func (in *PostgresConfig) DeepCopy() *PostgresConfig {
	if in == nil {
		return nil
	}
	out := new(PostgresConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestSpec) DeepCopyInto(out *RestSpec) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestSpec.
func (in *RestSpec) DeepCopy() *RestSpec {
	if in == nil {
		return nil
	}
	out := new(RestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyValue) DeepCopyInto(out *SecretKeyValue) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyValue.
func (in *SecretKeyValue) DeepCopy() *SecretKeyValue {
	if in == nil {
		return nil
	}
	out := new(SecretKeyValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerConfig) DeepCopyInto(out *ServerConfig) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSServerSettings)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerConfig.
func (in *ServerConfig) DeepCopy() *ServerConfig {
	if in == nil {
		return nil
	}
	out := new(ServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSServerSettings) DeepCopyInto(out *TLSServerSettings) {
	*out = *in
	if in.CredentialName != nil {
		in, out := &in.CredentialName, &out.CredentialName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSServerSettings.
func (in *TLSServerSettings) DeepCopy() *TLSServerSettings {
	if in == nil {
		return nil
	}
	out := new(TLSServerSettings)
	in.DeepCopyInto(out)
	return out
}
