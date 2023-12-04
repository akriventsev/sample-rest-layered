package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/akriventsev/sample-rest-layered/source/internal/app/application"
	"github.com/akriventsev/sample-rest-layered/source/internal/common"
	"github.com/akriventsev/sample-rest-layered/source/internal/storage"
	"github.com/akriventsev/sample-rest-layered/source/internal/transport/http"
	"github.com/ilyakaznacheev/cleanenv"
)

func setLogger() {
	var l *slog.Logger

	switch common.MainConfig.LogFormat {
	case "text":
		l = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: common.MainConfig.LogLevel,
		}))
	default:
		l = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: common.MainConfig.LogLevel,
		}))
	}

	slog.SetDefault(l)
	slog.Info("start logger", slog.Any("format", common.MainConfig.LogFormat), slog.String("level", common.MainConfig.LogLevel.String()))
}

func main() {
	err := cleanenv.ReadEnv(&common.MainConfig)
	if err != nil {
		slog.Error("cannot load config", slog.Any("error", err))
	}

	setLogger()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		common.MainConfig.DBHost,
		common.MainConfig.DBUser,
		common.MainConfig.DBPassword,
		common.MainConfig.DBName,
		common.MainConfig.DBPort,
	)

	slog.Info("started")

	s, err := storage.NewGormPgStorage(dsn)

	if err != nil {
		slog.Error("cannot initialize app", "error", err)

		return
	}

	app, err := application.NewApplication(s)

	if err != nil {
		return
	}

	tr, err := http.NewTransport(app, http.WithConfig(http.Config{
		ListenAddress: common.MainConfig.ListenAddr,
		Port:          common.MainConfig.ListenPort,
		JwtSecret:     "secret",
	}))

	if err != nil {
		return
	}

	err = tr.Start(context.Background())

	if err != nil {
		return
	}
}
