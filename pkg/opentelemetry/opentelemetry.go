package opentelemetry

import (
	"context"
	"errors"
	"time"

	"github.com/andiksetyawan/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
)

type Otel struct {
	serviceName string
	logger      log.Logger
	shutdownFn  []func(ctx context.Context) error

	traceProvider *trace.TracerProvider
	meterProvider *metric.MeterProvider
}

func New(opts ...OptFunc) (ot *Otel, err error) {
	ot = new(Otel)

	for _, opt := range opts {
		if err = opt(ot); err != nil {
			return
		}
	}

	if ot.serviceName == "" {
		err = errors.New("missing service name")
		return
	}

	if ot.logger == nil {
		err = errors.New("missing logger")
		return
	}

	if ot.traceProvider == nil {
		traceProvider, err := ot.initTraceProvider()
		if err != nil {
			return nil, err
		}
		ot.traceProvider = traceProvider
	}

	if ot.meterProvider == nil {
		meterProvider, err := ot.initMeterProvider()
		if err != nil {
			return nil, err
		}
		ot.meterProvider = meterProvider
	}

	return ot, err
}

func (o *Otel) initTraceProvider() (*trace.TracerProvider, error) {
	traceExporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter,
			// Default is 5s. Set to 1s for demonstrative purposes.
			trace.WithBatchTimeout(time.Second)),
	)

	return traceProvider, nil
}

func (o *Otel) initMeterProvider() (*metric.MeterProvider, error) {
	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter,
			// Default is 1m. Set to 3s for demonstrative purposes.
			metric.WithInterval(3*time.Second))),
	)

	return meterProvider, nil
}

func (o *Otel) Setup() {
	if o.traceProvider != nil {
		otel.SetTracerProvider(o.traceProvider)
		o.shutdownFn = append(o.shutdownFn, o.traceProvider.Shutdown)
	}

	if o.meterProvider != nil {
		otel.SetMeterProvider(o.meterProvider)
		o.shutdownFn = append(o.shutdownFn, o.meterProvider.Shutdown)
	}

	return
}

func (o *Otel) initJaegerTraceProvider(url string) (*trace.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	traceProvider := trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exp),
		// Record information about this application in a Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(o.serviceName),
		)),
	)

	return traceProvider, nil
}

func (o *Otel) Close(ctx context.Context) (err error) {
	for _, fn := range o.shutdownFn {
		err = errors.Join(err, fn(ctx))
	}

	if err != nil {
		return
	}

	o.logger.Info(ctx, "opentelemetry is shutdown correctly")
	return
}
