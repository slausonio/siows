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
	AppName    = "APP_NAME"
	CurrentEnv = "ENV"
	Port       = "PORT"
)

var (
	ErrNoEnvFile    = errors.New("no .env file found in root of project")
	ErrNoAppName    = errors.New("no APP_NAME env var found")
	ErrNocurrentEnv = errors.New("no ENV env var found")
)

type Environment map[string]string

func (e Environment) Value(key string) string {
	return e[key]
}

func (e Environment) Update(key, value string) {
	e[key] = value
}

func NewEnvironment() Environment {
	env := make(Environment)
	env = env.readEnvironment()
	env.setEnvToSystem()

	return env
}

func (e Environment) readEnvironment() Environment {
	currentEnv := readCurrentEnv()
	defaultEnvMap := readDefaultEnvFile()
	currentEnvMap := readEnvironmentSpecificEnvFile(currentEnv)

	mergedEnv := siogo.MergeMaps(defaultEnvMap, currentEnvMap)

	return mergedEnv
}

func (e Environment) setEnvToSystem() {
	for key, value := range e {
		os.Setenv(key, value)
	}
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

func readCurrentEnv() string {
	appName, ok := os.LookupEnv(CurrentEnv)
	if !ok {
		err := fmt.Errorf("new environment: %w", ErrNocurrentEnv)

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
