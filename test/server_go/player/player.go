package player

import "github.com/google/uuid"

type PlayerData struct {
	uuid uuid.UUID
	username string
	name string
	x,y int
	gender uint8
}

func NewPlayer(uuid uuid.UUID, username string, gender uint8) *PlayerData {
	return &PlayerData{uuid: uuid, username: username, name: username, gender: gender, x: 0, y: 0}
}

func (p *PlayerData) GetUUID() uuid.UUID {
	return p.uuid
}

func (p *PlayerData) GetUsername() string {
	return p.username
}

func (p *PlayerData) GetName() string {
	return p.name
}

func (p *PlayerData) SetName(name string) {
	p.name = name
	return
}