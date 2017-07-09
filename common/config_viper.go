package common

import (
	"os"
	"path"

	"github.com/spf13/viper"
	"github.com/zignd/errors"
)

// ConfigViper performs the initial viper set up.
func ConfigViper() error {
	// This application conforms with the XDG Base Directory Specification
	// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html

	viper.SetConfigName("config")
	if xdgConfigHome := os.Getenv("XDG_CONFIG_HOME"); xdgConfigHome != "" {
		viper.AddConfigPath(path.Join(xdgConfigHome, "sculp-manga-here"))
	} else {
		viper.AddConfigPath("$HOME/.config/sculp-manga-here")
	}

	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrapcf(err, map[string]interface{}{
			"XDG_CONFIG_HOME": os.Getenv("XDG_CONFIG_HOME"),
		}, "failed to read the configuration file")
	}

	return nil
}
