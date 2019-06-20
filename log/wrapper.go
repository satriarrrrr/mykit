package log

import (
	"context"

	internallog "github.com/satriarrrrr/mykit/internal/log"
	"github.com/satriarrrrr/mykit/pkg/ctxwrapper"
)

// Wrap wrap several data as a list of fields
func Wrap(ctx context.Context, event string, data interface{}, err error) []internallog.Field {
	fields := []internallog.Field{}

	fields = append(fields, internallog.String("trace_id", ctxwrapper.TraceID(ctx)))
	fields = append(fields, internallog.String("event", event))
	fields = append(fields, internallog.Any("data", data))

	if err != nil {
		fields = append(fields, internallog.Error(err))
	}

	return fields
}
