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
	"time"
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
	connOpts := newCallOptions(defaultOptions, callOpts)
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		grpcOpts, retryOpts := filterCallOptions(opts)
		callOpts := newCallOptions(connOpts, retryOpts)
		b := backoff.NewExponentialBackOff()
		if callOpts.initialInterval != nil {
			b.InitialInterval = *callOpts.initialInterval
		}
		if callOpts.maxInterval != nil {
			b.MaxInterval = *callOpts.maxInterval
		}
		return backoff.Retry(func() error {
			log.Debugf("SendMsg %s", req)
			callCtx := newCallContext(ctx, callOpts)
			if err := invoker(callCtx, method, req, reply, cc, grpcOpts...); err != nil {
				if isContextError(err) {
					if ctx.Err() != nil {
						log.Debugf("SendMsg %s: error", req, err)
						return backoff.Permanent(err)
					} else if callOpts.perCallTimeout != nil {
						log.Debugf("SendMsg %s: error", req, err)
						return err
					}
				}
				if isRetryable(callOpts, err) {
					log.Debugf("SendMsg %s: error", req, err)
					return err
				}
				log.Warnf("SendMsg %s: error", req, err)
				return backoff.Permanent(err)
			}
			log.Debugf("RecvMsg %s", reply)
			return nil
		}, b)
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
	connOpts := newCallOptions(defaultOptions, callOpts)
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		grpcOpts, retryOpts := filterCallOptions(opts)
		callOpts := newCallOptions(connOpts, retryOpts)
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
	connOpts := newCallOptions(defaultOptions, callOpts)
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		grpcOpts, retryOpts := filterCallOptions(opts)
		callOpts := newCallOptions(connOpts, retryOpts)
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
	connOpts := newCallOptions(defaultOptions, callOpts)
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		grpcOpts, retryOpts := filterCallOptions(opts)
		callOpts := newCallOptions(connOpts, retryOpts)
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
		log.Warn("CloseSend: error", err)
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
	return s.retrySendMsg(m)
}

func (s *retryingClientStream) retrySendMsg(m interface{}) error {
	return backoff.RetryNotify(func() error {
		return s.trySendMsg(m)
	}, backoff.NewExponentialBackOff(), func(err error, duration time.Duration) {
		log.Debugf("SendMsg %s: retry after %s", m, duration, err)
	})
}

func (s *retryingClientStream) trySendMsg(m interface{}) error {
	err := s.getStream().SendMsg(m)
	if err == nil {
		s.buffer.append(m)
		return nil
	}
	if isContextError(err) {
		if s.ctx.Err() != nil {
			log.Debugf("SendMsg %s: error", m, err)
			return backoff.Permanent(err)
		} else if s.opts.perCallTimeout != nil {
			log.Debugf("SendMsg %s: error", m, err)
			return s.tryStream()
		}
	}
	if isRetryable(s.opts, err) {
		log.Debugf("SendMsg %s: error", m, err)
		return s.tryStream()
	}
	log.Warnf("SendMsg %s: error", m, err)
	return backoff.Permanent(err)
}

func (s *retryingClientStream) RecvMsg(m interface{}) error {
	return s.retryRecvMsg(m)
}

func (s *retryingClientStream) retryRecvMsg(m interface{}) error {
	return backoff.RetryNotify(func() error {
		return s.tryRecvMsg(m)
	}, backoff.NewExponentialBackOff(), func(err error, duration time.Duration) {
		log.Debugf("RecvMsg: retry after %s", duration, err)
	})
}

func (s *retryingClientStream) tryRecvMsg(m interface{}) error {
	err := s.getStream().RecvMsg(m)
	if err == nil {
		log.Debugf("RecvMsg %s", m)
		return nil
	}
	if err == io.EOF {
		log.Debug("RecvMsg: EOF")
		return backoff.Permanent(err)
	}
	if isContextError(err) {
		if s.ctx.Err() != nil {
			log.Debug("RecvMsg: error", err)
			return backoff.Permanent(err)
		} else if s.opts.perCallTimeout != nil {
			log.Debug("RecvMsg: error", err)
			return s.tryStream()
		}
	}
	if isRetryable(s.opts, err) {
		log.Debug("RecvMsg: error", err)
		return s.tryStream()
	}
	log.Warn("RecvMsg: error", err)
	return backoff.Permanent(err)
}

func (s *retryingClientStream) retryStream() error {
	b := backoff.NewExponentialBackOff()
	if s.opts.initialInterval != nil {
		b.InitialInterval = *s.opts.initialInterval
	}
	if s.opts.maxInterval != nil {
		b.MaxInterval = *s.opts.maxInterval
	}
	return backoff.RetryNotify(s.tryStream, backoff.NewExponentialBackOff(), func(err error, duration time.Duration) {
		log.Debugf("Stream: retry after %s", duration, err)
	})
}

func (s *retryingClientStream) tryStream() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	stream, err := s.newStream(newCallContext(s.ctx, s.opts))
	if err != nil {
		if isContextError(err) {
			if s.ctx.Err() != nil {
				log.Debug("Stream: error", err)
				return backoff.Permanent(err)
			} else if s.opts.perCallTimeout != nil {
				log.Debug("Stream: error", err)
				return err
			}
		}
		if isRetryable(s.opts, err) {
			log.Debug("Stream: error", err)
			return err
		}
		log.Warn("Stream: error", err)
		return backoff.Permanent(err)
	}

	msgs := s.buffer.list()
	for _, m := range msgs {
		log.Debugf("SendMsg %s", m)
		if err := stream.SendMsg(m); err != nil {
			if isContextError(err) {
				if s.ctx.Err() != nil {
					log.Debugf("SendMsg %s: error", m, err)
					return backoff.Permanent(err)
				} else if s.opts.perCallTimeout != nil {
					log.Debugf("SendMsg %s: error", m, err)
					return err
				}
			}
			if isRetryable(s.opts, err) {
				log.Debugf("SendMsg %s: error", m, err)
				return err
			}
			log.Warnf("SendMsg %s: error", m, err)
			return backoff.Permanent(err)
		}
	}

	if s.closed {
		log.Debug("CloseSend")
		if err := stream.CloseSend(); err != nil {
			if isContextError(err) {
				if s.ctx.Err() != nil {
					log.Debug("CloseSend: error", err)
					return backoff.Permanent(err)
				} else if s.opts.perCallTimeout != nil {
					log.Debug("CloseSend: error", err)
					return err
				}
			}
			if isRetryable(s.opts, err) {
				log.Debug("CloseSend: error", err)
				return err
			}
			log.Warn("CloseSend: error", err)
			return backoff.Permanent(err)
		}
	}
	s.stream = stream
	return nil
}

func isContextError(err error) bool {
	code := status.Code(err)
	return code == codes.DeadlineExceeded || code == codes.Canceled
}

func isRetryable(opts *callOptions, err error) bool {
	code := status.Code(err)
	if code == codes.Canceled || code == codes.DeadlineExceeded {
		return false
	}
	for _, retryableCode := range opts.codes {
		if code == retryableCode {
			return true
		}
	}
	return false
}
