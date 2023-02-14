// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"context"
	"fmt"
	"hash/fnv"
	"time"

	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

const (
	defaultReconciliationTimeout = 30 * time.Second
)

type ID interface {
	~string | ~int | ~int32 | ~int64 | ~uint | ~uint32 | ~uint64
	fmt.Stringer
}

// Reconciler reconciles an object
type Reconciler[I ID] func(ctx context.Context, request Request[I]) Directive[I]

// Request is a reconciler request
type Request[I ID] struct {
	ID        I
	partition int
	attempt   int
}

func (r Request[I]) Complete() *Complete[I] {
	return &Complete[I]{
		request: r,
	}
}

func (r Request[I]) Requeue() *Requeue[I] {
	return &Requeue[I]{
		request: r,
	}
}

func (r Request[I]) Fail(err error) *Fail[I] {
	return &Fail[I]{
		request: r,
		Error:   err,
	}
}

func (r Request[I]) Retry(err error) *Retry[I] {
	return &Retry[I]{
		request: r,
		Error:   err,
	}
}

type Options struct {
	Log        logging.Logger
	Partitions int
	Timeout    *time.Duration
}

type Option func(*Options)

func WithOptions(options Options) Option {
	return func(opts *Options) {
		*opts = options
	}
}

func WithLog(log logging.Logger) Option {
	return func(options *Options) {
		options.Log = log
	}
}

func WithPartitions(partitions int) Option {
	return func(options *Options) {
		options.Partitions = partitions
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(options *Options) {
		options.Timeout = &timeout
	}
}

// NewController creates a new controller
func NewController[I ID](reconciler Reconciler[I], opts ...Option) *Controller[I] {
	var options Options
	for _, opt := range opts {
		opt(&options)
	}
	numPartitions := options.Partitions
	if numPartitions == 0 {
		numPartitions = 1
	}
	partitions := make([]chan Request[I], numPartitions)
	for i := 0; i < numPartitions; i++ {
		partitions[i] = make(chan Request[I])
	}
	rlog := options.Log
	if rlog == nil {
		rlog = log
	}
	timeout := defaultReconciliationTimeout
	if options.Timeout != nil {
		timeout = *options.Timeout
	}
	controller := &Controller[I]{
		partitions: partitions,
		reconciler: reconciler,
		Log:        rlog,
		timeout:    timeout,
	}
	controller.start()
	return controller
}

// Controller is a control loop
// The Controller is responsible for processing events provided by a Watcher. Events are processed by
// a configurable Reconciler. The controller processes events in a loop, retrying requests until the
// Reconciler can successfully process them.
// The Controller can be activated or deactivated by a configurable Activator. When inactive, the controller
// will ignore requests, and when active it processes all requests.
// For per-request filtering, a Filter can be provided which provides a simple bool to indicate whether a
// request should be passed to the Reconciler.
// Once the Reconciler receives a request, it should process the request using the current state of the cluster
// Reconcilers should not cache state themselves and should instead rely on stores for consistency.
// If a Reconciler returns false, the request will be requeued to be retried after all pending requests.
// If a Reconciler returns an error, the request will be retried after a backoff period.
// Once a Reconciler successfully processes a request by returning true, the request will be discarded.
// Requests can be partitioned among concurrent goroutines by configuring a WorkPartitioner. The controller
// will create a goroutine per PartitionKey provided by the WorkPartitioner, and requests to different
// partitions may be handled concurrently.
type Controller[I ID] struct {
	reconciler Reconciler[I]
	partitions []chan Request[I]
	Log        logging.Logger
	timeout    time.Duration
}

// start starts the controller
func (c *Controller[I]) start() {
	for _, partition := range c.partitions {
		go c.processRequests(partition)
	}
}

// Stop stops the controller
func (c *Controller[I]) Stop() {
	for _, partition := range c.partitions {
		close(partition)
	}
}

func (c *Controller[I]) Reconcile(id I) error {
	hash, err := computeHash(id)
	if err != nil {
		return err
	}
	request := Request[I]{
		ID:        id,
		partition: hash % len(c.partitions),
	}
	go c.enqueue(request)
	return nil
}

func (c *Controller[I]) enqueue(request Request[I]) {
	request.attempt++
	partition := c.partitions[request.partition]
	partition <- request
}

// processRequests reconciles requests from the given channel
func (c *Controller[I]) processRequests(ch chan Request[I]) {
	for request := range ch {
		c.processRequest(request)
	}
}

// processRequest reconciles the given request
func (c *Controller[I]) processRequest(request Request[I]) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	directive := c.reconciler(ctx, request)
	if directive != nil {
		directive.Do(c)
	}
}

func computeHash[I ID](id I) (int, error) {
	hash := fnv.New32a()
	if _, err := hash.Write([]byte(id.String())); err != nil {
		return 0, err
	}
	return int(hash.Sum32()), nil
}
