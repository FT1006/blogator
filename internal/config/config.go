package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }
    return filepath.Join(homeDir, "project", "blogator", configFileName), nil
}

func Read() (Config, error) {
    path, err := getConfigFilePath()
    if err != nil {
        return Config{}, err
    }
    
    file, err := os.Open(path)
    if err != nil {
        return Config{}, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    var cfg Config
    err = decoder.Decode(&cfg)
    if err != nil {
        return Config{}, err
    }
    fmt.Println(cfg)
    return cfg, nil
}

func write(cfg Config) error {
    path, err := getConfigFilePath()
    if err != nil {
        return err
    }
    
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    err = encoder.Encode(cfg)
    fmt.Println(cfg)
    if err != nil {
        return err
    }

    return nil
}

func (cfg *Config) SetUser(userName string) error {
    cfg.CurrentUserName = userName
    return write(*cfg)
}