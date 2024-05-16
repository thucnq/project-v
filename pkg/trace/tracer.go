package tracer

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"

	"project-v/pkg/l"
	logexporter "project-v/pkg/trace/log-exporter"
)

type tracerImpl struct {
	tp     trace.TracerProvider
	tr     trace.Tracer
	enable bool
}

var DefaultTracer *tracerImpl

type TelExporter struct {
	ID    int64
	Token string
}

type Option struct {
	TracerName  string
	L           l.Logger
	TelExporter *TelExporter
	Enable      bool
}

// New creates a new tracer
func New(
	enable bool, projectID string, TracerName string, logger l.Logger,
) *tracerImpl {
	os.Setenv("OTEL_SERVICE_NAME", TracerName)

	if !enable {
		initEmptyTrace()
		return DefaultTracer
	}

	lExporter := logexporter.New(logger)
	tp := sdktrace.NewTracerProvider(
		// For this example code we use sdktrace.AlwaysSample sampler to sample all traces.
		// In a production application, use sdktrace. ProbabilitySampler with a desired probability.
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(lExporter),
	)
	otel.SetTracerProvider(tp)
	tr := otel.Tracer(TracerName)
	DefaultTracer = &tracerImpl{
		tp:     tp,
		tr:     tr,
		enable: enable,
	}
	return DefaultTracer
}

// NewWithOption creates a new tracer
func NewWithOption(op *Option) *tracerImpl {
	os.Setenv("OTEL_SERVICE_NAME", op.TracerName)

	if !op.Enable {
		initEmptyTrace()
		return DefaultTracer
	}

	lExporter := logexporter.New(op.L)
	p := []sdktrace.TracerProviderOption{
		// For this example code we use sdktrace.AlwaysSample sampler to sample all traces.
		// In a production application, use sdktrace. ProbabilitySampler with a desired probability.
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(lExporter),
	}
	// if len(op.TelExporter.Token) > 0 {
	// 	telExporter := telexporter.New(
	// 		op.L, op.TelExporter.ID, op.TelExporter.Token,
	// 		telegram.NewTeleNotiService(op.TelExporter.Token),
	// 	)
	// 	p = append(p, sdktrace.WithBatcher(telExporter))
	// }

	tp := sdktrace.NewTracerProvider(
		p...,
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{}, propagation.Baggage{},
		),
	)

	tr := otel.Tracer(op.TracerName)
	DefaultTracer = &tracerImpl{
		tp:     tp,
		tr:     tr,
		enable: op.Enable,
	}
	return DefaultTracer
}

func Shutdown() {
	if DefaultTracer == nil || !DefaultTracer.enable {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = DefaultTracer.tp.(*sdktrace.TracerProvider).Shutdown(ctx)
}

func initEmptyTrace() {
	tp := trace.NewNoopTracerProvider()
	DefaultTracer = &tracerImpl{
		tp:     tp,
		tr:     tp.Tracer(""),
		enable: false,
	}
}

// StartSpan creates a span and context from exist context
func StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	if DefaultTracer == nil {
		initEmptyTrace()
	}
	return DefaultTracer.tr.Start(ctx, name)
}

// SetSpanAttributes Sets span attributes for input key-value pairs
func SetSpanAttributes(span trace.Span, input map[string]string) {
	for key, value := range input {
		span.SetAttributes(attribute.String(key, value))
	}
}
func SetTraceReq(span trace.Span, req interface{}) {
	span.SetAttributes(attribute.String("req", fmt.Sprintf("%+v", req)))
}
func SetAttribute(span trace.Span, key string, data interface{}) {
	span.SetAttributes(attribute.String(key, fmt.Sprintf("%+v", data)))
}
func SetError(span trace.Span, err error) {
	span.RecordError(err)
	span.SetStatus(1, err.Error())
}

func SetWarnDuration(span trace.Span, d time.Duration) {

}

func StartEmptySpan(ctx context.Context, name string) (
	context.Context, trace.Span,
) {
	tp := trace.NewNoopTracerProvider()
	return tp.Tracer("").Start(ctx, name)
}
