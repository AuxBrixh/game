package main

import (
	"enet-server/server"
	"log/slog"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	conf, err := server.ReadToml(slog.Default())
	if err != nil {
		panic(err)
	}

	srv := conf.New()
	srv.Listen()
	srv.Wait()
}