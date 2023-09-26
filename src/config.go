package config

import (
    "os"
)

type GitHubConfig struct {
    Username string
    APIKey string
}

type Config struct {
    GitHub GitHubConfig
}

func New() *Config {
    return &Config{
        GitHub: GitHubConfig{
            Username: getEnv("GITHUB_USERNAME", "")
            APIKey: getEnv("GITHUB_API_KEY", "")
        },
    }
}

func getEnv(key string, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }

    return defaultVal
}

