package config

import (
	"fmt"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type env string

const (
	envName            = "APPENV"
	envDevelopment env = "development"
)

func New(repoName string) (*Config, error) {
	filename := getConfigFile(repoName)
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// getConfigFile get  config file name
// - files/etc/affiliate/affiliate.development.yaml in dev
// - otherwise /etc/affiliate/affiliate.{TKPENV}.yaml
func getConfigFile(repoName string) string {
	currentEnv := os.Getenv(envName)
	filename := fmt.Sprintf("%s.%s.yaml", repoName, currentEnv)
	// for non dev env, use config in /etc
	if env(currentEnv) != envDevelopment {
		return fmt.Sprintf("/etc/%s/%s", repoName, filename)
	}
	// use local files in dev
	repoPath := filepath.Join(os.Getenv("GOPATH"), "src/github.com/kecci/", repoName)
	return filepath.Join(repoPath, "files/etc", repoName, filename)
}
