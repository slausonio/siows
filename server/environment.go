package server

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/slausonio/siogo"
)

const (
	AppName = "APP_NAME"
	Env     = "ENV"
	Port    = "PORT"
)

var (
	ErrNoEnvFile = errors.New("no .env file found in root of project")
	ErrNoAppName = errors.New("no APP_NAME env var found")
	ErrNoEnv     = errors.New("no ENV env var found")
)

type Environment map[string]string

func (e Environment) Value(key string) string {
	return e[key]
}

func (e Environment) Update(key, value string) {
	e[key] = value
}

func NewEnvironment() Environment {
	environment := make(Environment)

	return environment
}

func (e Environment) readEnvironment() Environment {
	defaultEnvFile := readDefaultEnvFile()
	mergedEnv := siogo.MergeMaps(defaultEnvFile, e)

	environment := make(Environment)
	environment["PATH"] = "/bin:/usr/bin"
	return environment
}

func readDefaultEnvFile() Environment {
	defaultEnvFile, err := godotenv.Read()
	if err != nil {
		dotEnvErr := fmt.Errorf("dot env err: %w", err)
		logrus.Error(dotEnvErr)

		panic(ErrNoEnvFile)
	}

	return defaultEnvFile
}

func readEnvironmentSpecificEnvFile(env string) Environment {
	fileName := fmt.Sprintf("env/-%s", env)
	defaultEnvFile, err := godotenv.Read(fileName)
	if err != nil {
		dotEnvErr := fmt.Errorf("dot env err: %w", err)
		logrus.Info(dotEnvErr)
	}

	return defaultEnvFile
}

func readEnv() string {
	appName, ok := os.LookupEnv(Env)
	if !ok {
		err := fmt.Errorf("new environment: %w", ErrNoEnv)

		logrus.Error(err)
		panic(err)
	}

	return appName
}

func readAppName() string {
	appName, ok := os.LookupEnv(AppName)
	if !ok {
		err := fmt.Errorf("new environment: %w", ErrNoAppName)

		logrus.Error(err)
		panic(err)
	}

	return appName
}
