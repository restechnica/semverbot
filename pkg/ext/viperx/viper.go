package viperx

import (
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// LoadConfig loads a configuration file.
// Returns an error if it fails.
func LoadConfig(path string) (err error) {
	viper.AddConfigPath(filepath.Dir(path))
	viper.SetConfigName(strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)))
	viper.SetConfigType(strings.Split(filepath.Ext(path), ".")[1])

	log.Debug().Str("path", path).Msg("loading config file...")

	viper.AutomaticEnv()

	return viper.ReadInConfig()
}
