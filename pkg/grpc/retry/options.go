// Copyright 2019-present Open Networking Foundation.
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

package retry

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

// WithPerCallTimeout sets the per-call retry timeout
func WithPerCallTimeout(t time.Duration) CallOption {
	return newCallOption(func(opts *callOptions) {
		opts.perCallTimeout = &t
	})
}

// WithInterval sets the base retry interval
func WithInterval(d time.Duration) CallOption {
	return newCallOption(func(opts *callOptions) {
		opts.initialInterval = &d
	})
}

// WithMaxInterval sets the maximum retry interval
func WithMaxInterval(d time.Duration) CallOption {
	return newCallOption(func(opts *callOptions) {
		opts.maxInterval = &d
	})
}

// WithRetryOn sets the codes on which to retry a request
func WithRetryOn(codes ...codes.Code) CallOption {
	return newCallOption(func(opts *callOptions) {
		opts.codes = codes
	})
}

func newCallOption(f func(opts *callOptions)) CallOption {
	return CallOption{
		applyFunc: f,
	}
}

// CallOption is a retrying interceptor call option
type CallOption struct {
	grpc.EmptyCallOption // make sure we implement private after() and before() fields so we don't panic.
	applyFunc            func(opts *callOptions)
}

type callOptions struct {
	perCallTimeout  *time.Duration
	initialInterval *time.Duration
	maxInterval     *time.Duration
	codes           []codes.Code
}

func newCallContext(ctx context.Context, opts *callOptions) context.Context {
	if opts.perCallTimeout != nil {
		ctx, _ = context.WithTimeout(ctx, *opts.perCallTimeout) //nolint:govet
	}
	return ctx
}

func newCallOptions(opts *callOptions, options []CallOption) *callOptions {
	if len(options) == 0 {
		return opts
	}
	optCopy := &callOptions{}
	*optCopy = *opts
	for _, f := range options {
		f.applyFunc(optCopy)
	}
	return optCopy
}

func filterCallOptions(options []grpc.CallOption) (grpcOptions []grpc.CallOption, retryOptions []CallOption) {
	for _, opt := range options {
		if co, ok := opt.(CallOption); ok {
			retryOptions = append(retryOptions, co)
		} else {
			grpcOptions = append(grpcOptions, opt)
		}
	}
	return grpcOptions, retryOptions
}
