package main

import (
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"github.com/zignd/errors"
	"github.com/zignd/sculp-manga-here/common"
)

func init() {
	if err := common.ConfigViper(); err != nil {
		panic(errors.Wrap(err, "failed to set up viper"))
	}

	if err := common.ConfigLogrus(); err != nil {
		panic(errors.Wrap(err, "failed to set up Logrus"))
	}
}

func main() {
	common.Log.Infoln("Running...")
	if err := run(); err != nil {
		common.Log.Fatalln(errors.Wrap(err, "something wrong happened"))
	} else {
		common.Log.Fatalln("Application stopping for no apparent reason")
	}
}

func run() error {
	_, err := amqp.Dial(viper.GetString("RabbitMQ.URL"))
	if err != nil {
		return errors.Wrap(err, "failed to connect the RabbitMQ server")
	}

	return nil
}
