package models

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/youssefsiam38/twist/reporter"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Execute    string   `yaml:"execute"`
	Order      []string `yaml:"order"`
	Timeout    string   `yaml:"timeout"`
	Output     string   `yaml:"output"`
	TimeoutDur time.Duration
}

func ParseConfig() (*Config, error) {
	config := Config{}

	configFile, err := ioutil.ReadFile("./twist/config.yml")
	if err != nil {

		// check if any fail in reading the config file Twist will use the default values
		if err.Error() == "open ./twist/config.yml: no such file or directory" {
			// return &Config{
			// 	Execute:    "in order",
			// 	Order:      []string{},
			// 	Timeout:    "",
			// 	TimeoutDur: 2 * time.Minute,
			// }, nil
			reporter.Write("You must add a twist/config.yaml file and specify an order", reporter.FLOW_ERROR)
			os.Exit(1)
		} else {
			return nil, err
		}
	}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}

	config.TimeoutDur, err = time.ParseDuration(config.Timeout)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
