/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-12-14 04:23:27
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-14 05:13:34
 * @FilePath: /lesson3/config/config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		Server  ServerConfig
		MongoDB MongoDBConfig
	}
	ServerConfig struct {
		Host string
		Port string // Server port
	}

	MongoDBConfig struct {
		URI string
	}
)

func (c *Config) Load() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	requiredVars := map[string]*string{
		"SERVER_HOST": &c.Server.Host,
		"SERVER_PORT": &c.Server.Port,
		//database
		"MONGODB_URI": &c.MongoDB.URI,
	}

	for envVar, fieldPtr := range requiredVars {
		value := os.Getenv(envVar)
		if value == "" {
			return fmt.Errorf("missing required environment variable: %s", envVar)
		}
		*fieldPtr = value
	}

	return nil
}

func New() (*Config, error) {
	var config Config
	if err := config.Load(); err != nil {
		return nil, err
	}
	return &config, nil
}
