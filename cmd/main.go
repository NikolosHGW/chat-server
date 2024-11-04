package main

import (
	"context"
	"log"

	"github.com/NikolosHGW/chat-server/internal/app"
)

func main() {
	app, err := app.NewApp(context.Background())
	if err != nil {
		log.Fatalf("ошибка инициализации приложения: %s", err.Error())
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("ошибка запуска приложения: %s", err.Error())
	}
}
