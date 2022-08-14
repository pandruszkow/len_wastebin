package apiv1

import (
	chromaLexers "github.com/alecthomas/chroma/lexers"
	"github.com/coolguy1771/wastebin/internal/config"
	"github.com/coolguy1771/wastebin/internal/logger"
	"github.com/coolguy1771/wastebin/internal/storage"
)

type Data struct {
	Log logger.Config
	DB  storage.DB

	Lexers *[]string

	Version *string

	TitleMaxLen *int
	BodyMaxLen  *int
	MaxLifeTime *int64

	ServerAbout *string
	ServerRules *string

	AdminName *string
	AdminMail *string
}

func Load(cfg config.Config) Data {
	lexers := chromaLexers.Names(false)

	return Data{
		DB:          cfg.DB,
		Log:         cfg.Log,
		Lexers:      &lexers,
		Version:     &cfg.Version,
		TitleMaxLen: &cfg.TitleMaxLen,
		BodyMaxLen:  &cfg.BodyMaxLen,
		MaxLifeTime: &cfg.MaxLifeTime,
		ServerAbout: &cfg.ServerAbout,
		ServerRules: &cfg.ServerRules,
		AdminName:   &cfg.AdminName,
		AdminMail:   &cfg.AdminMail,
	}
}
