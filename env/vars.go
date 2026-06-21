package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func getEnvVar(key string) string {
	if envVar, ok := os.LookupEnv(key); ok {
		return envVar
	}
	env := ".env"
	if appEnv := os.Getenv("ENV"); appEnv != "" {
		env = env + "." + appEnv
	}
	godotenv.Load(env)

	envVar, ok := os.LookupEnv(key)
	if !ok {
		fmt.Println(fmt.Errorf("Environment variable not set %s", key))
	}

	return envVar
}
func GetApiPort() string {
	if uri := getEnvVar("PORT"); uri != "" {
		if !strings.Contains(uri, ":") {
			uri = fmt.Sprintf(":%s", uri)
		}
		return uri
	}

	return ":8025"
}

func GetJwtSecret() string {
	return getEnvVar("API_JWT_SECRET")

}
