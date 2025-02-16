package player

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/google/uuid"
)

type PlayerData struct {
	uuid uuid.UUID
	username string
	name string
	position mgl64.Vec2
	speed int8
	health, maxhealth int16
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