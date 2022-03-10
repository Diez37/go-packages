package configurator

import (
	"github.com/spf13/viper"
	"strings"
)

// NewViper creating and configuration instance of viper.Viper implements Configurator
func NewViper() (Configurator, error) {
	vp := viper.New()

	vp.AddConfigPath(ConfigPathDefault)
	vp.SetConfigName(ConfigNameDefault)
	vp.SetConfigType(ConfigTypeDefault)

	vp.SetEnvKeyReplacer(strings.NewReplacer(EnvKeyReplaceFromDefault, EnvKeyReplaceToDefault))
	vp.AutomaticEnv()

	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	return vp, nil
}
