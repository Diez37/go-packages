package router

import (
	"compress/gzip"
	"github.com/diez37/go-packages/configurator"
	"github.com/diez37/go-packages/log"
	"github.com/diez37/go-packages/metrics"
	"github.com/diez37/go-packages/router/middlewares"
	"github.com/diez37/go-packages/tracer"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func WithConfigurator(configurator configurator.Configurator, traceConfig *tracer.Config, metrics *metrics.Metrics, logger log.Logger) chi.Router {
	configurator.SetDefault(tracer.UseProfileFieldName, tracer.UseProfileDefault)
	if isUseProfile := configurator.GetBool(tracer.UseProfileFieldName); isUseProfile != tracer.UseProfileDefault && traceConfig.UseProfile == tracer.UseProfileDefault {
		traceConfig.UseProfile = isUseProfile
	}

	return NewRouter(traceConfig, metrics, logger)
}

// NewRouter creating and configuration instance of mux.Router
func NewRouter(traceConfig *tracer.Config, metrics *metrics.Metrics, logger log.Logger) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Compress(gzip.BestSpeed))
	router.Use(middleware.RequestID)
	router.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger, NoColor: true}))
	router.Use(middleware.Recoverer)
	router.Use(middlewares.HttpDurationMiddleware(metrics, logger))
	router.Use(middlewares.HttpRequestTotalMiddleware(metrics, logger))

	logger.Info("router: add handle health")
	router.Use(middleware.Heartbeat("/health-check"))

	logger.Info("router: add handle metrics")
	router.Mount("/metrics", promhttp.Handler())

	if traceConfig.UseProfile {
		logger.Info("router: add handles pprof")

		router.Mount("/debug", middleware.Profiler())
	}

	return router
}
