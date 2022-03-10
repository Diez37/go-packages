package configurator

import "time"

// Configurator general interface for configurator a application
type Configurator interface {
	GetString(string) string
	GetBool(string) bool
	GetInt(string) int
	GetInt32(string) int32
	GetInt64(string) int64
	GetUint(string) uint
	GetUint32(string) uint32
	GetUint64(string) uint64
	GetFloat64(string) float64
	GetTime(string) time.Time
	GetDuration(string) time.Duration
	GetIntSlice(string) []int
	GetStringSlice(string) []string
	GetStringMap(string) map[string]interface{}
	GetStringMapString(string) map[string]string
	GetStringMapStringSlice(string) map[string][]string
	GetSizeInBytes(string) uint

	SetDefault(string, interface{})
}
