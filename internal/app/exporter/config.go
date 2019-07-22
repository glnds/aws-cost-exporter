package exporter

import (
	"log"
	"os"
	"os/user"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

// Config represents the application's config file
type Config struct {
	AccessKeyID     string `toml:"AccessKeyID"`
	SecretAccessKey string `toml:"SecretAccessKey"`
	Debug           bool   `toml:"Debug"`
}

// GetConfig reads the .dyndns-r53.toml configuration file for initialization.
func GetConfig(logger *logrus.Logger) Config {

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// Read dyndns-r53.toml config file for initialization
	conf := Config{Debug: false} // Set default values
	if _, err := toml.DecodeFile(usr.HomeDir+string(os.PathSeparator)+".aws-cost-exporter.toml", &conf); err != nil {
		log.Fatal(err.Error())
	}

	logger.WithFields(logrus.Fields{
		"AccessKeyID":     conf.AccessKeyID,
		"SecretAccessKey": conf.SecretAccessKey,
		"Debug":           conf.Debug,
	}).Info("Config settings")

	return conf
}
