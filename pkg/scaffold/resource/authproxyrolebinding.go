/*
Copyright 2018 The Kubernetes Authors.

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

package resource

import (
	"path/filepath"

	"sigs.k8s.io/kubebuilder/pkg/scaffold/input"
)

var _ input.File = &AuthProxyRoleBinding{}

// AuthProxyRoleBinding scaffolds the config/rbac/auth_proxy_role_binding_rbac.yaml file
type AuthProxyRoleBinding struct {
	input.Input

	// Resource is a resource in the API group
	Resource *Resource
}

// GetInput implements input.File
func (r *AuthProxyRoleBinding) GetInput() (input.Input, error) {
	if r.Path == "" {
		r.Path = filepath.Join("config", "rbac", "auth_proxy_role_binding.yaml")
	}
	r.TemplateBody = proxyRoleBindinggTemplate
	return r.Input, nil
}

// Validate validates the values
func (r *AuthProxyRoleBinding) Validate() error {
	return r.Resource.Validate()
}

var proxyRoleBindinggTemplate = `apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: proxy-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: proxy-role
subjects:
- kind: ServiceAccount
  name: default
  namespace: system
`
