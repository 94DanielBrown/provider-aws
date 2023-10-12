/*
Copyright 2023 The Crossplane Authors.

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

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"k8s.io/utils/ptr"
)

// ConfugurationARN returns the status.atProvider.ARN of a Configuration.
func ConfugurationARN() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		r, ok := mg.(*Configuration)
		if !ok {
			return ""
		}
		return ptr.Deref(r.Status.AtProvider.ARN, "")

	}
}
