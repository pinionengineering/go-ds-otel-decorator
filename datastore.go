package dsoteldecorator

// Wraps a datastore with OpenTelemetry instrumentation.

import (
	"context"

	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"

	"go.opentelemetry.io/otel/trace"
)

var (
	_ datastore.Datastore = (*OtelDatastore)(nil)
)

type OtelDatastore struct {
	ds     datastore.Datastore
	tracer trace.Tracer
}

func Wrap(ds datastore.Datastore, tracer trace.Tracer) *OtelDatastore {
	return &OtelDatastore{
		ds:     ds,
		tracer: tracer,
	}
}

func (o *OtelDatastore) Get(ctx context.Context, key datastore.Key) (value []byte, err error) {
	nctx, span := o.tracer.Start(ctx, "datastore.Get")
	defer span.End()
	return o.ds.Get(nctx, key)
}

func (o *OtelDatastore) Has(ctx context.Context, key datastore.Key) (exists bool, err error) {
	nctx, span := o.tracer.Start(ctx, "datastore.Has")
	defer span.End()
	return o.ds.Has(nctx, key)
}

func (o *OtelDatastore) GetSize(ctx context.Context, key datastore.Key) (size int, err error) {
	nctx, span := o.tracer.Start(ctx, "datastore.GetSize")
	defer span.End()
	return o.ds.GetSize(nctx, key)
}

func (o *OtelDatastore) Query(ctx context.Context, q query.Query) (query.Results, error) {
	nctx, span := o.tracer.Start(ctx, "datastore.Query")
	defer span.End()
	return o.ds.Query(nctx, q)
}

func (o *OtelDatastore) Put(ctx context.Context, key datastore.Key, value []byte) error {
	nctx, span := o.tracer.Start(ctx, "datastore.Put")
	defer span.End()
	return o.ds.Put(nctx, key, value)
}

func (o *OtelDatastore) Delete(ctx context.Context, key datastore.Key) error {
	nctx, span := o.tracer.Start(ctx, "datastore.Delete")
	defer span.End()
	return o.ds.Delete(nctx, key)
}
func (o *OtelDatastore) Sync(ctx context.Context, prefix datastore.Key) error {
	nctx, span := o.tracer.Start(ctx, "datastore.Sync")
	defer span.End()
	return o.ds.Sync(nctx, prefix)
}

func (o *OtelDatastore) Close() error {
	return o.ds.Close()
}
