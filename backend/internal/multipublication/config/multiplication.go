package config

import (
  "os"
)

// MultipublicationConfig is the configuration for multipublication.
type MultipublicationConfig struct {
  // Networks is a list of configured social networks.
  Networks []string
  // Tokens is a list of authorization tokens for the configured social networks.
  Tokens []string
}

// LoadMultipublicationConfig loads the multipublication configuration from a file.
func LoadMultipublicationConfig() (*MultipublicationConfig, error) {
  // Load the configuration from a file.
  config := &MultipublicationConfig{}
  err := os.ReadFile("./config/multipublication.json", config)
  if err != nil {
    return nil, err
  }

  // Validate the configuration.
  for _, network := range config.Networks {
    if network == "" {
      return nil, errors.New("social network cannot be empty")
    }
  }

  for _, token := range config.Tokens {
    if token == "" {
      return nil, errors.New("authorization token cannot be empty")
    }
  }

  return config, nil
}