package main

import (
	"enet-server/player"
	"enet-server/player/db"
	"fmt"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/google/uuid"
)


func main() {
	// db, err := leveldb.OpenFile("data", nil);
	// if err != nil {
	// 	panic(err)
	// } 

	// defer db.Close()

	// value,err := db.Get([]byte("key"), nil)
	// fmt.Println(string(value))
	// err = db.Delete([]byte("key"), nil)
	// value,err = db.Get([]byte("key"), nil)
	// fmt.Println(string(value))

	player := player.Config{
		Uuid: uuid.New(),
		Username: "hello",
		Name: "hello",
		Position: mgl64.Vec2{},
		Speed: 5,
		Health: 20,
		Maxhealth: 20,
	}

	db, err := db.NewPlayerDB("test")
	if err != nil {
		panic(err)
	}

	err = db.Save("hello", player);
	if err != nil {
		panic(err)
	}

	conf, err := db.Load("asa")
	if err != nil {
		panic(err)
	}

	fmt.Print(conf)
	return
} 