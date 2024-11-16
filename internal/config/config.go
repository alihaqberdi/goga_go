package config

import (
	"fmt"
	"github.com/go-shafaq/timep"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

var (
	PORT = "8080"

	//	Databases
	POSTGRES_URI = "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"

	// Chaching
	CACHING_EXPIRATION_DURATION = 16 * 24 * time.Hour
)

func LoadVarsFromEnv() {
	setIfExistsStr(&PORT, "PORT")

	setIfExistsStr(&POSTGRES_URI, "POSTGRES_URI")

	setIfExistsDur(&CACHING_EXPIRATION_DURATION, "CACHING_EXPIRATION_DURATION")

}

func setIfExists[V any](ptr *V, key string, parser func(string) (V, bool)) bool {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return false
	}

	val, ok := parser(envVal)
	if !ok {
		return false
	}

	*ptr = val

	return ok
}

func setIfExistsStr(ptr *string, key string) bool {
	return setIfExists(ptr, key,
		func(s string) (string, bool) { return s, true })
}

func setIfExistsBool(ptr *bool, key string) bool {
	return setIfExists(ptr, key,
		func(s string) (bool, bool) {
			b, err := strconv.ParseBool(s)
			return b, err == nil
		})
}

func setIfExistsDur(ptr *time.Duration, key string) bool {
	return setIfExists(ptr, key,
		func(s string) (time.Duration, bool) {
			dur, err := timep.ParseDuration(s)
			return dur, err == nil
		})
}

func init() {
	err := godotenv.Overload(".env")
	if err != nil {
		panic(fmt.Errorf("err in loading ENV: %s", err))
	}

	LoadVarsFromEnv()
}
