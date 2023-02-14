package controller

import (
	"math"
	"time"
)

type Directive[I ID] interface {
	Do(controller *Controller[I])
}

type Backoff func(attempt int) time.Duration

func ConstantBackoff(delay time.Duration, maxDelay time.Duration) Backoff {
	return func(attempt int) time.Duration {
		return time.Duration(math.Min(float64(int(delay)*attempt), float64(maxDelay)))
	}
}

func ExponentialBackoff(initialDelay time.Duration, maxDelay time.Duration) Backoff {
	return func(attempt int) time.Duration {
		maxExponent := math.Log2(float64(maxDelay) / float64(initialDelay))
		return initialDelay * time.Duration(math.Pow(2, math.Min(float64(attempt), maxExponent)))
	}
}

type Complete[I ID] struct {
	request Request[I]
}

func (c *Complete[I]) Do(controller *Controller[I]) {
	controller.Log.Debugf("Reconciliation of %s complete", c.request.ID)
}

type Requeue[I ID] struct {
	request Request[I]
}

func (r *Requeue[I]) Do(controller *Controller[I]) {
	controller.Log.Debugf("Requeueing %s", r.request.ID)
	go controller.enqueue(r.request)
}

type Fail[I ID] struct {
	request Request[I]
	Error   error
}

func (f *Fail[I]) Do(controller *Controller[I]) {
	controller.Log.Warnf("Reconciliation of %s failed: %s", f.request.ID, f.Error.Error())
}

type Retry[I ID] struct {
	request Request[I]
	Error   error
}

func (r *Retry[I]) Do(controller *Controller[I]) {
	controller.Log.Debugf("Reconciliation of %s failed: %s. Retrying", r.request.ID, r.Error.Error())
	go controller.enqueue(r.request)
}

func (r *Retry[I]) After(delay time.Duration) *RetryAfter[I] {
	return &RetryAfter[I]{
		Retry: r,
		delay: delay,
	}
}

func (r *Retry[I]) At(t time.Time) *RetryAt[I] {
	return &RetryAt[I]{
		Retry: r,
		t:     t,
	}
}

func (r *Retry[I]) With(backoff Backoff) *RetryWith[I] {
	return &RetryWith[I]{
		Retry:   r,
		backoff: backoff,
	}
}

type RetryAfter[I ID] struct {
	*Retry[I]
	delay time.Duration
}

func (r *RetryAfter[I]) Do(controller *Controller[I]) {
	controller.Log.Debugf("Reconciliation of %s failed: %s. Retrying after %s", r.request.ID, r.Error.Error(), r.delay)
	time.AfterFunc(r.delay, func() {
		controller.enqueue(r.request)
	})
}

type RetryAt[I ID] struct {
	*Retry[I]
	t time.Time
}

func (r *RetryAt[I]) Do(controller *Controller[I]) {
	controller.Log.Debugf("Reconciliation of %s failed: %s. Retrying at %s", r.request.ID, r.Error.Error(), r.t)
	time.AfterFunc(time.Until(r.t), func() {
		controller.enqueue(r.request)
	})
}

type RetryWith[I ID] struct {
	*Retry[I]
	backoff Backoff
}

func (r *RetryWith[I]) Do(controller *Controller[I]) {
	delay := r.backoff(r.request.attempt)
	controller.Log.Debugf("Reconciliation of %s failed: %s. Retrying after %s", r.request.ID, r.Error.Error(), delay)
	time.AfterFunc(delay, func() {
		controller.enqueue(r.request)
	})
}
