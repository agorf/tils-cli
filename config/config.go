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
	BaseURL  string `json:"base_url"`
}

func Load() (*Config, error) {
	config := &Config{
		APIToken: os.Getenv("TILS_CLI_API_TOKEN"),
		BaseURL:  os.Getenv("TILS_CLI_API_BASE_URL"),
	}
	if config.BaseURL == "" {
		config.BaseURL = defaultBaseURL
	}

	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	absoluteConfigPath := strings.Replace(configPath, "~", usr.HomeDir, 1)

	_, err = os.Stat(absoluteConfigPath)
	if err != nil {
		if os.IsNotExist(err) {
			return config, nil
		}
		return nil, err
	}

	file, err := os.Open(absoluteConfigPath)
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
	if config.BaseURL == defaultBaseURL && fileConfig.BaseURL != "" {
		config.BaseURL = fileConfig.BaseURL
	}

	return config, nil
}
