package db

import (
	"encoding/json"
	"enet-server/player"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type PlayerDB struct {
	db *leveldb.DB
}

func NewPlayerDB(path string) (*PlayerDB, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, 0777)
	}

	db, err := leveldb.OpenFile(path, &opt.Options{Compression: opt.SnappyCompression})
	if err != nil {
		return nil, err
	}

	return &PlayerDB{db: db}, nil
}


func(p *PlayerDB) Save(username string, data player.Config) error {
	d, err := json.Marshal(data);
	if err != nil {
		return err
	}
	return p.db.Put([]byte(username), d, nil)
}

func(p *PlayerDB) Load(username string) (player.Config, error){
	b, err := p.db.Get([]byte(username), nil);
	if err != nil {
		return player.Config{}, err
	}

	var d player.Config
	err = json.Unmarshal(b, &d)
	if err != nil {
		return player.Config{}, err
	}

	return d, nil
}