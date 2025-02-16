package server

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/BurntSushi/toml"
)

type UserConfig struct {
	Server struct {
		Port uint16
		Name string
	}
	Players struct {
		Folder_player string
		Max_count uint
	}
}

type ServerConfig struct {
	Log *slog.Logger
	Name string
	MaxCount uint
	Listeners []func(conf ServerConfig) (Listener, error)
}

func(c ServerConfig) New() *Server {
	srv := &Server{conf: c}
	for _, lf := range c.Listeners {
		l, err := lf(c)
		if err != nil {
			c.Log.Error("create listener: " + err.Error())
		}
		srv.listeners = append(srv.listeners, l)
	}
	return srv
}

func (uc *UserConfig) Config(log *slog.Logger) (ServerConfig, error) {
	conf := ServerConfig{
		Log: log,
		Name: uc.Server.Name,
		MaxCount: uc.Players.Max_count,
	}

	conf.Listeners = append(conf.Listeners, uc.Listeners)
	return conf, nil
}

func DefaultConfig() UserConfig {
	Sc := UserConfig{}
	Sc.Server.Port = 8888
	Sc.Server.Name = "test"
	Sc.Players.Folder_player = "Players"
	Sc.Players.Max_count = 10
	return Sc
}

func ReadToml(log *slog.Logger) (ServerConfig, error) {
	c := DefaultConfig()
	var zero ServerConfig
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}
		return c.Config(log)
	}
	data, err := os.ReadFile("config.toml")
	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}
	return c.Config(log)
}