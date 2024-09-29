package configs

import (
	"context"
	"distributed-task-queue/pkg/utils"
	"fmt"
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
)

type config struct {
	Api    apiserver    `yaml:"api_server"`
	Worker workerserver `yaml:"worker_server"`
	Redis  redisserver  `yaml:"redis_server"`
	Docker dockerserver `yaml:"docker_server"`
}

type apiserver struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type workerserver struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type redisserver struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type dockerserver struct {
}

var (
	once sync.Once
	rdb  *redis.Client
)

var Env = func() config {
	data, err := os.ReadFile("pkg/configs/config.yaml")
	utils.LogFatalError(err, "Error occured")

	var cfg config
	err = yaml.Unmarshal(data, &cfg)
	utils.LogFatalError(err, "Error occured")

	return cfg
}()

func InitRedis() (*redis.Client, error) {
	var err error
	once.Do(func() {
		ctx := context.Background()

		rdb = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", Env.Redis.Host, Env.Redis.Port),
			Password: Env.Redis.Password,
			DB:       Env.Redis.DB,
		})

		_, err = rdb.Ping(ctx).Result()
	})

	if err != nil {
		return nil, err
	}

	return rdb, nil
}
