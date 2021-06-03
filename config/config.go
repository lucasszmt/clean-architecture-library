package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type envConfig struct {
	DbName     string
	DbHost     string
	DbUser     string
	DbPassword string
	DbPort     string
}

func (e envConfig) String() string {
	return fmt.Sprintf("[DbName: %s, DbHost: %s, DbUser: %s, DbPassword: %s, DbPort: %s]",
		e.DbName, e.DbHost, e.DbUser, e.DbPassword, e.DbPort)
}

var Envs *envConfig
var ErrDb = errors.New("error loading database")

func Init(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil {
		return err
	}
	Envs = &envConfig{
		DbName:     os.Getenv("DB_DB"),
		DbHost:     os.Getenv("DB_HOST"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbPort:     os.Getenv("DB_PORT"),
	}
	if Envs.DbName == "" || Envs.DbHost == "" || Envs.DbUser == "" || Envs.DbPassword == "" || Envs.DbPort == "" {
		return ErrDb
	}
	return nil
}
