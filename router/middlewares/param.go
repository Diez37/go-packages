package middlewares

import (
	"context"
	"github.com/diez37/go-packages/log"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"net/http"
	"strings"
)

const (
	PageFieldName  = "page"
	LimitFieldName = "limit"

	CountHeaderName = "X-Pagination-Count"
	PageHeaderName  = "X-Pagination-Page"
	LimitHeaderName = "X-Pagination-Limit"

	LimitDefault = uint(20)
	PageDefault  = uint(1)
)

type Option func(param *param) *param

type Caster func(string) (interface{}, error)

type Getter func(request *http.Request) string

func WithUri(name string) Option {
	return func(param *param) *param {
		param.getters = append(param.getters, func(request *http.Request) string {
			return request.URL.Query().Get(name)
		})

		return param
	}
}

func WithQuery(name string) Option {
	return func(param *param) *param {
		param.getters = append(param.getters, func(request *http.Request) string {
			return chi.URLParam(request, name)
		})

		return param
	}
}

func WithHeader(name string) Option {
	return func(param *param) *param {
		param.getters = append(param.getters, func(request *http.Request) string {
			return request.Header.Get(name)
		})

		return param
	}
}

func WithCookie(name string) Option {
	return func(param *param) *param {
		param.getters = append(param.getters, func(request *http.Request) string {
			for _, cookie := range request.Cookies() {
				if cookie.Name == name {
					return cookie.Value
				}
			}

			return ""
		})

		return param
	}
}

func WithName(name string) Option {
	return func(param *param) *param {
		param.name = name

		return param
	}
}

func WithDefault(value interface{}) Option {
	return func(param *param) *param {
		param._defaultValue = value

		return param
	}
}

type param struct {
	logger log.Logger

	name string

	getters       []Getter
	caster        Caster
	_defaultValue interface{}
}

func NewParam(logger log.Logger, caster Caster, options ...Option) Middleware {
	middleware := &param{logger: logger, caster: caster, _defaultValue: nil}

	for _, option := range options {
		option(middleware)
	}

	return middleware
}

func (middleware *param) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var value string

		for _, getter := range middleware.getters {
			value = getter(request)
			if strings.TrimSpace(value) != "" {
				break
			}
		}

		if value == "" && middleware._defaultValue == nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			middleware.logger.Errorf("middleware.param: name - '%s', not found", middleware.name)
			return
		}

		ctx := request.Context()

		if value != "" {
			castValue, err := middleware.caster(value)
			if err != nil {
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				middleware.logger.Error(err)
				return
			}

			ctx = context.WithValue(request.Context(), middleware.name, castValue)
		} else {
			ctx = context.WithValue(request.Context(), middleware.name, middleware._defaultValue)
		}

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
