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
	"github.com/cenkalti/backoff"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"sync"
)

var log = logging.GetLogger("onos", "grpc", "retry")

var defaultOptions = &callOptions{
	codes: []codes.Code{
		codes.Unavailable,
		codes.Unknown,
	},
}

// RetryingUnaryClientInterceptor returns a UnaryClientInterceptor that retries requests
func RetryingUnaryClientInterceptor(callOpts ...CallOption) func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	connOpts := reuseOrNewWithCallOptions(defaultOptions, callOpts)
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		grpcOpts, retryOpts := filterCallOptions(opts)
		callOpts := reuseOrNewWithCallOptions(connOpts, retryOpts)
		b := backoff.NewExponentialBackOff()
		if callOpts.initialInterval != nil {
			b.InitialInterval = *callOpts.initialInterval
		}
		if callOpts.maxInterval != nil {
			b.MaxInterval = *callOpts.maxInterval
		}
		return backoff.Retry(func() error {
			log.Debugf("Sending %s", req)
			callCtx := perCallContext(ctx, callOpts)
			if err := invoker(callCtx, method, req, reply, cc, grpcOpts...); err != nil {
				if isRetryable(ctx, callOpts, err) {
					log.Debugf("Sending %s failed", req, err)
					return err
				}
				log.Warnf("Sending %s failed", req, err)
				return backoff.Permanent(err)
			}
			return nil
		}, backoff.WithContext(b, ctx))
	}
}

// RetryingStreamClientInterceptor returns a ClientStreamInterceptor that retries both requests and responses
func RetryingStreamClientInterceptor(callOpts ...CallOption) func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		if desc.ClientStreams && desc.ServerStreams {
			return newBiDirectionalStreamClientInterceptor(callOpts...)(ctx, desc, cc, method, streamer, opts...)
		} else if desc.ClientStreams {
			return newClientStreamClientInterceptor(callOpts...)(ctx, desc, cc, method, streamer, opts...)
		} else if desc.ServerStreams {
			return newServerStreamClientInterceptor(callOpts...)(ctx, desc, cc, method, streamer, opts...)
		}
		panic("Invalid StreamDesc")
	}
}

// newClientStreamClientInterceptor returns a ClientStreamInterceptor that retries both requests and responses
func newClientStreamClientInterceptor(callOpts ...CallOption) func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	connOpts := reuseOrNewWithCallOptions(defaultOptions, callOpts)
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		grpcOpts, retryOpts := filterCallOptions(opts)
		callOpts := reuseOrNewWithCallOptions(connOpts, retryOpts)
		stream := &retryingClientStream{
			ctx:    ctx,
			buffer: &retryingClientStreamBuffer{},
			opts:   callOpts,
			newStream: func(ctx context.Context) (grpc.ClientStream, error) {
				return streamer(ctx, desc, cc, method, grpcOpts...)
			},
		}
		return stream, stream.retryStream()
	}
}

// newServerStreamClientInterceptor returns a ClientStreamInterceptor that retries both requests and responses
func newServerStreamClientInterceptor(callOpts ...CallOption) func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	connOpts := reuseOrNewWithCallOptions(defaultOptions, callOpts)
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		grpcOpts, retryOpts := filterCallOptions(opts)
		callOpts := reuseOrNewWithCallOptions(connOpts, retryOpts)
		stream := &retryingClientStream{
			ctx:    ctx,
			buffer: &retryingServerStreamBuffer{},
			opts:   callOpts,
			newStream: func(ctx context.Context) (grpc.ClientStream, error) {
				return streamer(ctx, desc, cc, method, grpcOpts...)
			},
		}
		return stream, stream.retryStream()
	}
}

// newBiDirectionalStreamClientInterceptor returns a ClientStreamInterceptor that retries both requests and responses
func newBiDirectionalStreamClientInterceptor(callOpts ...CallOption) func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	connOpts := reuseOrNewWithCallOptions(defaultOptions, callOpts)
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		grpcOpts, retryOpts := filterCallOptions(opts)
		callOpts := reuseOrNewWithCallOptions(connOpts, retryOpts)
		stream := &retryingClientStream{
			ctx:    ctx,
			buffer: &retryingBiDirectionalStreamBuffer{},
			opts:   callOpts,
			newStream: func(ctx context.Context) (grpc.ClientStream, error) {
				return streamer(ctx, desc, cc, method, grpcOpts...)
			},
		}
		return stream, stream.retryStream()
	}
}

type retryingStreamBuffer interface {
	append(interface{})
	list() []interface{}
}

type retryingClientStreamBuffer struct {
	buffer []interface{}
	mu     sync.RWMutex
}

func (b *retryingClientStreamBuffer) append(msg interface{}) {
	b.mu.Lock()
	b.buffer = append(b.buffer, msg)
	b.mu.Unlock()
}

func (b *retryingClientStreamBuffer) list() []interface{} {
	b.mu.RLock()
	buffer := make([]interface{}, len(b.buffer))
	copy(buffer, b.buffer)
	b.mu.RUnlock()
	return buffer
}

type retryingServerStreamBuffer struct {
	msg interface{}
	mu  sync.RWMutex
}

func (b *retryingServerStreamBuffer) append(msg interface{}) {
	b.mu.Lock()
	b.msg = msg
	b.mu.Unlock()
}

func (b *retryingServerStreamBuffer) list() []interface{} {
	b.mu.RLock()
	msg := b.msg
	b.mu.RUnlock()
	if msg != nil {
		return []interface{}{msg}
	}
	return []interface{}{}
}

type retryingBiDirectionalStreamBuffer struct{}

func (b *retryingBiDirectionalStreamBuffer) append(interface{}) {

}

func (b *retryingBiDirectionalStreamBuffer) list() []interface{} {
	return []interface{}{}
}

type retryingClientStream struct {
	ctx       context.Context
	stream    grpc.ClientStream
	opts      *callOptions
	mu        sync.RWMutex
	buffer    retryingStreamBuffer
	newStream func(ctx context.Context) (grpc.ClientStream, error)
	closed    bool
}

func (s *retryingClientStream) setStream(stream grpc.ClientStream) {
	s.mu.Lock()
	s.stream = stream
	s.mu.Unlock()
}

func (s *retryingClientStream) getStream() grpc.ClientStream {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.stream
}

func (s *retryingClientStream) Context() context.Context {
	return s.ctx
}

func (s *retryingClientStream) CloseSend() error {
	log.Debug("CloseSend")
	s.mu.Lock()
	s.closed = true
	s.mu.Unlock()
	if err := s.getStream().CloseSend(); err != nil {
		log.Debug("Received stream error", err)
		return err
	}
	return nil
}

func (s *retryingClientStream) Header() (metadata.MD, error) {
	return s.getStream().Header()
}

func (s *retryingClientStream) Trailer() metadata.MD {
	return s.getStream().Trailer()
}

func (s *retryingClientStream) SendMsg(m interface{}) error {
	log.Debugf("SendMsg %s", m)
	err := s.getStream().SendMsg(m)
	if err == nil {
		s.buffer.append(m)
		return nil
	}

	if err == io.EOF {
		s.mu.RLock()
		closed := s.closed
		s.mu.RUnlock()
		if closed {
			log.Debugf("SendMsg %s: EOF", m)
			return err
		}
	} else if !isRetryable(s.ctx, s.opts, err) {
		log.Warnf("SendMsg %s: error", m, err)
		return err
	}

	log.Debugf("SendMsg %s: error", err)
	err = backoff.Retry(func() error {
		log.Debugf("SendMsg %s: retry", m)
		if err := s.retryStream(); err != nil {
			if err == io.EOF {
				s.mu.RLock()
				closed := s.closed
				s.mu.RUnlock()
				if !closed {
					log.Debugf("SendMsg %s: EOF", m)
					return err
				}
			} else if isRetryable(s.ctx, s.opts, err) {
				log.Debugf("SendMsg %s: error", m, err)
				return err
			}
			log.Warnf("SendMsg %s: error", m, err)
			return backoff.Permanent(err)
		}
		if err := s.getStream().SendMsg(m); err != nil {
			if err == io.EOF {
				s.mu.RLock()
				closed := s.closed
				s.mu.RUnlock()
				if !closed {
					log.Debugf("SendMsg %s: EOF", m)
					return err
				}
			} else if isRetryable(s.ctx, s.opts, err) {
				log.Debugf("SendMsg %s: error", m, err)
				return err
			}
			log.Warnf("SendMsg %s: error", m, err)
			return backoff.Permanent(err)
		}
		return nil
	}, backoff.WithContext(backoff.NewExponentialBackOff(), s.ctx))
	if err == nil {
		s.buffer.append(m)
		return nil
	}
	return err
}

func (s *retryingClientStream) RecvMsg(m interface{}) error {
	if err := s.getStream().RecvMsg(m); err != nil {
		if err == io.EOF {
			log.Debug("RecvMsg: EOF")
			return err
		}
		return backoff.Retry(func() error {
			if err := s.retryStream(); err != nil {
				if isRetryable(s.ctx, s.opts, err) {
					log.Debug("RecvMsg: error", err)
					return err
				}
				log.Warn("RecvMsg: error", err)
				return backoff.Permanent(err)
			}
			if err := s.getStream().RecvMsg(m); err != nil {
				if isRetryable(s.ctx, s.opts, err) {
					log.Debugf("RecvMsg: error", err)
					return err
				}
				log.Warn("RecvMsg: error", err)
				return backoff.Permanent(err)
			}
			return nil
		}, backoff.WithContext(backoff.NewExponentialBackOff(), s.ctx))
	}
	log.Debugf("RecvMsg %s", m)
	return nil
}

func (s *retryingClientStream) retryStream() error {
	b := backoff.NewExponentialBackOff()
	if s.opts.initialInterval != nil {
		b.InitialInterval = *s.opts.initialInterval
	}
	if s.opts.maxInterval != nil {
		b.MaxInterval = *s.opts.maxInterval
	}
	return backoff.Retry(func() error {
		stream, err := s.newStream(s.ctx)
		if err != nil {
			log.Debug("Received stream error", err)
			return err
		}

		s.mu.RLock()
		closed := s.closed
		s.mu.RUnlock()
		msgs := s.buffer.list()
		for _, m := range msgs {
			if err := stream.SendMsg(m); err != nil {
				if isRetryable(s.ctx, s.opts, err) {
					log.Debug("Received stream error", err)
					return err
				}
				log.Warn("Received stream error", err)
				return backoff.Permanent(err)
			}
		}

		if closed {
			if err := stream.CloseSend(); err != nil {
				if isRetryable(s.ctx, s.opts, err) {
					log.Debug("Received stream error", err)
					return err
				}
				log.Warn("Received stream error", err)
				return backoff.Permanent(err)
			}
		}

		s.setStream(stream)
		return nil
	}, backoff.WithContext(b, s.ctx))
}

func isRetryable(ctx context.Context, opts *callOptions, err error) bool {
	st := status.Code(err)
	if opts.perCallTimeout != nil && st == codes.DeadlineExceeded {
		return ctx.Err() == nil
	}
	for _, code := range opts.codes {
		if st == code {
			return true
		}
	}
	return false
}
