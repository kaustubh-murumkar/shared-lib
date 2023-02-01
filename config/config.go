package config

import (
	"errors"
	"os"

	"github.com/tidwall/gjson"
)

type ConfigFileType int

const (
	JSON ConfigFileType = iota
	TOML                = iota
	YAML                = iota
)

type Config struct {
	config gjson.Result
}

// errors
var (
	ErrInvalidFileType = errors.New("config: invalid file type")
)

var (
	osReadFile = os.ReadFile
)

// Init - read a config file optionally decrypt it and construct config struct
func Init(configFilePath string, fileType ConfigFileType, isEncrypted bool) (Config, error) {
	if fileType > YAML || fileType < JSON {
		return Config{}, ErrInvalidFileType
	}

	content, err := osReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}

	if isEncrypted {
		// remove when implemented
		panic("encrypted config is not implemented yet!")
	}

	switch fileType {
	case JSON:
		{
			config := Config{
				config: gjson.ParseBytes(content),
			}
			return config, nil
		}
	default:
		panic("file type not implemented yet!")
	}
}

// GetString - get value from config and type cast to string
func (c Config) GetString(path string) string {
	return c.config.Get(path).String()
}

// GetInt - get value from config and type cast to int
func (c Config) GetInt(path string) int64 {
	return c.config.Get(path).Int()
}

// GetBool - get value from config and type cast to boolean
func (c Config) GetBool(path string) bool {
	return c.config.Get(path).Bool()
}

// Get - get value from config
func (c Config) Get(path string) gjson.Result {
	return c.config.Get(path)
}
