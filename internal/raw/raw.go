package raw

import (
	"github.com/coolguy1771/wastebin/internal/config"
	"github.com/coolguy1771/wastebin/internal/logger"
	"github.com/coolguy1771/wastebin/internal/storage"
)

type Data struct {
	DB  storage.DB
	Log logger.Config
}

func Load(cfg config.Config) Data {
	return Data{
		DB:  cfg.DB,
		Log: cfg.Log,
	}
}
