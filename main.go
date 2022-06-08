package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"discount_service/internal/endpoint"
	"discount_service/internal/pkg"
	"discount_service/internal/repository"
	"discount_service/internal/transport"
	"discount_service/internal/webapi"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func main() {
	var httpAddr = flag.String("http", ":8086", "http listen address")

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "discount",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	flag.Parse()
	ctx := context.Background()

	var srv pkg.Service
	{
		repository := repository.NewRepo(logger)
		webAPI := webapi.NewWebAPI(logger)
		srv = pkg.NewService(repository, webAPI, logger)
	}

	errs := make(chan error)

	endpoints := endpoint.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := transport.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
