package config

import (
	"github.com/coolguy1771/wastebin/internal/logger"
	"github.com/coolguy1771/wastebin/internal/storage"
)

type Config struct {
	DB  storage.DB
	Log logger.Config

	Version string

	TitleMaxLen int
	BodyMaxLen  int
	MaxLifeTime int64

	ServerAbout string
	ServerRules string

	AdminName string
	AdminMail string

	RobotsDisallow bool
}
