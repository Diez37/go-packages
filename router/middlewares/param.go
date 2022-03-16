package middlewares

import (
	"context"
	"github.com/diez37/go-packages/log"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"net/http"
)

type Option func(param *param) *param

type Caster func(string) (interface{}, error)

func WithUri(name string) Option {
	return func(param *param) *param {
		param.uriParamName = name

		return param
	}
}

func WithQuery(name string) Option {
	return func(param *param) *param {
		param.queryParamName = name

		return param
	}
}

func WithHeader(name string) Option {
	return func(param *param) *param {
		param.headerName = name

		return param
	}
}

func WithName(name string) Option {
	return func(param *param) *param {
		param.name = name

		return param
	}
}

type param struct {
	logger log.Logger

	uriParamName   string
	queryParamName string
	headerName     string
	name           string

	caster Caster
}

func NewParam(logger log.Logger, caster Caster, options ...Option) Middleware {
	middleware := &param{logger: logger, caster: caster}

	for _, option := range options {
		option(middleware)
	}

	return middleware
}

func (middleware *param) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		value := request.Header.Get(middleware.headerName)
		if value == "" {
			value = request.URL.Query().Get(middleware.queryParamName)
		}
		if value == "" {
			value = chi.URLParam(request, middleware.uriParamName)
		}

		if value == "" {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			middleware.logger.Errorf(
				"middleware.param: header - '%s', query - '%s', uri - '%s' not found",
				middleware.headerName,
				middleware.queryParamName,
				middleware.uriParamName,
			)
			return
		}

		castValue, err := middleware.caster(value)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			middleware.logger.Error(err)
			return
		}

		ctx := context.WithValue(request.Context(), middleware.name, castValue)
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}

func NewUUID(logger log.Logger, options ...Option) Middleware {
	return NewParam(
		logger,
		func(value string) (interface{}, error) { return uuid.Parse(value) },
		options...,
	)
}

func NewString(logger log.Logger, options ...Option) Middleware {
	return NewParam(
		logger,
		func(value string) (interface{}, error) { return value, nil },
		options...,
	)
}

func NewInt64(logger log.Logger, options ...Option) Middleware {
	return NewParam(
		logger,
		func(value string) (interface{}, error) { return cast.ToInt64E(value) },
		options...,
	)
}

func NewFloat64(logger log.Logger, options ...Option) Middleware {
	return NewParam(
		logger,
		func(value string) (interface{}, error) { return cast.ToFloat64E(value) },
		options...,
	)
}

func NewUint64(logger log.Logger, options ...Option) Middleware {
	return NewParam(
		logger,
		func(value string) (interface{}, error) { return cast.ToUint64E(value) },
		options...,
	)
}

func NewBool(logger log.Logger, options ...Option) Middleware {
	return NewParam(
		logger,
		func(value string) (interface{}, error) { return cast.ToBoolE(value) },
		options...,
	)
}
