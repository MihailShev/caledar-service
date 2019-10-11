package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"time"
)

func interceptorWithLogger(logger grpclog.LoggerV2) grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {

		start := time.Now()

		h, err := handler(ctx, req)

		logger.Infof("Request - Method:%s\tDuration:%s\tError:%v\n",
			info.FullMethod,
			time.Since(start),
			err)

		return h, err
	}
}
