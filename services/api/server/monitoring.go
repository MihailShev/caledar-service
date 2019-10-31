package main

import (
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net/http"
)

func enableMonitoring(server *grpc.Server, addr string, logger grpclog.LoggerV2) {
	grpc_prometheus.Register(server)
	grpc_prometheus.EnableHandlingTimeHistogram()

	logger.Infof("Monitoring export listen %s", addr)
	err := http.ListenAndServe(addr, promhttp.Handler())

	logger.Error(err)
}
