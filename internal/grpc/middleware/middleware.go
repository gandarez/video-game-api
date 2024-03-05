package middleware

import (
	"log/slog"
	"runtime/debug"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// WithPanicRecovery returns a new grpc.ServerOption that recovers from panics and logs them.
func WithPanicRecovery(logger *slog.Logger) grpc.ServerOption {
	return grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(
				grpc_recovery.WithRecoveryHandler(func(p any) (err error) {
					logger.Error("service panicked",
						slog.Any("error", err),
						slog.String("stack", string(debug.Stack())),
					)

					return status.Errorf(codes.Internal, "service panicked: %s", p)
				}),
			),
		),
	)
}

// WithUnknownServiceHandler returns a new grpc.ServerOption that handles unknown routes.
func WithUnknownServiceHandler(logger *slog.Logger) grpc.ServerOption {
	return grpc.UnknownServiceHandler(func(_ any, stream grpc.ServerStream) error {
		m, _ := grpc.MethodFromServerStream(stream)

		logger.Error("unknown method in request", slog.String("method", m))

		return status.Errorf(codes.Unimplemented, "unknown route")
	})
}
