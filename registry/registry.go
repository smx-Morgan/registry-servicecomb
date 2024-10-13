// Copyright 2022 CloudWeGo Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package registry

import (
	scoptions "github.com/cloudwego-contrib/cwgo-pkg/registry/servicecomb/options"
	scregistry "github.com/cloudwego-contrib/cwgo-pkg/registry/servicecomb/servicecombkitex/registry"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/go-chassis/sc-client"
)

// Option is ServiceComb option.
type Option = scoptions.Option

// WithAppId with app id option
func WithAppId(appId string) Option {
	return scoptions.WithAppId(appId)
}

// WithVersionRule with version rule option
func WithVersionRule(versionRule string) Option {
	return scoptions.WithVersionRule(versionRule)
}

// WithHostName with host name option
func WithHostName(hostName string) Option {
	return scoptions.WithHostName(hostName)
}

func WithHeartbeatInterval(second int32) Option {
	return scoptions.WithHeartbeatInterval(second)
}

// NewDefaultSCRegistry create a new default ServiceComb registry
func NewDefaultSCRegistry(opts ...Option) (registry.Registry, error) {
	return scregistry.NewDefaultSCRegistry(opts...)
}

// NewSCRegistry create a new ServiceComb registry
func NewSCRegistry(client *sc.Client, opts ...Option) registry.Registry {
	return scregistry.NewSCRegistry(client, opts...)
}
