package dsoteldecorator

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	ds "github.com/ipfs/go-datastore"
	dstest "github.com/ipfs/go-datastore/test"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"

	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// Creates an in-memory opentelemetry exporter and runs the whole test suite
// using a known in-memory datastore.
// This doesn't actually test that the telemetry is correct, just that something is happening.
func TestSuiteOtel(t *testing.T) {
	buf := &bytes.Buffer{}
	exp, err := stdouttrace.New(stdouttrace.WithWriter(buf))
	if err != nil {
		t.Fatal(err)
	}
	r, err := resource.Merge(resource.Default(), resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName("test"),
		semconv.ServiceVersion("0.0.1"),
		attribute.String("environment", "test"),
	))
	tp := tracesdk.NewTracerProvider(tracesdk.WithBatcher(exp), tracesdk.WithResource(r))
	tracer := tp.Tracer("test")
	mds := ds.NewMapDatastore()
	w := Wrap(mds, tracer)
	dstest.SubtestAll(t, w)
	var ntraces int
	for {
		b, err := buf.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		assert.Greater(t, len(b), 0)
		ntraces++
	}
	assert.Greater(t, ntraces, 10000)
}
