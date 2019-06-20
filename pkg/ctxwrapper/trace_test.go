package ctxwrapper

import (
	"context"
	"reflect"
	"testing"
)

func TestSetTraceID(t *testing.T) {
	type args struct {
		ctx     context.Context
		traceID string
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "1. Set empty string",
			args: args{
				ctx:     context.Background(),
				traceID: "",
			},
			want: context.WithValue(context.Background(), contextKeyTraceID, ""),
		},
		{
			name: "2. Set non-empty string",
			args: args{
				ctx:     context.Background(),
				traceID: "traceID",
			},
			want: context.WithValue(context.Background(), contextKeyTraceID, "traceID"),
		},
		{
			name: "3. Set empty string to non empty context",
			args: args{
				ctx:     context.WithValue(context.Background(), contextKeyTraceID, "traceID-1234"),
				traceID: "",
			},
			want: context.WithValue(context.Background(), contextKeyTraceID, ""),
		},
		{
			name: "4. Set non-empty string to non-empty context",
			args: args{
				ctx:     context.WithValue(context.Background(), contextKeyTraceID, "traceID-1234"),
				traceID: "traceID",
			},
			want: context.WithValue(context.Background(), contextKeyTraceID, "traceID"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetTraceID(tt.args.ctx, tt.args.traceID)
			if !reflect.DeepEqual(got.Value(contextKeyTraceID), tt.want.Value(contextKeyTraceID)) {
				t.Errorf(
					"SetTraceID() = %v, want %v",
					got.Value(contextKeyTraceID),
					tt.want.Value(contextKeyTraceID),
				)
			}
		})
	}
}

func TestTraceID(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1. Get traceID from default context background",
			args: args{
				ctx: context.Background(),
			},
			want: "",
		},
		{
			name: "2. Get traceID empty string",
			args: args{
				ctx: context.WithValue(context.Background(), contextKeyTraceID, ""),
			},
			want: "",
		},
		{
			name: "3. Get traceID non-empty string",
			args: args{
				ctx: context.WithValue(context.Background(), contextKeyTraceID, "traceID"),
			},
			want: "traceID",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TraceID(tt.args.ctx); got != tt.want {
				t.Errorf("TraceID() = %v, want %v", got, tt.want)
			}
		})
	}
}
