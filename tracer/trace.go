package tracer

import (
	"github.com/diez37/go-packages/app"
	"github.com/diez37/go-packages/configurator"
	"github.com/diez37/go-packages/log"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

func WithConfigurator(configurator configurator.Configurator, config *Config, appConfig *app.Config, informer log.Informer) (trace.Tracer, error) {
	if jaegerDSN := configurator.GetString(JaegerDSNFieldName); jaegerDSN != "" && config.JaegerDSN == "" {
		config.JaegerDSN = jaegerDSN
	}

	return NewTrace(config, appConfig, informer)
}

// NewTrace creating and configuration instance of traceSdk.TracerProvider for send trace to Jaeger
// and return instance of trace.Tracer for application name
func NewTrace(config *Config, appConfig *app.Config, informer log.Informer) (trace.Tracer, error) {
	tracerProvider := traceSdk.NewTracerProvider(
		traceSdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(appConfig.Name),
		)),
	)

	if config.JaegerDSN != "" {
		informer.Infof("tracer: jaeger - %s", config.JaegerDSN)

		exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(config.JaegerDSN)))
		if err != nil {
			return nil, err
		}
		tracerProvider.RegisterSpanProcessor(traceSdk.NewBatchSpanProcessor(exporter))
	} else {
		informer.Info("tracer: empty")
	}

	return tracerProvider.Tracer(appConfig.Name), nil
}
