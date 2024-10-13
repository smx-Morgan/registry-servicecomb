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

package resolver

import (
	"github.com/go-chassis/sc-client"

	scoptions "github.com/cloudwego-contrib/cwgo-pkg/registry/servicecomb/options"
	scresolver "github.com/cloudwego-contrib/cwgo-pkg/registry/servicecomb/servicecombkitex/resolver"
	"github.com/cloudwego/kitex/pkg/discovery"
)

// Option is service-comb resolver option.
type Option = scoptions.ResolverOption

// WithAppId with appId option.
func WithAppId(appId string) Option {
	return scoptions.WithResolverAppId(appId)
}

// WithVersionRule with versionRule option.
func WithVersionRule(versionRule string) Option {
	return scoptions.WithResolverVersionRule(versionRule)
}

// WithConsumerId with consumerId option.
func WithConsumerId(consumerId string) Option {
	return scoptions.WithResolverConsumerId(consumerId)
}

func NewDefaultSCResolver(opts ...Option) (discovery.Resolver, error) {
	return scresolver.NewDefaultSCResolver(opts...)
}

func NewSCResolver(cli *sc.Client, opts ...Option) discovery.Resolver {
	return scresolver.NewSCResolver(cli, opts...)
}
