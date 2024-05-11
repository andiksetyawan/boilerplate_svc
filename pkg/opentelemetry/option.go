package opentelemetry

import (
	"github.com/andiksetyawan/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

type OptFunc func(o *Otel) error

// WithJaegerTracerProvider
// Deprecated
//
// docker container:
//
// url:http://localhost:14268/api/traces
//
//	docker run -d --name jaeger \
//		-e COLLECTOR_OTLP_ENABLED=true \
//		-p 16686:16686 \
//		-p 4317:4317 \
//		-p 4318:4318 \
//		-p 14268:14268 \
//		jaegertracing/all-in-one:latest
func WithJaegerTracerProvider(url string) OptFunc {
	return func(o *Otel) (err error) {
		tp, err := o.initJaegerTraceProvider(url)
		o.traceProvider = tp
		return
	}
}

func WithServiceName(name string) OptFunc {
	return func(o *Otel) (err error) {
		o.serviceName = name
		return
	}
}

func WithLogger(log log.Logger) OptFunc {
	return func(o *Otel) (err error) {
		o.logger = log
		return
	}
}

func WithDefaultTraceProvider() OptFunc {
	return func(o *Otel) (err error) {
		tp, err := o.initTraceProvider()
		o.traceProvider = tp
		return
	}
}

func WithDefaultMeterProvider() OptFunc {
	return func(o *Otel) (err error) {
		mp, err := o.initMeterProvider()
		o.meterProvider = mp
		return
	}
}

func WithTraceProvider(tp *trace.TracerProvider) OptFunc {
	return func(o *Otel) (err error) {
		o.traceProvider = tp
		return
	}
}

func WithMeterProvider(mp *metric.MeterProvider) OptFunc {
	return func(o *Otel) (err error) {
		o.meterProvider = mp
		return
	}
}
