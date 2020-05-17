package config

import (
	"encoding/json"
	"os"
	"os/user"
	"strings"
)

const (
	configPath     = "~/.config/tils-cli.json"
	defaultBaseURL = "https://tils.dev/api/"
)

type Config struct {
	APIToken string `json:"api_token"`
	BaseURL  string `json:"-"`
}

func Load() (*Config, error) {
	config := &Config{
		APIToken: os.Getenv("TILS_CLI_API_TOKEN"),
		BaseURL:  os.Getenv("TILS_CLI_API_BASE_URL"),
	}
	if config.BaseURL == "" {
		config.BaseURL = defaultBaseURL
	}

	cfgPath, err := absoluteConfigPath()
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(cfgPath)
	if err != nil {
		if os.IsNotExist(err) {
			return config, nil
		}
		return nil, err
	}

	file, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileConfig := &Config{}
	err = json.NewDecoder(file).Decode(fileConfig)
	if err != nil {
		return nil, err
	}

	if config.APIToken == "" && fileConfig.APIToken != "" {
		config.APIToken = fileConfig.APIToken
	}

	return config, nil
}

func Write(apiToken string) error {
	config := &Config{
		APIToken: apiToken,
	}

	cfgPath, err := absoluteConfigPath()
	if err != nil {
		return err
	}

	file, err := os.Create(cfgPath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(config)
	if err != nil {
		return err
	}

	return nil
}

func absoluteConfigPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return strings.Replace(configPath, "~", usr.HomeDir, 1), nil
}
