package storage

import (
	"database/sql"
	"github.com/imperatorofdwelling/payment-svc/internal/config"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/postgres"
	"github.com/imperatorofdwelling/payment-svc/internal/storage/redis"
	"github.com/imperatorofdwelling/payment-svc/pkg"
	r "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Storage struct {
	Psql  *sql.DB
	Redis *r.Client
}

func GetStorages(cfg *config.Config, log *zap.SugaredLogger) *Storage {
	psqlStorage, err := postgres.NewPsqlStorage(cfg.Db.Postgres)
	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}

	if cfg.Env != pkg.ProdEnv {
		log.Infof("successfully connected to postgres with %s:%d", cfg.Db.Postgres.Host, cfg.Db.Postgres.Port)

	}

	redisClient, err := redis.NewRedisClient(cfg.Redis)
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}

	if cfg.Env != pkg.ProdEnv {
		log.Infof("successfully connected to redis with %s:%d", cfg.Redis.Host, cfg.Redis.Port)
	}

	return &Storage{
		Psql:  psqlStorage,
		Redis: redisClient,
	}
}
