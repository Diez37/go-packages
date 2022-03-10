package bind_flags

import (
	"fmt"
	"github.com/diez37/go-packages/clients/cache"
	"github.com/diez37/go-packages/clients/cache/gocache"
	"github.com/diez37/go-packages/clients/cache/redis"
	"github.com/diez37/go-packages/clients/db"
	"github.com/diez37/go-packages/clients/db/mysql"
	"github.com/diez37/go-packages/clients/db/sqlite"
	httpClient "github.com/diez37/go-packages/clients/http"
	"github.com/diez37/go-packages/container"
	"github.com/diez37/go-packages/log"
	"github.com/diez37/go-packages/migrator"
	"github.com/diez37/go-packages/server/http"
	"github.com/diez37/go-packages/tracer"
	"github.com/spf13/cobra"
	"strings"
)

func CobraCmd(container container.Container, cmd *cobra.Command, components ...Component) (*cobra.Command, error) {
	for _, component := range components {
		switch component {
		case Cache:
			err := container.Invoke(func(cacheConfig *cache.Config, goCacheConfig *gocache.Config, redisConfig *redis.Config) {
				cmd.PersistentFlags().StringVar(&redisConfig.Host, redis.HostFieldName, redis.HostDefault, "")
				cmd.PersistentFlags().UintVar(&redisConfig.Port, redis.PortFieldName, redis.PortDefault, "")
				cmd.PersistentFlags().StringVar(&redisConfig.User, redis.UserFieldName, "", "")
				cmd.PersistentFlags().StringVar(&redisConfig.Password, redis.PasswordFieldName, "", "")
				cmd.PersistentFlags().DurationVar(&redisConfig.ReadTimeout, redis.ReadTimeoutFieldName, redis.ReadTimeoutDefault, "")
				cmd.PersistentFlags().DurationVar(&redisConfig.WriteTimeout, redis.WriteTimeoutFieldName, redis.WriteTimeoutDefault, "")
				cmd.PersistentFlags().DurationVar(&redisConfig.ConnectTimeout, redis.ConnectTimeoutFieldName, redis.ConnectTimeoutDefault, "")
				cmd.PersistentFlags().DurationVar(&goCacheConfig.Expiration, gocache.ExpirationFieldName, gocache.ExpirationDefault, "")
				cmd.PersistentFlags().DurationVar(&goCacheConfig.CleanupInterval, gocache.CleanupIntervalFieldName, gocache.CleanupIntervalDefault, "")
				cmd.PersistentFlags().StringVar(&cacheConfig.Type, cache.TypeFieldName, cache.TypeDefault, fmt.Sprintf(
					"cache type, available values (%s)",
					strings.Join([]string{cache.GoCacheType, cache.RedisType}, ", "),
				))
			})
			if err != nil {
				return nil, err
			}
		case CacheOnlyRedis:
			err := container.Invoke(func(cacheConfig *cache.Config, redisConfig *redis.Config) {
				cmd.PersistentFlags().StringVar(&redisConfig.Host, redis.HostFieldName, redis.HostDefault, "")
				cmd.PersistentFlags().UintVar(&redisConfig.Port, redis.PortFieldName, redis.PortDefault, "")
				cmd.PersistentFlags().StringVar(&redisConfig.User, redis.UserFieldName, "", "")
				cmd.PersistentFlags().StringVar(&redisConfig.Password, redis.PasswordFieldName, "", "")
				cmd.PersistentFlags().DurationVar(&redisConfig.ReadTimeout, redis.ReadTimeoutFieldName, redis.ReadTimeoutDefault, "")
				cmd.PersistentFlags().DurationVar(&redisConfig.WriteTimeout, redis.WriteTimeoutFieldName, redis.WriteTimeoutDefault, "")
				cmd.PersistentFlags().DurationVar(&redisConfig.ConnectTimeout, redis.ConnectTimeoutFieldName, redis.ConnectTimeoutDefault, "")
				cacheConfig.Type = cache.RedisType
				cacheConfig.OnlyType = cache.RedisType
			})
			if err != nil {
				return nil, err
			}
		case CacheOnlyGoCache:
			err := container.Invoke(func(cacheConfig *cache.Config, goCacheConfig *gocache.Config) {
				cmd.PersistentFlags().DurationVar(&goCacheConfig.Expiration, gocache.ExpirationFieldName, gocache.ExpirationDefault, "")
				cmd.PersistentFlags().DurationVar(&goCacheConfig.CleanupInterval, gocache.CleanupIntervalFieldName, gocache.CleanupIntervalDefault, "")
				cacheConfig.Type = cache.GoCacheType
				cacheConfig.OnlyType = cache.GoCacheType
			})
			if err != nil {
				return nil, err
			}
		case DataBase:
			err := container.Invoke(func(mysqlConfig *mysql.Config, sqliteConfig *sqlite.Config, dbConfig *db.Config) {
				cmd.PersistentFlags().StringVar(&mysqlConfig.Host, mysql.HostFieldName, mysql.HostDefault, "")
				cmd.PersistentFlags().Uint32Var(&mysqlConfig.Port, mysql.PortFieldName, mysql.PortDefault, "")
				cmd.PersistentFlags().StringVar(&mysqlConfig.User, mysql.UserFieldName, "", "")
				cmd.PersistentFlags().StringVar(&mysqlConfig.Password, mysql.PasswordFieldName, "", "")
				cmd.PersistentFlags().StringVar(&mysqlConfig.DataBase, mysql.DataBaseFieldName, "", "")

				cmd.PersistentFlags().StringVar(&sqliteConfig.Dsn, sqlite.DsnFieldName, "", "")

				cmd.PersistentFlags().StringVar(&dbConfig.Driver, db.DriverFieldName, "", fmt.Sprintf(
					"type db usage, available values (%s)",
					strings.Join([]string{db.MySQLDriver, db.SQLiteDriver}, ", "),
				))
			})
			if err != nil {
				return nil, err
			}
		case MySQL:
			err := container.Invoke(func(mysqlConfig *mysql.Config, dbConfig *db.Config) {
				cmd.PersistentFlags().StringVar(&mysqlConfig.Host, mysql.HostFieldName, mysql.HostDefault, "")
				cmd.PersistentFlags().Uint32Var(&mysqlConfig.Port, mysql.PortFieldName, mysql.PortDefault, "")
				cmd.PersistentFlags().StringVar(&mysqlConfig.User, mysql.UserFieldName, "", "")
				cmd.PersistentFlags().StringVar(&mysqlConfig.Password, mysql.PasswordFieldName, "", "")
				cmd.PersistentFlags().StringVar(&mysqlConfig.DataBase, mysql.DataBaseFieldName, "", "")
				dbConfig.Driver = db.MySQLDriver
			})
			if err != nil {
				return nil, err
			}
		case SQLite:
			err := container.Invoke(func(sqliteConfig *sqlite.Config, dbConfig *db.Config) {
				cmd.PersistentFlags().StringVar(&sqliteConfig.Dsn, sqlite.DsnFieldName, "", "")
				dbConfig.Driver = db.SQLiteDriver
			})
			if err != nil {
				return nil, err
			}
		case Logger:
			err := container.Invoke(func(logConfig *log.Config) {
				cmd.PersistentFlags().BoolVarP(&logConfig.InfoLevel, log.InfoLevelFieldName, "v", log.VerboseDefault, "log info level")
				cmd.PersistentFlags().BoolVarP(&logConfig.DebugLevel, log.DebugLevelFieldName, "d", log.VerboseDefault, "log debug level")
				cmd.PersistentFlags().StringVar(&logConfig.SentryDSN, log.SentryDSNFieldName, "", "sentry dsn")
			})
			if err != nil {
				return nil, err
			}
		case HttpServer:
			err := container.Invoke(func(httpServerConfig *http.Config) {
				cmd.PersistentFlags().StringVar(&httpServerConfig.Interface, http.InterfaceFieldName, http.InterfaceDefault, "http server listener interface")
				cmd.PersistentFlags().UintVar(&httpServerConfig.Port, http.PortFieldName, http.PortDefault, "http server listener port")
				cmd.PersistentFlags().DurationVar(&httpServerConfig.ShutdownTimeout, http.ShutdownTimeoutFieldName, http.ShutdownTimeoutDefault, "http server timeout for shutdown")
			})
			if err != nil {
				return nil, err
			}
		case HttpClient:
			err := container.Invoke(func(httpClientConfig *httpClient.Config) {
				cmd.PersistentFlags().DurationVar(&httpClientConfig.Timeout, httpClient.TimeoutFieldName, httpClient.TimeoutDefault, "")
				cmd.PersistentFlags().IntVar(&httpClientConfig.MaxIdleConns, httpClient.MaxIdleConnsFieldName, httpClient.MaxIdleConnsDefault, "")
				cmd.PersistentFlags().IntVar(&httpClientConfig.MaxIdleConnsPerHost, httpClient.MaxIdleConnsPerHostFieldName, httpClient.MaxIdleConnsPerHostDefault, "")
				cmd.PersistentFlags().IntVar(&httpClientConfig.MaxConnsPerHost, httpClient.MaxConnsPerHostFieldName, httpClient.MaxConnsPerHostDefault, "")
			})
			if err != nil {
				return nil, err
			}
		case Tracer:
			err := container.Invoke(func(traceConfig *tracer.Config) {
				cmd.PersistentFlags().BoolVar(&traceConfig.UseProfile, tracer.UseProfileFieldName, tracer.UseProfileDefault, "adding '/debug/pprof/...' handlers")
				cmd.PersistentFlags().StringVar(&traceConfig.JaegerDSN, tracer.JaegerDSNFieldName, "", "jaeger dsn")
			})
			if err != nil {
				return nil, err
			}
		case Migrator:
			err := container.Invoke(func(migratorConfig *migrator.Config) {
				cmd.PersistentFlags().StringVar(&migratorConfig.Source, migrator.SourceFieldName, migrator.SourceDefault, "directory of migrations")
			})
			if err != nil {
				return nil, err
			}
		}
	}

	return cmd, nil
}
