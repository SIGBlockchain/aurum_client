package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/SIGBlockchain/aurum_client/internal/constants"
)

type Config struct {
	Version         uint16
	ProducerAddress string
}

func LoadConfiguration() (*Config, error) {
	configFile, err := os.Open(constants.ConfigurationFile)
	if err != nil {
		return nil, errors.New("Failed to load configuration file : " + err.Error())
	}
	defer configFile.Close()

	cfgData, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, errors.New("Failed to read configuration file : " + err.Error())
	}

	cfg := Config{}
	if err := json.Unmarshal(cfgData, &cfg); err != nil {
		return nil, errors.New("Failed to unmarshall configuration data : " + err.Error())
	}

	return &cfg, nil
}

//SaveConfiguration saves a configuration to file
func SaveConfiguration(cfg *Config) error {
	cfgData, err := json.Marshal(&cfg)
	if err != nil {
		return errors.New("Failed to marhsal the config data structure")
	}
	if ioutil.WriteFile(constants.ConfigurationFile, cfgData, 0644) != nil {
		return errors.New("Failed to write configuration to file")
	}
	return nil
}
