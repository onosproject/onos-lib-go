package controller

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"hash/fnv"
	"sync/atomic"
	"testing"
	"time"
)

type testID string

func (id testID) Hash() (int, error) {
	hash := fnv.New32a()
	if _, err := hash.Write([]byte(id)); err != nil {
		return 0, err
	}
	return int(hash.Sum32()), nil
}

func (id testID) String() string {
	return string(id)
}

func TestController(t *testing.T) {
	var reconciler atomic.Value
	controller := NewController[testID](func(ctx context.Context, request Request[testID]) Directive[testID] {
		reconcile := reconciler.Load().(func(ctx context.Context, request Request[testID]) Directive[testID])
		return reconcile(ctx, request)
	}, WithPartitions(10))
	defer controller.Stop()

	done := make(chan struct{})
	reconciler.Store(func(ctx context.Context, request Request[testID]) Directive[testID] {
		close(done)
		return request.Complete()
	})
	assert.NoError(t, controller.Reconcile("foo"))
	<-done

	done = make(chan struct{})
	reconciler.Store(func(ctx context.Context, request Request[testID]) Directive[testID] {
		reconciler.Store(func(ctx context.Context, request Request[testID]) Directive[testID] {
			close(done)
			return request.Complete()
		})
		return request.Retry(errors.New("test"))
	})
	assert.NoError(t, controller.Reconcile("bar"))
	<-done

	done = make(chan struct{})
	reconciler.Store(func(ctx context.Context, request Request[testID]) Directive[testID] {
		reconciler.Store(func(ctx context.Context, request Request[testID]) Directive[testID] {
			close(done)
			return request.Complete()
		})
		return request.Retry(errors.New("test")).With(ExponentialBackoff(time.Second, 10*time.Second))
	})
	assert.NoError(t, controller.Reconcile("baz"))
	<-done

	done = make(chan struct{})
	reconciler.Store(func(ctx context.Context, request Request[testID]) Directive[testID] {
		reconciler.Store(func(ctx context.Context, request Request[testID]) Directive[testID] {
			close(done)
			return request.Complete()
		})
		return request.Retry(errors.New("test")).After(time.Second)
	})
	assert.NoError(t, controller.Reconcile("foo"))
	<-done

	done = make(chan struct{})
	reconciler.Store(func(ctx context.Context, request Request[testID]) Directive[testID] {
		reconciler.Store(func(ctx context.Context, request Request[testID]) Directive[testID] {
			close(done)
			return request.Complete()
		})
		return request.Retry(errors.New("test")).At(time.Now().Add(time.Second))
	})
	assert.NoError(t, controller.Reconcile("bar"))
	<-done
}
