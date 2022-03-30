package v1

import (
	"github.com/husanmusa/uusd-uz/config"
	"github.com/husanmusa/uusd-uz/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type handlerV1 struct {
	db  *sqlx.DB
	log logger.Logger
	cfg config.Config
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Db     *sqlx.DB
	Logger logger.Logger
	Cfg    config.Config
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		db:  c.Db,
		log: c.Logger,
		cfg: c.Cfg,
	}
}
