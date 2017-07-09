package common

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zignd/errors"
)

// Log is a pre-configured Logrus Entry that should be used for logging throughout the application.
var Log *logrus.Entry

// ConfigLogrus performs the initial Logrus set up.
func ConfigLogrus() error {
	rabbitMQURL := viper.GetString("RabbitMQ.URL")
	if rabbitMQURL == "" {
		return errors.New("RabbitMQ.URL is not set")
	}

	Log = logrus.WithField("application", "sculp-manga-here")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.SetOutput(os.Stdout)

	return nil
}
