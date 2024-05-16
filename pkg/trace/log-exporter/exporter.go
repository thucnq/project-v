package logexporter

import (
	"context"
	"encoding/json"

	"go.opentelemetry.io/otel/attribute"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"project-v/pkg/l"
)

// Span log span structure
type Span struct {
	// Name The resource name of the span
	Name string `json:"name,omitempty"`
	// SpanId The [SPAN_ID] portion of the span's resource name.
	SpanId string `json:"span_id,omitempty"`
	// ParentSpanId The [SPAN_ID] of this span's parent span. If this is a root span,
	// then this field must be empty.
	ParentSpanId string `json:"parent_span_id,omitempty"`
	// DisplayName A description of the span's operation (up to 128 bytes).
	// For example, the display name can be a qualified method name or a file name
	// and a line number where the operation is called. A best practice is to use
	// the same display name within an application and at the same call point.
	// This makes it easier to correlate spans in different traces.
	DisplayName string `json:"display_name,omitempty"`
	// StartTime The start time of the span. On the client side, this is the time kept by
	// the local machine where the span execution starts. On the server side, this
	// is the time when the server's application handler starts running.
	StartTime string `json:"start_time,omitempty"`
	// EndTime The end time of the span. On the client side, this is the time kept by
	// the local machine where the span execution ends. On the server side, this
	// is the time when the server application handler stops running.
	EndTime string `json:"end_time,omitempty"`
	// Attributes A set of attributes on the span. You can have up to 32 attributes per span.
	Attributes []attribute.KeyValue `json:"attributes,omitempty"`

	Duration string           `json:"duration"`
	Events   []sdktrace.Event `json:"events"`
	Links    []sdktrace.Link  `json:"link"`
	// ChildSpanCount The number of child spans that were generated while this span
	// was active.
	ChildSpanCount int  `json:"child_span_count,omitempty"`
	IsError        bool `json:"is_error,omitempty"`
}

// LogExporter is a log exporter that implement of SpanExporter.
// this exporter will print the span data to the log output. default is stdout
type LogExporter struct {
	l l.Logger
}

// ExportSpans ...exports a batch of spans to the log output.
func (e *LogExporter) ExportSpans(
	ctx context.Context, spans []sdktrace.ReadOnlySpan,
) error {
	results := make([]*Span, len(spans))
	for i, sd := range spans {
		results[i] = e.ConvertSpan(ctx, sd)
	}

	out, err := json.Marshal(results)
	if err != nil {
		e.l.Error("failed to export span", l.Error(err))
		return err
	}
	var fields []map[string]interface{}
	err = json.Unmarshal(out, &fields)
	if err != nil {
		e.l.Error("failed to export span", l.Error(err))
		return err
	}
	for _, item := range fields {
		displayName, ok := item["display_name"].(string)
		if !ok || len(displayName) == 0 {
			displayName = item["name"].(string)
		}
		isError, ok := item["is_error"].(bool)
		if isError {
			e.l.Error(displayName, l.Any("item", item))
		} else {
			e.l.Info(displayName, l.Any("item", item))
		}
		// e.l.Info(item, displayName)
	}
	return nil
}

// ConvertSpan converts a ReadOnlySpan to log Span.
func (e *LogExporter) ConvertSpan(
	_ context.Context, sd sdktrace.ReadOnlySpan,
) *Span {
	return protoFromReadOnlySpan(sd)
}

func (e *LogExporter) Shutdown(ctx context.Context) error {
	return nil
}

// If there are duplicate keys present in the list of attributes,
// then the first value found for the key is preserved.
func attributeWithLabelsFromResources(sd sdktrace.ReadOnlySpan) []attribute.KeyValue {
	ignoreKey := map[attribute.Key]struct{}{
		"telemetry.sdk.name":     {},
		"telemetry.sdk.language": {},
		"telemetry.sdk.version":  {},
		"telemetry.auto.version": {},
	}
	attributes := sd.Attributes()
	if sd.Resource().Len() == 0 {
		return attributes
	}
	uniqueAttrs := make(map[attribute.Key]bool, len(sd.Attributes()))
	for _, attr := range sd.Attributes() {
		uniqueAttrs[attr.Key] = true
	}
	for _, attr := range sd.Resource().Attributes() {
		if _, ig := ignoreKey[attr.Key]; ig {
			continue
		}
		if uniqueAttrs[attr.Key] {
			continue // skip resource attributes which conflict with span attributes
		}
		uniqueAttrs[attr.Key] = true
		attributes = append(attributes, attr)
	}

	return attributes
}

// protoFromReadOnlySpan ...
func protoFromReadOnlySpan(s sdktrace.ReadOnlySpan) *Span {
	if s == nil {
		return nil
	}

	traceIDString := s.SpanContext().TraceID().String()
	spanIDString := s.SpanContext().SpanID().String()

	sp := &Span{
		Name:           "traces/" + traceIDString + "/spans/" + spanIDString,
		SpanId:         spanIDString,
		DisplayName:    s.Name(),
		StartTime:      s.StartTime().String(),
		EndTime:        s.EndTime().String(),
		ChildSpanCount: s.ChildSpanCount(),
		IsError:        int(s.Status().Code) == 1,
		Duration:       s.EndTime().Sub(s.StartTime()).String(),
	}
	if s.Parent().SpanID() != s.SpanContext().SpanID() && s.Parent().SpanID().IsValid() {
		sp.ParentSpanId = s.Parent().SpanID().String()
	}

	sp.Attributes = attributeWithLabelsFromResources(s)

	sp.Events = s.Events()
	sp.Links = s.Links()

	return sp
}

// New creates a new log exporter
func New(ll l.Logger) *LogExporter {
	return &LogExporter{
		l: ll,
	}
}
