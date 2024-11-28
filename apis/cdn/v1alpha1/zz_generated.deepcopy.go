//go:build !ignore_autogenerated

/*
Copyright 2022 YANDEX LLC
This is modified version of the software, made by the Crossplane Authors
and available at: https://github.com/crossplane-contrib/provider-jet-template

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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddressACLInitParameters) DeepCopyInto(out *IPAddressACLInitParameters) {
	*out = *in
	if in.ExceptedValues != nil {
		in, out := &in.ExceptedValues, &out.ExceptedValues
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PolicyType != nil {
		in, out := &in.PolicyType, &out.PolicyType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddressACLInitParameters.
func (in *IPAddressACLInitParameters) DeepCopy() *IPAddressACLInitParameters {
	if in == nil {
		return nil
	}
	out := new(IPAddressACLInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddressACLObservation) DeepCopyInto(out *IPAddressACLObservation) {
	*out = *in
	if in.ExceptedValues != nil {
		in, out := &in.ExceptedValues, &out.ExceptedValues
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PolicyType != nil {
		in, out := &in.PolicyType, &out.PolicyType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddressACLObservation.
func (in *IPAddressACLObservation) DeepCopy() *IPAddressACLObservation {
	if in == nil {
		return nil
	}
	out := new(IPAddressACLObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddressACLParameters) DeepCopyInto(out *IPAddressACLParameters) {
	*out = *in
	if in.ExceptedValues != nil {
		in, out := &in.ExceptedValues, &out.ExceptedValues
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PolicyType != nil {
		in, out := &in.PolicyType, &out.PolicyType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddressACLParameters.
func (in *IPAddressACLParameters) DeepCopy() *IPAddressACLParameters {
	if in == nil {
		return nil
	}
	out := new(IPAddressACLParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionsInitParameters) DeepCopyInto(out *OptionsInitParameters) {
	*out = *in
	if in.AllowedHTTPMethods != nil {
		in, out := &in.AllowedHTTPMethods, &out.AllowedHTTPMethods
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BrowserCacheSettings != nil {
		in, out := &in.BrowserCacheSettings, &out.BrowserCacheSettings
		*out = new(float64)
		**out = **in
	}
	if in.CacheHTTPHeaders != nil {
		in, out := &in.CacheHTTPHeaders, &out.CacheHTTPHeaders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Cors != nil {
		in, out := &in.Cors, &out.Cors
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CustomHostHeader != nil {
		in, out := &in.CustomHostHeader, &out.CustomHostHeader
		*out = new(string)
		**out = **in
	}
	if in.CustomServerName != nil {
		in, out := &in.CustomServerName, &out.CustomServerName
		*out = new(string)
		**out = **in
	}
	if in.DisableCache != nil {
		in, out := &in.DisableCache, &out.DisableCache
		*out = new(bool)
		**out = **in
	}
	if in.DisableProxyForceRanges != nil {
		in, out := &in.DisableProxyForceRanges, &out.DisableProxyForceRanges
		*out = new(bool)
		**out = **in
	}
	if in.EdgeCacheSettings != nil {
		in, out := &in.EdgeCacheSettings, &out.EdgeCacheSettings
		*out = new(float64)
		**out = **in
	}
	if in.EnableIPURLSigning != nil {
		in, out := &in.EnableIPURLSigning, &out.EnableIPURLSigning
		*out = new(bool)
		**out = **in
	}
	if in.FetchedCompressed != nil {
		in, out := &in.FetchedCompressed, &out.FetchedCompressed
		*out = new(bool)
		**out = **in
	}
	if in.ForwardHostHeader != nil {
		in, out := &in.ForwardHostHeader, &out.ForwardHostHeader
		*out = new(bool)
		**out = **in
	}
	if in.GzipOn != nil {
		in, out := &in.GzipOn, &out.GzipOn
		*out = new(bool)
		**out = **in
	}
	if in.IPAddressACL != nil {
		in, out := &in.IPAddressACL, &out.IPAddressACL
		*out = make([]IPAddressACLInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IgnoreCookie != nil {
		in, out := &in.IgnoreCookie, &out.IgnoreCookie
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreQueryParams != nil {
		in, out := &in.IgnoreQueryParams, &out.IgnoreQueryParams
		*out = new(bool)
		**out = **in
	}
	if in.ProxyCacheMethodsSet != nil {
		in, out := &in.ProxyCacheMethodsSet, &out.ProxyCacheMethodsSet
		*out = new(bool)
		**out = **in
	}
	if in.QueryParamsBlacklist != nil {
		in, out := &in.QueryParamsBlacklist, &out.QueryParamsBlacklist
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.QueryParamsWhitelist != nil {
		in, out := &in.QueryParamsWhitelist, &out.QueryParamsWhitelist
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RedirectHTTPSToHTTP != nil {
		in, out := &in.RedirectHTTPSToHTTP, &out.RedirectHTTPSToHTTP
		*out = new(bool)
		**out = **in
	}
	if in.RedirectHTTPToHTTPS != nil {
		in, out := &in.RedirectHTTPToHTTPS, &out.RedirectHTTPToHTTPS
		*out = new(bool)
		**out = **in
	}
	if in.SecureKey != nil {
		in, out := &in.SecureKey, &out.SecureKey
		*out = new(string)
		**out = **in
	}
	if in.Slice != nil {
		in, out := &in.Slice, &out.Slice
		*out = new(bool)
		**out = **in
	}
	if in.StaticRequestHeaders != nil {
		in, out := &in.StaticRequestHeaders, &out.StaticRequestHeaders
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.StaticResponseHeaders != nil {
		in, out := &in.StaticResponseHeaders, &out.StaticResponseHeaders
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionsInitParameters.
func (in *OptionsInitParameters) DeepCopy() *OptionsInitParameters {
	if in == nil {
		return nil
	}
	out := new(OptionsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionsObservation) DeepCopyInto(out *OptionsObservation) {
	*out = *in
	if in.AllowedHTTPMethods != nil {
		in, out := &in.AllowedHTTPMethods, &out.AllowedHTTPMethods
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BrowserCacheSettings != nil {
		in, out := &in.BrowserCacheSettings, &out.BrowserCacheSettings
		*out = new(float64)
		**out = **in
	}
	if in.CacheHTTPHeaders != nil {
		in, out := &in.CacheHTTPHeaders, &out.CacheHTTPHeaders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Cors != nil {
		in, out := &in.Cors, &out.Cors
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CustomHostHeader != nil {
		in, out := &in.CustomHostHeader, &out.CustomHostHeader
		*out = new(string)
		**out = **in
	}
	if in.CustomServerName != nil {
		in, out := &in.CustomServerName, &out.CustomServerName
		*out = new(string)
		**out = **in
	}
	if in.DisableCache != nil {
		in, out := &in.DisableCache, &out.DisableCache
		*out = new(bool)
		**out = **in
	}
	if in.DisableProxyForceRanges != nil {
		in, out := &in.DisableProxyForceRanges, &out.DisableProxyForceRanges
		*out = new(bool)
		**out = **in
	}
	if in.EdgeCacheSettings != nil {
		in, out := &in.EdgeCacheSettings, &out.EdgeCacheSettings
		*out = new(float64)
		**out = **in
	}
	if in.EnableIPURLSigning != nil {
		in, out := &in.EnableIPURLSigning, &out.EnableIPURLSigning
		*out = new(bool)
		**out = **in
	}
	if in.FetchedCompressed != nil {
		in, out := &in.FetchedCompressed, &out.FetchedCompressed
		*out = new(bool)
		**out = **in
	}
	if in.ForwardHostHeader != nil {
		in, out := &in.ForwardHostHeader, &out.ForwardHostHeader
		*out = new(bool)
		**out = **in
	}
	if in.GzipOn != nil {
		in, out := &in.GzipOn, &out.GzipOn
		*out = new(bool)
		**out = **in
	}
	if in.IPAddressACL != nil {
		in, out := &in.IPAddressACL, &out.IPAddressACL
		*out = make([]IPAddressACLObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IgnoreCookie != nil {
		in, out := &in.IgnoreCookie, &out.IgnoreCookie
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreQueryParams != nil {
		in, out := &in.IgnoreQueryParams, &out.IgnoreQueryParams
		*out = new(bool)
		**out = **in
	}
	if in.ProxyCacheMethodsSet != nil {
		in, out := &in.ProxyCacheMethodsSet, &out.ProxyCacheMethodsSet
		*out = new(bool)
		**out = **in
	}
	if in.QueryParamsBlacklist != nil {
		in, out := &in.QueryParamsBlacklist, &out.QueryParamsBlacklist
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.QueryParamsWhitelist != nil {
		in, out := &in.QueryParamsWhitelist, &out.QueryParamsWhitelist
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RedirectHTTPSToHTTP != nil {
		in, out := &in.RedirectHTTPSToHTTP, &out.RedirectHTTPSToHTTP
		*out = new(bool)
		**out = **in
	}
	if in.RedirectHTTPToHTTPS != nil {
		in, out := &in.RedirectHTTPToHTTPS, &out.RedirectHTTPToHTTPS
		*out = new(bool)
		**out = **in
	}
	if in.SecureKey != nil {
		in, out := &in.SecureKey, &out.SecureKey
		*out = new(string)
		**out = **in
	}
	if in.Slice != nil {
		in, out := &in.Slice, &out.Slice
		*out = new(bool)
		**out = **in
	}
	if in.StaticRequestHeaders != nil {
		in, out := &in.StaticRequestHeaders, &out.StaticRequestHeaders
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.StaticResponseHeaders != nil {
		in, out := &in.StaticResponseHeaders, &out.StaticResponseHeaders
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionsObservation.
func (in *OptionsObservation) DeepCopy() *OptionsObservation {
	if in == nil {
		return nil
	}
	out := new(OptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionsParameters) DeepCopyInto(out *OptionsParameters) {
	*out = *in
	if in.AllowedHTTPMethods != nil {
		in, out := &in.AllowedHTTPMethods, &out.AllowedHTTPMethods
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BrowserCacheSettings != nil {
		in, out := &in.BrowserCacheSettings, &out.BrowserCacheSettings
		*out = new(float64)
		**out = **in
	}
	if in.CacheHTTPHeaders != nil {
		in, out := &in.CacheHTTPHeaders, &out.CacheHTTPHeaders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Cors != nil {
		in, out := &in.Cors, &out.Cors
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CustomHostHeader != nil {
		in, out := &in.CustomHostHeader, &out.CustomHostHeader
		*out = new(string)
		**out = **in
	}
	if in.CustomServerName != nil {
		in, out := &in.CustomServerName, &out.CustomServerName
		*out = new(string)
		**out = **in
	}
	if in.DisableCache != nil {
		in, out := &in.DisableCache, &out.DisableCache
		*out = new(bool)
		**out = **in
	}
	if in.DisableProxyForceRanges != nil {
		in, out := &in.DisableProxyForceRanges, &out.DisableProxyForceRanges
		*out = new(bool)
		**out = **in
	}
	if in.EdgeCacheSettings != nil {
		in, out := &in.EdgeCacheSettings, &out.EdgeCacheSettings
		*out = new(float64)
		**out = **in
	}
	if in.EnableIPURLSigning != nil {
		in, out := &in.EnableIPURLSigning, &out.EnableIPURLSigning
		*out = new(bool)
		**out = **in
	}
	if in.FetchedCompressed != nil {
		in, out := &in.FetchedCompressed, &out.FetchedCompressed
		*out = new(bool)
		**out = **in
	}
	if in.ForwardHostHeader != nil {
		in, out := &in.ForwardHostHeader, &out.ForwardHostHeader
		*out = new(bool)
		**out = **in
	}
	if in.GzipOn != nil {
		in, out := &in.GzipOn, &out.GzipOn
		*out = new(bool)
		**out = **in
	}
	if in.IPAddressACL != nil {
		in, out := &in.IPAddressACL, &out.IPAddressACL
		*out = make([]IPAddressACLParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IgnoreCookie != nil {
		in, out := &in.IgnoreCookie, &out.IgnoreCookie
		*out = new(bool)
		**out = **in
	}
	if in.IgnoreQueryParams != nil {
		in, out := &in.IgnoreQueryParams, &out.IgnoreQueryParams
		*out = new(bool)
		**out = **in
	}
	if in.ProxyCacheMethodsSet != nil {
		in, out := &in.ProxyCacheMethodsSet, &out.ProxyCacheMethodsSet
		*out = new(bool)
		**out = **in
	}
	if in.QueryParamsBlacklist != nil {
		in, out := &in.QueryParamsBlacklist, &out.QueryParamsBlacklist
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.QueryParamsWhitelist != nil {
		in, out := &in.QueryParamsWhitelist, &out.QueryParamsWhitelist
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RedirectHTTPSToHTTP != nil {
		in, out := &in.RedirectHTTPSToHTTP, &out.RedirectHTTPSToHTTP
		*out = new(bool)
		**out = **in
	}
	if in.RedirectHTTPToHTTPS != nil {
		in, out := &in.RedirectHTTPToHTTPS, &out.RedirectHTTPToHTTPS
		*out = new(bool)
		**out = **in
	}
	if in.SecureKey != nil {
		in, out := &in.SecureKey, &out.SecureKey
		*out = new(string)
		**out = **in
	}
	if in.Slice != nil {
		in, out := &in.Slice, &out.Slice
		*out = new(bool)
		**out = **in
	}
	if in.StaticRequestHeaders != nil {
		in, out := &in.StaticRequestHeaders, &out.StaticRequestHeaders
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.StaticResponseHeaders != nil {
		in, out := &in.StaticResponseHeaders, &out.StaticResponseHeaders
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionsParameters.
func (in *OptionsParameters) DeepCopy() *OptionsParameters {
	if in == nil {
		return nil
	}
	out := new(OptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginGroup) DeepCopyInto(out *OriginGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginGroup.
func (in *OriginGroup) DeepCopy() *OriginGroup {
	if in == nil {
		return nil
	}
	out := new(OriginGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OriginGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginGroupInitParameters) DeepCopyInto(out *OriginGroupInitParameters) {
	*out = *in
	if in.FolderID != nil {
		in, out := &in.FolderID, &out.FolderID
		*out = new(string)
		**out = **in
	}
	if in.FolderIDRef != nil {
		in, out := &in.FolderIDRef, &out.FolderIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FolderIDSelector != nil {
		in, out := &in.FolderIDSelector, &out.FolderIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Origin != nil {
		in, out := &in.Origin, &out.Origin
		*out = make([]OriginInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UseNext != nil {
		in, out := &in.UseNext, &out.UseNext
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginGroupInitParameters.
func (in *OriginGroupInitParameters) DeepCopy() *OriginGroupInitParameters {
	if in == nil {
		return nil
	}
	out := new(OriginGroupInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginGroupList) DeepCopyInto(out *OriginGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OriginGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginGroupList.
func (in *OriginGroupList) DeepCopy() *OriginGroupList {
	if in == nil {
		return nil
	}
	out := new(OriginGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OriginGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginGroupObservation) DeepCopyInto(out *OriginGroupObservation) {
	*out = *in
	if in.FolderID != nil {
		in, out := &in.FolderID, &out.FolderID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Origin != nil {
		in, out := &in.Origin, &out.Origin
		*out = make([]OriginObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UseNext != nil {
		in, out := &in.UseNext, &out.UseNext
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginGroupObservation.
func (in *OriginGroupObservation) DeepCopy() *OriginGroupObservation {
	if in == nil {
		return nil
	}
	out := new(OriginGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginGroupParameters) DeepCopyInto(out *OriginGroupParameters) {
	*out = *in
	if in.FolderID != nil {
		in, out := &in.FolderID, &out.FolderID
		*out = new(string)
		**out = **in
	}
	if in.FolderIDRef != nil {
		in, out := &in.FolderIDRef, &out.FolderIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FolderIDSelector != nil {
		in, out := &in.FolderIDSelector, &out.FolderIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Origin != nil {
		in, out := &in.Origin, &out.Origin
		*out = make([]OriginParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UseNext != nil {
		in, out := &in.UseNext, &out.UseNext
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginGroupParameters.
func (in *OriginGroupParameters) DeepCopy() *OriginGroupParameters {
	if in == nil {
		return nil
	}
	out := new(OriginGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginGroupSpec) DeepCopyInto(out *OriginGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginGroupSpec.
func (in *OriginGroupSpec) DeepCopy() *OriginGroupSpec {
	if in == nil {
		return nil
	}
	out := new(OriginGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginGroupStatus) DeepCopyInto(out *OriginGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginGroupStatus.
func (in *OriginGroupStatus) DeepCopy() *OriginGroupStatus {
	if in == nil {
		return nil
	}
	out := new(OriginGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginInitParameters) DeepCopyInto(out *OriginInitParameters) {
	*out = *in
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(bool)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginInitParameters.
func (in *OriginInitParameters) DeepCopy() *OriginInitParameters {
	if in == nil {
		return nil
	}
	out := new(OriginInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginObservation) DeepCopyInto(out *OriginObservation) {
	*out = *in
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(bool)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.OriginGroupID != nil {
		in, out := &in.OriginGroupID, &out.OriginGroupID
		*out = new(float64)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginObservation.
func (in *OriginObservation) DeepCopy() *OriginObservation {
	if in == nil {
		return nil
	}
	out := new(OriginObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginParameters) DeepCopyInto(out *OriginParameters) {
	*out = *in
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(bool)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginParameters.
func (in *OriginParameters) DeepCopy() *OriginParameters {
	if in == nil {
		return nil
	}
	out := new(OriginParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Resource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInitParameters) DeepCopyInto(out *ResourceInitParameters) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.Cname != nil {
		in, out := &in.Cname, &out.Cname
		*out = new(string)
		**out = **in
	}
	if in.FolderID != nil {
		in, out := &in.FolderID, &out.FolderID
		*out = new(string)
		**out = **in
	}
	if in.FolderIDRef != nil {
		in, out := &in.FolderIDRef, &out.FolderIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FolderIDSelector != nil {
		in, out := &in.FolderIDSelector, &out.FolderIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]OptionsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OriginGroupID != nil {
		in, out := &in.OriginGroupID, &out.OriginGroupID
		*out = new(float64)
		**out = **in
	}
	if in.OriginGroupIDRef != nil {
		in, out := &in.OriginGroupIDRef, &out.OriginGroupIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.OriginGroupIDSelector != nil {
		in, out := &in.OriginGroupIDSelector, &out.OriginGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.OriginGroupName != nil {
		in, out := &in.OriginGroupName, &out.OriginGroupName
		*out = new(string)
		**out = **in
	}
	if in.OriginProtocol != nil {
		in, out := &in.OriginProtocol, &out.OriginProtocol
		*out = new(string)
		**out = **in
	}
	if in.SSLCertificate != nil {
		in, out := &in.SSLCertificate, &out.SSLCertificate
		*out = make([]SSLCertificateInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecondaryHostnames != nil {
		in, out := &in.SecondaryHostnames, &out.SecondaryHostnames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInitParameters.
func (in *ResourceInitParameters) DeepCopy() *ResourceInitParameters {
	if in == nil {
		return nil
	}
	out := new(ResourceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceList) DeepCopyInto(out *ResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Resource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceList.
func (in *ResourceList) DeepCopy() *ResourceList {
	if in == nil {
		return nil
	}
	out := new(ResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceObservation) DeepCopyInto(out *ResourceObservation) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.Cname != nil {
		in, out := &in.Cname, &out.Cname
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.FolderID != nil {
		in, out := &in.FolderID, &out.FolderID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]OptionsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OriginGroupID != nil {
		in, out := &in.OriginGroupID, &out.OriginGroupID
		*out = new(float64)
		**out = **in
	}
	if in.OriginGroupName != nil {
		in, out := &in.OriginGroupName, &out.OriginGroupName
		*out = new(string)
		**out = **in
	}
	if in.OriginProtocol != nil {
		in, out := &in.OriginProtocol, &out.OriginProtocol
		*out = new(string)
		**out = **in
	}
	if in.ProviderCname != nil {
		in, out := &in.ProviderCname, &out.ProviderCname
		*out = new(string)
		**out = **in
	}
	if in.SSLCertificate != nil {
		in, out := &in.SSLCertificate, &out.SSLCertificate
		*out = make([]SSLCertificateObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecondaryHostnames != nil {
		in, out := &in.SecondaryHostnames, &out.SecondaryHostnames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceObservation.
func (in *ResourceObservation) DeepCopy() *ResourceObservation {
	if in == nil {
		return nil
	}
	out := new(ResourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceParameters) DeepCopyInto(out *ResourceParameters) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.Cname != nil {
		in, out := &in.Cname, &out.Cname
		*out = new(string)
		**out = **in
	}
	if in.FolderID != nil {
		in, out := &in.FolderID, &out.FolderID
		*out = new(string)
		**out = **in
	}
	if in.FolderIDRef != nil {
		in, out := &in.FolderIDRef, &out.FolderIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FolderIDSelector != nil {
		in, out := &in.FolderIDSelector, &out.FolderIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]OptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OriginGroupID != nil {
		in, out := &in.OriginGroupID, &out.OriginGroupID
		*out = new(float64)
		**out = **in
	}
	if in.OriginGroupIDRef != nil {
		in, out := &in.OriginGroupIDRef, &out.OriginGroupIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.OriginGroupIDSelector != nil {
		in, out := &in.OriginGroupIDSelector, &out.OriginGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.OriginGroupName != nil {
		in, out := &in.OriginGroupName, &out.OriginGroupName
		*out = new(string)
		**out = **in
	}
	if in.OriginProtocol != nil {
		in, out := &in.OriginProtocol, &out.OriginProtocol
		*out = new(string)
		**out = **in
	}
	if in.SSLCertificate != nil {
		in, out := &in.SSLCertificate, &out.SSLCertificate
		*out = make([]SSLCertificateParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecondaryHostnames != nil {
		in, out := &in.SecondaryHostnames, &out.SecondaryHostnames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceParameters.
func (in *ResourceParameters) DeepCopy() *ResourceParameters {
	if in == nil {
		return nil
	}
	out := new(ResourceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpec) DeepCopyInto(out *ResourceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpec.
func (in *ResourceSpec) DeepCopy() *ResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceStatus) DeepCopyInto(out *ResourceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceStatus.
func (in *ResourceStatus) DeepCopy() *ResourceStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSLCertificateInitParameters) DeepCopyInto(out *SSLCertificateInitParameters) {
	*out = *in
	if in.CertificateManagerID != nil {
		in, out := &in.CertificateManagerID, &out.CertificateManagerID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSLCertificateInitParameters.
func (in *SSLCertificateInitParameters) DeepCopy() *SSLCertificateInitParameters {
	if in == nil {
		return nil
	}
	out := new(SSLCertificateInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSLCertificateObservation) DeepCopyInto(out *SSLCertificateObservation) {
	*out = *in
	if in.CertificateManagerID != nil {
		in, out := &in.CertificateManagerID, &out.CertificateManagerID
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSLCertificateObservation.
func (in *SSLCertificateObservation) DeepCopy() *SSLCertificateObservation {
	if in == nil {
		return nil
	}
	out := new(SSLCertificateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSLCertificateParameters) DeepCopyInto(out *SSLCertificateParameters) {
	*out = *in
	if in.CertificateManagerID != nil {
		in, out := &in.CertificateManagerID, &out.CertificateManagerID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSLCertificateParameters.
func (in *SSLCertificateParameters) DeepCopy() *SSLCertificateParameters {
	if in == nil {
		return nil
	}
	out := new(SSLCertificateParameters)
	in.DeepCopyInto(out)
	return out
}