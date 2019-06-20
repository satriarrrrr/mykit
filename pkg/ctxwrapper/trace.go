package ctxwrapper

import "context"

var (
	contextKeyTraceID = contextKey("traceID")
)

// TraceID return value of traceID attached in the context
// Empty string will be returned if no traceID attached in the context
func TraceID(ctx context.Context) string {
	traceID := ctx.Value(contextKeyTraceID)
	if traceID == nil {
		return ""
	}
	return traceID.(string)
}

// SetTraceID return context with value traceID attached on them
// Accept empty string as the value of traceID
func SetTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, contextKeyTraceID, traceID)
}
