package log

import (
	"github.com/diez37/go-packages/app"
	"github.com/diez37/go-packages/configurator"
	"github.com/diez37/go-packages/metrics"
	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"
	"gopkg.in/errgo.v2/errors"
	"log"
)

func WithConfigurator(config *Config, appConfig *app.Config, configurator configurator.Configurator, metrics *metrics.Metrics) (*logrus.Logger, error) {
	configurator.SetDefault(InfoLevelFieldName, VerboseDefault)
	if verbose := configurator.GetBool(InfoLevelFieldName); verbose != VerboseDefault && config.InfoLevel == VerboseDefault {
		config.InfoLevel = verbose
	}

	configurator.SetDefault(DebugLevelFieldName, VerboseDefault)
	if verbose := configurator.GetBool(DebugLevelFieldName); verbose != VerboseDefault && config.DebugLevel == VerboseDefault {
		config.DebugLevel = verbose
	}

	if sentryDSN := configurator.GetString(SentryDSNFieldName); sentryDSN != "" && config.SentryDSN == "" {
		config.SentryDSN = sentryDSN
	}

	return NewLogrus(config, appConfig, metrics)
}

// NewLogrus creating and configuration instance of logrus.Logger which implements Logger
func NewLogrus(config *Config, appConfig *app.Config, metrics *metrics.Metrics) (*logrus.Logger, error) {
	logger := logrus.New()

	logger.SetLevel(logrus.ErrorLevel)
	logger.AddHook(&metricHook{metrics: metrics})

	if config.InfoLevel {
		logger.SetLevel(logrus.InfoLevel)
	}

	if config.DebugLevel {
		logger.SetLevel(logrus.DebugLevel)
	}

	logger.SetFormatter(&logrus.TextFormatter{})

	// if Sentry dsn not equal empty string then
	// adding hook to logrus for sending errors and warnings to Sentry
	if config.SentryDSN != "" {
		client, err := raven.New(config.SentryDSN)
		if err != nil {
			return nil, errors.Wrap(err)
		}

		hook, err := logrus_sentry.NewWithClientSentryHook(client, []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel})
		if err != nil {
			return nil, errors.Wrap(err)
		}

		logger.Hooks.Add(hook)
	}

	log.SetOutput(logger.Writer())

	return logger, nil
}

type metricHook struct {
	metrics *metrics.Metrics
}

func (hook *metricHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (hook *metricHook) Fire(entry *logrus.Entry) error {
	switch entry.Level {
	case logrus.ErrorLevel:
		hook.metrics.ErrorsCounter.Inc()
	case logrus.WarnLevel:
		hook.metrics.WarningsCounter.Inc()
	}

	return nil
}
