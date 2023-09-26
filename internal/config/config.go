package config

import (
	"errors"
	"reflect"
	"strconv"
)

type Config struct {
	Port int `env:"PORT"`

	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresDB       string `env:"POSTGRES_DB"`
	DBHost           string `env:"DB_HOST"`
	DBPort           string `env:"DB_PORT"`

	DatabaseURL string `env:"DATABASE_URL"`

	TestPostgresUser     string `env:"TEST_POSTGRES_USER"`
	TestPostgresPassword string `env:"TEST_POSTGRES_PASSWORD"`
	TestPostgresDB       string `env:"TEST_POSTGRES_DB"`
	TestDBHost           string `env:"TEST_DB_HOST"`
	TestDBPort           string `env:"TEST_DB_PORT"`

	RedisHost      string `env:"REDIS_HOST"`
	RedisPort      string `env:"REDIS_PORT"`
	RedisIndexName string `env:"REDIS_INDEX_NAME"`

	Secret string `env:"SECRET"`
}

func NewFromMap(envMap map[string]string) (*Config, error) {
	c := Config{}
	ct := reflect.TypeOf(c)
	cv := reflect.ValueOf(&c).Elem()

	for i := 0; i < ct.NumField(); i++ {
		field := ct.Field(i)
		envFieldName, ok := field.Tag.Lookup("env")
		if !ok {
			continue
		}

		if _, ok := envMap[envFieldName]; !ok {
			continue
		}

		for k, e := range envMap {
			if k != envFieldName {
				continue
			}

			fieldValue := cv.Field(i)

			switch fieldValue.Kind() {
			case reflect.String:
				fieldValue.SetString(e)
			case reflect.Int:
				eInt, err := strconv.Atoi(e)
				if err != nil {
					return nil, err
				}
				fieldValue.SetInt(int64(eInt))
				break
			default:
				return &Config{}, errors.New("field should be a string or int")
			}
		}

	}

	return &c, nil
}
