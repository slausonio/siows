package environment

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
	CurrentEnv = "CURRENT_ENV"
	Port       = "PORT"

	DefaultFilePath    = "env/.env"
	CurrentEnvFilePath = "env/%s.env"
)

var (
	ErrNoEnvFile    = errors.New("no .env file found in root of project")
	ErrNoAppName    = errors.New("no APP_NAME env var found")
	ErrNoCurrentEnv = errors.New("no CURRENT_ENV env var found")
)

type SioGoEnv map[string]string

func (e SioGoEnv) Value(key string) string {
	return e[key]
}

func (e SioGoEnv) Update(key, value string) {
	e[key] = value
}

func NewEnvironment() SioGoEnv {
	env := make(SioGoEnv)
	env = env.readEnvironment()
	env.setEnvToSystem()

	return env
}

func (e SioGoEnv) readEnvironment() SioGoEnv {
	defaultEnvMap := readDefaultEnvFile()
	defaultEnvMap.setEnvToSystem()

	currentEnv := readCurrentEnv()
	currentEnvMap := readEnvironmentSpecificFile(currentEnv)

	mergedEnv := siogo.MergeMaps(defaultEnvMap, currentEnvMap)

	return mergedEnv
}

func (e SioGoEnv) setEnvToSystem() {
	for key, value := range e {
		err := os.Setenv(key, value)
		if err != nil {
			panic(err)
		}
	}
}

func readDefaultEnvFile() SioGoEnv {
	defaultEnvFile, err := godotenv.Read(DefaultFilePath)
	if err != nil {
		dotEnvErr := fmt.Errorf("dot env err: %w", err)

		logrus.Error(dotEnvErr)
		panic(ErrNoEnvFile)
	}

	return defaultEnvFile
}

func readEnvironmentSpecificFile(env string) SioGoEnv {
	fileName := fmt.Sprintf(CurrentEnvFilePath, env)

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
		err := fmt.Errorf("new environment: %w", ErrNoCurrentEnv)

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
