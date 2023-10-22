package middleware

import (
	"context"

	"babycare/internal/common"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
)

func AddTraceToRequest() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var traceId string
			if trace, ok := tracing.TraceID()(ctx).(string); ok {
				traceId = trace
			}
			if header, ok := transport.FromServerContext(ctx); ok {
				header.RequestHeader().Set(common.TraceId, traceId)
			}
			return handler(ctx, req)
		}
	}
}
